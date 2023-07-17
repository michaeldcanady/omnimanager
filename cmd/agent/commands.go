package main

import (
	"github.com/michaeldcanady/omnimanager/internal/command"
	"github.com/mitchellh/cli"
)

var (
	Commands map[string]cli.CommandFactory
)

func initCommands() {
	Commands = map[string]cli.CommandFactory{
		"sync": func() (cli.Command, error) {
			return &command.SyncCommand{}, nil
		},
	}
}
