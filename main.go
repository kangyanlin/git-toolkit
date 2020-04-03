package main

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/cmd"
	"github.com/tonydeng/git-toolkit/utils"
	"os"
	"path/filepath"
)

func commandFor(basename string, rootCommand *cobra.Command) *cobra.Command {
	c, _, _ := rootCommand.Find([]string{basename})

	if c != nil {
		rootCommand.RemoveCommand(c)
		return c
	}
	return rootCommand
}

func main() {
	basename := filepath.Base(os.Args[0])
	utils.CheckAndExit(commandFor(basename, cmd.RootCmd).Execute())
}
