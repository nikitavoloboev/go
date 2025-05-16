package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "go",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test.")
	},
}

func init() {
	// Global flags
	RootCmd.PersistentFlags().CountP("verbose", "v", "verbose output")
	RootCmd.PersistentFlags().Bool("no-color", false, "disable colored output")
}
