package gitinit

import "github.com/WindomZ/go-commander"

var FlagRemoteAction = func(c commander.Context) error {
	if repo := c.MustString("--remote"); len(repo) != 0 {
		if _, err := execCommand("git", "remote", "add", "origin", repo); err != nil {
			return err
		}
	}
	return nil
}
