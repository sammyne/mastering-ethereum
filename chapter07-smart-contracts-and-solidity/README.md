# Chapter 07. Smart Contracts and Solidity

- EOAs are controlled by users, often via software such as a wallet application that is external to the Ethereum platform
- Contract accounts are controlled by program code (also commonly referred to as "smart contracts") that is executed by the Ethereum Virtual Machine

## What Is a Smart Contract?

- **DEFINITION**: Immutable computer programs that run deterministically in the context of an Ethereum Virtual Machine as part of the Ethereum network protocol -- i.e., on the decentralized Ethereum world computer
  - **Computer programs**: Smart contracts are simply computer program
  - **Immutable**: Once deployed, the code of a smart contract cannot change
  - **Deterministic**: Always the same output given the same tx and blockchain state
  - **EVM context**: Smart contracts can access their own state, the context of the transaction that called them, and some information about the most recent blocks
  - **Decentralized world computer**: All instances of the EVM operate on the same initial state and produce the same final state, the system as a whole operates as a single "world computer"

## Life Cycle of a Smart Contract

- Life cycle
  - Written in a high-level language, such as Solidity
  - Compiled to the low-level bytecode that runs in the EVM
  - Deployed on the Ethereum platform using a special contract creation transaction
- **Contracts only run if they are called by a transaction**
- Transactions are **atomic**, regardless of how many contracts they call or what those contracts do when called. Transactions execute in their entirety, with any changes in the global state (contracts, accounts, etc.) recorded only if all execution terminates successfully
- Contract deletion
  - A contract's code cannot be changed but can be "**deleted**", removing the code and its internal state (storage) from its address, leaving a blank account
  - Any txs sent to the deleted contract account address will result in no code execution
  - **HOW**: execute the EVM opcode called `SELFDESTRUCT`
  - `SELFDESTRUCT` costs "negative gas," a gas refund, thereby incentivizing the release of network client resources from the deletion of stored state
  - The `SELFDESTRUCT` capability will only be available if the contract author programmed the smart contract to have that functionality

## Introduction to Ethereum High-Level Languages

- The EVM is a virtual machine that runs a special form of code called **EVM bytecode**
- **WHY NOT BYTECODE**
  - EVM bytecodes is rather unwieldy and very difficult for programmers to read and understand
  - Programs writting in a high-level language can be compiled down to bytecodes
- Programming languages can be classified into two broad programming paradigms

  - Declarative (a.k.a, functional), including Haskell and SQL
    - Makes it easier to understand how a program will behave: since it has no side effects, any part of a program can be understood in isolation
  - Imperative (a.k.a. procedural), including C++ and Java

    - More commonly used by programmers
    - It can be very difficult to write programs that execute exactly as expected
    - The ability of any part of the program to change the state of any other makes it difficult to reason about a program's execution and introduces many opportunities for bugs

    > The hybrid ones include Lisp, JavaScript and Python

- In smart contracts, bugs literally cost money, so declarative languages is preferred
- Supported high-level languages are

| Language | Description                                                                                                                                                                                                             |
| -------: | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|      LLL | A functional (declarative) programming language, with Lisp-like syntax                                                                                                                                                  |
|  Serpent | A procedural (imperative) programming language with a syntax similar to Python. Can also be used to write functional (declarative) code, though it is not entirely free of side effects                                 |
| Solidity | A procedural (imperative) programming language with a syntax similar to JavaScript, C++, or Java. **The most popular and frequently used language for Ethereum smart contracts**                                        |
|    Vyper | A more recently developed language, similar to Serpent and again with Python-like syntax. Intended to get closer to a pure-functional Python-like language than Serpent, but not to replace Serpent                     |
|   Bamboo | A newly developed language, influenced by Erlang, with explicit state transitions and without iterative flows (loops). Intended to reduce side effects and increase auditability. Very new and yet to be widely adopted |

## Building a Smart Contract with Solidity

