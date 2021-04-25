package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func run_command(str string) (string, error) {
	cmd := exec.Command(str)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	outStr := string(stdout.String())
	return outStr, err
}

func main() {
	var out, err = run_command("tasklist")
	if err == nil {
		fmt.Println(out)
	} else {
		fmt.Println(err.Error())
	}

}
