# Address-cap redesign — design brief

**Status:** Pre-implementation analysis. No code yet — locking the design first.
**Owner:** bauti
**Date:** 2026-04-27

The contract caps `attackers.length + victims.length` at **100** per post via `MAX_ADDRESSES_PER_POST`. The operator wants to push this higher and asks: how high can we theoretically go, and how?

---

## Where the cap comes from

The cap is **NOT** about post() gas. It's about **voting**. Look at `vote()`:

```solidity
uint256 aLen = p.attackers.length;
for (uint256 i; i < aLen; ++i) {
    attackerScore[p.attackers[i]] += delta;
}
```

Every vote on a post iterates the post's full attacker list and updates each one's running score. Same pattern in `_removePost` (retract) — reverses the entire array.

That O(attackers) per-vote cost is the binding constraint, not the O(addrs) cost of `post()` itself.

### Gas math (rough, 2026 prices)

| Operation | Gas cost | Notes |
|---|---|---|
| `post()` fixed overhead | ~100k | header SSTOREs, entry validation, list insert |
| Per attacker in `post()` | ~50k | array push (SSTORE) + appearances counter (SSTORE-cold first time) |
| Per victim in `post()` | ~50k | array push + active-post counter + isVictim flag |
| Per attacker in `vote()` | ~7k | SLOAD attacker addr + SSTORE updated score (warm) |
| Per attacker in `retract()` | ~7k | reverse of vote |
| Per address calldata | ~512 | 32 bytes × 16 gas/byte (non-zero) |

### Where current `MAX_ADDRESSES_PER_POST = 100` sits

| N | post() cost | vote() cost | retract() cost |
|---|---|---|---|
| 100 | ~5.1M | ~700k | ~700k |
| 200 | ~10.1M | ~1.4M | ~1.4M |
| 500 | ~25.1M | ~3.5M | ~3.5M |
| 1000 | ~50M (over block limit) | ~7M | ~7M |
| 2000 | impossible (over block) | ~14M | ~14M |

Block gas limits: Ethereum mainnet 30M, Base 60M (post-EIP), Arbitrum effectively much higher (no per-block cap, but pays for L1 calldata).

---

## The three options

### Option A — Just raise the cap (cheap to implement, painful to use)

Bump `MAX_ADDRESSES_PER_POST` to ~**500**.

**What changes:**
- Single constant in the contract.
- `post()` still fits in a block (~25M gas at N=500).
- `vote()` and `retract()` fit too (~3.5M gas).
- API unchanged. Integrators see no difference.

**The cost:**
- Voting on a max-size post burns 3.5M gas. At 30 gwei mainnet that's ~$300 per vote. Whitelisters won't vote on big posts; the voting signal flatlines for them.
- Post creation cost is similar — posters pay ~$2k to post a 500-addr alert on mainnet. Probably OK because posting big things is rare and important.

**Verdict:** acceptable up to ~500 but **economically degrades the voting signal**, which is the whole point of the karma model.

### Option B — Two-tier address lists (medium effort, clean separation)

Split the address arrays into "primary" (voted-on) and "extended" (associated, but not vote-scored).

```solidity
// Storage:
mapping(uint256 => address[]) public primaryAttackers;     // capped at e.g. 50
mapping(uint256 => address[]) public extendedAttackers;    // capped at e.g. 5000
// (same split for victims)

// On vote: only iterate primaryAttackers[].
// On addExtended: just append, no score impact.
```

**What changes:**
- New API: `addExtendedAttackers(postId, addr[])`, etc.
- New event: `ExtendedAttackersAdded(postId, ...)`.
- `attackerScore(addr)` only reflects primary listings. Add `attackerAppearances(addr)` count includes both. Add new view `attackerExtendedAppearances(addr)`.
- Integrators must opt into "treat extended addresses as suspect too" — gives them flexibility.

**The cost:**
- API surface roughly doubles for any address-related function.
- Posters need to think about which list an address belongs in (UX overhead).

**Verdict:** good model for "definitive attackers we want to give a score signal" vs "addresses associated with the incident, FYI". But it's a real product decision, not a numbers tweak.

### Option C — Lazy attacker scoring (clean, big behavior change)

Stop maintaining `attackerScore[addr]` as a running aggregate. Instead, derive it on-demand from the per-post netScores.

```solidity
// Storage:
mapping(uint256 => Post) _posts;                    // unchanged
mapping(address => uint256[]) _attackerPosts;       // post ids the addr appears in (NEW)

// vote() becomes O(1):
function vote(uint256 postId, VoteDirection direction) external {
    // ... validate ...
    int256 delta = ...;
    p.upvotes/downvotes += 1;
    // NO LOOP — just update the per-post counter.
}

// attackerScore() becomes O(posts containing this attacker):
function attackerScore(address a) external view returns (int256 score) {
    uint256[] memory posts = _attackerPosts[a];
    for (uint256 i; i < posts.length; ++i) {
        Post storage p = _posts[posts[i]];
        if (!p.removed) score += int256(uint256(p.upvotes)) - int256(uint256(p.downvotes));
    }
}
```

**What changes:**
- `vote()` is O(1) — voting on a 5,000-addr post costs the same as voting on a 1-addr post.
- `attackerScore(addr)` is O(N) where N = posts that include this attacker. For typical attackers that's 1-5 posts; for a notorious address that's maybe 50.
- New `_attackerPosts` mapping must be maintained on `post()`, `addAttackers()`, `retract()` — slight `post()` overhead increase (one more SSTORE per attacker), no big deal.

