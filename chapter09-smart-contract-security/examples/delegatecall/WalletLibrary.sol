// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract WalletLibrary is WalletEvents {
    //...

    // throw unless the contract is not yet initialized.
    modifier only_uninitialized {
        if (m_numOwners > 0) throw;
        _;
    }

    // constructor - just pass on the owner array to multiowned and
    // the limit to daylimit
    function initWallet(
        address[] _owners,
        uint256 _required,
        uint256 _daylimit
    ) only_uninitialized {
        initDaylimit(_daylimit);
        initMultiowned(_owners, _required);
    }

    // kills the contract sending everything to `_to`.
    function kill(address _to) external onlymanyowners(sha3(msg.data)) {
        selfdestruct(_to);
    }

    //...
}
