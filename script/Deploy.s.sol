// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {Script, console2} from "forge-std/Script.sol";
import {ThatsRekt} from "../src/ThatsRekt.sol";

/// @notice Cross-chain deterministic deploy via the singleton CREATE2 factory.
///         Uses CREATE2_FACTORY (0x4e59...) inherited from forge-std/Base.sol.
///         Reverts if the GOVERNANCE constant is still the dev placeholder.
contract Deploy is Script {
    address public constant DEV_PLACEHOLDER = 0x000000000000000000000000000000000000ABcD;
    bytes32 public constant SALT            = keccak256("thatsRekt.v1");

    function run() external {
        // Read GOVERNANCE constant via a simulation-only sentinel instance.
        // This is local to the Foundry runner — no real-chain effect.
        ThatsRekt sentinel = new ThatsRekt();
        address gov = sentinel.GOVERNANCE();

        require(gov != DEV_PLACEHOLDER, "GOVERNANCE is still the dev placeholder");
        require(gov != address(0),       "GOVERNANCE is zero");
        require(gov.code.length > 0,     "GOVERNANCE has no code (must be a Safe / contract)");

        bytes memory initCode = type(ThatsRekt).creationCode;
        bytes32 initCodeHash = keccak256(initCode);
        address predicted = computeCreate2Address(SALT, initCodeHash, CREATE2_FACTORY);

        console2.log("Predicted address:", predicted);
        console2.log("Governance owner:",  gov);

        if (predicted.code.length > 0) {
            console2.log("Already deployed at predicted address - skipping.");
            return;
        }

        vm.startBroadcast();
        bytes memory payload = abi.encodePacked(SALT, initCode);
        (bool ok, bytes memory ret) = CREATE2_FACTORY.call(payload);
        require(ok, "CREATE2 deploy failed");
        address deployed = address(uint160(bytes20(ret)));
        require(deployed == predicted, "deployed address != predicted");
        vm.stopBroadcast();

        console2.log("Deployed at:", deployed);
    }
}
