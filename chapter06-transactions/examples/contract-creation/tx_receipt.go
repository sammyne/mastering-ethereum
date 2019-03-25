// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	ethereum "github.com/sammyne/2018-mastering-ethereum"
)

func main() {
	eth, err := ethclient.Dial(ethereum.INFURA)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

	// adapt to your own case
	const txHash = `0x0d546865e2ca26e4442491b854dd75d61f2fa6ce0a9c4db94027b007ff78f838`

	receipt, err := eth.TransactionReceipt(context.TODO(), common.HexToHash(txHash))
	if nil != err {
		panic(err)
	}

	receiptJSON, err := json.MarshalIndent(receipt, "", " ")
	if nil != err {
		panic(err)
	}
	fmt.Printf("%s\n", receiptJSON)
}
