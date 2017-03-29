package gitinit

import (
	"fmt"
	"github.com/WindomZ/go-commander"
)

var FlagInitAction = func(c commander.Context) error {
	fmt.Println("--init...")
	return nil
}
