# Makefile - Cross-platform version

.PHONY: all build install clean test deps

# Detect OS
ifeq ($(OS),Windows_NT)
    SHELL := powershell.exe
    BIN_DIR := bin
    EXE := clause.exe
    RM := Remove-Item -Recurse -Force
    MKDIR := New-Item -ItemType Directory -Force
else
    BIN_DIR := bin
    EXE := clause
    RM := rm -rf
    MKDIR := mkdir -p
endif

all: deps build

deps:
	@echo "Installing dependencies..."
	go mod download
	go mod tidy

build:
	@echo "Building clause..."
	@if not exist "$(BIN_DIR)" mkdir "$(BIN_DIR)" 2>nul || echo "" >nul
	go build -o $(BIN_DIR)/$(EXE) ./cmd/clause

install: build
	@echo "Installing clause..."
	go install ./cmd/clause

clean:
	@echo "Cleaning..."
	@if exist "$(BIN_DIR)" rmdir /s /q "$(BIN_DIR)" 2>nul || rm -rf $(BIN_DIR) 2>nul || true

test:
	@echo "Running tests..."
	go test -v ./...

run: build
	@echo "Running clause..."
	./$(BIN_DIR)/$(EXE)

fmt:
	go fmt ./...

lint:
	golangci-lint run

.PHONY: all build install clean test deps run fmt lint
