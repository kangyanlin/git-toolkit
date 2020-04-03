package utils

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
)

var GitToolkitHome string
var InstallPath string
var HooksPath string
var GitCMHookPath string
var CurrentPath string

func init() {

	var err error

	home, err := homedir.Dir()
	CheckAndExit(err)

	GitToolkitHome = filepath.Join(home, "git-toolkit")
	InstallPath = filepath.Join(GitToolkitHome, "git-toolkit")
	HooksPath = filepath.Join(GitToolkitHome, "hooks")
	GitCMHookPath = filepath.Join(HooksPath, "commit-msg")

	CurrentPath, err = exec.LookPath(os.Args[0])

	CheckAndExit(err)
}

func CheckErr(err error) bool {
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CheckAndExit(err error) {
	if !CheckErr(err) {
		os.Exit(1)
	}
}

func MustExecRtOut(name string, arg ...string) string {
	cmd := exec.Command(name, arg...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	b, err := cmd.Output()

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	return string(b)
}

func MustExecNoOut(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stderr = os.Stderr
	CheckAndExit(cmd.Run())
}

func TryExec(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return cmd.Run()
}

func Root() bool {
	u, err := user.Current()
	CheckAndExit(err)
	return u.Uid == "0" || u.Gid == "0"
}

func CheckOS() {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		fmt.Println("Platform not support!")
		os.Exit(1)
	}
}
