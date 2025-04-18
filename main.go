package main

import (
	"fmt"
	"github.com/nikitavoloboev/go/cmd"
	"os"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}
