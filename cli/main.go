package main

import (
	"github.com/spf13/cobra"
)

func main() {
	var cmdCommit = createCommitCmd()
	var rootCmd = &cobra.Command{Use: "igit"}
	rootCmd.AddCommand(cmdCommit)
	rootCmd.Execute()
}
