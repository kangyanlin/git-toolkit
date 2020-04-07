package cmd

import "github.com/spf13/cobra"

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "git-toolkit",
	Short: "Git工具集",
	Long:  "一个用于提高Git命令使用效率和规范Git Commit Message的工具集",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

func init() {
	// 添加install & uninstall 命令
	installCmd := NewInstall()
	installCmd.PersistentFlags().StringVar(&installDir, "dir", "/usr/local/bin", "install dir")

	uninstallCmd := NewUninstall()
	uninstallCmd.PersistentFlags().StringVar(&installDir, "dir", "/usr/local/bin/", "install dir")

	// 添加自定义Git命令

	// 安装 & 卸载
	RootCmd.AddCommand(installCmd)
	RootCmd.AddCommand(uninstallCmd)

	// Commit Message相关命令
	RootCmd.AddCommand(NewCi())

	// 创建分支相关命令
	RootCmd.AddCommand(NewFeatBranch())
	RootCmd.AddCommand(NewFixBranch())
	RootCmd.AddCommand(NewDocsBranch())
	RootCmd.AddCommand(NewStyleBranch())
	RootCmd.AddCommand(NewRefactorBranch())
	RootCmd.AddCommand(NewPerfBranch())
	RootCmd.AddCommand(NewHotFixBranch())
	RootCmd.AddCommand(NewTestBranch())
	RootCmd.AddCommand(NewChoreBranch())

	// 版本信息
	RootCmd.AddCommand(NewVersion())
}
