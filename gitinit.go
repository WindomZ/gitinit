package main

import (
	"github.com/WindomZ/gitinit/gitinit"
	"github.com/WindomZ/go-commander"
)

func main() {
	commander.Program.
		Version("0.0.1").
		Description("hello gitinit").
		Action(gitinit.GitInitAction)

	commander.Program.
		Option("--bare", "", gitinit.FlagBareAction).
		Action(gitinit.FlagBareAction)

	commander.Program.
		Command("repo").
		Action(gitinit.RepoAction)

	commander.Program.Parse()
}
