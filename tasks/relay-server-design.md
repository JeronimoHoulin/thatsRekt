# thatsRekt Relay Server — Design (Phase 10+)

**Status:** Design draft. Implementation deferred to a follow-up phase after Phase 9 (on-chain title) lands.
**Owner:** bauti
**Date:** 2026-04-27

A Go websocket service that receives hack alerts from an external AI detection provider and submits them on-chain to thatsRekt as a whitelisted poster.

## Goals

1. **Pure relay.** The server submits exactly what the AI provider tells it to. It does not enrich, deduplicate, retry-merge, batch, or guess any field.
2. **Whitelisted poster.** The server's signing key is a single EOA that has been added to the on-chain whitelist by governance on each chain it posts to.
3. **Multi-chain dispatch.** The AI provider tags each alert with one or more target chains (`["base"]`, `["base", "eth"]`, or `["all"]`). The server submits one transaction per target chain.
4. **Updatable.** The protocol must accept follow-up update messages for an existing post (amend title, amend note, add attackers, add victims). Implementation of update handling deferred to a later phase, but the message shape and dispatch mechanism must be reserved in the wire protocol now.

## Non-goals

- The server does NOT make trust decisions. If the AI provider says "post this", it posts. Filtering, ranking, or refusing alerts is the AI provider's job.
- The server does NOT subscribe to on-chain events. It is a one-way relay (off-chain → on-chain).
- The server does NOT vote, retract, or remove posts. Those flows belong to other whitelisters / governance.
- No backfill. The server processes only live websocket messages; if it disconnects, the AI provider is responsible for redelivery.

## Wire protocol (websocket)

Direction: **AI provider → relay server**, persistent websocket connection. Heartbeat every 30s (provider sends `{"type":"ping"}`, server responds `{"type":"pong"}`).

### Message envelope (every payload)

```jsonc
{
  "type": "post.create" | "post.amend_title" | "post.amend_note" | "post.add_attackers" | "post.add_victims" | "ping",
  "id": "msg-<uuid>",            // provider-assigned, unique per message; server uses for dedup + ack
  "timestamp": "2026-04-27T22:00:00Z",  // ISO 8601, when provider issued the message
  "payload": { ... }             // type-specific (see below)
}
```

The server replies on the same websocket with:

```jsonc
{
  "type": "ack" | "nack",
  "msg_id": "msg-<uuid>",        // echo of the request id
  "results": [                    // one entry per target chain
    {
      "chain": "base",
      "status": "submitted" | "skipped" | "failed",
      "tx_hash": "0x...",          // when status = submitted
      "post_id": "42",             // when known (from receipt log decoding)
      "error": "..."               // when status = failed
    },
    ...
  ]
}
```

The server's response is best-effort in the sense that on-chain confirmation of inclusion can take seconds; the server submits the tx and returns `submitted` as soon as the RPC accepts it. Confirmation is the AI provider's concern (it can re-query Mesh).

### `post.create` payload

```jsonc
{
  "chains": ["base", "ethereum"],   // OR "all"
  "title": "Aave drainer detected",
  "attackers": ["0xdead..."],
  "victims": ["0xbeef..."],
  "note": "Initial scoping — drainer pulled ~$200k...",
  "attacked_at": 1777340000          // unix seconds, on-chain attack time
}
```

The server submits one `post(title, attackers, victims, note, attackedAt)` tx per chain in `chains`. **Nothing in the payload is filled in or defaulted by the server.** If `note` is missing or `chains` is empty, that's a malformed message and the server returns `nack` for it.

`"chains": "all"` resolves to the registered chain set at the time of receipt. The server's chain registry is configured at startup (env-driven, mirrors `indexer/src/chains.ts`).

### `post.amend_title` (update — deferred implementation, reserved now)

```jsonc
{
  "chains": ["base"],
  "post_id": "42",                   // on-chain post id on each target chain
  "new_title": "Aave drainer — confirmed multi-vault"
}
```

The server submits `amendTitle(post_id, new_title)`.

> **Note:** post ids on each chain are independent. The AI provider must track per-chain post ids. The server's `ack` from the original `post.create` returns `post_id` per chain — the provider stores those and references them in subsequent updates.

### `post.amend_note` (deferred)

```jsonc
{
  "chains": ["base"],
  "post_id": "42",
  "new_note": "Updated scoping with new addresses..."
}
```

### `post.add_attackers` / `post.add_victims` (deferred)

