/**
 * Component tests for ConfirmVoteButtons error display.
 *
 * Verifies that when `useConfirmPost` returns a non-null `error`, the
 * component renders an inline error string so the user can see what
 * went wrong (fixes the silent no-op regression described in #138).
 *
 * Wagmi hooks are mocked at the module level — we don't need a real
 * chain connection to verify the error display path.
 */
import { describe, expect, test, mock, beforeAll, afterAll } from 'bun:test'
import { render, screen, cleanup } from '@testing-library/react'
import React from 'react'

// ---------------------------------------------------------------------------
// Module-level mocks — must be set up before imports that transitively use
// wagmi. Bun's mock.module applies synchronously before the module graph
// resolves for this file.
// ---------------------------------------------------------------------------

// Track calls to useConfirmPost so tests can assert on hook return values.
let mockConfirmPostReturn: ReturnType<typeof defaultConfirmPostReturn>

function defaultConfirmPostReturn() {
  return {
    submit: mock(() => Promise.resolve(true)),
    reset: mock(() => undefined),
    hash: undefined as `0x${string}` | undefined,
    isBroadcasting: false,
    isMining: false,
    isSuccess: false,
    error: null as Error | null,
    isPending: false,
  }
}

mock.module('../src/hooks/useConfirmPost', () => ({
  useConfirmPost: (_chainId: number) => mockConfirmPostReturn,
}))

// Minimal wagmi mocks — the component also uses useAccount.
mock.module('wagmi', () => ({
  useAccount: () => ({
    address: '0xdeadbeef00000000000000000000000000000001',
    isConnected: true,
    chainId: 1,
  }),
  useWriteContract: () => ({
    writeContract: mock(() => undefined),
    data: undefined,
    isPending: false,
    error: null,
    reset: mock(() => undefined),
  }),
  useWaitForTransactionReceipt: () => ({
    isLoading: false,
    isSuccess: false,
    error: null,
  }),
  useSwitchChain: () => ({
    switchChainAsync: mock(() => Promise.resolve()),
  }),
}))

// Mock hooks that touch chain / contract state — not needed for error display
// test, return stable empty values.
mock.module('../src/hooks/useIsWhitelisted', () => ({
  useIsWhitelisted: () => ({ isWhitelisted: true, isLoading: false }),
}))

mock.module('../src/hooks/useUserVote', () => ({
  useUserVote: () => ({
    direction: 0,
    isUp: false,
    isDown: false,
    refetch: mock(() => Promise.resolve()),
  }),
}))

mock.module('@tanstack/react-query', () => ({
  useQueryClient: () => ({
    invalidateQueries: mock(() => Promise.resolve()),
  }),
}))

// WhitelistGateModal — not relevant to error display tests.
mock.module('../src/components/WhitelistGateModal', () => ({
  WhitelistGateModal: () => null,
}))

// ---------------------------------------------------------------------------
// Import component AFTER mocks are registered.
// ---------------------------------------------------------------------------
// eslint-disable-next-line import/first
import { ConfirmVoteButtons } from '../src/components/ConfirmVoteButtons'

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

function renderButtons(props: Partial<Parameters<typeof ConfirmVoteButtons>[0]> = {}) {
  return render(
    React.createElement(ConfirmVoteButtons, {
      chainId: 1,
      postId: 7n,
      upCount: 3,
      downCount: 1,
      posterAddress: '0x0000000000000000000000000000000000000002',
      ...props,
    }),
  )
}

beforeAll(() => {
  mockConfirmPostReturn = defaultConfirmPostReturn()
})

afterAll(() => {
  cleanup()
})

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

describe('ConfirmVoteButtons — error display', () => {
  test('renders no error element when error is null', () => {
    mockConfirmPostReturn = { ...defaultConfirmPostReturn(), error: null }
    renderButtons()
    expect(screen.queryByTestId('vote-error')).toBeNull()
    cleanup()
  })

  test('renders inline error when useConfirmPost returns a non-null error', () => {
    const err = Object.assign(new Error('nonce too low'), {
      shortMessage: 'nonce too low',
    })
    mockConfirmPostReturn = { ...defaultConfirmPostReturn(), error: err }

    renderButtons()

    const errorEl = screen.getByTestId('vote-error')
    expect(errorEl).not.toBeNull()
    expect(errorEl.textContent).toContain('nonce too low')
    cleanup()
  })

  test('prefers shortMessage over message when available', () => {
    const err = Object.assign(new Error('ContractFunctionExecutionError: long message here'), {
      shortMessage: 'chain mismatch',
    })
    mockConfirmPostReturn = { ...defaultConfirmPostReturn(), error: err }

    renderButtons()

    const errorEl = screen.getByTestId('vote-error')
    expect(errorEl.textContent).toContain('chain mismatch')
    expect(errorEl.textContent).not.toContain('ContractFunctionExecutionError')
    cleanup()
  })

  test('falls back to error.message when shortMessage is absent', () => {
    const err = new Error('user rejected transaction')
    mockConfirmPostReturn = { ...defaultConfirmPostReturn(), error: err }

    renderButtons()

    const errorEl = screen.getByTestId('vote-error')
    expect(errorEl.textContent).toContain('user rejected transaction')
    cleanup()
  })

  test('error element has role=alert for screen readers', () => {
    const err = new Error('tx failed')
    mockConfirmPostReturn = { ...defaultConfirmPostReturn(), error: err }

    renderButtons()

    const errorEl = screen.getByRole('alert')
    expect(errorEl).not.toBeNull()
    cleanup()
  })
})
