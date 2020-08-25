// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract Lotto {
    bool public payedOut = false;
    address public winner;
    uint256 public winAmount;

    // ... extra functionality here

    function sendToWinner() public {
        require(!payedOut);
        payable(winner).send(winAmount);
        payedOut = true;
    }

    function withdrawLeftOver() public {
        require(payedOut);
        msg.sender.send(address(this).balance);
    }
}
