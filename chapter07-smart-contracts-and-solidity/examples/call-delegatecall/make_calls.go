// +build ignore

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	flag "github.com/spf13/pflag"
)

var (
	chainID      int64
	privKey      []byte
	ganacheURL   string
	nonce        uint64
	calleeTxHash string
	callerTxHash string
)

func main() {
	flag.Parse()

	eth, err := ethclient.Dial(ganacheURL)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

	/* construct the tx meta data */
	to, err := readContractAddress(eth, common.HexToHash(callerTxHash))
	if err != nil {
		panic(fmt.Sprintf("fail to read caller contract address: %v", err))
	}
	amount := toWei(0)
	const gasLimit = 2000000

	gasPrice, err := eth.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	methodSig := crypto.Keccak256([]byte("makeCalls(address)"))[:4]
	//fmt.Printf("%x\n", methodID)

	calleeAddress, err := readContractAddress(eth, common.HexToHash(calleeTxHash))
	if err != nil {
		panic(fmt.Sprintf("fail to read callee contract address: %v", err))
	}
	// the left padding is a MUST
	methodSig = append(methodSig, common.LeftPadBytes(calleeAddress.Bytes(), 32)...)
	/* end construct the tx meta data */
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, methodSig)

	chainID := big.NewInt(chainID)

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

	fmt.Println("      gasPrice =", gasPrice)
	fmt.Println("       account =", sender.Address.Hex())
	fmt.Println("calledContract =", calleeAddress.Hex())
	fmt.Println("        txHash =", tx.Hash().Hex())
}

func init() {
	flag.Int64VarP(&chainID, "chain", "c", 5777, "ID of chain bootstraped by Ganache")
	flag.StringVarP(&ganacheURL, "ganache-url", "g", "http://127.0.0.1:7545", "receiver's address")
	flag.BytesHexVarP(&privKey, "key", "k", nil, "sender's key")
	flag.Uint64Var(&nonce, "nonce", 0, "nonce of tx")
	flag.StringVar(&calleeTxHash, "callee-tx", "", "hash of tx deploying the calledContract")
	flag.StringVar(&callerTxHash, "caller-tx", "", "hash of tx deploying the caller contract")
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
