package main

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/cmd"
	"os"
	"path/filepath"
)

func commandFor(basename string, rootCommand *cobra.Command) *cobra.Command {

	print("path: ", basename)

	return rootCommand
}

func main() {
	basename := filepath.Base(os.Args[0])
	commandFor(basename,cmd.RootCmd)
}
