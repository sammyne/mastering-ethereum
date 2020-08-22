// SPDX-License-Identifier: ISC
// Version of Solidity compiler this program was written for
pragma solidity ^0.7.0;

contract calledContract {
    event callEvent(address sender, address origin, address from);

    function calledFunction() public {
        emit callEvent(msg.sender, tx.origin, address(this));
    }
}

library calledLibrary {
    event callEvent(address sender, address origin, address from);

    function calledFunction() public {
        emit callEvent(msg.sender, tx.origin, address(this));
    }
}

contract caller {
    function makeCalls(calledContract _calledContract) public {
        // Calling calledContract and calledLibrary directly
        _calledContract.calledFunction();
        calledLibrary.calledFunction();

        // Low-level calls using the address object for calledContract
        bytes memory methodSig = abi.encodeWithSignature("calledFunction()");

        (bool ok, bytes memory _) = address(_calledContract).call(methodSig);
        require(ok, "call failed");

        //require(address(_calledContract).delegatecall(methodSig));
        (ok, _) = address(_calledContract).delegatecall(methodSig);
        require(ok, "delegatecall failed");
    }
}
