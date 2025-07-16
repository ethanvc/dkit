package dkit

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path"
	"slices"
	"strings"

	"github.com/ethanvc/dkit/base"
	"github.com/ethanvc/dkit/dgit"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/codes"
)

func AddLintCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "lint",
	}
	targetBranch := cmd.Flags().String("target-branch", "",
		"branch to merge, will do increment check based on this branch")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return RunLintCmd(*targetBranch)
	}
	rootCmd.AddCommand(cmd)
}

func RunLintCmd(targetBranch string) error {
	c := context.Background()
	exist, err := dgit.IsBranchExist(c, targetBranch)
	if err != nil {
		return err
	}
	if !exist {
		return base.NewErr(codes.FailedPrecondition, "TargetBranchNotExist").
			SetMsg("target branch=%s", targetBranch)
	}
	mergeBase, err := dgit.GetMergeBase(c, "HEAD", targetBranch)
	if err != nil {
		return err
	}
	fmt.Printf("merge base is %s\n", mergeBase)
	files, err := dgit.ListAllChangeFiles(c, "HEAD", targetBranch)
	if err != nil {
		return err
	}
	files, err = removeUnrelatedFiles(files)
	if err != nil {
		return err
	}
	printChangedFiles(files)
	err = runLint(c, mergeBase, getDirs(files))
	if err != nil {
		return err
	}
	return nil
}

func runLint(c context.Context, mergeBase string, files []string) error {
	for i := 0; i < len(files); {
		args := []string{"run", "--new-from-rev=" + mergeBase}
		for end := min(len(files), i+100); i < end; i++ {
			args = append(args, files[i])
		}
		cmd := exec.Command("golangci-lint", args...)
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		fmt.Printf("run: %s\n", cmd.String())
		err := cmd.Run()
		if err != nil {
			return err
		}
	}
	return nil
}

func printChangedFiles(files []string) {
	fmt.Println("Changed files:")
	for _, file := range files {
		fmt.Printf("\t%s\n", file)
	}
}

func removeUnrelatedFiles(files []string) ([]string, error) {
	var newFiles []string
	for _, file := range files {
		if !strings.HasSuffix(file, ".go") {
			continue
		}
		_, err := os.Stat(file)
		if err == nil {
			newFiles = append(newFiles, file)
			continue
		}
		if os.IsNotExist(err) {
			continue
		}
		return nil, err
	}
	return newFiles, nil
}

func getDirs(files []string) []string {
	dirDict := make(map[string]bool)
	for _, file := range files {
		dirDict[path.Dir(file)] = true
	}
	dirList := make([]string, 0, len(dirDict))
	for dir := range dirDict {
		dirList = append(dirList, dir)
	}
	slices.Sort(dirList)
	return dirList
}
