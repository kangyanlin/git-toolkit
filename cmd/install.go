package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/utils"
)

var installDir string

// 构建安装命令
func NewInstall() *cobra.Command {
	return &cobra.Command{
		Use:     "git-toolkit install",
		Short:   "Install git-toolkit",
		Long:    "Install git-toolkit(only support *Unix).",
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Aliases: []string{"install"},
		Run: func(cmd *cobra.Command, args []string) {
			utils.Install(installDir)
		},
	}
}
