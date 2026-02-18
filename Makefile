.PHONY: all build build-windows build-all test test-coverage lint fmt install clean run help

# Binary names
BINARY_NAME=clause

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod

# Detect OS
ifeq ($(OS),Windows_NT)
	BINARY_EXT=.exe
	RM=del /q
	RMDIR=rmdir /s /q
	MKDIR=if not exist
else
	BINARY_EXT=
	RM=rm -f
	RMDIR=rm -rf
	MKDIR=mkdir -p
endif

# Build flags
VERSION?=$(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
BUILD_TIME=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ" 2>/dev/null || date /t)
COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
LDFLAGS=-ldflags "-s -w -X github.com/clause-cli/clause/internal/cmd.version=$(VERSION) -X github.com/clause-cli/clause/internal/cmd.buildTime=$(BUILD_TIME) -X github.com/clause-cli/clause/internal/cmd.commit=$(COMMIT)"

# Main build target
all: clean deps build

# Build for current platform
build:
	@echo "Building $(BINARY_NAME)..."
ifeq ($(OS),Windows_NT)
	@if not exist bin mkdir bin
else
	@mkdir -p bin
endif
	$(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)$(BINARY_EXT) ./cmd/clause
	@echo "Build complete: bin/$(BINARY_NAME)$(BINARY_EXT)"

# Build for Windows specifically
build-windows:
	@echo "Building for Windows..."
	@if not exist bin mkdir bin
	$(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME).exe ./cmd/clause
	@echo "Build complete: bin/$(BINARY_NAME).exe"

# Build for all platforms
build-all:
	@echo "Building for all platforms..."
ifeq ($(OS),Windows_NT)
	@if not exist bin mkdir bin
else
	@mkdir -p bin
endif
	GOOS=darwin GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-amd64 ./cmd/clause
	GOOS=darwin GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-darwin-arm64 ./cmd/clause
	GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-amd64 ./cmd/clause
	GOOS=linux GOARCH=arm64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-linux-arm64 ./cmd/clause
	GOOS=windows GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o bin/$(BINARY_NAME)-windows-amd64.exe ./cmd/clause
	@echo "All builds complete in bin/"

# Run tests
test:
	$(GOTEST) -v ./...

# Run tests with coverage
test-coverage:
	$(GOTEST) -v -coverprofile=coverage.out ./...
	$(GOCMD) tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Run linter
lint:
	$(GOCMD) vet ./...

# Format code
fmt:
	$(GOCMD) fmt ./...

# Install dependencies
deps:
	$(GOMOD) download

# Install locally (Windows)
install: build
	@echo "Installing $(BINARY_NAME)..."
ifeq ($(OS),Windows_NT)
	@copy bin\$(BINARY_NAME).exe $(GOPATH)\bin\$(BINARY_NAME).exe >nul 2>&1 || echo "Could not copy to GOPATH/bin. Run: go install ./cmd/clause"
else
	@cp bin/$(BINARY_NAME) $(GOPATH)/bin/$(BINARY_NAME) 2>/dev/null || sudo cp bin/$(BINARY_NAME) /usr/local/bin/$(BINARY_NAME)
endif
	@echo "Installed $(BINARY_NAME)"

# Install using go install
install-go:
	$(GOCMD) install ./cmd/clause

# Clean build artifacts
clean:
	@echo "Cleaning..."
ifeq ($(OS),Windows_NT)
	@if exist bin rmdir /s /q bin
	@if exist coverage.out del /q coverage.out
	@if exist coverage.html del /q coverage.html
else
	@rm -rf bin/
	@rm -f coverage.out coverage.html
endif
	$(GOCLEAN)

# Run the binary
run: build
ifeq ($(OS),Windows_NT)
	@.\bin\$(BINARY_NAME).exe
else
	@./bin/$(BINARY_NAME)
endif

# Create release
release:
	@echo "Creating release..."
	goreleaser release --clean

# Local release snapshot
snapshot:
	@echo "Creating snapshot release..."
	goreleaser release --snapshot --clean

# Development with hot reload
dev:
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
	@echo "  build-windows Build for Windows"
	@echo "  build-all     Build for all platforms"
	@echo "  test          Run tests"
	@echo "  test-coverage Run tests with coverage report"
	@echo "  lint          Run linter"
	@echo "  fmt           Format code"
	@echo "  deps          Download dependencies"
	@echo "  install       Install binary locally"
	@echo "  install-go    Install using go install"
	@echo "  clean         Remove build artifacts"
	@echo "  run           Build and run the binary"
	@echo "  release       Create a release (requires tag)"
	@echo "  snapshot      Create a snapshot release"
	@echo "  dev           Run with hot reload"
	@echo "  help          Show this help message"
