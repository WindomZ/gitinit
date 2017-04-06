package gitinit

import (
	"errors"
	"github.com/WindomZ/go-commander"
)

var FlagOriginAction = func(c commander.Context) error {
	if repo := c.MustString("--origin"); len(repo) != 0 {
		if !ValidGitAddress(repo) {
			return errors.New("ERROR: Invalid git address: " + repo)
		} else if _, err := execCommand("git", "remote", "add", "origin", repo); err != nil {
			return err
		}
	}
	return nil
}
