pragma solidity ^0.5.6;

import "./node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";

// A faucet for ERC20 token MET
contract METFaucet { 
  ERC20 public METoken;
  address public METOwner;

  // METFaucet constructor, provide the address of METoken contract and
  // the owner address we will be approved to transferFrom
  constructor(address _METoken, address _METOwner) public {
    // Initialize the METoken from the address provided
    METoken = ERC20(_METoken);
    METOwner = _METOwner;
  }

  function withdraw(uint amount) public {
    // Limit withdrawal amount to 10 MET
    require(amount <= 1000);

    // Use the transferFrom function of METoken
    METoken.transferFrom(METOwner, msg.sender, amount);
  }

  // REJECT any incoming ether
  function () external payable { revert(); }
}