package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gomock-proj",
	Short: "gomock-proj traverses directories while calling mockgen to mock entire project",
}

// Execute cobra root cmd entry
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("call rootCmd error")
		os.Exit(1)
	}
}
