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

func withdrawalLogging() {}

func TestWithdrawalLogging(t *testing.T) {
	txHash := common.HexToHash("0x727f474f3143a821bcfaa4af9e06b345c5553146dc51cc08bd5de9324471b222")

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
		payee  string
		amount []byte
	}{
		crypto.Keccak256([]byte("Withdrawal(address,uint256)")),
		"0xc2C16b5A7BFD3E13A99eB42aB0d03e19F66c6Fd2",
		abi.U256(eth.ToWei(0.005)),
	}

	if got := logs[0].Topics[0].Bytes(); !bytes.Equal(got, expect.method) {
		t.Fatalf("invalid method sig: got %x, expect %x", got, expect.method)
	}

	if got := common.BytesToAddress(
		logs[0].Topics[1].Bytes()).Hex(); got != expect.payee {
		t.Fatalf("invalid payee address: got %s, expect %s", got, expect.payee)
	}

	if !bytes.Equal(logs[0].Data, expect.amount) {
		t.Fatalf("invalid amount: got %x, expect %x", logs[0].Data, expect.amount)
	}
}
