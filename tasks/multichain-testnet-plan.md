# Multichain-Ready Testnet + LAN Stack — Design & Plan

**Status:** Design ready — operator approved 2026-04-27.
**Owner:** bauti
**Date:** 2026-04-27
**Predecessor PRs:** #4 (UUPS proxy), #5 (monorepo restructure), #6 (Subsquid indexer), #7 (frontend).
**Successor:** Implementation phases land as separate PRs (Phase 1 → Phase 7).

---

## 1. Goal

Run the full thatsRekt stack against **three chains in parallel** on a single host, presented to consumers as a **single multichain GraphQL endpoint** via a stitching gateway, while keeping each chain's indexer canonical (so it can lift to SQD Cloud unchanged):

1. **Anvil fork of Sepolia** — instant feedback for contract changes; no faucet, no rate limits.
2. **Sepolia testnet** — realistic public-chain conditions; shareable links.
3. **Base mainnet** — already-supported production target.

Expose the public surface (Mesh + frontend) on the **LAN** so a teammate on the same Wi-Fi can hit them.

**Non-goals (this plan):**
- Public-internet hosting / TLS / domains (deferred).
- IPFS / SQD Network deploy.
- Real Safe as governance on testnets — dummy EOA on Anvil + Sepolia. Mainnet stays Safe-governed.

---

## 2. Architecture

### 2.1 The single-host stack diagram

```
                          ┌────────────────────────────────────────┐
                          │   GraphQL Mesh gateway   :4350         │ ←── public
                          │   stitched + per-chain root fields     │     (LAN-bound)
                          └────────────────────────────────────────┘
                              ↑           ↑           ↑
              ┌───────────────┘           │           └───────────────┐
              │                           │                           │
   ┌──────────┴──────────┐  ┌─────────────┴───────────┐  ┌────────────┴────────┐
   │ squid:anvil         │  │ squid:sepolia            │  │ squid:base          │
   │   processor         │  │   processor              │  │   processor         │
   │   graphql  :4351    │  │   graphql  :4352         │  │   graphql  :4353    │   ←── internal only
   │   db: thatsrekt_anvil  │   db: thatsrekt_sepolia  │  │   db: thatsrekt_base│       (compose net,
   └─────────────────────┘  └──────────────────────────┘  └─────────────────────┘        not LAN)
              ↓                           ↓                           ↓
              └───────────────┬───────────┴───────────┬───────────────┘
                              ↓                       ↓
                       ┌──────────────────────────────────────┐
                       │  postgres :5432                       │ ←── 127.0.0.1 only
                       │  ├── thatsrekt_anvil                  │
                       │  ├── thatsrekt_sepolia                │
                       │  └── thatsrekt_base                   │
                       └──────────────────────────────────────┘
                              ↑
                    ┌─────────┘
   ┌────────────────┴─────┐
   │ anvil RPC :8545      │ ←── LAN-bound (fork of Sepolia)
   └──────────────────────┘

   ┌──────────────────────┐
   │ frontend dev :5173   │ ←── LAN-bound; talks to Mesh at :4350 only
   └──────────────────────┘
```

### 2.2 The decision rationale (locked)

| | Sovereign squids + Mesh + 1×pg/N×dbs | Shared DB + chainId | N independent stacks |
|---|---|---|---|
| Each squid is SQD-Cloud canonical | **✓** | ✗ | ✓ |
| Frontend sees one GraphQL endpoint | **✓** (via Mesh) | ✓ | ✗ |
| Failure isolation (one chain down ≠ all down) | **✓** (Mesh returns partial) | ✗ | ✓ |
| Per-chain canary deploys | **✓** | ✗ | ✓ |
| Schema stays single-chain clean | **✓** | ✗ (chainId everywhere) | ✓ |
| Cross-chain queries native | **✓** (Mesh stitch) | ✓ (SQL) | ✗ (off-chain fan-out) |
| One Postgres process to operate | **✓** (N logical dbs) | ✓ | ✗ (N pg) |
| Adding a 4th chain = config change only | **✓** | ✓ | ✓ |

### 2.3 Why one Postgres process with N databases (not N processes, not N schemas)

- **N databases (not N schemas):** Subsquid's `@subsquid/typeorm-store` connects via one `DATABASE_URL`. Database-level isolation is bulletproof; schema-switching via `search_path` is supported but flaky. This is also exactly what SQD Cloud gives each squid (one db per squid), so migration is a `pg_dump` per db → restore.
- **One process (not N):** the isolation we care about (one chain's bad migration can't touch another) is at the **database level**, not the **process level**. One pg process serving N databases gives us that for free; separate processes just add ops surface (3× memory, 3× backup volumes, 3× monitoring).
- **Init:** an `init.sql` baked into the pg container creates `thatsrekt_anvil`, `thatsrekt_sepolia`, `thatsrekt_base` on first boot. Idempotent (`CREATE DATABASE … IF NOT EXISTS`-equivalent pattern).

