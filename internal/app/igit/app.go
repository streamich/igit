package igit

import (
	"cmdcommit"

	"github.com/spf13/cobra"
)

// Run executes igit CLI command.
func Run() {
	var cmdCommit = cmdcommit.CreateCommitCmd()
	var rootCmd = &cobra.Command{Use: "igit"}
	rootCmd.AddCommand(cmdCommit)
	rootCmd.Execute()
}
