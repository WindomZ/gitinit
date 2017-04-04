package gitinit

import (
	"strings"
	"testing"
)

func TestPipeCommand(t *testing.T) {
	stdout, stderr, err := pipeCommand("ls")
	if err != nil {
		t.Fatal(err)
	} else if len(stderr) != 0 {
		t.Fatal("ERROR:", stderr)
	} else if len(stdout) != 0 {
		if !strings.Contains(stdout, ".go") {
			t.Fatal("ERROR:", stdout)
		}
	}
}

func TestExecCommand(t *testing.T) {
	stdout, err := execCommand("ls")
	if err != nil {
		t.Fatal(err)
	} else if len(stdout) != 0 {
		if !strings.Contains(stdout, ".go") {
			t.Fatal("ERROR:", stdout)
		}
	}
}
