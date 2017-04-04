package gitinit

import (
	"errors"
	"github.com/WindomZ/go-commander"
	"strings"
)

var FlagBareAction = func(c commander.Context) error {
	if out, err := execCommand("git", "init", "--bare"); err != nil {
		return err
	} else if len(out) != 0 {
		if !strings.Contains(out, "Initialized empty Git repository") {
			return errors.New(out)
		}
	}
	return nil
}
