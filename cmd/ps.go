package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/git"
)

func NewPs() *cobra.Command {
	return &cobra.Command{
		Use:                        "ps",
		Aliases:                    []string{"git-ps"},
		Short:                      "Push local branch to remote git server",
		Long:                       "Push local branch to remote git server.",
		Run: func(cmd *cobra.Command, args []string) {
			git.Push()
		},
	}
}