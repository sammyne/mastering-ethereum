# Chapter 11. Oracles

- In the context of blockchains, an oracle is a system that can answer questions that are external to Ethereum
- Ideally oracles are systems that are trustless, meaning that they do not need to be trusted because they operate on decentralized principles.

## Why Oracles Are Needed

- **WHY**: In order to maintain consensus, EVM execution must be totally deterministic and based only on the shared context of the Ethereum state and signed transactions, which indicates
  - The first is that there can be no intrinsic source of randomness for the EVM and smart contracts to work with
  - Extrinsic data can only be introduced as the data payload of a transaction
- Pseudorandom functions, such as cryptographically secure hash functions are not enough for many applications
  - E.g. gambling is vulnerable to miners
- Extrinsic information, including sources of randomness, price information, weather forecasts, etc., can be introduced as the data part of transactions sent to the network
  - Such data simply cannot be trusted, because it comes from unverifiable sources

## Oracle Use Cases and Examples

- Use cases
  - Oracles, ideally, provide a trustless (or at least near-trustless) way of getting extrinsic (i.e., "rea-world" or off-chain) information, such as the results of football games, the price of gold, or truly random numbers, onto the Ethereum platform for smart contracts to use
  - They can also be used to relay data securely to DApp frontends directly
- Example: "smart will" contract that distributes assets when a person dies
  - If the inheritance amount controlled by such a contract is high enough, the incentive to hack the oracle and trigger distribution of assets before the owner dies is very high
- Some oracles provide data that is particular to a specific private data source, such as academic certificates or government IDs
  - The truth of the data is subjective
- More examples
  - Random numbers/entropy from physical sources such as quantum/thermal processes: e.g., to fairly select a winner in a lottery smart contract
  - Parametric triggers indexed to natural hazards: e.g., triggering of catastrophe bond smart contracts, such as Richter scale measurements for an earthquake bond
  - Exchange rate data: e.g., for accurate pegging of cryptocurrencies to fiat currency
  - Capital markets data: e.g., pricing baskets of tokenized assets/securities
  - Benchmark reference data: e.g., incorporating interest rates into smart financial derivatives
  - Static/pseudostatic data: security identifiers, country codes, currency codes, etc.
  - Time and interval data: for event triggers grounded in precise time measurements
  - Weather data: e.g., insurance premium calculations based on weather forecasts
  - Political events: for prediction market resolution
  - Sporting events: for prediction market resolution and fantasy sports contracts
  - Geolocation data: e.g., as used in supply chain tracking
  - Damage verification: for insurance contracts
  - Events occurring on other blockchains: interoperability functions
  - Ether market price: e.g., for fiat gas price oracles
  - Flight statistics: e.g., as used by groups and clubs for flight ticket pooling

## Oracle Design Patterns

- Basic abilities
  - Collect data from an off-chain source
  - Transfer the data on-chain with a signed message
  - Make the data available by putting it in a smart contract's storage
- 3 patterns
  - Request–response
  - Publish-subscribe
  - Immediate-read

### Immediate-read

- Immediate-read oracles are those that provide data that is only needed for an immediate decision
  - Example queries
    - What is the address for ethereum‐book.info?
    - Is this person over 18?
- An example oracle
  - Hold data about or issued by organizations, such as academic certificates, dial codes, institutional memberships, airport identifiers, self-sovereign IDs
  - Place/Update the data in the contract storage for public access
  - The data stored by the oracle is likely not to be the raw data that the oracle is serving, e.g., for efficiency or privacy reasons
  - A hash of the raw is sufficient. Hashing the data (more carefully, in Merkle trees with salts) and only storing the root hash in the smart contract's storage would be an efficient way to organize such a service

### Publish-subscribe

- **HOW**: An oracle that effectively provides a broadcast service for data that is expected to change (perhaps both regularly and frequently) is either polled by a smart contract on-chain, or watched by an off-chain daemon for updates
- Examples: price feeds, weather information, economic or social statistics, traffic data, etc
- Polling for data changes is a local call to a synced client
- Ethereum event logs make it particularly easy for applications to look out for oracle updates

### Request-response

