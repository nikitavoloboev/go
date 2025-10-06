package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dzonerzy/go-snap/snap"
)

const (
	flowVersion = "1.0.0"
)

func main() {
	app := snap.New("flow", "flow is CLI to do things fast").
		Version(flowVersion).
		DisableHelp()

	app.Command("updateGoVersion", "Upgrade Go using the workspace script").
		Action(func(ctx *snap.Context) error {
			scriptPath, err := determineUpgradeScriptPath()
			if err != nil {
				return err
			}

			if _, err := os.Stat(scriptPath); err != nil {
				return fmt.Errorf("unable to access %s: %w", scriptPath, err)
			}

			cmd := exec.Command(scriptPath)
			cmd.Stdout = ctx.Stdout()
			cmd.Stderr = ctx.Stderr()
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("running %s: %w", scriptPath, err)
			}

			return nil
		})

	app.Command("version", "Reports the current version of flow").
		Action(func(ctx *snap.Context) error {
			fmt.Fprintln(ctx.Stdout(), flowVersion)
			return nil
		})

	args := os.Args[1:]
	if handled := handleTopLevel(args, os.Stdout); handled {
		return
	}

	app.RunAndExit()
}

func handleTopLevel(args []string, out io.Writer) bool {
	if len(args) == 0 {
		if err := openCurrentDirectory(out); err != nil {
			fmt.Fprintf(out, "open . failed: %v\n", err)
			printRootHelp(out)
		}
		return true
	}

	switch args[0] {
	case "--help", "-h", "h":
		printRootHelp(out)
		return true
	case "--version":
		fmt.Fprintln(out, flowVersion)
		return true
	case "help":
		if len(args) == 1 {
			printRootHelp(out)
			return true
		}
		if printCommandHelp(args[1], out) {
			return true
		}
		fmt.Fprintf(out, "Unknown help topic %q\n", args[1])
		return true
	}

	if len(args) > 1 {
		last := args[len(args)-1]
		if last == "--help" || last == "-h" {
			if printCommandHelp(args[0], out) {
				return true
			}
			printRootHelp(out)
			return true
		}
	}

	return false
}

func printCommandHelp(name string, out io.Writer) bool {
	switch name {
	case "updateGoVersion":
		fmt.Fprintln(out, "Upgrade Go using the workspace script")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "Usage:")
		fmt.Fprintln(out, "  flow updateGoVersion")
		return true
	case "version":
		fmt.Fprintln(out, "Reports the current version of flow")
		fmt.Fprintln(out)
		fmt.Fprintln(out, "Usage:")
		fmt.Fprintln(out, "  flow version")
		return true
	}

	return false
}

func printRootHelp(out io.Writer) {
	fmt.Fprintln(out, "flow is CLI to do things fast")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Usage:")
	fmt.Fprintln(out, "  flow [command]")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Available Commands:")
	fmt.Fprintln(out, "  help             Help about any command")
	fmt.Fprintln(out, "  updateGoVersion  Upgrade Go using the workspace script")
	fmt.Fprintln(out, "  version          Reports the current version of flow")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Flags:")
	fmt.Fprintln(out, "  -h, --help   help for flow")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Use \"flow [command] --help\" for more information about a command.")
}

func determineUpgradeScriptPath() (string, error) {
	if path := os.Getenv("FLOW_UPGRADE_SCRIPT_PATH"); path != "" {
		return path, nil
	}

	if root := os.Getenv("FLOW_CONFIG_ROOT"); root != "" {
		return filepath.Join(root, "sh", "upgrade-go-version.sh"), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("determine home directory: %w", err)
	}

	return filepath.Join(home, "src", "config", "sh", "upgrade-go-version.sh"), nil
}

func openCurrentDirectory(out io.Writer) error {
	cmd := exec.Command("open", ".")
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
