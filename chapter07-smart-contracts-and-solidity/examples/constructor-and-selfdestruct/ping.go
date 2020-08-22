// +build ignore

package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	flag "github.com/spf13/pflag"
)

var (
	ganacheURL string
	txHash     string
)

// @TODO: this example code has been broken
func main() {
	flag.Parse()

	eth, err := ethclient.Dial(ganacheURL)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

	// fetch the receipt to decode the contract address
	receipt, err := eth.TransactionReceipt(context.TODO(), common.HexToHash(txHash))
	if nil != err {
		panic(err)
	}

	c, err := rpc.Dial(ganacheURL)
	if err != nil {
		panic(err)
	}

	pos := common.HexToHash("0x00")
	var result string
	err = c.CallContext(context.TODO(), &result, "eth_getStorageAt", receipt.ContractAddress, pos,
		"latest")
	if err != nil {
		panic(err)
	}
	fmt.Println("result:", result)
}

func init() {
	flag.StringVarP(&ganacheURL, "ganache-url", "g", "http://127.0.0.1:7545", "receiver's address")
	flag.StringVar(&txHash, "tx", "", "hash of tx deploying the called contract")
}

func newKeyStore(key []byte, dir string) (*keystore.KeyStore, error) {
	const passphrase = "hello-world"

	privKey, err := crypto.ToECDSA(key)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal private key: %w", err)
	}

	store := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := store.ImportECDSA(privKey, passphrase)
	if err != nil {
		return nil, fmt.Errorf("fail to import private key: %w", err)
	}

	if err := store.Unlock(account, passphrase); err != nil {
		return nil, fmt.Errorf("fail to unlock account: %w", err)
	}

	return store, nil
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}
