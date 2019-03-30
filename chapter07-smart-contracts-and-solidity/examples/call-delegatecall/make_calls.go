// +build ignore

package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func loadContractAddress(c *ethclient.Client,
	txHash common.Hash) common.Address {
	// fetch the receipt to decode the contract address
	receipt, err := c.TransactionReceipt(context.TODO(), txHash)
	if nil != err {
		panic(err)
	}

	return receipt.ContractAddress
}

func main() {
	txHash := common.HexToHash("0x0dfbd9d7992f433fff03f3e6326b8f6b72e8e29e6ba258d1fd7ae0e1813e3956")

	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	/* construct the tx meta data */
	const nonce = 22 // this should adapt to the specified account
	to := loadContractAddress(c, txHash)
	amount := eth.ToWei(0)
	const gasLimit = 2000000

	gasPrice, err := c.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	// destroy() is method sig
	//methodID := crypto.Keccak256([]byte("makeCalls(address)"))[:4]
	methodID := crypto.Keccak256([]byte("makeCalls(address)"))[:4]
	//fmt.Printf("%x\n", methodID)

	deployCalledContractTxHash := common.HexToHash("0x6ea482e07a361bc21cf4081b9e2bb232f5b18be156dd9f587e27b0e57641c722")
	// the left padding is a MUST
	calledContract := common.LeftPadBytes(
		loadContractAddress(c, deployCalledContractTxHash).Bytes(), 32)
	methodID = append(methodID, calledContract...)
	/* end construct the tx meta data */
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, methodID)

	store, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	account := accounts[0]
	ropstenChainID := big.NewInt(3)
	if tx, err = store.SignTx(account, tx, ropstenChainID); nil != err {
		panic(err)
	}

	if err := c.SendTransaction(context.TODO(), tx); nil != err {
		panic(err)
	}

	fmt.Println("      gasPrice =", gasPrice)
	fmt.Println("       account =", account.Address.Hex())
	//fmt.Println("calledContract =", calledContract.Hex())
	fmt.Println("        txHash =", tx.Hash().Hex())
}
