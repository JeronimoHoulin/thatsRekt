# thatsRekt monorepo

Public-good on-chain hack alert registry. Whitelisted operators post structured alerts about active DeFi exploits; other whitelisters race to vouch (confirm) or refute (disconfirm). Other contracts (DEX routers, wallets, stablecoins) plug in and inline-blacklist live attacker addresses.

## Layout

| Directory | Purpose |
|-----------|---------|
| [`contracts/`](contracts/) | Solidity smart contracts (Foundry). UUPS upgradeable proxy with TimelockController. |
| [`indexer/`](indexer/) | Subsquid indexer (TypeScript). Per-chain stack indexes contract events into postgres. |
| [`mesh/`](mesh/) | GraphQL stitching gateway (TypeScript + `@graphql-tools/stitch` + `graphql-yoga`). Single public GraphQL surface; fans out to per-chain squids. |
| [`frontend/`](frontend/) | Static IPFS-compatible web app (Vite + React + Tailwind). Browses the registry. |
| [`relay/`](relay/) | Optional Go service for whitelisted posters who want a webhook-driven submission path (e.g. an automated detection pipeline). Single-tenant per deployment — operator brings their own EOA. |
| [`ops/`](ops/) | Local-dev orchestration: Makefile + scripts to bring up the dual-anvil + Mesh + frontend stack on a laptop. |
| [`data/`](data/) | Reference data (e.g. historic incidents) for seeding/testing. |

## Quick start

- Smart contracts: see [contracts/README.md](contracts/README.md).
- Indexer: see [indexer/README.md](indexer/README.md).
- Mesh gateway: see [mesh/README.md](mesh/README.md).
- Frontend: see [frontend/README.md](frontend/README.md).
- Relay (optional): see [relay/README.md](relay/README.md).
- Full-stack local walkthrough: see [ops/README.md](ops/README.md).

## License

[MIT](./LICENSE).
