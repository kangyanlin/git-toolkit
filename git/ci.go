package git

import (
	"errors"
	"fmt"
	"github.com/mritd/promptx"
	"github.com/tonydeng/git-toolkit/utils"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type TypeMessage struct {
	Type          CommitType
	ZHDescription string
	ENDescription string
}

type Message struct {
	Type    CommitType
	Emoji   string
	Scope   string
	Subject string
	Body    string
	Footer  string
	Sob     string
}

// 选择提交类型
func SelectCommitType() CommitType {
	commitTypes := []TypeMessage{
		{Type: FEAT, ZHDescription: "新功能", ENDescription: "Introducing new features"},
		{Type: FIX, ZHDescription: "修复Bug", ENDescription: "Fixing a bug"},
		{Type: DOCS, ZHDescription: "添加文档", ENDescription: "Writing docs"},
		{Type: STYLE, ZHDescription: "调整格式", ENDescription: "Improving structure/format of the code"},
		{Type: REFACTOR, ZHDescription: "重构代码", ENDescription: "Refactoring code"},
		{Type: TEST, ZHDescription: "添加测试", ENDescription: "When adding missing tests"},
		{Type: CHORE, ZHDescription: "构建过程或辅助工具的变动", ENDescription: "Changing configuration files"},
		{Type: PERF, ZHDescription: "性能优化", ENDescription: "Improving performance"},
		{Type: HOTFIX, ZHDescription: "关键的热修复补丁", ENDescription: "Critical hotfix"},
		{Type: EXIT, ZHDescription: "退出", ENDescription: "Exit commit"},
	}

	cfg := &promptx.SelectConfig{
		ActiveTpl:    "» {{ .Type | cyan }} ({{ .ENDescription | cyan }})",
		InactiveTpl:  " {{ .Type | white }} ({{ .ENDescription | white }})",
		SelectPrompt: "Commit Type",
		SelectedTpl:  "{{ \"» \" | green }}{{\"Type:\" | cyan }} {{ .Type }}",
		DisPlaySize:  9,
		DetailsTpl: `
--------- Commit Type ----------
{{ "Type:" | faint }}	{{ .Type }}
{{ "Description:" | faint }}	{{ .ZHDescription }}({{ .ENDescription }})`,
	}

	s := &promptx.Select{
		Config: cfg,
		Items:  commitTypes,
	}

	idx := s.Run()

	if commitTypes[idx].Type == EXIT {
		fmt.Println("Talk is cheap. Show me the code.")
		os.Exit(0)
	}

	return commitTypes[idx].Type
}

// 获取Commit Type对应的Emoji
func GenEmoji(commitType CommitType) string {
	return CommitEmoji[commitType]
}

// 输入影响范围
func InputScope() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("Input is empty!")
		} else {
			return nil
		}
	}, "Scope(本次提交的范围，建议填写版本号):")

	return strings.TrimSpace(p.Run())
}

// 输入提交主题
func InputSubject() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		if strings.TrimSpace(string(line)) == "" {
			return errors.New("Input is empty!")
		} else if len(line) > 50 {
			return errors.New("Input length must < 25!")
		} else {
			return nil
		}
	}, "Subject(请添加简短的、必要的对本次提交的描述):")
	return strings.TrimSpace(p.Run())
}

// 输入完整提交信息
func InputBody() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "Body(添加一个完整提交描述):")

	body := strings.TrimSpace(p.Run())

	if body == "big" {
		return utils.OSEditInput()
	}

	return body
}

// 输入提交关联信息
func InputFooter() string {
	p := promptx.NewDefaultPrompt(func(line []rune) error {
		return nil
	}, "Footer(列出本次解决的、可以关闭的所有问题，建议使用关键字refs、close):")
	return strings.TrimSpace(p.Run())
}

// 生成SOB
func GenSOB() string {

	author := "Undefined"
	email := "Undefined"
	emailIndex := 0

	output := utils.ExecRtOut("git", "var", "GIT_AUTHOR_IDENT")
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

//  提交
func Commit(cm *Message) {
	if cm.Body == "" {
		cm.Body = cm.Subject
	}

	t, err := template.New("commitMessage").Parse(CommitTpl)
	utils.CheckAndExit(err)

	f, err := ioutil.TempFile("", "git-commit")
	utils.CheckAndExit(err)
	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()

	_ = t.Execute(f, cm)
	utils.Exec("git", "commit", "-F", f.Name())

	fmt.Println("\n✔ Always code as if the guy who ends up maintaining your code will be a violent psychopath who knows where you live.")
}
