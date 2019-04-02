pragma solidity ^0.5.6;


import "./node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20.sol";
import "./node_modules/openzeppelin-solidity/contracts/token/ERC20/ERC20Detailed.sol";

contract METoken is ERC20, ERC20Detailed {
  uint constant _initialSupply = 2_100_000_000;

  constructor() public ERC20Detailed("Mastering Ethereum Token", "MET", 2) {
    _mint(msg.sender, _initialSupply);
  }
}