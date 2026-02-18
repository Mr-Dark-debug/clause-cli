# 🚀 Clause CLI - Comprehensive Build Instructions for Claude Code

---

## 📋 OVERVIEW

You are tasked with building **Clause** - a Go-based, cross-platform CLI tool for AI-native project scaffolding. This is a production-grade tool that will be distributed globally via Homebrew, Winget, APT, and direct download.

**Project Name:** Clause (Framework for Organized, Reproducible, and Guided Engineering)
**CLI Command:** `Clause`
**Language:** Go 1.21+
**Key Libraries:** Cobra (CLI framework), Bubble Tea (TUI), Lip Gloss (styling)

---

## 🎯 CORE PHILOSOPHY - READ THIS FIRST

### Code Organization Principles

1. **Small Files Only:** Each file should do ONE thing and do it well. Maximum ~150-200 lines per file. If a file grows larger, split it.

2. **Utils-Based Architecture:** Create reusable utility packages. Code should never be duplicated. Every piece of logic should exist in exactly one place.

3. **Modular Design:** Each feature is a self-contained module. Dependencies flow inward. Modules communicate through interfaces.

4. **Documentation First:** Every exported function, type, and constant MUST have documentation comments. Every package MUST have a doc.go file explaining its purpose.

5. **Test Coverage:** Every utility function must have unit tests. Every module must have integration tests.

### UI/UX Principles

1. **NOT ASCII Art:** We want beautiful, modern terminal UI using Lip Gloss styling with colors, gradients, and smooth animations.

2. **Responsive Design:** The UI must adapt to terminal size changes. Listen for resize events and re-render appropriately.

3. **Colorful & Professional:** Think Claude Code or Gemini CLI aesthetics - clean, modern, with accent colors. Not retro/ASCII style.

4. **Smooth Interactions:** Use animations for transitions. Show loading states. Provide visual feedback for every action.

---

## 🎨 THEME & COLORS - USE THESE EXACTLY

```go
// Primary Brand Colors
PrimaryOrange    = "#FF6B35"  // Main brand color, CTAs, highlights
PrimaryOrangeDim = "#CC5529"  // Dimmed version for secondary elements

// Background Colors
BackgroundNavy   = "#0D1117"  // Main background (dark mode)
BackgroundCard   = "#161B22"  // Card/component backgrounds
BackgroundHover  = "#21262D"  // Hover states

// Text Colors
TextPrimary      = "#F0F6FC"  // Primary text
TextSecondary    = "#8B949E"  // Secondary/muted text
TextDim          = "#484F58"  // Very dim text (hints, placeholders)

// Semantic Colors
SuccessGreen     = "#3FB950"  // Success states
WarningAmber     = "#D29922"  // Warnings
ErrorRed         = "#F85149"  // Errors
InfoBlue         = "#58A6FF"  // Info states

// Accent Colors
AccentPurple     = "#A371F7"  // Special highlights
AccentCyan       = "#39C5CF"  // Secondary accents
AccentPink       = "#DB61A2"  // Tertiary accents

// Border Colors
BorderDefault    = "#30363D"  // Default borders
BorderMuted      = "#21262D"  // Subtle borders
BorderAccent     = "#FF6B35"  // Accent borders
```

---

## 📁 PROJECT STRUCTURE - CREATE THIS EXACTLY

