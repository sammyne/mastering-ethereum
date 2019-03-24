# Chapter 04. Cryptography

- Cryptography can, for example, also be used to prove
  - Knowledge of a secret without revealing that secret (e.g., with a digital signature), or
  - The authenticity of data (e.g., with digital fingerprints, also known as "hashes")

## Keys and Addresses

- Ownership of ether by EOAs is established through digital **private keys**, **Ethereum addresses**, and **digital signatures**
- Private keys are not used directly in the Ethereum system in any way; they are never transmitted or stored on Ethereum
- Access and control of funds is achieved with digital signatures created using the private key
- In the payment portion of an Ethereum transaction, the intended recipient is represented by an Ethereum address
- An Ethereum address for an **EOA** is generated from the public key portion of a key pair

## Public Key Cryptography and Cryptocurrency

- The key exchange protocol, first published in the 1970s by Martin Hellman, Whitfield Diffie, and Ralph Merkle
- Public key cryptography uses unique keys to secure information
  - These keys are based on mathematical functions that have a special property: it is easy to calculate them, but hard to calculate their inverse
- **Trapdoor functions**: Functions are very difficult to invert unless you are given a piece of secret information as a shortcut
- The basis of Ethereum's (and other cryptocurrencies') use of private keys and digital signatures: the **elliptic curve cryptography** secured by the discrete logarithm problem

## Private Keys

- **DEFINITION**: Simply a randomly picked number
- The private key is used to create signatures required to spend ether by proving ownership of funds used in a transaction
- The private key MUST
  - Always remain secret
  - Be backed up and protected from accidental loss

### Generating a Private Key from a Random Number

- **The first and most important step** in generating keys is to find a secure source of entropy, or randomness
- A private key can be any nonzero number up to a very large number (defined as the the order of the elliptic curve) slightly less than 2<sup>256</sup>
- Private keys should be generated offline
- Manual choice and bad random number generator are bad for helping generating secure private keys

## Public Keys

- An Ethereum public key is a point `K=(x, y)` on an elliptic curve
- The point `K=k*G` is produced from the private key `k` by an one-way calculation, where
  - `G` is the a constant point called the **generator point**
  - `*` is the special elliptic curve "multiplication" operator

### Elliptic Curve Cryptography Explained

> Ethereum uses the exact same elliptic curve, called `secp256k1`, as Bitcoin

- [Example 4-1: Using Go to confim that a point is on the elliptic curve](examples/1_test.go)

### Elliptic Curve Arithmetic Operations

- Elliptic curve addition is defined such that given two points P<sub>1</sub> and P<sub>2</sub> on the elliptic curve, there is a third point P<sub>3</sub>=P<sub>1</sub>+P<sub>2</sub>, also on the elliptic curve
- A point called the "point at infinity," which roughly corresponds to the role of the number zero in addition
- For a point `P` on the elliptic curve, if `k` is a whole number, then `k * P = P + P + P + ... + P` (`k` times)

### Generating a Public Key

- The Ethereum's public keys represented as a serialization of 130 hexadecimal characters (65 bytes) according to [SECG standard](http://www.secg.org/sec1-v2.pdf)
- The SECG standrad specify 4 possible public key formats as

  | Prefix | Description                  | Length |
  | ------ | ---------------------------- | ------ |
  | 0x00   | Point at Infinity            | 1      |
  | 0x04   | Uncompressed point           | 65     |
  | 0x02   | Compressed point with even y | 33     |
  | 0x03   | Compressed point with odd y  | 33     |

  > As seen, **Ethereum only uses uncompressed public keys**
- An example goes as [Generating a Public Key](examples/generate_pubkey_test.go)

### Elliptic Curve Libraries

