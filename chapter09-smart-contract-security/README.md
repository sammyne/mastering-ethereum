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

### Real-World Examples: PoWHC and Batch Transfer Overflow (CVE-2018–10299)

## Unexpected Ether

### Further Examples

## DELEGATECALL

### Real-World Example: Parity Multisig Wallet (Second Hack)

## Default Visibilities

### Real-World Example: Parity Multisig Wallet (First Hack)

## Entropy Illusion

### Real-World Example: PRNG Contracts

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
