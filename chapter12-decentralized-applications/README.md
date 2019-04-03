# Chapter 12. Decentralized Applications (DApps)

## Auction DApp: Backend Smart Contracts

1. Change to `auction-dapp/backend` directory and compile the contracts

```bash
./solc --bin --optimize AuctionRepository.sol -o build
```

2. Deploy both the `AuctionRepository` and `DeedRepository` contract by running [deploy.go](auction-dapp/backend/deploy.go)
