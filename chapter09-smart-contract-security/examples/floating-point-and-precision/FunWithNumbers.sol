// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract FunWithNumbers {
    uint256 public constant tokensPerEth = 10;
    uint256 public constant weiPerEth = 1e18;
    mapping(address => uint256) public balances;

    function buyTokens() public payable {
        // convert wei to eth, then multiply by token rate
        uint256 tokens = (msg.value / weiPerEth) * tokensPerEth;
        balances[msg.sender] += tokens;
    }

    function sellTokens(uint256 tokens) public {
        require(balances[msg.sender] >= tokens);
        uint256 eth = tokens / tokensPerEth;
        balances[msg.sender] -= tokens;
        msg.sender.transfer(eth * weiPerEth);
    }
}
