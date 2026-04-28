#!/usr/bin/env bash
# =============================================================================
# Start an ngrok tunnel pointing at the local relay's port and capture
# the resulting public URL.
#
# Idempotent:
#   - If $NGROK_PID_FILE points to a running ngrok, this is a no-op (we
#     just re-read the URL from ngrok's local API).
#   - Otherwise: launch fresh, wait for the tunnel to come up, capture URL.
#
# Outputs:
#   $NGROK_PID_FILE   — pid of the running ngrok process
#   $NGROK_URL_FILE   — the https:// public URL
#   $NGROK_LOG_FILE   — ngrok stdout/stderr
#
# Prereqs: ngrok on PATH, with an authtoken configured (`ngrok config
# add-authtoken <token>` once per machine). Free-tier ngrok works for
# this — the URL just rotates on restart.
# =============================================================================

set -euo pipefail

NGROK_PID_FILE="${NGROK_PID_FILE:-/tmp/thatsrekt-ngrok.pid}"
NGROK_URL_FILE="${NGROK_URL_FILE:-/tmp/thatsrekt-ngrok.url}"
NGROK_LOG_FILE="${NGROK_LOG_FILE:-/tmp/thatsrekt-ngrok.log}"
RELAY_PORT="${RELAY_PORT:-8080}"
NGROK_API="${NGROK_API:-http://127.0.0.1:4040}"

if ! command -v ngrok >/dev/null 2>&1; then
    cat >&2 <<EOF
ERROR: ngrok not on PATH.
Install: https://ngrok.com/download
Then run once: ngrok config add-authtoken <your_token>
EOF
    exit 1
fi

# --- Already running? --------------------------------------------------------
if [[ -f "$NGROK_PID_FILE" ]] && kill -0 "$(cat "$NGROK_PID_FILE")" 2>/dev/null; then
    URL="$(curl -sf "$NGROK_API/api/tunnels" 2>/dev/null | jq -r '.tunnels[0].public_url' 2>/dev/null || echo "")"
    if [[ -n "$URL" && "$URL" != "null" ]]; then
        echo "$URL" > "$NGROK_URL_FILE"
        echo "    ngrok already running (pid $(cat "$NGROK_PID_FILE"))"
        echo "    public URL: $URL"
        exit 0
    fi
    echo "    ngrok pid file is stale; starting fresh"
fi

# --- Launch ------------------------------------------------------------------
echo "    starting ngrok http $RELAY_PORT"
nohup ngrok http "$RELAY_PORT" --log=stdout > "$NGROK_LOG_FILE" 2>&1 &
NGROK_PID=$!
echo "$NGROK_PID" > "$NGROK_PID_FILE"

# Wait for ngrok's local API to come up + report a tunnel.
echo -n "    waiting for ngrok tunnel"
URL=""
for _ in $(seq 1 20); do
    URL="$(curl -sf "$NGROK_API/api/tunnels" 2>/dev/null | jq -r '.tunnels[0].public_url' 2>/dev/null || echo "")"
    if [[ -n "$URL" && "$URL" != "null" ]]; then
        echo " ok"
        echo "$URL" > "$NGROK_URL_FILE"
        echo "    pid=$NGROK_PID logs=$NGROK_LOG_FILE"
        echo "    public URL: $URL"
        exit 0
    fi
    echo -n "."
    sleep 1
done
echo " FAILED"
echo "ERROR: ngrok didn't come up — last log lines:" >&2
tail -20 "$NGROK_LOG_FILE" >&2
exit 1
