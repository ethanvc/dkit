package dkit

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func AddDiffCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "diff",
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return CmdDiff()
	}
	rootCmd.AddCommand(cmd)
}

func CmdDiff() error {
	content1, err := readInputFromConsole("read_first_text")
	if err != nil {
		return err
	}
	content2, err := readInputFromConsole("read_second_text")
	if err != nil {
		return err
	}
	com := NewDiffCompare()
	return com.ShowDiff(content1, content2)
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
