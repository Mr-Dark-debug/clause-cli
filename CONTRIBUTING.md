# Contributing to Clause CLI

Thank you for your interest in contributing to Clause CLI! This document provides guidelines and instructions for contributing.

## 📜 Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment. Please be considerate of others and follow standard open-source community guidelines.

## 🐛 Reporting Bugs

Before creating a bug report, please check existing issues to avoid duplicates.

### Bug Report Template

```markdown
**Description**
A clear description of the bug.

**Steps to Reproduce**
1. Run `clause init test`
2. Select options X, Y, Z
3. Observe error

**Expected Behavior**
What should happen instead.

**Actual Behavior**
What actually happens.

**Environment**
- OS: [e.g., macOS 14]
- Clause version: [e.g., 1.0.0]
- Terminal: [e.g., iTerm2]

**Screenshots**
If applicable, add screenshots.

**Additional Context**
Any other relevant information.
```

## 💡 Requesting Features

Feature requests are welcome! Please use the feature request template:

```markdown
**Is your feature request related to a problem?**
A clear description of the problem.

**Describe the solution you'd like**
A clear description of what you want to happen.

**Describe alternatives you've considered**
Any alternative solutions or features you've considered.

**Additional context**
Any other context, screenshots, or examples.
```

## 🔧 Development Setup

### Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile)
- Git

### Getting Started

```bash
# Fork and clone
git clone https://github.com/YOUR_USERNAME/clause-cli.git
cd clause-cli

# Create a branch
git checkout -b feature/your-feature-name

# Install dependencies
go mod download

# Build
make build
# or
go build -o bin/clause ./cmd/forge

# Run tests
make test
# or
go test ./...

# Run locally
./bin/clause init test-project
```

### Project Structure

```
clause-cli/
├── cmd/forge/main.go        # Entry point
├── internal/                # Private packages
│   ├── cmd/                 # CLI commands
│   ├── wizard/              # TUI wizard
│   ├── config/              # Configuration
│   ├── template/            # Template engine
│   ├── generator/           # Project generation
│   └── governance/          # AI governance
├── pkg/                     # Public packages
│   ├── styles/              # Styling
│   ├── tui/                 # Terminal UI
│   └── utils/               # Utilities
└── templates/               # Embedded templates
```

## 📝 Coding Standards

### Go Style

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Use `golint` for linting
- Run `go vet` before committing

### Code Organization

- Keep files under 200 lines
- One responsibility per file
- Use clear, descriptive names
- Add documentation comments

### Documentation

```go
// Package example provides example functionality.
//
// This package demonstrates documentation standards.
// Use it as a reference for documenting your code.
package example

// FunctionName does something specific.
//
// More detailed description if needed.
//
// Example:
//
//   result := FunctionName(arg1, arg2)
//   fmt.Println(result)
func FunctionName(arg1, arg2 Type) ReturnType {
    // implementation
}
```

### Testing

- Write tests for all new functionality
- Aim for 80%+ coverage on new code
- Use table-driven tests for multiple cases

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    Type
        expected ReturnType
        wantErr  bool
    }{
        {
            name:     "valid input",
            input:    validInput,
            expected: expectedOutput,
            wantErr:  false,
        },
        {
            name:     "invalid input",
            input:    invalidInput,
            expected: zeroValue,
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := FunctionName(tt.input)
            if tt.wantErr {
                assert.Error(t, err)
                return
            }
            assert.NoError(t, err)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

## 📋 Commit Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding/updating tests
- `chore`: Maintenance tasks

### Examples

```
feat(wizard): add backend framework selection screen

Add a new wizard screen for selecting backend frameworks
including FastAPI, Express, and Django.

Closes #123
```

```
fix(generator): handle empty project name correctly

Fix a panic that occurred when project name was empty.
Now shows a validation error instead.
```

## 🔀 Pull Request Process

1. **Fork and Branch**: Create a feature branch from `main`
2. **Make Changes**: Implement your changes with tests
3. **Run Tests**: Ensure all tests pass
4. **Update Docs**: Update documentation if needed
5. **Commit**: Use conventional commit messages
6. **Push**: Push to your fork
7. **Open PR**: Create a pull request

### PR Checklist

- [ ] Code compiles correctly
- [ ] All tests pass
- [ ] New code has tests
- [ ] Documentation updated
- [ ] Commit messages follow conventions
- [ ] PR description is clear

### PR Template

```markdown
## Description
Brief description of changes.

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
Describe how you tested these changes.

## Checklist
- [ ] Tests pass
- [ ] Code follows style guidelines
- [ ] Documentation updated
```

## 🏗️ Release Process

1. Version bumps follow [SemVer](https://semver.org/)
2. Releases are automated via GoReleaser
3. Changelog is updated automatically

## 📞 Getting Help

- Open a GitHub Discussion for questions
- Join our Discord for real-time help
- Check existing issues/PRs before creating new ones

---

Thank you for contributing! 🎉