```
Clause/
├── cmd/
│   └── Clause/
│       └── main.go                 # Entry point - minimal, just setup and execute
│
├── internal/
│   ├── cmd/
│   │   ├── root.go                 # Root command setup
│   │   ├── init.go                 # Init command
│   │   ├── add.go                  # Add command
│   │   ├── update.go               # Self-update command
│   │   ├── validate.go             # Validation command
│   │   ├── config.go               # Config command
│   │   └── version.go              # Version command
│   │
│   ├── wizard/
│   │   ├── wizard.go               # Main wizard orchestrator
│   │   ├── screen.go               # Screen interface and base
│   │   ├── screens/
│   │   │   ├── welcome.go          # Welcome screen
│   │   │   ├── project.go          # Project metadata screen
│   │   │   ├── frontend.go         # Frontend config screen
│   │   │   ├── backend.go          # Backend config screen
│   │   │   ├── infrastructure.go   # Infrastructure config screen
│   │   │   ├── governance.go       # AI governance config screen
│   │   │   └── summary.go          # Summary/confirmation screen
│   │   ├── components/
│   │   │   ├── input.go            # Text input component
│   │   │   ├── select.go           # Single select component
│   │   │   ├── multiselect.go      # Multi-select component
│   │   │   ├── progress.go         # Progress bar component
│   │   │   └── spinner.go          # Loading spinner component
│   │   └── navigation.go           # Screen navigation logic
│   │
│   ├── config/
│   │   ├── config.go               # Config struct and defaults
│   │   ├── loader.go               # Load config from files
│   │   ├── saver.go                # Save config to files
│   │   ├── validator.go            # Config validation
│   │   └── schema.go               # Config schema definitions
│   │
│   ├── template/
│   │   ├── engine.go               # Template processing engine
│   │   ├── parser.go               # Template DSL parser
│   │   ├── registry.go             # Template registry
│   │   ├── loader.go               # Load templates from various sources
│   │   └── types.go                # Template-related types
│   │
│   ├── generator/
│   │   ├── generator.go            # Main file generator
│   │   ├── writer.go               # File writing logic
│   │   ├── conflict.go             # Conflict resolution strategies
│   │   └── hooks.go                # Post-generation hooks
│   │
│   ├── governance/
│   │   ├── prompts.go              # Prompt generation
│   │   ├── brainstorm.go           # Brainstorm.md system
│   │   ├── registry.go             # Component registry
│   │   ├── rules.go                # Governance rule definitions
│   │   └── formats/
│   │       ├── universal.go        # Universal format output
│   │       ├── cursor.go           # Cursor Rules format
│   │       ├── continue.go         # Continue.dev format
│   │       └── claude.go           # Claude Code format
│   │
│   ├── selfupdate/
│   │   ├── checker.go              # Version checking
│   │   ├── downloader.go           # Binary download
│   │   ├── verifier.go             # Checksum verification
│   │   └── installer.go            # Atomic installation
│   │
│   └── validation/
│       ├── validator.go            # Project validation
│       ├── checker.go              # Compliance checker
│       └── reporter.go             # Validation reports
│
├── pkg/                            # Public packages (can be imported by others)
│   ├── styles/
│   │   ├── colors.go               # Color definitions
│   │   ├── theme.go                # Theme setup
│   │   ├── typography.go           # Text styles
│   │   └── layout.go               # Layout utilities
│   │
│   ├── tui/
│   │   ├── base.go                 # Base TUI utilities
│   │   ├── responsive.go           # Responsive layout handling
│   │   ├── animations.go           # Animation utilities
│   │   └── render.go               # Rendering helpers
│   │
│   ├── utils/
│   │   ├── file.go                 # File utilities
│   │   ├── path.go                 # Path utilities
│   │   ├── string.go               # String utilities
│   │   ├── slice.go                # Slice utilities
│   │   ├── map.go                  # Map utilities
│   │   ├── terminal.go             # Terminal detection/utilities
│   │   └── version.go              # Version comparison
│   │
│   └── output/
│       ├── printer.go              # Output printing
│       ├── logger.go               # Logging utilities
│       └── table.go                # Table rendering
│
├── templates/                      # Embedded templates
│   ├── frontend/
│   │   ├── nextjs/
│   │   ├── react/
│   │   └── vue/
│   ├── backend/
│   │   ├── fastapi/
│   │   ├── express/
│   │   └── django/
│   ├── fullstack/
│   │   └── nextjs-fastapi/
│   └── governance/
│       └── ai_prompt_guidelines/
│
├── docs/
│   ├── README.md                   # User documentation
│   ├── ARCHITECTURE.md             # Architecture overview
│   ├── CONTRIBUTING.md             # Contribution guidelines
│   └── templates/
│       └── creating-templates.md   # Template creation guide
│
├── test/
│   ├── integration/
│   │   └── ...                     # Integration tests
│   └── fixtures/
│       └── ...                     # Test fixtures
│
├── .goreleaser.yaml               # GoReleaser configuration
├── Makefile                        # Build automation
├── go.mod                          # Go module definition
├── go.sum                          # Dependency checksums
└── README.md                       # Project README
```

---

## 🔧 BUILD INSTRUCTIONS - STEP BY STEP

### Step 1: Project Initialization

1. Initialize the Go module: `go mod init github.com/Clause-cli/Clause`
2. Create the directory structure as shown above
3. Add dependencies:
   - `github.com/spf13/cobra` - CLI framework
   - `github.com/charmbracelet/bubbletea` - TUI framework
   - `github.com/charmbracelet/lipgloss` - Styling
   - `github.com/charmbracelet/bubbles` - Pre-built components
   - `github.com/spf13/viper` - Configuration management
   - `github.com/go-git/go-git/v5` - Git operations
   - `github.com/Masterminds/sprig/v3` - Template functions
   - `embed` - For embedding templates

### Step 2: Foundation Packages (Build These First)

#### 2.1 pkg/styles/ - Theme and Styling

Create the styling foundation that all other packages will use:

