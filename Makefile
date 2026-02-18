.PHONY: all build build-all test test-coverage lint fmt install clean run help

# Binary names
BINARY_NAME=clause
BINARY_UNIX=$(BINARY_NAME)_unix

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Build flags
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS=-ldflags "-s -w -X github.com/clause-cli/clause/internal/cmd.version=$(VERSION) -X github.com/clause-cli/clause/internal/cmd.buildTime=$(BUILD_TIME) -X github.com/clause-cli/clause/internal/cmd.commit=$(COMMIT)"

# Main build target
all: clean deps build

# Build for current platform
build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p bin
	$(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME) ./cmd/clause
	@echo "Build complete: bin/$(BINARY_NAME)"

# Build for all platforms
build-all:
	@echo "Building for all platforms..."
	@mkdir -p bin
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 ./cmd/clause
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-arm64 ./cmd/clause
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 ./cmd/clause
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-arm64 ./cmd/clause
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe ./cmd/clause
	GOOS=windows GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-arm64.exe ./cmd/clause
	@echo "All builds complete in bin/"

# Run tests
test:
	$(GOTEST) -v -race ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -race -coverprofile=coverage.out -covermode=atomic ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linter
lint:
	@which golangci-lint > /dev/null || go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run ./...

# Format code
fmt:
	$(GOCMD) fmt ./...

# Install dependencies
deps:
	$(GOMOD) download
	$(GOMOD) verify

# Install locally
install: build
	@echo "Installing $(BINARY_NAME)..."
	@cp bin/$(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME) 2>/dev/null || sudo cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
	@echo "Installed $(BINARY_NAME)"

# Clean build artifacts
clean:
	@echo "Cleaning..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html
	$(GOCLEAN)

# Run the binary
run: build
	@./bin/$(BINARY_NAME)

# Create release
release:
	@echo "Creating release..."
	@which goreleaser > /dev/null || go install github.com/goreleaser/goreleaser/v2@latest
	goreleaser release --clean

# Local release snapshot
snapshot:
	@echo "Creating snapshot release..."
	@which goreleaser > /dev/null || go install github.com/goreleaser/goreleaser/v2@latest
	goreleaser release --snapshot --clean

# Development with hot reload
dev:
	@which air > /dev/null || go install github.com/air-verse/air@latest
	air

# Help
help:
	@echo "Clause CLI - Build System"
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@echo "  all           Clean, download deps, and build (default)"
	@echo "  build         Build for current platform"
	@echo "  build-all     Build for all platforms"
	@echo "  test          Run tests"
	@echo "  test-coverage Run tests with coverage report"
	@echo "  lint          Run linter"
	@echo "  fmt           Format code"
	@echo "  deps          Download dependencies"
	@echo "  install       Install binary locally"
	@echo "  clean         Remove build artifacts"
	@echo "  run           Build and run the binary"
	@echo "  release       Create a release (requires tag)"
	@echo "  snapshot      Create a snapshot release"
	@echo "  dev           Run with hot reload"
	@echo "  help          Show this help message"
