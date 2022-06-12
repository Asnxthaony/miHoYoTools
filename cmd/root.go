package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	isOversea bool
)

var rootCmd = &cobra.Command{
	Use:   "miHoYoTools",
	Short: "A useful tool for Genshin Impact",
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
