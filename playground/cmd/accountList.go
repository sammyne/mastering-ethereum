// Copyright © 2019 sammyne <xiangminli@alumni.sjtu.edu.cn>
//
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file
//

package cmd

import (
	"fmt"

	"github.com/sammyne/mastering-ethereum/playground/eth"

	"github.com/spf13/cobra"
)

// accountListCmd represents the accountList command
var accountListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print summary of existing accounts",
	Long:  `Print a short summary of all accounts`,
	Run: func(cmd *cobra.Command, args []string) {
		_, accounts, err := eth.UnlockAccounts(keyStoreDir, passphrase)

		if nil != err {
			fmt.Println(err)
			return
		}

		for i, a := range accounts {
			fmt.Printf("[%d]: %s\n", i, a.Address.Hex())
		}
	},
}

func init() {
	accountCmd.AddCommand(accountListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
