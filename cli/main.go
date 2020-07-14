package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	state := CommitState{
		settings: defaultSettings,
		info:     CommitInfo{},
	}

	collectUserInput(&state)
	message, _ := formatCommitMessage(&state)

	cmd := exec.Command("git", "commit", "-m", message)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}

	fmt.Printf("Message: %q\n", message)
}
