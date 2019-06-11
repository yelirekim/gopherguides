package cmd

import "github.com/yelirekim/gopherguides/src/ui"

type RunCommand func(*Command) (int, error)

type Command struct {
	Args    []string
	Execute RunCommand
	io      ui.IO
}

func (cmd *Command) Run() {
	cmd.Execute(cmd)
}
