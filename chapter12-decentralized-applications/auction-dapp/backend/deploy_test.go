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
		common.HexToHash("0x4a9cac1fe13793b30f7ce269cf4c78c20d0ca96e8d93de89e50b3d2bbc3b3429"),
		common.HexToHash("0x71e99b07a0512c76252ece765c5fc05c8de95c2b70b0153f8836b5cab2b98433"),
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
