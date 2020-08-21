// +build ignore

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	flag "github.com/spf13/pflag"
)

var (
	chainID    int64
	ganacheURL string
	txHash     string
)

func main() {
	flag.Parse()

	eth, err := ethclient.Dial(ganacheURL)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

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

func init() {
	flag.Int64VarP(&chainID, "chain", "c", 5777, "ID of chain bootstraped by Ganache")
	flag.StringVarP(&ganacheURL, "ganache-url", "g", "http://127.0.0.1:7545", "receiver's address")
	flag.StringVar(&txHash, "txhash", "", "tx hash")
}
