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
		nonce    = 5
		gasLimit = 2000000
	)

	var (
		ropstenChainID = big.NewInt(3)
		code, _        = hex.DecodeString(`6080604052348015600f57600080fd5b5060a18061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c80630d8ea18014602d575b600080fd5b60336035565b005b60408051338152326020820152308183015290517fc90ed2a0abbce0ddd0cfd1fd303f15bf1a3860b32002dc1ed37873ea889488309181900360600190a156fea165627a7a723058207b3cd8ec15e555784190b0c13de1b171f1728602d59d9f1c1c898e5c8efc11680029`)
	)

	// TODO: compare with NewTransaction
	tx := types.NewContractCreation(nonce, eth.ToWei(0), gasLimit, gasPrice, code)

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
