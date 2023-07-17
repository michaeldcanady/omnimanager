package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/cli"
)

func main() {
	os.Exit(actualMain())
}

func actualMain() int {

	// Get the command line args.
	binName := filepath.Base(os.Args[0])
	args := os.Args[1:]

	// In tests, Commands may already be set to provide mock commands
	if Commands == nil {
		// Commands get to hold on to the original working directory here,
		// in case they need to refer back to it for any special reason, though
		// they should primarily be working with the override working directory
		// that we've now switched to above.
		initCommands()
	}

	// Build the CLI so far, we do this so we can query the subcommand.
	cliRunner := &cli.CLI{
		Name:       binName,
		Args:       args,
		Commands:   Commands,
		HelpFunc:   helpFunc,
		HelpWriter: os.Stdout,
	}

	exitCode, err := cliRunner.Run()
	if err != nil {
		fmt.Printf("Error executing CLI: %s", err.Error())
		return 1
	}

	// if we are exiting with a non-zero code, check if it was caused by any
	// plugins crashing
	if exitCode != 0 {
		//for _, panicLog := range logging.PluginPanics() {
		///	Ui.Error(panicLog)
		//}
	}

	return exitCode
}
