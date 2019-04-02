package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func readInMETokenBytecodes() ([]byte, error) {
	const src = "contracts/build/METoken.bin"

	data, err := ioutil.ReadFile(src)
	if nil != err {
		return nil, err
	}

	return hex.DecodeString(string(data))
}

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
		nonce    = 24
		gasLimit = 2000000
	)

	var ropstenChainID = big.NewInt(3)
	bytecodes, err := readInMETokenBytecodes()
	if nil != err {
		panic(err)
	}
	//fmt.Printf("%s\n", bytecodes)

	// TODO: compare with NewTransaction
	tx := types.NewContractCreation(nonce, eth.ToWei(0), gasLimit,
		gasPrice, bytecodes)

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
