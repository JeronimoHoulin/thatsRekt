import type { EditEntity, VoteEntity } from '../lib/queries'
import { AddressLabel } from './AddressLabel'
import { formatTimestamp, relativeTime } from '../lib/format'

type TimelineItem =
  | { kind: 'vote'; data: VoteEntity }
  | { kind: 'edit'; data: EditEntity }

export function Timeline({
  votes,
  edits,
  chainSlug,
}: {
  votes: VoteEntity[]
  edits: EditEntity[]
  chainSlug?: string
}) {
  const items: TimelineItem[] = [
    ...votes.map((v): TimelineItem => ({ kind: 'vote', data: v })),
    ...edits.map((e): TimelineItem => ({ kind: 'edit', data: e })),
  ].sort((a, b) => a.data.blockNumber - b.data.blockNumber)

  if (items.length === 0) {
    return (
      <p className="text-xs uppercase tracking-widest text-neutral-700">
        no activity on this post yet.
      </p>
    )
  }

  return (
    <ol className="space-y-4">
      {items.map((item, i) => (
        <li
          key={`${item.kind}-${item.data.id}`}
          className="border-l-2 border-black pl-4"
        >
          <div className="flex items-baseline gap-2 text-[10px] uppercase tracking-widest text-neutral-700">
            <span className="font-black text-black">{String(i + 1).padStart(2, '0')}</span>
            <span title={formatTimestamp(item.data.timestamp)}>
              block {item.data.blockNumber} · {relativeTime(item.data.timestamp)}
            </span>
          </div>
          {item.kind === 'vote' ? (
            <VoteRow vote={item.data} chainSlug={chainSlug} />
          ) : (
            <EditRow edit={item.data} />
          )}
        </li>
      ))}
    </ol>
  )
}

function VoteRow({ vote, chainSlug }: { vote: VoteEntity; chainSlug?: string }) {
  const action = describeVote(vote.oldDirection, vote.newDirection)
  return (
    <p className="mt-1 text-sm">
      <AddressLabel addr={vote.voter.id} chainSlug={chainSlug} />{' '}
      <span className={`font-black uppercase tracking-tight ${voteColor(vote.newDirection)}`}>
        {action.icon} {action.label}
      </span>
    </p>
  )
}

function EditRow({ edit }: { edit: EditEntity }) {
  return (
    <div className="mt-1 space-y-1">
      <p className="text-sm font-black uppercase tracking-tight">
        {describeEditKind(edit.kind)}
      </p>
      {edit.kind === 'AmendNote' && edit.newNote != null && (
        <p className="text-sm leading-relaxed text-neutral-800">{edit.newNote}</p>
      )}
      {edit.kind === 'AddAttackers' && edit.addedAttackers && (
        <ul className="space-y-0.5 text-xs font-mono">
          {edit.addedAttackers.map((a) => (
            <li key={a}>+ {a}</li>
          ))}
        </ul>
      )}
      {edit.kind === 'AddVictims' && edit.addedVictims && (
        <ul className="space-y-0.5 text-xs font-mono">
          {edit.addedVictims.map((v) => (
            <li key={v}>+ {v}</li>
          ))}
        </ul>
      )}
    </div>
  )
}

function describeVote(
  oldDir: VoteEntity['oldDirection'],
  newDir: VoteEntity['newDirection'],
): { icon: string; label: string } {
  if (oldDir === 'None' && newDir === 'Upvote') return { icon: '↑', label: 'upvoted' }
  if (oldDir === 'None' && newDir === 'Downvote') return { icon: '↓', label: 'downvoted' }
  if (newDir === 'None') return { icon: '×', label: 'cleared their vote' }
  if (oldDir === 'Upvote' && newDir === 'Downvote')
    return { icon: '↓', label: 'switched to downvote' }
  if (oldDir === 'Downvote' && newDir === 'Upvote')
    return { icon: '↑', label: 'switched to upvote' }
  return { icon: '·', label: `${oldDir} → ${newDir}` }
}

function voteColor(newDir: string): string {
  if (newDir === 'None') return 'text-neutral-700'
  if (newDir === 'Upvote') return 'text-emerald-700'
  if (newDir === 'Downvote') return 'text-red-600'
  return 'text-neutral-700'
}

function describeEditKind(kind: EditEntity['kind']): string {
  switch (kind) {
    case 'AmendNote':
      return 'note amended'
    case 'AddAttackers':
      return 'attackers added'
    case 'AddVictims':
      return 'victims added'
  }
}
