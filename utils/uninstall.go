package utils

import (
	"fmt"
	"os"
	"os/exec"
)

//卸载
func Uninstall(dir string) {
	CheckOS()

	currentPath, err := exec.LookPath(os.Args[0])
	CheckAndExit(err)

	if !Root() {
		cmd := exec.Command("sudo", currentPath, "uninstall", "--dir", dir)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		CheckAndExit(cmd.Run())
	} else {
		fmt.Printf("👉 remove %s\n", GitToolkitHome)
		_ = os.RemoveAll(GitToolkitHome)

		var binPaths = GenBinPaths(dir)

		for _, binPath := range binPaths {
			_ = os.Remove(binPath)
		}

		UnsetGlobalConfig()
	}
}

// 删除Git整体定制化设置
func UnsetGlobalConfig() {
	// 删除Alias配置
	Exec(Cmd, "config", "--global", "--unset", "alias.co")
	Exec(Cmd, "config", "--global", "--unset", "alias.br")
	Exec(Cmd, "config", "--global", "--unset", "alias.st")
	Exec(Cmd, "config", "--global", "--unset", "alias.lg")
	Exec(Cmd, "config", "--global", "--unset", "alias.mrg")

	// 删除hooksPath配置
	Exec("git", "config", "--global", "--unset", "core.hooksPath")
}