**The numerical effect:**

| N | post() cost | vote() cost | retract() cost |
|---|---|---|---|
| 100 | ~5.5M | ~50k | ~50k |
| 1,000 | ~55M (over block) | ~50k | ~50k |
| 5,000 | impossible | ~50k | ~50k |

Voting cost stops being the binding constraint. **The remaining constraint becomes `post()` itself** — you can't post more than ~500-600 addresses in a single tx because the block gas limit caps the post() call. Posts of 5,000+ addresses would need an iterative `addAttackers()` flow (one tx of post() with the headline attackers, then several `addAttackers()` follow-ups).

**The cost:**
- `attackerScore(addr)` view is no longer O(1). Integrators with hot-path checks need to cache or use the indexer.
- Slightly more storage per post (the `_attackerPosts` reverse index).
- Read-side complexity bumps from "constant" to "linear in posts touching this address". For 99% of addresses this is fine; for a few mega-popular attackers it could be a problem.

**Verdict:** the right architectural answer if we want the address cap to stop mattering for voting. The integrator-side cost is real but **the existing on-chain-readable design isn't load-bearing for any current integrator** — they all pull from Mesh/the indexer. The on-chain `attackerScore()` is a courtesy for future direct-integrator use.

### Option D — Event-only addresses (max scale, biggest API change)

Move attacker/victim address arrays out of storage entirely. They live only in `PostCreated` / `AttackersAdded` / `VictimsAdded` events. On-chain state holds:
- Per-post `netScore` (no per-attacker aggregate)
- Per-attacker `score` mapping, updated on vote (with delta lookup via event scan? — needs careful design)

**Verdict:** essentially a v2 contract. Out of scope for "raise the cap"; would be a from-scratch rewrite.

---

## Constraint table summary

| Approach | Vote cost @ 1k addrs | Post() cost @ 1k addrs | Read attackerScore | API change |
|---|---|---|---|---|
| Current | 7M (impossible at scale) | 50M (block-limit blocked) | O(1) | none |
| **A: raise cap** | 7M | 50M | O(1) | none — but reuses the same gas-capped patterns |
| **B: two-tier** | 350k (50 primary) | 50k + (50 × 50k) primary, extended cheap | O(1) | medium — doubles address-related surface |
| **C: lazy scoring** | 50k | 50M (still block-limited at 1k) | O(posts containing addr) | low — single function semantics change |
| **D: event-only** | 50k | (small) | indexer-only | high (v2 territory) |

---

## My recommendation: **Option C (lazy attacker scoring)**

Why:
1. **Voting cost stops scaling with address count**, which is the actual problem.
2. **`post()` becomes the only cap**, and posts of >500 addresses are a UX problem more than a data-model problem — addressing them via `addAttackers()` follow-ups is cleaner anyway (avoids monster posts).
3. **The on-chain `attackerScore()` O(1) guarantee was load-bearing for an integrator pattern that doesn't currently exist** — every integrator (existing and planned) reads through the indexer/Mesh. We pay a real complexity cost (`O(posts containing addr)`) for a feature nobody uses; let's drop it.
4. Storage layout is upgrade-compatible — add the new `_attackerPosts` mapping, leave the old `attackerScore` mapping for now (deprecated; could be removed in a later upgrade once we're sure no integrator depends on it).

### What "raise the cap to 5,000" looks like under Option C

- `MAX_ADDRESSES_PER_POST` becomes a soft guidance constant (or removed entirely — let `post()` revert with `OutOfGas` if you ask for too many).
- `addAttackers()` is the path for big alerts: post the headline + a primary set of 100 attackers, then follow-up txs add hundreds more in batches.
- The vote signal stays meaningful: voting on a post is cheap regardless of how many addresses it lists.

### Migration path

This is an upgrade, but a structurally simple one:

1. **New impl** with lazy-scoring math + `_attackerPosts` index.
2. **Migration script** populates `_attackerPosts` for existing posts on each chain (one-time pass over the indexer's data).
3. **Deprecate `attackerScore` mapping** as a stored field; switch the `attackerScore(addr)` view to compute lazily.
4. **Bump `MAX_ADDRESSES_PER_POST`** to whatever feels right — recommend **1,000** as initial cap (still block-bounded for `post()`, plenty of headroom for typical alerts).

The migration is timelocked, so integrators see "behavior unchanged but the cap moved up". Since no integrator depends on the O(1) read guarantee in practice, the deprecation is silent.

---

## Open questions for operator

1. **Are we confident no integrator uses `attackerScore(addr)` as an on-chain hot path** (i.e., calls it from another contract during a swap / liquidation / signature check)? If yes, Option C is fine. If no, Option B is the safer move.

2. **What's the target N?** 500? 1,000? 5,000? Different targets justify different effort levels:
    - N ≤ 500: Option A is cheapest.
    - N ≤ 5,000 with cheap voting: Option C.
    - N ≤ 50,000+: requires Option D (full v2 redesign).

3. **Should the cap be a soft ceiling (revert) or a soft guideline (let gas decide)?** Recommend keeping a hard ceiling — protects posters from accidentally creating an un-voteable post. But the value can be much higher than today.

---

## What this doc doesn't decide

- The actual new cap value (depends on operator answer to Q2).
- Whether to keep `attackerScore` storage as a cached optimization vs always lazy-compute.
- Specific timelock + migration ordering — mechanical, follow once the design is locked.

Implementation lands in a new phase once the operator picks an option.
