package llm

import (
	"context"
	"fmt"
	"os"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/googleai"
	"github.com/tmc/langchaingo/llms/ollama"
)

// Generator is an interface for generating commit messages.
type Generator interface {
	Generate(ctx context.Context, diff string) (string, error)
}

// LangChainGenerator is an implementation of Generator using LangChainGo.
type LangChainGenerator struct {
	llm llms.Model
}

// NewGenerator creates a new Generator based on the LLM_PROVIDER environment variable.
func NewGenerator() (Generator, error) {
	provider := os.Getenv("LLM_PROVIDER")
	if provider == "" {
		provider = "ollama" // Default to ollama
	}

	var llm llms.Model
	var err error

	switch provider {
	case "ollama":
		model := os.Getenv("OLLAMA_MODEL")
		if model == "" {
			model = "llama3"
		}
		llm, err = ollama.New(ollama.WithModel(model))
	case "gemini":
		llm, err = googleai.New(context.Background())
	default:
		return nil, fmt.Errorf("unsupported LLM provider: %s", provider)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to initialize LLM: %w", err)
	}

	return &LangChainGenerator{llm: llm}, nil
}

// Generate generates a commit message from a git diff.
func (g *LangChainGenerator) Generate(ctx context.Context, diff string) (string, error) {
	prompt := fmt.Sprintf("Generate a concise and descriptive git commit message for the following diff. Only return the commit message itself, nothing else.\n\nDiff:\n%s", diff)

	completion, err := llms.GenerateFromSinglePrompt(ctx, g.llm, prompt)
	if err != nil {
		return "", fmt.Errorf("failed to generate commit message: %w", err)
	}

	return completion, nil
}
