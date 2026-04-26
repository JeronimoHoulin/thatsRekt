# thatsRekt

A public-good on-chain registry of in-progress and confirmed DeFi exploits.

Whitelisted operators (typically Twitter-monitor bots watching threat-intel firms like SlowMist, BlockSec, PeckShield) post structured alerts naming attacker addresses, victim contracts, and free-form context. Other whitelisters race to vouch (upvote) or refute (downvote). Aggregates are exposed as O(1) reads so any contract — DEX router, wallet, stablecoin issuer, risk dashboard — can plug in and inline-blacklist live attacker addresses.

Designed as a public good: no economic admin power, no upgradeability, no proxies. Cross-chain identical-address deploy via the singleton CREATE2 factory.

## Architecture

- **Owner** (Safe multisig, hardcoded constant) — only role with whitelist write authority. Can be transferred via OpenZeppelin `Ownable2Step`.
- **Whitelisted addresses** — can post alerts and vote up/down on others' alerts. Cannot vote on own posts.
- **Anyone** — can read posts, attacker scores, victim flags, and the active-post linked list.

Posts contain: `address[] attackers`, `address[] victims`, `string note`. At least one must be non-empty. Up to 32 addresses total per post. Notes live in `PostCreated` events, never in storage.

## Public reads (for integrators)

```solidity
function attackerScore(address) external view returns (int256);     // signed: pick your threshold
function attackerAppearances(address) external view returns (uint256);
function isVictim(address) external view returns (bool);
function attackerReport(address) external view returns (int256 score, uint256 appearances);
```

A DEX router can `require(reg.attackerScore(user) <= 0)` before allowing a swap. A stablecoin issuer might require `attackerScore <= -2` (must have been actively refuted). The threshold is the integrator's choice; the registry is just data.

## Posting + voting

```solidity
function post(address[] attackers, address[] victims, string note) external returns (uint256 id);
function vote(uint256 postId, int8 direction) external;             // direction in {-1, 0, +1}; 0 = retract
function retract(uint256 postId) external;                          // poster only
```

`vote()` accepts `+1` (upvote), `-1` (downvote), or `0` (retract previous vote). The poster cannot vote on their own post. Same direction twice in a row reverts (`NoVoteChange`).

## Removal

A post is removed automatically when `downvotes - upvotes >= 3`, or by the poster calling `retract(id)`. Removal reverses all aggregate contributions and unlinks from the active-post list. Posts cannot be un-removed.

## On-chain feed enumeration

```solidity
function recentActivePosts(uint256 limit) external view returns (uint256[]);   // newest first
function activePostsBefore(uint256 beforeId, uint256 limit) external view returns (uint256[]);
```

Walks a doubly-linked list of non-removed posts. `MAX_VIEW_LIMIT = 100` per call. For richer queries (full-text search on notes, per-attacker post lists), consume `PostCreated` / `Voted` / `PostRemoved` events via an off-chain indexer.

## Cross-chain deploy

The contract is deployed at the same address on every supported EVM chain using the CREATE2 deployer at `0x4e59b44847b379578588920cA78FbF26c0B4956C`. Each chain has its own sovereign state — own whitelist, own posts, own karma. Cross-chain aggregation is an off-chain concern.

The governance Safe must also exist at the same address on every chain (see Safe Singleton Factory) since its address is baked into the contract bytecode as a constant.

## Build / test / deploy

```bash
forge build
forge test -vv
forge test --match-contract ThatsRektInvariants -vv

# Pre-deploy: replace GOVERNANCE constant with the real Safe address.
# The deploy script refuses to run while the dev placeholder is in place.
cp .env.example .env  # fill in PRIVATE_KEY, RPC_URL, ETHERSCAN_API_KEY
forge script script/Deploy.s.sol \
    --rpc-url <chain-rpc> \
    --broadcast \
    --verify \
    -vvvv
```

## Spec + design history

- Implementation plan: `tasks/v0-impl-plan.md` (this branch)
- Canonical design spec: DAMM Capital knowledge base (`threads/bauti/thatsrekt.md`)
- Predecessor (flat-set `addRekt(address[])` with 2-of-N propose/execute removal): see `git log master`. The current design replaces it wholesale.
