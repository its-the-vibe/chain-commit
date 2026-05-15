BINARY_NAME=chain-commit

.PHONY: all build test clean install

all: build

build:
	go build -o $(BINARY_NAME) ./cmd/chain-commit

test:
	go test ./...

clean:
	rm -f $(BINARY_NAME)

install:
	go install ./cmd/chain-commit
