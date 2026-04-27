# thatsRekt Frontend — Implementation Plan

**Date:** 2026-04-27
**Status:** APPROVED (operator green-lit stack picks inline)
**Owner:** bauti

## 1. Goal

A simple, IPFS-compatible static web app that browses the thatsRekt registry. Two views:

1. **Feed** — list of recent (non-removed) posts, sorted by creation time. Click a card to drill into a post.
2. **Post detail** — full post fields + a chronological timeline merging Vote events and Edit events.

Not hosted on IPFS yet — but the build output must be **IPFS-compatible** so we can `ipfs add -r dist/` whenever ready.

## 2. Stack

| Concern | Choice |
|---------|--------|
| Build tool | Vite (predictable static output, no SSR baggage) |
| Framework | React 19 + TypeScript |
| Styling | Tailwind v4 |
| Routing | `react-router-dom` with `HashRouter` (IPFS gateway safe — no `index.html` fallback dependency) |
| Data | `@tanstack/react-query` for cache + state, `graphql-request` for fetch |
| GraphQL endpoint | env-configurable via `VITE_GRAPHQL_ENDPOINT`; default `http://localhost:4350/graphql` for local dev |

## 3. IPFS-compatibility rules (enforce throughout)

- `vite.config.ts` sets `base: './'` so all asset URLs are relative.
- HashRouter — no client routing depends on the gateway returning `index.html` for arbitrary paths.
- No SSR, no server-side rendering hooks. Pure CSR.
- All env config (GraphQL endpoint, etc.) **inlined at build time**. The deployed bundle is fully self-contained.
- No image optimization that requires a server. Use plain `<img>` tags or inline SVG.
- No analytics, no fonts loaded from external CDNs (or if used, document that IPFS deploys go offline if the CDN is down).

## 4. Repo layout

```
frontend/
├── README.md
├── package.json
├── tsconfig.json
├── vite.config.ts
├── postcss.config.js (if needed for Tailwind v4)
├── index.html
├── public/
│   └── (static assets — favicon, etc.)
├── src/
│   ├── main.tsx          (entry; mounts <App /> with QueryClientProvider + HashRouter)
│   ├── App.tsx           (route definitions)
│   ├── index.css         (Tailwind imports + minimal globals)
│   ├── lib/
│   │   ├── client.ts     (GraphQL client setup; reads VITE_GRAPHQL_ENDPOINT)
│   │   ├── queries.ts    (typed GraphQL queries as strings + their TS result types)
│   │   └── format.ts     (helpers: shortAddress, relativeTime, etc.)
│   ├── pages/
│   │   ├── Feed.tsx
│   │   └── PostDetail.tsx
│   └── components/
│       ├── PostCard.tsx
│       ├── Timeline.tsx
│       ├── VoteRow.tsx
│       ├── EditRow.tsx
│       ├── AddressLabel.tsx
│       └── EmptyState.tsx
└── tasks/
    └── frontend-plan.md
```

## 5. Pages

### Feed (`/`)

Header with title + brief tagline + the GraphQL endpoint info. Body: vertical list of `PostCard`s.

**Query:**
```graphql
{
  posts(orderBy: createdAtBlock_DESC, limit: 50, where: { removed_eq: false }) {
    id
    poster { id }
    attackedAt
    note
    upvotes
    downvotes
    netScore
    createdAtTimestamp
    attackerLinks { address { id attackerScore } }
    victimLinks { address { id } }
  }
}
```

`PostCard` shows: ID, poster (truncated), `attackedAt` as relative time, note (truncated to ~200 chars), netScore + up/down breakdown, attacker/victim address counts. Links to `/#/post/:id`.

Empty state: helpful message when no posts (likely the day-1 state — contract not deployed yet, or just no activity on the indexed chain).

### PostDetail (`/#/post/:id`)

**Query** — single post with full attacker/victim list + all votes + all edits:
```graphql
query Post($id: String!) {
  postById(id: $id) {
    id poster { id } attackedAt lastUpdatedAt note upvotes downvotes netScore removed createdAtTimestamp removedAtTimestamp
    attackerLinks { address { id attackerScore attackerAppearances } }
    victimLinks   { address { id isVictim } }
    votes { id voter { id } oldDirection newDirection blockNumber timestamp }
    edits { id kind newNote addedAttackers addedVictims blockNumber timestamp }
  }
}
```

Sections:
1. **Header** — postId, poster, attackedAt, lastUpdatedAt, removed-or-active badge, netScore.
2. **Note** — full note text, monospace-friendly.
3. **Attackers** — list with each address's `attackerScore` and appearances.
4. **Victims** — list with `isVictim` flag.
5. **Timeline** — merges `votes` + `edits` into a single chronological feed (oldest → newest), each row labeled with its kind. Renders `VoteRow` for vote events (showing direction transition `oldDir → newDir`) and `EditRow` for edits (kind + payload preview).

## 6. Local dev

```bash
cd frontend
pnpm install
cp .env.example .env  # set VITE_GRAPHQL_ENDPOINT if not http://localhost:4350/graphql
pnpm dev
# open http://localhost:5173
```

Assumes the indexer's docker stack is running (`cd indexer && docker compose up -d --build`).

## 7. Production build

```bash
pnpm build
# outputs to dist/

# Smoke test the static bundle locally:
pnpm preview
```

`dist/` is what gets `ipfs add -r dist/` whenever Phase 7 of indexer hosting is ready.

## 8. Out of scope

- IPFS pinning service (Pinata, Web3.Storage). Captured for future workstream — no automation here.
- ENS contenthash automation (`thatsrekt.eth` → IPFS CID).
- E2E tests against the indexer (manual smoke only — defer to dedicated test PR if needed).
- Auth / wallet connect (no write actions yet — read-only browser).
- Cross-chain UI (single-chain for now per indexer scope).
- Server-side rendering of any kind.
- Analytics or third-party tracking.

## 9. Decision log

- **Vite over Next.js**: Next's static export has gotchas on IPFS (trailing slashes, `next/image` server requirements, prefetching). Vite produces clean static output that's been battle-tested on IPFS.
- **HashRouter over BrowserRouter**: IPFS HTTP gateways (and ENS' `*.eth.limo` mirror) have varying handling for SPA fallback to `index.html`. Hash routing sidesteps the entire question.
- **graphql-request over Apollo Client**: 1.5KB vs 60KB+. We don't need Apollo's cache — TanStack Query's cache is enough.
- **No GraphQL codegen yet**: queries are short and few. Codegen adds tooling overhead. Revisit if the schema grows complex.

## 10. Workflow

Single PR via worktree (`bauti/frontend`). Commits roughly:

1. `chore(frontend): scaffold vite + react + ts + tailwind`
2. `feat(frontend): graphql client, queries, types`
3. `feat(frontend): feed page`
4. `feat(frontend): post detail page + timeline`
5. `docs(frontend): README quickstart + IPFS notes`

Build smoke-tested locally (`pnpm build` + `pnpm preview`) before pushing.

PR target: `master`. Operator merges.
