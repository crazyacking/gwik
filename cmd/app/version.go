package app

import (
	"fmt"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the gwik version number",
	Long:  `Print the gwik version number`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gwik %s\n", "v1.0.0")
	},
}
