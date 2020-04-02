package main

import (
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"qianxin.com/dengtao/git-toolkit/cmd"
)

func commandFor(basename string, rootCommand *cobra.Command) *cobra.Command {

	print("path: ", basename)

	return rootCommand
}

func main() {
	basename := filepath.Base(os.Args[0])
	commandFor(basename,cmd.RootCmd)
}
