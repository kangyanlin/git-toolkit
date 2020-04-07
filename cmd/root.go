package cmd

import "github.com/spf13/cobra"

var cfgFile string

var RootCmd = &cobra.Command{
	Use:   "git-toolkit",
	Short: "git工具集",
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

	RootCmd.AddCommand(installCmd)
	RootCmd.AddCommand(uninstallCmd)

	// 添加自定义Git命令
	RootCmd.AddCommand(NewCi())
}
