package git

import (
	"os"
	"os/exec"
	"testing"
)

func setupGitRepo(t *testing.T) string {
	dir, err := os.MkdirTemp("", "git-test-*")
	if err != nil {
		t.Fatal(err)
	}

	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}

	exec.Command("git", "init").Run()
	exec.Command("git", "config", "user.email", "test@example.com").Run()
	exec.Command("git", "config", "user.name", "Test User").Run()

	return dir
}

func TestGitOperations(t *testing.T) {
	origDir, _ := os.Getwd()
	defer os.Chdir(origDir)

	tempDir := setupGitRepo(t)
	defer os.RemoveAll(tempDir)

	// Test HasStagedChanges - should be false
	hasChanges, err := HasStagedChanges()
	if err != nil {
		t.Fatalf("HasStagedChanges failed: %v", err)
	}
	if hasChanges {
		t.Error("Expected no staged changes")
	}

	// Test GetStagedDiff - should be empty
	diff, err := GetStagedDiff()
	if err != nil {
		t.Fatalf("GetStagedDiff failed: %v", err)
	}
	if diff != "" {
		t.Errorf("Expected empty diff, got: %s", diff)
	}

	// Create and stage a file
	err = os.WriteFile("test.txt", []byte("hello"), 0644)
	if err != nil {
		t.Fatal(err)
	}
	exec.Command("git", "add", "test.txt").Run()

	// Test HasStagedChanges - should be true
	hasChanges, err = HasStagedChanges()
	if err != nil {
		t.Fatalf("HasStagedChanges failed: %v", err)
	}
	if !hasChanges {
		t.Error("Expected staged changes")
	}

	// Test GetStagedDiff - should not be empty
	diff, err = GetStagedDiff()
	if err != nil {
		t.Fatalf("GetStagedDiff failed: %v", err)
	}
	if diff == "" {
		t.Error("Expected non-empty diff")
	}

	// Test Commit
	err = Commit("initial commit")
	if err != nil {
		t.Fatalf("Commit failed: %v", err)
	}

	// Test HasStagedChanges - should be false again
	hasChanges, err = HasStagedChanges()
	if err != nil {
		t.Fatalf("HasStagedChanges failed: %v", err)
	}
	if hasChanges {
		t.Error("Expected no staged changes after commit")
	}
}
