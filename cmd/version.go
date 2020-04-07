package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/cobra"
	"runtime"
)

var bannerBase64 = "CiAgIF9fX19fIF8gXyAgICAgX19fX19fXyAgICAgICAgICBfIF8gICAgXyBfICAgCiAgLyBfX19fKF8pIHwgICB8X18gICBfX3wgICAgICAgIHwgfCB8ICAoXykgfCAgCiB8IHwgIF9fIF98IHxfICAgICB8IHwgX19fICAgX19fIHwgfCB8IF9fX3wgfF8gCiB8IHwgfF8gfCB8IF9ffCAgICB8IHwvIF8gXCAvIF8gXHwgfCB8LyAvIHwgX198CiB8IHxfX3wgfCB8IHxfICAgICB8IHwgKF8pIHwgKF8pIHwgfCAgIDx8IHwgfF8gCiAgXF9fX19ffF98XF9ffCAgICB8X3xcX19fLyBcX19fL3xffF98XF9cX3xcX198Cg=="

var versionTpl = `
%s

Name: git-toolkit
Version: %s
Arch: %s
BuildTime: %s
CommitID: %s
`

var (
	Version   string
	BuildTime string
	CommitID  string
)

// 构建打印版本命令
func NewVersion() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Long: `
Print version.`,
		Run: func(cmd *cobra.Command, args []string) {
			banner, _ := base64.StdEncoding.DecodeString(bannerBase64)
			fmt.Printf(versionTpl, banner, Version, runtime.GOOS+"/"+runtime.GOARCH, BuildTime, CommitID)
		},
	}
}
