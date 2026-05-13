/**
 * Canonical EIP-712 typed-data fingerprint tests.
 *
 * Cross-repo drift sentinel: this file pins the typed-data hashes for a
 * fixed set of fixtures. The mesh backend has a parallel test
 * (`mesh/test/comments.test.ts`) that produces typed-data with the same
 * fixtures and recovers signers via `verifyTypedData` against the same
 * domain + types. Both halves must compute the same hash for a given
 * fixture, otherwise client-side signatures won't verify on the server.
 *
 * If this test starts failing, suspect a contract-shape change:
 *   - Domain fields (name / version / chainId / verifyingContract)
 *   - Type orderings (field order in COMMENT_TYPES)
 *   - Field types (e.g. uint256 vs uint64 for commentId)
 *
 * Bump the domain version on BOTH halves in lockstep. Never patch one
 * side in isolation.
 */
import { describe, expect, test } from 'bun:test'
import { hashTypedData } from 'viem'
import {
  buildCreateTypedData,
  buildEditTypedData,
  buildDeleteTypedData,
  COMMENT_TYPES,
} from '../src/lib/comments'

// Fixed fixtures — must match the mesh test fixtures byte-for-byte.
// `base-1` → chainId 8453, registry proxy is the v1.1.0 Base mainnet
// canonical address (also the same address on Optimism mainnet —
// cross-chain CREATE2). See `frontend/src/lib/contracts.ts`.
const FIX_POST_ID = 'base-1'
const FIX_BODY = 'gm'
// Must match the mesh fixture in `mesh/test/comments.test.ts`
// byte-for-byte — the cross-repo fingerprint sentinel relies on this.
const FIX_NEW_BODY = 'gm edited'
const FIX_COMMENT_ID = '42'
const FIX_SIGNED_AT = '2026-05-01T00:00:00.000Z'

const EXPECTED_DOMAIN = {
  name: 'thatsRekt',
  version: '1',
  chainId: 8453,
  verifyingContract: '0xBfaEEE9662b4c037De24e5Caa65815350d57b89A',
} as const

describe('canonical typed-data builders', () => {
  test('domain binds to chain + verifying contract from postId', () => {
    const td = buildCreateTypedData(FIX_POST_ID, FIX_BODY, FIX_SIGNED_AT)
    expect(td.domain).toEqual(EXPECTED_DOMAIN)
    expect(td.primaryType).toBe('CreateComment')
    expect(td.types).toEqual(COMMENT_TYPES)
    expect(td.message).toEqual({
      postId: FIX_POST_ID,
      body: FIX_BODY,
      signedAt: FIX_SIGNED_AT,
    })
  })

  test('edit typed data converts commentId string to bigint', () => {
    const td = buildEditTypedData(
      FIX_COMMENT_ID,
      FIX_POST_ID,
      FIX_NEW_BODY,
      FIX_SIGNED_AT,
    )
    expect(td.primaryType).toBe('EditComment')
    expect(td.message.commentId).toBe(42n)
    expect(td.message.postId).toBe(FIX_POST_ID)
    expect(td.message.newBody).toBe(FIX_NEW_BODY)
    expect(td.message.signedAt).toBe(FIX_SIGNED_AT)
  })

  test('delete typed data has no body field', () => {
    const td = buildDeleteTypedData(FIX_COMMENT_ID, FIX_POST_ID, FIX_SIGNED_AT)
    expect(td.primaryType).toBe('DeleteComment')
    expect(td.message.commentId).toBe(42n)
    expect(td.message.postId).toBe(FIX_POST_ID)
    expect(td.message.signedAt).toBe(FIX_SIGNED_AT)
    expect(td.message).not.toHaveProperty('body')
    expect(td.message).not.toHaveProperty('newBody')
  })
})

describe('canonical typed-data fingerprints', () => {
  // Backend fingerprints must match — see `mesh/test/comments.test.ts`.
  // Recorded values below were produced from the frontend builders
  // against the fixed fixtures + EXPECTED_DOMAIN. The mesh test is
  // expected to compute identical values via viem's hashTypedData.
  //
  // Regenerated 2026-05-13 for the v1.2.0 multichain deploy — the
  // verifyingContract moved from the v1.1.0 Base proxy to the new
  // canonical CREATE2 address 0xBfaEEE…b89A, so every hash here had
  // to be recomputed. Mesh's mirror test was updated to the same
  // values in the same commit; do not bump one side in isolation.

  test('CreateComment fixture hash is stable', () => {
    const td = buildCreateTypedData(FIX_POST_ID, FIX_BODY, FIX_SIGNED_AT)
    expect(hashTypedData(td)).toBe(
      '0x018722c91126010662d10f1306418de299ea1cb18d691e5d9890fd9b518317a6',
    )
  })

  test('EditComment fixture hash is stable', () => {
    const td = buildEditTypedData(
      FIX_COMMENT_ID,
      FIX_POST_ID,
      FIX_NEW_BODY,
      FIX_SIGNED_AT,
    )
    expect(hashTypedData(td)).toBe(
      '0x2952089c6480cfd3ae728288011dc57e993e2519c99477bda4959851ba76e53d',
    )
  })

  test('DeleteComment fixture hash is stable', () => {
    const td = buildDeleteTypedData(FIX_COMMENT_ID, FIX_POST_ID, FIX_SIGNED_AT)
    expect(hashTypedData(td)).toBe(
      '0xa2972bc91b80acf85f6c11f38a6cc4683e04be04ad0c5287e378eb65747bbde6',
    )
  })
})

describe('chain slug derivation', () => {
  test('rejects unknown chain slugs', () => {
    expect(() =>
      buildCreateTypedData('mars-1', FIX_BODY, FIX_SIGNED_AT),
    ).toThrow(/cannot derive chain slug/i)
  })

  test('matches longer slug first (base-sepolia, not base)', () => {
    // base-sepolia has chainId 84532 + a different verifyingContract;
    // if we accidentally matched `base-` first, the domain would point
    // at Base mainnet which is wrong.
    const td = buildCreateTypedData('base-sepolia-7', FIX_BODY, FIX_SIGNED_AT)
    expect(td.domain.chainId).toBe(84532)
    expect(td.domain.verifyingContract).toBe(
      '0x5278dD25e8551Cc98f2dC89791f5C89a9C83F695',
    )
  })
})
