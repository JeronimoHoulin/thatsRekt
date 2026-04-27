export function EmptyState({ title, hint }: { title: string; hint?: string }) {
  return (
    <div className="border-2 border-dashed border-neutral-400 py-16 text-center">
      <p className="text-lg font-black uppercase tracking-tight">{title}</p>
      {hint && (
        <p className="mx-auto mt-3 max-w-md text-xs uppercase tracking-widest text-neutral-700">
          {hint}
        </p>
      )}
    </div>
  )
}
