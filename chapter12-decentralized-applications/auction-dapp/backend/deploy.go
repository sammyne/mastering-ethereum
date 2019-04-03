package main

import (
	"context"
	"fmt"
	"math/big"
	"path/filepath"

	"github.com/ethereum/go-ethereum/accounts/abi"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/sammyne/mastering-ethereum/playground/eth"
	"github.com/sammyne/mastering-ethereum/playground/io"
)

func packDeedArgs() ([]byte, error) {
	strType, err := abi.NewType("string", nil)
	if nil != err {
		return nil, err
	}

	args := abi.Arguments{
		abi.Argument{Name: "_name", Type: strType},
		abi.Argument{Name: "_symbol", Type: strType},
	}

	data, err := args.Pack("Ultra Auction NFT", "UANFT")
	if nil != err {
		return nil, err
	}

	return data, nil
}

func main() {
	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	gasPrice, err := c.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	const (
		nonce    = 30
		gasLimit = 2000000
	)

	var ropstenChainID = big.NewInt(3)

	store, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}
	account := accounts[0]

	deploy := func(contract string, args []byte, i uint64) {
		defer fmt.Println("--- done deploying contract:", contract, " ---")

		contract = filepath.Join(".", "contracts", "build", contract+".bin")
		bytecodes, err := io.LoadContractCode(contract)
		if nil != err {
			panic(err)
		}
		bytecodes = append(bytecodes, args...)

		// TODO: compare with NewTransaction
		tx := types.NewContractCreation(nonce+i, eth.ToWei(0), gasLimit,
			gasPrice, bytecodes)

		if tx, err = store.SignTx(account, tx, ropstenChainID); nil != err {
			panic(err)
		}

		if err := c.SendTransaction(context.TODO(), tx); nil != err {
			panic(err)
		}

		fmt.Println("--- deploying contract:", contract, " ---")
		fmt.Println("gasPrice =", gasPrice)
		fmt.Println(" account =", account.Address.Hex())
		fmt.Println("  txHash =", tx.Hash().Hex())
	}

	deedArgs, err := packDeedArgs()
	if nil != err {
		panic(err)
	}

	contracts := []struct {
		name string
		args []byte
	}{
		{"AuctionRepository", nil},
		{"DeedRepository", deedArgs},
	}

	for i, c := range contracts {
		deploy(c.name, c.args, uint64(i))
	}
}
