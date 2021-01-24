package cmd

import (
	"fmt"

	"github.com/shanqincheng/gomock-proj/internal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(mockCmd)

	mockCmd.Flags().StringVarP(&dir, "dir", "d", ".", "dir to mock from")
}

var dir string

var mockCmd = &cobra.Command{
	Use:   "mock",
	Short: "mock use gomock to mock interface",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("mock dir: %s\n", dir)

		mkdirs, done := internal.NewMkProject()
		mkdirs.QueueDirs() <- dir
		mkdirs.QueueClose()
		<-done
		fmt.Printf("mockCmd dir done\n")

		if err := internal.GoimportsMockDir(); err != nil {
			fmt.Printf("internal.GoimportsMockDir: %s", err)
			return
		}
		fmt.Printf("goimports mock %s done\n", internal.MockDir)
	},
}
