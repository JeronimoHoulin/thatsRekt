// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Script, console} from "forge-std/Script.sol";
import {ThatsRekt} from "../src/ThatsRekt.sol";

/// @notice Deploys ThatsRekt. No constructor arguments — the two whitelisted
///         addresses (JERRYTHEKID & BAUTI) are hardcoded in the contract.
///
///         forge script script/Deploy.s.sol \
///           --rpc-url $RPC_URL \
///           --broadcast \
///           --verify \
///           -vvvv
contract Deploy is Script {
    function run() external returns (ThatsRekt instance) {
        vm.startBroadcast();
        instance = new ThatsRekt();
        vm.stopBroadcast();

        console.log("ThatsRekt  :", address(instance));
        console.log("JERRYTHEKID:", instance.JERRYTHEKID());
        console.log("BAUTI      :", instance.BAUTI());
    }
}
