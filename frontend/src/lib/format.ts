/** Truncate a hex address to e.g. `0x1234…abcd`. */
export function shortAddress(addr: string, head = 6, tail = 4): string {
  if (!addr) return ''
  if (addr.length <= head + tail + 2) return addr
  return `${addr.slice(0, head)}…${addr.slice(-tail)}`
}

/** Human-readable relative time (e.g. "5 minutes ago"). */
export function relativeTime(iso: string): string {
  const t = new Date(iso).getTime()
  const now = Date.now()
  const diff = Math.max(0, now - t)
  const sec = Math.floor(diff / 1000)
  if (sec < 60) return `${sec}s ago`
  const min = Math.floor(sec / 60)
  if (min < 60) return `${min}m ago`
  const hr = Math.floor(min / 60)
  if (hr < 24) return `${hr}h ago`
  const day = Math.floor(hr / 24)
  if (day < 30) return `${day}d ago`
  const mo = Math.floor(day / 30)
  if (mo < 12) return `${mo}mo ago`
  const yr = Math.floor(mo / 12)
  return `${yr}y ago`
}

/** Format a timestamp as a sortable readable date. */
export function formatTimestamp(iso: string): string {
  const d = new Date(iso)
  return d.toISOString().slice(0, 19).replace('T', ' ') + ' UTC'
}

/** Format a timestamp as dd/mm/yyyy. */
export function formatDateOnly(iso: string): string {
  const d = new Date(iso)
  const day = String(d.getUTCDate()).padStart(2, '0')
  const month = String(d.getUTCMonth() + 1).padStart(2, '0')
  const year = d.getUTCFullYear()
  return `${day}/${month}/${year}`
}

/** Score color: positive = green, negative = red, zero = gray. */
export function scoreColor(score: number): string {
  if (score > 0) return 'text-emerald-400'
  if (score < 0) return 'text-rose-400'
  return 'text-neutral-400'
}

/** Resolve an X / Twitter handle (or a full URL) to a profile URL. */
export function twitterUrl(handleOrUrl: string): string {
  if (handleOrUrl.startsWith('http://') || handleOrUrl.startsWith('https://')) {
    return handleOrUrl
  }
  const cleaned = handleOrUrl.replace(/^@/, '')
  return `https://x.com/${cleaned}`
}

/** Extract Twitter/X source URL from note if available.
 * Note format from relay: "tweet_content\ntweet_url\n[image_urls...]"
 * Returns null if no URL found or note is empty. */
export function extractSourceUrl(note: string | undefined): string | null {
  if (!note || typeof note !== 'string') return null
  const lines = note.split('\n')
  if (lines.length < 2) return null
  const potentialUrl = lines[1].trim()
  if (potentialUrl.startsWith('http://') || potentialUrl.startsWith('https://')) {
    return potentialUrl
  }
  return null
}
