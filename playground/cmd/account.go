// Copyright © 2019 sammyne <xiangminli@alumni.sjtu.edu.cn>
//
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file
//

package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/sammyne/mastering-ethereum/playground/eth"

	"github.com/spf13/cobra"
)

var keyStoreDir string

// accountCmd represents the account command
var accountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("account called")
		fmt.Println(os.LookupEnv("HOME"))
	},
}

func init() {
	rootCmd.AddCommand(accountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	accountCmd.PersistentFlags().StringVar(&keyStoreDir, "keystore",
		filepath.Join(eth.DefaultDataDir(), "keystore"),
		"Directory for the keystore")
}
