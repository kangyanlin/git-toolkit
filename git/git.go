package git

import (
	"qianxin.com/dengtao/git-toolkit/utils"
	"strings"
)

type CommitType string
type RepoType string

const Cmd = "git"

const (
	FEAT     CommitType = "feat"
	FIX      CommitType = "fix"
	DOCS     CommitType = "docs"
	STYLE    CommitType = "style"
	REFACTOR CommitType = "refactor"
	TEST     CommitType = "test"
	CHORE    CommitType = "chore"
	PERF     CommitType = "perf"
	HOTFIX   CommitType = "hotfix"
	EXIT     CommitType = "exit"
)

const CommitTpl = `{{ .Type }}({{ .Scope }}): {{ .Subject }}

{{ .Body }}

{{ .Footer }}

{{ .sub }}
`

const CommitMessagePattern = `^(?:fixup!\s*)?(\w*)(\(([\w\$\.\*/-].*)\))?\: (.*)|^Merge\ branch(.*)`

func CheckGitProject() {
	utils.MustExecNoOut(Cmd, "rev-parse", "--show-toplevel")
}

func CheckStageFiles() bool {
	output := utils.MustExecRtOut(Cmd, "diff", "--cached", "--name-only")
	return strings.TrimSpace(output) != ""
}

func CheckLastCommitInfo() *[]string {
	title := utils.MustExecRtOut(Cmd, "log", "-1", "--pretty=format:%s")
	desc := utils.MustExecRtOut(Cmd, "log", "-1", "--pretty=format:%b")
	return &[]string{title, desc}
}
