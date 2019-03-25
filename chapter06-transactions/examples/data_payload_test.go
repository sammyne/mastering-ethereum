package main_test

import (
	"bytes"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/crypto"
)

func pad32(x []byte) []byte {
	if n := 32 - len(x); n > 0 {
		x = append(bytes.Repeat([]byte{0}, n), x...)
	}

	return x
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}

func dataPayload() {}

func Example_dataPayload() {
	hash := crypto.Keccak256([]byte("withdraw(uint256)"))
	fmt.Println(hexutil.Encode(hash))

	amt := toWei(0.01).Bytes()
	fmt.Println(hexutil.Encode(amt))
	amt = pad32(amt)

	selector := append(hash[:4], amt...)
	fmt.Println(hexutil.Encode(selector))

	// Output:
	// 0x2e1a7d4d13322e7b96f9a57413e1525c250fb7a9021cf91d1540d5b69f16a49f
	// 0x2386f26fc10000
	// 0x2e1a7d4d000000000000000000000000000000000000000000000000002386f26fc10000
}
