package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func Fatal(args ...interface{}) {
	red := color.New(color.FgRed)
	red.Fprint(os.Stderr, "error: ")
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	Fatal(fmt.Sprintf(format, args...))
}
