#!/usr/bin/env bash
# =============================================================================
# Write detector/.env from the live local stack:
#   - WEBHOOK_BASE_URL = current ngrok public URL
#   - WEBHOOK_TOKEN    = current relay bearer token
#   - WEBHOOK_CHAIN    = anvil-eth (single chain for v1)
#
# This DOES NOT set OTOMATO_API_KEY — that's an operator secret. If
# detector/.env already exists with an OTOMATO_API_KEY value, we
# preserve it; otherwise we leave it blank and the deploy will fail
# fast with a clear error.
# =============================================================================

set -euo pipefail

NGROK_URL_FILE="${NGROK_URL_FILE:-/tmp/thatsrekt-ngrok.url}"
RELAY_TOKEN_FILE="${RELAY_TOKEN_FILE:-/tmp/thatsrekt-relay.token}"
WEBHOOK_CHAIN="${WEBHOOK_CHAIN:-anvil-eth}"

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "$SCRIPT_DIR/../.." && pwd)"
DETECTOR_ENV="$REPO_ROOT/detector/.env"

if [[ ! -f "$NGROK_URL_FILE" ]]; then
    echo "ERROR: $NGROK_URL_FILE missing — run \`make ngrok-up\` first" >&2
    exit 1
fi
if [[ ! -f "$RELAY_TOKEN_FILE" ]]; then
    echo "ERROR: $RELAY_TOKEN_FILE missing — run \`make relay-up\` first" >&2
    exit 1
fi

NGROK_URL="$(cat "$NGROK_URL_FILE")"
RELAY_TOKEN="$(cat "$RELAY_TOKEN_FILE")"

# --- Preserve OTOMATO_API_KEY if it's set in the existing file --------------
EXISTING_API_KEY=""
if [[ -f "$DETECTOR_ENV" ]]; then
    EXISTING_API_KEY="$(awk -F= '/^OTOMATO_API_KEY=/{print substr($0,16); exit}' "$DETECTOR_ENV" || true)"
fi

cat > "$DETECTOR_ENV" <<EOF
# Auto-written by ops/scripts/write-detector-env.sh.
# Re-run \`make detector-env\` to refresh after relay/ngrok restarts.
# Manual edits are preserved EXCEPT for the three webhook fields below.

OTOMATO_API_KEY=${EXISTING_API_KEY}
OTOMATO_API_URL=https://api.otomato.xyz/api

WEBHOOK_BASE_URL=${NGROK_URL}
WEBHOOK_TOKEN=${RELAY_TOKEN}
WEBHOOK_CHAIN=${WEBHOOK_CHAIN}
EOF

chmod 600 "$DETECTOR_ENV"

echo "    wrote $DETECTOR_ENV"
echo "      WEBHOOK_BASE_URL=$NGROK_URL"
echo "      WEBHOOK_TOKEN=<32 chars>"
echo "      WEBHOOK_CHAIN=$WEBHOOK_CHAIN"
if [[ -z "$EXISTING_API_KEY" ]]; then
    echo "      OTOMATO_API_KEY=<empty — set it before \`npm run deploy\`>"
else
    echo "      OTOMATO_API_KEY=<preserved from existing file>"
fi
