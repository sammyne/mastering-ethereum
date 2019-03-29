package main_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func depositLogging() {}

func TestDepositLogging(t *testing.T) {
	txHash := common.HexToHash("0x9cfe8a38fb9eef9ed55b65effe3572acb81c5e0d78ef7ef2097c7b6d58d71c87")

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

	logs := receipt.Logs
	expect := struct {
		method []byte
		payer  string
		amount []byte
	}{
		crypto.Keccak256([]byte("Deposit(address,uint256)")),
		"0x141a6AA1FA91F88442CF6C57f6b08b5EA63609bf",
		abi.U256(eth.ToWei(0.01)),
	}

	if got := logs[0].Topics[0].Bytes(); !bytes.Equal(got, expect.method) {
		t.Fatalf("invalid method sig: got %x, expect %x", got, expect.method)
	}

	if got := common.BytesToAddress(
		logs[0].Topics[1].Bytes()).Hex(); got != expect.payer {
		t.Fatalf("invalid payer address: got %s, expect %s", got, expect.payer)
	}

	if !bytes.Equal(logs[0].Data, expect.amount) {
		t.Fatalf("invalid amount: got %x, expect %x", logs[0].Data, expect.amount)
	}
}
