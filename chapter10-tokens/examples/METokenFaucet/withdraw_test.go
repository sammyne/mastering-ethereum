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

func withdraw() {}

func assertWithdrawTransfer(t *testing.T, log *types.Log) {
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
		hex.EncodeToString(crypto.Keccak256([]byte("Transfer(address,address,uint256)"))),
		accounts[0].Address.Hex(),
		accounts[1].Address.Hex(),
		"1000",
	}

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

func assertWithdrawApproval(t *testing.T, log *types.Log) {
	// check against event Approval(address indexed from, address indexed to, uint256 value) in IERC20

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
		"0", // no amount remains
	}

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

func assertBalances(t *testing.T, contract common.Address,
	c *ethclient.Client) {
	_, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		t.Fatal(err)
	}
	payer, payee := accounts[0].Address, accounts[1].Address

	// in my case, the withdrawal tx is mined on block 5311866
	before := big.NewInt(5311865)
	after := new(big.Int).Add(before, big.NewInt(1))

	a1 := balanceOf(t, c, contract, payer, before)
	b1 := balanceOf(t, c, contract, payee, before)

	a2 := balanceOf(t, c, contract, payer, after)
	b2 := balanceOf(t, c, contract, payee, after)

	delta := new(big.Int).Add(a1, a2.Neg(a2))
	if "0" == delta.String() {
		t.Fatalf("should be non-zero delta")
	}

	// decrease in payer should be equal to the increase in payee
	if expect := new(big.Int).Add(b1, delta); 0 != expect.Cmp(b2) {
		t.Fatalf("invalid balance for payee: got %s, expect %s", b2, expect)
	}
}

func balanceOf(t *testing.T, c *ethclient.Client,
	token, who common.Address, block *big.Int) *big.Int {
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
		To:       &token,
		Gas:      2000000,
		GasPrice: gasPrice,
		Data:     data,
	}

	// in my case, the withdrawal tx is mined on block 5311866
	response, err := c.CallContract(ctx, msg, block)
	if nil != err {
		t.Fatal(err)
	}

	return new(big.Int).SetBytes(response)
}

func Test_withdraw(t *testing.T) {
	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	txHash := common.HexToHash("0xa10b150f5a658ea2948210c2b23de924875d09b1f5464dc1bab913c4a9000959")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
	defer cancel()

	receipt, err := c.TransactionReceipt(ctx, txHash)
	if nil != err {
		t.Fatal(err)
	}

	if 2 != len(receipt.Logs) {
		t.Fatalf("invalid #(logs): got %d, expect 2", len(receipt.Logs))
	}

	assertWithdrawTransfer(t, receipt.Logs[0])
	assertWithdrawApproval(t, receipt.Logs[1])

	// the address of the deployed METoken
	token := common.HexToAddress("0x13b3D3e67a963Ef5Ad94d9C47dEbd503Ff04dfE6")
	assertBalances(t, token, c)
}