### 2.4 Why GraphQL Mesh

- Subsquid auto-generates GraphQL schemas from the entities — they're not federation-aware. Apollo Federation v2 would require wrapping each squid; not worth it.
- Mesh handles introspection-based stitching natively; ~50 LOC of YAML.
- Mesh exposes both **per-chain root fields** (`anvil.posts(...)`, `sepolia.posts(...)`, `base.posts(...)`) and **stitched unified queries** (`posts(...)` returning all chains, plus a `chain` field per result). The frontend gets to pick.
- Mesh is stateless — trivially replaceable, no data to lose if it crashes.

### 2.5 Mesh schema strategy

Each upstream squid's `Post`, `Whitelister`, etc. needs renaming to avoid collisions in the stitched schema. Two reasonable patterns; we go with **prefix transforms**:

```
squid:anvil   → AnvilPost,   AnvilWhitelister,   …
squid:sepolia → SepoliaPost, SepoliaWhitelister, …
squid:base    → BasePost,    BaseWhitelister,    …
```

Plus a Mesh-side **additional resolver** that fans out to all three for the unified `posts` query (returning a union or a normalized common shape with a `chain: Chain!` field). This keeps both per-chain-explicit and cross-chain queries first-class.

### 2.6 LAN-mode networking

| Service | Default | LAN-mode | Why |
|---|---|---|---|
| Mesh gateway | `0.0.0.0:4350` | same | The only public GraphQL surface. |
| Frontend dev server | `0.0.0.0:5173` | same | Public-good UI. |
| Anvil RPC | `0.0.0.0:8545` | same | Local-only chain; LAN access is the point. |
| Squid GraphQLs (`:4351-:4353`) | compose-internal only | unchanged | Mesh consumes them; nobody else should. |
| Postgres `:5432` | `127.0.0.1` only | unchanged | DB has no LAN business. |

**Discovery:** the frontend computes the Mesh endpoint as `http://${window.location.hostname}:4350/graphql`. Same build works for `localhost`, LAN, or future public hosting.

### 2.7 Chain registry — the single source of truth

Three parallel registries (one per layer that needs chain knowledge):

| File | Shape | Knows about |
|---|---|---|
| `indexer/src/chains.ts` | Backend | RPC URL env-var, gateway URL, finality, contract address env-var, start block env-var |
| `mesh/.meshrc.yaml` (config) | Mesh | Upstream GraphQL URL per chain, prefix transform |
| `frontend/src/lib/chains.ts` | Frontend | Display name, slug, chainId, badge color, explorer URL |

Keyed on numeric `chainId` so they stay in sync mechanically. A `// keep in sync …` comment marks the boundary; CI grep-check in Phase 7+ if drift becomes a concern.

### 2.8 The `CHAIN` env var (per processor service)

Each processor docker service reads `CHAIN` (e.g., `CHAIN=anvil`), looks up the registry entry, and configures itself. `processor.ts` becomes a `buildProcessor(chain)` factory; `main.ts` reads the env and runs.

```yaml
# excerpt: indexer/docker-compose.yml
processor-anvil:    { environment: { CHAIN: anvil,    DB_NAME: thatsrekt_anvil   } }
processor-sepolia:  { environment: { CHAIN: sepolia,  DB_NAME: thatsrekt_sepolia } }
processor-base:     { environment: { CHAIN: base,     DB_NAME: thatsrekt_base    } }
graphql-anvil:      { environment: { CHAIN: anvil,    DB_NAME: thatsrekt_anvil,    GQL_PORT: 4351 } }
graphql-sepolia:    { environment: { CHAIN: sepolia,  DB_NAME: thatsrekt_sepolia,  GQL_PORT: 4352 } }
graphql-base:       { environment: { CHAIN: base,     DB_NAME: thatsrekt_base,     GQL_PORT: 4353 } }
```

Three services from one image, three configs.

### 2.9 Anvil fork strategy (locked)

Fork from Sepolia (gives us the CREATE2 factory + Safe singleton factory + realistic state for free):

```bash
anvil --fork-url $SEPOLIA_RPC --host 0.0.0.0 --chain-id 31337 --block-time 2
```

