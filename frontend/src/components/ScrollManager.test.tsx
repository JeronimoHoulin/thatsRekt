/**
 * Unit tests for ScrollManager.
 *
 * Strategy: mount ScrollManager inside a MemoryRouter whose initial entries
 * and nav type we control via react-router-dom test helpers. Assert that
 * window.scrollTo is called correctly on PUSH (→ top) and on POP
 * (→ saved position or 0 when no position was saved).
 *
 * happy-dom provides window/document globals via test/setup.ts.
 * We spy on window.scrollTo and clear it between each test.
 */
import { describe, it, expect, beforeEach, afterEach, mock, spyOn } from 'bun:test'
import { render, act, cleanup, within } from '@testing-library/react'
import { MemoryRouter, Route, Routes, useNavigate } from 'react-router-dom'
import { ScrollManager } from './ScrollManager'

// ---------------------------------------------------------------------------
// ResizeObserver spy infrastructure
// ---------------------------------------------------------------------------

// Captured callbacks for each ResizeObserver instance instantiated during a
// test. Keyed sequentially so multi-observer scenarios remain deterministic.
let resizeCallbacks: ResizeObserverCallback[] = []
let origResizeObserver: typeof ResizeObserver

function installResizeObserverSpy(): void {
  resizeCallbacks = []
  // In the test environment, ResizeObserver lives on window (the happy-dom
  // GlobalWindow hoisted to globalThis.window) not on bare globalThis.
  // ScrollManager accesses it via window.ResizeObserver, so we must spy
  // on window — not globalThis — to intercept the constructor calls.
  origResizeObserver = (window as unknown as { ResizeObserver: typeof ResizeObserver }).ResizeObserver

  // Spy class records each callback so tests can fire resize events manually
  // (happy-dom's ResizeObserver doesn't auto-fire on defineProperty changes).
  const FakeResizeObserver = class implements ResizeObserver {
    #cb: ResizeObserverCallback
    constructor(cb: ResizeObserverCallback) {
      this.#cb = cb
      resizeCallbacks.push(cb)
    }
    observe(_target: Element, _options?: ResizeObserverOptions): void {}
    unobserve(_target: Element): void {}
    disconnect(): void {}
  }

  Object.defineProperty(window, 'ResizeObserver', {
    value: FakeResizeObserver,
    writable: true,
    configurable: true,
  })
}

function restoreResizeObserver(): void {
  Object.defineProperty(window, 'ResizeObserver', {
    value: origResizeObserver,
    writable: true,
    configurable: true,
  })
  resizeCallbacks = []
}

// ---------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------

let scrollToSpy: ReturnType<typeof spyOn>

beforeEach(() => {
  // Spy on window.scrollTo before each test.
  // happy-dom attaches scrollTo to the window; spyOn lets us track calls.
  scrollToSpy = spyOn(window, 'scrollTo').mockImplementation(() => undefined)
})

afterEach(() => {
  scrollToSpy.mockRestore()
  cleanup()
})

/**
 * A helper page component that exposes a navigation trigger for testing.
 */
function NavigatorPage({ to, label }: { to: string; label: string }) {
  const navigate = useNavigate()
  return (
    <button type="button" onClick={() => navigate(to)}>
      {label}
    </button>
  )
}

// ---------------------------------------------------------------------------
// Tests
// ---------------------------------------------------------------------------

