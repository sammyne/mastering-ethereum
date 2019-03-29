pragma solidity ^0.5.6; 

contract owned {
  address payable owner;

  // Contract constructor: set owner
  constructor() public {
    owner = msg.sender;
  }

  // Access control modifier
  modifier onlyOwner {
    require(msg.sender == owner, "Only the contract owner can call this function");
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
  event Withdrawal(address indexed to, uint amount);
  event Deposit(address indexed from, uint amount);

  // Give out ether to anyone who asks
  function withdraw(uint amount) public {
    // Limit withdrawal amount
    require(amount <= 0.1 ether, "Insufficient balance in faucet for withdrawal request");
    // Send the amount to the address that requested it
    msg.sender.transfer(amount);
    emit Withdrawal(msg.sender, amount);
  }

  // Accept any incoming amount
  function() external payable {
    emit Deposit(msg.sender, msg.value);
  }
}