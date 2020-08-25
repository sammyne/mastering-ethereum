// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

// encryption contract
contract Rot13Encryption {
    event Result(string convertedString);

    // rot13-encrypt a string
    function rot13Encrypt(string memory text) public {
        uint256 length = bytes(text).length;
        for (uint256 i = 0; i < length; i++) {
            bytes1 char = bytes(text)[i];
            // inline assembly to modify the string
            assembly {
                // get the first byte
                char := byte(0, char)
                // if the character is in [n,z], i.e. wrapping
                if and(gt(char, 0x6D), lt(char, 0x7B)) {
                    // the difference between character <char> and 'z' // subtract from the ASCII number 'a',
                    char := sub(0x60, sub(0x7A, char))
                }
                if iszero(eq(char, 0x20)) {
                    // add 13 to char // ignore spaces
                    mstore8(add(add(text, 0x20), mul(i, 1)), add(char, 13))
                }
            }
        }
        emit Result(text);
    }

    // rot13-decrypt a string
    function rot13Decrypt(string memory text) public {
        uint256 length = bytes(text).length;
        for (uint256 i = 0; i < length; i++) {
            bytes1 char = bytes(text)[i];
            assembly {
                char := byte(0, char)
                if and(gt(char, 0x60), lt(char, 0x6E)) {
                    char := add(0x7B, sub(char, 0x61))
                }
                if iszero(eq(char, 0x20)) {
                    mstore8(add(add(text, 0x20), mul(i, 1)), sub(char, 13))
                }
            }
        }
        emit Result(text);
    }
}
