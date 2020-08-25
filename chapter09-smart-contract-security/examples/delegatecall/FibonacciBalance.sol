// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

contract FibonacciBalance {
    address public fibonacciLibrary;
    // the current Fibonacci number to withdraw
    uint256 public calculatedFibNumber;
    // the starting Fibonacci sequence number
    uint256 public start = 3;
    uint256 public withdrawalCounter;
    // the Fibonacci function selector
    bytes4 constant fibSig = bytes4(keccak256("setFibonacci(uint256)"));

    // constructor - loads the contract with ether
    constructor(address _fibonacciLibrary) payable {
        fibonacciLibrary = _fibonacciLibrary;
    }

    function withdraw() public {
        withdrawalCounter += 1;
        // calculate the Fibonacci number for the current withdrawal user-
        // this sets calculatedFibNumber
        (bool ok, ) = fibonacciLibrary.delegatecall(
            abi.encodeWithSelector(fibSig, withdrawalCounter)
        );
        require(ok);
        msg.sender.transfer(calculatedFibNumber * 1 ether);
    }

    // allow users to call Fibonacci library functions
    receive() external payable {
        (bool ok, ) = fibonacciLibrary.delegatecall(msg.data);
        require(ok);
    }
}
