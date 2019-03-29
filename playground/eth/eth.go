package eth

import (
	"context"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// DecodeContractAddresses decode the contract address out of the given txs
// indexed by the hashes
func DecodeContractAddresses(tx []common.Hash) ([]common.Address, error) {
	c, err := Dial()
	if nil != err {
		return nil, err
	}
	defer c.Close()

	contracts := make([]common.Address, len(tx))

	var wg sync.WaitGroup
	wg.Add(len(tx))
	for i, hash := range tx {
		go func(i int, hash common.Hash) {
			defer wg.Done()

			receipt, err := c.TransactionReceipt(context.TODO(), hash)
			if nil != err {
				return
			}

			contracts[i] = receipt.ContractAddress
		}(i, hash)
	}

	wg.Wait()

	return contracts, nil
}
