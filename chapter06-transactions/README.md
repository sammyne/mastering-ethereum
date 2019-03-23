# Chapter 06. Transactions

- Transactions are signed messages
  - originated by an externally owned account
  - transmitted by the Ethereum network, and
  - recorded on the Ethereum blockchain
- Ethereum is a global singleton state machine, and transactions are what make that state machine "tick," changing its state

## The Structure of a Transaction

- The network-serialization is the only standard form of a transaction as

  |     Field | Description                                                                      |
  | --------: | :------------------------------------------------------------------------------- |
  |     Nonce | A sequence number, issued by the originating EOA, used to prevent message replay |
  | Gas price | The price of gas (in wei) the originator is willing to pay                       |
  | Gas limit | The maximum amount of gas the originator is willing to buy for this transaction  |
  | Recipient | The destination Ethereum address                                                 |
  |     Value | The amount of ether to send to the destination                                   |
  |      Data | The variable-length binary data payload                                          |
  |   `v,r,s` | The three components of an ECDSA digital signature of the originating EOA        |

- The transaction message's structure is serialized using the **Recursive Length Prefix** (RLP) encoding scheme
- All numbers in Ethereum are encoded as big-endian integers, of lengths that are multiples of 8 bits
- RLP does not contain any field delimiters or labels
- Since the address can be recovered from `v,r,s`, no address of the payer is included

## The Transaction Nonce

- **WHAT**: A scalar value equal to the number of transactions sent from this address or, in the case of accounts with associated code, the number of contract-creations made by this account
- ??The nonce is an attribute of the originating address; that is, it only has meaning in the context of the sending address
- Nonce is calculated dynamically, by counting the number of **confirmed** transactions that have originated from an address.
- 2 scenarios
  - The usability feature of transactions being included in the order of creation
    - Tx with larger nonces will be processed after those with smaller nonces
  - The vital feature of transaction duplication protection
    - By having the incrementing nonce as part of the transaction, it is simply impossible for anyone to "duplicate" a payment you have made

### Keeping Track of Nonces

- The nonce is an up-to-date zero-based count of the number of **confirmed** (i.e., on-chain) transactions that have originated from an account
- When you create a new transaction, you assign the next nonce in the sequence. But until it is confirmed, it will not count toward the `getTransactionCount` total

- TODO: example demonstrating the `getTransactionCount`

- Only when the pending and confirmed counts are equal (all outstanding transactions are confirmed) can you trust the output of `getTransactionCount` to start your nonce counter
- Parity's JSON RPC interface offers the `parity_nextNonce` function, which returns the next nonce that should be used in a transaction

### Gaps in Nonces, Duplicate Nonces, and Confirmation

- The Ethereum network processes transactions sequentially, based on the nonce
  - The txs of larger nonces will always remain pending until those of smaller nonces have been confirmed
- Invalid tx or out-of-gas tx will induce gaps in nonces, making all the subsequent transactions "stuck" in waiting for the missing nonce
- Once the gaps of nonces is validated by the network, all the broadcast transactions with subsequent nonces will incrementally become valid; it is not possible to "recall" a transaction

### Concurrency, Transaction Origination, and Nonces

- Concurrency is when you have simultaneous computation by multiple independent systems in form of
  - Multithreading in the same program
  - Multiprocessing on the same CPU
  - Distributed systems on different computers
- How do multiple computers generating, signing, and broadcasting transactions from the same hot wallet account coordinate?
  - A single computer assigns nonces on a first-come first-served basis, to computers signing transactions. However, this computer is now a single point of failure
  - Forward generated transactions (leaving out nonces) to a single node which
    - Keeps track of nonces
    - Assign the nonce and sign the txs
    - is likely to become congested under load
- The difficulties of synchronizing nonces forces most implementations toward avoiding concurrency and creating bottlenecks such as
  - a single process handling all withdrawal transactions in an exchange
  - setting up multiple hot wallets that can work completely independently for withdrawals and only need to be intermittently rebalanced

## Transaction Gas

## Transaction Recipient

## Transaction Value and Data

### Transmitting Value to EOAs and Contracts

### Transmitting a Data Payload to an EOA or Contract

## Special Transaction: Contract Creation

## Digital Signatures

### The Elliptic Curve

### Digital Signature Algorithm

### How Digital Signatures Work

### Verifying the Signature

### ECDSA Math

### Transaction Signing in Practice

### Raw Transaction Creation and Signing Raw Transaction Creation with EIP-155

## The Signature Prefix Value (v) and Public Key Recovery

## Separating Signing and Transmission (Offline Signing)

## Transaction Propagation

## Recording on the Blockchain

## Multiple-Signature (Multisig) Transactions

## Conclusions
