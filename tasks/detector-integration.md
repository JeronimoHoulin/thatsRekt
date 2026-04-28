# Detector integration — plan

**Branch:** `bauti/detector-integration` (worktree at `thatsRekt-detector-integration/`)
**Base:** `origin/master` @ `6fcd15c` (post-Phase-10A)
**Status:** in-progress
**Date:** 2026-04-28

Vendor Jerry's `rektSDK` into the monorepo as `detector/`, upgrade the AI step
to structured JSON output, route detections both to email (regression-preserved)
and to the local Go relay over an ngrok tunnel, run the full local stack so a
prod Otomato detection lands as a post on `localhost:5173`.

---

## Architecture (target end state)

The relay is DAMM-internal (single-tenant). The contract has many possible
submitters; the relay is just DAMM's automated AI-detected one.

**Local dev path (no Otomato, no ngrok):**

```
detector/src/mock.ts (CLI)
        │
        │ POST /detect (same shape as Otomato sends)
        ▼
relay (localhost:8080) ──► anvil-eth ──► PostCreated event
                                                 │
                                                 ▼
                                       indexer ──► Mesh ──► frontend
```

**Prod path (real Otomato):**

```
8 X account triggers → Otomato cloud (18 AI branches → IF eq "true")
                                          │
                  ┌───────────────────────┴────────────┐
                  │                                    │
                  ▼ HTTP_REQUEST                       ▼ 5× SEND_EMAIL
                ngrok / public DNS                   alert recipients
                  │
                  ▼
                relay (DAMM's server) ──► chain (anvil-eth in dev,
                                                 mainnet/Base later)
                  │
                  ▼
                indexer ──► Mesh ──► frontend
```

**Other submitters (not via the relay):**

- Other whitelisted EOAs (Jerry, future) → frontend "create post" form
  → MetaMask → contract directly
- Anyone reading: GraphQL Mesh, direct contract `cast call`, etc.

Single chain for v1 (anvil-eth). Multichain fan-out is sub-phase B of the
relay (already designed in `tasks/relay-server-design.md`).

---

## Tasks

### Phase 1 — Relay HTTP transports (DONE)

- [x] **1.1** Extracted `ProcessEnvelope` as a shared method on `*Server`.
      WS handler now a thin wrapper. Auth, dedup, submitter all shared.
- [x] **1.2** Added `relay/internal/ws/http.go` with `HandleHTTP`:
      bearer auth, 1 MiB body cap, status mapping (200/400/401/405/413/502),
      ping-over-HTTP rejection. Accepts the **raw envelope** shape — used
      for direct integrators, smoke tests, future provider sources.
- [x] **1.3** `RELAY_HTTP_PATH` env var (default `/post`), wired into
      `cmd/relay/main.go`. Refuses to start if WS == HTTP path.
- [x] **1.4** 11 HTTP tests including cross-transport dedup
      (WS→HTTP same id only submits once).
- [x] **1.5** Build + test + race clean.

### Phase 1.5 — Otomato-shaped `/detect` adapter (NEW)

Otomato's `HTTP_REQUEST` action does dumb variable substitution in the
body string. Building a JSON envelope inside Otomato breaks the moment a
tweet contains a `"` (no JSON escaping). So we accept an **adapter**
shape: body = the AI's JSON output verbatim, headers carry the metadata.

- [ ] **1.5.1** Add `relay/internal/ws/detect.go` with `HandleDetect`:
       - bearer auth (reuse `checkAuth`)
       - required headers: `X-Idempotency-Key`, `X-Tweet-URL`, `X-Chain`,
         `X-Tweet-Timestamp` (RFC3339 or unix seconds), `X-Tweet-Account`
       - body = AI's JSON `{hacked, title, attackers, victims}`,
         strict-decoded with unknown-fields rejected
       - validate `hacked === true` AND `title` non-empty (≤200 bytes —
         the contract caps but we want a fast reject for the obvious case)
       - validate chain is one of the relay's configured chains
       - build internal `PostCreatePayload` from headers + body, run
         through `ProcessEnvelope` (gets dedup, submission, response
         caching for free)
       - return same JSON `Response` shape as `/post`
