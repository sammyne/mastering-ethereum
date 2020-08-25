// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract HashForEther {
    function withdrawWinnings() external {
        // Winner if the last 8 hex characters of the address are 0
        require(uint32(msg.sender) == 0);
        _sendWinnings();
    }

    function _sendWinnings() {
        msg.sender.transfer(address(this).balance);
    }
}
