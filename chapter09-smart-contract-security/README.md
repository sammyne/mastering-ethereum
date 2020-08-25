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
- Employ the **[checks-effects-interactions]** pattern
  - For any code that performs external calls to unknown addresses to be the last operation in a localized function or piece of code execution
- Use mutex to prevent reentrances
- A fixed `EtherStore` with all above applied go as [EtherStoreOK.sol](examples/reentrancy/EtherStoreOK.sol)

[checks-effects-interactions]: https://solidity.readthedocs.io/en/latest/security-considerations.html#use-the-checks-effects-interactions-pattern

### Real-World Example: The DAO

- Consequence: \$150 million loss and a hark fork creating ETC
- Analysis as [analysis-of-the-dao-exploit]

[analysis-of-the-dao-exploit]: https://hackingdistributed.com/2016/06/18/analysis-of-the-dao-exploit

## Arithmetic Over/Underflows

- Related links (TODO: links)
  - [How to Secure Your Smart Contracts]
  - [Ethereum Smart Contract Best Practices]
  - [Ethereum, Solidity and integer overflows]: programming blockchains like 1970

[How to Secure Your Smart Contracts]: https://medium.com/loom-network/how-to-secure-your-smart-contracts-6-solidity-vulnerabilities-and-how-to-avoid-them-part-1-c33048d4d17d
[Ethereum Smart Contract Best Practices]: https://consensys.github.io/smart-contract-best-practices/known_attacks/#integer-overflow-and-underflow
[Ethereum, Solidity and integer overflows]: https://randomoracle.wordpress.com/2018/04/27/ethereum-solidity-and-integer-overflows-programming-blockchains-like-1970/

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
- Recommendation: OpenZeppelin
  - The [SafeMath] to tackle over/underflows as [OverflowFreeTimeLock.sol](examples/overflow-underflow/OverflowFreeTimeLock.sol)

[OpenZeppelin]: https://github.com/OpenZeppelin/openzeppelin-contracts
[SafeMath]: https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/math/SafeMath.sol

### Real-World Examples: PoWHC and Batch Transfer Overflow (CVE-2018–10299)

- Proof of Weak Hands Coin (PoWHC), originally devised as a joke of sorts, was a Ponzi scheme written by an internet collective
- PoWHC suffers from underflow as [explained][attck on PoWHC] by Eric Banisadr (TODO: link)
- PeckShield's account suffers from overflow

