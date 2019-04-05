package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/compiler"
)

func compileSolidity(names []string, sources ...string) (
	[][]byte, error) {
	contracts, err := compiler.CompileSolidity("solc", sources...)
	if nil != err {
		panic(err)
	}

	codes := make([][]byte, 0, len(names))
	for _, name := range names {

		var ok bool
		for path, v := range contracts {
			if strings.Contains(path, name) {
				fmt.Println(name, v.Code)
				code, err := hexutil.Decode(v.Code)
				if nil != err {
					return nil, err
				}

				codes, ok = append(codes, code), true
				break
			}
		}
		if !ok {
			return nil, errors.New("missing code for " + name)
		}
	}

	return codes, nil
}

func packDeedArgs() ([]byte, error) {
	strType, err := abi.NewType("string", nil)
	if nil != err {
		return nil, err
	}

	args := abi.Arguments{
		abi.Argument{Name: "_name", Type: strType},
		abi.Argument{Name: "_symbol", Type: strType},
	}

	data, err := args.Pack("Ultra Auction NFT", "UANFT")
	if nil != err {
		return nil, err
	}

	return data, nil
}

func main() {
	const contract = "AuctionRepository"
	codes, err := compileSolidity(
		[]string{contract, "DeedRepository"},
		filepath.Join("contracts", contract+".sol"))
	if nil != err {
		panic(err)
	}

	hello := codes[0]
	bin, err := ioutil.ReadFile("AuctionRepository.bin")
	if nil != err {
		panic(err)
	}

	world, err := hex.DecodeString(string(bin))
	if nil != err {
		panic(err)
	}
	//fmt.Println(hex.EncodeToString(world))

	//fmt.Println(len(hello), len(world))
	for i := range hello {
		if hello[i] != world[i] {
			x := hex.EncodeToString(hello[i : i+8])
			y := hex.EncodeToString(world[i : i+8])
			fmt.Println(i, x, y)
			break
		}
	}

	/*
		c, err := eth.Dial()
		if nil != err {
			panic(err)
		}
		defer c.Close()

		gasPrice, err := c.SuggestGasPrice(context.TODO())
		if nil != err {
			panic(err)
		}

		const (
			nonce    = 0
			gasLimit = 2000000
		)

		var ropstenChainID = big.NewInt(3)

		store, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
			eth.DefaultPassphrase())
		if nil != err {
			panic(err)
		}
		account := accounts[0]

		deploy := func(contract string, bytecode []byte, i uint64) {
			defer fmt.Println("--- done deploying contract:", contract, " ---")

			// TODO: compare with NewTransaction
			tx := types.NewContractCreation(nonce+i, eth.ToWei(0), gasLimit,
				gasPrice, bytecode)

			if tx, err = store.SignTx(account, tx, ropstenChainID); nil != err {
				panic(err)
			}

			if err := c.SendTransaction(context.TODO(), tx); nil != err {
				panic(err)
			}

			fmt.Println("--- deploying contract:", contract, " ---")
			fmt.Println("gasPrice =", gasPrice)
			fmt.Println(" account =", account.Address.Hex())
			fmt.Println("  txHash =", tx.Hash().Hex())
		}

		deedArgs, err := packDeedArgs()
		if nil != err {
			panic(err)
		}

		contracts := []struct {
			name     string
			bytecode []byte
		}{
			{"AuctionRepository", codes[0]},
			{"DeedRepository", append(codes[1], deedArgs...)},
		}

		for i, c := range contracts {
			deploy(c.name, c.bytecode, uint64(i))
		}
	*/
}
