import { Link } from 'react-router-dom'
import type { FeedPost } from '../lib/queries'
import { relativeTime, shortAddress } from '../lib/format'

export function PostCard({ post }: { post: FeedPost }) {
  const tags = [
    `poster: ${shortAddress(post.poster.id)}`,
    `${post.attackerLinks.length} attacker${post.attackerLinks.length === 1 ? '' : 's'}`,
    `${post.victimLinks.length} victim${post.victimLinks.length === 1 ? '' : 's'}`,
  ]

  return (
    <article className="space-y-3">
      <Link
        to={`/post/${post.id}`}
        className="block group"
      >
        <h2 className="font-black uppercase tracking-tight text-3xl leading-tight group-hover:text-red-600">
          #{post.id} —{' '}
          <span className="text-neutral-800 group-hover:text-red-600">
            {firstLine(post.note) || 'untitled alert'}
          </span>
        </h2>
      </Link>

      <div className="flex flex-wrap items-baseline gap-x-2 gap-y-1 text-xs uppercase tracking-widest text-neutral-700">
        <span>attacked {relativeTime(post.attackedAt)}</span>
        {tags.map((tag, i) => (
          <span key={i}>
            <span className="text-neutral-400">·</span>{' '}
            <span>[{tag}]</span>
          </span>
        ))}
        <span>
          <span className="text-neutral-400">·</span>{' '}
          <ScoreBadge net={post.netScore} up={post.upvotes} down={post.downvotes} />
        </span>
      </div>

      <p className="text-base leading-relaxed text-neutral-800 line-clamp-3">
        {post.note || '(no note)'}
      </p>

      <Link
        to={`/post/${post.id}`}
        className="inline-block text-xs font-black uppercase tracking-widest rekt-link"
      >
        more →
      </Link>
    </article>
  )
}

function ScoreBadge({ net, up, down }: { net: number; up: number; down: number }) {
  const color = net > 0 ? 'text-emerald-700' : net < 0 ? 'text-red-600' : 'text-neutral-700'
  return (
    <span className={`font-mono ${color}`}>
      {net >= 0 ? `+${net}` : net} ({up}↑/{down}↓)
    </span>
  )
}

function firstLine(s: string): string {
  if (!s) return ''
  const idx = s.indexOf('.')
  const slice = idx > 0 ? s.slice(0, idx) : s
  return slice.length > 100 ? slice.slice(0, 97) + '…' : slice
}
