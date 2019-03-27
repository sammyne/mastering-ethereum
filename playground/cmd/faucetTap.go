// Copyright © 2019 sammyne <xiangminli@alumni.sjtu.edu.cn>
//
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file
//

package cmd

import (
	"github.com/spf13/cobra"
)

var address string

// faucetTapCmd represents the faucetTap command
var faucetTapCmd = &cobra.Command{
	Use:   "tap",
	Short: "Tap some ethers into a given address",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: fill in the implementation
	},
}

func init() {
	faucetCmd.AddCommand(faucetTapCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// faucetTapCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// faucetTapCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	faucetTapCmd.Flags().StringVar(&address, "address", "",
		"address to received the tapped ethers")

	faucetTapCmd.MarkFlagRequired("address")
}
