/**
 * Token allowlist for the donations indexer.
 *
 * Slice #205: Ethereum only, native ETH only.
 * Slice #207: Ethereum ERC20 allowlist added (~9 major tokens).
 * Later slices add additional chains.
 *
 * Design:
 * - Pure module — no I/O, no side effects. Testable in total isolation.
 * - `nativeFloor(chainId)` filters 1-wei spam; returns 0n for unknown chains
 *   (fail-open on unfamiliar chains — the processor validates chain before
 *   calling, so returning 0n for unknown is safe).
 * - `isAllowed(chainId, tokenAddress)` returns true for the native sentinel
 *   (null/undefined/'') and for any whitelisted ERC20. Returns false for
 *   unknown chains or unknown tokens.
 * - `tokenMeta(chainId, tokenAddress)` returns symbol + decimals for allowed
 *   tokens. Returns null for unknown tokens (processor skips them).
 * - `erc20Addresses(chainId)` returns all allowlisted ERC20 addresses (lowercased)
 *   for a given chain — used by the processor to register addLog subscriptions.
 *
 * The native sentinel is represented as null (tokenAddress === null).
 * ERC20 entries are lowercased addresses.
 *
 * ERC20 decimals/symbols were verified on-chain via cast call before being
 * committed here. See PR body for the exact cast commands and results.
 *
 * Transfer(address,address,uint256) topic0:
 *   keccak256("Transfer(address,address,uint256)")
 *   = 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef
 * Pre-computed and exported as TRANSFER_TOPIC0 to avoid runtime hashing.
 */

export interface TokenMeta {
  readonly symbol: string
  readonly decimals: number
}

export interface ChainAllowlist {
  /** Native coin entry. */
  readonly native: TokenMeta
  /** Dust floor in native base units (wei). Transfers below this are dropped. */
  readonly nativeFloorWei: bigint
  /** ERC20 allowlist: lowercased address -> meta. */
  readonly erc20: Readonly<Record<string, TokenMeta>>
}

/**
 * ERC-20 Transfer event topic0.
 * keccak256("Transfer(address,address,uint256)")
 * Verified with: cast keccak "Transfer(address,address,uint256)"
 */
export const TRANSFER_TOPIC0 =
  '0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef'

// ---------------------------------------------------------------------------
// Ethereum mainnet ERC20 allowlist (slice #207)
//
// Decimals and symbols verified on-chain (Ethereum mainnet, block latest)
// via cast call <address> "decimals()" --rpc-url <routeme-eth> and
//      cast call <address> "symbol()" --rpc-url <routeme-eth>
// See PR #207 body for the exact commands and raw outputs.
//
// Token          | Address                                    | decimals | symbol
// ---------------|--------------------------------------------|----------|-------
// USDC           | 0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48 |    6     | USDC
// USDT           | 0xdac17f958d2ee523a2206206994597c13d831ec7 |    6     | USDT
// DAI            | 0x6b175474e89094c44da98b954eedeac495271d0f |   18     | DAI
// WETH           | 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2 |   18     | WETH
// WBTC           | 0x2260fac5e5542a773aa44fbcfedf7c193bc2c599 |    8     | WBTC
// stETH (Lido)   | 0xae7ab96520de3a18e5e111b5eaab095312d7fe84 |   18     | stETH
// wstETH (Lido)  | 0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0 |   18     | wstETH
// LINK           | 0x514910771af9ca656af840dff83e8264ecf986ca |   18     | LINK
// AAVE           | 0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9 |   18     | AAVE
// ---------------------------------------------------------------------------

