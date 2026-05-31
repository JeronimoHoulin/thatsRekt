/**
 * Unit tests for tokenAllowlist — pure module, no I/O.
 * Written test-first (TDD).
 *
 * Slice #207 additions: ERC20 allowlist lookups, erc20Addresses(), TRANSFER_TOPIC0.
 */
import { describe, expect, test } from 'bun:test'
import {
  allowlistFor,
  isAllowed,
  nativeFloor,
  tokenMeta,
  erc20Addresses,
  TRANSFER_TOPIC0,
} from '../src/tokenAllowlist.ts'

describe('allowlistFor', () => {
  test('returns non-null for Ethereum mainnet (chainId 1)', () => {
    expect(allowlistFor(1)).not.toBeNull()
  })

  test('returns null for unknown chain', () => {
    expect(allowlistFor(999999)).toBeNull()
  })

  test('Ethereum allowlist has native entry with ETH symbol', () => {
    const list = allowlistFor(1)
    expect(list?.native.symbol).toBe('ETH')
    expect(list?.native.decimals).toBe(18)
  })

  test('Ethereum nativeFloorWei is greater than 0', () => {
    const list = allowlistFor(1)
    expect(list!.nativeFloorWei).toBeGreaterThan(0n)
  })

  test('Ethereum allowlist has exactly 9 ERC20 entries', () => {
    const list = allowlistFor(1)
    expect(Object.keys(list!.erc20)).toHaveLength(9)
  })
})

describe('isAllowed', () => {
  test('native (null tokenAddress) is allowed on Ethereum', () => {
    expect(isAllowed(1, null)).toBe(true)
  })

  test('unknown token address is NOT allowed on Ethereum', () => {
    expect(isAllowed(1, '0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef')).toBe(false)
  })

  test('native is NOT allowed on unknown chain', () => {
    expect(isAllowed(999999, null)).toBe(false)
  })

  // ERC20 allowlist — canonical addresses (checksummed form — isAllowed lowercases internally)
  test('USDC (0xA0b86991...) is allowed on Ethereum', () => {
    expect(isAllowed(1, '0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48')).toBe(true)
  })

  test('USDT (0xdAC17F95...) is allowed on Ethereum', () => {
    expect(isAllowed(1, '0xdAC17F958D2ee523a2206206994597C13D831ec7')).toBe(true)
  })

  test('DAI (0x6B175474...) is allowed on Ethereum', () => {
    expect(isAllowed(1, '0x6B175474E89094C44Da98b954EedeAC495271d0F')).toBe(true)
  })

  test('WETH (0xC02aaA39...) is allowed on Ethereum', () => {
    expect(isAllowed(1, '0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2')).toBe(true)
  })

  test('WBTC (0x2260FAC5...) is allowed on Ethereum', () => {
    expect(isAllowed(1, '0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599')).toBe(true)
  })

  test('stETH (0xae7ab965...) is allowed on Ethereum', () => {
    expect(isAllowed(1, '0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84')).toBe(true)
  })

  test('wstETH (0x7f39C581...) is allowed on Ethereum', () => {
    expect(isAllowed(1, '0x7f39C581F595B53c5cb19bD0b3f8dA6c935E2Ca0')).toBe(true)
  })

  test('LINK (0x514910771...) is allowed on Ethereum', () => {
    expect(isAllowed(1, '0x514910771AF9Ca656af840dff83E8264EcF986CA')).toBe(true)
  })

  test('AAVE (0x7Fc66500...) is allowed on Ethereum', () => {
    expect(isAllowed(1, '0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9')).toBe(true)
  })

  test('lowercased address variant is also allowed (case-insensitive lookup)', () => {
    expect(isAllowed(1, '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48')).toBe(true)
  })
})

describe('nativeFloor', () => {
  test('Ethereum floor is positive (100000000000000 wei)', () => {
    const floor = nativeFloor(1)
    expect(floor).toBe(100_000_000_000_000n)
  })

  test('unknown chain returns 0n (fail-open)', () => {
    expect(nativeFloor(999999)).toBe(0n)
  })
})

