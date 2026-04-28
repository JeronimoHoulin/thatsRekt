# Rename: votes → confirmations

**Status:** Idea, not started. Captured 2026-04-28 mid-detector-integration.
Defer until after the detector PR ships.

## What

Rename the "vote" concept everywhere to "confirmation" / "confirm".
Mechanism is unchanged — same upvote/downvote semantics, same on-chain
events, same scoring. The change is purely user-facing nomenclature.

The reframing: a confirmation is "I agree this is a real incident", a
disconfirmation is "I think this is wrong / not real". This reads more
naturally for a hack-alert registry than the generic "voting" frame.

## Where the work lives

- **Contract (`contracts/src/ThatsRekt.sol`):** rename `Voted` →
  `Confirmed`, `vote()` / `retract()` → `confirm()` / `unconfirm()`,
  `attackerScore` → still scores but conceptually it's "confirmation
  count", maybe rename to keep things honest (not strictly needed).
  Decide whether to ALSO rename storage fields (breaks any direct
  consumers of `attackerScore`) or only event/function names. Likely
  only the function/event names — the storage layout migration is more
  surgery than the rename is worth.
- **Indexer (`indexer/src/processors/`):** new event names if we change
  them at the contract level. Also rename `Vote` entity → `Confirmation`
  in the schema.
- **Mesh (`mesh/src/`):** rename `votes` query / `Vote` type → those.
- **Frontend (`frontend/src/`):** every "upvote" / "downvote" / "vote"
  → "confirm" / "disconfirm" / "confirmation". UI labels, toast
  messages, hover text.
- **Relay (`relay/`):** no changes — the relay only does `post.create`.
- **Detector (`detector/`):** no changes.

## Migration path

This is a **v2 contract** kind of change. Two paths:

1. **In-place rename without contract migration.** Keep contract
   functions/events as-is (Voted, vote, retract). Rename ONLY in the
   indexer schema, mesh, and frontend. Pros: no contract upgrade.
   Cons: people reading the contract see "vote" but the UI says
   "confirm" — confusing.
2. **Contract upgrade.** UUPS upgrade ships new function selectors
   (confirm/unconfirm). Old vote() still callable but deprecated.
   Indexer + mesh + frontend rename for free since events change.
   Pros: clean naming end-to-end. Cons: timelocked upgrade ceremony.

Recommend option 2 — the contract is the source of truth and the
naming should match what users read on-chain. Bundle this with any
other contract changes (e.g. address-cap-redesign) so we only do the
upgrade ceremony once.

## Open questions

1. **Score nomenclature.** "attackerScore" implies voting. Rename to
   "attackerConfirmationScore" or "attackerNetConfirmations"? Or keep
   "attackerScore" (briefer)? Probably keep — the score IS just a
   number; "confirmation" is the action that produces it.
2. **Negative confirmations.** A "downvote" in the new frame is a
   "disconfirmation". Two function names (`confirm` / `disconfirm`)
   vs. one with a direction enum? Match the current `vote(direction)`
   pattern → one function `confirm(direction)`. Cleaner, less surface.
3. **Backwards compatibility.** Do we keep the old `vote` selector
   live for a deprecation window so anything using the old ABI still
   works? Probably yes — UUPS lets us add new functions without
   removing old ones.

## Estimate

- Contract: ~1h (rename + tests + upgrade dry-run)
- Indexer schema migration: ~2h
- Mesh: ~1h
- Frontend: ~3h (lots of small UI strings)
- Total: ~half a day, but bundled with the address-cap-redesign for
  a single upgrade ceremony makes it ~1 day.
