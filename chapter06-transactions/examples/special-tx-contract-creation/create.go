// +build ignore

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	flag "github.com/spf13/pflag"
)

var (
	chainID    int64
	privKey    []byte
	ganacheURL string
	nonce      uint64
)

func main() {
	flag.Parse()

	eth, err := ethclient.Dial(ganacheURL)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

	gasPrice, err := eth.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	fmt.Println("gasPrice =", gasPrice)

	const gasLimit = 2000000

	var (
		amount        = big.NewInt(0)
		chainID       = big.NewInt(chainID)
		faucetCode, _ = hexutil.Decode(`608060405234801561001057600080fd5b5060cc8061001f6000396000f3fe608060405260043610601f5760003560e01c80632e1a7d4d14602a576025565b36602557005b600080fd5b348015603557600080fd5b50605060048036036020811015604a57600080fd5b50356052565b005b67016345785d8a0000811115606657600080fd5b604051339082156108fc029083906000818181858888f193505050501580156092573d6000803e3d6000fd5b505056fea2646970667358221220e61ab153361b198b141e914f883bfd3fad37fb3d4ee2c526845f1530700da08e64736f6c63430007000033`)
	)

	// TODO: compare with NewTransaction
	tx := types.NewContractCreation(nonce, amount, gasLimit, gasPrice, faucetCode)

	workingDir, err := ioutil.TempDir("", "mastering-eth-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(workingDir)

	store, err := newKeyStore(privKey, workingDir)
	if err != nil {
		panic(fmt.Sprintf("fail to new key store: %v", err))
	}
	sender := store.Accounts()[0]

	if tx, err = store.SignTx(sender, tx, chainID); nil != err {
		panic(err)
	}

	if err := eth.SendTransaction(context.TODO(), tx); nil != err {
		panic(err)
	}

	fmt.Println("tx hash:", tx.Hash().Hex())
}

func init() {
	flag.Int64VarP(&chainID, "chain", "c", 5777, "ID of chain bootstraped by Ganache")
	flag.StringVarP(&ganacheURL, "ganache-url", "g", "http://127.0.0.1:7545", "receiver's address")
	flag.BytesHexVarP(&privKey, "key", "k", nil, "sender's key")
	flag.Uint64Var(&nonce, "nonce", 0, "nonce of tx")
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
