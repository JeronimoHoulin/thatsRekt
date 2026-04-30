/**
 * Single-page docs for integrators. Plain JSX (no MDX, no syntax
 * highlighter) — keeps the bundle lean. Migrate to a proper docs site
 * (Docusaurus / Mintlify / vitepress) when this page outgrows the format.
 */
export function Docs() {
  return (
    <article className="space-y-12">
      <Hero />
      <WhatIs />
      <HowItWorks />
      <Integrating />
      <Reference />
    </article>
  )
}

function Hero() {
  return (
    <header className="space-y-3 border-b-2 border-black pb-6">
      <h1 className="font-black uppercase tracking-tighter text-4xl sm:text-5xl leading-none">
        docs
      </h1>
      <p className="text-xs uppercase tracking-widest text-neutral-700">
        [for protocols, integrators, and the curious]
      </p>
    </header>
  )
}

// =============================================================================
// What is thatsRekt
// =============================================================================

function WhatIs() {
  return (
    <Section heading="what is thatsRekt">
      <p className="text-base leading-relaxed text-neutral-800">
        thatsRekt is an{' '}
        <strong className="font-black">on-chain hack alert registry</strong>.
        Whitelisted operators post structured alerts about active DeFi
        exploits — attacker addresses, victim contracts, and a
        free-form note. Other whitelisters race to{' '}
        <em>vouch</em> (confirm) or <em>refute</em> (disconfirm).
      </p>
      <p className="text-base leading-relaxed text-neutral-800">
        Other contracts read this state directly: a DEX router can
        block a swap when the recipient's{' '}
        <code className="font-mono text-sm">attackerScore</code> is too
        negative, a wallet can warn the user, a stablecoin can
        circuit-break. The registry is permissioned to write but{' '}
        <strong className="font-black">open to read</strong> — every
        score, post, and confirmer set is queryable from any contract
        or app.
      </p>
    </Section>
  )
}

// =============================================================================
// How posts work
// =============================================================================

function HowItWorks() {
  return (
    <Section heading="how posts work">
      <SubSection heading="whitelisters">
        Authorized addresses (the "contributors" listed under{' '}
        <Inline>/about</Inline>). They can call{' '}
        <Code>post(...)</Code>, <Code>confirm(...)</Code>, and{' '}
        <Code>disconfirm(...)</Code>. Posts include a title, attacker
        addresses, victim contracts, and a free-form note. Confirmer
        identities are public on-chain.
      </SubSection>
      <SubSection heading="governance">
        A multisig controls the whitelist and can upgrade the contract
        — but every change goes through a{' '}
        <strong className="font-black">7-day TimelockController</strong>.
        Integrators always have a week to disengage if a malicious
        change is queued. The whitelist is the only mutable state
        managed by governance; posts themselves are
        whitelister-authored and not curated.
      </SubSection>
      <SubSection heading="integrators">
        Anyone reading the registry. Two main signals: an address's{' '}
        <Code>attackerScore</Code> (signed integer — sum of
        confirmations minus disconfirmations across every post that
        names the address as an attacker) and an address's{' '}
        <Code>isVictim</Code> flag (true if the address is currently
        the target of an active alert). Both are readable on-chain in
        a single view call.
      </SubSection>
    </Section>
  )
}

// =============================================================================
// Integrating
// =============================================================================

function Integrating() {
  return (
    <Section heading="integrating">
      <p className="text-base leading-relaxed text-neutral-800">
        Three integration paths depending on where your code lives —
        on-chain, in a dApp / indexer, or as an off-chain detection
        pipeline.
      </p>

      <SubSection heading="from a Solidity contract">
        <p className="text-sm leading-relaxed text-neutral-800 mb-3">
          Read the registry directly from your contract. The proxy is
          the same address on every chain (deterministic CREATE2
          deploy):
        </p>
        <CodeBlock>{`interface IThatsRekt {
    function attackerReport(address a)
        external view returns (int256 score, uint256 appearances);
    function isVictim(address a) external view returns (bool);
}

contract MySwapRouter {
    IThatsRekt constant rekt = IThatsRekt(0x0000000000000000000000000000000000000000);
    int256 constant ATTACKER_THRESHOLD = -3;

    function swap(address recipient, ...) external {
        (int256 score, ) = rekt.attackerReport(recipient);
        require(score > ATTACKER_THRESHOLD, "recipient flagged");
        // ... rest of swap
    }
}`}</CodeBlock>
        <p className="text-xs leading-relaxed text-neutral-700 mt-3">
          The <Inline>0x000…000</Inline> placeholder is the proxy
          address — see the <strong>reference</strong> section below
          for the live deployment per chain.
        </p>
      </SubSection>

      <SubSection heading="from a dApp or indexer">
        <p className="text-sm leading-relaxed text-neutral-800 mb-3">
          Query the public Mesh GraphQL gateway. It stitches every
          per-chain squid into a single endpoint and merges results
          across chains automatically.
        </p>
        <CodeBlock>{`# fetch the latest 10 posts across all indexed chains
query LatestPosts {
  posts(limit: 10) {
    items {
      id
      chain { slug name }
      poster
      title
      note
      attackedAt
      attackers
      victims
      confirmations
      disconfirmations
      netScore
    }
    totalCount
  }
}`}</CodeBlock>
        <p className="text-xs leading-relaxed text-neutral-700 mt-3">
          Per-chain queries are also available with the prefixed
          <Inline>{'<Chain>_postById'}</Inline> root fields — useful
          when you need the full post-detail view including
          confirmation log + edit history.
        </p>
      </SubSection>

      <SubSection heading="from an off-chain detection pipeline">
        <p className="text-sm leading-relaxed text-neutral-800">
          If you're a whitelisted operator running an automated
          detector and want a webhook-driven submission path, the{' '}
          <Inline>relay/</Inline> service in the monorepo provides
          one. Single-tenant per deployment — bring your own EOA + a
          bearer token. See{' '}
          <a
            href="https://github.com/JeronimoHoulin/thatsRekt/blob/master/relay/README.md"
            target="_blank"
            rel="noopener noreferrer"
            className="rekt-link"
          >
            relay/README.md ↗
          </a>{' '}
          for the full spec.
        </p>
      </SubSection>
    </Section>
  )
}

