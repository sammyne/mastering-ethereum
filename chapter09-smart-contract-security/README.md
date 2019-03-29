# Chapter 09. Smart Contract Security

- In the field of smart contract programming, mistakes are costly and easily exploited
- The execution of smart contracts isn't always expected
- All smart contracts are public, so any vulnerability can be exploited, and losses are almost always impossible to recover

## Security Best Practices

- **Defensive programming** is a style of programming that is particularly well suited to smart contracts
- Best practices
  - Minimalism/simplicity
    - The simpler the code, and the less it does, the lower the chances are of a bug or unforeseen effect occurring
  - Code reuse
    - Follow DRY
    - Code that has been extensively used and tested is likely more secure than any new code you write
    - Beware of "Not Invented Here" syndrome
  - Code quality
    - Every bug can lead to monetary loss
    - Once you "launch" your code, there's little you can do to fix any problems
  - Readability/auditability
    - The easier it is to read, the easier it is to audit
    - Code should be well documented and easy to read, following the style and naming conventions that are part of the Ethereum community
  - Test coverage
    - Test all arguments to make sure they are within expected ranges and properly formatted before allowing execution of your code to continue

## Security Risks and Antipatterns

- **WHY** care: To detect and avoid the programming patterns that leave your contracts exposed to these risks

## Reentrancy

- **WHAT**: External calls originated from contracts can be hijacked by attackers, who can force the contracts to execute further code (through a fallback function), including calls back into themselves

### The Vulnerability

- **WHEN**: Contracts pay to unknown addresses which contains malicious codes in the fallback functions
- **HOW**: The external malicious contract calls a function on the vulnerable contract and the path of code execution "reenters" it
- Example as [EtherStore.sol](examples/reentrancy/EtherStore.sol) and [Attack.sol](examples/reentrancy/Attack.sol)
  - **HOW**: The external calling of `Attack` contract would drain the `EtherStore` until the store has balance no more than 1 ether
  - **WHY**: During the external calling, execution would reenter `EtherStore` without updating the `sender`'s amount, enabling him to act like the initial call

### Preventative Techniques

- Use the built-in `transfer` function which limit external calls to 2300 gas, rendering them unable to call other contracts
- Employ the **checks-effects-interactions** pattern (TODO: add the link)
  - For any code that performs external calls to unknown addresses to be the last opera‐ tion in a localized function or piece of code execution
- Use mutex to prevent reentrances
- A fixed `EtherStore` with all above applied go as [EtherStoreOK.sol](examples/reentrancy/EtherStoreOK.sol)

### Real-World Example: The DAO

- Consequence: \$150 million loss and a hark fork creating ETC
- TODO: link

## Arithmetic Over/Underflows

- Related links (TODO: links)
  - How to Secure Your Smart Contracts
  - Ethereum Smart Contract Best Practices
  - Ethereum, Solidity and integer overflows: programming blockchains like 1970

### The Vulnerability

- An over/underflow occurs when an operation is performed that requires a fixed-size variable to store a number (or piece of data) that is outside the range of the variable's data type
- Examples
  - [TimeLock.sol](examples/overflow-underflow/TimeLock.sol)
    - TODO: Fire an issue reporting the missing balance
    - **HOW**: Call the `increaseLockTime(2^256 - userLockTime)` to reset `lockTime[msg.sender]` as `0`, and then withdraw all reward
  - [Token.sol](examples/overflow-underflow/Token.sol)
    - **HOW**: Bypass the Line 13 with `_value>balances[msg.sender]` to trigger underflow thus stealing free tokens

### Preventative Techniques

- Use or build mathematical libraries that replace the standard math operators addition, subtraction, and multiplication (division is excluded as it does not cause over/underflows and the EVM reverts on division by 0)
- Recommendation: OpenZeppelin (TODO: link)
  - The SafeMath (TODO: link) to tackle over/underflows as [OverflowFreeTimeLock.sol](examples/overflow-underflow/OverflowFreeTimeLock.sol)

### Real-World Examples: PoWHC and Batch Transfer Overflow (CVE-2018–10299)

- Proof of Weak Hands Coin (PoWHC), originally devised as a joke of sorts, was a Ponzi scheme written by an internet collective
- PoWHC suffers from underflow as explained by Eric Banisadr (TODO: link)
- PeckShield's account (TODO: link) suffers from overflow

## Unexpected Ether

- Contracts that rely on code execution for all ether sent to them can be vulnerable to attacks where ether is forcibly sent
- Related links (TODO: link)
  - How to Secure Your Smart Contracts
  - Solidity Security Patterns - Forcing Ether to a Contract

### The Vulnerability

- A common defensive programming technique that is useful in enforcing correct state transitions or validating operations is **invariant checking**. This technique involves defining a set of invariants (metrics or parameters that should not change) and checking that they remain unchanged after a single (or many) operation(s)
- **WHY**: Misconception that a contract can only accept or obtain ether via payable functions
  - Example is the use of `this.balance`
- Contracts can receive ethers without payable functions or executing any code
  - As receipent of `selfdestruct()` forced by attackers demo as (TODO: link)
  - As receipent of pre-sent ether before deployment due to the determinstic contract address as
    ```js
    address = sha3(rlp.encode([account_address, transaction_nonce]))
    ```
- Examples
  - [EtherGame](examples/unexpected-ethers/EtherGame.sol)
    - **CAUSE**: Line 14 and 32
    - **HOW**
      - Fund the contract by `selfdestruct()` to make the `this.balance` be non-multiple of 0.5 ether
      - Lock all ethers by forcibly funding 10 ethers due to missing milestones