describe('ScrollManager', () => {
  it('scrolls to top on PUSH navigation', async () => {
    // Mount on "/" — the initial render is a PUSH
    const { container } = render(
      <MemoryRouter initialEntries={['/']}>
        <ScrollManager />
        <Routes>
          <Route path="/" element={<NavigatorPage to="/post/base-1" label="go to post" />} />
          <Route path="/post/base-1" element={<div>post page</div>} />
        </Routes>
      </MemoryRouter>,
    )

    // Initial mount fires a scroll-to-top (PUSH on first entry)
    expect(scrollToSpy).toHaveBeenCalledWith(0, 0)
    scrollToSpy.mockClear()

    // Navigate forward — PUSH
    await act(async () => {
      within(container).getByText('go to post').click()
    })

    expect(scrollToSpy).toHaveBeenCalledWith(0, 0)
  })

  it('restores scroll position on POP navigation after a saved position', async () => {
    // Simulate: user was on "/" with scrollY = 400, navigated to a post,
    // then pressed back. ScrollManager should restore 400 on the POP.
    //
    // Approach: mount with entries ["/", "/post/base-1"] so the second
    // entry is active (index 1). Then trigger a back navigation (POP).
    // Before that, we need to prime the saved position map.
    //
    // Because the Map is internal to the component we cannot inject it
    // directly. Instead we:
    //   1. Mount with "/" as the initial route.
    //   2. Navigate to "/post/base-1" (PUSH — ScrollManager saves "/" key).
    //      We fake window.scrollY = 400 before the save fires (cleanup).
    //   3. Navigate back (POP) — ScrollManager should restore 400.
    //
    // The save happens in the effect cleanup of the previous location
    // (i.e., when the location changes away from "/"), so we set
    // window.scrollY *after* the initial render and *before* navigating.

    // happy-dom supports scrollY as a getter; override it to control the value.
    Object.defineProperty(window, 'scrollY', { value: 400, writable: true, configurable: true })

    let navigateRef: ((delta: number) => void) | null = null

    function BackNavigatorPage() {
      const navigate = useNavigate()
      navigateRef = (delta: number) => navigate(delta)
      return <div>post page</div>
    }

    const { container } = render(
      <MemoryRouter initialEntries={['/', '/post/base-1']} initialIndex={0}>
        <ScrollManager />
        <Routes>
          <Route
            path="/"
            element={
              <NavigatorPage to="/post/base-1" label="go to post" />
            }
          />
          <Route path="/post/base-1" element={<BackNavigatorPage />} />
        </Routes>
      </MemoryRouter>,
    )

    // Navigate to post (PUSH) — cleanup of "/" will save scrollY=400
    await act(async () => {
      within(container).getByText('go to post').click()
    })

    scrollToSpy.mockClear()

    // Reset scrollY to 0 (simulating the post page top)
    Object.defineProperty(window, 'scrollY', { value: 0, writable: true, configurable: true })

    // Navigate back (POP)
    await act(async () => {
      navigateRef!(-1)
    })

    // Should restore the saved 400
    expect(scrollToSpy).toHaveBeenCalledWith(0, 400)
  })

  it('scrolls to top (fallback) on POP when no position was saved', async () => {
    // Navigate forward and back without a pre-existing saved position.
    // This covers a cold deep-link → back scenario.
    let navigateRef: ((delta: number) => void) | null = null

    function BackPage() {
      const navigate = useNavigate()
      navigateRef = (delta: number) => navigate(delta)
      return <div>back page</div>
    }

    const { container } = render(
      <MemoryRouter initialEntries={['/other', '/post/base-1']} initialIndex={0}>
        <ScrollManager />
        <Routes>
          <Route path="/other" element={<NavigatorPage to="/post/base-1" label="go to post" />} />
          <Route path="/post/base-1" element={<BackPage />} />
        </Routes>
      </MemoryRouter>,
    )

    // Navigate forward — use container-scoped query to avoid cross-test DOM interference
    await act(async () => {
      within(container).getByText('go to post').click()
    })

    scrollToSpy.mockClear()

    // Navigate back — no saved position for '/other' after a cold start
    await act(async () => {
      navigateRef!(-1)
    })

    // Falls back to top
    expect(scrollToSpy).toHaveBeenCalledWith(0, 0)
  })

  it('re-applies scroll on POP when the document grows after initial restore (async content race)', async () => {
    // This is the cold-path race: the feed list renders asynchronously AFTER
    // the POP restore fires. At restore time the document is short, so
    // scrollTo(0, 700) clamps to ~219 and stays there. The fix must watch
    // for document growth via ResizeObserver and re-apply until reached.
    //
    // Test strategy:
    //   1. Install fake ResizeObserver so we can fire its callback manually.
    //   2. Stub document.documentElement.scrollHeight = 300 (target 700 unreachable).
    //   3. POP to feed with saved position 700.
    //   4. Assert initial scrollTo(0,700) was attempted.
    //   5. Grow document: scrollHeight = 11000 + stub window.scrollY = 700.
    //   6. Fire the captured ResizeObserver callback.
    //   7. Assert scrollTo(0,700) was re-applied.

    installResizeObserverSpy()

    // Make the document start short — target Y=700 is unreachable.
    Object.defineProperty(document.documentElement, 'scrollHeight', {
      value: 300,
      writable: true,
      configurable: true,
    })
    Object.defineProperty(window, 'innerHeight', {
      value: 768,
      writable: true,
      configurable: true,
    })
    Object.defineProperty(window, 'scrollY', { value: 400, writable: true, configurable: true })

    let navigateRef: ((delta: number) => void) | null = null

    function BackNavigatorPage() {
      const navigate = useNavigate()
      navigateRef = (delta: number) => navigate(delta)
      return <div>post page</div>
    }

    const { container } = render(
      <MemoryRouter initialEntries={['/', '/post/base-1']} initialIndex={0}>
        <ScrollManager />
        <Routes>
          <Route
            path="/"
            element={<NavigatorPage to="/post/base-1" label="go to post" />}
          />
          <Route path="/post/base-1" element={<BackNavigatorPage />} />
        </Routes>
      </MemoryRouter>,
    )

    // Navigate to post (PUSH) — cleanup of "/" will save scrollY=400.
    await act(async () => {
      within(container).getByText('go to post').click()
    })

    // Simulate arriving at the post page at the top.
    Object.defineProperty(window, 'scrollY', { value: 0, writable: true, configurable: true })

    scrollToSpy.mockClear()

    // Navigate back (POP) — ScrollManager restores saved Y=400 but the document
    // is short so the browser would clamp it.  The manager must also register
    // a ResizeObserver to re-apply once the document grows.
    await act(async () => {
      navigateRef!(-1)
    })

    // Initial restore was attempted.
    expect(scrollToSpy).toHaveBeenCalledWith(0, 400)
    scrollToSpy.mockClear()

    // Grow the document — feed list has rendered.
    Object.defineProperty(document.documentElement, 'scrollHeight', {
      value: 11000,
      writable: true,
      configurable: true,
    })
    // Simulate browser NOT having scrolled yet (still at 0 because the page
    // was short when the restore fired).
    Object.defineProperty(window, 'scrollY', { value: 0, writable: true, configurable: true })

    // At least one ResizeObserver must have been registered by the fix.
    // The POP effect (restoring saved=400) registers the LAST observer —
    // earlier observers may have been registered during the initial render
    // (where saved=0) and were torn down before we got here.
    expect(resizeCallbacks.length).toBeGreaterThan(0)

    // Fire the most-recently-registered ResizeObserver callback — this is the
    // one the POP effect installed while the page was still short.
    await act(async () => {
      const cb = resizeCallbacks[resizeCallbacks.length - 1]
      cb([], {} as ResizeObserver)
    })

    // Fix re-applied the target position now that the page is tall enough.
    expect(scrollToSpy).toHaveBeenCalledWith(0, 400)

    restoreResizeObserver()
  })
})
