package gitinit

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strings"
)

func pipeCommand(name string, arg ...string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	cmd := exec.Command(name, arg...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return stdout.String(), stderr.String(), err
	}
	if err := cmd.Wait(); err != nil {
		return stdout.String(), stderr.String(), err
	}
	return stdout.String(), stderr.String(), nil
}

func execCommand(name string, arg ...string) (string, error) {
	stdout, stderr, err := pipeCommand(name, arg...)
	os.Stdout.WriteString(stdout)
	os.Stderr.WriteString(stderr)
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return strings.TrimSpace(stdout + stderr), err
	} else if len(stderr) != 0 {
		return stdout, errors.New(stderr)
	}
	return stdout, nil
}
