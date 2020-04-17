package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/git"
	"github.com/tonydeng/git-toolkit/utils"
)

func NewPs() *cobra.Command {
	return &cobra.Command{
		Use:                        "git-ps",
		Aliases:                    []string{"git-ps"},
		Short:                      "Push local branch to remote git server",
		Long:                       "Push local branch to remote git server.",
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Run: func(cmd *cobra.Command, args []string) {
			git.Push()
		},
	}
}