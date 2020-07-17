package igit

import (
	"github.com/spf13/cobra"
)

// Run executes igit CLI command.
func Run() {
	var cmdCommit = createCommitCmd()
	var rootCmd = &cobra.Command{Use: "igit"}
	rootCmd.AddCommand(cmdCommit)
	rootCmd.Execute()
}
