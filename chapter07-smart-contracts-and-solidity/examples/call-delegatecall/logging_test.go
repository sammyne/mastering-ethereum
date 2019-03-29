package main_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sammyne/mastering-ethereum/playground/eth"
)

func logging() {
}

func assert(t *testing.T, data []byte, addresses []common.Address) {
	if expect := len(addresses) * 32; len(data) != expect {
		t.Fatalf("invalid data length: got %d, expect %d", len(data), expect)
	}

	for i, j := 0, 0; j < len(addresses); i, j = i+32, j+1 {
		addr := common.BytesToAddress(data[i:(i + 32)])
		if addr != addresses[j] {
			t.Fatalf("#%d invalid address: got %s, expect %s", j, addr.Hex(),
				addresses[j].Hex())
		}
	}
}

func statMakeCalls(txHash common.Hash) ([]*types.Log, error) {
	c, err := eth.Dial()
	if nil != err {
		return nil, err
	}
	defer c.Close()

	receipt, err := c.TransactionReceipt(context.TODO(), txHash)
	if nil != err {
		return nil, err
	}

	return receipt.Logs, nil
}

func TestLogging(t *testing.T) {
	calledContractTx := common.HexToHash("0x6ea482e07a361bc21cf4081b9e2bb232f5b18be156dd9f587e27b0e57641c722")
	calledLibraryTx := common.HexToHash("0xa254c0d0bc5a9484aad891739b72bbdf9ec154d7fc95ad3b678614430ea6a0db")
	callerTx := common.HexToHash("0x0dfbd9d7992f433fff03f3e6326b8f6b72e8e29e6ba258d1fd7ae0e1813e3956")

	contracts, err := eth.DecodeContractAddresses([]common.Hash{
		calledContractTx, calledLibraryTx, callerTx,
	})
	if nil != err {
		t.Fatal(err)
	}
	//fmt.Println(contracts)
	//for _, c := range contracts {
	//	fmt.Println(c.Hex())
	//}
	calledContract, caller := contracts[0], contracts[2]

	_, accounts, err := eth.UnlockAccounts(eth.DefaultKeyDir(),
		eth.DefaultPassphrase())
	if nil != err {
		t.Fatal(err)
	}
	payer := accounts[0].Address
	//fmt.Println(payer.Hex())

	makeCallTx := common.HexToHash("0x3969914cd0e74445bc4b5ff2eb1239439e22fae6281c2b1b209dd1fabb97c847")
	logs, err := statMakeCalls(makeCallTx)
	if nil != err {
		t.Fatal(err)
	}

	topic := crypto.Keccak256([]byte("callEvent(address,address,address)"))
	for i, l := range logs {
		if !bytes.Equal(l.Topics[0].Bytes(), topic) {
			t.Fatalf("#%d invalid topic: got %s, expect %x", i,
				l.Topics[0].Hex(), topic)
		}
	}

	t.Log("checking the event emitted by _callContract.calledFunction()")
	assert(t, logs[0].Data, []common.Address{caller, payer, calledContract})

	t.Log("checking the event emitted by calledLibrary.calledFunction()")
	assert(t, logs[1].Data, []common.Address{payer, payer, caller})

	t.Log("check the event emitted by address(_calledContract).call(methodSig)")
	assert(t, logs[2].Data, []common.Address{caller, payer, calledContract})
	t.Log("check the event emitted by address(_calledContract).delegatecall(methodSig)")
	assert(t, logs[3].Data, []common.Address{payer, payer, caller})
}
