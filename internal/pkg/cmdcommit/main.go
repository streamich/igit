package cmdcommit

import (
	"fmt"

	"gitutils"

	"github.com/spf13/cobra"
)

var description = `commit prompts you for information to construct
a semantic Angular commit, and Git commits it.`

// CreateCommitCmd creates Cobra CLI command for semantic commit prompt.
func CreateCommitCmd() *cobra.Command {
	flags := commitFlags{}
	cmd := &cobra.Command{
		Aliases: []string{"cz", "c"},
		Use:     "commit [title to commit]",
		Short:   "Create a semantic commit",
		Long:    description,
		Args:    cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			flags.args = args
			execCommitCmd(&flags)
		},
	}

	cmd.Flags().BoolVarP(&(flags.scope), "scope", "s", false, "whether to prompt user for commit scope")

	return cmd
}

func addScopeToPromptList(state *CommitState) {
	if state.flags.scope == false {
		return
	}

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
		state.settings.prompts = append(state.settings.prompts[:typeIndex+2], state.settings.prompts[typeIndex+1:]...)
		state.settings.prompts[typeIndex+1] = "scope"
	}
}

func execCommitCmd(flags *commitFlags) {
	state := &CommitState{
		flags:    flags,
		settings: defaultSettings,
		info:     CommitInfo{},
	}

	addScopeToPromptList(state)
	collectUserInput(state)
	message, _ := formatCommitMessage(state)

	// Print the generated git message to user so they see what
	// the committed, and we need to print this before actually
	// executing the "git commit ..." command so if error happens
	// user would still see their message. Also, if there are
	// pre-commit hooks running user can read their messsage
	// and decide to cancel the commit.
	fmt.Printf("\n%s\n", message)

	gitutils.MakeGitCommit(message)
}
