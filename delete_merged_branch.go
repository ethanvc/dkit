package dkit

import (
	"context"
	"fmt"

	"github.com/ethanvc/dkit/dgit"
	"github.com/spf13/cobra"
)

func AddDeleteMergedBranchCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "delete-merged-branch",
	}
	dryRunFlag := cmd.Flags().Bool("dry-run", false, "dry run")
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		return DeleteMergedBranch(&DeleteMergedBranchReq{
			DryRun: *dryRunFlag,
		})
	}
	rootCmd.AddCommand(cmd)
}

type DeleteMergedBranchReq struct {
	DryRun bool
}

func DeleteMergedBranch(req *DeleteMergedBranchReq) error {
	if req.DryRun {
		fmt.Printf("Notice: dry run mode\n")
	}
	c := context.Background()
	productionBranches := map[string]struct{}{
		"origin/master":  {},
		"origin/main":    {},
		"origin/release": {},
	}
	for k, _ := range productionBranches {
		exist, err := dgit.IsRemoteBranchExist(c, k)
		if err != nil {
			return err
		}
		if !exist {
			delete(productionBranches, k)
		}
	}
	mergedBranches := make(map[string]struct{})
	for targetBranch := range productionBranches {
		branches, err := dgit.ListMergedBranches(c, targetBranch)
		if err != nil {
			return err
		}
		for _, branch := range branches {
			mergedBranches[branch] = struct{}{}
		}
	}
	for branch, _ := range mergedBranches {
		if !req.DryRun {
			err := dgit.DeleteBranch(c, branch)
			if err != nil {
				return err
			}
		}
		fmt.Printf("delete local branch: %s\n", branch)
	}
	return nil
}
