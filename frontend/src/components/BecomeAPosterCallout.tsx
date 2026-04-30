/**
 * Email address for prospective vetted posters to apply. Defined in one
 * place so swapping it out is a single edit.
 */
export const BECOME_POSTER_EMAIL = 'thatsrekt@protonmail.com'

interface BecomeAPosterCalloutProps {
  /**
   * Visual treatment:
   *   - `card`  — full bordered card, used on /docs
   *   - `inline`— compact prose paragraph, used on /about under the hero
   */
  variant?: 'card' | 'inline'
}

export function BecomeAPosterCallout({
  variant = 'card',
}: BecomeAPosterCalloutProps) {
  const subject = encodeURIComponent('thatsRekt — vetted poster application')
  const body = encodeURIComponent(
    [
      'Team / detector name:',
      'Public profile (X / GitHub / website):',
      'Detection focus (which protocols, which chains, which exploit classes):',
      'Existing track record (writeups, prior incidents flagged, etc.):',
      'Address you want whitelisted:',
      '',
      "We'll review and reply with next steps.",
    ].join('\n'),
  )
  const mailto = `mailto:${BECOME_POSTER_EMAIL}?subject=${subject}&body=${body}`

  if (variant === 'inline') {
    return (
      <p className="text-sm leading-relaxed text-neutral-800">
        Run a security team or automated detector?{' '}
        <a href={mailto} className="rekt-link font-black uppercase tracking-widest">
          apply to post →
        </a>
      </p>
    )
  }

  return (
    <section className="border-2 border-black bg-yellow-50 p-5 space-y-3">
      <h3 className="font-black uppercase tracking-widest text-xs">
        become a poster
      </h3>
      <p className="text-sm leading-relaxed text-neutral-800">
        Posting is permissioned — addresses are added to the whitelist
        by governance after a vetting review. If you run a security
        team, an exploit detector, or any pipeline that produces
        timely on-chain incident signals, get in touch.
      </p>
      <p className="text-sm leading-relaxed text-neutral-800">
        Email{' '}
        <a
          href={mailto}
          className="rekt-link font-mono text-sm break-all"
        >
          {BECOME_POSTER_EMAIL}
        </a>{' '}
        with a short pitch — track record, detection focus, and the
        address you want whitelisted. We'll review and propose your
        addition through the on-chain governance flow (subject to the
        7-day timelock).
      </p>
    </section>
  )
}
