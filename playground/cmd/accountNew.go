// Copyright © 2019 sammyne <xiangminli@alumni.sjtu.edu.cn>
//
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file
//

package cmd

import (
	"fmt"

	"github.com/sammyne/2018-mastering-ethereum/playground/eth"
	"github.com/spf13/cobra"
)

var (
	accountNewPassphrase string
)

// accountNewCmd represents the accountNew command
var accountNewCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new account",
	Long:  `Creates a new account and prints the address`,
	Run: func(cmd *cobra.Command, args []string) {
		_, account, err := eth.NewAccount(keyStoreDir, accountNewPassphrase)

		if nil != err {
			fmt.Println(err)
			return
		}
		fmt.Println(account.Address.Hex())
	},
}

func init() {
	accountCmd.AddCommand(accountNewCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountNewCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountNewCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	accountCmd.Flags().StringVar(&accountNewPassphrase, "passphrase", "hello",
		"passphrase to encrypted the generated account")
}
