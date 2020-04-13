package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tonydeng/git-toolkit/git"
	"github.com/tonydeng/git-toolkit/utils"
	"os"
)

var fastCommit = false

// 构造git ci命令，并执行
func NewCi() *cobra.Command {
	ciCmd := &cobra.Command{
		Use:   "git-ci",
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
		Version: utils.GenVersion(Version, BuildTime, CommitID),
		Aliases: []string{"git-ci"},
		Run: func(cmd *cobra.Command, args []string) {
			git.CheckGitProject()

			if !git.CheckStageFiles() {
				fmt.Println("No staged any files")
				os.Exit(1)
			}

			cm := &git.Message{Sob: git.GenSOB()}

			if fastCommit {
				cm.Type = git.FEAT
				cm.Emoji = git.GenEmoji(cm.Type)
				cm.Scope = "Undefined"
				cm.Subject = git.InputSubject()
				git.Commit(cm)
			} else {
				cm.Type = git.SelectCommitType()
				cm.Emoji = git.GenEmoji(cm.Type)
				cm.Scope = git.InputScope()
				cm.Subject = git.InputSubject()
				cm.Body = git.InputBody()
				cm.Footer = git.InputFooter()
				cm.Sob = git.GenSOB()
				git.Commit(cm)
			}
		},
	}

	ciCmd.Flags().BoolVarP(&fastCommit, "fast", "f", false, "快速提交")

	return ciCmd
}
