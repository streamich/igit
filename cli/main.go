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
	MakeGitCommit(message)

	fmt.Printf("Message: %q\n", message)
}
