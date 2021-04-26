package shell

import (
	"bytes"
	"os/exec"
)

func RunCommand(str string, shell bool) (string, error) {
	var cmd *exec.Cmd
	if shell == true {
		cmd = exec.Command("/bin/sh", "-c", str)
	} else {
		cmd = exec.Command(str)
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr := string(stdout.String())
	return outStr, err
}
