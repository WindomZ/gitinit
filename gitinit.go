package main

import (
	"github.com/WindomZ/gitinit/gitinit"
	"github.com/WindomZ/go-commander"
)

func main() {
	// gitinit
	commander.Program.
		Version("0.0.1").
		Description("A cli tool, easy way to 'git init' a new repository.").
		Action(gitinit.GitInitAction)

	// gitinit --init
	commander.Program.
		Command("-i --init").
		Description("create an empty Git repository or reinitialize an existing one").
		Action(gitinit.FlagInitAction).
		Option("--remote=<repo>", "manage 'origin' repository to '<repo>'")

	// gitinit --bare
	commander.Program.
		Command("-b --bare").
		Description("create a bare repository").
		Action(gitinit.FlagBareAction)

	commander.Program.Parse()
}