- [ ] **1.5.2** New env var `RELAY_DETECT_PATH` (default `/detect`).
      Refuses to start if any two of WS/HTTP/DETECT paths collide.
- [ ] **1.5.3** Tests: bearer auth, missing/bad headers, body schema
      mismatches, `hacked: false` rejection, unknown chain rejection,
      happy path, cross-transport dedup with `/post`.
- [ ] **1.5.4** Update relay README with the `/detect` shape and a curl
      example.

### Phase 2 — Build `detector/` from scratch (NOT a fork of rektSDK)

We write fresh TypeScript designed for **our** pipeline (HTTP webhook +
email fan-out, structured AI output, our config schema). rektSDK is read
ONLY for `otomato-sdk` API patterns (TRIGGERS, ACTIONS namespaces,
Workflow/Action/Edge construction) — those come from the third-party SDK
and we'd discover them by reading SDK docs anyway. No copy-paste of
Jerry's config, README, or domain logic.

- [ ] **2.1** New file `detector/package.json`:
       - `name`: `thatsrekt-detector`
       - `type: "module"`, `private: true`
       - exact-version-pinned deps: `otomato-sdk@2.0.0`, `dotenv@16.4.5`
       - dev deps: `tsx@4.15.0`, `typescript@5.4.5`, `@types/node@20.14.0`
       - scripts: `deploy`, `check`, `redeploy` (stops + redeploys),
         `clean`
- [ ] **2.2** New file `detector/tsconfig.json` — strict mode,
      `module: "ESNext"`, `target: "ES2022"`, `moduleResolution: "bundler"`,
      `noUncheckedIndexedAccess: true`, `exactOptionalPropertyTypes: true`.
- [ ] **2.3** New file `detector/src/config.ts` — Zod-validated config
      schema: `MonitoredAccount`, `Protocol`, `DetectorConfig`. Loads
      `tracking.json` + env. `webhookUrl`, `webhookToken`, `chains[]`,
      `alertEmail`, `protocols[]`, `monitoredAccounts[]`. Fail-fast on
      validation error.
- [ ] **2.4** New file `detector/src/prompt.ts` — exports a single function
      `buildDetectionPrompt(protocol)` returning the AI prompt string. AI
      schema lives here in one place, used by both prompt construction
      and (later) test fixtures.
- [ ] **2.5** New file `detector/src/workflow.ts` — `buildWorkflow(config)`
      constructs the Otomato Workflow (triggers, split, AI, IF, parallel
      Email + HTTP_REQUEST terminals). Pure function; no side effects.
- [ ] **2.6** New file `detector/src/deploy.ts` — entrypoint for
      `npm run deploy`: loads config, builds workflow, calls
      `workflow.create() + workflow.run()`, writes
      `workflow-ids.local.json`.
- [ ] **2.7** New file `detector/src/check.ts` — entrypoint for
      `npm run check`: loads workflow id from `workflow-ids.local.json`,
      reports state.
