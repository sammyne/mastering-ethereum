package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	// note the https:// prefix is a MUST
	const endpoint = "https://ropsten.infura.io/v3/f3df74d615a74774821985274dedcc9e"

	c, err := rpc.Dial(endpoint)
	if nil != err {
		panic(err)
	}

	var quantity string
	if err := c.Call(&quantity, "eth_gasPrice"); nil != err {
		panic(err)
	}

	fmt.Println(quantity)
}