**pkg/styles/colors.go:**
- Define all color constants from the theme section above
- Provide helper functions to get colors as lipgloss.Color
- Handle terminal color capability detection (fallback for limited terminals)
- Support for both 24-bit color and 256-color terminals

**pkg/styles/theme.go:**
- Define a Theme struct containing all style definitions
- Create DefaultTheme with our brand colors
- Support theme switching (future: light/dark mode)
- Provide helper methods to apply theme to components

**pkg/styles/typography.go:**
- Define text styles (headers, body, code, etc.)
- Create styled text helpers (Title, Subtitle, Body, Code, etc.)
- Handle text wrapping and truncation
- Support for bold, italic, underline combinations

**pkg/styles/layout.go:**
- Define spacing constants and helpers
- Create border styles
- Padding and margin utilities
- Flexbox-like layout helpers for terminal

#### 2.2 pkg/tui/ - Terminal UI Utilities

Create reusable TUI components and utilities:

**pkg/tui/base.go:**
- Define common TUI types and interfaces
- Create base model with common functionality
- Provide message types used across the app
- Define common key bindings

**pkg/tui/responsive.go:**
- Implement terminal size detection
- Handle resize events gracefully
- Define breakpoints: Compact (<80), Standard (80-120), Wide (>120)
- Provide layout adaptation helpers
- Listen for SIGWINCH signals for resize

**pkg/tui/animations.go:**
- Implement smooth transition effects
- Create fade in/out animations
- Progress bar animations
- Typing effect for text
- Particle/confetti effects for celebrations

**pkg/tui/render.go:**
- Rendering pipeline utilities
- Frame composition helpers
- Double buffering for smooth updates
- Dirty rectangle optimization

#### 2.3 pkg/utils/ - General Utilities

Create small, focused utility files:

**pkg/utils/file.go:**
- FileExists(path string) bool
- IsDirectory(path string) bool
- EnsureDirectory(path string) error
- CopyFile(src, dst string) error
- WriteFile(path string, content []byte) error
- ReadFile(path string) ([]byte, error)

**pkg/utils/path.go:**
- ExpandHome(path string) string
- IsAbsPath(path string) bool
- JoinPath(parts ...string) string
- GetWorkingDirectory() string
- FindFileUp(name string, start string) string

**pkg/utils/string.go:**
- Truncate(s string, max int) string
- PadLeft(s string, pad string, length int) string
- PadRight(s string, pad string, length int) string
- TitleCase(s string) string
- KebabCase(s string) string
- SnakeCase(s string) string
- CamelCase(s string) string

**pkg/utils/slice.go:**
- Contains[T](slice []T, item T) bool
- IndexOf[T](slice []T, item T) int
- Remove[T](slice []T, index int) []T
- Map[T, U](slice []T, fn func(T) U) []U
- Filter[T](slice []T, fn func(T) bool) []T
- Unique[T](slice []T) []T

**pkg/utils/map.go:**
- Keys[M ~map[K]V, K comparable, V any](m M) []K
- Values[M ~map[K]V, K comparable, V any](m M) []V
- Merge[K comparable, V any](maps ...map[K]V) map[K]V
- GetOrDefault[K comparable, V any](m map[K]V, key K, default V) V

**pkg/utils/terminal.go:**
- IsTerminal() bool
- GetTerminalSize() (width, height int, err error)
- SupportsTrueColor() bool
- Supports256Colors() bool
- ClearScreen()
- MoveCursor(x, y int)

**pkg/utils/version.go:**
- ParseVersion(v string) (major, minor, patch int, err error)
- CompareVersions(v1, v2 string) int (-1, 0, 1)
- IsNewer(current, latest string) bool

#### 2.4 pkg/output/ - Output Handling

**pkg/output/printer.go:**
- Print functions with styling
- Printf variants with color support
- PrintSuccess, PrintError, PrintWarning, PrintInfo helpers
- Indented printing
- Boxed output for important messages

**pkg/output/logger.go:**
- Structured logging with levels (debug, info, warn, error)
- Color-coded log output
- File logging support
- Verbose mode support

**pkg/output/table.go:**
- Beautiful table rendering with Lip Gloss
- Column alignment options
- Header styling
- Row styling (alternating colors)
- Responsive column sizing

---

### Step 3: Core Commands (internal/cmd/)

Build command structure using Cobra:

#### 3.1 cmd/Clause/main.go

Keep this file MINIMAL - it should only:
- Import the root command from internal/cmd
- Execute the root command
- Handle any top-level errors
- Setup any global defer/cleanup

