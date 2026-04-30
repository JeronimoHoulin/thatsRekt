import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { WagmiProvider } from 'wagmi'
import { App } from './App'
import { wagmiConfig } from './lib/wagmi'
import './index.css'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 30_000,
      refetchOnWindowFocus: false,
      retry: 1,
    },
  },
})

const rootElement = document.getElementById('root')
if (!rootElement) throw new Error('No root element found')

// Provider order matters: Wagmi must be inside QueryClientProvider (it
// uses TanStack Query under the hood for read-contract caching).
createRoot(rootElement).render(
  <StrictMode>
    <QueryClientProvider client={queryClient}>
      <WagmiProvider config={wagmiConfig}>
        {/*
         * BrowserRouter (pathname routing) is required for crawler-scraped
         * OG cards: social-media bots don't fetch URL fragments, so a
         * shared `/#/post/...` link has no per-post metadata to render.
         * With pathname routing, Mesh can intercept `/post/:chain/:postId`
         * server-side and emit a meta-tagged HTML page; nginx must proxy
         * `/post/*` to the Mesh container (see damm-cloud nginx.conf).
         *
         * The static asset server (nginx) does the standard SPA fallback
         * (`try_files $uri /index.html`) so deep-links still hydrate the
         * SPA even without server-side rendering for non-/post/ routes.
         */}
        <BrowserRouter>
          <App />
        </BrowserRouter>
      </WagmiProvider>
    </QueryClientProvider>
  </StrictMode>,
)
