// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract Print {
    event Print(string text);

    function rot13Encrypt(string text) public {
        emit Print(text);
    }
}
