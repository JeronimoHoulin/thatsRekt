/**
 * Slug-coverage invariant tests.
 *
 * The frontend has two parallel views of "what chains are live":
 *   - `chains.ts::CHAINS` — UI registry. Entries with `liveIndexed: true`
 *     are the chains whose posts the live feed renders.
 *   - `queries.ts::SLUG_TO_PREFIX` — bridge map. The per-chain detail
 *     query `<Prefix>_postById(id:...)` looks up its GraphQL prefix
 *     here.
 *
 * These must stay in lockstep: every live-indexed chain needs a prefix
 * entry, or `/post/<slug>/<id>` 404s even though the unified `posts()`
 * feed returns the post fine. Same goes for the inverse — a prefix
 * entry pointing at a chain the UI doesn't render is dead config.
 *
 * This regression has now happened twice:
 *   - PR #111 (`base-sepolia` missing after v1.1.0 Base Sepolia deploy)
 *   - v1.2.0 cutover (`ethereum` + `arbitrum` missing after the
 *     multichain redeploy that landed the new mainnet/arb contracts)
 *
 * Both times mesh + indexer + `chains.ts` were updated in the same PR
 * but `SLUG_TO_PREFIX` was forgotten, so live-indexed posts surfaced in
 * the feed but couldn't be opened. This test fails loudly the next
 * time someone wires a new chain through `chains.ts` without also
 * registering its prefix.
 */
import { describe, expect, test } from 'bun:test'
import { CHAINS } from '../src/lib/chains'
import { SLUG_TO_PREFIX } from '../src/lib/queries'

describe('SLUG_TO_PREFIX coverage', () => {
  test('every liveIndexed chain has a SLUG_TO_PREFIX entry', () => {
    const missing = Object.values(CHAINS)
      .filter((c) => c.liveIndexed)
      .map((c) => c.slug)
      .filter((slug) => !(slug in SLUG_TO_PREFIX))
    expect(missing).toEqual([])
  })

  test('SLUG_TO_PREFIX has no orphans (every entry maps to a known chain)', () => {
    const orphan = Object.keys(SLUG_TO_PREFIX).filter(
      (slug) => !(slug in CHAINS),
    )
    expect(orphan).toEqual([])
  })

  test('every SLUG_TO_PREFIX entry points at a liveIndexed chain', () => {
    // Catch the inverse mistake: keeping a prefix entry for a chain
    // that's been demoted to archive-only. The detail page would then
    // attempt a `<Prefix>_postById` against a squid that no longer
    // exists. Better to fail loudly here than to ship a broken link.
    const archiveOnly = Object.keys(SLUG_TO_PREFIX).filter((slug) => {
      const chain = (CHAINS as Record<string, { liveIndexed: boolean } | undefined>)[slug]
      return chain !== undefined && chain.liveIndexed === false
    })
    expect(archiveOnly).toEqual([])
  })
})
