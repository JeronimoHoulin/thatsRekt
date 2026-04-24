// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Test} from "forge-std/Test.sol";
import {ThatsRekt} from "../src/ThatsRekt.sol";

contract ThatsRektTest is Test {
    ThatsRekt public rekt;

    address public signer1 = makeAddr("signer1");
    address public signer2 = makeAddr("signer2");
    address public alice   = makeAddr("alice");
    address public bob     = makeAddr("bob");
    address public charlie = makeAddr("charlie");
    address public rando   = makeAddr("rando");

    event WhitelistUpdated(address indexed account, bool status);
    event AddedRekt(address indexed account);
    event RemovedRekt(address indexed account);
    event RemovalProposed(uint256 indexed id, address indexed proposer, address[] targets);
    event RemovalExecuted(uint256 indexed id, address indexed executor);

    function setUp() public {
        rekt = new ThatsRekt();

        // Seed two test signers via the hardcoded JERRYTHEKID address
        vm.startPrank(rekt.JERRYTHEKID());
        rekt.addWhitelisted(signer1);
        rekt.addWhitelisted(signer2);
        vm.stopPrank();
    }

    /*//////////////////////////////////////////////////////////////
                            WHITELIST TESTS
    //////////////////////////////////////////////////////////////*/

    function test_InitialWhitelist() public view {
        assertTrue(rekt.isWhitelisted(rekt.JERRYTHEKID()));
        assertTrue(rekt.isWhitelisted(rekt.BAUTI()));
    }

    function test_WhitelistedCanAddToWhitelist() public {
        vm.prank(signer1);
        vm.expectEmit(true, false, false, true);
        emit WhitelistUpdated(rando, true);
        rekt.addWhitelisted(rando);

        assertTrue(rekt.isWhitelisted(rando));
    }

    function test_WhitelistedCanRemoveFromWhitelist() public {
        vm.prank(signer1);
        rekt.removeWhitelisted(signer2);

        assertFalse(rekt.isWhitelisted(signer2));
    }

    function test_RevertIf_NonWhitelistedAddsToWhitelist() public {
        vm.prank(rando);
        vm.expectRevert(ThatsRekt.NotWhitelisted.selector);
        rekt.addWhitelisted(rando);
    }

    function test_RevertIf_NonWhitelistedRemovesFromWhitelist() public {
        vm.prank(rando);
        vm.expectRevert(ThatsRekt.NotWhitelisted.selector);
        rekt.removeWhitelisted(signer1);
    }

    /*//////////////////////////////////////////////////////////////
                             addRekt TESTS
    //////////////////////////////////////////////////////////////*/

    function test_AddRekt_Single() public {
        vm.expectEmit(true, false, false, false);
        emit AddedRekt(alice);

        vm.prank(signer1);
        rekt.addRekt(_arr(alice));

        address[] memory list = rekt.isRekt();
        assertEq(list.length, 1);
        assertEq(list[0], alice);
    }

    function test_AddRekt_Multiple() public {
        vm.prank(signer1);
        rekt.addRekt(_arr(alice, bob, charlie));

        assertEq(rekt.isRekt().length, 3);
    }

    function test_AddRekt_SkipsDuplicateWithinCall() public {
        address[] memory targets = new address[](2);
        targets[0] = alice;
        targets[1] = alice;

        vm.prank(signer1);
        rekt.addRekt(targets);

        assertEq(rekt.isRekt().length, 1);
    }

    function test_AddRekt_SkipsAlreadyRekt() public {
        vm.prank(signer1);
        rekt.addRekt(_arr(alice));

        vm.prank(signer2);
        rekt.addRekt(_arr(alice));

        assertEq(rekt.isRekt().length, 1);
    }

    function test_AddRekt_WhitelistedAddressCanBeRekt() public {
        vm.prank(signer1);
        rekt.addRekt(_arr(signer2));

        assertEq(rekt.isRekt().length, 1);
    }

    function test_RevertIf_NonWhitelistedCallsAddRekt() public {
        vm.prank(rando);
        vm.expectRevert(ThatsRekt.NotWhitelisted.selector);
        rekt.addRekt(_arr(alice));
    }

    /*//////////////////////////////////////////////////////////////
                          proposeRemoval TESTS
    //////////////////////////////////////////////////////////////*/

    function test_ProposeRemoval_ReturnsId() public {
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        assertEq(id, 0);
        assertEq(rekt.proposalCount(), 1);
    }

    function test_ProposeRemoval_StoresProposal() public {
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice, bob));

        (address proposer, bool executed, address[] memory targets) = rekt.getProposal(id);
        assertEq(proposer, signer1);
        assertFalse(executed);
        assertEq(targets.length, 2);
        assertEq(targets[0], alice);
        assertEq(targets[1], bob);
    }

    function test_ProposeRemoval_IdsIncrement() public {
        vm.startPrank(signer1);
        assertEq(rekt.proposeRemoval(_arr(alice)), 0);
        assertEq(rekt.proposeRemoval(_arr(bob)),   1);
        assertEq(rekt.proposeRemoval(_arr(charlie)), 2);
        vm.stopPrank();
    }

    function test_ProposeRemoval_EmitsEvent() public {
        address[] memory targets = _arr(alice);
        vm.expectEmit(true, true, false, true);
        emit RemovalProposed(0, signer1, targets);

        vm.prank(signer1);
        rekt.proposeRemoval(targets);
    }

    function test_RevertIf_NonWhitelistedProposes() public {
        vm.prank(rando);
        vm.expectRevert(ThatsRekt.NotWhitelisted.selector);
        rekt.proposeRemoval(_arr(alice));
    }

    /*//////////////////////////////////////////////////////////////
                          executeRemoval TESTS
    //////////////////////////////////////////////////////////////*/

    function test_ExecuteRemoval_HappyPath() public {
        vm.prank(signer1);
        rekt.addRekt(_arr(alice, bob));

        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        vm.expectEmit(true, false, false, false);
        emit RemovedRekt(alice);

        vm.prank(signer2);
        rekt.executeRemoval(id);

        address[] memory list = rekt.isRekt();
        assertEq(list.length, 1);
        assertEq(list[0], bob);
    }

    function test_ExecuteRemoval_MarksProposalExecuted() public {
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        vm.prank(signer2);
        rekt.executeRemoval(id);

        (, bool executed,) = rekt.getProposal(id);
        assertTrue(executed);
    }

    function test_ExecuteRemoval_EmitsEvent() public {
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        vm.expectEmit(true, true, false, false);
        emit RemovalExecuted(id, signer2);

        vm.prank(signer2);
        rekt.executeRemoval(id);
    }

    function test_ExecuteRemoval_SkipsAddressNotInList() public {
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(makeAddr("notInList")));

        vm.prank(signer2);
        rekt.executeRemoval(id); // should not revert

        (, bool executed,) = rekt.getProposal(id);
        assertTrue(executed);
    }

    function test_ExecuteRemoval_RemoveAll() public {
        vm.prank(signer1);
        rekt.addRekt(_arr(alice, bob, charlie));

        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice, bob, charlie));

        vm.prank(signer2);
        rekt.executeRemoval(id);

        assertEq(rekt.isRekt().length, 0);
    }

    function test_ExecuteRemoval_SwapAndPopPreservesSet() public {
        vm.prank(signer1);
        rekt.addRekt(_arr(alice, bob, charlie));

        // Remove alice (slot 0) — charlie swaps in
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        vm.prank(signer2);
        rekt.executeRemoval(id);

        address[] memory list = rekt.isRekt();
        assertEq(list.length, 2);
        bool hasBob;
        bool hasCharlie;
        for (uint256 i; i < list.length; ++i) {
            if (list[i] == bob)     hasBob     = true;
            if (list[i] == charlie) hasCharlie = true;
        }
        assertTrue(hasBob);
        assertTrue(hasCharlie);
    }

    function test_ExecuteRemoval_ThirdWhitelistedCanExecute() public {
        address third = makeAddr("third");
        vm.prank(signer1);
        rekt.addWhitelisted(third);

        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        vm.prank(third); // not the proposer, but whitelisted
        rekt.executeRemoval(id);
    }

    /*//////////////////////////////////////////////////////////////
                     executeRemoval REVERT TESTS
    //////////////////////////////////////////////////////////////*/

    function test_RevertIf_ProposerExecutesOwnProposal() public {
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        vm.prank(signer1);
        vm.expectRevert(ThatsRekt.CannotSelfExecute.selector);
        rekt.executeRemoval(id);
    }

    function test_RevertIf_AlreadyExecuted() public {
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        vm.prank(signer2);
        rekt.executeRemoval(id);

        vm.prank(signer2);
        vm.expectRevert(ThatsRekt.AlreadyExecuted.selector);
        rekt.executeRemoval(id);
    }

    function test_RevertIf_ProposalNotFound() public {
        vm.prank(signer1);
        vm.expectRevert(ThatsRekt.ProposalNotFound.selector);
        rekt.executeRemoval(999);
    }

    function test_RevertIf_NonWhitelistedExecutes() public {
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        vm.prank(rando);
        vm.expectRevert(ThatsRekt.NotWhitelisted.selector);
        rekt.executeRemoval(id);
    }

    function test_RevertIf_DewhitelistedProposerBlocksNothing() public {
        // Proposer getting de-whitelisted after proposing should NOT block execution
        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(_arr(alice));

        vm.prank(signer2);
        rekt.removeWhitelisted(signer1);

        // signer2 can still execute — whitelist check is on the executor, not proposer
        vm.prank(signer2);
        rekt.executeRemoval(id);
    }

    /*//////////////////////////////////////////////////////////////
                              FUZZ TESTS
    //////////////////////////////////////////////////////////////*/

    function testFuzz_AddRekt_NoDuplicates(address[] calldata targets) public {
        vm.assume(targets.length < 64);

        vm.prank(signer1);
        rekt.addRekt(targets);

        address[] memory list = rekt.isRekt();
        for (uint256 i; i < list.length; ++i) {
            for (uint256 j = i + 1; j < list.length; ++j) {
                assertNotEq(list[i], list[j], "duplicate in rekt list");
            }
        }
    }

    function testFuzz_RemovalPreservesSetIntegrity(uint8 seed) public {
        address a = makeAddr("a");
        address b = makeAddr("b");
        address c = makeAddr("c");
        address d = makeAddr("d");

        vm.prank(signer1);
        rekt.addRekt(_arr(a, b, c, d));

        address[] memory toRemove = new address[](1);
        address[4] memory all = [a, b, c, d];
        toRemove[0] = all[seed % 4];

        vm.prank(signer1);
        uint256 id = rekt.proposeRemoval(toRemove);

        vm.prank(signer2);
        rekt.executeRemoval(id);

        address[] memory list = rekt.isRekt();
        assertEq(list.length, 3);
        for (uint256 i; i < list.length; ++i) {
            assertNotEq(list[i], toRemove[0], "removed address still in list");
        }
        for (uint256 i; i < list.length; ++i) {
            for (uint256 j = i + 1; j < list.length; ++j) {
                assertNotEq(list[i], list[j], "duplicate after removal");
            }
        }
    }

    /*//////////////////////////////////////////////////////////////
                              TEST HELPERS
    //////////////////////////////////////////////////////////////*/

    function _arr(address a) internal pure returns (address[] memory r) {
        r = new address[](1);
        r[0] = a;
    }

    function _arr(address a, address b) internal pure returns (address[] memory r) {
        r = new address[](2);
        r[0] = a; r[1] = b;
    }

    function _arr(address a, address b, address c) internal pure returns (address[] memory r) {
        r = new address[](3);
        r[0] = a; r[1] = b; r[2] = c;
    }

    function _arr(address a, address b, address c, address d) internal pure returns (address[] memory r) {
        r = new address[](4);
        r[0] = a; r[1] = b; r[2] = c; r[3] = d;
    }
}
