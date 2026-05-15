# chain-commit
[![CI](https://github.com/its-the-vibe/chain-commit/actions/workflows/ci.yaml/badge.svg)](https://github.com/its-the-vibe/chain-commit/actions/workflows/ci.yaml)

A Go CLI that reads the staged diff, asks AI for a commit message via LangChainGo, and commits.

## Features
- Generates descriptive commit messages based on staged changes.
- Supports multiple LLM providers (Ollama, Gemini).
- Simple CLI interface.
- Dry-run mode to preview generated messages.

## Installation

```bash
git clone https://github.com/jules/chain-commit.git
cd chain-commit
make build
# Optional: install to your GOPATH
make install
```

## Setup

### Environment Variables

Configure the tool using environment variables:

- `LLM_PROVIDER`: The LLM provider to use (`ollama` or `gemini`). Defaults to `ollama`.
- `OLLAMA_MODEL`: The Ollama model to use (e.g., `llama3`). Defaults to `llama3`.
- `GOOGLE_API_KEY`: Required when using the `gemini` provider.

## Usage

1. Stage your changes:
   ```bash
   git add <files>
   ```

2. Run `chain-commit`:
   ```bash
   # Generate and commit
   ./chain-commit

   # Generate and preview without committing
   ./chain-commit -dry-run
   ```

## Development

- `make build`: Build the binary.
- `make test`: Run tests.
- `make clean`: Remove the binary.
