package main_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func generatePubKey() {}

func Example_generatePubKey() {
	curve := secp256k1.S256()

	k, _ := new(big.Int).SetString("f8f8a2f43c8376ccb0871305060d7b27b0554d2cc72bccf41b2705608452f315", 16)

	Kx, Ky := curve.ScalarBaseMult(k.Bytes())

	fmt.Println("x =", Kx.Text(16))
	fmt.Println("y =", Ky.Text(16))
	fmt.Printf("%x\n", curve.Marshal(Kx, Ky))

	// Output:
	// x = 6e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b
	// y = 83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0
	// 046e145ccef1033dea239875dd00dfb4fee6e3348b84985c92f103444683bae07b83b5c38e5e2b0c8529d7fa3f64d46daa1ece2d9ac14cab9477d042c84c32ccd0
}
