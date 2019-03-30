package eth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
)

// Fund fill the address with some Ropsten ethers
func Fund(address common.Address) (string, error) {
	const ropsten = "https://faucet.ropsten.be/donate/"

	resp, err := http.Get(ropsten + address.Hex())
	if nil != err {
		return "", err
	}
	defer resp.Body.Close()

	if http.StatusOK != resp.StatusCode {
		return "", errors.New(resp.Status)
	}

	response := make(map[string]interface{})
	if err := json.NewDecoder(resp.Body).Decode(&response); nil != err {
		return "", err
	}

	fmt.Println(response)

	return response["txhash"].(string), nil
}