[attck on PoWHC]: https://medium.com/@ebanisadr/how-800k-evaporated-from-the-powh-coin-ponzi-scheme-overnight-1b025c33b530
[PeckShield's account]: https://medium.com/@peckshield/alert-new-batchoverflow-bug-in-multiple-erc20-smart-contracts-cve-2018-10299-511067db6536

## Unexpected Ether

- Contracts that rely on code execution for all ether sent to them can be vulnerable to attacks where ether is forcibly sent
- Related links
  - [How to Secure Your Smart Contracts - Part 2] 
  - [Solidity Security Patterns - Forcing Ether to a Contract]

[How to Secure Your Smart Contracts - Part 2]: https://medium.com/loom-network/how-to-secure-your-smart-contracts-6-solidity-vulnerabilities-and-how-to-avoid-them-part-2-730db0aa4834
[Solidity Security Patterns - Forcing Ether to a Contract]: http://danielszego.blogspot.com/2018/03/solidity-security-patterns-forcing.html

### The Vulnerability

- A common defensive programming technique that is useful in enforcing correct state transitions or validating operations is **invariant checking**. This technique involves defining a set of invariants (metrics or parameters that should not change) and checking that they remain unchanged after a single (or many) operation(s)
- **WHY**: Misconception that a contract can only accept or obtain ether via payable functions
  - Example is the use of `this.balance`
- Contracts can receive ethers without payable functions or executing any code
  - As receipent of `selfdestruct()` forced by attackers [demo][attack based on selfdestruct] by Martin Swende
  - As receipent of pre-sent ether before deployment due to the determinstic contract address as
    ```js
    address = sha3(rlp.encode([account_address, transaction_nonce]))
    ```
- Examples
  - [EtherGame](examples/unexpected-ethers/EtherGame.sol)
    - **CAUSE**: Line 17 and 35
    - **HOW**
      - Fund the contract by `selfdestruct()` to make the `this.balance` be non-multiple of 0.5 ether
      - Lock all ethers by forcibly funding 10 ethers due to missing milestones

[attack based on selfdestruct]: https://swende.se/blog/Ethereum_quirks_and_vulns.html

### Preventative Techniques

- Contract logic, when possible, should avoid being dependent on exact values of the balance of the contract, because it can be artificially manipulated
- Make a self-defined variable incremented in payable functions, to safely track the deposited ether as [EtherGameOK.sol](examples/unexpected-ethers/EtherGameOK.sol)

### Further Examples

- [Underhanded Solidity Coding Contest]

[Underhanded Solidity Coding Contest]: https://github.com/Arachnid/uscc/tree/master/submissions-2017/

## DELEGATECALL

- Standard external message calls to contracts are handled by the `CALL` opcode, whereby code is run in the context of the external contract/function
- The `DELEGATECALL` opcode is almost identical, except that the code executed at the targeted address is run in the context of the calling contract, and `msg.sender` and `msg.value` remain unchanged
- Related links (TODO: links)
  - [Loi.Luu's Ethereum Stack Exchange question on this topic][difference-between-call-callcode-and-delegatecall]
  - [The Solidity docs][delegatecall-callcode-and-libraries]

[difference-between-call-callcode-and-delegatecall]: https://ethereum.stackexchange.com/questions/3667/difference-between-call-callcode-and-delegatecall
[delegatecall-callcode-and-libraries]: https://solidity.readthedocs.io/en/latest/introduction-to-smart-contracts.html#delegatecall-callcode-and-libraries


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
- As a general rule of thumb, when using `DELEGATECALL`, pay careful attention to the possible calling context of both the library contract and the calling contract, and whenever possible, build stateless libraries.

### Real-World Example: Parity Multisig Wallet (Second Hack)

- Examples
  - [Parity Multisig Hacked. Again]
  - [An In-Depth Look at the Parity Multisig Bug]
- Example: [WalletLibrary.sol](examples/delegatecall/WalletLibrary.sol) and [Wallet.sol](examples/delegatecall/Wallet.sol)
  - The `WalletLibrary` contract is itself a **contract** and maintains its own state
  - The owner of `WalletLibrary` contract can destruct the `WalletLibrary`, rendering any later payment to the `Wallet` referencing the deleted `WalletLibrary` lost forever

[Parity Multisig Hacked. Again]: https://medium.com/chain-cloud-company-blog/parity-multisig-hack-again-b46771eaa838
[An In-Depth Look at the Parity Multisig Bug]: https://hackingdistributed.com/2017/07/22/deep-dive-parity-bug/

## Default Visibilities

### The Vulnerability

- The intended `private` API is missed out, and default to `public`
- Example: [HashForEther.sol](examples/default-visibility/HashForEther.sol)
  - The unintended public `_sendWinnings` function allows any address to steal the bounty

### Preventative Techniques

- Always specify the visibility of all functions in a contract, even if they are intentionally `public`

### Real-World Example: Parity Multisig Wallet (First Hack)

- Analysis as [Haseeb Qureshi][Parity Multisig Wallet (First Hack)]
- Example: [WalletLibrary.sol](examples/default-visibility/WalletLibrary.sol)
  - The unintended public `initMultiowned` enables an attacker to
    - Call these functions on deployed contracts, resetting the ownership to the attacker's address
    - Then drain the wallets of all their ether

[Parity Multisig Wallet (First Hack)]: https://www.freecodecamp.org/news/a-hacker-stole-31m-of-ether-how-it-happened-and-what-it-means-for-ethereum-9e5dc29e33ce/

## Entropy Illusion

- Every transaction modifies the global state of the Ethereum ecosystem in a calculable way, with no uncertainty
- This has the fundamental implication that there is no source of entropy or randomness in Ethereum
- Achieving decentralized entropy (randomness) is a well-known problem for which many solutions have been proposed, including [RANDAO], or using a chain of hashes, as described by Vitalik Buterin in the blog post "Validator Ordering and Randomness in PoS" (deprecated)

### The Vulnerability

- Gambling requires uncertainty (something to bet on), which makes building a gambling system on the blockchain (a deterministic system) rather difficult
- The uncertainty must come from a source external to the blockchain
- Option 1: future block variables
  - **Demerit**: Controlled by miners
- Option 2: past or present variables
  - Analysis as [Martin Swende][An Ethereum Roulette]
- Using solely block variables means that the pseudorandom number will be the same for all transactions in a block, so an attacker can multiply their wins by doing many transactions within a block (should there be a maximum bet)

[An Ethereum Roulette]: https://swende.se/blog/Breaking_the_house.html

### Preventative Techniques

- The source of entropy (randomness) must be external to the blockchain
- Options
  - [Commit–reveal]
  - Changing the trust model to a group of participants (as in [RandDAO])
  - A centralized entity that acts as a randomness oracle

[Commit–reveal]: https://ethereum.stackexchange.com/questions/191/how-can-i-securely-generate-a-random-number-in-my-smart-contract

### Real-World Example: PRNG Contracts

- Analyzed by [Arseny Reutov][Predicting Random Numbers in Ethereum Smart Contracts] in Feb. 2018

[Predicting Random Numbers in Ethereum Smart Contracts]: https://blog.positive.com/predicting-random-numbers-in-ethereum-smart-contracts-e5358c6b8620

## External Contract Referencing

### The Vulnerability

- In Solidity, any address can be cast to a contract, regardless of whether the code at the address represents the contract type being cast
- Example: [EncryptionContract.sol](examples/external-contracts-referencing/EncryptionContract.sol) to deploy with [Rot13Encryption.sol](examples/external-contracts-referencing/Rot13Encryption.sol)
  - **HOW**
    - Replace the provided `_encryptionLibrary` to `EncryptionContract` with `Rot26Encryption` or `Print` contract
    - Or with `Blank` contract by putting malicious codes in the fallback to be called due to no matching functions

### Preventative Techniques

- Use the `new` keyword to create contracts internally to render the referenced contract immutable
- Hardcode external contract addresses
- The address of the referenced external contracts should be public for auditing
- If a user can change a contract address that is used to call external functions, it can be important (in a decentralized system context) to implement a time-lock and/or voting mechanism to allow users to see what code is being changed, or to give participants a chance to opt in/out with the new contract address

### Real-World Example: Reentrancy Honey Pot

- **Honey Pot**: Contracts try to outsmart Ethereum hackers who try to exploit the contracts, but who in turn end up losing ether to the contract they expect to exploit
- Example: [Log.sol](examples/external-contracts-referencing/Log.sol)
  - Reentrance by exploiting line 29 would trigger OOG thus reverting any tx
  - Detail as [Reentrancy Honey Pot], but still **DON'T UNDERSTAND :(**

[Reentrancy Honey Pot]: https://www.reddit.com/r/ethdev/comments/7x5rwr/tricked_by_a_honeypot_contract_or_beaten_by/

## Short Address/Parameter Attack

- Related links
  - [The ERC20 Short Address Attack Explained]
  - [ICO Smart Contract Vulnerability: Short Address Attack]
  - [This Reddit post]

[The ERC20 Short Address Attack Explained]: https://vessenes.com/the-erc20-short-address-attack-explained/
[ICO Smart contract Vulnerability: Short Address Attack]: https://medium.com/huzzle/ico-smart-contract-vulnerability-short-address-attack-31ac9177eb6b
[This Reddit post]: https://www.reddit.com/r/ethereum/comments/6r9nhj/cant_understand_the_erc20_short_address_attack/

### The Vulnerability

- It is possible to send encoded parameters that are shorter than the expected parameter length according to the ABI specification
- **WHEN**: Third-party applications do not validate inputs
  - Example: The ERC20 Short Address Attack Explained
    - Given `transfer` function as
      ```solidity
      function transfer(address to, uint tokens) public returns (bool success);
      ```
    - To withdraw 100 tokens to address `0xdeaddeaddeaddeaddeaddeaddeaddeaddeaddead`, the correct ABI-encoded data should be (where `-` is added to ease readibility only, which doesn't actually exist)
      ```
      a9059cbb-000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddeaddead-0000000000000 000000000000000000000000000000000056bc75e2d63100000
      ```
    - A well-crafted data payload is
      ```
      a9059cbb-000000000000000000000000deaddeaddeaddeaddeaddeaddeaddeaddeadde00-00000000000000000000000000000000000000000000056bc75e2d6310000000
      ```
      by leaving out the last byte of the recipient address, thus rendering the tokens stealing possible if not watched carefully

### Preventative Techniques

- All input parameters in external applications should be validated before sending them to the blockchain

## Unchecked `CALL` Return Values

- The `call` and `send` functions return a Boolean indicating whether the call succeeded or failed, but not reverting the calling in case of failure
- A common error is that the developer expects a revert to occur if the external call fails, and does not check the return value
- Related links (TODO: links)
  - DASP Top 10 of 2018
  - Scanning Live Ethereum Contracts for the 'Unchecked-Send' Bug

[DASP Top 10 of 2018]: https://www.dasp.co/#item-4
[Scanning Live Ethereum Contracts for the 'Unchecked-Send' Bug]: https://hackingdistributed.com/2016/06/16/scanning-live-ethereum-contracts-for-bugs/

### The Vulnerability

- Example: [Lotto.sol](examples/unchecked-call-returned-value/Lotto.sol)
  - The missing failure check of line 13
    - `sendToWinner` would set on `payout` regardless of `winner.send(winAmount)`
    - Anyone after the `sendToWinner` can take all the leftover

### Preventative Techniques

- Prefer `transfer` than `send`
- If `send`, always check the returned value
- Recommendation: the withdrawal pattern
  - Each user must call an isolated withdraw function that handles the sending of ether out of the contract and deals with the consequences of failed send transactions. The idea is to logically isolate the external send functionality from the rest of the codebase, and place the burden of a potentially failed transaction on the end user calling the withdraw function.

### Real-World Example: Etherpot and King of the Ether

- **WHY**: incorrect use of block hashes as explained by [Aakil Fernandes](http://aakilfernandes.github.io/blockhashes-are-only-good-for-256-blocks)
- Sample code

  ```solidity
  ...
  function cash(uint roundIndex, uint subpotIndex){

      var subpotsCount = getSubpotsCount(roundIndex);

      if(subpotIndex>=subpotsCount)
          return;

      var decisionBlockNumber = getDecisionBlockNumber(roundIndex,subpotIndex);

      if(decisionBlockNumber>block.number)
          return;

      if(rounds[roundIndex].isCashed[subpotIndex])
          return;
      //Subpots can only be cashed once. This is to prevent double payouts

      var winner = calculateWinner(roundIndex,subpotIndex);
      var subpot = getSubpot(roundIndex);

      winner.send(subpot);

      rounds[roundIndex].isCashed[subpotIndex] = true;
      //Mark the round as cashed
  }
  ...
  ```

  - **HOW**: `winner.send(subpot)` isn't checked, which may produce a state where the winner does not receive their ether, but the state of the contract can indicate that the winner has already been paid

## Race Conditions/Front Running

- Related links
  - [Ethereum Wiki]
  - [\#7 on the DASP Top10 of 2018]
  - The Ethereum Smart Contract Best Practices

[race conditions on Ethereum Wiki]: https://github.com/ethereum/wiki/wiki/Safety#race-conditions
[#7 on the DASP Top10 of 2018]: https://www.dasp.co/#item-7
[The Ethereum Smart Contract Best Practices]: https://consensys.github.io/smart-contract-best-practices/known_attacks/#race-conditions

### The Vulnerability

- **HOW**: An attacker can watch the transaction pool for transactions that may contain solutions to problems, and modify or revoke the solver's permissions or change state in a contract detrimentally to the solver
- Example contract as [FindThisHash.sol](examples/race-condition-front-running/FindThisHash.sol)
  - For a solved solution submit as tx, anyone knowing the tx can make a higher-`gasPrice` tx to be favored for mining

### Preventative Techniques

- Two kinds of actors
  - Users modifying the `gasPrice` (This is worse than below)
  - Miners filtering tx
- To guard only against malicious users
  - Place an upper bound on the `gasPrice` to disable tx ordering by `gasPrice`
- Another option against both actors: [a commit–reveal scheme]
  - **HOW**
    - Users send transactions with hidden information (typically a hash)
    - After the transaction has been included in a block, the user sends a transaction revealing the data that was sent (the reveal phase)
  - **Cons**: Cannot conceal the transaction value
- Related links
  - [ENS contracts][ENS]
  - [submarine sends]

[a commit–reveal scheme]: https://ethereum.stackexchange.com/questions/191/how-can-i-securely-generate-a-random-number-in-my-smart-contract
[ENS]: https://ens.domains/
[submarine sends]: https://hackingdistributed.com/2017/08/28/submarine-sends/

### Real-World Examples: ERC20 and Bancor

- The ERC20 standard is quite well-known for building tokens on Ethereum. This standard has a potential front-running vulnerability that comes about due to the `approve` function as explained by [Mikhail Vladimirov and Dmitry Khovratovich][ERC20 API: An Attack Vector on Approve/TransferFrom Methods]
  - A approves B 100 tokens
  - B takes them
  - Then A want to reset tokens for B as 50
  - B can take another 50 more tokens
- Analysis for Bancor is given by [Ivan Bogatyy][Implementing Ethereum trading front-runs on the Bancor exchange in Python]

[ERC20 API: An Attack Vector on Approve/TransferFrom Methods]: https://docs.google.com/document/d/1YLPtQxZu1UAvO9cZ1O2RPXBbT0mooh4DYKjA_jp-RLM/edit#heading=h.wqhvh2y0obwt
[Implementing Ethereum trading front-runs on the Bancor exchange in Python]: https://hackernoon.com/front-running-bancor-in-150-lines-of-python-with-ethereum-api-d5e2bfd0d798

## Denial of Service (DoS)

### The Vulnerability

- Looping through externally manipulated mappings or arrays as [DistributeTokens.sol](examples/dos/DistributeTokens.sol)
  - An attacker can create many user accounts, making the investor array large. In principle this can be done such that the gas required to execute the for loop exceeds the block gas limit
- Owner operations

  ```solidity
  bool public isFinalized = false;
  address public owner; // gets set somewhere

  function finalize() public {
      require(msg.sender == owner);
      isFinalized == true;
  }

  // ... extra ICO functionality

  // overloaded transfer function
  function transfer(address _to, uint _value) returns (bool) {
      require(isFinalized);
      super.transfer(_to,_value)
  }

  ...
  ```

  - Owners have specific privileges in contracts and must perform some task in order for the contract to proceed to the next state
  - If the privileged user loses their private keys or becomes inactive, the entire token contract becomes inoperable

- Progressing state based on external calls
  - The external call fails or is prevented for external reasons

### Preventative Techniques

- For [DistributeTokens.sol](examples/dos/DistributeTokens.sol), a withdrawal pattern is recommended, whereby each of the investors call a withdraw function to claim tokens independently.
- For owner operations, a failsafe can be used in the event that the owner becomes incapacitated. Two solutions
  - Make the owner a multisig contract
  - Use a time-lock
- For external calls, account for their possible failure and potentially add a time-based state progression in the event that the desired call never comes

### Real-World Examples: GovernMental

- **HOW**: A [Reddit post][GovernMental's 1100 ETH jackpot payout is stuck because it uses too much gas] by etherik describes how the contract required the deletion of a large mapping in order to withdraw the ether.

[GovernMental's 1100 ETH jackpot payout is stuck because it uses too much gas]: https://www.reddit.com/r/ethereum/comments/4ghzhv/governmentals_1100_eth_jackpot_payout_is_stuck/

## Block Timestamp Manipulation

- Applications of timestamp
  - Entropy for random numbers
  - Locking funds for periods of time
  - Various state-changing conditional statements
- Related links (TODO: links)
  - The Solidity docs - [Block and Transaction Properties]
  - Joris Bontje's Ethereum Stack Exchange question: [Can a contract safely rely on block.timestamp]

[Block and Transaction Properties]: https://solidity.readthedocs.io/en/latest/units-and-global-variables.html#block-and-transaction-properties
[Can a contract safely rely on block.timestamp]: https://ethereum.stackexchange.com/questions/413/can-a-contract-safely-rely-on-block-timestamp?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa

### The Vulnerability

- Example contract as [Roulette.sol](examples/block-timestamp-manipulation/Roulette.sol)
  - If enough ether pools in the contract, a miner who solves a block is incentivized to choose a timestamp such that `block.timestamp` modulo 15 is 0
  - As there is only one person allowed to bet per block, this is also vulnerable to front-running attacks
- In practice, block timestamps are required to be
  - Monotonically increasing
  - Not too far in the future

### Preventative Techniques

- Block timestamps should not be used for entropy or generating random numbers
  - They should not be the deciding factor (either directly or through some derivation) for winning a game or changing an important state
- Time-sensitive applications
  - For unlocking contracts (time-locking), completing an ICO after a few weeks, or enforcing expiry dates
  - Are recommended to use `block.number` and an average block time to estimate times


### Real-World Example: GovernMental

- Link (deprecated)
- **HOW**
  - The contract paid out to the player who was the last player to join (for at least one minute) in a round
  - A miner who was a player could adjust the timestamp (to a future time, to make it look like a minute had elapsed) to make it appear that they were the last player to join for over a minute (even though this was not true in reality)

## Constructors with Care

### The Vulnerability

- **WHY**: If the contract name is modified, or there is a typo in the constructor's name such that it does not match the name of the contract, the constructor will behave like a normal function
- Example contract as [OwnerWallet.sol](examples/constructor-with-care/OwnerWallet.sol)

### Preventative Techniques

- Use the `constructor` keyword as of v0.4.22

### Real-World Example: Rubixi

- With renaming from `DynamicPyramid` to `Rubixi` without changing the constructor name, allows any user to become the creator
- Explanation goes as (TODO: link)

## Uninitialized Storage Pointers

- Related links in solidity doc

  - [Layout of State Variables in Storage]
  - [Layout in Memory]

[Layout of State Variables in Storage]: https://solidity.readthedocs.io/en/latest/internals/layout_in_storage.html
[Layout in Memory]: https://solidity.readthedocs.io/en/latest/internals/layout_in_memory.html

### The Vulnerability

- Local variables within functions default to storage or memory depending on their type
- Uninitialized local storage variables may contain the value of other storage variables in the contract; this fact can cause unintentional vulnerabilities, or be exploited deliberately
- Example contract as [NameRegistrar.sol](examples/uninitialized-storage-pointers/NameRegistrar.sol)
  - State variables are stored sequentially in slots (32 bytes each) as they appear in the contract
  - Solidity by default puts complex data types, such as structs, in storage when initializing them as local variables
  - The uninitialized `newRecord` (line 20) defaults to storage, and `newRecord.name` is mapped to storage slot[0], which currently contains a pointer to `unlocked`
    - If the last byte of `_name` argument is non-zero, `unlocked` is modified directly

### Preventative Techniques

- The Solidity compiler shows a warning for unintialized storage variables; developers should pay careful attention to these warnings when building smart contracts
- Explicitly use the `memory` or `storage` specifiers when dealing with complex types, to ensure they behave as expected

### Real-World Examples: OpenAddressLottery and CryptoRoulette Honey Pots

- [OpenAddressLottery]: A honey pot was deployed that used this uninitialized storage variable quirk to collect ether from some would-be hackers as analyzed by the [Reddit thread][How does this honeypot work? It seems like a private variable gets overwritten by another]
- CryptoRoulette: A honey pot utilizes this trick to try and collect some ether (see "[An Analysis of a Couple Ethereum Honeypot Contracts]")

[OpenAddressLottery]: https://etherscan.io/address/0x741f1923974464efd0aa70e77800ba5d9ed18902#code
[How does this honeypot work? It seems like a private variable gets overwritten by another]: https://www.reddit.com/r/ethdev/comments/7wp363/how_does_this_honeypot_work_it_seems_like_a/
[An Analysis of a Couple Ethereum Honeypot Contracts]: https://medium.com/coinmonks/an-analysis-of-a-couple-ethereum-honeypot-contracts-5c07c95b0a8d

## Floating Point and Precision

- As of this writing (v0.4.24), Solidity does not support fixed-point and floating-point numbers. This means that floating-point representations must be constructed with integer types in Solidity
- Related links
  - [Ethereum Contract Security Techniques and Tips wiki]

[Ethereum Contract Security Techniques and Tips wiki]: https://eth.wiki/en/howto/smart-contract-safety

### The Vulnerability

- Example contract as [FunWithNumbers.sol](examples/floating-point-and-precision/FunWithNumbers.sol)
  - Ignore over/underflow issues
  - **WHY**: The precision is only to the nearest ether
    - For `buyTokens`, if the `value` isn't a multiple of `weiPerEth`, the buyer would get less than expectation
    - For `sellTokens`, if `tokens` isn't a multiple of `tokensPerEth`, the seller would get fewer ether than expected

### Preventative Techniques

- Ensure that any ratios or rates you are using allow for large numerators in fractions
- Be mindful of order of operations to achieve greater precision
  - `a*b/c` would be more precise than `a/c*b`
- When defining arbitrary precision for numbers
  1. Convert values to higher precision
  2. Perform all mathematical operations
  3. Convert back down to the precision required for output

### Real-World Example: Ethstick

- Source of [Ethstick]
- Have issues of rounding at the wei level of precision
- Also suffers from the entropy illusion (TODO: link)
- More see "[Ethereum Contracts Are Going to Be Candy for Hackers]"

[Ethstick]: https://etherscan.io/address/0xbA6284cA128d72B25f1353FadD06Aa145D9095Af#code
[Ethereum Contracts Are Going to Be Candy for Hackers]: https://vessenes.com/ethereum-contracts-are-going-to-be-candy-for-hackers/

## Tx.Origin Authentication

- `tx.origin` traverses the entire call stack and contains the address of the account that originally sent the call (or transaction)
- Related links (TODO: link)
  - "[Tx.Origin and Ethereum Oh My!]" by Peter Vessenes
  - "Solidity: Tx Origin Attacks" by Chris Coverdale

[Tx.Origin and Ethereum Oh My!]: https://vessenes.com/tx-origin-and-ethereum-oh-my/
[Solidity: Tx Origin Attacks]: https://medium.com/coinmonks/solidity-tx-origin-attacks-58211ad95514

### The Vulnerability

- Contracts that authorize users using the `tx.origin` variable are typically vulnerable to phishing attacks that can trick users into performing authenticated actions on the vulnerable contract
- Example: [Phishable.sol](examples/tx-origin-authentication/Phishable.sol) is vulnerable to [AttackContract.sol](examples/tx-origin-authentication/AttackContract.sol)
  - The source code of public contracts is not available by default
  - **HOW**
    - If the victim sends a transaction with enough gas to the `AttackContract` address
    - The fallback function of `AttackContract` is invoked
    - The `withdrawAll` function of the `Phishable` contract with the parameter `attacker` is called
    - All funds from the `Phishable` contract are withdrawn to the attacker address due to the valid `tx.origin` check

### Preventative Techniques

- Never use `tx.origin` for authorization in smart contracts.
- To deny external contracts from calling the current contract, one could implement a require of the form `require(tx.origin == msg.sender)`

## Contract Libraries

- Advantages of using well-established existing on-platform libraries

  - Being able to benefit from the latest upgrades
  - Saves you money and benefits the Ethereum ecosystem by reducing the total number of live contracts in Ethereum

- Reliable sources
  - The most widely used resource is the [OpenZeppelin suite]
    - The contracts in this repository have been extensively tested and in some cases even function as de facto standard implementations
  - Zeppelin is [ZeppelinOS]
    - An open source platform of services and tools to develop and manage smart contract applications securely
- `ethpm` is a package management tool for libraries
  - Website: https://www.ethpm.com/
  - Repository link: https://www.ethpm.com/registry
  - GitHub link: https://github.com/ethpm
  - Documentation: https://www.ethpm.com/docs/integration-guide

[OpenZeppelin suite]: https://openzeppelin.com/contracts/
[ZeppelinOS]: https://openzeppelin.com/sdk/

## Conclusions

[RANDAO]: https://github.com/randao/randao
