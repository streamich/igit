package main

import (
	"fmt"
)

func main() {

	state := CommitState{
		settings: defaultSettings,
		info:     CommitInfo{},
	}

	collectUserInput(&state)
	message, _ := formatCommitMessage(&state)

	// Print the generate git message to user so they see what
	// the commited, and we need to print this before actually
	// executing the "git commit ..." command so if error happens
	// user would still see their message. Also, if there are
	// pre-commit hooks running user can read their messsage
	// and decide to cancel the commit.
	fmt.Printf("%s\n", message)

	MakeGitCommit(message)
}
