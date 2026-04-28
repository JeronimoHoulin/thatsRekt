#!/usr/bin/env bash
# =============================================================================
# Start the thatsRekt relay against the locally-deployed anvil-eth.
#
# Idempotent:
#   - If $RELAY_PID_FILE points to a running process, this is a no-op.
#   - Otherwise: ensure the EOA is whitelisted, generate (or reuse) the
#     bearer token, launch the relay in the background.
#
# Outputs:
#   $RELAY_PID_FILE        — pid of the running relay
#   $RELAY_TOKEN_FILE      — bearer token (also passed to detector/.env)
#   $RELAY_LOG_FILE        — relay stdout/stderr
#
# Prereqs: go, cast, jq on PATH; lan-up + anvil-bootstrap done.
# =============================================================================

set -euo pipefail

RELAY_PID_FILE="${RELAY_PID_FILE:-/tmp/thatsrekt-relay.pid}"
RELAY_TOKEN_FILE="${RELAY_TOKEN_FILE:-/tmp/thatsrekt-relay.token}"
RELAY_LOG_FILE="${RELAY_LOG_FILE:-/tmp/thatsrekt-relay.log}"
RELAY_PORT="${RELAY_PORT:-8080}"
RELAY_EOA_KEY_FILE="${RELAY_EOA_KEY_FILE:-/tmp/thatsrekt-relay-eoa.key}"
ANVIL_RPC="${ANVIL_RPC:-http://localhost:8545}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
RELAY_DIR="$REPO_ROOT/relay"
DEPLOYED_JSON="$REPO_ROOT/contracts/script/anvil/.deployed.anvil-eth.json"

# --- Already running? --------------------------------------------------------
if [[ -f "$RELAY_PID_FILE" ]] && kill -0 "$(cat "$RELAY_PID_FILE")" 2>/dev/null; then
    echo "    relay already running (pid $(cat "$RELAY_PID_FILE"))"
    echo "    bearer token: $(cat "$RELAY_TOKEN_FILE")"
    exit 0
fi

# --- Prereqs -----------------------------------------------------------------
if [[ ! -f "$DEPLOYED_JSON" ]]; then
    echo "ERROR: $DEPLOYED_JSON not found — run \`make anvil-bootstrap\` first" >&2
    exit 1
fi

# --- Ensure EOA + whitelist --------------------------------------------------
"$SCRIPT_DIR/whitelist-relay-eoa.sh"

# --- Generate or reuse token -------------------------------------------------
if [[ ! -f "$RELAY_TOKEN_FILE" ]]; then
    head -c 32 /dev/urandom | base64 | tr -d '/+=\n' > "$RELAY_TOKEN_FILE"
    chmod 600 "$RELAY_TOKEN_FILE"
fi
TOKEN="$(cat "$RELAY_TOKEN_FILE")"
KEY="$(cat "$RELAY_EOA_KEY_FILE")"
PROXY="$(jq -r .proxy "$DEPLOYED_JSON")"

# --- Launch ------------------------------------------------------------------
echo "    starting relay on :$RELAY_PORT"
cd "$RELAY_DIR"
RELAY_PROVIDER_TOKEN="$TOKEN" \
RELAY_PRIVATE_KEY="$KEY" \
RELAY_RPC_URL="$ANVIL_RPC" \
RELAY_CONTRACT_ADDRESS="$PROXY" \
RELAY_CHAIN_ID=31337 \
RELAY_CHAIN_NAME=anvil-eth \
RELAY_LISTEN_ADDR=":$RELAY_PORT" \
    nohup go run ./cmd/relay > "$RELAY_LOG_FILE" 2>&1 &
RELAY_PID=$!
echo "$RELAY_PID" > "$RELAY_PID_FILE"

# Wait briefly for /healthz to come up — go run takes a few seconds to
# compile and bind. 30s ceiling so we fail fast on actual problems.
echo -n "    waiting for relay /healthz"
for _ in $(seq 1 30); do
    if curl -sf "http://127.0.0.1:$RELAY_PORT/healthz" >/dev/null 2>&1; then
        echo " ok"
        echo "    relay pid=$RELAY_PID logs=$RELAY_LOG_FILE"
        echo "    bearer token: $TOKEN"
        exit 0
    fi
    echo -n "."
    sleep 1
done
echo " FAILED"
echo "ERROR: relay didn't come up — last log lines:" >&2
tail -20 "$RELAY_LOG_FILE" >&2
exit 1
