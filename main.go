package main

import (
	"cmdcommit"

	"github.com/spf13/cobra"
)

func main() {
	var cmdCommit = cmdcommit.CreateCommitCmd()
	var rootCmd = &cobra.Command{Use: "igit"}
	rootCmd.AddCommand(cmdCommit)
	rootCmd.Execute()
}
