// Copyright © 2019 sammyne <xiangminli@alumni.sjtu.edu.cn>
//
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file
//

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/sammyne/mastering-ethereum/playground/eth"

	"github.com/spf13/cobra"
)

var keysOutput string

// accountExportCmd represents the accountExport command
var accountExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export the all accounts",
	Run: func(cmd *cobra.Command, args []string) {
		keyJSON, err := eth.ExportAccounts(keyStoreDir, passphrase)

		if nil != err {
			fmt.Println(err)
			return
		}

		if err := os.RemoveAll(keysOutput); nil != err {
			fmt.Println(err)
			return
		}
		os.Mkdir(keysOutput, 0755)

		for i, data := range keyJSON {
			if err := ioutil.WriteFile(filepath.Join(keysOutput,
				fmt.Sprintf("key%02d.json", i)), data, 0644); nil != err {
				fmt.Println(err)
				return
			}
		}
	},
}

func init() {
	accountCmd.AddCommand(accountExportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// accountExportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// accountExportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	accountExportCmd.Flags().StringVar(&keysOutput, "out", "tmp",
		"output directory to place the exported keys")
}
