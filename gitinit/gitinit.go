package gitinit

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

var GitInitAction = func(c commander.Context) error {
	fmt.Println("gitinit...")
	return nil
}
