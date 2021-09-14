package installer

import (
	"bytes"
	"fmt"
	"os/exec"
)

type step struct {
	cmd string
}

func (s *step) execute() {
	stdOut, stdErr, err := shellExec(s.cmd)

	if len(stdOut) > 0 {
		fmt.Print(stdOut)
	}

	if len(stdErr) > 0 {
		fmt.Print(stdErr)
	}

	if err != nil {
		fmt.Print(err.Error())
	}
}

func shellExec(command string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	const defaultShell = "bash"

	cmd := exec.Command(defaultShell, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}
