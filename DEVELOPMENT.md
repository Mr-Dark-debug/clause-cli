# Development Guide

This guide provides detailed instructions for developing Clause CLI.

---

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [Development Environment](#development-environment)
- [Building](#building)
- [Debugging](#debugging)
- [Code Organization](#code-organization)
- [Package Development](#package-development)
- [Template Development](#template-development)
- [Testing Strategy](#testing-strategy)
- [Release Process](#release-process)

---

## Architecture Overview

### High-Level Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLI Layer                                │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                    cmd/clause/main.go                     │   │
│  └─────────────────────────────────────────────────────────┘   │
│                              │                                   │
│                              ▼                                   │
│  ┌─────────────────────────────────────────────────────────┐   │
│  │                    internal/cmd/                          │   │
│  │    Commands: init, add, validate, config, update         │   │
│  └─────────────────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                       Core Layer                                 │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐         │
│  │   Wizard     │  │  Generator   │  │  Governance  │         │
│  │  (UI Flow)   │  │ (Scaffold)   │  │  (AI Rules)  │         │
│  └──────────────┘  └──────────────┘  └──────────────┘         │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐         │
│  │   Config     │  │   Template   │  │              │         │
│  │  (Settings)  │  │   (Engine)   │  │              │         │
│  └──────────────┘  └──────────────┘  └──────────────┘         │
└─────────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────────┐
│                      Utility Layer                               │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐         │
│  │    Styles    │  │     TUI      │  │   Output     │         │
│  │ (Theming)    │  │ (Terminal)   │  │ (Formatting) │         │
│  └──────────────┘  └──────────────┘  └──────────────┘         │
│  ┌──────────────┐                                               │
│  │    Utils     │                                               │
│  │ (Helpers)    │                                               │
│  └──────────────┘                                               │
└─────────────────────────────────────────────────────────────────┘
```

### Key Components

| Component | Location | Responsibility |
|-----------|----------|----------------|
| Commands | `internal/cmd/` | CLI command implementations |
| Wizard | `internal/wizard/` | Interactive project setup |
| Generator | `internal/generator/` | Project file generation |
| Governance | `internal/governance/` | AI governance management |
| Config | `internal/config/` | Configuration handling |
| Templates | `internal/template/` | Template rendering |
| Styles | `pkg/styles/` | Color theming and typography |
| TUI | `pkg/tui/` | Terminal UI utilities |
| Output | `pkg/output/` | Formatted console output |
| Utils | `pkg/utils/` | General utilities |

---

## Development Environment

### Required Tools

| Tool | Version | Purpose |
|------|---------|---------|
| Go | 1.21+ | Build and run |
| Git | 2.x | Version control |
| Make | Any | Build automation |
| golangci-lint | 1.55+ | Linting |

### Recommended Tools

| Tool | Purpose |
|------|---------|
| Air | Hot reload development |
| gopls | Go language server |
| delve | Go debugger |

### Environment Variables

```bash
# Development mode
export CLAUSE_DEV=true

# Log level
export CLAUSE_LOG_LEVEL=debug

# Config file path
export CLAUSE_CONFIG=~/.clause/config.yaml
```

### IDE Configuration

#### VS Code (settings.json)

```json
{
  "go.useLanguageServer": true,
  "go.lintTool": "golangci-lint",
  "go.lintOnSave": "package",
  "go.formatTool": "gofmt",
  "editor.formatOnSave": true,
  "[go]": {
    "editor.codeActionsOnSave": {
      "source.organizeImports": true
    }
  }
}
```

#### GoLand

1. Enable Go Modules
2. Set Go SDK to 1.21+
3. Configure File Watchers for:
   - gofmt
   - goimports
   - golangci-lint

---

## Building

### Development Build

```bash
# Quick build for current platform
make build

# Or directly
go build -o bin/clause ./cmd/clause
```

### Production Build

```bash
# Build with version info
make build VERSION=v1.0.0

# Build for all platforms
make build-all
```

### Build Flags

| Flag | Description |
|------|-------------|
| `-s` | Strip symbol table |
| `-w` | Strip DWARF debug info |
| `-X` | Set variable value at link time |

### Version Injection

```bash
# Inject version info at build time
go build -ldflags "\
  -X github.com/clause-cli/clause/internal/cmd.version=v1.0.0 \
  -X github.com/clause-cli/clause/internal/cmd.buildTime=$(date -u +%Y-%m-%dT%H:%M:%SZ) \
  -X github.com/clause-cli/clause/internal/cmd.commit=$(git rev-parse --short HEAD)" \
  -o bin/clause ./cmd/clause
```

---

## Debugging

### Using Delve

```bash
# Install delve
go install github.com/go-delve/delve/cmd/dlv@latest

# Debug a command
dlv debug ./cmd/clause -- init my-project

# Attach to running process
dlv attach <pid>
```

### Common Debug Points

```go
import "fmt"

// Quick debug print
fmt.Printf("DEBUG: %+v\n", variable)

// Use log package for more context
log.Printf("DEBUG [functionName]: variable=%v", variable)
```

### Verbose Mode

```bash
# Run with verbose output
./bin/clause --verbose init my-project

# Or with environment
CLAUSE_LOG_LEVEL=debug ./bin/clause init my-project
```

---

## Code Organization

### Package Rules

1. **`pkg/` - Public packages**
   - Can be imported by external projects
   - Must have stable APIs
   - Require comprehensive documentation

2. **`internal/` - Private packages**
   - Cannot be imported externally
   - Internal implementation details
   - Can change without notice

3. **`cmd/` - Entry points**
   - Minimal code
   - Wire up dependencies
   - Handle CLI parsing

### File Naming Conventions

| Pattern | Usage |
|---------|-------|
| `foo.go` | Main implementation |
| `foo_test.go` | Tests |
| `foo_windows.go` | Windows-specific |
| `foo_unix.go` | Unix-specific |
| `doc.go` | Package documentation |

### Package Documentation

Every package should have a `doc.go`:

```go
// Package packagename provides [brief description].
//
// Detailed description of what the package does.
//
// Example usage:
//
//   item := packagename.New()
//   err := item.DoSomething()
package packagename
```

---

## Package Development

### Adding a New Command

1. Create command file in `internal/cmd/`:

```go
// internal/cmd/mycommand.go
package cmd

import (
    "github.com/spf13/cobra"
)

var myCommandCmd = &cobra.Command{
    Use:   "mycommand [args]",
    Short: "Brief description",
    Long:  `Detailed description.`,
    RunE:  runMyCommand,
}

func init() {
    rootCmd.AddCommand(myCommandCmd)
    // Add flags here
}

func runMyCommand(cmd *cobra.Command, args []string) error {
    // Implementation
    return nil
}
```

2. Register in `init()` function
3. Add tests in `mycommand_test.go`

### Adding a New Utility

1. Determine package: `pkg/utils/` for general, `pkg/tui/` for UI
2. Create file with clear name
3. Add comprehensive tests
4. Document exported functions

### Adding a New Wizard Screen

1. Create screen in `internal/wizard/screens/`:

```go
// internal/wizard/screens/myscreen.go
package screens

type MyScreen struct {
    BaseScreen
    // Screen-specific fields
}

func NewMyScreen() *MyScreen {
    return &MyScreen{
        BaseScreen: BaseScreen{
            name: "My Screen",
            id:   "my-screen",
        },
    }
}

func (s *MyScreen) Init() tea.Cmd { return nil }

func (s *MyScreen) Update(msg tea.Msg) tea.Cmd {
    // Handle input
    return nil
}

func (s *MyScreen) View() string {
    // Render screen
    return "Screen content"
}
```

2. Add to wizard screen list in `internal/wizard/wizard.go`

---

## Template Development

### Template Structure

Templates are stored in `templates/` and embedded in the binary.

```
templates/
├── embed.go              # Embed configuration
└── governance/           # Governance templates
    ├── context.yaml.tmpl
    ├── prompt-guidelines.md.tmpl
    ├── rules.yaml.tmpl
    └── brainstorm.md.tmpl
```

### Template Syntax

Uses Go text/template with custom functions:

```go
// Available functions:
// String: lower, upper, title, trim, replace, split, join
// Case: camelCase, pascalCase, snakeCase, kebabCase
// Other: default, ternary, quote, indent, nindent
```

### Template Data

```go
type TemplateData struct {
    Project   ProjectData
    Frontend  FrontendData
    Backend   BackendData
    Infra     InfraData
    Govern    GovernData
    Now       time.Time
    Env       map[string]string
}
```

### Creating a New Template

1. Create `.tmpl` file in appropriate directory
2. Use template data fields:

```go
# {{.Project.Name}}

{{if .Frontend.Enabled}}
Frontend: {{.Frontend.Framework}}
{{end}}

Created: {{.Now.Format "2006-01-02"}}
```

3. Add to generator if needed

---

## Testing Strategy

### Test Categories

1. **Unit Tests**
   - Test individual functions/methods
   - Mock dependencies
   - Fast execution

2. **Integration Tests**
   - Test component interactions
   - Use real dependencies when possible
   - Tagged with `// +build integration`

3. **E2E Tests**
   - Test full CLI commands
   - Use temp directories
   - Tagged with `// +build e2e`

### Test Organization

```
internal/
└── generator/
    ├── generator.go
    ├── generator_test.go       # Unit tests
    ├── generator_integration_test.go  # Integration tests
    └── testdata/               # Test fixtures
        └── sample_config.yaml
```

### Writing Good Tests

```go
func TestFunctionName(t *testing.T) {
    t.Run("describes what is being tested", func(t *testing.T) {
        // Arrange
        input := "test"
        expected := "TEST"

        // Act
        result := FunctionName(input)

        // Assert
        if result != expected {
            t.Errorf("FunctionName(%q) = %q, want %q",
                input, result, expected)
        }
    })
}
```

### Test Fixtures

Store test data in `testdata/` directories:

```
internal/config/
├── loader.go
├── loader_test.go
└── testdata/
    ├── valid_config.yaml
    ├── invalid_config.yaml
    └── empty_config.yaml
```

### Benchmarks

```go
func BenchmarkFunctionName(b *testing.B) {
    for i := 0; i < b.N; i++ {
        FunctionName("input")
    }
}
```

Run benchmarks:
```bash
go test -bench=. ./...
```

---

## Release Process

### Version Numbering

We follow [Semantic Versioning](https://semver.org/):
- `MAJOR.MINOR.PATCH`
- Major: Breaking changes
- Minor: New features (backward compatible)
- Patch: Bug fixes

### Release Checklist

1. **Prepare**
   ```bash
   # Update version in code
   # Update CHANGELOG.md
   # Run all tests
   make test
   make lint
   ```

2. **Tag**
   ```bash
   git tag -a v1.0.0 -m "Release v1.0.0"
   git push origin v1.0.0
   ```

3. **CI/CD**
   - GitHub Actions builds for all platforms
   - GoReleaser creates distribution packages
   - Homebrew formula updated automatically

4. **Post-Release**
   - Verify GitHub release
   - Test installation methods
   - Announce release

### GoReleaser Configuration

See `.goreleaser.yaml` for full configuration.

Key targets:
- macOS (Intel, ARM)
- Linux (amd64, arm64)
- Windows (amd64)

---

## Troubleshooting

### Common Issues

1. **"module not found"**
   ```bash
   go mod tidy
   go mod download
   ```

2. **"permission denied"**
   ```bash
   chmod +x bin/clause
   ```

3. **Tests failing locally**
   ```bash
   # Clear test cache
   go clean -testcache
   make test
   ```

4. **Import cycle**
   - Check package dependencies
   - Consider moving shared types to separate package

### Getting Help

- Check existing issues
- Ask in GitHub Discussions
- Join Discord community