`--chain-id 31337` (Anvil default) keeps it distinct from real Sepolia (11155111). Indexer config for Anvil is **RPC-only** (no Subsquid gateway — there's no archive for Anvil). `setRpcEndpoint(url).setFinalityConfirmation(0)` for instant indexing.

### 2.10 Dummy EOA governance for testnets

`Deploy.s.sol` requires `GOVERNANCE_OWNER` to be a contract — production-safe, blocker on testnets. A new **`DeployDev.s.sol`** mirrors it but takes an EOA owner and skips the contract-code check. Distinct salts (`thatsRekt.impl.dev.v1.0.0`, `thatsRekt.timelock.dev.v1`, `thatsRekt.proxy.dev`) so dev deploys can't collide with prod CREATE2 addresses.

Recommended dev EOA: **Anvil default account 0** (`0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266`, mnemonic `test test test test test test test test test test test junk`). Reproducible across machines, no secret to manage. Same EOA across Anvil + Sepolia → **same thatsRekt address on both** (convenient for parity testing). Mainnet's Safe-governed deploy will land at a different address — by design.

---

## 3. File Structure Changes

```
thatsRekt/
├── contracts/
│   └── script/
│       ├── Deploy.s.sol                      (unchanged — production / mainnet)
│       ├── DeployDev.s.sol                   (NEW — EOA owner, distinct salts)
│       └── anvil/                            (NEW)
│           ├── bootstrap.sh                  ← waits for Anvil, runs DeployDev, emits .deployed.json
│           ├── reset.sh                      ← tears down + replays
│           └── README.md
├── indexer/
│   ├── src/
│   │   ├── chains.ts                         (NEW — backend registry)
│   │   ├── processor.ts                      (REFACTORED — buildProcessor(chain) factory)
│   │   ├── main.ts                           (REFACTORED — reads CHAIN env)
│   │   └── handlers/                         (unchanged — single-chain handlers)
│   ├── schema.graphql                        (unchanged — single-chain shape per squid)
│   ├── db/                                   (unchanged)
│   ├── init.sql                              (NEW — creates 3 dbs on pg first boot)
│   ├── docker-compose.yml                    (UPDATED — pg + 3×proc + 3×gql + mesh)
│   ├── docker-compose.lan.yml                (NEW — overlay binding mesh + frontend to 0.0.0.0)
│   ├── docker-compose.anvil.yml              (NEW — adds anvil service)
│   ├── .env.example                          (UPDATED — per-chain RPC + contract + start block)
│   └── README.md                             (UPDATED)
├── mesh/                                     (NEW — GraphQL Mesh gateway)
│   ├── Dockerfile
│   ├── package.json
│   ├── .meshrc.yaml                          ← stitching config + prefix transforms + additional resolvers
│   ├── src/
│   │   └── resolvers/
│   │       └── unified-posts.ts              ← cross-chain fan-out resolver
│   └── README.md
├── frontend/
│   ├── src/
│   │   ├── lib/
│   │   │   ├── chains.ts                     (NEW — frontend registry)
│   │   │   ├── client.ts                     (UPDATED — endpoint from window.location)
│   │   │   └── queries.ts                    (UPDATED — Mesh root fields)
│   │   ├── hooks/
│   │   │   └── useChainFilter.ts             (NEW)
│   │   └── components/
│   │       ├── ChainSelector.tsx             (NEW — All / Anvil / Sepolia / Base)
│   │       └── ChainBadge.tsx                (NEW)
│   └── package.json                          (UPDATED — dev:lan)
├── ops/                                       (NEW)
│   ├── Makefile                              ← lan-up / lan-down / anvil-bootstrap / lan-info
│   └── README.md
├── tasks/
│   └── multichain-testnet-plan.md            (this file)
└── README.md                                  (UPDATED)
```

---

## 4. Phased Implementation

Each phase is a separate PR. Each leaves the codebase in a working state.

### Phase 1 — Backend chain registry + processor factory

**Goal:** indexer stops hard-coding Base; reads `CHAIN` env, picks config from registry. Still a single processor running in single-chain mode (regression-equivalent to today).

**Files:** `indexer/src/chains.ts`, refactor `processor.ts` to `buildProcessor(chain)`, refactor `main.ts` to read `CHAIN`. Update `.env.example`.

**Tests:**
- Unit: registry lookup throws on unknown slug, returns typed entry on known.
- Manual: `CHAIN=base RPC_BASE_HTTP=… docker compose up` matches today's behavior (regression).

### Phase 2 — Multi-squid compose (one image, three services, one Postgres, N databases)

**Goal:** the existing single-squid compose grows to: 1×postgres (with `init.sql` provisioning 3 dbs), 3×processor service (one per chain), 3×graphql service (one per chain). No Mesh yet; the frontend would point at one of the three squid endpoints directly.

**Files:** `indexer/init.sql`, expand `indexer/docker-compose.yml`. Per-chain `.env` blocks.

**Tests:**
- Manual: `CHAINS=base docker compose up postgres processor-base graphql-base` matches today (regression).
- Manual: `docker compose up` brings all three squids up; each squid's GraphQL serves its chain's data on its own internal port.
- Negative: `nc -vz localhost 4351` from the host fails (squid GraphQLs internal-only).

### Phase 3 — DeployDev + Sepolia deploy

**Goal:** thatsRekt deployed to Sepolia using the dummy EOA owner; indexer config wired up.

**Files:** `contracts/script/DeployDev.s.sol`, deploy wrapper script, fill in `.env` for Sepolia.

**Tests:**
- Forge: `DeployDev` produces a working proxy + timelock + impl with the EOA as owner; `multisig.code.length > 0` check is correctly removed/bypassed.
- Manual: thatsRekt deployed on Sepolia at the dev-CREATE2 address; events from a `cast send` show up in `graphql-sepolia` within seconds.

### Phase 4 — Anvil bootstrap

**Goal:** one command spins a forked Anvil + thatsRekt deployed + indexer config emitted.

**Files:** `contracts/script/anvil/bootstrap.sh`, `reset.sh`, `indexer/docker-compose.anvil.yml` overlay (adds `anvil` service).

**Tests:**
- Bootstrap is idempotent on a clean Anvil session.
- Manual: post a test event from the EOA via `cast send`, watch it surface in `graphql-anvil` within seconds. **Anvil and Sepolia share the dev-CREATE2 address** (parity check).

### Phase 5 — GraphQL Mesh gateway

**Goal:** Mesh stitches the three squid GraphQLs into one schema, exposed on `:4350`. Per-chain root fields (`anvil.posts(...)`) and stitched unified queries (`posts(...)` with `chain` field) both work.

**Files:** `mesh/` directory (Dockerfile, package.json, `.meshrc.yaml`, `src/resolvers/unified-posts.ts`), add `mesh` service to `indexer/docker-compose.yml`.

**Tests:**
- Manual: `mesh` service comes up after all three squid GraphQLs are healthy.
- Manual: query at `http://localhost:4350/graphql`:
  - `anvil { posts(limit: 5) { id ... } }` returns Anvil's posts.
  - `posts(limit: 5)` (unified) returns posts from all three chains, each with `chain { slug name chainId }`.
- Manual: kill `processor-sepolia`; `posts` still returns Anvil + Base (Mesh partial-result behavior).

### Phase 6 — LAN-mode overlay

**Goal:** one command exposes Mesh + frontend to the LAN; squid GraphQLs and Postgres stay internal.

**Files:** `indexer/docker-compose.lan.yml` (overlay; Mesh and frontend bind to `0.0.0.0`, others unchanged), `frontend/package.json` (`dev:lan` script), `ops/Makefile` (`lan-up`, `lan-down`, `lan-info`), `ops/README.md`.

**Tests:**
- Manual: from a separate LAN device, hit `http://<host-lan-ip>:5173`, frontend loads, GraphQL queries succeed against `http://<host-lan-ip>:4350/graphql`.
- Negative: `nc -vz <host-lan-ip> 5432` (postgres), `nc -vz <host-lan-ip> 4351` (squid graphql) both refused.

### Phase 7 — Frontend chain filter

**Goal:** frontend can scope by chain (or show all) without rebuild.

**Files:** `frontend/src/lib/chains.ts`, `client.ts`, `queries.ts` updates, `useChainFilter.ts`, `ChainSelector.tsx`, `ChainBadge.tsx`.

**Behavior:** dropdown in header — `All`, `Anvil`, `Sepolia`, `Base`. "All" = unified `posts` query. Per-chain selection = scoped query (e.g., `anvil { posts(...) }`). Selection persisted to `localStorage`. `ChainBadge` renders next to each post (essential when "All" is active).

**Tests:**
- Unit: hook persists / rehydrates from `localStorage`.
- Unit: `queries.ts` builds correct GraphQL document for All vs scoped.
- Manual: switch chains, see TanStack cache reset, see correct query (devtools).

### Phase 8 — End-to-end LAN verification

**Steps (manual playbook in `ops/README.md`):**
1. `make anvil-bootstrap` → Anvil up, thatsRekt deployed, env files written.
2. `make lan-up` → postgres + 3 processors + 3 squid graphqls + mesh up.
3. `cd frontend && pnpm dev:lan`.
4. From a phone on the same Wi-Fi: visit `http://<host-lan-ip>:5173`, see "All chains" feed mixing Anvil/Sepolia/Base, switch filters, confirm scoping.
5. Kill one processor; the unified feed continues serving the other two (Mesh failure isolation).

**Exit criteria:** all three chains indexed concurrently; one Postgres, three squid GraphQLs, one Mesh endpoint; frontend chain filter works; LAN access works without rebuild; killing one squid degrades gracefully.

---

## 5. Risks & Open Questions

| # | Risk | Mitigation |
|---|---|---|
| 1 | Mesh schema collisions if upstream squids share entity names. | Prefix transforms (`AnvilPost`, etc.) handle this mechanically. |
| 2 | Mesh additional-resolver for unified `posts` is hand-rolled — risk of subtle bugs (sort order, pagination across chains). | Phase 5 includes explicit tests for sort-by-timestamp-desc and limit-N across all three. Pagination cursor-based (timestamp + chainId tiebreaker), not offset-based. |
| 3 | Anvil cold-start has no CREATE2 factory. | Fork-from-Sepolia gives it for free. Cold-start is documented fallback. |
| 4 | Subsquid Sepolia/Base gateway lag on very recent blocks. | Auto-RPC fallback. |
| 5 | `chains.ts` (backend) / `.meshrc.yaml` / `chains.ts` (frontend) drift. | Phase 8+ optional CI grep check on chainSlug list. Mechanical drift only — no logic duplication. |
| 6 | Per-chain `pg_dump` for backups grows N-fold over time. | Each db starts at deploy block; total volume small. Logical backup, not physical. |
| 7 | Mesh adds a hop of latency. | Negligible for our query volume. |
| 8 | Vite dev server on LAN exposes HMR endpoints. | Acceptable for trusted-LAN dev; production deploy uses static `dist/`. |
| 9 | Dummy-EOA testnet deploy uses well-known mnemonic (Anvil default account 0). | Documented and intentional. The EOA is `0xf39Fd6…2266`; private key is the published Anvil test mnemonic. **Never use this on mainnet.** |

**Open questions (decide during implementation):**

1. **Mesh image: official or self-built Dockerfile?** Recommend self-built (control over node version, deps, alpine base).
2. **Mesh additional resolver in TypeScript or Mesh's YAML-defined resolver hooks?** Recommend TypeScript — easier to test, easier to evolve.
3. **Pagination strategy for unified `posts`?** Recommend cursor `(timestamp, chainId, postId)` lexicographic — stable across chain restarts and additions.

---

## 6. What This Plan Doesn't Do (intentionally)

- **No production deploy / TLS / domain.** LAN-only.
- **No real Safe on testnets.** EOA on Anvil + Sepolia. Mainnet keeps Safe.
- **No shared `chains` package across layers.** Three parallel registries; hoist when duplication actually hurts.
- **No CI changes (Phase 1–7).** Phase 8+ may add a chain-slug drift check.
- **No SQD Cloud deploy yet.** But every squid is shaped to lift to Cloud unchanged; Mesh self-hosts (it isn't an SQD Cloud thing).

---

## 7. Success Definition

This plan is "done" when:

1. `cd contracts/script/anvil && ./bootstrap.sh` produces a working Anvil with thatsRekt deployed via `DeployDev`.
2. `make lan-up` brings up: 1×postgres + 3×processor + 3×squid-graphql + 1×mesh + (optional) 1×anvil.
3. `http://<host-lan-ip>:4350/graphql` serves a unified schema with both `posts(...)` (all chains, sort-merged) and `anvil.posts(...)` / `sepolia.posts(...)` / `base.posts(...)` (per-chain).
4. Frontend at `http://<host-lan-ip>:5173`, accessed from a phone on the same Wi-Fi, shows the cross-chain feed by default and a working chain filter.
5. `ops/README.md` walks a fresh teammate from `git clone` to "the whole stack on my LAN" in under 30 minutes.
6. Killing any single processor leaves the other two chains' data still queryable through Mesh (graceful degradation).
7. **Adding a 4th chain is:** one entry in each of three registries (backend / mesh / frontend), one `.env` block, one routeme.sh key, two new compose services (processor + graphql) defined as parameterized copies. **No code changes.**

The seventh point is the real test: the design is good if and only if the next chain is config, not code.
