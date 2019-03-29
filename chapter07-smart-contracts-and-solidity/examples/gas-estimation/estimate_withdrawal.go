// +build ignore

package main

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func loadContractAddress(c *ethclient.Client,
	txHash common.Hash) common.Address {
	// fetch the receipt to decode the contract address
	receipt, err := c.TransactionReceipt(context.TODO(), txHash)
	if nil != err {
		panic(err)
	}

	return receipt.ContractAddress
}

func main() {
	txHash := common.HexToHash("0xdd0ce568a3a88646feb5634f7a08f00b0292e79731399e82b8a436f0115bdd79")

	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	_, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	/* construct the CallMsg meta data */
	const nonce = 24 // this should adapt to the specified account
	from := accounts[0].Address
	to := loadContractAddress(c, txHash)
	const gasLimit = 2000000

	gasPrice, err := c.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	// no value

	data := crypto.Keccak256([]byte("withdraw(uint256)"))[:4]
	data = append(data, common.LeftPadBytes(eth.ToWei(0.01).Bytes(), 32)...)

	/* end construct the CallMsg meta data */
	callMsg := ethereum.CallMsg{from, &to, gasLimit, gasPrice, nil, data}

	gas, err := c.EstimateGas(context.TODO(), callMsg)
	if nil != err {
		panic(err)
	}

	cost := new(big.Int).Mul(gasPrice, big.NewInt(int64(gas)))

	fmt.Println("Gas price is", gasPrice, " wei")
	fmt.Println("Gas estimation =", gas, " units")
	fmt.Println("Gas cost estimation =", cost, " wei")
}
