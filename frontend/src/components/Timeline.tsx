import type { EditEntity, VoteEntity } from '../lib/queries'
import { AddressLabel } from './AddressLabel'
import { formatTimestamp, relativeTime } from '../lib/format'

type TimelineItem =
  | { kind: 'vote'; data: VoteEntity }
  | { kind: 'edit'; data: EditEntity }

export function Timeline({
  votes,
  edits,
}: {
  votes: VoteEntity[]
  edits: EditEntity[]
}) {
  const items: TimelineItem[] = [
    ...votes.map((v): TimelineItem => ({ kind: 'vote', data: v })),
    ...edits.map((e): TimelineItem => ({ kind: 'edit', data: e })),
  ].sort((a, b) => a.data.blockNumber - b.data.blockNumber)

  if (items.length === 0) {
    return (
      <p className="text-sm text-neutral-500">No activity on this post yet.</p>
    )
  }

  return (
    <ul className="space-y-2">
      {items.map((item) => (
        <li
          key={`${item.kind}-${item.data.id}`}
          className="rounded-md border border-neutral-800 bg-neutral-950 p-3"
        >
          {item.kind === 'vote' ? <VoteRow vote={item.data} /> : <EditRow edit={item.data} />}
        </li>
      ))}
    </ul>
  )
}

function VoteRow({ vote }: { vote: VoteEntity }) {
  const action = describeVote(vote.oldDirection, vote.newDirection)
  return (
    <div className="flex items-baseline justify-between gap-3 text-sm">
      <div className="flex items-center gap-2">
        <span className={`font-mono text-xs ${voteColor(vote.newDirection, vote.oldDirection)}`}>
          {action.icon}
        </span>
        <AddressLabel addr={vote.voter.id} />
        <span className="text-neutral-400">{action.label}</span>
      </div>
      <span
        className="font-mono text-xs text-neutral-500"
        title={formatTimestamp(vote.timestamp)}
      >
        block {vote.blockNumber} · {relativeTime(vote.timestamp)}
      </span>
    </div>
  )
}

function EditRow({ edit }: { edit: EditEntity }) {
  return (
    <div className="space-y-1">
      <div className="flex items-baseline justify-between gap-3 text-sm">
        <span className="text-neutral-300">{describeEditKind(edit.kind)}</span>
        <span
          className="font-mono text-xs text-neutral-500"
          title={formatTimestamp(edit.timestamp)}
        >
          block {edit.blockNumber} · {relativeTime(edit.timestamp)}
        </span>
      </div>
      {edit.kind === 'AmendNote' && edit.newNote != null && (
        <p className="text-sm text-neutral-200">{edit.newNote}</p>
      )}
      {edit.kind === 'AddAttackers' && edit.addedAttackers && (
        <div className="flex flex-wrap gap-1 text-xs">
          {edit.addedAttackers.map((a) => (
            <code key={a} className="rounded bg-neutral-900 px-2 py-0.5 font-mono">
              {a}
            </code>
          ))}
        </div>
      )}
      {edit.kind === 'AddVictims' && edit.addedVictims && (
        <div className="flex flex-wrap gap-1 text-xs">
          {edit.addedVictims.map((v) => (
            <code key={v} className="rounded bg-neutral-900 px-2 py-0.5 font-mono">
              {v}
            </code>
          ))}
        </div>
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

function voteColor(newDir: string, oldDir: string): string {
  if (newDir === 'None') return 'text-neutral-500'
  if (newDir === 'Upvote') return 'text-emerald-400'
  if (newDir === 'Downvote') return 'text-rose-400'
  return 'text-neutral-400'
}

function describeEditKind(kind: EditEntity['kind']): string {
  switch (kind) {
    case 'AmendNote':
      return 'Note amended'
    case 'AddAttackers':
      return 'Attackers added'
    case 'AddVictims':
      return 'Victims added'
  }
}
