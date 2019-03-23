# Chapter 05. Wallets

- At a high level, a wallet is a software application that
  - Serves as the primary user interface to Ethereum
  - Controls access to a user's money by
    - Managing keys and addresses
    - Tracking the balance
    - Creating and signing transactions
- From a programmer's perspective, the word wallet refers to the system used to store and manage a user's keys

## Wallet Technology Overview

- One key consideration in designing wallets is balancing convenience and privacy
- Reusing private keys or addresses suffering from the vulnerability that anyone can easily track and correlate all your transactions
- Users control the tokens on the network by signing transactions with the keys in their wallets
- **An Ethereum wallet is a keychain**

> **Ethereum wallets contain keys, not ether or tokens**. Wallets are like keychains containing pairs of private and public keys. Users sign transactions with the private keys, thereby proving they own the ether. The ether is stored on the blockchain

- 2 types of wallets
  - **Nondeterministic wallet**: Each key is independently generated from a different random number, a.k.a., JBOK ("Just a Bunch of Keys") wallet
  - **Deterministic wallet**: All the keys are derived from a single master key, known as the **seed**
- Mnemonic code words
  - **WHY**: To make deterministic wallets slightly more secure against data-loss accidents, such as having your phone stolen or dropping it in the toilet
  - **WHAT**: a list of words (in English or another language) encoding the seed for you to write down and use in the event of an accident

### Nondeterministic (Random) Wallets

