// Copyright © 2019 sammyne <xiangminli@alumni.sjtu.edu.cn>
//
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file
//

package cmd

import (
	"context"
	"fmt"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/sammyne/mastering-ethereum/playground/eth"
	"github.com/spf13/cobra"
)

// accountBalanceCmd represents the accountBalance command
var accountBalanceCmd = &cobra.Command{
	Use:   "balance",
	Short: "List balances for all accounts",
	Run: func(cmd *cobra.Command, args []string) {
		_, accounts, err := eth.UnlockAccounts(keyStoreDir, passphrase)

		if nil != err {
			fmt.Println(err)
			return
		}

		c, err := eth.Dial()
		if nil != err {
			fmt.Println(err)
			return
		}

		var wg sync.WaitGroup
		balances := make([]*big.Int, len(accounts))

		wg.Add(len(accounts))
		for i, a := range accounts {
			//fmt.Printf("[%d]: %s: %s\n", i, )
			go func(i int, address common.Address) {
				defer wg.Done()

				var err error
				balances[i], err = c.BalanceAt(context.TODO(), address, nil)
				if nil != err {
					fmt.Printf("failed to fetch balance for [%d]:%s\n", i, address.Hex())
				}
			}(i, a.Address)
		}

		wg.Wait()
		for i, a := range accounts {
			fmt.Printf("[%d] %s: %s\n", i, a.Address.Hex(), balances[i])
		}
	},
}

func init() {
	accountCmd.AddCommand(accountBalanceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountBalanceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountBalanceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
