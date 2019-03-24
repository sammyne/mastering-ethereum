package main_test

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func encodeEIP55() {}

func Example_encodeEIP55() {
	address := common.HexToAddress("001d3f1ef827552ae1114027bd3ecf1f086ba0f9")

	//crypto.Keccak256(address)
	fmt.Println(address.Hex())

	// Output:
	// 0x001d3F1ef827552Ae1114027BD3ECF1f086bA0F9
}

func Example_encodeEIP55_Error() {
	// ok: 001d3F1ef827552Ae1114027BD3ECF1f086bA0F9
	// the character 'F' before the last one is changed to 'E'
	bad := "001d3F1ef827552Ae1114027BD3ECF1f086bA0E9"

	digest := crypto.Keccak256([]byte(strings.ToLower(bad)))

	fmt.Println(bad)
	fmt.Printf("%x\n", digest)

	// As seen, several letters have been of wrong cases

	// Output:
	// 001d3F1ef827552Ae1114027BD3ECF1f086bA0E9
	// 5429b5d9460122fb4b11af9cb88b7bb76d8928862e0a57d46dd18dd8e08a6927
}
