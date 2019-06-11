package main

import (
	"os"
	"github.com/yelirekim/gopherguides/src/cmd"
)

func main() {
	command := &cmd.Command{
		Args: os.Args[1:],
		Execute: runGurl,
	}
	command.Run()
}

func runGurl(cmd *cmd.Command) (int, error) {
	return 0, nil
}
