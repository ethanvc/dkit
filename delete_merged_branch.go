package dkit

import (
	"context"
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
	c := context.Background()
	localBranches, err := dgit.ListLocalBranches(c)
	if err != nil {
		return err
	}
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
	for _, localBranch := range localBranches {
		merged := false
		for productionBranch, _ := range productionBranches {
			merged, err = dgit.IsBranchMerged(c, productionBranch, localBranch)
			if err != nil {
				return err
			}
			if merged {
				break
			}
		}
		if merged && !req.DryRun {
			err = dgit.DeleteBranch(c, localBranch)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
