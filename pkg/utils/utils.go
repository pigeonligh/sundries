package utils

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func SysExecShell(shell string) {
	err := syscall.Exec("/usr/bin/env", []string{"env", shell}, append(os.Environ(), "SKIP_ENTRANCE=1", "TERM=xterm-256color"))
	if err != nil {
		panic(err)
	}
}

func ExecCommand(ctx context.Context, cmdStr string) (string, string, error) {
	var stderr bytes.Buffer
	var stdout bytes.Buffer

	cmd := exec.CommandContext(ctx, "/bin/bash", "-c", cmdStr)
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()

	return stdout.String(), stderr.String(), err
}

func Fatal(err error) {
	fmt.Printf("There's been an error: %v", err)
	os.Exit(1)
}
