// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract Roulette {
    uint256 public pastBlockTime; // forces one bet per block

    constructor() payable {} // initially fund contract

    // fallback function used to make a bet
    receive() external payable {
        require(msg.value == 10 ether); // must send 10 ether to play
        require(block.timestamp != pastBlockTime); // only 1 transaction per block
        pastBlockTime = block.timestamp;
        if (block.timestamp % 15 == 0) {
            // winner
            msg.sender.transfer(address(this).balance);
        }
    }
}
