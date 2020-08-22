// +build ignore

package main

import (
	"bytes"
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
	privKey        []byte
	ganacheURL     string
	calleeTxHash   string
	callerTxHash   string
	makeCallTxHash string
)

func main() {
	flag.Parse()

	eth, err := ethclient.Dial(ganacheURL)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

	calleeTxHash, callerTxHash := common.HexToHash(calleeTxHash), common.HexToHash(callerTxHash)
	callee, err := readContractAddress(eth, calleeTxHash)
	if err != nil {
		panic(fmt.Sprintf("fail to read calledContract's address: %w", err))
	}
	caller, err := readContractAddress(eth, callerTxHash)
	if err != nil {
		panic(fmt.Sprintf("fail to read caller's address: %w", err))
	}

	workingDir, err := ioutil.TempDir("", "mastering-eth-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(workingDir)

	store, err := newKeyStore(privKey, workingDir)
	if err != nil {
		panic(fmt.Sprintf("fail to new key store: %v", err))
	}
	sender := store.Accounts()[0].Address
	//fmt.Println(payer.Hex())

	logs, err := statMakeCalls(eth, common.HexToHash(makeCallTxHash))
	if nil != err {
		panic(fmt.Sprintf("fail to stat makeCalls tx: %v", err))
	}

	if len(logs) != 4 {
		panic("should be exactly 4 events")
	}

	topic := crypto.Keccak256([]byte("callEvent(address,address,address)"))
	for i, l := range logs {
		if !bytes.Equal(l.Topics[0].Bytes(), topic) {
			panic(fmt.Sprintf("#%d invalid topic: got %s, expect %x", i, l.Topics[0].Hex(), topic))
		}
	}

	fmt.Println("checking the event emitted by _callContract.calledFunction()")
	assert(logs[0].Data, []common.Address{caller, sender, callee})

	fmt.Println("checking the event emitted by calledLibrary.calledFunction()")
	assert(logs[1].Data, []common.Address{sender, sender, caller})

	fmt.Println("check the event emitted by address(_calledContract).call(methodSig)")
	assert(logs[2].Data, []common.Address{caller, sender, callee})

	fmt.Println("check the event emitted by address(_calledContract).delegatecall(methodSig)")
	assert(logs[3].Data, []common.Address{sender, sender, caller})
}

func assert(data []byte, addresses []common.Address) {
	if expect := len(addresses) * 32; len(data) != expect {
		panic(fmt.Sprintf("invalid data length: got %d, expect %d", len(data), expect))
	}

	for i, j := 0, 0; j < len(addresses); i, j = i+32, j+1 {
		addr := common.BytesToAddress(data[i:(i + 32)])
		if addr != addresses[j] {
			panic(fmt.Sprintf("#%d invalid address: got %s, expect %s", j, addr.Hex(),
				addresses[j].Hex()))
		}
	}
}

func init() {
	//flag.Int64VarP(&chainID, "chain", "c", 5777, "ID of chain bootstraped by Ganache")
	flag.StringVarP(&ganacheURL, "ganache-url", "g", "http://127.0.0.1:7545", "receiver's address")
	flag.StringVar(&calleeTxHash, "callee-tx", "", "hash of tx deploying the calledContract")
	flag.StringVar(&callerTxHash, "caller-tx", "", "hash of tx deploying the caller contract")
	flag.StringVar(&makeCallTxHash, "make-calls-tx", "", "hash of tx deploying makeCalls")
	flag.BytesHexVarP(&privKey, "key", "k", nil, "sender's key")
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

func statMakeCalls(c *ethclient.Client, txHash common.Hash) ([]*types.Log, error) {
	receipt, err := c.TransactionReceipt(context.TODO(), txHash)
	if nil != err {
		return nil, err
	}

	return receipt.Logs, nil
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}
