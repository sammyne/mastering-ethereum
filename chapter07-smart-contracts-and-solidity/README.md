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

- Variables as
  - `block`
  - `msg`
  - `tx`
- Functions exposed as EVM opcodes

#### Transaction/message call context

- The `msg` object is the transaction call (EOA originated) or message call (contract originated) that launched this contract execution
- It contains a number of useful attributes

| Attribute | Description                                                                                                                                                      |
| --------: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|  `sender` | It represents the `address` that initiated this contract call, not necessarily the originating EOA that sent the transaction (contract address is also possible) |
|   `value` | The value of ether sent with this call (in `wei`)                                                                                                                |
|     `gas` | The amount of gas left in the gas supply of this execution environment. This was deprecated in Solidity **v0.4.21** and replaced by the `gasleft` function       |
|    `data` | The data payload of this call into our contract                                                                                                                  |
|     `sig` | The first four bytes of the data payload, which is the function selector (??)                                                                                    |

> Whenever a contract calls another contract, the values of all the attributes of `msg` change to reflect the new caller's information. The only exception to this is the `delegatecall` function

#### Transaction context

- Expressed as the `tx` object with information

|  Attribute | Description                                                                   |
| ---------: | :---------------------------------------------------------------------------- |
| `gasprice` | The gas price in the calling transaction                                      |
|   `origin` | The address of the originating EOA for this transaction. **WARNING: unsafe!** |

#### Block context

- Expressed as the `block` object containing information

|                Attribute | Description                                                                                                                                               |
| -----------------------: | :-------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `blockhash(blockNumber)` | The block hash of the specified block number, up to 256 blocks in the past. **Deprecated and replaced with the `blockhash` function in Solidity v0.4.22** |
|               `coinbase` | The **address** of the recipient of the current block's fees and block reward                                                                             |
|             `difficulty` | The difficulty (proof of work) of the current block                                                                                                       |
|               `gaslimit` | The maximum amount of gas that can be spent across all transactions included in the current block                                                         |
|                 `number` | The current block number (blockchain height)                                                                                                              |
|              `timestamp` | The timestamp placed in the current block by the miner (number of seconds since the Unix epoch)                                                           |

#### `address` object

|           Attribute | Description                                                                                                                                                                          |
| ------------------: | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|           `balance` | The balance of the address, in `wei`. For example, the current contract balance is `address(this).balance`                                                                           |
|  `transfer(amount)` | Transfers the amount (in `wei`) to this address, throwing an exception on any error                                                                                                  |
|      `send(amount)` | Similar to `transfer`, only instead of throwing an exception, it returns `false` on error. **WARNING: always check the return value of send**                                        |
|     `call(payload)` | Low-level `CALL` function -- can construct an arbitrary message call with a data payload. Returns `false` on error. **WARNING: unsafe**                                              |
| `callcode(payload)` | Low-level `CALLCODE` function, like `address(this).call(...)` but with this contract's code replaced with that of address. Returns `false` on error. **WARNING: advanced use only!** |
|    `delegatecall()` | Low-level `DELEGATECALL` function, like `callcode(...)` but with the full `msg` context seen by the current contract. Returns `false` on error. **WARNING: advanced use only!**      |

#### Built-in functions

|                                   Function | Description                                                                                       |
| -----------------------------------------: | :------------------------------------------------------------------------------------------------ |
|                          `addmod`,`mulmod` | For modulo addition and multiplication                                                            |
| `keccak256`, `sha256`, `sha3`, `ripemd160` | Functions to calculate hashes with various standard hash algorithms                               |
|                                `ecrecover` | Recovers the address used to sign a message from the signature                                    |
|         `selfdestrunct(recipient_address)` | Deletes the current contract, sending any remaining ether in the account to the recipient address |
|                                     `this` | The address of the currently executing contract account                                           |

### Contract Definition

- Signaled by the `contract` keyword
- 2 other similar objects are
  - `interface` is structured exactly like a `contract`, except none of the functions are defined, they are only declared. This type of declaration is often called a `stub`
  - `library` is meant to be deployed only once and used by other contracts, using the `delegatecall` method

### Functions

- Syntax of declaration goes as

  ```solidity
  function FunctionName([parameters]) {public|private|internal|external} [pure|constant|view|payable] [modifiers] [returns (return types)]
  ```

  - `FunctionName`
    - Defines the name of the function, which is used to call the function in a transaction (from an EOA), from another contract, or even from within the same contract.
    - Especially, the so-called `fallback` function is defined without a name, which is called when no other function is called. The `fallback` function cannot have any arguments or return anything
  - `parameters`
    - The arguments that must be passed to the function, with their **names** and **types**
  - `public`/`private`/`internal`/`external` specify the function's visibility
    - `public` is the default; such functions can be called by other contracts or EOA transactions, or from within the contract
    - `external` is like `public`, except the decorated functions cannot be called from within the contract unless explicitly prefixed with the keyword `this`
    - `internal` make functions only accessible from within the contract -- they cannot be called by another contract or EOA transaction. They can be called by derived contracts (those that inherit this one)
    - `private` is like `internal`, except the decorated functions cannot be called by derived contracts
  - `pure`/`constant`/`view`/`payable` affects behaviors of the functions
    - `constant`/`view` marks a function promising not to modify any state, where `constant` will be deprecated in a future release
    - `pure` marks functions neither reading nor writing any variables in storage
    - `payable` marks the only functions for accepting incoming payments. 2 exceptions are coinbase payments and `SELFDESTRUCT` inheritance will be paid even if the fallback function is not declared as `payable`

### Contract Constructor and `selfdestruct`

- When a contract is created, it also runs the `constructor` function if one exists, to initialize the state of the contract
- The constructor function is optional
- Constructors can be specified in two ways

  - The constructor is a function whose name matches the name of the contract

    ```
    contract MEContract {
      function MEContract() {
        // This is the constructor
      }
    }
    ```

  - Solidity **v0.4.22** introduces a `constructor` keyword that operates like a constructor function but does not have a name

    ```
    pragma ^0.4.22

    contract MEContract {
      constructor () {
        // This is the constructor
      }
    }
    ```

- To summarize, a contract's life cycle
  - Starts with a creation transaction from an EOA or contract account
  - If there is a constructor, it is executed as part of contract creation, to initialize the state of the contract as it is being created, and is then discarded.
  - Finally, contract can be destructed
- Contracts are destroyed by a special EVM opcode called `SELFDESTRUCT` exposed as a high-level built-in function as

  ```
  // recipient is the address to receive any remaining ether balance
  selfdestruct(address recipient)
  ```

- Only the explicitly declared `selfdestruct` command can delete a contract

### Adding a `Constructor` and `selfdestruct` to Our Faucet Example

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
