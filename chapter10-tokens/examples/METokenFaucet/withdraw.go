// +build ignore

package main

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func main() {
	// tx deploying the METFaucet contract
	txHash := common.HexToHash("0x88b9bec762386023b06d95d1c25de6a7482c2852dfac8d931d72113420becbce")

	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*16)
	defer cancel()

	receipt, err := c.TransactionReceipt(ctx, txHash)
	if nil != err {
		panic(err)
	}

	store, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	/* tx metadata */
	const nonce = 2
	to := receipt.ContractAddress
	amount := big.NewInt(0)
	const gasLimit = 2000000

	gasPrice, err := c.SuggestGasPrice(ctx)
	if nil != err {
		panic(err)
	}

	data := crypto.Keccak256([]byte("withdraw(uint256)"))[:4]
	data = append(data, abi.U256(big.NewInt(1000))...)
	/* end tx metadata */
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)

	ropstenChainID := big.NewInt(3)
	tx, err = store.SignTx(accounts[1], tx, ropstenChainID)
	if nil != err {
		panic(err)
	}

	if err := c.SendTransaction(context.TODO(), tx); nil != err {
		panic(err)
	}

	fmt.Println("gasPrice =", gasPrice)
	fmt.Println("    from =", accounts[1].Address.Hex())
	fmt.Println("      to =", to.Hex())
	fmt.Println("  txHash =", tx.Hash().Hex())
}
