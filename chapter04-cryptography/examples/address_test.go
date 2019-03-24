package main_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func pubKey2Address() {}

func Example_pubKey2Address() {
	curve := secp256k1.S256()

	k, _ := new(big.Int).SetString("f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315", 16)

	K := curve.Marshal(curve.ScalarBaseMult(k.Bytes()))
	// don't forget the trim out the 0x04 prefix

	out := crypto.Keccak256(K[1:])

	fmt.Printf("%x\n", out[len(out)-20:])

	// Output
	// 001d3f1ef827552ae1114027bd3ecf1f086ba0f9
}
