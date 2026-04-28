#!/usr/bin/env bash
# =============================================================================
# Anvil reset — wipes Anvil's state, drops + recreates thatsrekt_anvil
# database, restarts Anvil, and re-runs bootstrap.sh.
# =============================================================================
# Use after schema changes, or when Anvil's accumulated state is interfering
# with a fresh test run.
# =============================================================================

set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
INDEXER_DIR="$(cd "$SCRIPT_DIR/../../../indexer" && pwd)"
COMPOSE_FILES=(-f docker-compose.yml -f docker-compose.anvil.yml)

cd "$INDEXER_DIR"

echo "==> Stopping anvil + per-anvil services"
docker compose "${COMPOSE_FILES[@]}" stop \
    anvil migrate-anvil processor-anvil graphql-anvil 2>/dev/null || true
docker compose "${COMPOSE_FILES[@]}" rm -fv anvil 2>/dev/null || true

echo "==> Resetting thatsrekt_anvil database"
# DROP+CREATE rather than \dt-level cleanup so any schema drift gets nuked.
docker compose exec -T db psql -U postgres <<'EOF'
DROP DATABASE IF EXISTS thatsrekt_anvil;
CREATE DATABASE thatsrekt_anvil;
EOF

echo "==> Restarting anvil"
docker compose "${COMPOSE_FILES[@]}" up -d anvil

echo "==> Waiting for anvil healthcheck"
for i in $(seq 1 20); do
    if cast chain-id --rpc-url http://localhost:8545 >/dev/null 2>&1; then
        break
    fi
    sleep 1
done

echo "==> Re-bootstrapping"
"$SCRIPT_DIR/bootstrap.sh"
