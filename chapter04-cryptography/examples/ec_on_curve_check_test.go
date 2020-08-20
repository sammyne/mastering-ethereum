package main_test

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func onCurveChecking() {}

// ExampleOnCurveChecking implements Example 4-1. Using Python to confim that this point is on the elliptic curve
func Example_onCurveChecking() {
	p := secp256k1.S256().P

	x, _ := new(big.Int).SetString("49790390825249384486033144355916864607616083520101638681403973749255924539515", 10)
	y, _ := new(big.Int).SetString("59574132161899900045862086493921015780032175291755807399284007721050341297360", 10)

	// the equation is y^2 = x^3 + 7
	x3 := new(big.Int).Mul(x, x)
	x3.Mul(x3, x)

	y2 := new(big.Int).Mul(y, y)

	z := new(big.Int).Add(x3, big.NewInt(7))
	z.Add(z, new(big.Int).Neg(y2))

	fmt.Println(z.Mod(z, p))

	// Output:
	// 0
}