```go
// Structure (do not copy this, understand the pattern):
package main

import (
    "github.com/Clause-cli/Clause/internal/cmd"
    "os"
)

func main() {
    if err := cmd.Execute(); err != nil {
        // Handle error with styled output
        os.Exit(1)
    }
}
```

#### 3.2 internal/cmd/root.go

- Define the root command with Cobra
- Setup persistent flags (verbose, quiet, config file path)
- Initialize logging based on flags
- Load global configuration
- Add all subcommands
- Handle global key bindings if needed

#### 3.3 internal/cmd/init.go

The init command is the PRIMARY command. It should:
- Parse any provided flags (non-interactive mode, template override)
- If interactive mode: Launch the TUI wizard
- If non-interactive mode: Parse configuration from flags/config file
- Call the generator with the configuration
- Display progress and results
- Show success animation and next steps

#### 3.4 internal/cmd/add.go

- Parse component type flag (frontend, backend, infrastructure, governance)
- Determine what can be added based on existing project
- Launch appropriate sub-wizard or parse flags
- Update project configuration
- Generate the new component
- Update component registry

#### 3.5 internal/cmd/update.go

- Check GitHub releases for latest version
- Compare with current version
- If update available:
  - Download new binary with progress indication
  - Verify checksum
  - Perform atomic replacement
  - Verify new binary works
- Support channel selection (stable, beta, nightly)

#### 3.6 internal/cmd/validate.go

- Load project configuration
- Run all validation checks
- Display compliance report
- Show any violations with suggestions
- Return appropriate exit code

#### 3.7 internal/cmd/config.go

- Subcommands: get, set, list, init
- Handle global vs project config
- Pretty print configuration
- Validate configuration changes

#### 3.8 internal/cmd/version.go

- Display version information beautifully
- Show build date, commit, go version
- Check for updates and show notification
- Display in a nice box with branding

---

### Step 4: The Interactive Wizard (internal/wizard/)

This is the CORE experience. Build it with exceptional care.

#### 4.1 Design Philosophy for Wizard

The wizard should feel like a premium, modern application:
- Smooth transitions between screens
- Clear visual hierarchy
- Helpful contextual information
- Immediate validation feedback
- Beautiful loading states
- Celebratory completion animation

#### 4.2 internal/wizard/wizard.go

The main wizard orchestrator:
- Define the Wizard struct that holds all state
- Implement bubbletea.Model interface (Init, Update, View)
- Manage screen stack for navigation
- Handle global key events (quit, back, help)
- Coordinate between screens
- Track wizard completion state
- Return final configuration when complete

#### 4.3 internal/wizard/screen.go

Define the screen interface:
- Create Screen interface with standard methods
- All screens must implement: Init, Update, View, Title, Help
- Define common screen state (focused, completed, hasErrors)
- Create base screen implementation for embedding
- Define transition types (next, back, skip)

#### 4.4 internal/wizard/screens/ - Individual Screens

Each screen should be its own file, under 200 lines:

**welcome.go:**
- Display Clause logo with gradient colors (NOT ASCII art, use styled text/blocks)
- Quick start options as cards/buttons
- Keyboard shortcuts displayed
- Smooth entrance animation
- Option to start custom configuration

**project.go:**
- Collect project name (with validation)
- Project description
- Author name (default from git config)
- License selection (MIT, Apache, GPL, None)
- Each field as a styled input component

**frontend.go:**
- Framework selection (Next.js, React, Vue, Svelte, Angular, None)
- Show cards with framework logos/names (use colored boxes)
- State management preference
- Styling approach selection
- Component library preference
- Display technology preview as selections are made

**backend.go:**
- Framework selection (FastAPI, Django, Flask, Express, NestJS, None)
- Database selection with descriptions
- Authentication approach
- API style (REST, GraphQL, tRPC)
- Show architecture preview

**infrastructure.go:**
- Containerization options
- Deployment target selection
- CI/CD preference
- Environment management (Docker, Docker Compose, None)
- Show infrastructure diagram preview

**governance.go:**
- Strictness level slider (Visual: Permissive → Balanced → Strict)
- Documentation requirements
- Custom rules input
- Technology whitelist/blacklist
- Preview of generated governance structure

**summary.go:**
- Display all selections in organized cards
- Allow editing any section
- Show file tree preview of what will be generated
- Estimated file count
- Final confirmation button (styled prominently)

#### 4.5 internal/wizard/components/ - Reusable UI Components

**input.go:**
- Styled text input field
- Placeholder text support
- Validation state (normal, error, success)
- Character limit with counter
- Password mode (masked input)

**select.go:**
- Single selection list
- Navigate with arrows
- Search/filter functionality
- Descriptions for options
- Grouped options support
- Styled selection indicator

