# Mastering Ethereum

## Ethereum Client

- The Geth client in use is the official docker container [ethereum/client-go](https://hub.docker.com/r/ethereum/client-go) tagged at `alltools-v1.8.23`
- The installation goes as

  ```bash
  docker pull ethereum/client-go:alltools-v1.8.23
  ```

- Basic usage

  ```bash
  docker run --rm ethereum/client-go:alltools-v1.8.23 [executable] [options]
  ```

  - where the avaiable `executable` is listed as the official [Executables](https://github.com/ethereum/go-ethereum#executables) section

## Work in Progress

- [x] Chapter 01. What Is Ethereum?
- [x] Chapter 02. Ethereum Basics
- [x] Chapter 03. Ethereum Clients
- [ ] Chapter 04. Cryptography
  - Missing the ICAP demo
- [x] Chapter 05. Wallets
- [x] Chapter 06. Transactions
  - Maybe js version for demo later
- [ ] Chapter 07. Smart Contracts and Solidity
- [ ] Chapter 09. Smart Contract Security
- [ ] Chapter 10. Tokens
