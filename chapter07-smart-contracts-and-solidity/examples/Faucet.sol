pragma solidity ^0.5.6;

// Example 7-1. Faucet.sol: A Solidity contract implementing a faucet

// Our first contract is a faucet!
contract Faucet {

  // Give out ether to anyone who asks
  function withdraw(uint amount) public {
    // improve our code by using the unit multiplier ether, 
    // to express the value in ether instead of wei
    //require(amount <= 100000000000000000000);
    require(amount <= 0.1 ether);

    // Send the amount to the address that requested it
    msg.sender.transfer(amount);
  }

  // Accept any incoming amount
  function () external payable { }
}