package git

type CommitType string
type RepoType string

const Cmd  = "git"

const (
	FEAT CommitType = "feat"
	FIX CommitType = "fix"
	DOCS CommitType = "docs"
	STYLE CommitType = "style"
	REFACTOR CommitType = "refactor"
	TEST CommitType = "test"
	CHORE CommitType = "chore"
	PERF CommitType = "perf"
	HOTFIX CommitType = "hotfix"
	EXIT CommitType = "exit"
)

const CommitTpl  = `{{ .Type }}({{ .Scope }}): {{ .Subject }}

{{ .Body }}

{{ .Footer }}

{{ .sub }}
`

const CommitMessagePattern  = `^(?:fixup!\s*)?(\w*)(\(([\w\$\.\*/-].*)\))?\: (.*)|^Merge\ branch(.*)`

func CheckGitProject()  {

}