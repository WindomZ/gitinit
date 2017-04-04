package gitinit

import (
	"errors"
	"github.com/WindomZ/go-commander"
	"strings"
)

var FlagInitAction = func(c commander.Context) error {
	if out, err := execCommand("git", "init"); err != nil {
		return err
	} else if len(out) != 0 {
		if !strings.Contains(out, "Initialized empty Git repository") {
			return errors.New(out)
		}
	}
	if err := FlagRemoteAction(c); err != nil {
		return err
	}
	return nil
}
