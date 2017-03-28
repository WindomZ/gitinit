package gitinit

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

var FlagBareAction = func(c commander.Context) error {
	fmt.Println("--bare...")
	return nil
}
