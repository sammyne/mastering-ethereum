# Chapter 01. What is Ethereum

## Compared to Bitcoin

- Common elements
  - a P2P network
  - a BFT consensus algorithm
  - cryptographic primitives such digital signatures and hashes
  - a digital currency
- Difference
  - **Ether** serves both as a digital currency and a utility currency to pay for maintaining the Ethereum platform as the world computer
  - Ethereum's language is **Turing complete**, but not Bitcoin's

## Components of a Blockchain

- **A peer-to-peer (P2P) network** connecting participants and propagating transactions and blocks of verified transactions, based on a standardized "gossip" protocol
- Messages, in the form of **transactions**, representing state transitions
- A set of **econsensus rules**, governing what constitutes a transaction and what makes for a valid state transition
- **A state machine** that processes transactions according to the consensus rules
- **A chain of cryptographically secured blocks** that acts as a journal of all the verified and accepted state transitions
- **A consensus algorithm** that decentralizes control over the blockchain, by forcing participants to cooperate in the enforcement of the consensus rules
- A game-theoretically sound **incentivization scheme** (e.g., proof-of-work costs plus block rewards) to economically secure the state machine in an open environment
- One or more open source software implementations of the above (**"clients"**)

## The Birth of Ethereum

- **WHY**: people recognized the power of the Bitcoin model, and were trying to move beyond cryptocurrency applications
- **WHY NOT BITCOIN**: Building upon Bitcoin meant living within the intentional constraints of the network and trying to find workarounds
- In December 2013, **Vitalik** started sharing a whitepaper that outlined the idea behind Ethereum: a Turing-complete, general-purpose blockchain
- Starting in December 2013, **Vitalik** and **Gavin** refined and evolved the idea, together building the protocol layer that became Ethereum
- Ethereum's founders were thinking about a blockchain without a specific purpose, that could support a broad variety of applications by being programmed
- And on July 30, 2015, the first Ethereum block was mined

## Ethereum's Four Stages of Development

- Ethereum's development was planned over four distinct stages, with major changes occurring at each stage
- A stage may include subreleases, known as **"hard forks,"** that change functionality in a way that is not backward compatible
- 4 Stages are codenamed as
  - Frontier
  - Homestead
  - Metropolis
  - Serenity
- Timeline

|     Block | Releases                                                                                                                                                                                                                         |
| --------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
|         0 | **Frontier** -- The initial stage of Ethereum, lasting from July 30, 2015, to March 2016                                                                                                                                         |
|   200,000 | **Ice Age** -- A hard fork to introduce an exponential difficulty increase, to motivate a transition to PoS when ready                                                                                                           |
| 1,150,000 | **Homestead** -- The second stage of Ethereum, launched in March 2016                                                                                                                                                            |
| 1,192,000 | **DAO** -- A hard fork that reimbursed victims of the hacked DAO contract and caused Ethereum and Ethereum Classic to split into two competing systems                                                                           |
| 2,463,000 | **Tangerine Whistle** -- A hard fork to change the gas calculation for certain I/O-heavy operations and to clear the accumulated state from a denial-of-service (DoS) attack that exploited the low gas cost of those operations |
| 2,675,000 | **Spurious Dragon** -- A hard fork to address more DoS attack vectors, and another state clearing. Also, a replay attack protection mechanism                                                                                    |
| 4,370,000 | **Metropolis Byzantium** -- Metropolis is the third stage of Ethereum, current at the time of writing this book, launched in October 2017. Byzantium is the first of two hard forks planned for Metropolis.                      |

## Ethereum: A General-Purpose Blockchain

- Consensue model for bitcoin as a distributed consensus state machine, where
  - Transactions cause a global state transition, altering the ownership of coins
  - The state transitions are constrained by the rules of consensus, allowing all participants to (eventually) converge on a common (consensus) state of the system, after several blocks are mined
- Ethereum is also a distributed state machine
  - Ethereum tracks the state transitions of a general-purpose data store, i.e., a store that can hold any data expressible as a key–value tuple

## Ethereum's Components

- P2P network
- Consensus rules
- Transactions
- State machine
- Data structures
- Consensus algorithm: PoW
- Economic security: moving from PoW to PoS
- Clients

