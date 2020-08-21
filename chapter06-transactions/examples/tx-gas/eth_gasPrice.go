package main

import (
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	ganacheURL := "http://127.0.0.1:7545"
	if len(os.Args) > 2 {
		ganacheURL = os.Args[1]
	}

	fmt.Printf("Ganache at %s is used\n", ganacheURL)

	c, err := rpc.Dial(ganacheURL)
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
