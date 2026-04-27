import { shortAddress } from '../lib/format'

export function AddressLabel({ addr, full = false }: { addr: string; full?: boolean }) {
  return (
    <span className="font-mono text-sm" title={addr}>
      {full ? addr : shortAddress(addr)}
    </span>
  )
}
