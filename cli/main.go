package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var scope bool
	var cmdCommit = &cobra.Command{
		Aliases: []string{"cz", "c"},
		Use:     "commit [title to commit]",
		Short:   "Create a semantic commit",
		Long: `commit prompts you for information to construct
a semantic Angular commit, and Git commits it.`,
		Args: cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			state := CommitState{
				settings: defaultSettings,
				info:     CommitInfo{},
			}

			if scope {
				hasScope := false
				typeIndex := 0
				for i := range state.settings.prompts {
					if state.settings.prompts[i] == "type" {
						typeIndex = i
					}
					if state.settings.prompts[i] == "scope" {
						hasScope = true
						break
					}
				}
				if hasScope == false {
					state.settings.prompts = append(state.settings.prompts[:typeIndex+2], state.settings.prompts[typeIndex + 1:]...)
					state.settings.prompts[typeIndex + 1] = "scope"
				}
			}

			collectUserInput(&state)
			message, _ := formatCommitMessage(&state)

			// Print the generated git message to user so they see what
			// the commited, and we need to print this before actually
			// executing the "git commit ..." command so if error happens
			// user would still see their message. Also, if there are
			// pre-commit hooks running user can read their messsage
			// and decide to cancel the commit.
			fmt.Printf("%s\n", message)

			MakeGitCommit(message)
		},
	}
	cmdCommit.Flags().BoolVarP(&scope, "scope", "s", false, "whether to prompt user for commit scope")

	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmdCommit)
	rootCmd.Execute()
}
