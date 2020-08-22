// SPDX-License-Identifier: ISC
//pragma solidity ^0.7.0;

// Example 7-2. Faucet02.sol: Adding the version pragma to Faucet

// Our first contract is a faucet!
contract Faucet {
    // Give out ether to anyone who asks
    function withdraw(uint256 amount) public {
        // Limit withdrawal amount
        require(amount <= 100000000000000000000);

        // Send the amount to the address that requested it
        msg.sender.transfer(amount);
    }

    // Accept any incoming amount
    receive() external payable {}
}