- Solidity is now developed and maintained as an independent project on [GitHub](https://github.com/ethereum/solidity)
- The main "product" of the Solidity project is the Solidity compiler, `solc`, which converts programs written in the Solidity language to EVM bytecode
- The project also manages the important application binary interface (ABI) standard for Ethereum smart contracts
- Each version of the Solidity compiler corresponds to and compiles a specific version of the Solidity language.

### Selecting a Version of Solidity

- Solidity follows [semantic versioning](https://semver.org/), which specifies version numbers structured as three numbers separated by dots: `MAJOR.MINOR.PATCH`

  - The "major" number is incremented for major and backward-incompatible changes
  - The "minor" number is incremented as backward-compatible features are added in between major releases
  - The "patch" number is incremented for backward-compatible bug fixes

- Solidity programs can contain a pragma directive that specifies the minimum and maximum versions of Solidity that it is compatible with, and can be used to compile your contract

### Download and Install

- Detailed in [the Solidity documentation](https://solidity.readthedocs.io/en/v0.5.6/installing-solidity.html)

### Development Environment

- Desktop-based text editors
  - Emacs
  - Vim
  - Atom
- Web-based text editors
  - [Remix IDE](https://remix.ethereum.org/#optimize=false&version=soljson-v0.5.1+commit.c8a2cb62.js)
  - [EthFiddle](https://ethfiddle.com/)
- Simply save your program source code with a `.sol` extension and it will be recognized by the Solidity compiler as a Solidity program

### Writing a Simple Solidity Program

- As [Faucet.sol](examples/Faucet.sol)

### Compiling with the Solidity Compiler (solc)

- Use the `--bin` and `--optimize` arguments of solc to produce an optimized binary of our example contract

## The Ethereum Contract ABI

- In computer software, an **application binary interface** is an interface between two program modules; often, between the **operating system** and **user programs**
- An ABI defines how data structures and functions are accessed in **machine code**
- In Ethereum, the ABI is used to encode contract calls for the EVM and to read data out of transactions. The purpose of an ABI is to define the functions in the contract that can be invoked and describe how each function will accept arguments and return its result
- A contract's ABI is specified as a JSON array of
  - Function descriptions as `(type, name, inputs, outputs, constant, payable)`
  - Events as `(type, name, inputs, anonymous)`
- ABI produced by `solc` with `--abi` option
- All that is needed for an application to interact with a contract is
  - An ABI
  - The address where the contract has been deployed

### Selecting a Solidity Compiler and Language Version

- **Problem**: A contract written in a specific version of Solidity is given to different version of Solidity compilers
- **Solution**: Solidity offers a `compiler directive` known as a `version pragma` that instructs the compiler that the program expects a specific compiler (and lan‐ guage) version
- Pragma directives are not compiled into EVM bytecode. They are only used by the compiler to check compatibility. If missing, a warning will be reported
  - Tested with [Faucet.sol](examples/Faucet.sol) commenting out the `pragma` directive
  - Adding a version pragma is a best practice, as it avoids problems with mismatched compiler and language versions

## Programming with Solidity

### Data Types

|                           Type | Description                                                                                                                                                                   |
| -----------------------------: | :---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|               Boolean (`bool`) | `true`/`false`, with logical operators `!`(not), `&&`(add), `||`(or), `==`(equal), `!=`(not equal)                                                                            |
|         Integer (`int`,`uint`) | Signed (int) and unsigned (uint) integers, declared in increments of 8 bits from `int8` to `uint256`. `int`/`uint` aliases `int256`/`uint256`                                 |
| Fixed point (`fixed`/`ufixed`) | Fixed-point numbers, declared with `(u)fixedMxN` where `M` is the size in bits (increments of `8` up to `256`) and `N` is the number of decimals after the point (up to `18`) |
|            Address (`address`) | A 20-byte Ethereum address with many helpful member functions, the main ones being `balance` (returns the account balance) and `transfer`                                     |
|           Byte array (_fixed_) | Fixed-size arrays of bytes, declared with `bytes1` up to `bytes32`                                                                                                            |
|         Byte array (_dynamic_) | Variable-sized arrays of bytes, declared with `bytes` or `string`                                                                                                             |
|                           Enum | User-defined type for enumerating discrete values: `enum NAME {LABEL1, LABEL 2, ...}`                                                                                         |
|                         Arrays | An array of any type, either fixed or dynamic: `uint32[][5]` is a fixed-size array of five dynamic arrays of unsigned integers                                                |
|                         Struct | User-defined data containers for grouping variables: `struct NAME {TYPE1 VARIABLE1; TYPE2 VARIABLE2; ...}`                                                                    |
|                        Mapping | Hash lookup tables for `key => value` pairs: `mapping(KEY_TYPE => VALUE_TYPE) NAME`                                                                                           |

And various value literals as

|     Literal | Description                                                                                                                     |
| ----------: | :------------------------------------------------------------------------------------------------------------------------------ |
|  Time units | The units `seconds`, `minutes`, `hours`, and `days` can be used as suffixes, converting to multiples of the base unit `seconds` |
| Ether units | The units `wei`, `finney`, `szabo`, and `ether` can be used as suffixes, converting to multiples of the base unit `wei`         |

> Improve our code by using the unit multiplier ether, to express the value in `ether` instead of `wei`

### Predefined Global Variables and Functions

### Contract Definition

### Functions

### Contract Constructor and selfdestruct

### Adding a Constructor and selfdestruct to Our Faucet Example

### Function Modifiers

### Contract Inheritance

### Error Handling (assert, require, revert)

### Events

### Calling Other Contracts (send, call, callcode, delegatecall)

## Gas Considerations

### Avoid Dynamically Sized Arrays

### Avoid Calls to Other Contracts

### Estimating Gas Cost

## Conclusions
