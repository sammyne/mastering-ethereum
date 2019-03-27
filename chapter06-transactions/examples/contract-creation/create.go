// +build ignore

package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
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
		nonce    = 1
		gasLimit = 2000000
	)

	var (
		amount         = big.NewInt(0)
		ropstenChainID = big.NewInt(3)
		faucetCode, _  = hexutil.Decode(`0x608060405234801561001057600080fd5b5060b08061001f6000396000f3fe608060405260043610601c5760003560e01c80632e1a7d4d14601e575b005b348015602957600080fd5b50601c60048036036020811015603e57600080fd5b503567016345785d8a0000811115605457600080fd5b604051339082156108fc029083906000818181858888f193505050501580156080573d6000803e3d6000fd5b505056fea165627a7a72305820d77165d5b972de3b3604c31e2dfd8cf2dd0c901b623521d1444453fbb12d52160029`)
	)

	// TODO: compare with NewTransaction
	tx := types.NewContractCreation(nonce, amount, gasLimit, gasPrice, faucetCode)

	store, account := loadKeyStore()
	if tx, err = store.SignTx(account, tx, ropstenChainID); nil != err {
		panic(err)
	}

	if err := eth.SendTransaction(context.TODO(), tx); nil != err {
		panic(err)
	}

	//fmt.Println(tx.Hash())
}
