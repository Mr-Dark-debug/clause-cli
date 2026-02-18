# Contributing to Clause CLI

First off, thank you for considering contributing to Clause CLI! It's people like you that make Clause such a great tool.

---

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [Making Changes](#making-changes)
- [Coding Standards](#coding-standards)
- [Testing](#testing)
- [Commit Guidelines](#commit-guidelines)
- [Pull Request Process](#pull-request-process)
- [Reporting Issues](#reporting-issues)

---

## Code of Conduct

This project and everyone participating in it is governed by our Code of Conduct. By participating, you are expected to uphold this code. Please report unacceptable behavior to team@clause.dev.

### Our Standards

- Be respectful and inclusive
- Welcome newcomers and help them get started
- Focus on what is best for the community
- Show empathy towards other community members

---

## Getting Started

### Prerequisites

- **Go 1.21+**: For building the CLI
- **Git**: For version control
- **Make**: For build commands (optional but recommended)

### Quick Start

```bash
# Fork and clone the repository
git clone https://github.com/YOUR_USERNAME/clause.git
cd clause

# Install dependencies
make deps

# Build the binary
make build

# Run tests
make test

# Run the CLI
./bin/clause version
```

---

## Development Setup

### 1. Fork and Clone

```bash
# Fork the repo on GitHub, then:
git clone https://github.com/YOUR_USERNAME/clause.git
cd clause

# Add upstream remote
git remote add upstream https://github.com/clause-cli/clause.git
```

### 2. Install Dependencies

```bash
# Download Go modules
go mod download

# Verify dependencies
go mod verify
```

### 3. Build

```bash
# Build for your platform
make build

# Or build for all platforms
make build-all
```

### 4. Development Mode

```bash
# Install air for hot reload (optional)
go install github.com/air-verse/air@latest

# Run with hot reload
make dev
```

### 5. IDE Setup

#### VS Code

Recommended extensions:
- Go (golang.go)
- YAML (redhat.vscode-yaml)
- Markdown All in One

#### GoLand/IntelliJ

Enable Go modules support and configure:
- Go 1.21+ SDK
- Enable goimports on save
- Enable golint

---

## Project Structure

```
clause/
├── cmd/
│   └── clause/           # Main entry point
│       └── main.go
├── internal/             # Private application code
│   ├── cmd/              # CLI commands
│   ├── config/           # Configuration management
│   ├── generator/        # Project generation
│   ├── governance/       # AI governance system
│   ├── template/         # Template engine
│   └── wizard/           # Interactive wizard
│       ├── screens/      # Wizard screens
│       └── components/   # UI components
├── pkg/                  # Public packages (can be imported)
│   ├── output/           # Output formatting
│   ├── styles/           # Styling and theming
│   ├── tui/              # Terminal UI utilities
│   └── utils/            # General utilities
├── templates/            # Template files
│   └── governance/       # Governance templates
├── docs/                 # Documentation
├── bin/                  # Compiled binaries (gitignored)
└── dist/                 # Distribution packages (gitignored)
```

### Package Guidelines

- **`cmd/`**: Contains only `main.go` - minimal entry point
- **`internal/`**: Private code, cannot be imported by external projects
- **`pkg/`**: Public code that can be imported by external projects
- **`templates/`**: Embedded template files

---

## Making Changes

### 1. Create a Branch

```bash
# Sync with upstream
git fetch upstream
git checkout main
git merge upstream/main

# Create a feature branch
git checkout -b feature/my-new-feature
```

### 2. Make Your Changes

- Write clean, readable code
- Follow the [coding standards](#coding-standards)
- Add tests for new functionality
- Update documentation

### 3. Test Your Changes

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage

# Run linter
make lint

# Format code
make fmt
```

### 4. Commit Your Changes

```bash
git add .
git commit -m "feat: add amazing feature"
```

---

## Coding Standards

### Go Style Guide

We follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments) and [Effective Go](https://golang.org/doc/effective_go).

#### Key Points

1. **Formatting**
   ```bash
   # Format code before committing
   gofmt -w .
   ```

2. **Imports**
   ```go
   // Group imports:
   // 1. Standard library
   // 2. Third-party packages
   // 3. Local packages
   import (
       "context"
       "fmt"

       "github.com/spf13/cobra"
       "github.com/clause-cli/clause/pkg/utils"
   )
   ```

3. **Error Handling**
   ```go
   // Always handle errors explicitly
   if err != nil {
       return fmt.Errorf("failed to do something: %w", err)
   }
   ```

4. **Comments**
   ```go
   // FunctionName does something.
   // It takes x and returns y.
   //
   // Example:
   //   result := FunctionName("input")
   func FunctionName(x string) string {
       // Implementation
   }
   ```

5. **Naming**
   ```go
   // Use descriptive names
   func calculateTotalPrice(items []Item) float64 {
       totalPrice := 0.0
       for _, item := range items {
           totalPrice += item.Price
       }
       return totalPrice
   }
   ```

### File Organization

Keep files small and focused:
- Aim for 150-300 lines per file
- One primary type/function per file when possible
- Group related functionality in subpackages

### Code Quality

Run these before committing:

```bash
# Format
go fmt ./...

# Vet
go vet ./...

# Lint (requires golangci-lint)
golangci-lint run ./...
```

---

## Testing

### Running Tests

```bash
# Run all tests
make test

# Run specific package tests
go test ./internal/generator/...

# Run with verbose output
go test -v ./...

# Run with coverage
go test -cover ./...

# Generate coverage report
make test-coverage
```

### Writing Tests

1. **Test File Naming**
   - Test files should end with `_test.go`
   - Place in the same package as the code being tested

2. **Test Structure (AAA Pattern)**
   ```go
   func TestSomething(t *testing.T) {
       // Arrange
       input := "test"
       expected := "TEST"

       // Act
       result := strings.ToUpper(input)

       // Assert
       if result != expected {
           t.Errorf("expected %s, got %s", expected, result)
       }
   }
   ```

3. **Table-Driven Tests**
   ```go
   func TestCalculatePrice(t *testing.T) {
       tests := []struct {
           name     string
           quantity int
           price    float64
           expected float64
       }{
           {"single item", 1, 10.0, 10.0},
           {"multiple items", 5, 10.0, 50.0},
           {"zero quantity", 0, 10.0, 0.0},
       }

       for _, tt := range tests {
           t.Run(tt.name, func(t *testing.T) {
               result := CalculatePrice(tt.quantity, tt.price)
               if result != tt.expected {
                   t.Errorf("expected %f, got %f", tt.expected, result)
               }
           })
       }
   }
   ```

4. **Coverage Requirements**
   - New code should have at least 80% coverage
   - Critical paths should have 100% coverage

---

## Commit Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/).

### Format

```
<type>(<scope>): <subject>

[optional body]

[optional footer(s)]
```

### Types

| Type | Description |
|------|-------------|
| `feat` | New feature |
| `fix` | Bug fix |
| `docs` | Documentation changes |
| `style` | Code style changes (formatting, etc.) |
| `refactor` | Code refactoring |
| `perf` | Performance improvements |
| `test` | Adding/updating tests |
| `chore` | Maintenance tasks |
| `ci` | CI/CD changes |

### Examples

```bash
# Feature
git commit -m "feat(wizard): add support for Vue 3"

# Bug fix
git commit -m "fix(generator): correct Dockerfile template path"

# Breaking change
git commit -m "feat(api)!: change config file format to YAML

BREAKING CHANGE: Configuration files must now use YAML format"
```

### Commit Message Rules

1. Use the imperative mood ("add feature" not "added feature")
2. Limit the subject line to 72 characters
3. Separate subject from body with a blank line
4. Use the body to explain what and why, not how

---

## Pull Request Process

### 1. Push Your Changes

```bash
git push origin feature/my-new-feature
```

### 2. Create a Pull Request

1. Go to [GitHub](https://github.com/clause-cli/clause/pulls)
2. Click "New Pull Request"
3. Select your branch
4. Fill in the PR template

### 3. PR Requirements

- [ ] All tests pass
- [ ] Code follows style guidelines
- [ ] New code has tests
- [ ] Documentation is updated
- [ ] Commit messages follow conventions

### 4. PR Review Process

1. At least one approval required
2. All CI checks must pass
3. No merge conflicts
4. PR must be up to date with main

### 5. After Approval

- Squash and merge is preferred
- Delete your branch after merge

---

## Reporting Issues

### Before Creating an Issue

1. Search existing issues
2. Try the latest version
3. Gather relevant information

### Issue Template

```markdown
**Description**
[Clear description of the issue]

**Steps to Reproduce**
1. [First step]
2. [Second step]
3. [Expected behavior]

**Expected Behavior**
[What should happen]

**Actual Behavior**
[What actually happened]

**Environment**
- OS: [e.g., macOS 14]
- Clause version: [e.g., v1.0.0]
- Go version: [e.g., 1.21.0]

**Additional Context**
[Screenshots, logs, or other relevant information]
```

### Feature Requests

```markdown
**Is your feature request related to a problem?**
[Description of the problem]

**Describe the Solution**
[Description of what you want to happen]

**Describe Alternatives**
[Description of alternative solutions]

**Additional Context**
[Any other context or screenshots]
```

---

## Getting Help

- **GitHub Discussions**: [github.com/clause-cli/clause/discussions](https://github.com/clause-cli/clause/discussions)
- **Discord**: [Join our server](https://discord.gg/clause)
- **Email**: team@clause.dev

---

## Recognition

Contributors are recognized in:
- Our [Contributors page](https://github.com/clause-cli/clause/graphs/contributors)
- Release notes for significant contributions
- Annual contributor spotlight

---

Thank you for contributing to Clause CLI! 🎉
