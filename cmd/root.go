package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "mycli",
	Short: "My CLI application",
	Long:  `A longer description of my CLI application.`,
}

func init() {
	// Global flags
	RootCmd.PersistentFlags().CountP("verbose", "v", "verbose output")
	RootCmd.PersistentFlags().Bool("no-color", false, "disable colored output")
}
