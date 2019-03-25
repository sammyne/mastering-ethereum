// +build ignore

package main

import (
	"math/big"
	"net/rpc"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

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

func main() {

}
