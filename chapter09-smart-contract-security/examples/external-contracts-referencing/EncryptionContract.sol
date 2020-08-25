// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

import "Rot13Encryption.sol";

// encrypt your top-secret info
contract EncryptionContract {
    // library for encryption
    Rot13Encryption encryptionLibrary;

    // constructor - initialize the library
    constructor(Rot13Encryption _encryptionLibrary) {
        encryptionLibrary = _encryptionLibrary;
    }

    function encryptPrivateData(string memory privateInfo) public {
        // potentially do some operations here
        encryptionLibrary.rot13Encrypt(privateInfo);
     }
 }