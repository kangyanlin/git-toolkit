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
	emailIndex := 0

	output := utils.MustExecRtOut("git", "var", "GIT_AUTHOR_IDENT")
	authorInfo := strings.Fields(output)

	for i := 0; i < len(authorInfo); i++ {

		if i < len(authorInfo)-2 {
			//fmt.Println(authorInfo[i])
			if strings.Index(authorInfo[i], "<") >= 0 {
				email = authorInfo[i]
				emailIndex = i
			}
		}
	}

	if emailIndex > 0 {
		author = strings.Join(authorInfo[0:emailIndex], " ")
	}

	return strings.Join([]string{"Signed-off-by:", author, email}, " ")
}
