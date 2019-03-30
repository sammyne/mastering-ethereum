// +build ignore

package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	ethereum "github.com/sammyne/mastering-ethereum"
)

func loadKeyStore() (*keystore.KeyStore, accounts.Account) {
	store := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	accounts := store.Accounts()
	if 0 == len(accounts) {
		panic("please create an account")
	}
	//for _, a := range accounts {
	//	fmt.Println(a.Address.Hex())
	//}
	if err := store.Unlock(accounts[0], "@hello-infora"); nil != err {
		panic(err)
	}

	return store, accounts[0]
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}

func main() {
	eth, err := ethclient.Dial(ethereum.INFURA)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

	gasPrice, err := eth.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	fmt.Println("gasPrice =", gasPrice)

	const (
		nonce         = 3
		gasLimit      = 2000000
		faucetAddress = `0x910928acd685fb3844c8f599476c1af31788bd64`
	)

	var (
		ropstenChainID = big.NewInt(3)
		// data is to invoke withdraw(0.01 ether)
		data, _ = hexutil.Decode("0x2e1a7d4d000000000000000000000000000000000000000000000000002386f26fc10000")
	)

	tx := types.NewTransaction(nonce, common.HexToAddress(faucetAddress),
		toWei(0), gasLimit, gasPrice, data)

	store, account := loadKeyStore()
	if tx, err = store.SignTx(account, tx, ropstenChainID); nil != err {
		panic(err)
	}

	if err := eth.SendTransaction(context.TODO(), tx); nil != err {
		panic(err)
	}

	fmt.Println(tx.Hash().Hex())
}
