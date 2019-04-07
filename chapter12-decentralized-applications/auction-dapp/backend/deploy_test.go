package main_test

import (
	"context"
	"fmt"
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

	names := []string{"AuctionRepository", "DeedRepository"}

	txHashes := []common.Hash{
		common.HexToHash("0xb0a8c365a55163874d9237f8862aa3142086c1f829682d7dbd401c7fe31c1e5d"),
		common.HexToHash("0x4c6bc5621b0c82507531bc4180b1fb5d2f784e12ea2e2af32dcbe41650c9fad6"),
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
	defer cancel()

	for i, txHash := range txHashes {
		receipt, err := c.TransactionReceipt(ctx, txHash)
		if nil != err {
			t.Fatal(err)
		}
		fmt.Printf("[%17s]: %s\n", names[i], receipt.ContractAddress.Hex())
	}

}
