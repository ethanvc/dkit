package dgit

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func splitString(s string, sep string) []string {
	parts := strings.Split(s, sep)
	result := make([]string, 0, len(parts))
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		result = append(result, part)
	}
	return result
}

func ListLocalBranches(c context.Context) ([]string, error) {
	buf, _, err := runCommand(c, "git", "branch", "--format=%(refname:short)")
	if err != nil {
		return nil, err
	}
	branches := splitString(string(buf), "\n")
	return branches, nil
}

func runCommand(c context.Context, name string, args ...string) ([]byte, int, error) {
	cmd := exec.Command(name, args...)
	buf := bytes.NewBuffer(nil)
	cmd.Stdout = buf
	cmd.Stderr = buf
	err := cmd.Run()
	bytesData := buf.Bytes()
	bytesData = bytes.Trim(bytesData, "\n ")
	return bytesData, cmd.ProcessState.ExitCode(), err
}

func ListMergedBranches(c context.Context, targetBranch string) ([]string, error) {
	buf, _, err := runCommand(c, "git", "branch", "--merged", targetBranch, `--format=%(refname:short)`)
	if err != nil {
		return nil, err
	}
	return splitString(string(buf), "\n"), nil
}

func IsRemoteBranchExist(c context.Context, remoteBranch string) (bool, error) {
	_, exitCode, err := runCommand(c, "git", "show-ref", "--quiet",
		fmt.Sprintf("refs/remotes/%s", remoteBranch))
	if err != nil {
		if exitCode == 1 {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func DeleteBranch(c context.Context, branchName string) error {
	_, _, err := runCommand(c, "git", "branch", "-d", branchName)
	if err != nil {
		return err
	}
	return nil
}
