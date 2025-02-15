package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "dev" // This can be set during build

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("mycli version %s\n", Version)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
}
