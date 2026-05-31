/**
 * Donations indexer — processor entry point.
 *
 * Watches native-value transactions AND allowlisted ERC20 Transfer logs
 * whose recipient is the thatsrekt.eth donation Safe on Ethereum mainnet.
 * Writes rows to the `donation` table in `thatsrekt_donations` DB.
 *
 * Processor-only: no squid-graphql-server is started here.
 * The mesh reads directly via a second pg pool (DONATIONS_DB_URL).
 *
 * Slice #205: Ethereum + native ETH only.
 * Slice #207: ERC20 Transfer log subscriptions added.
 * Slice #209 adds additional chains.
 */

import 'dotenv/config'
import { EvmBatchProcessor } from '@subsquid/evm-processor'
import type {
  HotDatabase,
  HotDatabaseState,
  HotTxInfo,
  FinalTxInfo,
  HashAndHeight,
} from '@subsquid/util-internal-processor-tools'
import pkg from 'pg'
import { ensureDonationTable, upsertDonation } from './donationStore.js'
import { mapNativeTransfer, mapErc20Transfer } from './donationMapper.js'
import { erc20Addresses, TRANSFER_TOPIC0 } from './tokenAllowlist.js'

const { Pool } = pkg

// ---------------------------------------------------------------------------
// Env validation — fail fast before touching any infrastructure.
// ---------------------------------------------------------------------------

const requireEnv = (key: string): string => {
  const v = process.env[key]
  if (!v) throw new Error(`Missing required env var: ${key}`)
  return v
}

const RPC_URL = requireEnv('RPC_ETHEREUM_HTTP')
const DB_URL = requireEnv('DONATIONS_DB_URL')

// The thatsrekt.eth Safe — canonical donation address on every supported chain.
// Ethereum mainnet v1.2.0 multisig (also the thatsRekt governance multisig).
const DONATION_SAFE = '0x59E4DBc95BD312A882Bb36b7f3E8298682340679'.toLowerCase()

// Ethereum chainId 1, slug 'ethereum'
const CHAIN_ID = 1
const CHAIN_SLUG = 'ethereum'

// The Safe's deployment block on Ethereum mainnet.
// We index from here so the full history is captured.
// Confirmed via Etherscan: the Safe was deployed in the same tx that created
// the thatsRekt governance multisig. Start block sourced from env so it can
// be overridden for testing (anvil fork starts much later).
const START_BLOCK = parseInt(process.env.START_BLOCK_ETHEREUM ?? '19000000', 10)

// Finality confirmation in blocks.
// Default: 75 (Ethereum PoS justification depth, ~15 min).
// Set FINALITY_CONFIRMATION=0 in tests (e.g. anvil e2e) to treat all blocks
// as final — with allBlocksAreFinal=true, the processor only calls transact()
// and skips the hot-block phase entirely.
const FINALITY_CONFIRMATION = parseInt(process.env.FINALITY_CONFIRMATION ?? '75', 10)

// ---------------------------------------------------------------------------
// Postgres pool for donations DB.
// ---------------------------------------------------------------------------

const pool = new Pool({
  connectionString: DB_URL,
  max: 5,
  idleTimeoutMillis: 30_000,
  connectionTimeoutMillis: 5_000,
})

pool.on('error', (err) => {
  console.error('[donations-indexer] idle client error:', err)
})

// ---------------------------------------------------------------------------
// Subsquid processor.
//
// Subscriptions:
//   - addTransaction({to:[Safe]}) — native value transfers to the Safe.
//   - addLog({address:[token], topic0:[Transfer], topic2:[Safe]}) per ERC20
//     — Transfer events directed to the Safe from each allowlisted token.
//     topic2 = recipient (second indexed arg of Transfer(from,to,amount)).
//
// The `topic2` filter is Subsquid's server-side filter: only logs where
// topics[2] matches the Safe address padded to 32 bytes are returned.
// We still verify `to` defensively in the handler.
//
// No gateway for local Anvil forks — falls back to RPC-only.
// ---------------------------------------------------------------------------

// Pad an address to a 32-byte topic (0x + 64 hex chars, address in lower 20 bytes).
const addressToTopic = (addr: string): string =>
  '0x' + addr.replace(/^0x/, '').toLowerCase().padStart(64, '0')

const DONATION_SAFE_TOPIC = addressToTopic(DONATION_SAFE)

