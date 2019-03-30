// +build ignore

package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func main() {
	txHash := common.HexToHash("0x639972a46ed3bd36bf06077a8e83de844e19b3f9b9388e3f64a4550c1afb53c5")

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
	amount := eth.ToWei(0.01)
	const gasLimit = 2000000

	gasPrice, err := c.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}
	/* end construct the tx meta data */

	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, nil)

	store, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	//account := accounts[1]
	account := accounts[1]
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