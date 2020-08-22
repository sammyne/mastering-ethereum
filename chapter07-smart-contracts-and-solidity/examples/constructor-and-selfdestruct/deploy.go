package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	flag "github.com/spf13/pflag"
)

var (
	chainID    int64
	privKey    []byte
	ganacheURL string
	nonce      uint64
)

func main() {
	flag.Parse()

	faucetCode, err := generateBinaryABI("Faucet03.sol")
	if err != nil {
		panic(err)
	}
	fmt.Printf("faucet code: %x\n", faucetCode)

	eth, err := ethclient.Dial(ganacheURL)
	if nil != err {
		panic(err)
	}
	defer eth.Close()

	gasPrice, err := eth.SuggestGasPrice(context.TODO())
	if nil != err {
		panic(err)
	}

	chainID := big.NewInt(chainID)
	const gasLimit = 2000000

	workingDir, err := ioutil.TempDir("", "mastering-eth-*")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(workingDir)

	store, err := newKeyStore(privKey, workingDir)
	if err != nil {
		panic(fmt.Sprintf("fail to new key store: %v", err))
	}
	sender := store.Accounts()[0]

	// TODO: compare with NewTransaction
	tx := types.NewContractCreation(nonce, toWei(0), gasLimit, gasPrice, faucetCode)
	if tx, err = store.SignTx(sender, tx, chainID); nil != err {
		panic(err)
	}

	if err := eth.SendTransaction(context.TODO(), tx); nil != err {
		panic(err)
	}

	fmt.Println("gasPrice =", gasPrice)
	fmt.Println(" account =", sender.Address.Hex())
	fmt.Println("  txHash =", tx.Hash().Hex())
}

func generateBinaryABI(contract string) ([]byte, error) {
	workingDir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("fail to get current working dir: %w", err)
	}

	cmdLine := fmt.Sprintf("docker run --rm -v %s/contracts:/contracts --workdir /contracts ethereum/solc:0.7.0 --bin --optimize %s", workingDir, contract)

	cmdAndArgs := strings.Split(cmdLine, " ")
	cmd := exec.Command(cmdAndArgs[0], cmdAndArgs[1:]...)

	var stdout strings.Builder
	cmd.Stdout, cmd.Stderr = &stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("fail to run cmd: %w", err)
	}

	// output should be like
	//
	// ======= Faucet03.sol:Faucet =======
	// Binary:
	// [the actual abi binary code]
	//
	raw := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	//fmt.Println(raw[2])
	out, err := hex.DecodeString(raw[2])
	if err != nil {
		return nil, fmt.Errorf("fail to decode ABI code: %w", err)
	}

	return out, nil
}

func init() {
	flag.Int64VarP(&chainID, "chain", "c", 5777, "ID of chain bootstraped by Ganache")
	flag.StringVarP(&ganacheURL, "ganache-url", "g", "http://127.0.0.1:7545", "receiver's address")
	flag.BytesHexVarP(&privKey, "key", "k", nil, "sender's key")
	flag.Uint64Var(&nonce, "nonce", 0, "nonce of tx")
}

func newKeyStore(key []byte, dir string) (*keystore.KeyStore, error) {
	const passphrase = "hello-world"

	privKey, err := crypto.ToECDSA(key)
	if err != nil {
		return nil, fmt.Errorf("fail to unmarshal private key: %w", err)
	}

	store := keystore.NewKeyStore(dir, keystore.StandardScryptN, keystore.StandardScryptP)
	account, err := store.ImportECDSA(privKey, passphrase)
	if err != nil {
		return nil, fmt.Errorf("fail to import private key: %w", err)
	}

	if err := store.Unlock(account, passphrase); err != nil {
		return nil, fmt.Errorf("fail to unlock account: %w", err)
	}

	return store, nil
}

func toWei(ethers float64) *big.Int {
	// 1 ether = 10^18 wei
	orders, _ := new(big.Float).SetString("1000000000000000000")

	x := big.NewFloat(ethers)
	x.Mul(x, orders)

	wei, _ := x.Int(nil)

	return wei
}
