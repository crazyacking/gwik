package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "default",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gwik")
		_ = cmd.Help()
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
