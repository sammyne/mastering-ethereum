// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract Phishable {
    address public owner;

    constructor (address _owner) {
        owner = _owner;
    }

    receive () external payable {} // collect ether

    function withdrawAll(address _recipient) public {
        require(tx.origin == owner);
        payable(_recipient).transfer(address(this).balance);
    }
}