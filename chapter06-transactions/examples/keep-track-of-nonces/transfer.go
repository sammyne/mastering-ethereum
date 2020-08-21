package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	flag "github.com/spf13/pflag"
)

const gasLimit = 2000000

var (
	chainID    int64
	privKey    []byte
	to         string
	ganacheURL string
)

func main() {
	flag.Parse()

	workingDir, err := ioutil.TempDir("", "mastering-eth-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(workingDir)

	store, err := newKeyStore(privKey, workingDir)
	if err != nil {
		panic(fmt.Sprintf("fail to new key store: %v", err))
	}

	toAddr := common.HexToAddress(to)
	const nonce = 0
	fmt.Println("---")
	fmt.Println("> one sending")
	if err := sendTx(store, nonce, toAddr, toWei(8)); err != nil {
		panic(fmt.Sprintf("fail to send tx: %v", err))
	}

	fmt.Println("---")
	fmt.Println("> batch sending")
	if err := batchSendTx(store, nonce+1, toAddr, toWei(2)); err != nil {
		panic(fmt.Sprintf("fail to batch tx: %v", err))
	}
}

func batchSendTx(store *keystore.KeyStore, nonce uint64, to common.Address, amt *big.Int) error {
	c, err := rpc.Dial(ganacheURL)
	if nil != err {
		return fmt.Errorf("fail to connect Ganache: %w", err)
	}
	defer c.Close()

	gasPrice, sender, chainID := mustGetGasPrice(c), store.Accounts()[0], big.NewInt(chainID)

	var txs [3]string
	for i := range txs {
		tx := types.NewTransaction(nonce+uint64(i), to, amt, gasLimit,
			gasPrice, nil)

		var err error
		tx, err = store.SignTx(sender, tx, chainID)
		if nil != err {
			return fmt.Errorf("fail to sign %d-th tx: %w", i, err)
		}

		wireTx, err := rlp.EncodeToBytes(tx)
		if nil != err {
			return fmt.Errorf("fail to encode %d-th tx: %w", i, err)
		}

		txs[i] = hexutil.Encode(wireTx)
	}

	for i, tx := range txs {
		var txHash string
		if err := c.Call(&txHash, "eth_sendRawTransaction", tx); nil != err {
			return fmt.Errorf("fail to send %d-th raw tx: %w", i, err)
		}

		fmt.Println("hash", txHash)

		var nTx string
		if err := c.Call(&nTx, "eth_getTransactionCount", sender.Address.Hex(), "pending"); nil != err {
			return fmt.Errorf("fail to get #(pending tx): %w", err)
		}

		fmt.Println("#(tx) =", nTx)
	}

	return nil
}

func init() {
	flag.Int64VarP(&chainID, "chain", "c", 5777, "ID of chain bootstraped by Ganache")
	flag.BytesHexVarP(&privKey, "key", "k", nil, "sender's key")
	flag.StringVarP(&to, "to", "t", "", "receiver's address")
	flag.StringVarP(&ganacheURL, "ganache-url", "g", "http://127.0.0.1:7545", "receiver's address")
}

func mustGetGasPrice(c *rpc.Client) *big.Int {
	var price string
	if err := c.Call(&price, "eth_gasPrice"); nil != err {
		panic(err)
	}

	return hexutil.MustDecodeBig(price)
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

func sendTx(store *keystore.KeyStore, nonce uint64, to common.Address, amt *big.Int) error {
	c, err := rpc.Dial(ganacheURL)
	if nil != err {
		return fmt.Errorf("fail to connect to Ganache: %w", err)
	}
	defer c.Close()

	tx := types.NewTransaction(nonce, to, amt, gasLimit, mustGetGasPrice(c), nil)

	sender := store.Accounts()[0]
	tx, err = store.SignTx(sender, tx, big.NewInt(chainID))
	if nil != err {
		return fmt.Errorf("fail to sign tx: %w", err)
	}

	wireTx, err := rlp.EncodeToBytes(tx)
	if nil != err {
		return fmt.Errorf("fail to marshal tx: %w", err)
	}

	txHex := hexutil.Encode(wireTx)

	var txHash string
	if err := c.Call(&txHash, "eth_sendRawTransaction", txHex); nil != err {
		panic(err)
	}

	fmt.Println("from", sender.Address.Hex())
	fmt.Println("hash", txHash)

	var nTx string
	if err := c.Call(&nTx, "eth_getTransactionCount", sender.Address.Hex(), "pending"); nil != err {
		panic(err)
	}

	fmt.Println("#(tx) =", nTx)

	return nil
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}

/*
func main() {
	// introduce a nonce gap would produce the same effect as demo in the book

	//sendTx(12, throttledTestnetFaucet, toWei(0.3))
	batchSendTx(7, throttledTestnetFaucet, toWei(0.1))
}
*/
