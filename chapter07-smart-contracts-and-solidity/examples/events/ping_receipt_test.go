package main_test

import (
	"context"
	"encoding/hex"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func PingReceipt() {}

func TestPingReceipt(t *testing.T) {
	_, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		panic(err)
	}

	txHash := common.HexToHash("0x639972a46ed3bd36bf06077a8e83de844e19b3f9b9388e3f64a4550c1afb53c5")

	expect := struct {
		address string
		code    string
		owner   string
	}{
		"0x443158884343fd60f5de7eE0DaEdb331D3Cd716C",
		"6080604052600436106100295760003560e01c80632e1a7d4d1461006157806383197ef01461008d575b60408051348152905133917fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c919081900360200190a2005b34801561006d57600080fd5b5061008b6004803603602081101561008457600080fd5b50356100a2565b005b34801561009957600080fd5b5061008b610153565b67016345785d8a00008111156100ec57604051600160e51b62461bcd0281526004018080602001828103825260358152602001806101dc6035913960400191505060405180910390fd5b604051339082156108fc029083906000818181858888f19350505050158015610119573d6000803e3d6000fd5b5060408051828152905133917f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65919081900360200190a250565b6000546001600160a01b0316331461019f57604051600160e51b62461bcd02815260040180806020018281038252602e8152602001806101ae602e913960400191505060405180910390fd5b6000546001600160a01b0316fffe4f6e6c792074686520636f6e7472616374206f776e65722063616e2063616c6c20746869732066756e6374696f6e496e73756666696369656e742062616c616e636520696e2066617563657420666f72207769746864726177616c2072657175657374a165627a7a72305820abaf16ce06faf1ff11d2bc61d3728f94655eca1972e696bb29a69a97cf5850cc0029",
		accounts[0].Address.Hex(),
	}

	c, err := eth.Dial()
	if nil != err {
		panic(err)
	}
	defer c.Close()

	// fetch the receipt to decode the contract address
	receipt, err := c.TransactionReceipt(context.TODO(), txHash)
	if nil != err {
		panic(err)
	}

	code, err := c.CodeAt(context.TODO(), receipt.ContractAddress, nil)
	if nil != err {
		panic(err)
	}

	pos := common.HexToHash("0x00")
	owner, err := c.StorageAt(context.TODO(), receipt.ContractAddress, pos, nil)
	if nil != err {
		panic(err)
	}

	if got := receipt.ContractAddress.Hex(); got != expect.address {
		t.Fatalf("invalid address: got %s, expect %s", got, expect.address)
	}

	if got := hex.EncodeToString(code); got != expect.code {
		t.Fatalf("invalid code: got %s, expect %s", got, expect.code)
	}

	if got := common.BytesToAddress(owner).Hex(); got != expect.owner {
		t.Fatalf("invalid owner: got %s, expect %s", got, expect.owner)
	}
}
