package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func main() {
	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	gasPrice, err := c.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	const (
		nonce    = 23
		gasLimit = 2000000
	)

	var (
		ropstenChainID = big.NewInt(3)
		code, _        = hex.DecodeString(`608060405234801561001057600080fd5b5060b18061001f6000396000f3fe608060405260043610601c5760003560e01c80632e1a7d4d14601e575b005b348015602957600080fd5b50601c60048036036020811015603e57600080fd5b503568056bc75e2d63100000811115605557600080fd5b604051339082156108fc029083906000818181858888f193505050501580156081573d6000803e3d6000fd5b505056fea165627a7a723058203e0fd3e2dd356cca16234c9694818592b25230f09763db6796dd47babac46aa20029`)
	)

	tx := types.NewContractCreation(nonce, eth.ToWei(0.1), gasLimit, gasPrice,
		code)

	store, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	account := accounts[0]
	if tx, err = store.SignTx(account, tx, ropstenChainID); nil != err {
		panic(err)
	}

	if err := c.SendTransaction(context.TODO(), tx); nil != err {
		panic(err)
	}

	fmt.Println("gasPrice =", gasPrice)
	fmt.Println(" account =", account.Address.Hex())
	fmt.Println("  txHash =", tx.Hash().Hex())
}
