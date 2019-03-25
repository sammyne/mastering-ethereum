package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rpc"
	ethereum "github.com/sammyne/2018-mastering-ethereum"
)

func main() {
	c, err := rpc.Dial(ethereum.INFURA)
	if nil != err {
		panic(err)
	}
	defer c.Close()

	var gasPrice string
	if err := c.Call(&gasPrice, "eth_gasPrice"); nil != err {
		panic(err)
	}

	fmt.Println(gasPrice)
}
