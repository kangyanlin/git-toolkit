package git

import (
	"qianxin.com/dengtao/git-toolkit/utils"
	"strings"
)

type Message struct {
	Type    CommitType
	Scope   string
	Subject string
	Body    string
	Footer  string
	Sob     string
}

// 生成SOB
func GenSOB() string {

	author := "Undefined"
	email := "Undefined"

	output := utils.MustExecRtOut("git", "var", "GIT_AUTHOR_IDENT")
	authorInfo := strings.Fields(output)

	if len(authorInfo) > 1 && authorInfo[0] != "" {
		author = authorInfo[0]
	}

	if len(authorInfo) > 2 && authorInfo[1] != "" {
		email = authorInfo[1]
	}

	return "Signed-off-by: " + author + " " + email
}
