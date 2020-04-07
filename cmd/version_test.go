package cmd

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestNewVersion(t *testing.T) {
	version :=    NewVersion()
	fmt.Println(version.Execute())
}

func TestBanner(t *testing.T)  {
	banner :=`
 ██████╗ ██╗████████╗    ████████╗ ██████╗  ██████╗ ██╗     ██╗  ██╗██╗████████╗
██╔════╝ ██║╚══██╔══╝    ╚══██╔══╝██╔═══██╗██╔═══██╗██║     ██║ ██╔╝██║╚══██╔══╝
██║  ███╗██║   ██║          ██║   ██║   ██║██║   ██║██║     █████╔╝ ██║   ██║   
██║   ██║██║   ██║          ██║   ██║   ██║██║   ██║██║     ██╔═██╗ ██║   ██║   
╚██████╔╝██║   ██║          ██║   ╚██████╔╝╚██████╔╝███████╗██║  ██╗██║   ██║   
 ╚═════╝ ╚═╝   ╚═╝          ╚═╝    ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝╚═╝   ╚═╝
`
	fmt.Println(banner)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(banner)))
}