```jsonc
{
  "chains": ["base"],
  "post_id": "42",
  "addresses": ["0x..."]
}
```

## Architecture

```
┌─────────────────────┐    websocket     ┌──────────────────────┐
│  AI detection       │◀────────────────▶│  thatsRekt-relay     │
│  provider           │   alerts + acks  │  (Go)                │
└─────────────────────┘                  │                      │
                                         │  - ws server (gorilla│
                                         │    /websocket or std │
                                         │    net/http)         │
                                         │  - per-message       │
                                         │    dispatcher        │
                                         │  - per-chain queue   │
                                         │    + signer          │
                                         │  - structured logs   │
                                         └──────┬───────────────┘
                                                │
                       ┌────────────────────────┼────────────────────────┐
                       ↓                        ↓                        ↓
                ┌────────────┐           ┌────────────┐           ┌────────────┐
                │  base RPC  │           │  arbitrum  │           │  optimism  │
                │  (signer)  │           │  RPC       │           │  RPC       │
                └────────────┘           └────────────┘           └────────────┘
```

### Goroutines

- **1 ws-listener goroutine** — reads incoming messages, validates envelope, hands off to dispatcher.
- **1 dispatcher goroutine** — fans incoming messages out to per-chain queues based on `chains` field.
- **N per-chain worker goroutines** — one per configured chain. Each owns its own RPC client + signing nonce manager. Serializes tx submission per chain to keep nonces clean.

### Nonce management

Each chain worker maintains its own `pending nonce` cursor:
- On startup, query `eth_getTransactionCount(latest)` for the relay EOA.
- Increment local counter on each submission.
- On RPC error indicating nonce mismatch, refresh from chain (`eth_getTransactionCount(pending)`) and retry once.
- No mempool-replacement logic — the relay is single-writer per chain, so this is sufficient.

### Idempotency

Same `id` (envelope-level uuid) seen twice within the dedup window (15min, in-memory ring buffer) → return cached ack, do not re-submit. This protects against websocket reconnects where the provider replays the buffer.

The relay does NOT enforce on-chain idempotency — if the provider sends two distinct `id`s with identical payload, two posts get created. That's the provider's responsibility.

## Security & ops

### Trust model

- **Inbound (provider → relay):** the websocket connection is authenticated via a shared secret in the upgrade-request `Authorization: Bearer ...` header. TLS required. Single provider per deployment. (Multi-tenant relay is out of scope for v1.)
- **Outbound (relay → chain):** the relay holds a signing key that is whitelisted on each chain. The key is hot — necessarily — but its blast radius is limited to "can post arbitrary alerts" (no value held, no upgrade authority). Compromise = governance revokes the whitelist via the timelock; integrators see the bad poster's downvotes accumulate.

### Key custody

Three configurable backends, picked by env:
1. **Env var** (`RELAY_PRIVATE_KEY`) — local dev only.
2. **AWS KMS** — production; the relay never has the raw key, just signs via KMS calls.
3. **HSM via PKCS#11** — for higher trust deployments.

KMS is the recommended default for any deployment that posts to mainnet.

### Rate limits

- **Per-chain submission rate:** 1 tx every 2 seconds (configurable). Hard cap to avoid burning gas if the provider misbehaves.
- **Per-provider message rate:** 30 messages/minute (configurable). Excess returns `nack` with `error: "rate_limited"`.

### Observability

Structured JSON logs (zerolog or slog). Required fields per log line: `msg_id`, `type`, `chain`, `event` (`received` / `submitting` / `submitted` / `failed`), `tx_hash` if applicable, `error` if applicable.

Prometheus metrics:
- `relay_messages_received_total{type}`
- `relay_submissions_total{chain, status}`
- `relay_submission_latency_seconds{chain}` (histogram)
- `relay_nonce_resyncs_total{chain}`
- `relay_ws_connection_state` (0/1)

## Configuration

```yaml
# relay.yaml
listen_addr: ":8080"
ws_path: "/ws"
provider_token: "${PROVIDER_TOKEN}"   # shared secret

contract:
  abi_path: ./abi/ThatsRekt.json

chains:
  base:
    rpc: "${BASE_RPC}"
    chain_id: 8453
    contract: "0x..."
  ethereum:
    rpc: "${ETHEREUM_RPC}"
    chain_id: 1
    contract: "0x..."
  arbitrum: ...

signer:
  type: "kms"                         # env | kms | pkcs11
  key_id: "alias/thatsrekt-relay"     # AWS KMS

dedup_window: "15m"
per_chain_rate_limit: "0.5/s"
per_provider_rate_limit: "30/m"

log_level: "info"
```

