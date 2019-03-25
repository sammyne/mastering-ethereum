// +build ignore

package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	store := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := store.NewAccount("@hello-infora")
	if nil != err {
		panic(err)
	}

	fmt.Println(account.Address.Hex())
}
