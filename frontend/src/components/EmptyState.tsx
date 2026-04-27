export function EmptyState({ title, hint }: { title: string; hint?: string }) {
  return (
    <div className="flex flex-col items-center justify-center rounded-md border border-dashed border-neutral-800 py-16 text-center">
      <p className="text-sm text-neutral-300">{title}</p>
      {hint && <p className="mt-2 text-xs text-neutral-500">{hint}</p>}
    </div>
  )
}
