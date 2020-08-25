// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract WalletLibrary is WalletEvents {
    //...

    // METHODS

    //...

    // constructor is given number of sigs required to do protected
    // "onlymanyowners" transactionsas well as the selection of addresses
    // capable of confirming them
    function initMultiowned(address[] _owners, uint256 _required) {
        m_numOwners = _owners.length + 1;
        m_owners[1] = uint256(msg.sender);
        m_ownerIndex[uint256(msg.sender)] = 1;
        for (uint256 i = 0; i < _owners.length; ++i) {
            m_owners[2 + i] = uint256(_owners[i]);
            m_ownerIndex[uint256(_owners[i])] = 2 + i;
        }
        m_required = _required;
    }

    //...

    // constructor - just pass on the owner array to multiowned and
    // the limit to daylimit
    function initWallet(
        address[] _owners,
        uint256 _required,
        uint256 _daylimit
    ) {
        initDaylimit(_daylimit);
        initMultiowned(_owners, _required);
    }
}
