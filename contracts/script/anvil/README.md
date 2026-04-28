# Anvil bootstrap

Local Anvil fork of Sepolia for tight dev loops on the thatsRekt contract + indexer stack. Forking gives us the CREATE2 singleton factory + Safe singleton factory + realistic block numbers for free, while running locally with no faucet, no rate limits, and instant blocks.

## Prerequisites

- `foundry` installed (`anvil`, `cast`, `forge`).
- A Sepolia RPC URL for Anvil to fork from. Set `ANVIL_FORK_URL` in `indexer/.env` (typically the same value as `RPC_SEPOLIA_HTTP`).
- Docker + docker compose (or run anvil natively — see below).

## Quickstart (compose)

```bash
# 1. Start Anvil (fork of Sepolia, chainId 31337)
cd indexer
docker compose -f docker-compose.yml -f docker-compose.anvil.yml up -d anvil

# 2. Deploy thatsRekt onto Anvil via DeployDev (EOA owner = Anvil account 0)
../contracts/script/anvil/bootstrap.sh

# 3. Copy the printed CONTRACT_ANVIL + START_BLOCK_ANVIL values into indexer/.env

# 4. Bring up the rest of the Anvil indexer slice
docker compose -f docker-compose.yml -f docker-compose.anvil.yml up -d \
    db migrate-anvil processor-anvil graphql-anvil
```

The bootstrap is idempotent — re-running it on the same Anvil instance is a no-op (DeployDev detects existing CREATE2 deploys and short-circuits).

## Reset

Wipe Anvil's state, drop+recreate the `thatsrekt_anvil` database, and re-bootstrap:

```bash
contracts/script/anvil/reset.sh
```

Use after schema changes, or when accumulated Anvil state interferes with a fresh test run.

## Configuration

Both scripts pick up overrides from env:

| Var | Default | Purpose |
|---|---|---|
| `ANVIL_RPC` | `http://localhost:8545` | Where the bootstrap script reaches Anvil. |
| `DEV_EOA` | `0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266` | Anvil default account 0. The owner / proposer / executor of the timelock. |
| `DEV_KEY` | `0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80` | Anvil default account 0 private key. **Public test mnemonic — never use on mainnet.** |
| `ANVIL_FORK_URL` | (required) | Sepolia RPC URL Anvil forks state from. Set in `indexer/.env`. |

## Running anvil natively (no docker)

```bash
anvil \
    --fork-url $ANVIL_FORK_URL \
    --host 0.0.0.0 \
    --chain-id 31337 \
    --block-time 2 \
    --port 8545
```

`--chain-id 31337` (Anvil default) is intentional — keeps Anvil's chain id distinct from real Sepolia (11155111) so the indexer treats it as a separate chain even though the forked state is from Sepolia.

## Output

`bootstrap.sh` writes `contracts/script/anvil/.deployed.json` (gitignored) with the deployed addresses + start block:

```json
{
  "chainId": 31337,
  "blockNumber": 12345678,
  "implementation": "0x...",
  "timelock": "0x...",
  "proxy": "0x...",
  "owner": "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
}
```

This is the source of truth for the local Anvil deploy. Copy `proxy` → `CONTRACT_ANVIL` and `blockNumber` → `START_BLOCK_ANVIL` in `indexer/.env`.

## Why fork from Sepolia (not cold-start)?

Forking gives us:
- **CREATE2 singleton factory** at `0x4e59…` already deployed → existing deploy script works unchanged.
- **Safe singleton factory** in case we ever want a real Safe in the loop for parity testing.
- **Realistic block numbers + gas pricing** instead of arbitrary single-digit numbers.
- The indexer can fetch real historical data when sanity-checking event handlers.

## Why chain id 31337 (not Sepolia's 11155111)?

The indexer's chain registry treats each `chainId` as a unique chain. If Anvil reported chain id `11155111`, the indexer's `anvil` processor and `sepolia` processor would both think they're indexing chain `11155111` — colliding on per-address keys and confusing every cross-chain query.

Anvil's default `31337` is well-known and reserved for local development. Keeping that default is the right call.
