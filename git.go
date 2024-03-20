package main

import (
	"fmt"
	"os/exec"
)

func checkGitInPath() error {
	if _, err := exec.LookPath("git"); err != nil {
		return fmt.Errorf("git not found in PATH: [%w]", err)
	}
	return nil
}

func updateGitTags() error {
	cmd := exec.Command("git", "pull", "--tags")
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("git couldn't pull tags: [%w]", err)
	}
	return nil
}

func getGitTags() ([]byte, error) {
	cmd := exec.Command("git", "tag")
	return cmd.Output()
}
