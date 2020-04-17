package cmd

import (
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/utils"
)

// 构建卸载命令
func NewUninstall() *cobra.Command {
	return &cobra.Command{
		Use:     "git-toolkit uninstall",
		Short:   "Uninstall git-toolkit",
		Long:    "Uninstall git-toolkit.",
		Aliases: []string{"uninstall"},
		Run: func(cmd *cobra.Command, args []string) {
			utils.Uninstall(installDir)
		},
	}
}
