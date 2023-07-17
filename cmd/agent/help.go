package main

import (
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
)

func helpFunc(commands map[string]cli.CommandFactory) string {

	helpText := fmt.Sprintf(`Ain't much help yet...`)
	return strings.TrimSpace(helpText)
}
