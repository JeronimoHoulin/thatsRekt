import { useParams, Link } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { fetchPostDetail } from '../lib/queries'
import { AddressLabel } from '../components/AddressLabel'
import { Timeline } from '../components/Timeline'
import { EmptyState } from '../components/EmptyState'
import { formatTimestamp, relativeTime, scoreColor } from '../lib/format'

export function PostDetail() {
  const { id } = useParams<{ id: string }>()
  const postId = id ?? ''

  const { data, isLoading, error } = useQuery({
    queryKey: ['post', postId],
    queryFn: () => fetchPostDetail(postId),
    enabled: postId.length > 0,
  })

  if (isLoading) {
    return <p className="text-sm text-neutral-500">Loading post #{postId}…</p>
  }

  if (error) {
    return (
      <EmptyState
        title="Couldn't load this post."
        hint={`Error: ${(error as Error).message}`}
      />
    )
  }

  if (!data) {
    return (
      <EmptyState
        title={`Post #${postId} not found.`}
        hint="The id may be wrong, or the post hasn't been indexed yet."
      />
    )
  }

  return (
    <article className="space-y-8">
      <Link to="/" className="text-xs text-neutral-500 hover:text-neutral-300">
        ← back to feed
      </Link>

      <header className="flex flex-col gap-2">
        <div className="flex items-center gap-3">
          <h1 className="text-xl font-semibold">Post #{data.id}</h1>
          {data.removed ? (
            <span className="rounded bg-rose-900/40 px-2 py-0.5 text-xs text-rose-300">
              retracted
            </span>
          ) : (
            <span className="rounded bg-emerald-900/30 px-2 py-0.5 text-xs text-emerald-300">
              active
            </span>
          )}
          <div className={`ml-auto text-2xl font-semibold tabular-nums ${scoreColor(data.netScore)}`}>
            {data.netScore > 0 ? `+${data.netScore}` : data.netScore}
            <span className="ml-2 font-mono text-xs text-neutral-500">
              <span className="text-emerald-400">{data.upvotes}↑</span>{' '}
              <span className="text-rose-400">{data.downvotes}↓</span>
            </span>
          </div>
        </div>
        <dl className="grid grid-cols-1 gap-1 text-xs text-neutral-500 sm:grid-cols-2">
          <Field label="Posted by">
            <AddressLabel addr={data.poster.id} />
          </Field>
          <Field label="Attacked at" tooltip={formatTimestamp(data.attackedAt)}>
            {relativeTime(data.attackedAt)}
          </Field>
          <Field label="Posted on-chain" tooltip={formatTimestamp(data.createdAtTimestamp)}>
            {relativeTime(data.createdAtTimestamp)}
          </Field>
          <Field label="Last updated" tooltip={formatTimestamp(data.lastUpdatedAt)}>
            {relativeTime(data.lastUpdatedAt)}
          </Field>
          {data.removed && data.removedAtTimestamp && (
            <Field label="Retracted" tooltip={formatTimestamp(data.removedAtTimestamp)}>
              {relativeTime(data.removedAtTimestamp)}
            </Field>
          )}
        </dl>
      </header>

      <section className="space-y-2">
        <h2 className="text-sm font-semibold uppercase tracking-wide text-neutral-400">
          Note
        </h2>
        <p className="rounded-md border border-neutral-800 bg-neutral-950 p-4 text-sm text-neutral-200 whitespace-pre-wrap">
          {data.note || '(no note)'}
        </p>
      </section>

      <section className="space-y-2">
        <h2 className="text-sm font-semibold uppercase tracking-wide text-neutral-400">
          Attackers ({data.attackerLinks.length})
        </h2>
        {data.attackerLinks.length === 0 ? (
          <p className="text-sm text-neutral-500">None listed.</p>
        ) : (
          <ul className="space-y-1">
            {data.attackerLinks.map((link) => (
              <li
                key={link.address.id}
                className="flex items-baseline justify-between rounded border border-neutral-800 bg-neutral-950 px-3 py-2 text-sm"
              >
                <AddressLabel addr={link.address.id} full />
                <div className="flex gap-3 font-mono text-xs">
                  <span className={scoreColor(Number(link.address.attackerScore))}>
                    score {link.address.attackerScore}
                  </span>
                  {link.address.attackerAppearances != null && (
                    <span className="text-neutral-500">
                      {link.address.attackerAppearances} post(s)
                    </span>
                  )}
                </div>
              </li>
            ))}
          </ul>
        )}
      </section>

      <section className="space-y-2">
        <h2 className="text-sm font-semibold uppercase tracking-wide text-neutral-400">
          Victims ({data.victimLinks.length})
        </h2>
        {data.victimLinks.length === 0 ? (
          <p className="text-sm text-neutral-500">None listed.</p>
        ) : (
          <ul className="space-y-1">
            {data.victimLinks.map((link) => (
              <li
                key={link.address.id}
                className="flex items-baseline justify-between rounded border border-neutral-800 bg-neutral-950 px-3 py-2 text-sm"
              >
                <AddressLabel addr={link.address.id} full />
                <span className="font-mono text-xs text-neutral-500">
                  {link.address.isVictim ? 'flagged' : 'cleared'}
                </span>
              </li>
            ))}
          </ul>
        )}
      </section>

      <section className="space-y-2">
        <h2 className="text-sm font-semibold uppercase tracking-wide text-neutral-400">
          Timeline
        </h2>
        <Timeline votes={data.votes} edits={data.edits} />
      </section>
    </article>
  )
}

function Field({
  label,
  children,
  tooltip,
}: {
  label: string
  children: React.ReactNode
  tooltip?: string
}) {
  return (
    <div className="flex gap-2" title={tooltip}>
      <dt className="w-32 shrink-0 text-neutral-500">{label}</dt>
      <dd className="text-neutral-300">{children}</dd>
    </div>
  )
}
