// +build ignore

package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	keyStorePath = "./keystore"
	pwd          = "@hello-infora"
)

func loadKeyStore() (*keystore.KeyStore, accounts.Account) {
	store := keystore.NewKeyStore(keyStorePath, keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := store.NewAccount(pwd)
	if nil != err {
		panic(err)
	}

	if err := store.Unlock(account, pwd); nil != err {
		panic(err)
	}

	return store, account
}

func main() {
	faucetCode, _ := hexutil.Decode(`0x608060405234801561001057600080fd5b50600080546001600160a01b0319163317905560f5806100316000396000f3fe60806040526004361060265760003560e01c80632e1a7d4d14602857806383197ef014604e575b005b348015603357600080fd5b50602660048036036020811015604857600080fd5b50356060565b348015605957600080fd5b50602660a5565b68056bc75e2d63100000811115607557600080fd5b604051339082156108fc029083906000818181858888f1935050505015801560a1573d6000803e3d6000fd5b5050565b6000546001600160a01b0316331460bb57600080fd5b6000546001600160a01b0316fffea165627a7a72305820408e4857ac94d3ea89fa681706c0b3768df6408d06af6f72593947a1c5a70eaf0029`)
	fmt.Printf("%x\n", faucetCode)

	/*
		eth, err := ethclient.Dial(ethereum.INFURA)
		if nil != err {
			panic(err)
		}
		defer eth.Close()

		//_, account := loadKeyStore()

		address := common.HexToAddress("0x0392334A1Ee195851d884b05c74aAa094cb23ab8")

		code, err := eth.CodeAt(context.TODO(), address, nil)
		if nil != err {
			panic(err)
		}

		fmt.Printf("%x\n", code)
	*/
}