- [OpenSSL](https://www.openssl.org/)
- [libsecp256k1](https://github.com/bitcoin-core/secp256k1)

## Cryptographic Hash Functions

- **WHY**
  - Aid at transforming Ethereum public keys into addresses
  - Used to create digital fingerprints aiding in the verification of data
- A hash function is "any function that can be used to map data of arbi‐ trary size to data of fixed size."
  - The input is called a **pre-image**, **the message**, or simply the **input data**
  - The output is called the **hash**
- A cryptographic hash function is a **one-way** hash function that maps data of arbitrary size to a fixed-size string of bits
- **Hash Collision**: Two sets of input data that hash to the same output
- The main properties of cryptographic hash functions

|            Property | Description                                                                                                         |
| ------------------: | :------------------------------------------------------------------------------------------------------------------ |
|         Determinism | Same input, always same output                                                                                      |
|       Verifiability | Computing the hash of a message is efficient (linear complexity).                                                   |
|      Noncorrelation | A small change to the message, an extensive change in output                                                        |
|     Irreversibility | Computing the message from its hash is infeasible, equivalent to a brute-force search through all possible messages |
| Collision resistant | It should be infeasible to find hash collision. **IMPORTANT FOR AVOIDING SIGNATURE FORGERY**                        |

- Applications
  - Data fingerprinting
  - Message integrity (error detection)
  - Proof of work
  - Authentication (password hashing and key stretching)
  - Pseudorandom number generators
  - Message commitment (commit–reveal mechanisms)
  - Unique identifiers

### Ethereum's Cryptographic Hash Function: Keccak-256

- **WHY NOT THE STANDARD**
  - During the period when Ethereum was developed, the NIST standardizawas not yet finalized
  - Edward Snowden revealed documents that imply that NIST may have been improperly influenced by the National Security Agency to intentionally weaken the `Dual_EC_DRBG` random-number generator standard, effectively placing a backdoor in the standard random number generator

> **Many if not all of "SHA-3" mentioned throughout Ethereum documents and code actually refer to Keccak-256, not the finalized FIPS-202 SHA-3 standard**

### Which Hash Function Am I Using?

- Distinguish the standard SHA3 and Keccak-256 by checking the hash of the empty input

```go
Keccak256("") = c5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470
SHA3("") = a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a
```

## Ethereum Addresses

- **DEFINITION**: Unique identifier that are derived from **public keys** or **contracts** using the Keccak-256 one-way hash function

> It is worth noting that the public key is not formatted with the prefix (hex) 04 when the address is calculated

### Ethereum Address Formats

- Ethereum addresses are hexadecimal numbers, identifiers derived from **the last 20 bytes of the Keccak-256 hash** of the public key or contract

  - [An example deriving an address from public key](examples/address_test.go)

- Ethereum addresses are presented as raw hexadecimal **without any checksum**
  - **WHY**
    - Ethereum addresses would eventually be hidden behind abstractions (such as name services) at higher layers of the system
    - Checksums should be added at higher layers if necessary

### Inter Exchange Client Address Protocol (ICAP)

- **WHAT**: An Ethereum address encoding that is partly compatible with the International Bank Account Number (IBAN) encoding, offering a versatile, checksummed, and interoperable encoding for Ethereum addresses
- IBAN is an international standard for identifying bank account numbers, mostly used for wire transfers
  - **FORMAT**: A string of up to **34** alphanumeric characters (case-insensitive) comprising
    - A country code
    - A checksum
    - A country-specific bank account identifier
- ICAP uses the same structure by introducing a nonstandard country code, "`XE`," (aliasing Ethereum) followed by a **two-character checksum** and three possible variations of an account identifier

  - **Direct**
    - A big-endian base-**36** integer comprised of up to 30 alphanumeric characters, representing the **155** LSB of an Ethereum address
    - Disadvantage: since **155<160**, it **only works for Ethereum addresses that start with one or more zero bytes**
    - Advantage: compatible with IBAN, in terms of the field length and checksum
    - Example: `XE60HAMICDXSV5QXVJA7TJW47Q9CHWKJD`
  - **Basic**
    - Same as the Direct encoding, except that it is 31 characters long
    - Advantage: can encode any Ethereum address
    - Disadvantage: incompatible with IBAN field validation
    - Example: `XE18CHDJBPLTBCJ03FE9O2NS0BPOJVQCU2P`
  - **Indirect**
    - Encodes an identifier that resolves to an Ethereum address through a name regisprovider
    - It uses 16 alphanumeric characters, comprising
      - An asset identier (e.g., `ETH`)
      - A name service (e.g., `XREG`)
      - A 9-character human-readable name (e.g., `KITTYCATS`)
    - Example: `XE##ETHXREGKITTYCATS`, whre `##` should be replaced with checksum

- TODO: add an example for the direct and basic variations

> At this time, ICAP is unfortunately only supported by a few wallets

### Hex Encoding with Checksum in Capitalization (EIP-55)

- [EIP-55](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md) offers a backward-compatible checksum for Ethereum addresses by modifying the capitalization of the hexadecimal address
  - Wallets that do not support EIP-55 checksums simply ignore the fact that the address contains mixed capitalization
  - Those supporting it can validate it and detect errors with a 99.986% accuracy
- **HOW** (TODO: code)
  1. Keccak-256 hash the lowercase address, without the `0x` prefix
  2. Capitalize each alphabetic address character if the corresponding hex digit of the hash is greater than or equal to `0x8`
  - Only the first 20 bytes (40 hex characters) of the hash is employed as a checksum due the address length
- Detecting an error in an EIP-55 encoded address (TODO: code)

## Conclusions