const ETHEREUM_ERC20: Readonly<Record<string, TokenMeta>> = Object.freeze({
  // USDC — 6 decimals (verified on-chain)
  '0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48': Object.freeze({ symbol: 'USDC', decimals: 6 }),
  // USDT — 6 decimals (verified on-chain)
  '0xdac17f958d2ee523a2206206994597c13d831ec7': Object.freeze({ symbol: 'USDT', decimals: 6 }),
  // DAI — 18 decimals (verified on-chain)
  '0x6b175474e89094c44da98b954eedeac495271d0f': Object.freeze({ symbol: 'DAI', decimals: 18 }),
  // WETH — 18 decimals (verified on-chain)
  '0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2': Object.freeze({ symbol: 'WETH', decimals: 18 }),
  // WBTC — 8 decimals (verified on-chain; NOT 18 — same category of gotcha as syrupUSDC)
  '0x2260fac5e5542a773aa44fbcfedf7c193bc2c599': Object.freeze({ symbol: 'WBTC', decimals: 8 }),
  // stETH (Lido) — 18 decimals (verified on-chain)
  '0xae7ab96520de3a18e5e111b5eaab095312d7fe84': Object.freeze({ symbol: 'stETH', decimals: 18 }),
  // wstETH (Lido wrapped stETH) — 18 decimals (verified on-chain)
  '0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0': Object.freeze({ symbol: 'wstETH', decimals: 18 }),
  // LINK (Chainlink) — 18 decimals (verified on-chain)
  '0x514910771af9ca656af840dff83e8264ecf986ca': Object.freeze({ symbol: 'LINK', decimals: 18 }),
  // AAVE — 18 decimals (verified on-chain)
  '0x7fc66500c84a76ad7e9c93437bfc5ac33e2ddae9': Object.freeze({ symbol: 'AAVE', decimals: 18 }),
})

// Ethereum mainnet (chainId 1).
const ETHEREUM_ALLOWLIST: ChainAllowlist = Object.freeze({
  native: Object.freeze({ symbol: 'ETH', decimals: 18 }),
  // 0.0001 ETH dust floor (1e14 wei). Filters 1-wei spam while still
  // allowing any meaningful micro-donation.
  nativeFloorWei: 100_000_000_000_000n,
  erc20: ETHEREUM_ERC20,
})

// Allowlist registry keyed by EIP-155 chain id.
// Chains absent from this map are not indexed by the donations processor.
const ALLOWLISTS: Readonly<Record<number, ChainAllowlist>> = Object.freeze({
  1: ETHEREUM_ALLOWLIST,
})

/** Return the allowlist for a chain, or null if the chain is not indexed. */
export const allowlistFor = (chainId: number): ChainAllowlist | null =>
  ALLOWLISTS[chainId] ?? null

/**
 * Is `tokenAddress` allowlisted on `chainId`?
 * `tokenAddress` is null for native-coin transfers.
 */
export const isAllowed = (chainId: number, tokenAddress: string | null): boolean => {
  const list = allowlistFor(chainId)
  if (!list) return false
  if (tokenAddress === null) return true
  return Object.prototype.hasOwnProperty.call(list.erc20, tokenAddress.toLowerCase())
}

/**
 * Return the native dust floor (in wei) for a chain.
 * Returns 0n for unknown chains (fail-open; caller validates chain before use).
 */
export const nativeFloor = (chainId: number): bigint =>
  allowlistFor(chainId)?.nativeFloorWei ?? 0n

/**
 * Return token metadata for an allowlisted token.
 * `tokenAddress` is null for the native coin.
 * Returns null for unknown chains or tokens (caller skips the transfer).
 */
export const tokenMeta = (chainId: number, tokenAddress: string | null): TokenMeta | null => {
  const list = allowlistFor(chainId)
  if (!list) return null
  if (tokenAddress === null) return list.native
  return list.erc20[tokenAddress.toLowerCase()] ?? null
}

/**
 * Return all allowlisted ERC20 addresses (lowercased) for a chain.
 * Returns empty array for unknown chains.
 * Used by the processor to register addLog subscriptions.
 */
export const erc20Addresses = (chainId: number): readonly string[] => {
  const list = allowlistFor(chainId)
  if (!list) return []
  return Object.keys(list.erc20)
}