### Preventative Techniques

- Contract logic, when possible, should avoid being dependent on exact values of the balance of the contract, because it can be artificially manipulated
- Make a self-defined variable incremented in payable functions, to safely track the deposited ether as [EtherGameOK.sol](examples/unexpected-ethers/EtherGameOK.sol)

### Further Examples

- Underhanded Solidity Coding Contest (TODO: link)

## DELEGATECALL

- Standard external message calls to contracts are handled by the `CALL` opcode, whereby code is run in the context of the external contract/function
- The `DELEGATECALL` opcode is almost identical, except that the code executed at the targeted address is run in the context of the calling contract, and `msg.sender` and `msg.value` remain unchanged
- Related links (TODO: links)
  - Loi.Luu's Ethereum Stack Exchange question on this topic
  - The Solidity docs

### The Vulnerability

- The code in libraries themselves can be secure and vulnerability-free
- The context-preserving nature of `DELEGATECALL` introduce new vulnerabilities
- Demo as [FibonacciBalance.sol](examples/delegatecall/FibonacciBalance.sol) against [FibonacciLib.sol](examples/delegatecall/FibonacciLib.sol)
  - The fallback function in the `FibonacciBalance` contract allows all calls to be passed to the library contract
  - The storage layout of state variables in contract make it possible to modify the states the `FibonacciBalance` contract due to context preserving of `delegatecall`
    - `setStart()` allows modify the `slot[0]` of `FibonacciBalance`, thus enabling changing the address of `fibLibrary` (e.g. as [`Attack` contract](examples/delegatecall/Attack.sol))
- When we say that delegatecall is state-preserving, we are not talking about the variable names of the contract, but rather the actual storage slots to which those names point

### Preventative Techniques

- Employ `library` keyword for implementing library contracts
  - Forcing libraries to be stateless mitigates the complexities of storage context
  - Stateless libraries also prevent attacks wherein attackers modify the state of the library directly in order to affect the contracts that depend on the library's code
- As a general rule of thumb, when using `DELEGATECALL` pay careful attention to the possible calling context of both the library contract and the calling contract, and whenever possible build stateless libraries.

### Real-World Example: Parity Multisig Wallet (Second Hack)

- Examples (TODO: links)
  - Parity Multisig Hacked. Again
  - An In-Depth Look at the Parity Multisig Bug
- Example: [WalletLibrary.sol](examples/delegatecall/WalletLibrary.sol) and [Wallet.sol](examples/delegatecall/Wallet.sol)
  - The `WalletLibrary` contract is itself a **contract** and maintains its own state
  - The owner of `WalletLibrary` contract can destruct the `WalletLibrary`, rendering any later payment to the `Wallet` referencing the deleted `WalletLibrary` lost forever

## Default Visibilities

### The Vulnerability

- The intended `private` API is missed out, and default to `public`
- Example: [HashForEther.sol](examples/default-visibility/HashForEther.sol)
  - The unintended public `_sendWinnings` function allows any address to steal the bounty

### Preventative Techniques

- Always specify the visibility of all functions in a contract, even if they are intentionally `public`

### Real-World Example: Parity Multisig Wallet (First Hack)

- Analysis as Haseeb Qureshi (TODO: link)
- Example: [WalletLibrary.sol](examples/default-visibility/WalletLibrary.sol)
  - The unintended public `initMultiowned` enables an attacker to
    - Call these functions on deployed contracts, resetting the ownership to the attacker's address
    - Then drain the wallets of all their ether

## Entropy Illusion

- Every transaction modifies the global state of the Ethereum ecosystem in a calculable way, with no uncertainty
- This has the fundamental implication that there is no source of entropy or randomness in Ethereum
- Achieving decentralized entropy (randomness) is a well-known problem for which many solutions have been proposed, including RANDAO (TODO: link), or using a chain of hashes, as described by Vitalik Buterin in the blog post "Validator Ordering and Randomness in PoS" (TODO: link)

### The Vulnerability

- Gambling requires uncertainty (something to bet on), which makes building a gambling system on the blockchain (a deterministic system) rather difficult
- The uncertainty must come from a source external to the blockchain
- Option 1: future block variables
  - **Demerit**: Controlled by miners
- Option 2: past or present variables
  - Analysis as Martin Swende (TODO: link)
- Using solely block variables means that the pseudorandom number will be the same for all transactions in a block, so an attacker can multiply their wins by doing many transactions within a block (should there be a maximum bet)

### Preventative Techniques

- The source of entropy (randomness) must be external to the blockchain
- Options (TODO: links)
  - Commit–reveal
  - Changing the trust model to a group of participants (as in RandDAO)
  - A centralized entity that acts as a randomness oracle

### Real-World Example: PRNG Contracts

- Analyzed by Arseny Reutov in Feb. 2018 (TODO: link)

## External Contract Referencing

### Real-World Example: Reentrancy Honey Pot

## Short Address/Parameter Attack

## Unchecked CALL Return Values

### Real-World Example: Etherpot and King of the Ether

## Race Conditions/Front Running

### Real-World Examples: ERC20 and Bancor

## Denial of Service (DoS)

### Real-World Examples: GovernMental

## Block Timestamp Manipulation

### Real-World Example: GovernMental

## Constructors with Care

### Real-World Example: Rubixi

## Uninitialized Storage Pointers

### Real-World Examples: OpenAddressLottery and CryptoRoulette Honey Pots

## Floating Point and Precision

### Real-World Example: Ethstick

## Tx.Origin Authentication

## Contract Libraries

## Conclusions