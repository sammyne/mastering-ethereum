// +build ignore

package main

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func recoverPrivKey() *ecdsa.PrivateKey {
	D, _ := new(big.Int).SetString("91c8360c4cb4b5fac45513a7213f31d4e4a7bfcb4630e9fbf074f42a203ac0b9", 16)

	curve := secp256k1.S256()

	X, Y := curve.ScalarBaseMult(D.Bytes())

	return &ecdsa.PrivateKey{
		D: D,
		PublicKey: ecdsa.PublicKey{
			Curve: curve,
			X:     X,
			Y:     Y,
		},
	}
}

func main() {
	var (
		nonce       uint64 = 0
		gasPrice, _        = new(big.Int).SetString("09184e72a000", 16)
		gasLimit    uint64 = 0x30000
		to                 = common.HexToAddress("0xb0920c523d582040f2bcb1bd7fb1c7c1ecebdb34")
		value              = big.NewInt(0)
	)

	tx := types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil)

	mainnetChainID := big.NewInt(1)

	priv := recoverPrivKey()

	/*
		msg := sha256.Sum256([]byte("hello world"))
		r, s, err := ecdsa.Sign(rand.Reader, priv, msg[:])
		if nil != err {
			panic(err)
		}

		if !ecdsa.Verify(&priv.PublicKey, msg[:], r, s) {
			fmt.Println("failed to verify")
		}
	*/

	tx, err := types.SignTx(tx, types.NewEIP155Signer(mainnetChainID), priv)
	if nil != err {
		panic(err)
	}

	encodedTx, err := rlp.EncodeToBytes(tx)
	if nil != err {
		panic(err)
	}

	// NOTICE: the hash isn't the same with that on book due to addon after signing
	fmt.Println("Tx Hash:", tx.Hash().Hex())
	fmt.Printf("Signed Raw Transaction: %x\n", encodedTx)
}
