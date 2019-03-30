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
		nonce    = 0
		gasLimit = 2000000
	)

	var (
		ropstenChainID = big.NewInt(3)
		faucetCode, _  = hex.DecodeString(`608060405234801561001057600080fd5b50600080546001600160a01b0319163317905560f5806100316000396000f3fe60806040526004361060265760003560e01c80632e1a7d4d14602857806383197ef014604e575b005b348015603357600080fd5b50602660048036036020811015604857600080fd5b50356060565b348015605957600080fd5b50602660a5565b68056bc75e2d63100000811115607557600080fd5b604051339082156108fc029083906000818181858888f1935050505015801560a1573d6000803e3d6000fd5b5050565b6000546001600160a01b0316331460bb57600080fd5b6000546001600160a01b0316fffea165627a7a72305820408e4857ac94d3ea89fa681706c0b3768df6408d06af6f72593947a1c5a70eaf0029`)
	)

	// TODO: compare with NewTransaction
	tx := types.NewContractCreation(nonce, eth.ToWei(0), gasLimit,
		gasPrice, faucetCode)

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
