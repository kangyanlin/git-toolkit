package utils

import (
	"fmt"
	"testing"
)

func TestGenVersion(t *testing.T) {

	versionStr := GenVersion("1.2","2019-11-11","adfasdfasdffee")

	fmt.Println(versionStr)
}
