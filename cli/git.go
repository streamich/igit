package main

import (
	"log"
	"os/exec"
)

func gitCommit(message string) error {
	cmd := exec.Command("git", "commit", "-m", message)
	return cmd.Run()
}

// MakeGitCommit executes "git commit -m ..." command.
func MakeGitCommit(message string) {
	err := gitCommit(message)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}
