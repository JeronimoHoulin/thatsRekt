import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// IPFS-friendly build:
//   - base: './'   relative asset URLs work under any IPFS gateway path.
//   - sourcemaps off in prod to keep the bundle slim.
//   - emit assets into dist/ ready for `ipfs add -r dist`.
export default defineConfig({
  base: './',
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
