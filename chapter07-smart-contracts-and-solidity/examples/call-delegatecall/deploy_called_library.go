// +build ignore

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
		nonce    = 6
		gasLimit = 2000000
	)

	var (
		ropstenChainID = big.NewInt(3)
		code, _        = hex.DecodeString(`60b8610024600b82828239805160001a607314601757fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361060335760003560e01c80630d8ea180146038575b600080fd5b818015604357600080fd5b50604a604c565b005b60408051338152326020820152308183015290517fc90ed2a0abbce0ddd0cfd1fd303f15bf1a3860b32002dc1ed37873ea889488309181900360600190a156fea165627a7a72305820704e1b75a6f88411b8eaf38470b06a064a44fc085e8ae097fb08d4a6ed8130cb0029`)
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
