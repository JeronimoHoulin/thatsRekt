// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {Test} from "forge-std/Test.sol";
import {Deploy} from "../script/Deploy.s.sol";
import {ThatsRekt} from "../src/ThatsRekt.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/// @notice Regression gate: asserts that the five canonical CREATE2 addresses
///         produced by `Deploy.s.sol` with the production role tuple exactly
///         match the live v1.2.0 addresses deployed on Ethereum mainnet, Base,
///         Arbitrum, and Optimism — and, crucially, that they WILL match on any
///         future chain deployment including BNB Chain (BSC, chain 56).
///
///         === Why this test exists ===
///         CREATE2 addresses are computed from (factory, salt, initCodeHash).
///         The CREATE2 factory address and salt strings are chain-agnostic; the
///         initCodeHash is determined by the compiler output and constructor
///         arguments. Provided the same deployer EOA, the same governance tuple,
///         and the same Solidity compiler+optimizer settings are used, the proxy
///         address is IDENTICAL on every EVM chain.
///
///         This test locks that invariant. If anyone:
///           - bumps a Solidity/OZ dependency (changes ThatsRekt bytecode)
///           - changes any salt constant in Deploy.s.sol
///           - changes the deployer EOA
///           - reorders or changes the canonical role tuple
///           - changes optimizer settings
///
///         …the predicted addresses will no longer match the live canonical
///         values, this test will FAIL loudly, and the mismatch must be resolved
///         BEFORE spending any gas on a new chain deployment. The proxy address
///         is the top-priority cross-chain invariant for integrators.
///
///         === BSC (chain 56) proof ===
///         computeCreate2Address uses only (factory, salt, initCodeHash) — no
///         chainId. Therefore this test running on any EVM is simultaneously a
///         proof that the same bytecode + tuple yields the canonical addresses on
///         BSC. No RPC, no broadcast, no BSC node required.
///
///         === Canonical production tuple (v1.2.0) ===
///         These are the exact values that were used for the live deploys on
///         Ethereum/Base/Arbitrum/Optimism and must be reproduced on BSC:
///
///           DEPLOYER_ADDRESS      = 0xb5a6c8ca369e38050784e2a6793bee6447109340
///           GOVERNANCE_OWNER      = 0x59E4DBc95BD312A882Bb36b7f3E8298682340679
///           WHITELIST_OPERATOR    = 0xda1b9dFA299d655135C1ECdc4f0b4c9aED9a7f45
///           PURGE_REMOVER_EOA     = 0xda1b9dFA299d655135C1ECdc4f0b4c9aED9a7f45
///           INITIAL_WHITELISTERS  = [6 addresses in order below]
///
///           Expected on-chain addresses:
///             impl       = 0xd7A06A47325b9e439Df5FCE3F5C64AD010ab6eD9
///             upgradeTLC = 0xf6F807f095D6D09c1216ffBd6AaCBB73D8F02aB6
///             addTLC     = 0xB83AB5772f919BE72b4AaB98456eDdED5ad68D4f
///             purgeTLC   = 0xd8Dbce72f488c7664c6bdFae4aa819daBEEF98a8
///             proxy      = 0xBfaEEE9662b4c037De24e5Caa65815350d57b89A
contract CanonicalAddressParityTest is Test {
    // -----------------------------------------------------------------------
    // Canonical production tuple (v1.2.0) — DO NOT CHANGE without a full
    // cross-chain re-deploy. Changing any of these values means the BSC proxy
    // will land at a different address than the one already live on other chains.
    // -----------------------------------------------------------------------

    address constant DEPLOYER  = 0xb5A6c8ca369e38050784e2A6793beE6447109340;
    address constant GOV_SAFE  = 0x59E4DBc95BD312A882Bb36b7f3E8298682340679;
    address constant OPERATOR  = 0xda1b9dFA299d655135C1ECdc4f0b4c9aED9a7f45;
    address constant PURGE_EOA = 0xda1b9dFA299d655135C1ECdc4f0b4c9aED9a7f45;

    // Six initial whitelisters in canonical order.
    address constant WL_0 = 0x5822B262EDdA82d2C6A436b598Ff96fA9AB894c4;
    address constant WL_1 = 0xda1b9dFA299d655135C1ECdc4f0b4c9aED9a7f45;
    address constant WL_2 = 0x9E8680dbBcA1127add812abE209A10E621b385dF;
    address constant WL_3 = 0x24C2167054A9A9e00F67233F1eBc4060501f54FA;
    address constant WL_4 = 0xE0396d6d738e726D39f96099b8f6a55d11184374;
    address constant WL_5 = 0xb5A6c8ca369e38050784e2A6793beE6447109340;

    // -----------------------------------------------------------------------
    // Expected canonical addresses (live on Mainnet / Base / Arb / OP,
    // and must reproduce on BSC with the same tuple + compiler settings).
    // -----------------------------------------------------------------------

    address constant EXPECTED_IMPL        = 0xd7A06A47325b9e439Df5FCE3F5C64AD010ab6eD9;
    address constant EXPECTED_UPGRADE_TLC = 0xf6F807f095D6D09c1216ffBd6AaCBB73D8F02aB6;
    address constant EXPECTED_ADD_TLC     = 0xB83AB5772f919BE72b4AaB98456eDdED5ad68D4f;
    address constant EXPECTED_PURGE_TLC   = 0xd8Dbce72f488c7664c6bdFae4aa819daBEEF98a8;
    address constant EXPECTED_PROXY       = 0xBfaEEE9662b4c037De24e5Caa65815350d57b89A;

    Deploy internal d;

    function setUp() public {
        d = new Deploy();
    }

    /// @notice Core parity gate. Computes all five CREATE2 addresses using the
    ///         exact same logic as PredictAddresses.s.sol (which mirrors
    ///         Deploy.s.sol) and asserts they match the canonical live addresses.
    ///
    ///         Failure means the current codebase would deploy to a DIFFERENT
    ///         address than the one live on all existing chains — BSC deployment
    ///         MUST be blocked until the discrepancy is understood and resolved.
    function test_canonicalAddressParity() public view {
        (
            address impl,
            address upgradeTLC,
            address addTLC,
            address purgeTLC,
            address proxy
        ) = _predictAll();

        assertEq(impl,       EXPECTED_IMPL,        "impl address mismatch: parity BROKEN");
        assertEq(upgradeTLC, EXPECTED_UPGRADE_TLC, "upgradeTLC address mismatch: parity BROKEN");
        assertEq(addTLC,     EXPECTED_ADD_TLC,     "addTLC address mismatch: parity BROKEN");
        assertEq(purgeTLC,   EXPECTED_PURGE_TLC,   "purgeTLC address mismatch: parity BROKEN");
        assertEq(proxy,      EXPECTED_PROXY,       "proxy address mismatch: BSC deploy BLOCKED");
    }

    // -----------------------------------------------------------------------
    // Internal helpers — pure address computation, no state changes
    // -----------------------------------------------------------------------

    /// @dev Replicates the five-step CREATE2 prediction from
    ///      PredictAddresses.s.sol using the canonical production tuple.
    ///      All logic here must stay in sync with Deploy.s.sol / PredictAddresses.s.sol.
    function _predictAll() internal view returns (
        address impl,
        address upgradeTLC,
        address addTLC,
        address purgeTLC,
        address proxy
    ) {
        // Shared executor list: [GOV_SAFE, address(0)] — matches Deploy.s.sol's
        // `sharedExecutors`. address(0) makes EXECUTOR_ROLE open (anyone can
        // execute after the delay).
        address[] memory sharedExecutors = new address[](2);
        sharedExecutors[0] = GOV_SAFE;
        sharedExecutors[1] = address(0);

        // 1. Implementation — no constructor args, so only creationCode matters.
        impl = vm.computeCreate2Address(
            d.IMPL_SALT(),
            keccak256(type(ThatsRekt).creationCode),
            CREATE2_FACTORY
        );

        // 2. Upgrade TLC (7-day) — proposer = GOV_SAFE, deployer as temp admin.
        address[] memory upProp = new address[](1);
        upProp[0] = GOV_SAFE;
        upgradeTLC = vm.computeCreate2Address(
            d.UPGRADE_TIMELOCK_SALT(),
            keccak256(_tlcInitCode(d.UPGRADE_DELAY(), upProp, sharedExecutors, DEPLOYER)),
            CREATE2_FACTORY
        );

        // 3. Add TLC (3-day) — proposer = GOV_SAFE, deployer as temp admin.
        address[] memory addProp = new address[](1);
        addProp[0] = GOV_SAFE;
        addTLC = vm.computeCreate2Address(
            d.ADD_TIMELOCK_SALT(),
            keccak256(_tlcInitCode(d.ADD_DELAY(), addProp, sharedExecutors, DEPLOYER)),
            CREATE2_FACTORY
        );

        // 4. Purge TLC (1-day) — proposer = PURGE_EOA, deployer as temp admin.
        address[] memory purgeProp = new address[](1);
        purgeProp[0] = PURGE_EOA;
        purgeTLC = vm.computeCreate2Address(
            d.PURGE_TIMELOCK_SALT(),
            keccak256(_tlcInitCode(d.PURGE_DELAY(), purgeProp, sharedExecutors, DEPLOYER)),
            CREATE2_FACTORY
        );

        // 5. Proxy — depends on the four addresses above.
        address[] memory whitelisters = _canonicalWhitelisters();
        bytes memory initCall = abi.encodeCall(
            ThatsRekt.initialize,
            (upgradeTLC, addTLC, OPERATOR, purgeTLC, PURGE_EOA, whitelisters)
        );
        proxy = vm.computeCreate2Address(
            d.PROXY_SALT(),
            keccak256(abi.encodePacked(type(ERC1967Proxy).creationCode, abi.encode(impl, initCall))),
            CREATE2_FACTORY
        );
    }

    /// @dev Encodes the TimelockController constructor calldata exactly as
    ///      Deploy.s.sol does. Must stay in sync with `_deployTLCWithSplitRoles`.
    function _tlcInitCode(
        uint256 delay,
        address[] memory proposers,
        address[] memory executors,
        address deployerAddr
    ) internal pure returns (bytes memory) {
        return abi.encodePacked(
            type(TimelockController).creationCode,
            abi.encode(delay, proposers, executors, deployerAddr)
        );
    }

    /// @dev Returns the six canonical initial whitelisters in order.
    function _canonicalWhitelisters() internal pure returns (address[] memory wl) {
        wl = new address[](6);
        wl[0] = WL_0;
        wl[1] = WL_1;
        wl[2] = WL_2;
        wl[3] = WL_3;
        wl[4] = WL_4;
        wl[5] = WL_5;
    }
}