## File layout

```
relay/
├── cmd/
│   └── relay/
│       └── main.go                  ← entrypoint
├── internal/
│   ├── ws/
│   │   ├── server.go                ← gorilla/websocket handler
│   │   ├── auth.go                  ← bearer-token check
│   │   └── codec.go                 ← envelope encode/decode
│   ├── dispatcher/
│   │   ├── dispatcher.go            ← fans to per-chain workers
│   │   └── chain_worker.go          ← per-chain submitter
│   ├── signer/
│   │   ├── signer.go                ← interface
│   │   ├── env.go                   ← env-key implementation
│   │   ├── kms.go                   ← AWS KMS implementation
│   │   └── pkcs11.go                ← PKCS#11 implementation
│   ├── chain/
│   │   └── registry.go              ← per-chain config + ABI binding
│   ├── thatsrekt/                   ← abigen-generated bindings
│   │   └── thatsrekt.go
│   └── metrics/
│       └── metrics.go               ← prometheus collectors
├── abi/
│   └── ThatsRekt.json               ← copied from contracts/out/
├── relay.yaml.example
├── go.mod
├── go.sum
└── README.md
```

## Implementation phases

The websocket server lands in three sub-phases to keep PRs reviewable:

### Sub-phase A — `post.create` only

- Wire protocol envelope handling (decode, validate, ack/nack)
- Single-chain happy path (`base` only, env-key signer)
- gorilla/websocket server with bearer-token auth
- Logs + basic metrics

### Sub-phase B — multichain + nonce mgmt + KMS

- Multi-chain dispatcher with per-chain worker pool
- `chains: "all"` resolution
- Nonce cursor + resync on mismatch
- AWS KMS signer
- Prometheus metrics

### Sub-phase C — update messages

- `post.amend_title`, `post.amend_note`, `post.add_attackers`, `post.add_victims`
- Per-chain post-id tracking is provider-side; server is stateless
- Same dispatcher pattern, different ABI calls

## Open questions for operator decision

1. **Provider authentication:** bearer token enough, or do we want mTLS / signed messages too?
2. **Single provider or multi-tenant?** Current design assumes one provider per relay deployment. Multi-tenant adds auth scoping + per-tenant rate limits.
3. **Failed submissions:** retry policy. Current design says "submit once, return failed status, provider re-sends if it wants". Alternative: in-process retry with backoff. Recommend keeping it stateless (provider re-sends).
4. **Message persistence:** none in v1 (in-memory only). If the relay crashes between `submitted` and the websocket ack reaching the provider, the provider may resubmit and we'd dedup on `id`. Anything richer requires a DB.
5. **Confirmation tracking:** should the relay watch tx receipts and emit `post.confirmed` messages back through the websocket? Current design says no — Mesh + the indexer is the source of confirmation truth, and the provider can subscribe to that.

## What this design defers

- Re-org handling (RPC re-org of an already-submitted post) — current design says "the indexer's reorg handling makes the post disappear if it doesn't survive; the AI provider re-issues if it cares".
- Cost accounting per provider / per tenant.
- Geographic redundancy (multiple relay instances, leader election, shared nonce manager).
- A web UI for the relay's status (logs + metrics over Prometheus + Grafana cover the operator's needs for v1).

## Why Go?

- Strong static typing for the wire protocol decoder + ABI bindings (`abigen`).
- Goroutines cleanly model "one ws-reader + N per-chain submitters + 1 dispatcher".
- `go-ethereum`/`abigen` toolchain is the gold standard for typed contract calls.
- Single static binary deploys easily to ECS / Fargate / a basic Linux box.
- KMS + PKCS#11 SDKs are mature.

## Success criteria

When implemented:
1. AI provider connects via websocket, sends `post.create`, gets `ack` with `submitted` for each target chain within 1s.
2. Posts appear in Mesh's unified feed within 5s of submission (anvil instant-block dev), or within typical chain finality on real chains.
3. Reconnect + replay of buffered messages produces zero duplicates (dedup by `id`).
4. Killing the relay mid-burst loses zero confirmed posts (in-flight submissions complete because the chain has them; only un-acked ones might re-fire on reconnect).
5. Adding a new chain is a `chains:` block in `relay.yaml` and a one-time on-chain governance call to add the relay's EOA to the new chain's whitelist. No code changes.
