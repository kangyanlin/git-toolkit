package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

const Cmd = "git"

var GitToolkitHome string
var InstallPath string
var HooksPath string
var GitCMHookPath string
var CurrentPath string

// 安装路径等配置初始化
func init() {

	var err error

	home := "/usr/local"
	CheckAndExit(err)

	GitToolkitHome = filepath.Join(home, "git-toolkit")
	InstallPath = filepath.Join(GitToolkitHome, "git-toolkit")
	HooksPath = filepath.Join(GitToolkitHome, "hooks")
	GitCMHookPath = filepath.Join(HooksPath, "commit-msg")

	CurrentPath, err = exec.LookPath(os.Args[0])

	CheckAndExit(err)
}

// 判断是否有错误
func IskErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// 检查错误并退出
func CheckAndExit(err error) {
	if !IskErr(err) {
		os.Exit(1)
	}
}

// 执行命令，忽略错误
func MustExec(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// 执行命令，有错误则退出
func Exec(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	CheckAndExit(cmd.Run())
}

// 执行命令并返回标准输出，有错误则退出
func ExecRtOut(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	b, err := cmd.Output()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(b)
}

// 执行命令忽略标准输出，有错误则退出
func ExecNoOut(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stderr = os.Stderr
	CheckAndExit(cmd.Run())
}

// 执行命令，当有错误将错误返回给调用者
func TryExec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return cmd.Run()
}

// 判断当前用户是否有root权限
func Root() bool {
	u, err := user.Current()
	CheckAndExit(err)
	return u.Uid == "0" || u.Gid == "0"
}

// 基于操作系统选择编辑器输入
func OSEditInput() string {
	f, err := ioutil.TempFile("", "git-toolkit")
	CheckAndExit(err)

	defer func() {
		_ = f.Close()
		_ = os.Remove(f.Name())
	}()

	// write utf8 bom
	bom := []byte{0xef, 0xbb, 0xbf}
	_, err = f.Write(bom)
	CheckAndExit(err)

	//获取系统编辑器
	editor := "vim"
	if runtime.GOOS == "windows" {
		editor = "notepad"
	}

	if v := os.Getenv("VISUAL"); v != "" {
		editor = v
	} else if e := os.Getenv("EDITOR"); e != "" {
		editor = e
	}

	//执行编辑器
	Exec(editor, f.Name())

	raw, err := ioutil.ReadFile(f.Name())
	CheckAndExit(err)
	input := string(bytes.TrimPrefix(raw, bom))

	return input
}

// 检查操作系统
func CheckOS() {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		fmt.Println("Platform not support!")
		os.Exit(1)
	}
}

// 构建命令目录
func GenBinPaths(dir string) []string {
	return []string{
		filepath.Join(dir, "git-ci"),
		filepath.Join(dir, "git-cm"),
		filepath.Join(dir, "git-feat"),
		filepath.Join(dir, "git-fix"),
		filepath.Join(dir, "git-docs"),
		filepath.Join(dir, "git-style"),
		filepath.Join(dir, "git-refactor"),
		filepath.Join(dir, "git-test"),
		filepath.Join(dir, "git-chore"),
		filepath.Join(dir, "git-pref"),
		filepath.Join(dir, "git-hotfix"),
		filepath.Join(dir, "git-ps"),
	}
}
