package gitinit

import (
	"errors"
	"github.com/WindomZ/go-commander"
	"regexp"
)

var FlagRemoteAction = func(c commander.Context) error {
	if repo := c.MustString("--remote"); len(repo) != 0 {
		if !validGitAddress(repo) {
			return errors.New("ERROR: Invalid git address.")
		} else if _, err := execCommand("git", "remote", "add", "origin", repo); err != nil {
			return err
		}
	}
	return nil
}

func validGitAddress(repo string) bool {
	matched, err := regexp.MatchString(
		`((git|ssh|http(s)?)|(git@[\w\.]+))(:(//)?)([\w\.@\:/\-~]+)(\.git)(/)?`,
		repo)
	return matched && err == nil
}
