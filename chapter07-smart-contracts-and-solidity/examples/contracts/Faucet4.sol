//  Version of Solidity compiler this program was written for
pragma solidity ^0.5.6;

// Our first contract is a faucet!
contract Faucet {
  address payable owner;

  modifier onlyOwner {
    require(msg.sender == owner);
    _;
  }

  // Initialize Faucet contract: set owner
  constructor() public {
    owner = msg.sender;
  }

  // Contract destructor
  function destroy() public onlyOwner{
    selfdestruct(owner);
  }

  // Give out ether to anyone who asks
  function withdraw(uint amount) public {
    // Limit withdrawal amount
    require(amount <= 0.1 ether);

    // Send the amount to the address that requested it
    msg.sender.transfer(amount);
  }

  // Accept any incoming amount
  function () external payable { }
}