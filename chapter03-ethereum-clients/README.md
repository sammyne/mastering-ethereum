# Chapter 03. Ethereum Clients

- An Ethereum client is a **software application** that
  - Implements the Ethereum specification
  - Communicates over the peer-to-peer network with other Ethereum clients
- Different Ethereum clients **interoperate** if they comply with the reference specification and the standardized communications protocols

## Ethereum Networks

- There exist a variety of Ethereum-based networks that largely conform to the formal specification defined in the Ethereum Yellow Paper, but which may or may not interoperate with each other
- Examples
  - Ethereum
  - Ethereum Classic
  - Ella
  - Expanse
  - Ubiq
  - Musicoin
- 6 main implementations
  - Parity, written in Rust
  - Geth, written in Go
  - cpp-ethereum, written in C++
  - pyethereum, written in Python
  - Mantis, written in Scala
  - Harmony, written in Java

### Should I Run a Full Node?

- The health, resilience, and censorship resistance of blockchains depend on them having many independently operated and geographically dispersed full nodes
- Each full node can help other new nodes obtain the block data to bootstrap their operation, as well as offering the operator an authoritative and independent verification of all transactions and contracts
- Running a full node will incur a cost in hardware resources (80-100GB as of Sep. 2018) and bandwidth
- We can do almost everything you need to do with a **testnet** node (which connects you to one of the smaller public test blockchains), with a local private blockchain like Ganache, or with a cloud-based Ethereum client offered by a service provider like Infura
- **Remote clients** are those who does not store a local copy of the blockchain or validate blocks and transactions
  - Examples: MetaMask, Emerald Wallet, MyEtherWallet, or MyCrypto
- **REMOTE CLIENTS ARE NOT LIGHT CLIENT**:
  - Light clients validate block headers and use Merkle proofs to validate the inclusion of transactions in the blockchain and determine their effects, giving them a similar level of security to a full node
  - Remote clients do not validate block headers or transactions. They entirely trust a full client to give them access to the blockchain, and hence lose significant security and anonymity guarantees

### Full Node Advantages and Disadvantages

#### Advantages

- Supports the resilience and censorship resistance of Ethereum-based networks
- Authoritatively validates all transactions
- Can interact with any contract on the public blockchain without an intermediary
- Can directly deploy contracts into the public blockchain without an intermediary
- Can query (read-only) the blockchain status (accounts, contracts, etc.) offline
- Can query the blockchain without letting a third party know the information you’re reading

#### Disadvantages

- Requires significant and growing hardware and bandwidth resources
- May require several days to fully sync when first started
- Must be maintained, upgraded, and kept online to remain synced

### Public Testnet Advantages and Disadvantages

#### Advantages

- A testnet node needs to sync and store much less data—about 10 GB depending on the network (as of April 2018).
- A testnet node can sync fully in a few hours.
- Deploying contracts or making transactions requires test ether, which has no value and can be acquired for free from several "faucets."
- Testnets are public blockchains with many other users and contracts, running "live."

#### Disadvantages

- You can't use "real" money on a testnet; it runs on test ether. Consequently, you can't test security against real adversaries, as there is nothing at stake.
- There are some aspects of a public blockchain that you cannot test realistically on a testnet. Examples are
  - Transaction fees, although necessary to send transactions, are not a consideration on a testnet, since gas is free.
  - Not so real network congestion occurring in the public mainnet sometimes

### Local Blockchain Simulation Advantages and Disadvantages

- Ganache shares many of the advantages and disadvantages of the public testnet, but also has some differences

#### Advantages

- No syncing and almost no data on disk; you mine the first block yourself
- No need to obtain test ether; you "award" yourself mining rewards that you can use for testing
- No other users, just you
- No other contracts, just the ones you deploy after you launch it

#### Disadvantages

- Having no other users means that it doesn't behave the same as a public blockchain. There's no competition for transaction space or sequencing of transactions.
- No miners other than you means that mining is more predictable; therefore, you can't test some scenarios that occur on a public blockchain.
- Having no other contracts means you have to deploy everything that you want to test, including dependencies and contract libraries.
- You can't recreate some of the public contracts and their addresses to test some scenarios (e.g., the DAO contract).

## Running an Ethereum Client

- If you have the time and resources, you should attempt to run a full node

### Hardware Requirements for a Full Node

#### Minimum requirements:

- CPU with 2+ cores
- At least 80 GB free storage space
- 4 GB RAM minimum with an SSD, 8 GB+ if you have an HDD
- 8 MBit/sec download internet service

> The Parity codebase is lighter on resources, friendlier to limited hardwares

#### Recommended specifications

- Fast CPU with 4+ cores
- 16 GB+ RAM
- Fast SSD with at least 500 GB free space
- 25+ MBit/sec download internet service

