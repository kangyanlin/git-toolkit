package utils

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"os"
	"os/exec"
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

	home, err: homedir.Dir()
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

func CheckOS() {
	if runtime.GOOS != "linux" && runtime.GOOS != "darwin" {
		fmt.Println("Platform not support!")
		os.Exit(1)
	}
}
