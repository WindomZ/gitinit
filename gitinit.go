package main

import (
	"github.com/WindomZ/gitinit/gitinit"
	"github.com/WindomZ/go-commander"
)

func main() {
	// gitinit
	commander.Program.
		Version("0.0.1").
		Description("hello gitinit").
		Action(gitinit.GitInitAction)

	// gitinit --init
	commander.Program.
		Command("--init").
		Action(gitinit.FlagInitAction)

	// gitinit --bare
	commander.Program.
		Command("--bare").
		Action(gitinit.FlagBareAction)

	// gitinit repo
	commander.Program.
		Command("repo").
		Action(gitinit.RepoAction)

	commander.Program.Parse()
	//commander.Program.ShowHelpMessage()
}
