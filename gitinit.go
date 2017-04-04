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
		Command("-i --init").
		Action(gitinit.FlagInitAction).
		Option("--remote=<repo>")

	// gitinit --bare
	commander.Program.
		Command("-b --bare").
		Action(gitinit.FlagBareAction)

	commander.Program.Parse()
}
