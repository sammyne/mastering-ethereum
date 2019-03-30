package main_test

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func PingCode() {}

func TestPingCode(t *testing.T) {
	_, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	txHash := common.HexToHash("0x5dfc8c40a185dba7fd3cd1c2a09ebf87ef634bd3458a5e04f90d508f3521c380")

	expect := struct {
		address string
		code    string
		owner   string
	}{
		"0x64447Cf53aBd2ff422b8b51c547fBcECb8c3E49F",
		"60806040526004361060265760003560e01c80632e1a7d4d14602857806383197ef014604e575b005b348015603357600080fd5b50602660048036036020811015604857600080fd5b50356060565b348015605957600080fd5b50602660a5565b68056bc75e2d63100000811115607557600080fd5b604051339082156108fc029083906000818181858888f1935050505015801560a1573d6000803e3d6000fd5b5050565b6000546001600160a01b0316331460bb57600080fd5b6000546001600160a01b0316fffea165627a7a72305820408e4857ac94d3ea89fa681706c0b3768df6408d06af6f72593947a1c5a70eaf0029",
		accounts[0].Address.Hex(),
	}

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

	code, err := c.CodeAt(context.TODO(), receipt.ContractAddress, nil)
	if nil != err {
		panic(err)
	}

	pos := common.HexToHash("0x00")
	owner, err := c.StorageAt(context.TODO(), receipt.ContractAddress, pos, nil)
	if nil != err {
		panic(err)
	}

	if got := receipt.ContractAddress.Hex(); got != expect.address {
		t.Fatalf("invalid address: got %s, expect %s", got, expect.address)
	}

	if got := hex.EncodeToString(code); got != expect.code {
		t.Fatalf("invalid code: got %s, expect %s", got, expect.code)
	}

	if got := common.BytesToAddress(owner).Hex(); got != expect.owner {
		t.Fatalf("invalid owner: got %s, expect %s", got, expect.owner)
	}
}
