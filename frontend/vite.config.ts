import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// Build config:
//   - base: '/'  absolute asset URLs. Required because the SPA uses
//                BrowserRouter (pathname routing) so the same
//                index.html is served at deep paths like
//                `/post/base/42`. With `base: './'` the browser would
//                resolve `./assets/x.js` to `/post/base/assets/x.js`
//                — broken. Pathname routing is in turn required so
//                Mesh can SSR Open Graph cards (see mesh/src/og.ts);
//                social-media crawlers don't fetch URL fragments.
//                Tradeoff: IPFS-gateway hosting is no longer trivial
//                — the site assumes hosting at the domain root with
//                SPA-fallback (nginx `try_files $uri /index.html`).
//   - sourcemaps off in prod to keep the bundle slim.
//   - emit assets into dist/.
export default defineConfig({
  base: '/',
  plugins: [react()],
  build: {
    outDir: 'dist',
    sourcemap: false,
    target: 'es2022',
    rollupOptions: {
      output: {
        manualChunks: {
          react: ['react', 'react-dom', 'react-router-dom'],
          query: ['@tanstack/react-query', 'graphql-request'],
        },
      },
    },
  },
})