- **WHY**: The data space is too huge to be stored in a smart contract and users are expected to only need a small part of the overall dataset at a time
- Application: Data provider businesses
- Implementation: A system of on-chain smart contracts and off-chain infrastructure used to monitor requests and retrieve and return data
- **HOW** (in the perspective of the oracle smart contract)
  1. Receive a query from a DApp triggered by a tx from an EOA
  2. Parse the query.
  3. Check that payment and data access permissions are provided.
  4. Retrieve relevant data from an off-chain source (and encrypt it if necessary).
  5. Sign the transaction(s) with the data included.
  6. Broadcast the transaction(s) to the network.
  7. Schedule any further necessary transactions, such as notifications, etc.
- Other schemes
  - Data can be requested from and returned directly by an EOA, removing the need for an oracle smart contract

### Summary

- The request–response pattern described here is commonly seen in client–server architectures
- Publish–subscribe is a pattern where **publishers** (in this context, oracles) do not send messages directly to receivers, but instead categorize published messages into distinct classes. **Subscribers** are able to express an interest in one or more classes and retrieve only those messages that are of interest
- A broadcast pattern is appropriate where the oracle does not need to know the identity of the subscribing contract

## Data Authentication

- **WHY**: Achieve trust under the circumstance that the oracle and the request–response mechanism may be operated by distinct entities
- **HOW**: Off-chain methods are able to attest to the returned data's integrity. Two common approaches
  - Authenticity proofs
  - Trusted execution environments (TEEs)

### Authenticity proofs

- Authenticity proofs are cryptographic guarantees that data has not been tampered with
- Application: The TLSNotatory proofs of [Provable](https://provable.xyz/) (known as Oraclize before)
  - TLSNotary proofs allow a client to provide evidence to a third party that HTTPS web traffic occurred between the client and a server
  - Although it offers higher assurances against data tampering than a pure request–response mechanism, this approach does require the assumption that Amazon itself will not tamper with the VM instance

### TEE

- Application: [Town Crier](http://www.town-crier.org/) is an authenticated data feed oracle system based on the TEE approach
  - Such methods utilize hardware-based secure enclaves (Intel's Software Guard eXtensions (SGX)) to ensure data integrity
  - SGX provides
    - Guarantees of integrity
    - Confidentiality, ensuring that an application's state is opaque to other processes when running within the enclave
    - Attestation, by generating a digitally signed proof that an application—securely identified by a hash of its build—is actually running within an enclave

## Computation Oracles

- **WHAT**: Computation oracles can be used to perform computation on a set of inputs and return a calculated result that may have been infeasible to calculate on-chain
- [Provable](https://provable.xyz/) (known as Oraclize before) is centralized
- The concept of a cryptlet as a standard for verifiable oracle truths has been formalized as part of Microsoft's wider ESC Framework
- [TrueBit](https://truebit.io/) are more decentralized
  - Offers a solution for scalable and verifiable off-chain computation
  - As a computation market, allows decentralized applications to pay for verifiable computation to be performed outside of the network, but relying on Ethereum to enforce the rules of the verification game

## Decentralized Oracles

- Application: [ChainLink](https://www.smartcontract.com/link)

  - Contracts

    |       Contract | Description                                                                                                                                                                                             |
    | -------------: | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
    |     Reputation | Keep track of data provider's performance, and populate the corresponding scores in the **off-chain registry**                                                                                          |
    | Order-matching | Selects bids from oracles using the reputation contract, and finalize a service-level agreement, which includes query parameters and the number of oracles required                                     |
    |    Aggregation | Collects responses (submitted using a commit–reveal scheme) from multiple oracles, calculates the final collective result of the query, and finally feeds the results back into the reputation contract |

  - One major challenge is the formulation of the aggregation function tackling responst weighting
  - A related idea is the SchellingCoin protocol, where multiple participants report values and the median is taken as the "correct" answer

## Oracle Client Interfaces in Solidity

- A solidity example demonstrating how Oraclize can be used to continuously poll for the ETH/USD price from an API and store the result in a usable manner
  - Sample as [EthUsdPriceTicker.sol](examples/contracts/EthUsdPriceTicker.sol)
  - The data request is made using the `oraclize_query` function, which is inherited from the `usingOraclize` contract
  - In order to perform the query, Oraclize requires the payment of a small fee in ether, covering the gas cost for processing the result and transmitting it to the `__callback` function and an accompanying surcharge for the service
- Financial data provider Thomson Reuters also provides an oracle service for Ethereum, called BlockOne IQ, allowing market and reference data to be requested by smart contracts running on private or permissioned networks

## Conclusions
