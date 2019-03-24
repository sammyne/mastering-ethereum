// +build ignore

package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	c, err := rpc.Dial("http://localhost:8545")
	if nil != err {
		panic(err)
	}

	var quantity string
	if err := c.Call(&quantity, "eth_gasPrice"); nil != err {
		panic(err)
	}

	fmt.Println(quantity)
}
