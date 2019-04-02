package main_test

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func deploy() {}

func Test_deploy(t *testing.T) {
	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	txHash := common.HexToHash("0x88b9bec762386023b06d95d1c25de6a7482c2852dfac8d931d72113420becbce")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
	defer cancel()

	receipt, err := c.TransactionReceipt(ctx, txHash)
	if nil != err {
		t.Fatal(err)
	}

	account := receipt.ContractAddress

	/** check token */
	token, err := c.StorageAt(context.TODO(), account,
		common.BigToHash(big.NewInt(0)), nil)
	if nil != err {
		t.Fatal(err)
	}

	tokenAddress := "0x13b3D3e67a963Ef5Ad94d9C47dEbd503Ff04dfE6"
	if got := common.BytesToAddress(token).Hex(); got != tokenAddress {
		t.Fatalf("invalid token contract address: got %s, expect %s", got,
			tokenAddress)
	}
	/** done check token */

	owner, err := c.StorageAt(context.TODO(), account,
		common.BigToHash(big.NewInt(1)), nil)
	if nil != err {
		t.Fatal(err)
	}

	_, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		t.Fatal(err)
	}

	if got, expect := common.BytesToAddress(owner).Hex(),
		accounts[0].Address.Hex(); got != expect {
		t.Fatalf("invalid owner: got %s, expect %s", got, expect)
	}
}
