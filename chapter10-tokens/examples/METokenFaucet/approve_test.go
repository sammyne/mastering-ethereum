package main_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func approve() {}

func assertReceipt(t *testing.T, logs []*types.Log) {
	// check against event Transfer(address indexed from, address indexed to, uint256 value) in IERC20

	_, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		t.Fatal(err)
	}

	expect := struct {
		EventSig string
		Owner    string
		Spender  string
		Value    string
	}{
		hex.EncodeToString(crypto.Keccak256([]byte("Approval(address,address,uint256)"))),
		accounts[0].Address.Hex(),
		// address of faucet
		"0xB76d6936f92399D8F07B12D025072D0CF2BbCC96",
		"1000",
	}

	// should be only one event emitted by _mint()
	if 1 != len(logs) {
		t.Fatalf("unexpected #(logs): got %d, expect 1", len(logs))
	}
	log := logs[0]

	if 3 != len(log.Topics) {
		t.Fatalf("invalid #(topics): got %d, expect 3", len(log.Topics))
	}

	if got := log.Topics[0].Hex(); got[2:] != expect.EventSig {
		t.Fatalf("invalid event sig: got %s, expect %s", got, expect.EventSig)
	}

	if got := common.BytesToAddress(
		log.Topics[1].Bytes()).Hex(); got != expect.Owner {
		t.Fatalf("invalid from: got %s, expect %s", got, expect.Owner)
	}

	if got := common.BytesToAddress(
		log.Topics[2].Bytes()).Hex(); got != expect.Spender {
		t.Fatalf("invalid to: got %s, expect %s", got, expect.Spender)
	}

	if got := len(log.Data); got != 32 {
		t.Fatalf("invalid length of data: got %d, expect 32", got)
	}

	if got := new(big.Int).SetBytes(log.Data).String(); got != expect.Value {
		t.Fatalf("invalid value: got %s, expect %s", got, expect.Value)
	}
}

func assertOwnerBalance(t *testing.T, contract common.Address,
	c *ethclient.Client) {
	minter := common.HexToAddress("0xc2C16b5A7BFD3E13A99eB42aB0d03e19F66c6Fd2")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
	defer cancel()

	gasPrice, err := c.SuggestGasPrice(ctx)
	if nil != err {
		t.Fatal(err)
	}

	data := crypto.Keccak256([]byte("balanceOf(address)"))[:4]
	data = append(data, common.LeftPadBytes(minter.Bytes(), 32)...)

	msg := ethereum.CallMsg{
		From:     minter,
		To:       &contract,
		Gas:      2000000,
		GasPrice: gasPrice,
		Data:     data,
	}

	blockNumber := big.NewInt(5310724)
	response, err := c.CallContract(ctx, msg, blockNumber)
	if nil != err {
		t.Fatal(err)
	}

	const expect = "2100000000"

	if got := new(big.Int).SetBytes(response).String(); got != expect {
		t.Fatalf("invalid balance of minter: got %s, expect %s", got, expect)
	}
}

func Test_approve(t *testing.T) {
	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	txHash := common.HexToHash("0x516554f0c2618135e232aea9d65574cd277478d58377a5a7cbaf8d7c4021d089")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
	defer cancel()

	receipt, err := c.TransactionReceipt(ctx, txHash)
	if nil != err {
		t.Fatal(err)
	}

	assertReceipt(t, receipt.Logs)
}