// =============================================================================
// Reference
// =============================================================================

function Reference() {
  return (
    <Section heading="reference">
      <SubSection heading="deployments">
        <p className="text-sm leading-relaxed text-neutral-800 mb-3">
          The proxy address is{' '}
          <strong className="font-black">stable across chains</strong>{' '}
          via CREATE2 — same address everywhere. Per-chain status:
        </p>
        <div className="overflow-x-auto border-2 border-black">
          <table className="w-full text-left text-sm">
            <thead className="border-b-2 border-black bg-black/5 text-xs uppercase tracking-widest">
              <tr>
                <th className="px-3 py-2">chain</th>
                <th className="px-3 py-2">chain id</th>
                <th className="px-3 py-2">proxy</th>
                <th className="px-3 py-2">status</th>
              </tr>
            </thead>
            <tbody className="divide-y divide-black font-mono text-xs">
              <tr>
                <td className="px-3 py-2 font-black">base</td>
                <td className="px-3 py-2 tabular-nums">8453</td>
                <td className="px-3 py-2 text-neutral-600">— TBD —</td>
                <td className="px-3 py-2 uppercase tracking-widest text-amber-700">
                  pending deploy
                </td>
              </tr>
              <tr>
                <td className="px-3 py-2 font-black">optimism</td>
                <td className="px-3 py-2 tabular-nums">10</td>
                <td className="px-3 py-2 text-neutral-600">— TBD —</td>
                <td className="px-3 py-2 uppercase tracking-widest text-amber-700">
                  pending deploy
                </td>
              </tr>
            </tbody>
          </table>
        </div>
        <p className="text-xs leading-relaxed text-neutral-700 mt-3">
          Addresses will populate here once the production contracts
          are deployed. Until then, see{' '}
          <Inline>/contributors</Inline> on the contracts/ subdir of
          the monorepo for testnet addresses and{' '}
          <Inline>contracts/script/Deploy.s.sol</Inline> for the
          deploy script.
        </p>
      </SubSection>

      <SubSection heading="public endpoints">
        <ul className="space-y-2 text-sm">
          <li>
            <strong className="font-black">GraphQL gateway:</strong>{' '}
            <Inline>https://thatsrekt.com/graphql</Inline>
          </li>
          <li>
            <strong className="font-black">Frontend:</strong>{' '}
            <Inline>https://thatsrekt.com</Inline>
          </li>
          <li>
            <strong className="font-black">Source:</strong>{' '}
            <a
              href="https://github.com/JeronimoHoulin/thatsRekt"
              target="_blank"
              rel="noopener noreferrer"
              className="rekt-link"
            >
              github.com/JeronimoHoulin/thatsRekt ↗
            </a>
          </li>
        </ul>
      </SubSection>
    </Section>
  )
}

// =============================================================================
// Layout primitives
// =============================================================================

function Section({
  heading,
  children,
}: {
  heading: string
  children: React.ReactNode
}) {
  return (
    <section className="space-y-4">
      <h2 className="font-black uppercase tracking-tighter text-2xl sm:text-3xl leading-none">
        {heading}
      </h2>
      {children}
    </section>
  )
}

function SubSection({
  heading,
  children,
}: {
  heading: string
  children: React.ReactNode
}) {
  return (
    <div className="space-y-2">
      <h3 className="font-black uppercase tracking-widest text-xs">
        {heading}
      </h3>
      {children}
    </div>
  )
}

function CodeBlock({ children }: { children: string }) {
  return (
    <pre className="overflow-x-auto border-2 border-black bg-neutral-50 p-4 text-xs leading-relaxed font-mono">
      <code>{children}</code>
    </pre>
  )
}

function Code({ children }: { children: React.ReactNode }) {
  return (
    <code className="font-mono text-sm bg-neutral-100 border border-neutral-300 px-1 py-0.5">
      {children}
    </code>
  )
}

function Inline({ children }: { children: React.ReactNode }) {
  return <code className="font-mono text-sm">{children}</code>
}
