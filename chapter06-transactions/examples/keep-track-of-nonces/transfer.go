package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// INFURA don't miss the *https://* prefix
const (
	INFURA   = "https://ropsten.infura.io/v3/f3df74d615a74774821985274dedcc9e"
	gasLimit = 2000000
)

var (
	account accounts.Account
	store   *keystore.KeyStore

	ropstenChainID         = big.NewInt(3)
	throttledTestnetFaucet = common.HexToAddress("0x3b873a919aa0512d5a0f09e6dcceaa4a6727fafe")
)

func init() {
	store = keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	accounts := store.Accounts()
	if 0 == len(accounts) {
		panic("please create an account")
	}
	//for _, a := range accounts {
	//	fmt.Println(a.Address.Hex())
	//}
	account = accounts[0]
	if err := store.Unlock(account, "@hello-infora"); nil != err {
		panic(err)
	}
}

func batchSendTx(nonce uint64, to common.Address, amt *big.Int) {
	c, err := rpc.Dial(INFURA)
	if nil != err {
		panic(err)
	}
	defer c.Close()

	gasPrice := gasPrice(c)

	var txs [3]string
	for i := range txs {
		tx := types.NewTransaction(nonce+uint64(i), to, amt, gasLimit,
			gasPrice, nil)

		var err error
		tx, err = store.SignTx(account, tx, ropstenChainID)
		if nil != err {
			panic(err)
		}

		wireTx, err := rlp.EncodeToBytes(tx)
		if nil != err {
			panic(err)
		}

		txs[i] = hexutil.Encode(wireTx)
	}

	for _, tx := range txs {
		var txHash string
		if err := c.Call(&txHash, "eth_sendRawTransaction", tx); nil != err {
			panic(err)
		}

		fmt.Println("hash", txHash)

		var nTx string
		if err := c.Call(&nTx, "eth_getTransactionCount", account.Address.Hex(), "pending"); nil != err {
			panic(err)
		}

		fmt.Println("#(tx) =", nTx)
	}
}

func gasPrice(c *rpc.Client) *big.Int {
	var price string
	if err := c.Call(&price, "eth_gasPrice"); nil != err {
		panic(err)
	}

	return hexutil.MustDecodeBig(price)
}

func sendTx(nonce uint64, to common.Address, amt *big.Int) {
	c, err := rpc.Dial(INFURA)
	if nil != err {
		panic(err)
	}
	defer c.Close()

	tx := types.NewTransaction(nonce, to, amt, gasLimit, gasPrice(c), nil)

	tx, err = store.SignTx(account, tx, ropstenChainID)
	if nil != err {
		panic(err)
	}

	wireTx, err := rlp.EncodeToBytes(tx)
	if nil != err {
		panic(err)
	}

	txHex := hexutil.Encode(wireTx)

	var txHash string
	if err := c.Call(&txHash, "eth_sendRawTransaction", txHex); nil != err {
		panic(err)
	}

	fmt.Println("from", account.Address.Hex())
	fmt.Println("hash", txHash)

	var nTx string
	if err := c.Call(&nTx, "eth_getTransactionCount", account.Address.Hex(), "pending"); nil != err {
		panic(err)
	}

	fmt.Println("#(tx) =", nTx)
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}

func main() {
	// introduce a nonce gap would produce the same effect as demo in the book

	//sendTx(12, throttledTestnetFaucet, toWei(0.3))
	batchSendTx(7, throttledTestnetFaucet, toWei(0.1))
}
