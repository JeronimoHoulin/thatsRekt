# thatsRekt monorepo

Public-good on-chain hack alert registry. Whitelisted operators post structured alerts about active DeFi exploits; other whitelisters race to vouch (upvote) or refute (downvote). Other contracts (DEX routers, wallets, stablecoins) plug in and inline-blacklist live attacker addresses.

## Layout

| Directory | Purpose |
|-----------|---------|
| [`contracts/`](contracts/) | Solidity smart contracts (Foundry). UUPS upgradeable proxy with TimelockController. |
| [`indexer/`](indexer/) | Subsquid indexer (TypeScript). Indexes contract events, exposes GraphQL API. |
| [`frontend/`](frontend/) | Static IPFS-compatible web app (Vite + React + Tailwind). Browses the registry. |

## Quick start

- Smart contracts: see [contracts/README.md](contracts/README.md).
- Indexer: see [indexer/README.md](indexer/README.md).
- Frontend: see [frontend/README.md](frontend/README.md).

## Project status

Pre-deployment. Contract design complete (see [contracts/tasks/](contracts/tasks/)). Indexer scaffolded with full schema + event handlers (Phases 1-5). Frontend scaffolded with feed + post detail + edit/vote timeline. Awaiting first deployment to wire `CONTRACT_ADDRESS` + `START_BLOCK` for the indexer; frontend will then have live data.
