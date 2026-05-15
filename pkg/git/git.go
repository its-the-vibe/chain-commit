package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// GetStagedDiff returns the diff of the staged changes.
func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "--no-pager", "diff", "--cached")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get staged diff: %w", err)
	}
	return out.String(), nil
}

// HasStagedChanges checks if there are any staged changes.
func HasStagedChanges() (bool, error) {
	cmd := exec.Command("git", "diff", "--cached", "--quiet")
	err := cmd.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			if exitError.ExitCode() == 1 {
				return true, nil
			}
		}
		return false, fmt.Errorf("failed to check for staged changes: %w", err)
	}
	return false, nil
}

// Commit creates a commit with the given message.
func Commit(message string) error {
	cmd := exec.Command("git", "commit", "--message", message)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to commit changes: %s: %w", strings.TrimSpace(stderr.String()), err)
	}
	return nil
}
