package dgit

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
)

func runCommand(c context.Context, name string, args ...string) ([]byte, error) {
	cmd := exec.Command(name, args...)
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = buf
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	if cmd.ProcessState.ExitCode() != 0 {
		return nil, fmt.Errorf("command %q exited with %d", name, cmd.ProcessState.ExitCode())
	}
	bytesData := buf.Bytes()
	bytesData = bytes.Trim(bytesData, "\n ")
	return bytesData, nil
}
