// SPDX-License-Identifier: ISC
pragma solidity ^0.7.0;

// library contract - calculates Fibonacci-like numbers
contract FibonacciLib {
    // initializing the standard Fibonacci sequence
    uint256 public start;
    uint256 public calculatedFibNumber;

    // modify the zeroth number in the sequence
    function setStart(uint256 _start) public {
        start = _start;
    }

    function setFibonacci(uint256 n) public {
        calculatedFibNumber = fibonacci(n);
    }

    function fibonacci(uint256 n) internal returns (uint256) {
        if (n == 0) return start;
        else if (n == 1) return start + 1;
        else return fibonacci(n - 1) + fibonacci(n - 2);
    }
}