**multiselect.go:**
- Multiple selection with checkboxes
- Select all/none options
- Minimum/maximum selection limits
- Selected count display
- Toggle individual items

**progress.go:**
- Beautiful progress bar
- Percentage display
- Activity indicator
- Multiple color states
- Estimated time remaining

**spinner.go:**
- Multiple spinner styles
- Customizable frames
- Text label support
- Color support

#### 4.6 internal/wizard/navigation.go

- Handle screen transitions
- Implement forward/backward navigation
- Track navigation history
- Handle skip conditions
- Manage focus between elements

---

### Step 5: Configuration Management (internal/config/)

#### 5.1 internal/config/config.go

Define the complete configuration structure:
- ProjectConfig struct with all fields
- FrontendConfig, BackendConfig, InfrastructureConfig sub-structs
- GovernanceConfig with rules
- Provide sensible defaults
- Implement deep copy for config
- Validation methods

#### 5.2 internal/config/loader.go

- Load from multiple sources with priority
- Global config: ~/.Clause/config.yaml
- Project config: .Clause/config.yaml
- Environment variables: Clause_*
- Command-line flags (highest priority)
- Merge configs intelligently

#### 5.3 internal/config/saver.go

- Save config in YAML format
- Preserve comments where possible
- Atomic write (write to temp, then rename)
- Create backup of existing config

#### 5.4 internal/config/validator.go

- Validate all config fields
- Check for required fields
- Validate cross-field dependencies
- Return detailed validation errors
- Suggest corrections

#### 5.5 internal/config/schema.go

- Define JSON schema for config
- Enable auto-completion in editors
- Document all fields
- Version the schema

---

### Step 6: Template Engine (internal/template/)

#### 6.1 internal/template/engine.go

The template processing engine:
- Load and parse templates
- Apply configuration data to templates
- Handle template inheritance
- Support template composition
- Provide template functions (sprig + custom)
- Cache compiled templates

#### 6.2 internal/template/parser.go

Parse the template DSL:
- Variable substitution: {{.Field}}
- Conditionals: {{if .Condition}}...{{end}}
- Loops: {{range .Items}}...{{end}}
- File operations: {{file "path"}}
- Include other templates: {{template "name" .}}
- Custom Clause directives

#### 6.3 internal/template/registry.go

Manage available templates:
- List templates by category
- Get template by ID
- Register new templates
- Template metadata (name, description, author, version)
- Template dependencies

#### 6.4 internal/template/loader.go

Load templates from multiple sources:
- Embedded templates (compiled into binary)
- Remote templates (from Git repositories)
- Local templates (from project directory)
- Cache remote templates locally
- Validate template structure

#### 6.5 internal/template/types.go

Define template-related types:
- Template struct
- TemplateMetadata struct
- TemplateSource enum (embedded, remote, local)
- TemplateVariable struct
- TemplateFile struct

---

### Step 7: File Generator (internal/generator/)

#### 7.1 internal/generator/generator.go

The main generator orchestrator:
- Accept configuration and templates
- Determine file generation order (dependencies)
- Process each template
- Track generation progress
- Report errors and warnings
- Return generation summary

#### 7.2 internal/generator/writer.go

Handle file writing:
- Create directories as needed
- Set file permissions appropriately
- Handle file encoding
- Write with buffering
- Implement atomic writes

#### 7.3 internal/generator/conflict.go

Handle file conflicts:
- Detect existing files
- Apply conflict strategy (skip, backup, merge, overwrite)
- Implement smart merge for structured files (JSON, YAML)
- Log conflict resolution decisions

#### 7.4 internal/generator/hooks.go

Post-generation hooks:
- RunGitInit - Initialize git repository
- RunInitialCommit - Create initial commit
- RunDependencyInstall - Install dependencies
- RunValidation - Validate generated project
- Custom hooks from configuration

---

### Step 8: AI Governance System (internal/governance/)

#### 8.1 internal/governance/prompts.go

Generate governance prompts:
- Compose prompts from layers (base + tech + project + custom)
- Apply variable substitution
- Handle prompt length optimization
- Generate for different AI tools
- Include examples and constraints

#### 8.2 internal/governance/brainstorm.go

Implement Brainstorm.md system:
- Define Brainstorm.md structure
- Generate initial template
- Parse existing Brainstorm.md
- Provide update utilities
- Extract open questions
- Track decision history

#### 8.3 internal/governance/registry.go

Component registry management:
- Define registry structure
- Register new components
- Query components (by name, type, dependency)
- Update component metadata
- Validate registry integrity
- Generate registry file

#### 8.4 internal/governance/rules.go

