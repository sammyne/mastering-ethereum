package main_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"strconv"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func transfer() {}

func assertReceipt2(t *testing.T, logs []*types.Log) {
	// check against event Transfer(address indexed from, address indexed to, uint256 value) in IERC20

	_, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	expect := struct {
		EventSig string
		From     string
		To       string
		Value    string
	}{
		hex.EncodeToString(crypto.Keccak256([]byte("Transfer(address,address,uint256)"))),
		accounts[0].Address.Hex(),
		accounts[1].Address.Hex(),
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
		log.Topics[1].Bytes()).Hex(); got != expect.From {
		t.Fatalf("invalid from: got %s, expect %s", got, expect.From)
	}

	if got := common.BytesToAddress(
		log.Topics[2].Bytes()).Hex(); got != expect.To {
		t.Fatalf("invalid to: got %s, expect %s", got, expect.To)
	}

	if got := len(log.Data); got != 32 {
		t.Fatalf("invalid length of data: got %d, expect 32", got)
	}

	if got := new(big.Int).SetBytes(log.Data).String(); got != expect.Value {
		t.Fatalf("invalid value: got %s, expect %s", got, expect.Value)
	}
}

func assertBalance(t *testing.T, contract common.Address,
	c *ethclient.Client, who common.Address, expect string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
	defer cancel()

	gasPrice, err := c.SuggestGasPrice(ctx)
	if nil != err {
		t.Fatal(err)
	}

	data := crypto.Keccak256([]byte("balanceOf(address)"))[:4]
	data = append(data, common.LeftPadBytes(who.Bytes(), 32)...)

	msg := ethereum.CallMsg{
		From:     who,
		To:       &contract,
		Gas:      2000000,
		GasPrice: gasPrice,
		Data:     data,
	}

	// block height should be adjusted accordingly
	blockNumber := big.NewInt(5311385)
	response, err := c.CallContract(ctx, msg, blockNumber)
	if nil != err {
		t.Fatal(err)
	}

	if got := new(big.Int).SetBytes(response).String(); got != expect {
		t.Fatalf("invalid balance of %s: got %s, expect %s", who.Hex(), got, expect)
	}
}

func Test_transfer(t *testing.T) {
	txHash := common.HexToHash("0x186367fee060851ccf1e113aa9f05e6064844fdfa07bd61cdd78c11d4c4d1d62")

	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
	defer cancel()

	receipt, err := c.TransactionReceipt(ctx, txHash)
	if nil != err {
		t.Fatal(err)
	}

	// tx executes the transfering
	assertReceipt2(t, receipt.Logs)

	_, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	// the receipt.ContractAddress is nil
	contract := common.HexToAddress("0x13b3D3e67a963Ef5Ad94d9C47dEbd503Ff04dfE6")

	assertBalance(t, contract, c, accounts[0].Address,
		strconv.Itoa(2100000000-1000))
	assertBalance(t, contract, c, accounts[1].Address, "1000")
}
