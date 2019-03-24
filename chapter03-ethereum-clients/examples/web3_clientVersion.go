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

	var version string
	if err := c.Call(&version, "web3_clientVersion"); nil != err {
		panic(err)
	}

	fmt.Println(version)
}
