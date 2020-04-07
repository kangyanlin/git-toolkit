package utils

import (
	"fmt"
	"os/user"
	"testing"
)

func TestRoot(t *testing.T) {
	u, err := user.Current()
	if err != nil {

	}
	fmt.Printf("user is %s, uid='%s', gid='%s'\n", u.Name, u.Uid, u.Gid)

	fmt.Printf("%s = root is %t\n", u.Username, Root())
}
