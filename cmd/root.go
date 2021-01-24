package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomock-proj",
	Short: "gomock-proj continusly call gomock to mock whole project",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("call rootCmd error")
		os.Exit(1)
	}
}
