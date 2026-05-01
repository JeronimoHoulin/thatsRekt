import { useEnsName } from 'wagmi'
import { mainnet } from 'wagmi/chains'

/**
 * Reverse-resolve an EVM address to its ENS primary name.
 *
 * Backed by wagmi's `useEnsName` (which itself uses TanStack Query under
 * the hood, so deduplication + caching is automatic). We pin chainId to
 * mainnet — ENS primary names live on Ethereum regardless of which
 * chain the address is actually active on (so the same address on Base
 * still gets its mainnet ENS name).
 *
 * Caching strategy:
 *   - `staleTime: Infinity` — once resolved, never re-query for this
 *     address during the session. ENS primary names change rarely
 *     (manual `setName` tx on mainnet) and the cost of stale data is
 *     just showing the prior name briefly. A re-render of an
 *     AddressLabel won't trigger any RPC call after the first.
 *   - `gcTime: 1 day` — keep entries alive for 24h of inactivity so a
 *     user navigating between pages doesn't re-resolve familiar
 *     addresses. After 24h idle, the entry is GC'd and re-fetched on
 *     next view.
 *
 * Cross-session persistence (ie. survives a hard reload) is NOT enabled
 * here. ENS lookups via the routeme.sh load balancer are fast (~100ms
 * cold) and adding `localStorage`-backed persistence for the whole
 * TanStack Query cache is a bigger change than this PR warrants. If the
 * cold path becomes a problem, swap in `@tanstack/query-sync-storage-persister`
 * later.
 */
export function useEnsLookup(address: `0x${string}` | undefined | null) {
  const { data: name, isLoading } = useEnsName({
    address: address ?? undefined,
    chainId: mainnet.id,
    query: {
      enabled: !!address,
      staleTime: Infinity,
      gcTime: 24 * 60 * 60 * 1_000,
      retry: 1,
    },
  })
  return { name: name ?? null, isLoading }
}
