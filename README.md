# thatsRekt

A minimalist on-chain register of rekt addresses, deployed at `thatsrekt.eth`.

## How it works

### The rekt list

`isRekt()` returns the current array of addresses that have been flagged as rekt. It is a set — no address appears more than once.

### Who can write

A set of **whitelisted addresses** controls the contract. Two addresses are whitelisted at deploy time and hardcoded in the contract:

| Name | Address |
|---|---|
| jerrythekid.eth | `0x9E8680dbBcA1127add812abE209A10E621b385dF` |
| bauti.eth | `0xda1b9dFA299d655135C1ECdc4f0b4c9aED9a7f45` |

Any whitelisted address can add or remove addresses from the whitelist via `addWhitelisted(address)` and `removeWhitelisted(address)`.

---

### Adding to the rekt list — `addRekt(address[])`

Any single whitelisted address can call `addRekt` with an array of addresses. They are appended to the list immediately. Addresses already on the list are silently skipped.

```bash
cast send <contract> "addRekt(address[])" "[0xAlice,0xBob]" --private-key $YOUR_KEY
```

---

### Removing from the rekt list — two steps

Removal is intentionally slower than addition: it requires **two different whitelisted addresses** to act. One proposes, a different one executes.

#### Step 1 — propose

A whitelisted address calls `proposeRemoval(address[])` with the addresses to remove. This stores the proposal on-chain and returns a numeric `id`.

```bash
cast send <contract> "proposeRemoval(address[])" "[0xAlice,0xBob]" --private-key $KEY_A
```

Read the latest proposal id:
```bash
cast call <contract> "proposalCount()(uint256)"   # subtract 1 for the last id
```

Inspect a proposal before executing:
```bash
cast call <contract> "getProposal(uint256)" <id>
# returns: (proposer address, executed bool, targets address[])
```

#### Step 2 — execute

A **different** whitelisted address calls `executeRemoval(uint256 id)`. The removal happens immediately. The proposer cannot execute their own proposal.

```bash
cast send <contract> "executeRemoval(uint256)" <id> --private-key $KEY_B
```

Addresses in the proposal that have already been removed from the list (e.g. by a concurrent execution) are silently skipped — the call never reverts because of them.

---

### Whitelist management

```bash
# Add an address to the whitelist (caller must be whitelisted)
cast send <contract> "addWhitelisted(address)" <address> --private-key $YOUR_KEY

# Remove an address from the whitelist (caller must be whitelisted)
cast send <contract> "removeWhitelisted(address)" <address> --private-key $YOUR_KEY
```

---

## Contract reference

| Function | Access | Description |
|---|---|---|
| `isRekt()` | anyone | Returns the full rekt address array |
| `getProposal(uint256 id)` | anyone | Returns proposer, executed status, and targets for a proposal |
| `proposalCount()` | anyone | Total number of proposals created so far |
| `isWhitelisted(address)` | anyone | Check whether an address is whitelisted |
| `addRekt(address[])` | whitelisted | Add addresses to the rekt list immediately |
| `proposeRemoval(address[])` | whitelisted | Propose a removal; returns a proposal `id` |
| `executeRemoval(uint256 id)` | whitelisted (not the proposer) | Execute a pending removal proposal |
| `addWhitelisted(address)` | whitelisted | Add an address to the whitelist |
| `removeWhitelisted(address)` | whitelisted | Remove an address from the whitelist |

---

## Development

### Build

```bash
forge build
```

### Test

```bash
forge test -vv
```

### Deploy

```bash
cp .env.example .env   # fill in PRIVATE_KEY and RPC_URL

forge script script/Deploy.s.sol \
  --rpc-url $RPC_URL \
  --broadcast \
  --verify \
  -vvvv
```
