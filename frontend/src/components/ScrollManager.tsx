/**
 * ScrollManager — mounts once inside <BrowserRouter>, renders null.
 *
 * Behaviour:
 *   - POP (browser/OS back or forward): restore the saved scroll position
 *     for the destination history entry (keyed by location.key).
 *
 *     The restore is applied immediately, re-applied on the next animation
 *     frame (handles react-query synchronous cache-hit staleTime renders),
 *     and additionally watched via a ResizeObserver on document.documentElement.
 *     The observer re-applies scrollTo each time the document grows until the
 *     target position is actually reached (within 2px) — this handles the
 *     cold-path race where the feed list renders asynchronously AFTER the
 *     initial restore, causing scrollTo to clamp to the short document height.
 *
 *     A hard 1000ms deadline disconnects the observer to prevent leaks.
 *     The observer and deadline are also torn down in effect cleanup (i.e.,
 *     on the next navigation) so they never outlive the current history entry.
 *
 *   - PUSH / REPLACE (forward navigation): scroll to the top.
 *
 * Positions are saved keyed by location.key in a useRef<Map> so they
 * persist for the whole session without triggering re-renders.
 *
 * history.scrollRestoration is set to 'manual' once so the browser's own
 * heuristic doesn't fight us.
 */
import { useEffect, useRef } from 'react'
import { useLocation, useNavigationType } from 'react-router-dom'

/** Tolerance in pixels: within this distance we consider the target reached. */
const REACH_TOLERANCE_PX = 2

/** Maximum time (ms) we keep the ResizeObserver alive waiting for the page to grow. */
const OBSERVER_DEADLINE_MS = 1000

export function ScrollManager(): null {
  const location = useLocation()
  const navType = useNavigationType()

  // Keyed by location.key (unique per history entry, stable on revisit).
  const positions = useRef<Map<string, number>>(new Map())

  // Disable the browser's own scroll-restoration once, globally.
  useEffect(() => {
    if (typeof window !== 'undefined' && 'scrollRestoration' in window.history) {
      window.history.scrollRestoration = 'manual'
    }
  }, [])

  useEffect(() => {
    const key = location.key

    if (navType === 'POP') {
      const saved = positions.current.get(key) ?? 0

      // Apply immediately…
      window.scrollTo(0, saved)

      // …and re-apply one frame later so any layout that occurs during
      // react-query's synchronous cache-hit render doesn't clamp us short.
      const rafId = requestAnimationFrame(() => {
        window.scrollTo(0, saved)
      })

      // ResizeObserver guard: handles the async content race where the page
      // content (e.g. react-query feed list) renders multiple frames after
      // the POP, making the document taller than it was when the initial
      // scrollTo fired.  We watch document.documentElement for size changes
      // and re-apply scrollTo on each resize until we reach the target or
      // the deadline expires.
      let observer: ResizeObserver | null = null
      let deadlineId: ReturnType<typeof setTimeout> | null = null

      function teardown(): void {
        if (observer !== null) {
          observer.disconnect()
          observer = null
        }
        if (deadlineId !== null) {
          clearTimeout(deadlineId)
          deadlineId = null
        }
      }

      function onResize(): void {
        // Target not reachable yet? Keep waiting.
        const maxScroll =
          document.documentElement.scrollHeight - window.innerHeight
        if (maxScroll < saved) {
          return
        }

        // Page is now tall enough — re-apply and check if we actually reached it.
        window.scrollTo(0, saved)

        // Disconnect once we've landed within tolerance so we don't keep
        // calling scrollTo after the user might have scrolled themselves.
        if (Math.abs(window.scrollY - saved) <= REACH_TOLERANCE_PX) {
          teardown()
        }
      }

      if (typeof window.ResizeObserver !== 'undefined') {
        observer = new window.ResizeObserver(onResize)
        observer.observe(document.documentElement)

        // Hard deadline: disconnect even if we never reached the target.
        deadlineId = setTimeout(teardown, OBSERVER_DEADLINE_MS)
      }

      return () => {
        cancelAnimationFrame(rafId)
        teardown()
        // Save current scroll before leaving this entry.
        positions.current.set(key, window.scrollY)
      }
    } else {
      // PUSH or REPLACE — always start at the top.
      window.scrollTo(0, 0)
      return () => {
        // Save current scroll before leaving this entry.
        positions.current.set(key, window.scrollY)
      }
    }
  }, [location.key, navType])

  return null
}
