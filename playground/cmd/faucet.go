// Copyright © 2019 sammyne <xiangminli@alumni.sjtu.edu.cn>
//
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file
//

package cmd

import (
	"github.com/spf13/cobra"
)

// faucetCmd represents the faucet command
var faucetCmd = &cobra.Command{
	Use:   "faucet",
	Short: "Collection of faucet related commands",
}

func init() {
	rootCmd.AddCommand(faucetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// faucetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// faucetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
