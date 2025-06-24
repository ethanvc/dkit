package dkit

import "github.com/spf13/cobra"

func AddDeleteMergedBranchCmd(rootCmd *cobra.Command) {
	cmd := &cobra.Command{
		Use: "delete-merged-branch",
	}
	dryRunFlag := cmd.Flags().Bool("dry-run", false, "dry run")
	cmd.Run = func(cmd *cobra.Command, args []string) {
		DeleteMergedBranch(&DeleteMergedBranchReq{
			DryRun: *dryRunFlag,
		})
	}
	rootCmd.AddCommand(cmd)
}

type DeleteMergedBranchReq struct {
	DryRun bool
}

func DeleteMergedBranch(req *DeleteMergedBranchReq) {

}