- **Such walles are being replaced with the deterministic ones**
- **GOOD PRACTICE**: Avoid address reuse to maximize users's privacy by employing a new address for every receiveing fund
- Many Ethereum clients (including `geth`) use a **keystore** file, which is a JSON-encoded file that contains a single (randomly generated) private key, encrypted by a passphrase for extra security
- The keystore format uses a **key derivation function** (KDF) (a.k.a., password stretching algorithm), which protects against brute-force, dictionary, and rainbow table attacks
- Keystore encoding library: [keythereum](https://github.com/ethereumjs/keythereum)

### Deterministic (Seeded) Wallets

- Private keys that are all derived from a single master key, or **seed**
- The seed is sufficient to
  - Recover all the derived keys
  - Export a wallet
  - Import a wallet

### Hierarchical Deterministic Wallets (BIP-32/BIP-44)

![HD wallet: a tree of keys generated from a single seed](images/hd_wallets.png)

- Advantages
  - The tree structure can be used to express additional organizational meaning
  - Users can create a sequence of public keys without having access to the corresponding private keys

### Seeds and Mnemonic Codes (BIP-39)

- The use of a mnemonic word list to encode the seed for an HD wallet makes for the easiest way to safely
  - Export
  - Transcribe
  - Record on paper
  - Read without error
  - Import a private key set into another wallet

## Wallet Best Practices

- Common standards

  | Standard BIP | Description                            |
  | -----------: | :------------------------------------- |
  |           39 | Mnemonic code words                    |
  |           32 | HD wallets                             |
  |           43 | Multipurpose HD wallet structure       |
  |           44 | Multicurrency and multiaccount wallets |

- Compatible software wallets
  - Jaxx
  - MetaMask
  - MyCrypto
  - MyEtherWallet (MEW)
- Compatible hardware wallets
  - Keepkey
  - Ledger
  - Trezor

### Mnemonic Code Words (BIP-39)

- **WHAT**: Word sequences that encode a random number used as a seed to derive a deterministic wallet
  > The primary difference is that a brainwallet consists of words chosen by the user, whereas mnemonic words are created randomly by the wallet and presented to the user

#### Generating mnemonic words

- Generating entropy and encoding as mnemonic words is depicted as

  ![Generating entropy and encoding as mnemonic words](images/bip39-part1.png)

- Mnemonic codes: entropy and word length goes as

  | Entropy (bits) | Checksum (bits) | Mnemonic Length (words) |
  | -------------- | --------------- | ----------------------- |
  | 128            | 4               | 12                      |
  | 160            | 5               | 15                      |
  | 192            | 6               | 18                      |
  | 224            | 7               | 21                      |
  | 256            | 8               | 24                      |

#### From mnemonic to seed

- From mnemonic to seed is depicted as  
  ![From mnemonic to seed](images/bip39-part2.png)
- The key-stretching function takes two parameters: the **mnemonic** and a **salt**
- The salt serves for 2 purposes

  - Make it difficult to build a lookup table enabling a brute-force attack
  - Allow the introduction of a passphrase that serves as an additional security factor protecting the seed

- TODO: example 5-{2,3,4}

#### Optional passphrase in BIP-39

- **WHY**
  - **A second factor** (something memorized) that makes a mnemonic useless on its own, protecting mnemonic backups from compromise by a thief
  - **A form of plausible deniability** or "duress wallet," where a chosen passphrase leads to a wallet with a small amount of funds, used to distract an attacker from the "real" wallet that contains the majority of funds
- Risks
  - If the wallet owner is incapacitated or dead and no one else knows the passphrase, the seed is useless and all the funds stored in the wallet are lost forever
  - Conversely, if the owner backs up the passphrase in the same place as the seed, it defeats the purpose of a second factor

#### Working with mnemonic codes

- Libraries
  - [python-mnemonic](https://github.com/trezor/python-mnemonic)
  - [ConsenSys/eth-lightwallet](https://github.com/ConsenSys/eth-lightwallet)
  - [npm/bip39](https://www.npmjs.com/package/bip39)
  - the web-based [Mnemonic Code Converter](https://iancoleman.io/bip39/)

### Creating an HD Wallet from the Seed

- HD wallets are created from a single root seed, which is a 128-, 256-, or 512-bit random number

### HD Wallets (BIP-32) and Paths (BIP-43/44)

#### Extended public and private keys

- A very useful characteristic of HD wallets is the ability to derive child public keys from parent public keys, without having the private keys
  - An extended public key can be used to derive all of the public keys (and only the public keys) in that branch of the HD wallet structure
- Application
  - Public key–only deployments, where a server or application has a copy of an extended public key, but no private keys
  - Cold-storage or hardware wallets: private keys on hardware wallets, public keys online

#### Hardened child key derivation

- **WHY**
  - Since the `xpub` contains the chain code, a known/leaked child private key can be used with the chain code to derive
    - All the other child private keys
    - Deduce the parent private key
- The hardened derivation function uses the parent private key to derive the child chain code, instead of the parent public key
- Best practice is to have the level-1 children of the master keys always derived by hardened derivation, to prevent compromise of the master keys

#### Index numbers for normal and hardened derivation

- Each index number, when combined with a parent key using the special child derivation function, gives a different child key
- Index range for derivation
  - Normal: [0, 2<sup>31</sup>–1] (0x0 to 0x7FFFFFFF)
  - Hardended: [2<sup>31</sup>, 2<sup>32</sup>–1] (0x80000000 to 0xFFFFFFFF)

> To make the index numbers easier to read and display, the index numbers for hardened children are displayed starting from zero, but with a prime symbol

#### HD wallet key identifier (path)

- Private keys derived from the master private key start with "m"
- Public keys derived from the master public key start with "M"
- Examples go as

  |       HD path | Key described                                                                                                      |
  | ------------: | :----------------------------------------------------------------------------------------------------------------- |
  |         `m/0` | The first (`0`) child private key from the master private key (`m`)                                                |
  |       `m/0/0` | The first grandchild private key of the first child (`m/0`)                                                        |
  |      `m/0'/0` | The first normal grandchild of the first hardened child (`m/0'`)                                                   |
  |       `m/1/0` | The first grandchild private key of the second child (`m/1`)                                                       |
  | `M/23/17/0/0` | The first great-great-grandchild public key of the first great-grandchild of the 18th grandchild of the 24th child |

#### Navigating the HD wallet tree structure

- BIP-43 proposes the use of the first hardened child index as a special identifier that signifies the **"purpose"** of the tree structure
- BIP-44 proposes a multiaccount structure as **"purpose" number 44'** under BIP-43
- BIP-44 specifies the structure as consisting of five predefined tree levels:

  ```
  m / purpose' / coin_type' / account' / change / address_index
  ```

  |        Level | Description                                                                                                                                                            |
  | -----------: | :--------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
  |   `purpose'` | always set to `44'`                                                                                                                                                    |
  | `coin_type'` | detailed as [SLIP0044](https://github.com/satoshilabs/slips/blob/master/slip-0044.md), Ethereum is `m/44'/60'`, Ethereum Classic is `m/44'/61'`, Bitcoin is `m/44'/0'` |
  |   `account'` | allows users to subdivide their wallets into separate logical subaccounts, for accounting or organizational purposes                                                   |
  |     `change` | an HD wallet has two subtrees, one for creating receiving addresses (indexed by `0`) and one for creating change addresses (indexed by `1`)                            |

- Examples of BIP-44 HD wallet structure

  |            HD path | Key described                                                              |
  | -----------------: | :------------------------------------------------------------------------- |
  | `M/44'/60'/0'/0/2` | The third receiving public key for the primary Ethereum account            |
  | `M/44'/0'/3'/1/14` | The 15th change-address public key for the 4th Ethereum account            |
  |  `m/44'/2'/0'/0/1` | The 2nd private key in the Litecoin main account, for signing transactions |

## Conclusions
