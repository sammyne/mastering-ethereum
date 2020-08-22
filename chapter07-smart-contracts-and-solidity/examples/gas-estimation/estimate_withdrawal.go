// +build ignore

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	flag "github.com/spf13/pflag"
)

var (
	ganacheURL string
	nonce      uint64
	privKey    []byte
	txHash     string
)

func main() {
	flag.Parse()

	eth, err := ethclient.Dial(ganacheURL)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

	workingDir, err := ioutil.TempDir("", "mastering-eth-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(workingDir)

	store, err := newKeyStore(privKey, workingDir)
	if err != nil {
		panic(fmt.Sprintf("fail to new key store: %v", err))
	}

	/* construct the CallMsg meta data */
	from := store.Accounts()[0].Address
	to, err := readContractAddress(eth, common.HexToHash(txHash))
	if err != nil {
		panic(fmt.Sprintf("fail to read contract address: %v", err))
	}
	const gasLimit = 2000000

	gasPrice, err := eth.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	data := crypto.Keccak256([]byte("withdraw(uint256)"))[:4]
	// the contract must have enough amount for the estimation to work
	data = append(data, common.LeftPadBytes(toWei(0.01).Bytes(), 32)...)
	/* end construct the CallMsg meta data */
	callMsg := ethereum.CallMsg{from, &to, gasLimit, gasPrice, nil, data}

	gas, err := eth.EstimateGas(context.TODO(), callMsg)
	if nil != err {
		panic(err)
	}

	cost := new(big.Int).Mul(gasPrice, big.NewInt(int64(gas)))

	fmt.Println("Gas price is", gasPrice, " wei")
	fmt.Println("Gas estimation =", gas, " units")
	fmt.Println("Gas cost estimation =", cost, " wei")
}

func init() {
	flag.StringVarP(&ganacheURL, "ganache-url", "g", "http://127.0.0.1:7545", "receiver's address")
	flag.Uint64Var(&nonce, "nonce", 0, "nonce of tx")
	flag.BytesHexVarP(&privKey, "key", "k", nil, "sender's key")
	flag.StringVar(&txHash, "tx", "", "hash of tx deploying the contract")
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

func readContractAddress(c *ethclient.Client, txHash common.Hash) (common.Address, error) {
	// fetch the receipt to decode the contract address
	receipt, err := c.TransactionReceipt(context.TODO(), txHash)
	if nil != err {
		return common.Address{}, fmt.Errorf("fail to fetch tx receipt: %w", err)
	}

	return receipt.ContractAddress, nil
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}