describe('tokenMeta', () => {
  test('native token meta on Ethereum: ETH, 18 decimals', () => {
    const meta = tokenMeta(1, null)
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('ETH')
    expect(meta!.decimals).toBe(18)
  })

  test('unknown ERC20 on Ethereum returns null', () => {
    expect(tokenMeta(1, '0xdeadbeefdeadbeefdeadbeefdeadbeefdeadbeef')).toBeNull()
  })

  test('unknown chain returns null', () => {
    expect(tokenMeta(999999, null)).toBeNull()
  })

  // Per-token decimals — the syrupUSDC lesson: always verify, never assume 18.
  test('USDC decimals = 6 (verified on-chain)', () => {
    const meta = tokenMeta(1, '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48')
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('USDC')
    expect(meta!.decimals).toBe(6)
  })

  test('USDT decimals = 6 (verified on-chain)', () => {
    const meta = tokenMeta(1, '0xdac17f958d2ee523a2206206994597c13d831ec7')
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('USDT')
    expect(meta!.decimals).toBe(6)
  })

  test('DAI decimals = 18 (verified on-chain)', () => {
    const meta = tokenMeta(1, '0x6b175474e89094c44da98b954eedeac495271d0f')
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('DAI')
    expect(meta!.decimals).toBe(18)
  })

  test('WETH decimals = 18 (verified on-chain)', () => {
    const meta = tokenMeta(1, '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2')
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('WETH')
    expect(meta!.decimals).toBe(18)
  })

  test('WBTC decimals = 8 (verified on-chain; NOT 18)', () => {
    const meta = tokenMeta(1, '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599')
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('WBTC')
    expect(meta!.decimals).toBe(8)
  })

  test('stETH decimals = 18 (verified on-chain)', () => {
    const meta = tokenMeta(1, '0xae7ab96520de3a18e5e111b5eaab095312d7fe84')
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('stETH')
    expect(meta!.decimals).toBe(18)
  })

  test('wstETH decimals = 18 (verified on-chain)', () => {
    const meta = tokenMeta(1, '0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0')
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('wstETH')
    expect(meta!.decimals).toBe(18)
  })

  test('LINK decimals = 18 (verified on-chain)', () => {
    const meta = tokenMeta(1, '0x514910771af9ca656af840dff83e8264ecf986ca')
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('LINK')
    expect(meta!.decimals).toBe(18)
  })

  test('AAVE decimals = 18 (verified on-chain)', () => {
    const meta = tokenMeta(1, '0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9')
    expect(meta).not.toBeNull()
    expect(meta!.symbol).toBe('AAVE')
    expect(meta!.decimals).toBe(18)
  })
})

describe('erc20Addresses', () => {
  test('returns 9 addresses for Ethereum mainnet', () => {
    const addrs = erc20Addresses(1)
    expect(addrs).toHaveLength(9)
  })

  test('all addresses are lowercased', () => {
    const addrs = erc20Addresses(1)
    for (const addr of addrs) {
      expect(addr).toBe(addr.toLowerCase())
    }
  })

  test('all addresses start with 0x and have 42 characters', () => {
    const addrs = erc20Addresses(1)
    for (const addr of addrs) {
      expect(addr).toMatch(/^0x[0-9a-f]{40}$/)
    }
  })

  test('returns empty array for unknown chain', () => {
    expect(erc20Addresses(999999)).toHaveLength(0)
  })

  test('USDC address present in Ethereum list', () => {
    const addrs = erc20Addresses(1)
    expect(addrs).toContain('0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48')
  })

  test('WBTC address present in Ethereum list', () => {
    const addrs = erc20Addresses(1)
    expect(addrs).toContain('0x2260fac5e5542a773aa44fbcfedf7c193bc2c599')
  })
})

describe('TRANSFER_TOPIC0', () => {
  test('has correct keccak256 of Transfer(address,address,uint256)', () => {
    // keccak256("Transfer(address,address,uint256)") — canonical value
    expect(TRANSFER_TOPIC0).toBe(
      '0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef',
    )
  })

  test('is 66 characters (0x + 64 hex)', () => {
    expect(TRANSFER_TOPIC0).toHaveLength(66)
  })
})
