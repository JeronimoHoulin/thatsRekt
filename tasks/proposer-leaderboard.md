# Proposer leaderboard — soon-to-do (parking note)

**Status:** Idea, not started. Captured 2026-04-28 mid-detector-integration so
we don't lose it. Implementation comes after the detector PR ships.

## What

A leaderboard ranking whitelisted **posters** by the lifetime sum of votes
on posts they created. The post's voter karma already lives elsewhere
(`attackerScore`); this is the **author-side** mirror.

- **Score = Σ (upvotes − downvotes)** over every post the author created.
- Negative scores are valid and expected (low-quality posters get
  punished by the network).
- All-time, no decay (v1).
- Retracted posts: include their pre-retraction score, OR exclude entirely
  — open question, decide at design time. Excluding is more punishing
  (incentive to retract honestly), including preserves history.
- Surface on the Contributors page (already exists per chain) with a new
  "leaderboard" column. Sortable.

## Where the work lives

- **Indexer (`indexer/src/`):** add a `proposerScore` aggregate, derived
  from existing `Vote` + `Post` entities. Two options:
   1. New `Proposer` entity keyed by address+chain, updated on every
      `Vote.handle()`. Cheap reads, more write-time work.
   2. Computed view via Mesh-side aggregation over existing `Post` +
      `Vote`. No new storage, slower reads, more Mesh complexity.
   Recommend option 1 — matches the existing `attackerScore` pattern.
- **Mesh (`mesh/src/`):** expose `proposerLeaderboard(chain, limit, offset)`
  that orders proposers by `score` desc with paging.
- **Frontend (`frontend/src/`):** new section on Contributors page
  (already chain-tabbed) showing the leaderboard. Can also add a
  "score" badge next to a poster's address everywhere it's rendered
  (PostCard, PostDetail, AddressLabel) — that's polish, not v1.

## Open questions

1. **Retracted posts** — include or exclude? See above.
2. **Score on edit?** Edits don't change votes, so author score is
   unaffected by amends/title changes/added-attackers. Cosmetic
   confirmation only.
3. **Cross-chain aggregation** — show per-chain leaderboards (matches
   current Contributors layout) or a global leaderboard? Per-chain is
   simpler and matches the current UI; global needs a separate Mesh
   resolver that fans across chains.

## Estimate

- Indexer: ~2-3h (mirrors `attackerScore` exactly — same shape).
- Mesh: ~1h (one resolver, one query).
- Frontend: ~2h (use existing Contributors page tabs, add a leaderboard
  section + sortable table).
- Total: ~half a day for v1, single PR.

Defer until after the detector-integration branch lands.
