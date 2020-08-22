// +build ignore

package main

import (
	"bytes"
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	flag "github.com/spf13/pflag"
)

var (
	ganacheURL string
	txHash     string
	account    string
)

func main() {
	flag.Parse()

	eth, err := ethclient.Dial(ganacheURL)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

	// fetch the receipt to decode the contract address
	receipt, err := eth.TransactionReceipt(context.TODO(), common.HexToHash(txHash))
	if nil != err {
		panic(err)
	}

	logs := receipt.Logs
	expect := struct {
		method []byte
		payer  string
		amount []byte
	}{
		crypto.Keccak256([]byte("Withdrawal(address,uint256)")),
		account,
		math.PaddedBigBytes(toWei(0.01), 32),
	}

	if got := logs[0].Topics[0].Bytes(); !bytes.Equal(got, expect.method) {
		panic(fmt.Sprintf("invalid method sig: got %x, expect %x", got, expect.method))
	}

	if got := common.BytesToAddress(
		logs[0].Topics[1].Bytes()).Hex(); got != expect.payer {
		panic(fmt.Sprintf("invalid payer address: got %s, expect %s", got, expect.payer))
	}

	if !bytes.Equal(logs[0].Data, expect.amount) {
		panic(fmt.Sprintf("invalid amount: got %x, expect %x", logs[0].Data, expect.amount))
	}
}

func init() {
	flag.StringVarP(&ganacheURL, "ganache-url", "g", "http://127.0.0.1:7545", "receiver's address")
	flag.StringVar(&txHash, "tx", "", "hash of tx deploying the called contract")
	flag.StringVarP(&account, "account", "a", "", "the expected 'from' field within the event")
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}
