import { Routes, Route, Link } from 'react-router-dom'
import { Feed } from './pages/Feed'
import { PostDetail } from './pages/PostDetail'

export function App() {
  return (
    <div className="mx-auto flex min-h-full max-w-3xl flex-col px-4 py-8">
      <Header />
      <main className="flex-1 pt-8">
        <Routes>
          <Route path="/" element={<Feed />} />
          <Route path="/post/:id" element={<PostDetail />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </main>
      <Footer />
    </div>
  )
}

function Header() {
  return (
    <header className="flex items-baseline justify-between border-b border-neutral-800 pb-4">
      <Link to="/" className="text-2xl font-semibold tracking-tight">
        thats<span className="text-rose-500">Rekt</span>
      </Link>
      <span className="font-mono text-xs text-neutral-500">on-chain hack alert registry</span>
    </header>
  )
}

function Footer() {
  return (
    <footer className="border-t border-neutral-800 pt-4 mt-12 text-xs text-neutral-500">
      <p>
        Public good. Source:{' '}
        <a
          href="https://github.com/JeronimoHoulin/thatsRekt"
          target="_blank"
          rel="noreferrer"
          className="underline hover:text-neutral-300"
        >
          github.com/JeronimoHoulin/thatsRekt
        </a>
        .
      </p>
    </footer>
  )
}

function NotFound() {
  return (
    <div className="py-16 text-center">
      <p className="text-neutral-400">Page not found.</p>
      <Link to="/" className="mt-4 inline-block text-rose-400 underline">
        Back to feed
      </Link>
    </div>
  )
}
