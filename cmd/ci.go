package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"qianxin.com/dengtao/git-toolkit/git"
)

var fastCommit = false

func NewCi() *cobra.Command {
	ciCmd := &cobra.Command{
		Use:   "ci",
		Short: "交互式输入 commit message",
		Long: `
交互式输入 git commit message, commit message格式为

<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>

该格式来源于 Angular 社区提交规范
`,
		Aliases: []string{"git-ci"},
		Run: func(cmd *cobra.Command, args []string) {
			git.CheckGitProject()

			if !git.CheckStageFiles() {
				fmt.Println("No staged any files")
				os.Exit(1)
			}

			cm := &git.Message{Sob: git.GenSOB()}
		},
	}
}
