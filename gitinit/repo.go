package gitinit

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

var RepoAction = func(c commander.Context) error {
	fmt.Println("repo...")
	return nil
}
