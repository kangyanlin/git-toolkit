package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

// 安装
func Install(dir string) {
	currentPath, err := exec.LookPath(os.Args[0])
	CheckAndExit(err)

	if !Root() {
		cmd := exec.Command("sudo", currentPath, "install", "--dir", dir)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		CheckAndExit(cmd.Run())
	} else {
		Uninstall(dir)

		fmt.Printf("📥 mkdir %s\n", HooksPath)
		CheckAndExit(os.MkdirAll(HooksPath, 0755))

		fmt.Printf("📥 copy file %s\n", InstallPath)
		currentFile, err := os.Open(CurrentPath)
		CheckAndExit(err)
		defer func() { _ = currentFile.Close() }()

		installFile, err := os.OpenFile(InstallPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0755)
		CheckAndExit(err)
		defer func() { _ = installFile.Close() }()

		_, err = io.Copy(installFile, currentFile)
		CheckAndExit(err)

		var binPaths = GenBinPaths(dir)

		for _, binPath := range binPaths {
			fmt.Printf("📥 install symbolic %s\n", binPath)
			CheckAndExit(os.Symlink(InstallPath, binPath))
		}
		SetGlobalConfig()
	}
}

// 设置Git整体定制配置
func SetGlobalConfig() {
	// 设置Alias
	MustExec(Cmd, "config", "--global", "alias.co", "checkout")
	MustExec(Cmd, "config", "--global", "alias.br", "branch")
	MustExec(Cmd, "config", "--global", "alias.st", "status")
	MustExec(Cmd, "config", "--global", "alias.lg", "log --oneline --decorate --graph " +
		"--pretty=format:'%Cred%h%Creset -%C(yellow)%d%Creset %s %Cgreen (%cr) %C(bold blue)<%an>%Creset' " +
		"--abbrev-commit")
	MustExec(Cmd, "config", "--global", "alias.mrg", "merge --no-ff")

	// 设置自定义的HooksPath
	fmt.Printf("📥 config set core.hooksPath %s\n", HooksPath)
	MustExec(Cmd, "config", "--global", "core.hooksPath", HooksPath)
}