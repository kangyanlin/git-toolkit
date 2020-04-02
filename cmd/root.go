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

}
