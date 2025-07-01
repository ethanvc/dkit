package dkit

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func AddDiffCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "diff [file1.txt file2.txt]",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 0 || len(args) == 2 {
				return nil
			}
			return errors.New("zero or two arguments required")
		},
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return diffConsole()
		} else if len(args) == 2 {
			return DiffTwoFile(args[0], args[1])
		} else {
			return errors.New("invalid argument")
		}
	}
	rootCmd.AddCommand(cmd)
}

func DiffTwoFile(file1, file2 string) error {
	content1, err := os.ReadFile(file1)
	if err != nil {
		return err
	}
	content2, err := os.ReadFile(file2)
	if err != nil {
		return err
	}
	return diffString(string(content1), string(content2))
}

func diffString(content1, content2 string) error {
	diffCom := NewDiffCompare()
	return diffCom.ShowDiff(content1, content2)
}

func diffConsole() error {
	content1, err := readInputFromConsole("read_first_text")
	if err != nil {
		return err
	}
	content2, err := readInputFromConsole("read_second_text")
	if err != nil {
		return err
	}
	return diffString(content1, content2)
}

func readInputFromConsole(title string) (string, error) {
	tmpF, err := os.CreateTemp("", fmt.Sprintf("dkit_temp_%s_*.txt", title))
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpF.Name())
	cmd := exec.Command("vim", "-c", "startinsert", tmpF.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		return "", err
	}
	content, err := os.ReadFile(tmpF.Name())
	if err != nil {
		return "", err
	}
	return string(content), nil
}
