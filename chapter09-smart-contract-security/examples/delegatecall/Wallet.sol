pragma solidity ^0.5.6;

contract Wallet is WalletEvents {

  //...

  // METHODS

  // gets called when no other function matches
  function() external payable {
    // just being sent some cash?
    if (msg.value > 0)
      Deposit(msg.sender, msg.value);
    else if (msg.data.length > 0)
      _walletLibrary.delegatecall(msg.data);
  }

  //...

  // FIELDS
  address constant _walletLibrary =
    0xcafecafecafecafecafecafecafecafecafecafe;
}