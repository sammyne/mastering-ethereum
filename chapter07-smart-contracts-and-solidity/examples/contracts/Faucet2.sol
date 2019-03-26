//pragma solidity ^0.5.6;

// Example 7-2. Faucet2.sol: Adding the version pragma to Faucet

// Our first contract is a faucet!
contract Faucet {

  // Give out ether to anyone who asks
  function withdraw(uint amount) public {
    // Limit withdrawal amount
    require(amount <= 100000000000000000000);

    // Send the amount to the address that requested it
    msg.sender.transfer(amount);
  }

  // Accept any incoming amount
  function () external payable { }
}