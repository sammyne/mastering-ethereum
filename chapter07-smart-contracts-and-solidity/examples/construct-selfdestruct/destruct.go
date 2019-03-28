// +build ignore

package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
	Usage: app accountIdx
	  Where accountIdx=0 is for owner, and otherwise nonwoner
`)
		return
	}

	accountIdx, err := strconv.Atoi(os.Args[1])
	if nil != err {
		fmt.Println(err)
		return
	}
	if accountIdx != 0 {
		accountIdx = 1
	}

	txHash := common.HexToHash("0x5dfc8c40a185dba7fd3cd1c2a09ebf87ef634bd3458a5e04f90d508f3521c380")

	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	// fetch the receipt to decode the contract address
	receipt, err := c.TransactionReceipt(context.TODO(), txHash)
	if nil != err {
		panic(err)
	}

	/* construct the tx meta data */
	const nonce = 1 // this should adapt to the specified account
	to := receipt.ContractAddress
	amount := eth.ToWei(0)
	const gasLimit = 2000000

	gasPrice, err := c.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	// destroy() is method sig
	methodID := crypto.Keccak256([]byte("destroy()"))[:4]
	/* end construct the tx meta data */
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, methodID)

	store, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	//account := accounts[1]
	account := accounts[accountIdx]
	ropstenChainID := big.NewInt(3)
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
