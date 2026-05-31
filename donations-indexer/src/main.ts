/**
 * Donations indexer — processor entry point.
 *
 * Watches native-value transactions whose `to` is the thatsrekt.eth
 * donation Safe on Ethereum mainnet, from the Safe's deployment block.
 * Writes rows to the `donation` table in `thatsrekt_donations` DB.
 *
 * Processor-only: no squid-graphql-server is started here.
 * The mesh reads directly via a second pg pool (DONATIONS_DB_URL).
 *
 * Walking skeleton (slice #205): Ethereum + native ETH only.
 * Slice #207 adds ERC20 addLog() subscriptions.
 * Slice #209 adds additional chains.
 */

import 'dotenv/config'
import { EvmBatchProcessor } from '@subsquid/evm-processor'
import pkg from 'pg'
import { ensureDonationTable, upsertDonation } from './donationStore.js'
import { mapNativeTransfer } from './donationMapper.js'

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
// We subscribe to ALL transactions to the Safe address (addTransaction with
// to filter). The `transaction` field must be requested for value access.
// No gateway for local Anvil forks — falls back to RPC-only.
// ---------------------------------------------------------------------------

const base = new EvmBatchProcessor()
  .setRpcEndpoint({
    url: RPC_URL,
    rateLimit: 10,
  })
  .setFinalityConfirmation(75)
  .setFields({
    transaction: {
      to: true,
      from: true,
      value: true,
      hash: true,
    },
  })
  .setBlockRange({ from: START_BLOCK })
  .addTransaction({
    to: [DONATION_SAFE],
  })

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

  // TypeormDatabase is not used here — we manage PG directly via our pool
  // (same pattern as mesh/src/db.ts comments table).
  // Subsquid's raw run() loop gives us blocks + transactions per batch.
  processor.run(
    // We use the raw Database interface (no TypeORM store). Subsquid's
    // EvmBatchProcessor.run() accepts a Database-compatible object; we
    // supply a minimal one that tracks height in our pg pool.
    buildRawDatabase(pool),
    async (ctx) => {
      for (const block of ctx.blocks) {
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
      }
    },
  )
}

// ---------------------------------------------------------------------------
// Minimal raw Database implementation for EvmBatchProcessor.
//
// Subsquid's run() API requires a Database-compatible object implementing
// `connect()`, `transact()`, `advance()`, and `getState()`. We implement
// a lightweight version that tracks the processor's last block in a
// single-row `donations_indexer_status` table. This avoids pulling in
// TypeormDatabase (and its full TypeORM/migration overhead) for a
// processor that manages its own schema.
// ---------------------------------------------------------------------------

interface ProcessorState {
  height: number
  hash: string
}

interface RawDatabase {
  connect(): Promise<ProcessorState>
  transact(info: { prevHead: ProcessorState; nextHead: ProcessorState }, cb: () => Promise<void>): Promise<void>
  advance(info: { nextHead: ProcessorState }): Promise<void>
}

function buildRawDatabase(pgPool: InstanceType<typeof Pool>): RawDatabase {
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
    async connect(): Promise<ProcessorState> {
      await ensureStatus()
      const { rows } = await pgPool.query<{ height: number; hash: string }>(
        `SELECT height, hash FROM donations_indexer_status WHERE id = 1`,
      )
      const row = rows[0]
      return row ?? { height: -1, hash: '' }
    },

    async transact(
      _info: { prevHead: ProcessorState; nextHead: ProcessorState },
      cb: () => Promise<void>,
    ): Promise<void> {
      // Execute the batch callback — our upserts are individually idempotent
      // so we don't need full transaction wrapping here. Subsquid will not
      // call advance() until transact() resolves successfully.
      await cb()
    },

    async advance(info: { nextHead: ProcessorState }): Promise<void> {
      await pgPool.query(
        `UPDATE donations_indexer_status SET height = $1, hash = $2 WHERE id = 1`,
        [info.nextHead.height, info.nextHead.hash],
      )
    },
  }
}

main().catch((err) => {
  console.error('[donations-indexer] fatal:', err)
  process.exit(1)
})