- [ ] **2.8** New file `detector/tracking.json` — our config: monitored
      accounts list (pull current 8 from Jerry's prod for v1 parity),
      18 protocols (names + keywords + twitter handle, no contract stubs
      since we're not using them yet), `webhookUrl: "REPLACE_AT_DEPLOY"`,
      `chains: ["anvil-eth"]`, `alertEmail: "jerry@karpatkey.com"`.
- [ ] **2.9** New file `detector/.env.example` — keys: `OTOMATO_API_KEY`,
      `OTOMATO_API_URL`, `WEBHOOK_URL`, `WEBHOOK_TOKEN`. (env overrides
      tracking.json fields.)
- [ ] **2.10** New file `detector/README.md` — what it is, prereqs (API
      key + ngrok), how to deploy, link to plan + relay-server-design.
- [ ] **2.11** Update root `.gitignore`: `detector/node_modules`,
      `detector/.env`, `detector/workflow-ids.local.json`,
      `detector/dist`.
- [ ] **2.12** `cd detector && npm install && npx tsc --noEmit` clean.

### Phase 3 — AI prompt v2 (structured JSON)

- [ ] **3.1** Rewrite the AI action's prompt in `create-alert-workflow.ts`.
      New schema (output is a JSON string the AI emits as text):
       ```
       {"hacked": bool,
        "title": string (≤200 bytes; concise human description, NEVER a
                         verbatim tweet quote),
        "attackers": string[] (0x-prefixed, only if VERBATIM in tweet),
        "victims": string[]   (0x-prefixed, only if VERBATIM in tweet)}
       ```
      Conservative guardrails baked into the prompt:
       - "If you cannot verify an address came verbatim from the tweet,
         return an empty array. Do not invent or normalize addresses."
       - "Title must be ≤200 bytes. If you cannot fit, truncate."
       - "Return ONLY the JSON object. No prose, no markdown fences."
- [ ] **3.2** Change the IF gate from `eq 'true'` to substring contains
      `'"hacked":true'`. Otomato cannot strict-parse JSON; the relay does
      that. False positives at this gate are fine — relay rejects them.
- [ ] **3.3** Drop the per-protocol `keywordList` / `handleClause` from
      the prompt? — keep them, they reduce false positives at the AI
      step. The structured-JSON requirement is additive.

### Phase 4 — HTTP_REQUEST + N email recipients per branch

The webhook is the on-chain alert channel; emails are kept for human
visibility (Jerry already relies on them; team wants copies). All
terminals fire in parallel from the same IF gate.

- [ ] **4.1** Per branch, terminals from `IF(hacked)`:
       - 1× `ACTIONS.CORE.HTTP_REQUEST.HTTP_REQUEST` → relay `/detect`
       - N× `ACTIONS.NOTIFICATIONS.EMAIL.SEND_EMAIL` (one per address
         in `tracking.json#alertEmails`). Otomato's `to` is a single
         address; we fan out by adding one action per recipient.
- [ ] **4.2** Construct the HTTP action targeting the relay's `/detect`
      endpoint (Otomato adapter shape — see `tasks/relay-server-design.md`):
       - `url`: `${WEBHOOK_BASE_URL}/detect`
       - `method`: `"POST"`
       - `headers` (object literal, Otomato substitutes variable refs):
         ```
         Authorization:     Bearer ${WEBHOOK_TOKEN}
         Content-Type:      application/json
         X-Idempotency-Key: ${tweetId}
         X-Tweet-URL:       ${tweetURL}
         X-Tweet-Account:   ${account}
         X-Tweet-Timestamp: ${timestamp}
         X-Chain:           ${WEBHOOK_CHAIN}
         ```
       - `body`: `${aiResult}` — the AI's structured JSON output
         verbatim. The relay parses, validates, and submits.
- [ ] **4.3** Edge graph: `trigger → split → AI → IF(hacked) → HTTP_REQUEST`.
      One terminal per branch.
- [ ] **4.4** `tracking.json` does NOT carry `alertEmail` — it's gone.
      The webhook config lives in env: `WEBHOOK_BASE_URL`, `WEBHOOK_TOKEN`,
      `WEBHOOK_CHAIN`.

### Phase 5 — Local stack orchestration (DONE)

Extended `ops/Makefile` rather than creating a parallel root Makefile.
New scripts live under `ops/scripts/`.

- [x] **5.1** Documented ngrok install in `detector/README.md`.
      `start-ngrok.sh` checks for the binary and gives a clean install
      pointer if missing.
- [x] **5.2** New Makefile targets: `relay-up`, `relay-down`,
      `relay-logs`, `ngrok-up`, `ngrok-down`, `ngrok-url`,
      `detector-env`, `detector-up`, `detector-down`, `detector-deploy`,
      `detector-smoke`. `detector-up` is the composite that boots the
      detector pipeline on top of an existing `lan-up` + `anvil-bootstrap`.
- [x] **5.3** `ops/scripts/`:
       - `whitelist-relay-eoa.sh` — generates a fresh EOA, funds it,
         impersonates the timelock to whitelist. Idempotent.
       - `start-relay.sh` — reads `.deployed.anvil-eth.json`, boots the
         Go relay with the right env, waits for `/healthz`. Idempotent.
       - `start-ngrok.sh` — launches ngrok pointed at the relay, polls
         the ngrok local API for the public URL. Idempotent.
       - `write-detector-env.sh` — writes `detector/.env` from the
         current ngrok URL + relay token, preserving any existing
         `OTOMATO_API_KEY`.
       - `smoke-test-detector.sh` — curl-based end-to-end test of the
         relay pipeline (covers Phase 6's automated portion).

### Phase 6 — Smoke test

- [ ] **6.1** Start `make local-detector`, capture ngrok URL.
- [ ] **6.2** Set `detector/.env` with `OTOMATO_API_KEY`, `WEBHOOK_URL`,
      `WEBHOOK_TOKEN`. Run `cd detector && npm run create`.
- [ ] **6.3** Verify in Otomato UI that the new workflow is `active`.
- [ ] **6.4** Trigger a manual test execution from Otomato (paste a
      real tweet that should fire) OR wait for a real tweet to fire.
- [ ] **6.5** Verify:
       - relay logs show `http POST /detect` with the AI body + headers
       - submitter logs show tx submitted to anvil-eth
       - frontend at `localhost:5173` renders the post (title comes from
         the AI's JSON, not the tweet text)
       - **NO email** is sent (we deliberately removed that branch — if
         email goes out something is wrong)
- [ ] **6.6** Capture each layer's logs for the operator's review.

### Phase 7 — Commit + open PR

- [ ] **7.1** Conventional commits per phase (`feat(relay): http post
      transport`, `feat(detector): vendor and rename rektSDK`, etc.).
- [ ] **7.2** Author = `bauti.eth` (per-commit, never global).
- [ ] **7.3** Open PR `bauti/detector-integration → master` once user gives
      green light.

---

## Open questions / known caveats

- **Otomato body templating syntax**: I have not verified the exact
  template syntax for interpolating `ai.getOutputVariableName('result')`
  into a JSON body. There's an existing precedent in the rektSDK email
  body (uses bare `${tweetContentVar}` strings). Worth a quick docs
  scan during Phase 4. If templating doesn't work, fallback: have the
  AI also include the raw tweet content + url in its JSON, and we
  body: just the AI result verbatim.

- **Idempotency id**: relay uses `envelope.id` for the 15-min dedup
  ring. Best stable id per tweet is the tweet URL itself (or a hash of
  it). Otomato exposes `tweetURL` from the trigger output. Use it
  directly.

- **`attacked_at`**: relay rejects `0`. Otomato has helper actions for
  current timestamp. If unavailable, embed `Date.now()/1000` at deploy
  time → all posts get the same attacked_at. That's wrong — they need
  per-execution timestamps. Investigate Otomato `HELPER.TIMESTAMP`.

- **Address cap redesign**: empty `attackers[]` and `victims[]` are
  allowed in v1.1; v1 detector takes advantage of this. The cap-redesign
  doc (`tasks/address-cap-redesign.md`) is independent.

- **Multichain**: this PR ships single-chain detector. Multi-chain
  fan-out is relay sub-phase B.

---

## Risks

- **Touching prod Otomato workflow.** Operator decision was Option B —
  we redeploy Jerry's workflow with HTTP_REQUEST added. Any deploy stops
  the previous workflow first (rektSDK README says so). Brief outage
  window during redeploy. Email pipeline restored on first successful
  detection of the new workflow.
- **ngrok URL rotation.** Free-tier ngrok URL changes on each restart,
  which means re-running `npm run create` to update the workflow with
  the new URL. Acceptable for local dev; if it bites, get paid ngrok
  or switch to cloudflared with a stable hostname.
- **AI cost doubling.** Adding HTTP_REQUEST as a parallel terminal does
  not double AI cost — the AI runs once per branch; both terminals
  fan-out from the IF gate.
