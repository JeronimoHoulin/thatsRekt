// SPDX-License-Identifier: MIT
pragma solidity 0.8.25;

import {Script, console2} from "forge-std/Script.sol";
import {ThatsRekt} from "../src/ThatsRekt.sol";

/// @notice Cross-chain deterministic deploy via the singleton CREATE2 factory
///         (CREATE2_FACTORY = 0x4e59... inherited from forge-std/Base.sol).
///
///         The initial owner (typically the governance Safe multisig) is
///         passed as a constructor argument, read from the GOVERNANCE_OWNER
///         environment variable. The same value MUST be used on every chain
///         for cross-chain address parity.
///
///         Once deployed, the owner can rotate via Ownable2Step's two-step
///         transferOwnership / acceptOwnership flow — independent per chain.
contract Deploy is Script {
    bytes32 public constant SALT = keccak256("thatsRekt.v1");

    function run() external {
        address initialOwner = vm.envAddress("GOVERNANCE_OWNER");
        require(initialOwner != address(0), "GOVERNANCE_OWNER env var is zero");
        require(initialOwner.code.length > 0, "GOVERNANCE_OWNER has no code (must be a Safe / contract)");

        // Constructor arg gets encoded into the init code, so the same owner
        // value on every chain produces the same init code -> same CREATE2 address.
        bytes memory initCode = abi.encodePacked(
            type(ThatsRekt).creationCode,
            abi.encode(initialOwner)
        );
        bytes32 initCodeHash = keccak256(initCode);
        address predicted = computeCreate2Address(SALT, initCodeHash, CREATE2_FACTORY);

        console2.log("Predicted address:", predicted);
        console2.log("Initial owner:    ", initialOwner);

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
