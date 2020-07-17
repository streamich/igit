package cmdbranch

import (
	"fmt"
	"gitutils"

	"github.com/spf13/cobra"
)

var description = `Creates a new semantic Git branch.`

// CreateBranchCmd creates Cobra CLI command for semantic commit prompt.
func CreateBranchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Aliases: []string{"bz", "b"},
		Use:     "branch",
		Short:   "Create a semantic branch",
		Long:    description,
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			execBranchCmd()
		},
	}

	return cmd
}

func execBranchCmd() {
	state := &branchState{
		settings: defaultSettings,
		info:     branchInfo{},
	}

	collectUserInput(state)
	name := formatBranchName(state)
	fmt.Printf("\nCreating branch: %s\n", name)
	gitutils.MakeBranch(name)
}