Define governance rules:
- Architectural constraint rules
- Technology whitelist/blacklist rules
- Documentation requirement rules
- Code quality standard rules
- Rule validation logic
- Rule severity levels

#### 8.5 internal/governance/formats/ - Output Formats

Each file generates governance in a specific format:

**universal.go:**
- Generate ai_prompt_guidelines/ directory
- Create structured markdown files
- Include all governance aspects
- Make it human and AI readable

**cursor.go:**
- Generate .cursor/rules/*.md files
- Apply path pattern scoping
- Format for Cursor's rule system
- Include Cursor-specific features

**continue.go:**
- Generate .continue/config.json
- Format for Continue.dev's config
- Include custom instructions
- Setup context configuration

**claude.go:**
- Generate CLAUDE.md at project root
- Format for Claude Code's context
- Include structured sections
- Optimize for Claude's understanding

---

### Step 9: Self-Update System (internal/selfupdate/)

#### 9.1 internal/selfupdate/checker.go

Check for updates:
- Query GitHub Releases API
- Parse latest version
- Compare with current version
- Support update channels (stable, beta, nightly)
- Handle rate limiting
- Cache version info

#### 9.2 internal/selfupdate/downloader.go

Download updates:
- Download binary with progress reporting
- Handle network errors gracefully
- Support resumable downloads
- Verify download completion
- Support proxy configuration

#### 9.3 internal/selfupdate/verifier.go

Verify downloaded binaries:
- Verify SHA256 checksum
- Verify GPG signature (if available)
- Check binary is executable
- Verify binary runs correctly

#### 9.4 internal/selfupdate/installer.go

Install updates:
- Atomic binary replacement
- Handle permissions
- Backup previous version
- Rollback on failure
- Verify installation success

---

### Step 10: Validation System (internal/validation/)

#### 10.1 internal/validation/validator.go

Project validation:
- Load project configuration
- Run all registered checks
- Aggregate results
- Return pass/fail status
- Generate detailed report

#### 10.2 internal/validation/checker.go

Compliance checking:
- Check technology constraints
- Check architectural boundaries
- Check documentation completeness
- Check code quality standards
- Return specific violations

#### 10.3 internal/validation/reporter.go

Generate reports:
- Format validation results
- Create visual report with styling
- Include violation details
- Suggest corrections
- Support JSON output for CI

---

### Step 11: Templates (templates/)

Create the embedded templates:

#### 11.1 templates/governance/ai_prompt_guidelines/

Create the core governance template files:

**system_prompt.md.template:**
- Base behavioral instructions
- Role definition for the AI
- Working style guidelines
- Response format expectations

**architecture.md.template:**
- Architectural principles
- Layer separation rules
- Dependency direction rules
- Module boundary definitions

**technologies.md.template:**
- Approved technologies list
- Forbidden technologies list
- Version constraints
- Library usage guidelines

**documentation.md.template:**
- Documentation requirements
- Comment standards
- README requirements
- API documentation standards

**brainstorm.md.template:**
- Initial Brainstorm.md structure
- Section definitions
- Usage instructions for AI

**component_registry.json.template:**
- Initial registry structure
- Empty component list
- Schema definition

**brand_guidelines.md.template:**
- Theme and color guidelines
- UI component standards
- Consistency requirements
- Reusable component patterns

---

### Step 12: Build Configuration

#### 12.1 .goreleaser.yaml

Configure GoReleaser for multi-platform distribution:
- Build for darwin (amd64, arm64)
- Build for linux (amd64, arm64)
- Build for windows (amd64, arm64)
- Create archives (.tar.gz for unix, .zip for windows)
- Generate checksums
- Create Homebrew formula
- Create Winget manifest
- Create Snap package
- Sign binaries (macOS, Windows)

#### 12.2 Makefile

Build automation commands:
- `make build` - Build for current platform
- `make build-all` - Build for all platforms
- `make test` - Run tests
- `make test-coverage` - Run tests with coverage
- `make lint` - Run linters
- `make fmt` - Format code
- `make install` - Install locally
- `make release` - Create release

---

## 🎨 UI SPECIFICATIONS - IMPORTANT

### Welcome Screen Design

```
┌──────────────────────────────────────────────────────────────────┐
│                                                                  │
│     ███████╗ ██████╗ █████╗  ██████╗██╗  ██╗                    │
│     ██╔════╝██╔════╝██╔══██╗██╔════╝██║  ██║                    │
│     █████╗  ██║     ███████║██║     ███████║                    │
│     ██╔══╝  ██║     ██╔══██║██║     ██╔══██║                    │
│     ██║     ╚██████╗██║  ██║╚██████╗██║  ██║                    │
│     ╚═╝      ╚═════╝╚═╝  ╚═╝ ╚═════╝╚═╝  ╚═╝                    │
│                                                                  │
│            [Framework for Organized, Reproducible,               │
│                    and Guided Engineering]                       │
│                                                                  │
│   ┌─────────────────────────┐  ┌─────────────────────────┐      │
│   │   🚀 Quick Start        │  │   ⚙️  Custom Setup       │      │
│   │   Next.js + FastAPI     │  │   Configure everything  │      │
│   └─────────────────────────┘  └─────────────────────────┘      │
│                                                                  │
│   Press [Enter] to select    Press [Tab] to switch              │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

NOT ASCII art - use Lip Gloss to create:
- Colored background boxes
- Rounded corners
- Gradient effects where possible
- Smooth hover/focus states
- Professional typography

### Selection Screen Design

```
┌──────────────────────────────────────────────────────────────────┐
│  Frontend Framework                                    [ 2 / 6 ] │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│   Choose your frontend framework:                                │
│                                                                  │
│   ┌─────────────────────────────────────────────────────────┐   │
│   │ ◉ Next.js                                                │   │
│   │   React framework with SSR, SSG, and more               │   │
│   │   Recommended for: Full-stack applications              │   │
│   └─────────────────────────────────────────────────────────┘   │
│                                                                  │
│   ┌─────────────────────────────────────────────────────────┐   │
│   │ ○ React (Vite)                                           │   │
│   │   Fast, modern React development                         │   │
│   │   Recommended for: SPAs, lightweight apps               │   │
│   └─────────────────────────────────────────────────────────┘   │
│                                                                  │
│   ┌─────────────────────────────────────────────────────────┐   │
│   │ ○ Vue.js                                                 │   │
│   │   Progressive JavaScript framework                       │   │
│   └─────────────────────────────────────────────────────────┘   │
│                                                                  │
│   ─────────────────────────────────────────────────────────────  │
│   [↑↓ Navigate]  [Enter Select]  [? Help]  [Esc Back]           │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

### Progress/Generation Screen

```
┌──────────────────────────────────────────────────────────────────┐
│  Generating Project                                     45%     │
├──────────────────────────────────────────────────────────────────┤
│                                                                  │
│   ████████████████████░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░   │
│                                                                  │
│   ✓ Created project structure                                   │
│   ✓ Created configuration files                                  │
│   ✓ Created frontend scaffold                                    │
│   ◐ Creating backend scaffold...                                 │
│   ○ Creating governance files                                    │
│   ○ Installing dependencies                                      │
│                                                                  │
│   ─────────────────────────────────────────────────────────────  │
│   Creating: src/backend/api/routes/users.py                      │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

### Success Screen

```
┌──────────────────────────────────────────────────────────────────┐
│                                                                  │
│                        🎉 Project Created!                       │
│                                                                  │
│     my-awesome-project                                           │
│     ├── frontend (Next.js)                                       │
│     ├── backend (FastAPI)                                        │
│     ├── infrastructure (Docker)                                  │
│     └── ai_prompt_guidelines                                     │
│                                                                  │
│   ─────────────────────────────────────────────────────────────  │
│                                                                  │
│   Next steps:                                                    │
│                                                                  │
│   1. cd my-awesome-project                                       │
│   2. Clause validate                                              │
│   3. Start building with AI assistance!                          │
│                                                                  │
│   ─────────────────────────────────────────────────────────────  │
│                                                                  │
│   Documentation: https://Clause.dev/docs                          │
│   Community:     https://discord.gg/Clause                        │
│                                                                  │
└──────────────────────────────────────────────────────────────────┘
```

---

## 📝 DOCUMENTATION REQUIREMENTS

Every package MUST have a doc.go file:

```go
// Package wizard provides an interactive configuration wizard
// for creating new projects with Clause.
//
// The wizard uses Bubble Tea for TUI and presents a series of
// screens for collecting configuration from the user.
//
// Example usage:
//
//   cfg, err := wizard.Run()
//   if err != nil {
//       // handle error
//   }
//   // use cfg for project generation
//
// The wizard handles all user interaction and returns a complete
// configuration ready for use by the generator.
package wizard
```

Every exported function MUST have a doc comment:

```go
// Run executes the interactive wizard and returns the collected configuration.
// It returns an error if the wizard fails to initialize or if the user
// cancels the operation.
//
// The wizard runs in the terminal and handles all user interaction.
// It respects the terminal size and adapts its layout accordingly.
func Run() (*config.ProjectConfig, error) {
    // implementation
}
```

---

## ✅ TESTING REQUIREMENTS

### Unit Tests

Every utility function must have tests:

```go
// file: pkg/utils/string_test.go
package utils

import "testing"

func TestTruncate(t *testing.T) {
    tests := []struct {
        input    string
        max      int
        expected string
    }{
        {"hello", 10, "hello"},
        {"hello world", 5, "he..."},
        {"", 5, ""},
    }
    
    for _, tt := range tests {
        result := Truncate(tt.input, tt.max)
        if result != tt.expected {
            t.Errorf("Truncate(%q, %d) = %q, want %q", 
                tt.input, tt.max, result, tt.expected)
        }
    }
}
```

### Integration Tests

Create tests for complete workflows:

```go
// file: test/integration/init_test.go
package integration

import (
    "os"
    "path/filepath"
    "testing"
    
    "github.com/Clause-cli/Clause/internal/cmd"
)

func TestInitCommand(t *testing.T) {
    // Create temp directory
    tmpDir := t.TempDir()
    
    // Run init command
    // Verify project structure created
    // Verify governance files exist
    // Verify configuration saved
}
```

---

## 🚀 EXECUTION ORDER

Build in this order:

1. **pkg/styles/** - Foundation for all styling
2. **pkg/utils/** - Utilities used everywhere
3. **pkg/tui/** - TUI helpers for wizard
4. **pkg/output/** - Output formatting
5. **internal/config/** - Configuration management
6. **internal/cmd/root.go** - Basic command structure
7. **internal/wizard/** - Interactive wizard (biggest component)
8. **internal/template/** - Template engine
9. **internal/generator/** - File generation
10. **internal/governance/** - AI governance system
11. **templates/** - Actual template files
12. **internal/cmd/init.go** - Connect wizard to generator
13. **internal/selfupdate/** - Self-update mechanism
14. **internal/validation/** - Validation system
15. **internal/cmd/*.go** - Remaining commands
16. **Build configuration** - GoReleaser, Makefile

---

## 🎯 SUCCESS CRITERIA

The tool is complete when:

1. ✅ Running `Clause init` launches a beautiful, responsive TUI wizard
2. ✅ Completing the wizard generates a complete project structure
3. ✅ The ai_prompt_guidelines directory contains all governance files
4. ✅ The Brainstorm.md file is created with proper structure
5. ✅ Running `Clause validate` checks governance compliance
6. ✅ Running `Clause update` successfully updates the binary
7. ✅ The tool works on macOS, Linux, and Windows
8. ✅ The UI adapts to terminal size changes
9. ✅ All code is documented with comments
10. ✅ All utility functions have unit tests
11. ✅ The tool can be installed via Homebrew, Winget, or curl|bash

---

## 💡 KEY IMPLEMENTATION NOTES

### Responsive Design Implementation

```go
// Listen for terminal resize
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.WindowSizeMsg:
        m.width = msg.Width
        m.height = msg.Height
        m.breakpoint = calculateBreakpoint(msg.Width)
        return m, nil
    }
    // ... rest of update
}

func calculateBreakpoint(width int) Breakpoint {
    if width < 80 {
        return BreakpointCompact
    } else if width < 120 {
        return BreakpointStandard
    }
    return BreakpointWide
}
```

### Color Application Pattern

```go
// Use Lip Gloss for styling
var (
    titleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(lipgloss.Color(styles.PrimaryOrange)).
        Padding(0, 1)
    
    cardStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(lipgloss.Color(styles.BorderDefault)).
        Padding(1, 2).
        Background(lipgloss.Color(styles.BackgroundCard))
)
```

### Animation Pattern

```go
// Use Bubble Tea's tick command for animations
type tickMsg time.Time

func animate() tea.Cmd {
    return tea.Tick(time.Millisecond*16, func(t time.Time) tea.Msg {
        return tickMsg(t)
    })
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tickMsg:
        m.frame++
        return m, animate()
    }
    // ...
}
```

---

## 📌 FINAL NOTES

1. **Start small, iterate**: Build one component at a time, test thoroughly, then move on.

2. **Document as you go**: Don't leave documentation for later. Write doc comments immediately.

3. **Test everything**: If it's not tested, it's broken. Aim for 80%+ coverage on utilities.

4. **Keep files small**: If a file exceeds 200 lines, split it. Modularity is key.

5. **Make it beautiful**: Every screen should feel premium. Use colors, spacing, and animations.

6. **Handle errors gracefully**: Never crash. Show helpful error messages with suggestions.

7. **Think cross-platform**: Test on macOS, Linux, and Windows throughout development.

8. **Performance matters**: The wizard should feel instant. No blocking operations on the main thread.

---

**Now, Claude Code: Build Clause following these specifications. Create beautiful, modular, well-documented code. Make something developers will love to use.** 🚀
