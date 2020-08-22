// SPDX-License-Identifier: ISC
// Version of Solidity compiler this program was written for
pragma solidity ^0.7.0;

contract owned {
    address payable owner;

    // Contract constructor: set owner
    constructor() {
        owner = msg.sender;
    }

    // Access control modifier
    modifier onlyOwner {
        require(
            msg.sender == owner,
            "Only the contract owner can call this function"
        );
        _;
    }
}

contract mortal is owned {
    // Contract destructor
    function destroy() public onlyOwner {
        selfdestruct(owner);
    }
}

contract Faucet is mortal {
    event Withdrawal(address indexed to, uint256 amount);
    event Deposit(address indexed from, uint256 amount);

    // Give out ether to anyone who asks
    function withdraw(uint256 amount) public {
        // Limit withdrawal amount
        require(
            amount <= 0.1 ether,
            "Insufficient balance in faucet for withdrawal request"
        );
        // Send the amount to the address that requested it
        msg.sender.transfer(amount);
        emit Withdrawal(msg.sender, amount);
    }

    // Accept any incoming amount
    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }
}
