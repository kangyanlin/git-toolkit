package git

import (
	"fmt"
	"strings"
	"testing"
)

func TestGenSOB(t *testing.T) {
	sob := GenSOB()

	if len(sob) <= 0 {
		t.Fatalf("gen sob error")
	}
	fmt.Println(sob)

	GIT_AUTHOR_IDENT := "Eric W. Biederman <ebiederm@lnxi.com> 1121223278 -0600"
	authorInfo := strings.Fields(GIT_AUTHOR_IDENT)

	author := "Undefined"
	email := "Undefined"
	emailIndex := 0
	for i := 0; i < len(authorInfo); i++ {

		if i < len(authorInfo)-2 {
			//fmt.Println(authorInfo[i])
			if strings.Index(authorInfo[i], "<") >= 0 {
				email = authorInfo[i]
				emailIndex = i
			}
		}
	}

	if emailIndex > 0 {
		author = strings.Join(authorInfo[0:emailIndex], " ")
	}

	fmt.Println(author, email)
}
