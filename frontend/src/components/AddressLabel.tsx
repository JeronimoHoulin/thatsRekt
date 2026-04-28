import { useState } from 'react'
import { shortAddress } from '../lib/format'
import { explorerAddressUrl, getChainBySlug } from '../lib/chains'

interface AddressLabelProps {
  addr: string
  /** When set, an "open in explorer" link points at this chain's explorer. */
  chainSlug?: string
  full?: boolean
}

/**
 * Renders a monospace address with two affordances:
 *   - Click the address to copy it to the clipboard (visual confirm).
 *   - Hover to reveal an "↗" link to the chain's block explorer (when
 *     a chainSlug is supplied).
 */
export function AddressLabel({ addr, chainSlug, full = false }: AddressLabelProps) {
  const [copied, setCopied] = useState(false)
  const chain = chainSlug ? getChainBySlug(chainSlug) : undefined
  const explorerUrl = chain ? explorerAddressUrl(chain, addr) : null

  const onCopy = async () => {
    try {
      await navigator.clipboard.writeText(addr)
      setCopied(true)
      setTimeout(() => setCopied(false), 1200)
    } catch {
      // fall back: select the text so the user can ⌘C themselves
      // (older browsers / non-secure contexts don't expose clipboard API)
    }
  }

  return (
    <span className="inline-flex items-center gap-1 group">
      <button
        type="button"
        onClick={onCopy}
        title={copied ? 'copied!' : `click to copy ${addr}`}
        className="font-mono text-sm hover:bg-yellow-100 px-0.5 rounded transition-colors cursor-pointer"
      >
        {full ? addr : shortAddress(addr)}
      </button>
      {copied && (
        <span className="text-xs text-green-700 font-mono">✓ copied</span>
      )}
      {explorerUrl && !copied && (
        <a
          href={explorerUrl}
          target="_blank"
          rel="noopener noreferrer"
          title={`open on ${chain?.name ?? 'block explorer'}`}
          className="text-xs text-gray-500 hover:text-red-600 opacity-0 group-hover:opacity-100 transition-opacity"
          aria-label="Open in block explorer"
        >
          ↗
        </a>
      )}
    </span>
  )
}
