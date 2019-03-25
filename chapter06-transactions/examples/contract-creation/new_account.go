package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	store := keystore.NewKeyStore("./keystore", keystore.StandardScryptN, keystore.StandardScryptP)

	account, err := store.NewAccount("@hello-infora")
	if nil != err {
		panic(err)
	}

	fmt.Println(account.Address.Hex())

	data := strings.NewReader(fmt.Sprintf(`{"toWhom": "%s" }`, account.Address.Hex()))

	resp, err := http.Post("https://ropsten.faucet.b9lab.com/tap", "application/json", data)
	if nil != err {
		panic(err)
	}
	defer resp.Body.Close()

	response := make(map[string]string)
	if err := json.NewDecoder(resp.Body).Decode(&response); nil != err {
		panic(err)
	}

	fmt.Println("txHash", response["txHash"])
}
