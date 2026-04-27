import { useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import { fetchFeed, type FeedPost, type SortOption } from '../lib/queries'
import { PostCard } from '../components/PostCard'
import { EmptyState } from '../components/EmptyState'

const SORT_OPTIONS: { value: SortOption; label: string }[] = [
  { value: 'newest', label: 'newest' },
  { value: 'oldest', label: 'oldest' },
]

export function Feed() {
  const [sort, setSort] = useState<SortOption>('newest')

  const { data, isLoading, error } = useQuery({
    queryKey: ['feed', sort],
    queryFn: () => fetchFeed(50, sort),
  })

  return (
    <div>
      <SortBar current={sort} onChange={setSort} />
      <div className="mt-6">
        {isLoading ? (
          <p className="text-xs uppercase tracking-widest text-neutral-700">loading…</p>
        ) : error ? (
          <EmptyState
            title="couldn't load the feed."
            hint={`is the indexer running? ${(error as Error).message}`}
          />
        ) : !data || data.length === 0 ? (
          <EmptyState
            title="no posts yet."
            hint="contract not deployed on the indexed chain, or no whitelister has posted an alert yet."
          />
        ) : (
          <FeedList posts={data} />
        )}
      </div>
    </div>
  )
}

function SortBar({
  current,
  onChange,
}: {
  current: SortOption
  onChange: (s: SortOption) => void
}) {
  return (
    <div className="flex items-baseline gap-3 border-b border-black pb-3">
      <span className="text-[10px] uppercase tracking-widest text-neutral-700">sort:</span>
      <div className="flex gap-1">
        {SORT_OPTIONS.map((opt) => {
          const active = current === opt.value
          return (
            <button
              key={opt.value}
              type="button"
              onClick={() => onChange(opt.value)}
              className={
                'px-2 py-0.5 text-xs uppercase tracking-widest border ' +
                (active
                  ? 'border-black bg-black text-[#f5f4ee]'
                  : 'border-transparent text-neutral-700 hover:border-black hover:text-black')
              }
            >
              {opt.label}
            </button>
          )
        })}
      </div>
    </div>
  )
}

function FeedList({ posts }: { posts: FeedPost[] }) {
  return (
    <div>
      {posts.map((post, i) => (
        <div key={post.id}>
          {i > 0 && <hr className="my-8 border-t-2 border-black" />}
          <PostCard post={post} />
        </div>
      ))}
      <div className="rekt-divider mt-8">* * *</div>
      <p className="text-center text-xs uppercase tracking-widest text-neutral-700">
        end of feed · {posts.length} post{posts.length === 1 ? '' : 's'}
      </p>
    </div>
  )
}

