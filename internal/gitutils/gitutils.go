package gitutils

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

// CreateBranch creates a new Git branch.
func CreateBranch(name string) error {
	cmd := exec.Command("git", "checkout", "-b", name)
	return cmd.Run()
}

// MakeBranch creates a new Git branch or exits.
func MakeBranch(name string) {
	err := CreateBranch(name)
	if err != nil {
		log.Fatalf("Error: %s\n", err)
	}
}
