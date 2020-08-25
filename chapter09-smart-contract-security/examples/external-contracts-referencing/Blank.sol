// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract Blank {
    event Print(string text);

    fallback() external {
        emit Print("Here");
        // put malicious code here and it will run
    }
}
