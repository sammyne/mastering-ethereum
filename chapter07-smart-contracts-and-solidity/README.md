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

### Selecting a Version of Solidity

### Download and Install

### Development Environment

### Writing a Simple Solidity Program

### Compiling with the Solidity Compiler (solc)

## The Ethereum Contract ABI

### Selecting a Solidity Compiler and Language Version

## Programming with Solidity

### Data Types

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
