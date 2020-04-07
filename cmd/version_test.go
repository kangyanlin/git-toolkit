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
   _____ _ _     _______          _ _    _ _   
  / ____(_) |   |__   __|        | | |  (_) |  
 | |  __ _| |_     | | ___   ___ | | | ___| |_ 
 | | |_ | | __|    | |/ _ \ / _ \| | |/ / | __|
 | |__| | | |_     | | (_) | (_) | |   <| | |_ 
  \_____|_|\__|    |_|\___/ \___/|_|_|\_\_|\__|
`
	fmt.Println(banner)
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(banner)))
}