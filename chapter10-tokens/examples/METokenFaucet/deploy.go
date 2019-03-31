package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func readInMETokenBytecodes() ([]byte, error) {
	const src = "contracts/build/METFaucet.bin"

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
		nonce    = 26
		gasLimit = 2000000
	)

	var ropstenChainID = big.NewInt(3)
	bytecodes, err := readInMETokenBytecodes()
	if nil != err {
		panic(err)
	}
	//fmt.Printf("%s\n", bytecodes)

	store, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}
	account := accounts[0]

	METokenAddress := common.HexToAddress("0x13b3D3e67a963Ef5Ad94d9C47dEbd503Ff04dfE6")

	data := append(bytecodes, common.LeftPadBytes(METokenAddress.Bytes(), 32)...)
	data = append(data, common.LeftPadBytes(account.Address.Bytes(), 32)...)

	// TODO: compare with NewTransaction
	tx := types.NewContractCreation(nonce, eth.ToWei(0), gasLimit, gasPrice, data)

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