> The disk size requirements listed here assume you will be running a node with default settings, where the blockchain is "pruned" of old state data. If you instead run a full "archival" node, where all state is kept on disk, it will likely require more than 1 TB of disk space.

- Up-to-date estimates of the blockchain size
  - [Ethereum](https://bitinfocharts.com/ethereum/)
  - [Ethereum Classic](https://bitinfocharts.com/ethereum%20classic/)

### Software Requirements for Building and Running a Client (Node)

- Typically every blockchain will have its own version of Geth
- Parity provides support for multiple Ethereum-based blockchains (Ethereum, Ethereum Classic, Ellaism, Expanse, Musicoin) with the same client download
- Tools
  - [Git](https://git-scm.com/)
  - [Go](https://golang.google.cn/)

### ~~Parity~~

### Go-Ethereum (Geth)

- The "official" implementation of the Ethereum client
- Every Ethereum-based blockchain will have its own Geth implementation
  - [Ethereum](https://github.com/ethereum/go-ethereum)
  - [Ethereum Classic](https://github.com/ethereumproject/go-ethereum)
  - [Ellaism](https://github.com/ellaism/go-ellaism)
  - [Expanse](https://github.com/expanse-org/go-expanse)
  - [Musicoin](https://github.com/Musicoin/go-musicoin)
  - [Ubiq](https://github.com/ubiq/go-ubiq)

#### Cloning the repository

```bash
git clone <Repository Link>
```

#### Building Geth from source code

```bash
cd go-ethereum
make geth
```

#### Docker container is another option

- [`ethereum/client-go`](https://hub.docker.com/r/ethereum/client-go)
- Once installed (the detailed instruction goes as the **Ethereum Client** secion in the root [README](../README.md)), check the correctness by touching the version of `geth`

```bash
docker run --rm ethereum/client-go:alltools-v1.8.23 geth version
```

## The First Synchronization of Ethereum-Based Blockchains

- Many Ethereum-based blockchains were the victim of denial-of-service attacks at the end of 2016
  - Example: block 2,283,397 (September 18, 2016) to block 2,700,031 (November 26, 2016), solved by hard forks
- To speed up sync, most Ethereum clients include an option to perform a "fast" synchronization that skips the full validation of transactions until it has synced to the tip of the blockchain, then resumes full validation.
  - For Geth, the option is typically called `--fast`

### Running Geth ~~or Parity~~

- Use the `--help` option to see all the configuration parameters

### The JSON-RPC Interface

- The JSON-RPC API is an interface that allows us to write programs that use an Ethereum client as a gateway to an Ethereum network and blockchain.
- Usually, the RPC interface is offered as an HTTP service on port 8545
  - Querying the version of `geth` goes as [web3_clientVersion.go](examples/web3_clientVersion.go)
  - Querying the current gas price goes as [eth_gasPrice.go](examples/eth_gasPrice.go)
- For security reasons it is restricted, by default, to only accept connections from localhost
- Each request contains four elements

|   Element | Description                                                                                                                                                                                                                                                     |
| --------: | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `jsonrpc` | Version of the JSON-RPC protocol. This **MUST** be exactly "2.0"                                                                                                                                                                                                |
|  `method` | The name of the method to be invoked                                                                                                                                                                                                                            |
|  `params` | A **OPTIONAL** structured value that holds the parameter values to be used during the invocation of the method                                                                                                                                                  |
|      `id` | An identifier established by the client that MUST contain a `String`, `Number`, or `NULL` value if included. The server MUST reply with the same value in the response object if included. This member is used to correlate the context between the two objects |

## Remote Ethereum Clients

- They do not store the full Ethereum blockchain, so they are faster to set up and require far less data storage
- Ability
  - Manage private keys and Ethereum addresses in a wallet.
  - Create, sign, and broadcast transactions.
  - Interact with smart contracts, using the data payload.
  - Browse and interact with DApps.
  - Offer links to external services such as block explorers.
  - Convert ether units and retrieve exchange rates from external sources.
  - Inject a web3 instance into the web browser as a JavaScript object.
  - Use a web3 instance provided/injected into the browser by another client.
  - Access RPC services on a local or remote Ethereum node.

### Mobile (Smartphone) Wallets

- All mobile wallets are remote clients
- Examples
  - [Jaxx](https://jaxx.io/)
  - [Status](https://status.im/)
  - [Trust Wallet](https://trustwalletapp.com/)
  - [Cipher Browser](https://www.cipherbrowser.com/)

### Browser Wallets

- MetaMask
  - Unlike other browser wallets, MetaMask injects a web3 instance into the browser JavaScript context, acting as an RPC client that connects to a variety of Ethereum blockchains
- Jaxx
- MyEtherWallet
- MyCrypto
- Mist
  - 1st implementation of ERC20 token
  - 1st introduction of the camelCase checksum (EIP-55)
  - Runs a full node, and offers a full DApp browser with support for Swarm-based stor‐ age and ENS addresses

## Conclusions
