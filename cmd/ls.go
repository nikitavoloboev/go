package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:   "ls [path]",
	Short: "List files in a directory",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Default to current directory if no path is provided
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		// Get file info to check if it's a directory
		fileInfo, err := os.Stat(path)
		if err != nil {
			return fmt.Errorf("error accessing path %s: %w", path, err)
		}

		// If it's not a directory, just print the file name and return
		if !fileInfo.IsDir() {
			fmt.Println(fileInfo.Name())
			return nil
		}

		// Read directory contents
		entries, err := os.ReadDir(path)
		if err != nil {
			return fmt.Errorf("error reading directory %s: %w", path, err)
		}

		// Print each entry
		for _, entry := range entries {
			// Get file info for additional details
			info, err := entry.Info()
			if err != nil {
				// If we can't get info, just print the name
				fmt.Println(entry.Name())
				continue
			}

			// Format: [D/F] Name Size
			fileType := "F"
			if entry.IsDir() {
				fileType = "D"
			}

			fmt.Printf("[%s] %s\t%d bytes\n", fileType, info.Name(), info.Size())
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(lsCmd)
}
