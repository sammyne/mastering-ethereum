package main_test

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func deploy() {}

func assertReceipt(t *testing.T, logs []*types.Log) {
	// check against event Transfer(address indexed from, address indexed to, uint256 value) in IERC20

	expect := struct {
		EventSig string
		From     string
		To       string
		Value    *big.Int
	}{
		hex.EncodeToString(crypto.Keccak256([]byte("Transfer(address,address,uint256)"))),
		strings.Repeat("0", 64),
		"0xc2C16b5A7BFD3E13A99eB42aB0d03e19F66c6Fd2",
		big.NewInt(2100000000),
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

	if got := log.Topics[1].Hex(); got[2:] != expect.From {
		t.Fatalf("invalid from: got %s, expect %s", got, expect.From)
	}

	if got := common.BytesToAddress(log.Topics[2].Bytes()).Hex(); got != expect.To {
		t.Fatalf("invalid to: got %s, expect %s", got, expect.To)
	}

	if got := len(log.Data); got != 32 {
		t.Fatalf("invalid length of data: got %d, expect 32", got)
	}

	if got := new(big.Int).SetBytes(log.Data); nil == got ||
		got.String() != expect.Value.String() {
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

func Test_deploy(t *testing.T) {
	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	txHash := common.HexToHash("0xb837d2e2df48e79e82150fc30dea38b9ace853e000323d4d1c75eb65e465ca34")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
	defer cancel()

	receipt, err := c.TransactionReceipt(ctx, txHash)
	if nil != err {
		t.Fatal(err)
	}

	fmt.Println(receipt.ContractAddress.Hex())

	assertReceipt(t, receipt.Logs)
	assertOwnerBalance(t, receipt.ContractAddress, c)
}
