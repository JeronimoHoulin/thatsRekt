/**
 * Unit tests for the mesh chain registry.
 *
 * Asserts structural invariants and per-chain identity so that typos in
 * chainId/slug/endpoint don't slip through undetected.
 */
import { describe, expect, test } from 'bun:test'
import { CHAINS } from '../src/chains.ts'
import type { ChainEntry } from '../src/chains.ts'

// ---------------------------------------------------------------------------
// Registry invariants
// ---------------------------------------------------------------------------

describe('CHAINS registry invariants', () => {
  test('all slugs are unique', () => {
    const slugs = CHAINS.map((c) => c.slug)
    const unique = new Set(slugs)
    expect(unique.size).toBe(slugs.length)
  })

  test('all chainIds are unique', () => {
    const ids = CHAINS.map((c) => c.chainId)
    const unique = new Set(ids)
    expect(unique.size).toBe(ids.length)
  })

  test('all prefixes end with underscore', () => {
    for (const chain of CHAINS) {
      expect(chain.prefix.endsWith('_')).toBe(true)
    }
  })

  test('all endpoints start with http', () => {
    for (const chain of CHAINS) {
      expect(chain.endpoint.startsWith('http')).toBe(true)
    }
  })
})

// ---------------------------------------------------------------------------
// BSC entry (issue #118 — consumes graphql-bsc:4362 from #117)
// ---------------------------------------------------------------------------

describe('BSC chain entry', () => {
  const bsc = CHAINS.find((c) => c.slug === 'bsc') as ChainEntry | undefined

  test('bsc entry exists in the registry', () => {
    expect(bsc).toBeDefined()
  })

  test('chainId is 56 (BNB Chain mainnet)', () => {
    expect(bsc?.chainId).toBe(56)
  })

  test('slug is bsc', () => {
    expect(bsc?.slug).toBe('bsc')
  })

  test('prefix is Bsc_', () => {
    expect(bsc?.prefix).toBe('Bsc_')
  })

  test('default endpoint targets graphql-bsc:4362 (matches #117 graphql-bsc compose port)', () => {
    // When GRAPHQL_BSC_URL is unset, the fallback must point at the
    // graphql-bsc service on port 4362 — the port #117 exposes in prod.
    const originalEnv = process.env.GRAPHQL_BSC_URL
    delete process.env.GRAPHQL_BSC_URL

    // Re-import via dynamic import with a fresh module cache isn't
    // straightforward in bun:test without module isolation, so we inspect
    // the static default embedded in the entry directly by checking the
    // production hostname/port pattern matches the expectation.
    expect(bsc?.endpoint).toContain('graphql-bsc')
    expect(bsc?.endpoint).toContain('4362')
    expect(bsc?.endpoint).toContain('/graphql')

    if (originalEnv !== undefined) process.env.GRAPHQL_BSC_URL = originalEnv
  })

  test('registryAddress matches the v1.2.0 CREATE2 canonical proxy', () => {
    // Same address as Ethereum/Base/Arbitrum/Optimism — deployed via
    // CREATE2 to identical address on every supported chain.
    expect(bsc?.registryAddress).toBe('0xBfaEEE9662b4c037De24e5Caa65815350d57b89A')
  })
})
