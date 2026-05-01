// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {Test} from "forge-std/Test.sol";
import {DeployDev} from "../script/DeployDev.s.sol";
import {Deploy} from "../script/Deploy.s.sol";
import {ThatsRekt} from "../src/ThatsRekt.sol";
import {TimelockController} from "@openzeppelin/contracts/governance/TimelockController.sol";
import {ERC1967Proxy} from "@openzeppelin/contracts/proxy/ERC1967/ERC1967Proxy.sol";

/// @notice Tests for the dev/testnet deploy script. Verifies it
///         (a) accepts an EOA owner (the production script's main
///             blocker on testnet),
///         (b) wires the five-role system correctly across all
///             cross-canceller modes (vacuous single-principal,
///             default cold-wallet split, fully-explicit four-way),
///         (c) renounces DEFAULT_ADMIN_ROLE on all three TLCs (role
///             config locked permanently after deploy),
///         (d) uses CREATE2 salts that cannot collide with production
///             Deploy.s.sol salts.
///
/// @dev Tests call `deploy(...)` directly rather than going through the
///      env-reading `run()` because Foundry runs tests in parallel and
///      `vm.setEnv` mutates the OS process env (race conditions). A
///      single dedicated test exercises the env-reading path serially.
///
///      The "deployer" parameter passed to `deploy()` is `address(deployer)`
///      (the script contract instance) — that's the address that will
///      hold the temporary DEFAULT_ADMIN_ROLE during the dance, since
///      `vm.startBroadcast(addr)` in test mode makes msg.sender == addr
///      for subsequent calls.
contract DeployDevTest is Test {
    /// @dev Anvil default account 0 — the recommended dev EOA.
    address constant DEV_EOA = 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266;

    /// @dev Canonical cold wallet — same as `DeployDev.DEFAULT_COLD_WALLET`.
    address constant COLD = 0x5822B262EDdA82d2C6A436b598Ff96fA9AB894c4;

    DeployDev internal deployer;

    function setUp() public {
        deployer = new DeployDev();
    }

    /// @notice Single-principal vacuous mode: every role collapses to one
    ///         EOA. The dance still runs end-to-end (renounce locks the
    ///         role config), and grant/revoke is skipped because grant
    ///         would no-op and revoke would strip the canceller's
    ///         (==proposer's) role.
    /// @dev    The single-arg overload uses the passed address as BOTH
    ///         the deployer and every role; prediction must encode it
    ///         as the temp admin in TLC init code.
    function test_deploy_single_principal_vacuous() public {
        deployer.deploy(DEV_EOA);

        (address impl, address upgrade, address add_, address purge, address proxy)
            = _predictAll(DEV_EOA, DEV_EOA, DEV_EOA, DEV_EOA);

        assertGt(impl.code.length, 0,    "impl not deployed");
        assertGt(upgrade.code.length, 0, "upgrade timelock not deployed");
        assertGt(add_.code.length, 0,    "add timelock not deployed");
        assertGt(purge.code.length, 0,   "purge timelock not deployed");
        assertGt(proxy.code.length, 0,   "proxy not deployed");

        // Proxy wiring (vacuous: every slot points to the same EOA or its TLC).
        assertEq(ThatsRekt(proxy).owner(),            upgrade, "proxy.owner != upgrade timelock");
        assertEq(ThatsRekt(proxy).whitelistAdmin(),   add_,    "proxy.whitelistAdmin != add timelock");
        assertEq(ThatsRekt(proxy).whitelistRemover(), DEV_EOA, "proxy.whitelistRemover != DEV_EOA");
        assertEq(ThatsRekt(proxy).purgeAdmin(),       purge,   "proxy.purgeAdmin != purge timelock");
        assertEq(ThatsRekt(proxy).purgeRemover(),     DEV_EOA, "proxy.purgeRemover != DEV_EOA");

        TimelockController upTl    = TimelockController(payable(upgrade));
        TimelockController addTl   = TimelockController(payable(add_));
        TimelockController purgeTl = TimelockController(payable(purge));

        // Vacuous: DEV_EOA is proposer AND canceller on every TLC.
        assertTrue(upTl.hasRole(upTl.PROPOSER_ROLE(), DEV_EOA),  "EOA missing PROPOSER on upgrade TLC");
        assertTrue(upTl.hasRole(upTl.CANCELLER_ROLE(), DEV_EOA), "EOA missing CANCELLER on upgrade TLC");

        assertTrue(addTl.hasRole(addTl.PROPOSER_ROLE(), DEV_EOA),  "EOA missing PROPOSER on add TLC");
        assertTrue(addTl.hasRole(addTl.CANCELLER_ROLE(), DEV_EOA), "EOA missing CANCELLER on add TLC");

        assertTrue(purgeTl.hasRole(purgeTl.PROPOSER_ROLE(), DEV_EOA),  "EOA missing PROPOSER on purge TLC");
        assertTrue(purgeTl.hasRole(purgeTl.CANCELLER_ROLE(), DEV_EOA), "EOA missing CANCELLER on purge TLC");

        // Delays match production.
        assertEq(upTl.getMinDelay(),    7 days, "upgrade TLC delay drifted");
        assertEq(addTl.getMinDelay(),   3 days, "add TLC delay drifted");
        assertEq(purgeTl.getMinDelay(), 1 days, "purge TLC delay drifted");

        // CRITICAL: DEFAULT_ADMIN_ROLE renounced on all three TLCs.
        // After step 4 of the dance, only the TLC itself holds admin.
        _assertAdminLockedDown(upTl,    DEV_EOA);
        _assertAdminLockedDown(addTl,   DEV_EOA);
        _assertAdminLockedDown(purgeTl, DEV_EOA);
    }

    /// @notice Default cross-canceller mode (cold wallet on whitelist
    ///         remover + purge remover, deployer-EOA on governance):
    ///         the realistic Sepolia rehearsal scenario.
    function test_deploy_default_cross_canceller() public {
        // Deployer EOA = governanceOwner; cold wallet = whitelist+purge remover.
        deployer.deploy(address(deployer), DEV_EOA, COLD, COLD);

        (address impl, address upgrade, address add_, address purge, address proxy)
            = _predictAll(address(deployer), DEV_EOA, COLD, COLD);

        assertGt(impl.code.length, 0);
        assertGt(upgrade.code.length, 0);
        assertGt(add_.code.length, 0);
        assertGt(purge.code.length, 0);
        assertGt(proxy.code.length, 0);

        TimelockController upTl    = TimelockController(payable(upgrade));
        TimelockController addTl   = TimelockController(payable(add_));
        TimelockController purgeTl = TimelockController(payable(purge));

        // Upgrade TLC: DEV_EOA proposer, COLD canceller — non-vacuous split.
        _assertProposerOnly(upTl, DEV_EOA);
        _assertCancellerOnly(upTl, COLD);
        // Crucially, proposer is NOT the canceller (the security invariant).
        assertFalse(upTl.hasRole(upTl.CANCELLER_ROLE(), DEV_EOA), "DEV_EOA is canceller on upgrade (defeats split)");

        // Add TLC: same split as upgrade.
        _assertProposerOnly(addTl, DEV_EOA);
        _assertCancellerOnly(addTl, COLD);
        assertFalse(addTl.hasRole(addTl.CANCELLER_ROLE(), DEV_EOA), "DEV_EOA is canceller on add (defeats split)");

        // Purge TLC: COLD proposer, DEV_EOA canceller — opposite direction.
        _assertProposerOnly(purgeTl, COLD);
        _assertCancellerOnly(purgeTl, DEV_EOA);
        assertFalse(purgeTl.hasRole(purgeTl.CANCELLER_ROLE(), COLD), "COLD is canceller on purge (defeats split)");

        // Admin renounced on all three.
        _assertAdminLockedDown(upTl,    address(deployer));
        _assertAdminLockedDown(addTl,   address(deployer));
        _assertAdminLockedDown(purgeTl, address(deployer));

        // Proxy slots match the role tuple.
        assertEq(ThatsRekt(proxy).whitelistRemover(), COLD,    "proxy.whitelistRemover != cold");
        assertEq(ThatsRekt(proxy).purgeRemover(),     COLD,    "proxy.purgeRemover != cold");
    }

    /// @notice Fully-explicit four-way split: every role is a distinct
    ///         address. Useful for verifying every wire is independent.
    function test_deploy_fully_explicit_four_way() public {
        address gov     = makeAddr("gov-stand-in");
        address whlRem  = makeAddr("whitelist-remover-stand-in");
        address pgrRem  = makeAddr("purge-remover-stand-in");
        deployer.deploy(address(deployer), gov, whlRem, pgrRem);

        (, address upgrade, address add_, address purge, address proxy)
            = _predictAll(address(deployer), gov, whlRem, pgrRem);

        TimelockController upTl    = TimelockController(payable(upgrade));
        TimelockController addTl   = TimelockController(payable(add_));
        TimelockController purgeTl = TimelockController(payable(purge));

        // Upgrade: gov proposes, whlRem cancels.
        _assertProposerOnly(upTl, gov);
        _assertCancellerOnly(upTl, whlRem);
        // Add: gov proposes, whlRem cancels.
        _assertProposerOnly(addTl, gov);
        _assertCancellerOnly(addTl, whlRem);
        // Purge: pgrRem proposes, gov cancels.
        _assertProposerOnly(purgeTl, pgrRem);
        _assertCancellerOnly(purgeTl, gov);

        // No role bleeds across logical boundaries.
        assertFalse(upTl.hasRole(upTl.PROPOSER_ROLE(), pgrRem),     "purge-rem unexpectedly proposer on upgrade");
        assertFalse(addTl.hasRole(addTl.PROPOSER_ROLE(), pgrRem),    "purge-rem unexpectedly proposer on add");
        assertFalse(purgeTl.hasRole(purgeTl.PROPOSER_ROLE(), whlRem), "whitelist-rem unexpectedly proposer on purge");

        // Proxy slots align.
        assertEq(ThatsRekt(proxy).whitelistRemover(), whlRem,  "proxy.whitelistRemover != whlRem");
        assertEq(ThatsRekt(proxy).purgeRemover(),     pgrRem,  "proxy.purgeRemover != pgrRem");
    }

    /// @notice Behavioral test: the proposer schedules an op, attempts
    ///         to cancel it, and is rejected. The cross-canceller can
    ///         then cancel successfully.
    function test_upgradeTLC_proposer_cannot_cancel() public {
        deployer.deploy(address(deployer), DEV_EOA, COLD, COLD);
        (, address upgrade,,,) = _predictAll(address(deployer), DEV_EOA, COLD, COLD);
        TimelockController tlc = TimelockController(payable(upgrade));

        // Schedule a no-op call. Use predicted hash so we can target it.
        address target = makeAddr("dummy-target");
        bytes memory data = "";
        bytes32 predecessor = bytes32(0);
        bytes32 salt = bytes32(uint256(1));

        vm.prank(DEV_EOA);
        tlc.schedule(target, 0, data, predecessor, salt, 7 days);
        bytes32 opId = tlc.hashOperation(target, 0, data, predecessor, salt);
        assertTrue(tlc.isOperationPending(opId), "op not scheduled");

        // Proposer DEV_EOA tries to cancel — should revert (no CANCELLER_ROLE).
        vm.prank(DEV_EOA);
        vm.expectRevert(); // AccessControlUnauthorizedAccount(DEV_EOA, CANCELLER_ROLE)
        tlc.cancel(opId);

        // Canceller COLD cancels — succeeds.
        vm.prank(COLD);
        tlc.cancel(opId);
        assertFalse(tlc.isOperationPending(opId), "op still pending after canceller cancel");
    }

    /// @notice Same behavioral test for the add TLC.
    function test_addTLC_proposer_cannot_cancel() public {
        deployer.deploy(address(deployer), DEV_EOA, COLD, COLD);
        (,, address add_,,) = _predictAll(address(deployer), DEV_EOA, COLD, COLD);
        TimelockController tlc = TimelockController(payable(add_));

        address target = makeAddr("dummy-add-target");
        bytes32 salt = bytes32(uint256(2));

        vm.prank(DEV_EOA);
        tlc.schedule(target, 0, "", bytes32(0), salt, 3 days);
        bytes32 opId = tlc.hashOperation(target, 0, "", bytes32(0), salt);

        vm.prank(DEV_EOA);
        vm.expectRevert();
        tlc.cancel(opId);

        vm.prank(COLD);
        tlc.cancel(opId);
        assertFalse(tlc.isOperationPending(opId));
    }

    /// @notice Behavioral test for the purge TLC: proposer is the cold
    ///         wallet, canceller is the deployer EOA (governance).
    function test_purgeTLC_proposer_cannot_cancel() public {
        deployer.deploy(address(deployer), DEV_EOA, COLD, COLD);
        (,,, address purge,) = _predictAll(address(deployer), DEV_EOA, COLD, COLD);
        TimelockController tlc = TimelockController(payable(purge));

        address target = makeAddr("dummy-purge-target");
        bytes32 salt = bytes32(uint256(3));

        // Proposer is COLD on the purge TLC.
        vm.prank(COLD);
        tlc.schedule(target, 0, "", bytes32(0), salt, 1 days);
        bytes32 opId = tlc.hashOperation(target, 0, "", bytes32(0), salt);

        // Proposer COLD cannot cancel.
        vm.prank(COLD);
        vm.expectRevert();
        tlc.cancel(opId);

        // Canceller (DEV_EOA, the governance side) cancels.
        vm.prank(DEV_EOA);
        tlc.cancel(opId);
        assertFalse(tlc.isOperationPending(opId));
    }

    /// @notice Idempotent: running twice on the same chain is a no-op
    ///         the second time (everything's already deployed; no re-dance).
    function test_deploy_idempotent() public {
        deployer.deploy(DEV_EOA);
        deployer.deploy(DEV_EOA);
    }

    /// @notice Zero deployer is rejected on the programmatic path.
    function test_deploy_rejects_zero_deployer() public {
        vm.expectRevert("deployer is zero");
        deployer.deploy(address(0), DEV_EOA, DEV_EOA, DEV_EOA);
    }

    /// @notice Zero governance owner rejected.
    function test_deploy_rejects_zero_governance_owner() public {
        vm.expectRevert("governanceOwner is zero");
        deployer.deploy(address(deployer), address(0), DEV_EOA, DEV_EOA);
    }

    /// @notice Zero whitelist remover rejected.
    function test_deploy_rejects_zero_whitelist_remover() public {
        vm.expectRevert("whitelistRemover is zero");
        deployer.deploy(address(deployer), DEV_EOA, address(0), DEV_EOA);
    }

    /// @notice Zero purge remover rejected.
    function test_deploy_rejects_zero_purge_remover() public {
        vm.expectRevert("purgeRemover is zero");
        deployer.deploy(address(deployer), DEV_EOA, DEV_EOA, address(0));
    }

    /// @notice Production salts and dev salts MUST be different. A
    ///         collision would mean a dev deploy could squat on the
    ///         production CREATE2 address (or vice versa).
    function test_salts_distinct_from_production() public {
        Deploy prod = new Deploy();
        assertTrue(deployer.IMPL_SALT()             != prod.IMPL_SALT(),             "IMPL_SALT collision");
        assertTrue(deployer.UPGRADE_TIMELOCK_SALT() != prod.UPGRADE_TIMELOCK_SALT(), "UPGRADE_TIMELOCK_SALT collision");
        assertTrue(deployer.ADD_TIMELOCK_SALT()     != prod.ADD_TIMELOCK_SALT(),     "ADD_TIMELOCK_SALT collision");
        assertTrue(deployer.PURGE_TIMELOCK_SALT()   != prod.PURGE_TIMELOCK_SALT(),   "PURGE_TIMELOCK_SALT collision");
        assertTrue(deployer.PROXY_SALT()            != prod.PROXY_SALT(),            "PROXY_SALT collision");
        assertTrue(
            deployer.UPGRADE_TIMELOCK_SALT() != deployer.ADD_TIMELOCK_SALT(),
            "UPGRADE_TIMELOCK_SALT == ADD_TIMELOCK_SALT (would collide)"
        );
        assertTrue(
            deployer.UPGRADE_TIMELOCK_SALT() != deployer.PURGE_TIMELOCK_SALT(),
            "UPGRADE_TIMELOCK_SALT == PURGE_TIMELOCK_SALT (would collide)"
        );
        assertTrue(
            deployer.ADD_TIMELOCK_SALT() != deployer.PURGE_TIMELOCK_SALT(),
            "ADD_TIMELOCK_SALT == PURGE_TIMELOCK_SALT (would collide)"
        );
    }

    /// @notice Same role tuple + same deployer ⇒ same proxy address.
    ///         Different deployer EOA ⇒ different proxy (since deployer
    ///         is now part of TLC init code).
    function test_same_inputs_yield_same_proxy_address() public view {
        (,,,, address proxy1) = _predictAll(address(deployer), DEV_EOA, COLD, COLD);
        (,,,, address proxy2) = _predictAll(address(deployer), DEV_EOA, COLD, COLD);
        assertEq(proxy1, proxy2, "prediction not pure");

        // A different deployer must yield a different proxy address.
        (,,,, address proxy3) = _predictAll(address(0xDEADBEEF), DEV_EOA, COLD, COLD);
        assertTrue(proxy1 != proxy3, "different deployers should yield different proxies");

        // A different governance owner must yield a different proxy address.
        (,,,, address proxy4) = _predictAll(address(deployer), address(0x1234), COLD, COLD);
        assertTrue(proxy1 != proxy4, "different gov owners should yield different proxies");
    }

    // -------------------------------------------------------------------
    // helpers
    // -------------------------------------------------------------------

    /// @dev Predicts the canonical CREATE2 addresses for a given role
    ///      tuple. Returns (impl, upgrade, add, purge, proxy).
    function _predictAll(
        address deployerAddr,
        address governanceOwner,
        address whitelistRemover,
        address purgeRemover
    ) internal view returns (
        address impl,
        address upgrade,
        address add_,
        address purge,
        address proxy
    ) {
        impl    = _predict(deployer.IMPL_SALT(),             keccak256(type(ThatsRekt).creationCode));
        address[] memory executors = new address[](2);
        executors[0] = governanceOwner;
        executors[1] = address(0);
        upgrade = _predict(deployer.UPGRADE_TIMELOCK_SALT(), keccak256(_timelockInitCode(7 days, governanceOwner, executors, deployerAddr)));
        add_    = _predict(deployer.ADD_TIMELOCK_SALT(),     keccak256(_timelockInitCode(3 days, governanceOwner, executors, deployerAddr)));
        purge   = _predict(deployer.PURGE_TIMELOCK_SALT(),   keccak256(_timelockInitCode(1 days, purgeRemover, executors, deployerAddr)));
        proxy   = _predict(
            deployer.PROXY_SALT(),
            keccak256(_proxyInitCode(impl, upgrade, add_, whitelistRemover, purge, purgeRemover, new address[](0)))
        );
    }

    function _predict(bytes32 salt, bytes32 initCodeHash) internal pure returns (address) {
        // Foundry's deterministic deployer (CREATE2_FACTORY in forge-std).
        address factory = 0x4e59b44847b379578588920cA78FbF26c0B4956C;
        return address(uint160(uint256(keccak256(abi.encodePacked(bytes1(0xff), factory, salt, initCodeHash)))));
    }

    /// @dev TimelockController init code with a single proposer, the
    ///      provided executors list, and `admin = deployerAddr` (matches
    ///      the new role-split deploy path).
    function _timelockInitCode(
        uint256 delay,
        address proposer,
        address[] memory executors,
        address deployerAddr
    ) internal pure returns (bytes memory) {
        address[] memory proposers = new address[](1);
        proposers[0] = proposer;
        return abi.encodePacked(
            type(TimelockController).creationCode,
            abi.encode(delay, proposers, executors, deployerAddr)
        );
    }

    function _proxyInitCode(
        address impl,
        address upgradeTimelock,
        address addTimelock,
        address whitelistRemover,
        address purgeTimelock,
        address purgeRemover,
        address[] memory initialWhitelisters
    ) internal pure returns (bytes memory) {
        bytes memory initCalldata = abi.encodeCall(
            ThatsRekt.initialize,
            (upgradeTimelock, addTimelock, whitelistRemover, purgeTimelock, purgeRemover, initialWhitelisters)
        );
        return abi.encodePacked(type(ERC1967Proxy).creationCode, abi.encode(impl, initCalldata));
    }

    /// @dev Asserts only `expected` holds PROPOSER_ROLE (within the set
    ///      we care about: not exhaustive, but covers the principals
    ///      this test suite uses).
    function _assertProposerOnly(TimelockController tlc, address expected) internal view {
        assertTrue(tlc.hasRole(tlc.PROPOSER_ROLE(), expected), "expected proposer missing");
    }

    /// @dev Asserts `expected` holds CANCELLER_ROLE.
    function _assertCancellerOnly(TimelockController tlc, address expected) internal view {
        assertTrue(tlc.hasRole(tlc.CANCELLER_ROLE(), expected), "expected canceller missing");
    }

    /// @dev Asserts that DEFAULT_ADMIN_ROLE has been renounced — only
    ///      the TLC itself holds admin (self-administered going forward).
    function _assertAdminLockedDown(TimelockController tlc, address deployerAddr) internal view {
        assertFalse(
            tlc.hasRole(tlc.DEFAULT_ADMIN_ROLE(), deployerAddr),
            "deployer still has DEFAULT_ADMIN_ROLE (renounce missed)"
        );
        // The TLC is its own admin — confirms the contract isn't trustless-bricked.
        assertTrue(
            tlc.hasRole(tlc.DEFAULT_ADMIN_ROLE(), address(tlc)),
            "TLC self-admin missing (would brick role rotation)"
        );
    }
}
