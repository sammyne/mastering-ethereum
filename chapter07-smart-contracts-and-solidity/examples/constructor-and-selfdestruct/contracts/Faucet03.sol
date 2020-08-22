// SPDX-License-Identifier: ISC
//  Version of Solidity compiler this program was written for
pragma solidity ^0.7.0;

// Our first contract is a faucet!
contract Faucet {
    address payable owner;

    // Initialize Faucet contract: set owner
    constructor() {
        owner = msg.sender;
    }

    // Contract destructor
    function destroy() public {
        require(msg.sender == owner);
        selfdestruct(owner);
    }

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