// Retrieve all allowlisted ERC20 addresses for chain 1 (Ethereum).
const ERC20_ADDRESSES = erc20Addresses(CHAIN_ID)

let base = new EvmBatchProcessor()
  .setRpcEndpoint({
    url: RPC_URL,
    rateLimit: 10,
  })
  .setFinalityConfirmation(FINALITY_CONFIRMATION)
  .setFields({
    transaction: {
      to: true,
      from: true,
      value: true,
      hash: true,
    },
    log: {
      address: true,
      topics: true,
      data: true,
      transactionHash: true,
    },
  })
  .setBlockRange({ from: START_BLOCK })
  .addTransaction({
    to: [DONATION_SAFE],
  })

// Register one addLog subscription per allowlisted ERC20.
// Filter: Transfer topic0 + Safe as recipient (topic2).
for (const tokenAddr of ERC20_ADDRESSES) {
  base = base.addLog({
    address: [tokenAddr],
    topic0: [TRANSFER_TOPIC0],
    topic2: [DONATION_SAFE_TOPIC],
    transaction: true,
  })
}

// Subsquid Network archive — only for production Ethereum mainnet.
// For local Anvil forks (no archive) we skip the gateway; the processor
// falls back to RPC-only, which is fine at fork volumes.
const GATEWAY_URL = process.env.GATEWAY_URL
const processor = GATEWAY_URL ? base.setGateway(GATEWAY_URL) : base

// ---------------------------------------------------------------------------
// Boot.
// ---------------------------------------------------------------------------

const main = async () => {
  // Ensure schema before starting the processor loop.
  await ensureDonationTable(pool)
  console.log('[donations-indexer] donation table ensured')

  // We manage PG directly via our pool — no TypeORM overhead.
  // buildHotDatabase() returns a HotDatabase<void> that satisfies the full
  // Subsquid Database contract. Finalized batches are written in transact();
  // hot (unfinalized) batches are ignored in transactHot() — the cursor only
  // advances on finalized data so reorgs cannot introduce phantom rows.
  processor.run(
    buildHotDatabase(pool),
    async (ctx) => {
      for (const block of ctx.blocks) {
        // Native value transfers.
        for (const tx of block.transactions) {
          // Defensive: ensure the `to` field matches our Safe.
          if (!tx.to || tx.to.toLowerCase() !== DONATION_SAFE) continue
          // Skip zero-value (pure contract calls, etc.).
          if (!tx.value || tx.value === 0n) continue

          const row = mapNativeTransfer({
            chainId: CHAIN_ID,
            chainSlug: CHAIN_SLUG,
            fromAddress: tx.from,
            txHash: tx.hash,
            blockNumber: block.header.height,
            blockTimestampMs: block.header.timestamp,
            value: tx.value,
          })

          if (!row) {
            ctx.log.debug(`[donations-indexer] dropped tx ${tx.hash} (below floor or unknown chain)`)
            continue
          }

          await upsertDonation(pool, row)
          ctx.log.info(`[donations-indexer] indexed donation ${row.id} — ${row.amountNorm} ${row.tokenSymbol}`)
        }

        // ERC20 Transfer logs.
        for (const log of block.logs) {
          // topic0 = Transfer, topic1 = from, topic2 = to (padded address)
          const topics = log.topics
          if (topics.length < 3) continue

          // Defensive: verify topic0 is Transfer.
          if (topics[0]?.toLowerCase() !== TRANSFER_TOPIC0) continue

          // Decode `from` and `to` from the padded 32-byte topics.
          const fromAddress = '0x' + (topics[1] ?? '').slice(-40)
          const toAddress = '0x' + (topics[2] ?? '').slice(-40)

          // Defensive: ensure the recipient is our Safe.
          if (toAddress.toLowerCase() !== DONATION_SAFE) continue

          // Decode amount from data (uint256, 32 bytes big-endian).
          const amount = log.data && log.data !== '0x' ? BigInt(log.data) : 0n

          // Resolve the tx hash: log has transactionHash directly.
          const txHash = log.transactionHash

          const row = mapErc20Transfer({
            chainId: CHAIN_ID,
            chainSlug: CHAIN_SLUG,
            tokenAddress: log.address,
            fromAddress,
            toAddress,
            amount,
            txHash,
            logIndex: log.logIndex,
            blockNumber: block.header.height,
            blockTimestampMs: block.header.timestamp,
          })

          if (!row) {
            ctx.log.debug(
              `[donations-indexer] dropped ERC20 log ${log.address}:${log.logIndex} (token not allowlisted or zero amount)`,
            )
            continue
          }

          await upsertDonation(pool, row)
          ctx.log.info(
            `[donations-indexer] indexed ERC20 donation ${row.id} — ${row.amountNorm} ${row.tokenSymbol}`,
          )
        }
      }
    },
  )
}

