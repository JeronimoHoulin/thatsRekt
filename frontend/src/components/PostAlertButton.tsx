import { useEffect, useRef, useState } from 'react'
import { Link } from 'react-router-dom'

const APPLY_EMAIL = 'thatsrekt@protonmail.com'

/**
 * Header-mounted CTA: "post alert" button + a modal that explains the
 * whitelist gate. Two purposes:
 *
 *   1. **Discoverability.** Most visitors don't know the registry takes
 *      writes at all. A primary-color filled button in the header
 *      surfaces the action.
 *   2. **Onboarding.** Anyone who clicks gets routed: existing
 *      whitelisters get pointed at the posters list (where the actual
 *      composer will live in a follow-up); non-whitelisters see how to
 *      apply.
 *
 * Out of scope for v1: the actual on-chain post composer (form +
 * wallet connect + tx). That arrives once we wire wagmi/viem and pick
 * a connector. Today the modal is just informational.
 *
 * Style: brutalist, sharp corners, monospace uppercase. The trigger is
 * a red-fill rectangle (matching the `REKT` accent in the logo + the
 * demo banner) — red is reserved for primary CTAs site-wide so it
 * reads as "this is the action".
 */
export function PostAlertButton({
  variant = 'desktop',
  onAfterClick,
}: {
  /** `desktop` for the header strip; `mobile` for inside the mobile menu */
  variant?: 'desktop' | 'mobile'
  /** invoked after the modal opens — used by the mobile menu to close itself */
  onAfterClick?: () => void
}) {
  const [open, setOpen] = useState(false)

  const handleClick = () => {
    setOpen(true)
    onAfterClick?.()
  }

  return (
    <>
      <button
        type="button"
        onClick={handleClick}
        className={
          variant === 'desktop'
            ? 'inline-flex items-center gap-1 border-2 border-red-600 bg-red-600 text-white px-3 py-1 text-[11px] uppercase tracking-widest font-black hover:bg-red-700 hover:border-red-700 transition-colors focus:outline-none focus:ring-2 focus:ring-red-600 focus:ring-offset-1'
            : 'block w-full text-left px-4 py-3 text-sm uppercase tracking-widest font-black bg-red-600 text-white hover:bg-red-700 active:bg-red-800 transition-colors'
        }
      >
        <span aria-hidden="true">+</span>
        <span>post alert</span>
      </button>
      {open && <PostAlertModal onClose={() => setOpen(false)} />}
    </>
  )
}

function PostAlertModal({ onClose }: { onClose: () => void }) {
  const dialogRef = useRef<HTMLDivElement>(null)

  // Close on Escape; trap nothing else (this modal has no form fields
  // worth focus-trapping yet — keep it simple).
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if (e.key === 'Escape') onClose()
    }
    document.addEventListener('keydown', onKey)
    // Lock background scroll while open
    const prevOverflow = document.body.style.overflow
    document.body.style.overflow = 'hidden'
    return () => {
      document.removeEventListener('keydown', onKey)
      document.body.style.overflow = prevOverflow
    }
  }, [onClose])

  // Auto-focus the close button so Escape isn't the only exit for
  // keyboard users who tab in
  useEffect(() => {
    dialogRef.current?.querySelector<HTMLButtonElement>('[data-close]')?.focus()
  }, [])

  return (
    <div
      role="dialog"
      aria-modal="true"
      aria-labelledby="post-alert-modal-title"
      className="fixed inset-0 z-50 flex items-start justify-center bg-black/60 px-4 py-12 sm:py-20 overflow-y-auto"
      onClick={onClose}
    >
      <div
        ref={dialogRef}
        onClick={(e) => e.stopPropagation()}
        className="relative w-full max-w-md border-2 border-black bg-[#f5f4ee] shadow-[6px_6px_0_0_#000]"
      >
        {/* Title strip */}
        <header className="flex items-center justify-between border-b-2 border-black px-4 py-2 bg-black text-[#f5f4ee]">
          <h2
            id="post-alert-modal-title"
            className="text-[11px] uppercase tracking-widest font-black"
          >
            [post alert]
          </h2>
          <button
            type="button"
            data-close
            onClick={onClose}
            aria-label="close"
            className="text-[#f5f4ee] hover:text-red-500 -mr-1 px-1 leading-none text-lg"
          >
            ✕
          </button>
        </header>

        <div className="px-5 py-5 space-y-5">
          <p className="text-sm leading-relaxed text-neutral-800">
            Posting requires a{' '}
            <strong className="font-black">whitelisted address</strong>.
            Whitelisters are vetted security teams and automated
            detectors who can submit alerts on-chain.
          </p>

          {/* Path 1: existing whitelister */}
          <section className="border-2 border-black bg-white p-4 space-y-2">
            <p className="text-[10px] uppercase tracking-widest text-neutral-700">
              [already whitelisted?]
            </p>
            <p className="text-sm leading-relaxed text-neutral-800">
              You can post directly from any of the addresses listed
              under <Code>/posters</Code>. The on-chain composer is
              shipping soon — for now, post by calling{' '}
              <Code>post(...)</Code> on the registry contract from your
              whitelisted EOA.
            </p>
            <Link
              to="/posters"
              onClick={onClose}
              className="inline-block mt-2 text-xs uppercase tracking-widest font-black text-black hover:text-red-600 underline underline-offset-4 decoration-2"
            >
              see the poster list →
            </Link>
          </section>

          {/* Path 2: not whitelisted yet */}
          <section className="border-2 border-black bg-white p-4 space-y-2">
            <p className="text-[10px] uppercase tracking-widest text-neutral-700">
              [want to become a poster?]
            </p>
            <p className="text-sm leading-relaxed text-neutral-800">
              Tell us who you are and what you'd be reporting. Adding a
              new poster goes through a 3-day public timelock so the
              rotation is visible before it lands.
            </p>
            <a
              href={`mailto:${APPLY_EMAIL}?subject=poster%20application`}
              className="inline-block mt-2 text-xs uppercase tracking-widest font-black text-black hover:text-red-600 underline underline-offset-4 decoration-2"
            >
              email {APPLY_EMAIL} →
            </a>
          </section>
        </div>
      </div>
    </div>
  )
}

function Code({ children }: { children: React.ReactNode }) {
  return (
    <code className="font-mono text-xs bg-neutral-100 border border-neutral-300 px-1 py-0.5">
      {children}
    </code>
  )
}
