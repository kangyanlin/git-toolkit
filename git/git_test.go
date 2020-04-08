package git

import (
	"fmt"
	"testing"
)

func TestGetCurrentBranch(t *testing.T) {

	branch := GetCurrentBranch()

	if len(branch) > 0 {
		fmt.Printf("current branch %s\n", branch)
	}
}
