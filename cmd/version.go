package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/utils"
)

var (
	Version   string
	BuildTime string
	CommitID  string
)

// 构建打印版本命令
func NewVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "git-toolkit version",
		Short: "Print version",
		Long: `
Print version.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(utils.GenVersion(Version,BuildTime,CommitID))
		},
	}
}
