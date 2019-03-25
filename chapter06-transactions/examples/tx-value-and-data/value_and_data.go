package main

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	ethereum "github.com/sammyne/2018-mastering-ethereum"
)

const (
	amt      = 0.01
	gasLimit = 2000000
	nonce    = 0
)

var (
	account accounts.Account
	store   *keystore.KeyStore

	ropstenChainID = big.NewInt(3)
	to             = common.HexToAddress("0x3b873a919aa0512d5a0f09e6dcceaa4a6727fafe")
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

func gasPrice(c *rpc.Client) *big.Int {
	var price string
	if err := c.Call(&price, "eth_gasPrice"); nil != err {
		panic(err)
	}

	return hexutil.MustDecodeBig(price)
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}

func exampleDataOnly() {
	fmt.Println("--- data only ---")

	c, err := rpc.Dial(ethereum.INFURA)
	if nil != err {
		panic(err)
	}
	defer c.Close()

	data, _ := hexutil.Decode("0x1234")

	tx := types.NewTransaction(nonce+2, to, toWei(0), gasLimit, gasPrice(c), data)

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

	fmt.Println("--- END data only ---")
}

func exampleNone() {
	fmt.Println("--- no value or data only ---")

	c, err := rpc.Dial(ethereum.INFURA)
	if nil != err {
		panic(err)
	}
	defer c.Close()

	tx := types.NewTransaction(nonce+3, to, toWei(0), gasLimit, gasPrice(c), nil)

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

	fmt.Println("--- END no value or data ---")
}

func exampleValueAndData() {
	fmt.Println("--- value and data ---")

	c, err := rpc.Dial(ethereum.INFURA)
	if nil != err {
		panic(err)
	}
	defer c.Close()

	data, _ := hexutil.Decode("0x1234")

	tx := types.NewTransaction(nonce+1, to, toWei(amt), gasLimit, gasPrice(c), data)

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

	fmt.Println("--- END value and data ---")
}

func exampleValueOnly() {
	fmt.Println("--- value only ---")

	c, err := rpc.Dial(ethereum.INFURA)
	if nil != err {
		panic(err)
	}
	defer c.Close()

	tx := types.NewTransaction(nonce, to, toWei(amt), gasLimit, gasPrice(c), nil)

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

	fmt.Println("--- END value only ---")
}

func main() {
	//exampleValueOnly()
	//exampleValueAndData()
	//exampleDataOnly()
	exampleNone()
}
