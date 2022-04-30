package cmd

import (
	"miHoYoTools/utils"

	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show account data list",
	Run: func(cmd *cobra.Command, args []string) {
		encryptedAccountDataList := utils.GetAccountDataList()
		accountDataList := utils.DecodeString(encryptedAccountDataList)
		println(accountDataList)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}
