package config

import (
	"fmt"

	"github.com/nikitavoloboev/go/cmd"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage configuration",
}

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get a configuration value",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Implement config get logic
		fmt.Printf("Getting config for key: %s\n", args[0])
	},
}

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Implement config set logic
		fmt.Printf("Setting config %s = %s\n", args[0], args[1])
	},
}

func init() {
	configCmd.AddCommand(getCmd)
	configCmd.AddCommand(setCmd)
	cmd.RootCmd.AddCommand(configCmd)
}
