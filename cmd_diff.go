package dkit

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/tidwall/gjson"
	"os"
	"os/exec"
)

func AddDiffCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "diff INPUTS",
		Long: `Compare text differences in three ways:

1. From console input (no arguments):
   $ diff
   [Enter benchmark text]
   [Enter target text]

2. From two text files:
   $ diff benchmark.txt target.txt

3. From a JSON file:
   $ diff diff.json
   (diff.json format: {"benchmark":"text1", "target":"text2"})`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) >= 0 && len(args) <= 2 {
				return nil
			}
			return errors.New("at most two arguments required")
		},
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return diffConsole()
		} else if len(args) == 2 {
			return DiffTwoFile(args[0], args[1])
		} else if len(args) == 1 {
			return DiffDiffFile(args[0])
		} else {
			return errors.New("invalid argument")
		}
	}
	rootCmd.AddCommand(cmd)
}

func DiffDiffFile(diffFile string) error {
	diffContent, err := os.ReadFile(diffFile)
	if err != nil {
		return err
	}
	benchResult := gjson.GetBytes(diffContent, "benchmark")
	if !benchResult.Exists() {
		return errors.New("benchmark not found")
	}
	targetResult := gjson.GetBytes(diffContent, "target")
	if !targetResult.Exists() {
		return errors.New("target not found")
	}
	return DiffString(benchResult.String(), targetResult.String())
}

func DiffTwoFile(benchmarkFile, targetFile string) error {
	benchContent, err := os.ReadFile(benchmarkFile)
	if err != nil {
		return err
	}
	targetContent, err := os.ReadFile(targetFile)
	if err != nil {
		return err
	}
	return DiffString(string(benchContent), string(targetContent))
}

func DiffString(benchContent, targetContent string) error {
	diffCom := NewDiffCompare()
	return diffCom.ShowDiff(benchContent, targetContent)
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
	return DiffString(content1, content2)
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
