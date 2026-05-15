package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jules/chain-commit/pkg/git"
	"github.com/jules/chain-commit/pkg/llm"
)

func main() {
	dryRun := flag.Bool("dry-run", false, "Generate commit message without committing")
	flag.Parse()

	// 1. Check for staged changes
	hasChanges, err := git.HasStagedChanges()
	if err != nil {
		log.Fatalf("Error checking for staged changes: %v", err)
	}

	if !hasChanges {
		fmt.Println("No files are staged. Please stage your changes before running chain-commit.")
		os.Exit(0)
	}

	// 2. Get staged diff
	diff, err := git.GetStagedDiff()
	if err != nil {
		log.Fatalf("Error getting staged diff: %v", err)
	}

	// 3. Initialize LLM generator
	generator, err := llm.NewGenerator()
	if err != nil {
		log.Fatalf("Error initializing LLM generator: %v", err)
	}

	// 4. Generate commit message
	ctx := context.Background()
	fmt.Printf("[%s] Generating commit message...\n", generator.ModelName())
	message, err := generator.Generate(ctx, diff)
	if err != nil {
		log.Fatalf("Error generating commit message: %v", err)
	}

	message = strings.TrimSpace(message)
	fmt.Printf("\nGenerated commit message:\n%s\n\n", message)

	if *dryRun {
		fmt.Println("Dry run enabled. Skipping commit.")
		return
	}

	// 5. Commit changes
	err = git.Commit(message)
	if err != nil {
		log.Fatalf("Error committing changes: %v", err)
	}

	fmt.Println("Changes committed successfully!")
}
