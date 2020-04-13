package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/utils"
)

func NewClog() *cobra.Command {
	return &cobra.Command{
		Use:     "git-clog",
		Aliases: []string{"git-clog"},
		Short:   "Git Change logs Output",
		Long:    "Summarizes git log output in a format suitable for inclusion in release announcements. Each commit will be grouped by author and title.",
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("git clog %s",args[0])
		},
	}
}