### Further Reading

- [The Ethereum Yellow Paper](https://ethereum.github.io/yellowpaper/paper.pdf)
- [The Beige Paper, a rewrite of the Yellow Paper for a broader audience in less formal language](https://github.com/chronaeon/beigepaper)
- [ÐΞVp2p network protocol](http://bit.ly/2quAlTE)
- [Ethereum Virtual Machine list of resources](http://bit.ly/2PmtjiS)
- [LevelDB database (used most often to store the local copy of the blockchain)](http://leveldb.org)
- [Merkle Patricia trees](https://github.com/ethereum/wiki/wiki/Patricia-Tree)
- [Ethash PoW algorithm](https://github.com/ethereum/wiki/wiki/Ethash)
- [Casper PoS v1 Implementation Guide](http://bit.ly/2DyPr3l)
- [Go-Ethereum (Geth) client](https://geth.ethereum.org/)
- [Parity Ethereum client](https://parity.io/)

## Ethereum and Turing Completeness

- The **halting problem**
  - **DEFINITION**: Whether it is possible, given an arbitrary program and its input, to determine whether the program will eventually stop running
  - Turing proved that it is not solvable
- Alan Turing defined a system to be Turing complete if it can be used to simulate any Turing machine. Such a system is called a **Universal Turing machine (UTM)**
- Ethereum's ability to execute a stored program, in a state machine called the Ethereum Virtual Machine, while reading and writing data to memory makes it a Turingcomplete system and therefore a UTM
- Ethereum can compute any algorithm that can be computed by any Turing machine, given the limitations of finite memory
- Ethereum's groundbreaking innovation is to combine the general-purpose computing architecture of a stored-program computer with a decentralized blockchain, thereby creating a distributed single-state (singleton) world computer

### Turing Completeness as a "Feature"

- Turing completeness is very dangerous, particularly in open access systems like public blockchains, because of the halting problem
- The fact that Ethereum is Turing complete means that any program of any complexity can be computed by Ethereum

### Implications of Turing Completeness

- Turing proved that you cannot predict whether a program will terminate by simulating it on a computer. In simple terms, we cannot predict the path of a program without running it
- Unintended never-ending loops can arise without warning, due to complex interactions between the starting conditions and the code
- Whether by accident or on purpose, a smart contract can be created such that it runs forever when a node attempts to validate it. This is effectively a DoS attack
- To thwart DoS attacks, Ethereum introduces a metering mechanism called **gas**, to allow Turing-complete computation while limiting the resources that any pro‐ gram can consume
- **Way to get gas**: Ether needs to be sent along with a transaction and it needs to be explicitly earmarked for the purchase of gas, along with an acceptable gas price

## From General-Purpose Blockchains to Decentralized Applications (DApps)

- **DEFINITION**: A DApp is a web application that is built on top of open, decentralized, peer-to-peer infrastructure services
- A DApp is composed of at least:
  - Smart contracts on a blockchain
  - A web frontend user interface
- Other decentralized components
  - A decentralized (P2P) storage protocol and platform
  - A decentralized (P2P) messaging protocol and platform

## The Third Age of the Internet

- In 2004 the term **"Web 2.0"** came to prominence, describing an evolution of the web toward user-generated content, responsive interfaces, and interactivity
- **web3** represents a new vision and focus for web applications: from centrally owned and managed applications, to applications built on decentralized protocols

## Ethereum's Development Culture

- In Bitcoin, development is guided by conservative principles: all changes are carefully studied to ensure that none of the existing systems are disrupted
- Ethereum's development culture is characterized by rapid innovation, rapid evolution, and a willingness to deploy forward-looking improvements, even if this is at the expense of some backward compatibility
- One of the big challenges facing developers in Ethereum is the inherent contradiction between deploying code to an immutable system and a development platform that is still evolving

## Why Learn Ethereum?

- Blockchains have a very steep learning curve, which is made a lot less steep by Ethereum
- Ethereum is a great platform for learning about blockchains and it's building a massive community of developers, faster than any other blockchain platform

## What This Book Will Teach You

- Every component of Ethereum
