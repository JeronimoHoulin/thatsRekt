#!/usr/bin/env bash
# =============================================================================
# Anvil bootstrap — deploys thatsRekt onto a running Anvil and emits the
# resulting addresses + start block to .deployed.json.
# =============================================================================
# Prerequisite: Anvil must already be running and reachable at $ANVIL_RPC.
# Typical workflow:
#
#   cd indexer
#   docker compose -f docker-compose.yml -f docker-compose.anvil.yml up -d anvil
#   ../contracts/script/anvil/bootstrap.sh
#
# The bootstrap is idempotent — DeployDev.s.sol detects already-deployed
# CREATE2 contracts and short-circuits, so re-running is a no-op.
#
# Output: contracts/script/anvil/.deployed.json (gitignored). Read by
# downstream tools (or copy by hand) into indexer/.env as
# CONTRACT_ANVIL + START_BLOCK_ANVIL.
# =============================================================================

set -euo pipefail

# --- Config (override via env) -----------------------------------------------
ANVIL_RPC="${ANVIL_RPC:-http://localhost:8545}"
DEV_EOA="${DEV_EOA:-0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266}"
# Anvil default account 0 private key. Public on the foundry website; safe to
# hardcode because it's the canonical dev key. NEVER use this on mainnet.
DEV_KEY="${DEV_KEY:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
CONTRACTS_DIR="$(cd "$SCRIPT_DIR/../.." && pwd)"
OUT_JSON="$SCRIPT_DIR/.deployed.json"
TMP_LOG="$(mktemp -t thatsrekt-bootstrap.XXXXXX)"
trap 'rm -f "$TMP_LOG"' EXIT

# --- 1. Verify Anvil is reachable -------------------------------------------
echo "==> Checking Anvil at $ANVIL_RPC"
if ! CHAIN_ID=$(cast chain-id --rpc-url "$ANVIL_RPC" 2>/dev/null); then
    cat <<EOF >&2
ERROR: Anvil not reachable at $ANVIL_RPC.

Start Anvil first. Options:

  # via docker compose (preferred — uses the Phase 4 anvil service)
  cd indexer
  docker compose -f docker-compose.yml -f docker-compose.anvil.yml up -d anvil

  # native (requires foundry installed locally)
  anvil --fork-url \$ANVIL_FORK_URL --host 0.0.0.0 --chain-id 31337 --block-time 2

EOF
    exit 1
fi
echo "    chain-id = $CHAIN_ID"

# --- 2. Deploy via DeployDev.s.sol ------------------------------------------
echo "==> Deploying thatsRekt via DeployDev.s.sol (owner=$DEV_EOA)"
cd "$CONTRACTS_DIR"

# Pipe forge output through tee so we can both display it and grep the
# emitted addresses below. `--silent` suppresses forge's own progress
# noise but our DeployDev's console2.log lines still come through.
GOVERNANCE_OWNER="$DEV_EOA" \
forge script script/DeployDev.s.sol \
    --rpc-url "$ANVIL_RPC" \
    --private-key "$DEV_KEY" \
    --broadcast \
    --slow \
    -vvv 2>&1 | tee "$TMP_LOG"

# --- 3. Extract addresses + current block -----------------------------------
PROXY=$(grep -E '^\s*Proxy:' "$TMP_LOG" | tail -1 | awk '{print $NF}')
TIMELOCK=$(grep -E '^\s*TimelockController:' "$TMP_LOG" | tail -1 | awk '{print $NF}')
IMPL=$(grep -E '^\s*Implementation:' "$TMP_LOG" | tail -1 | awk '{print $NF}')
BLOCK=$(cast block-number --rpc-url "$ANVIL_RPC")

if [[ -z "$PROXY" || -z "$TIMELOCK" || -z "$IMPL" ]]; then
    echo "ERROR: failed to extract deployed addresses from forge script output." >&2
    echo "Inspect $TMP_LOG for details." >&2
    exit 2
fi

# --- 4. Write .deployed.json -------------------------------------------------
cat > "$OUT_JSON" <<EOF
{
  "chainId": $CHAIN_ID,
  "blockNumber": $BLOCK,
  "implementation": "$IMPL",
  "timelock": "$TIMELOCK",
  "proxy": "$PROXY",
  "owner": "$DEV_EOA"
}
EOF

echo
echo "==> Wrote $OUT_JSON"
cat "$OUT_JSON"
echo
echo "==> Paste into indexer/.env:"
echo "    CONTRACT_ANVIL=$PROXY"
echo "    START_BLOCK_ANVIL=$BLOCK"
