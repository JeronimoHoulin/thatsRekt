import { Link } from 'react-router-dom'
import type { FeedPost } from '../lib/queries'
import { relativeTime, scoreColor } from '../lib/format'
import { AddressLabel } from './AddressLabel'

export function PostCard({ post }: { post: FeedPost }) {
  return (
    <Link
      to={`/post/${post.id}`}
      className="block rounded-md border border-neutral-800 bg-neutral-950 p-4 transition hover:border-neutral-700"
    >
      <div className="flex items-start justify-between gap-4">
        <div className="flex flex-col gap-1">
          <div className="flex items-center gap-2 text-xs text-neutral-500">
            <span>#{post.id}</span>
            <span>·</span>
            <span>by</span>
            <AddressLabel addr={post.poster.id} />
            <span>·</span>
            <span>attacked {relativeTime(post.attackedAt)}</span>
          </div>
          <p className="line-clamp-3 text-sm text-neutral-200">{post.note || '(no note)'}</p>
          <div className="mt-1 flex flex-wrap gap-x-4 gap-y-1 text-xs text-neutral-500">
            <span>
              <span className="text-neutral-300">{post.attackerLinks.length}</span>{' '}
              attacker{post.attackerLinks.length === 1 ? '' : 's'}
            </span>
            <span>
              <span className="text-neutral-300">{post.victimLinks.length}</span>{' '}
              victim{post.victimLinks.length === 1 ? '' : 's'}
            </span>
            <span>{relativeTime(post.createdAtTimestamp)} on-chain</span>
          </div>
        </div>
        <div className="flex flex-col items-end gap-1 shrink-0">
          <div className={`text-2xl font-semibold tabular-nums ${scoreColor(post.netScore)}`}>
            {post.netScore > 0 ? `+${post.netScore}` : post.netScore}
          </div>
          <div className="font-mono text-xs text-neutral-500">
            <span className="text-emerald-400">{post.upvotes}↑</span>{' '}
            <span className="text-rose-400">{post.downvotes}↓</span>
          </div>
        </div>
      </div>
    </Link>
  )
}
