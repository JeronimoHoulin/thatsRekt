import { useQuery } from '@tanstack/react-query'
import { fetchFeed } from '../lib/queries'
import { PostCard } from '../components/PostCard'
import { EmptyState } from '../components/EmptyState'

export function Feed() {
  const { data, isLoading, error } = useQuery({
    queryKey: ['feed'],
    queryFn: () => fetchFeed(50),
  })

  if (isLoading) {
    return <p className="text-sm text-neutral-500">Loading feed…</p>
  }

  if (error) {
    return (
      <EmptyState
        title="Couldn't load the feed."
        hint={`Is the indexer running? Error: ${(error as Error).message}`}
      />
    )
  }

  if (!data || data.length === 0) {
    return (
      <EmptyState
        title="No posts yet."
        hint="Either the contract isn't deployed on the indexed chain, or no whitelister has posted an alert yet."
      />
    )
  }

  return (
    <ul className="space-y-3">
      {data.map((post) => (
        <li key={post.id}>
          <PostCard post={post} />
        </li>
      ))}
    </ul>
  )
}