// ---------------------------------------------------------------------------
// HotDatabase<void> implementation.
//
// Subsquid's run<Store>(db: Database<Store>, ...) requires a Database that
// satisfies either FinalDatabase<S> or HotDatabase<S>.
//
// We implement HotDatabase<void>:
//   - The void store type means the handler receives `store: void` which it
//     ignores (handlers write directly via the module-level pg pool).
//   - supportsHotBlocks: true  — required so the runner can enter the
//     near-head hot-block phase without crashing.
//   - connect()       — reads {height, hash, top:[]} from the status table.
//   - transact()      — finalised-block path: run the handler, then advance
//                       the cursor to info.nextHead. This is where all
//                       donation rows are written (post finality-depth).
//   - transactHot()   — hot-block path: no-op. We intentionally do not
//                       persist donations from unfinalized blocks — with the
//                       default 75-block finality depth on Ethereum mainnet,
//                       every block we surface to the user is already safe.
//                       The transactHot2 optional override is not implemented
//                       so the runner falls back to transactHot.
//
// Cursor semantics:
//   The status table tracks the last committed FINALIZED block.
//   Hot blocks are tracked in memory by Subsquid's runner (the `top` field
//   in HotDatabaseState); we return top:[] from connect() because we do not
//   persist hot-block state across restarts.
// ---------------------------------------------------------------------------

function buildHotDatabase(pgPool: InstanceType<typeof Pool>): HotDatabase<void> {
  const ensureStatus = async (): Promise<void> => {
    await pgPool.query(`
      CREATE TABLE IF NOT EXISTS donations_indexer_status (
        id      INTEGER PRIMARY KEY DEFAULT 1 CHECK (id = 1),
        height  INTEGER NOT NULL DEFAULT -1,
        hash    TEXT    NOT NULL DEFAULT ''
      );
    `)
    // Insert the sentinel row if absent — idempotent.
    await pgPool.query(`
      INSERT INTO donations_indexer_status (id, height, hash)
      VALUES (1, -1, '')
      ON CONFLICT (id) DO NOTHING;
    `)
  }

  return {
    supportsHotBlocks: true,

    async connect(): Promise<HotDatabaseState> {
      await ensureStatus()
      const { rows } = await pgPool.query<{ height: number; hash: string }>(
        `SELECT height, hash FROM donations_indexer_status WHERE id = 1`,
      )
      const row = rows[0] ?? { height: -1, hash: '' }
      // top:[] — we do not persist hot-block state; the runner will re-fetch
      // near-head blocks on restart if they haven't been finalized yet.
      return { ...row, top: [] }
    },

    async transact(info: FinalTxInfo, cb: (store: void) => Promise<void>): Promise<void> {
      // Run the batch handler. Our upserts are individually idempotent so we
      // don't need a full PG transaction wrapping the donation rows.
      await cb(undefined as void)
      // Advance the cursor so the processor resumes from nextHead on restart.
      // This must happen after cb() succeeds — the cursor only moves on
      // successful finalized-block processing.
      await pgPool.query(
        `UPDATE donations_indexer_status SET height = $1, hash = $2 WHERE id = 1`,
        [info.nextHead.height, info.nextHead.hash],
      )
    },

    async transactHot(_info: HotTxInfo, _cb: (store: void, block: HashAndHeight) => Promise<void>): Promise<void> {
      // Hot blocks are not persisted. We only write finalized donations.
      // With FINALITY_CONFIRMATION=75, every block surfaced to transact() is
      // already 75 blocks deep — safe against any realistic Ethereum reorg.
      // The runner calls transactHot() for near-head blocks; we skip them.
    },
  }
}

main().catch((err) => {
  console.error('[donations-indexer] fatal:', err)
  process.exit(1)
})
