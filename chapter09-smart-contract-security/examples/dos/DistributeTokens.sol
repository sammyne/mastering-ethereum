// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract DistributeTokens {
    address public owner; // gets set somewhere
    address[] investors; // array of investors
    uint256[] investorTokens; // the amount of tokens each investor gets

    // ... extra functionality, including transfertoken()

    function invest() public payable {
        investors.push(msg.sender);
        investorTokens.push(msg.value * 5); // 5 times the wei sent
    }

    function distribute() public {
        require(msg.sender == owner); // only owner
        for (uint256 i = 0; i < investors.length; i++) {
            // here transferToken(to,amount) transfers "amount" of
            // tokens to the address "to"
            transferToken(investors[i], investorTokens[i]);
        }
    }
}
