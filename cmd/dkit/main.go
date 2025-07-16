package main

import (
	"fmt"
	"os"

	"github.com/ethanvc/dkit"
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "dkit",
		Short: "dkit",
		Long:  `dkit`,
	}

	dkit.AddDeleteMergedBranchCmd(rootCmd)
	dkit.AddDiffCmd(rootCmd)
	dkit.AddLintCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
