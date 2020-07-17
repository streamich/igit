package cmdbranch

import (
	"fmt"

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

	fmt.Printf("\nbranching...\n")

}
