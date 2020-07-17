package main

import (
	"cmdbranch"
	"cmdcommit"

	"github.com/spf13/cobra"
)

func main() {
	cmdCommit := cmdcommit.CreateCommitCmd()
	cmdBranch := cmdbranch.CreateBranchCmd()

	rootCmd := &cobra.Command{Use: "igit"}
	rootCmd.AddCommand(cmdCommit)
	rootCmd.AddCommand(cmdBranch)

	rootCmd.Execute()
}
