package cmd

import (
	"encoding/json"
	"log"
	"miHoYoTools/shared"
	"miHoYoTools/utils"

	"github.com/spf13/cobra"
)

var (
	uid      string
	token    string
	account  string
	deviceId string
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set your account data list",
	Run: func(cmd *cobra.Command, args []string) {
		accountDataListItem := shared.AccountDataListItem{
			Uid:       uid,
			Token:     token,
			Account:   account,
			IsLogin:   true,
			LoginType: 1,
			DeviceId:  deviceId,
			Country:   "CN",
			AreaCode:  "**",
		}

		accountDataList := shared.AccountDataList{
			Data: []shared.AccountDataListItem{accountDataListItem},
		}

		data, err := json.Marshal(accountDataList)

		if err != nil {
			log.Fatal("[ERR] Failed to serialize AccountDataList data: ", err)
		}

		println("New Account Data List =", string(data))

		encryptedAccountDataList := utils.EncodeString(string(data))
		utils.SetAccountDataList(encryptedAccountDataList)

		println("OK")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&uid, "uid", "u", "", "UID (required)	")
	setCmd.Flags().StringVarP(&token, "token", "t", "", "Combo Token (required)")
	setCmd.Flags().StringVarP(&account, "account", "a", "", "Account (required)")
	setCmd.Flags().StringVarP(&deviceId, "deviceId", "d", "", "Device Id (required)")

	setCmd.MarkFlagRequired("uid")
	setCmd.MarkFlagRequired("token")
	setCmd.MarkFlagRequired("account")
	setCmd.MarkFlagRequired("deviceId")
}
