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

func ListAllBranches(c context.Context) ([]RefBranchName, error) {
	buf, _, err := runCommand(c, "git", "branch", "-a", "--format=%(refname)")
	if err != nil {
		return nil, err
	}
	branches := splitString(string(buf), "\n")
	refBranches := make([]RefBranchName, 0, len(branches))
	for _, branch := range branches {
		refBranches = append(refBranches, RefBranchName(branch))
	}
	return refBranches, nil
}

type RefBranchName string

func (ref RefBranchName) BranchName() string {
	name := strings.TrimPrefix(string(ref), "refs/heads/")
	name = strings.TrimPrefix(name, "refs/remotes/")
	return name
}

func IsBranchExist(c context.Context, name string) (bool, error) {
	branchList, err := ListAllBranches(c)
	if err != nil {
		return false, err
	}
	for _, branch := range branchList {
		if branch.BranchName() == name {
			return true, nil
		}
	}
	return false, nil
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

func GetMergeBase(c context.Context, srcRef, dstRef string) (string, error) {
	buf, _, err := runCommand(c, "git", "merge-base", srcRef, dstRef)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(buf)), nil
}

// ListAllChangeFiles return path relative to git root
func ListAllChangeFiles(c context.Context, srcRef, dstRef string) ([]string, error) {
	buf, _, err := runCommand(c, "git", "diff", "--name-only", fmt.Sprintf("%s...%s", dstRef, srcRef))
	if err != nil {
		return nil, err
	}
	files := splitString(string(buf), "\n")
	return files, nil
}
