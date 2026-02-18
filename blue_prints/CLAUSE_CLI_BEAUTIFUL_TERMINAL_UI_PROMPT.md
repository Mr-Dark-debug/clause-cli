# Clause CLI - Beautiful Professional Terminal UI Design Prompt

## Overview

You are tasked with redesigning the **Clause CLI** terminal interface to create a stunning, professional, and modern user experience. The current interface is plain and utilitarian. We need to transform it into a beautiful, visually appealing terminal UI that rivals industry-leading AI CLIs like **Claude Code**, **Gemini CLI**, and **Aider**.

---

## Reference Designs - Study These Carefully

### 1. Claude Code CLI (Anthropic)
- **Design Philosophy**: Clean, minimal, modern - "just the right amount of visual elements"
- **No excessive emojis** - selective and purposeful use
- **Subtle animations** for state changes
- **Professional color palette** with muted tones
- **Elegant typography hierarchy**
- **Smooth transitions** between states
- **Box-based layouts** with rounded corners when terminal supports it
- **Inline status indicators** with color-coded states

### 2. Gemini CLI (Google)
- **Key Innovation**: "Making the terminal beautiful one pixel at a time"
- **Eliminates visual noise** associated with terminal applications
- **Graphical-interface polish** in terminal environment
- **Responsive rendering** that adapts to terminal width
- **Clean separation** between different UI sections
- **Minimal flickering** and stable rendering
- **Professional gradient-like effects** using ANSI colors

### 3. Aider CLI
- **Real-time diff visualization** in terminal
- **Multi-pane layouts** for code and chat
- **Status indicators** for file tracking
- **Interactive elements** that respond naturally

### 4. Charmbracelet Ecosystem (Lipgloss + Bubble Tea)
- **CSS-like styling** for terminal applications
- **Responsive layouts** that adapt to terminal size
- **Beautiful tables** with proper alignment
- **Border styles** (rounded, double, thick, normal)
- **Color gradients** and sophisticated palettes
- **Padding and margin** control
- **Text alignment** and wrapping

---

## Current Clause CLI Output (What Needs Improvement)

```
Clause (Framework for Organized, Reproducible, and Guided Engineering)
is a cross-platform CLI tool for creating AI-ready project structures.

It generates not just folder structures, but comprehensive AI context systems      
that guide AI coding assistants toward consistent, maintainable code.

Run 'clause init' to create a new project with an interactive wizard.

Usage:
  clause [command]

Available Commands:
  add         Add a new component to an existing project
  completion  Generate the autocompletion script for the specified shell
  config      Manage Clause configuration
  help        Help about any command
  init        Initialize a new AI-ready project
  update      Update Clause to the latest version
  validate    Validate project governance compliance
  version     Print the version information

Flags:
      --config string   config file (default is $HOME/.clause/config.yaml)
  -h, --help            help for clause
      --no-color        disable colored output
  -q, --quiet           suppress non-essential output
  -v, --verbose         verbose output

Use "clause [command] --help" for more information about a command.
```

**Problems with current design:**
- No visual hierarchy - everything looks equally important
- No color or styling - plain monochrome text
- No branding or visual identity
- No box formatting or visual separation
- No icons or visual indicators
- No emphasis on key actions
- Boring command list without visual grouping

---

## Target Design - What We Want

### Design Principles

1. **Visual Hierarchy** - Important commands and actions should stand out
2. **Brand Identity** - Clause should have a recognizable visual style
3. **Professional Polish** - Clean lines, proper spacing, consistent styling
4. **Color Strategy** - Sophisticated color palette, not rainbow chaos
5. **Responsive Design** - Adapt to terminal width gracefully
6. **Accessibility** - Support `--no-color` flag, good contrast ratios
7. **Performance** - No flickering, instant rendering
8. **Minimalism** - No clutter, every element earns its place

### Color Palette (Use Lipgloss)

```go
// Primary Colors
var (
    // Brand colors - Clause identity
    ClausePrimary   = lipgloss.Color("#7C3AED")  // Vibrant purple - main brand
    ClauseSecondary = lipgloss.Color("#A78BFA")  // Lighter purple - accents
    ClauseAccent    = lipgloss.Color("#F472B6")  // Pink accent - highlights
    
    // Semantic colors - meaning
    SuccessColor    = lipgloss.Color("#10B981")  // Green - success
    WarningColor    = lipgloss.Color("#F59E0B")  // Amber - warnings
    ErrorColor      = lipgloss.Color("#EF4444")  // Red - errors
    InfoColor       = lipgloss.Color("#3B82F6")  // Blue - information
    
    // UI colors - structure
    BorderColor     = lipgloss.Color("#374151")  // Dark gray - borders
    MutedColor      = lipgloss.Color("#6B7280")  // Gray - muted text
    TextColor       = lipgloss.Color("#F9FAFB")  // White - primary text
    DimmedColor     = lipgloss.Color("#9CA3AF")  // Light gray - secondary text
    
    // Background colors - containers
    BgDark          = lipgloss.Color("#1F2937")  // Dark background
    BgDarker        = lipgloss.Color("#111827")  // Darker background
    BgCard          = lipgloss.Color("#1F2937")  // Card background
)
```

---

## Required UI Components

### 1. Welcome Header / Banner

Create a stunning ASCII art banner with the Clause logo:

```
╔═══════════════════════════════════════════════════════════════════╗
║                                                                   ║
║   ██████╗██╗      █████╗ ██╗   ██╗██████╗ ███████╗               ║
║  ██╔════╝██║     ██╔══██╗██║   ██║██╔══██╗██╔════╝               ║
║  ██║     ██║     ███████║██║   ██║██║  ██║█████╗                 ║
║  ██║     ██║     ██╔══██║██║   ██║██║  ██║██╔══╝                 ║
║  ╚██████╗███████╗██║  ██║╚██████╔╝██████╔╝███████╗               ║
║   ╚═════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝               ║
║                                                                   ║
║   Framework for Organized, Reproducible, and Guided Engineering  ║
║                        Version 1.0.0                              ║
╚═══════════════════════════════════════════════════════════════════╝
```

**Requirements:**
- Centered ASCII art logo with gradient effect (if terminal supports true color)
- Subtle tagline underneath
- Version number displayed elegantly
- Rounded border with primary color
- Responsive - shrinks gracefully on narrow terminals

### 2. Description Card

A styled card with the tool description:

```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│  Clause is a cross-platform CLI tool for creating AI-ready     │
│  project structures with comprehensive AI context systems that  │
│  guide AI coding assistants toward consistent, maintainable     │
│  code.                                                          │
│                                                                 │
│  ✓ Generates complete project scaffolding                       │
│  ✓ Creates AI governance guidelines                             │
│  ✓ Includes context files for AI assistants                     │
│  ✓ Supports Next.js, FastAPI, and more                          │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

**Requirements:**
- Box with subtle border
- Proper padding (left, right, top, bottom)
- Bullet points with checkmarks using success color
- Word wrapping that respects terminal width
- Muted border color, white text

### 3. Quick Start Call-to-Action

Prominent call-to-action section:

```
┌─────────────────────────────────────────────────────────────────┐
│  🚀 Quick Start                                                 │
│                                                                 │
│     clause init          Create a new project with wizard      │
│     clause init --quick  Skip wizard, use defaults             │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

**Requirements:**
- Eye-catching header with emoji or icon
- Command highlighted in distinct color
- Description in muted color
- Clear visual separation from other sections

### 4. Commands Grid/Table

Beautiful command listing with descriptions:

```
┌─────────────────────────────────────────────────────────────────┐
│  AVAILABLE COMMANDS                                             │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│  Project Commands                                               │
│  ───────────────                                                │
│    init       Initialize a new AI-ready project                │
│    add        Add a new component to an existing project       │
│    validate   Validate project governance compliance           │
│                                                                 │
│  Configuration                                                  │
│  ───────────────                                                │
│    config     Manage Clause configuration settings             │
│    update     Update Clause to the latest version              │
│                                                                 │
│  Utility Commands                                               │
│  ───────────────                                                │
│    version    Print the version information                    │
│    help       Help about any command                           │
│    completion Generate shell autocompletion script             │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

**Requirements:**
- Grouped by category with section headers
- Section dividers with muted color
- Commands in bold with accent color
- Descriptions in normal weight
- Proper alignment using tables
- Hover/selection effect in interactive mode

### 5. Global Flags Section

Styled flags documentation:

```
┌─────────────────────────────────────────────────────────────────┐
│  GLOBAL FLAGS                                                   │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│      --config <file>    Config file path                       │
│                         Default: $HOME/.clause/config.yaml     │
│                                                                 │
│  -h, --help             Show help for command                  │
│                                                                 │
│      --no-color         Disable colored output                 │
│                                                                 │
│  -q, --quiet            Suppress non-essential output          │
│                                                                 │
│  -v, --verbose          Enable verbose output                  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

**Requirements:**
- Two-column layout for flag and description
- Default values shown in italics/muted
- Short and long flags shown together
- Proper indentation for wrapped descriptions

### 6. Footer

Professional footer with helpful links:

```
───────────────────────────────────────────────────────────────────
  📚 Documentation: https://clause.dev/docs
  💡 Examples:      https://github.com/clause-cli/examples
  🐛 Issues:        https://github.com/clause-cli/clause/issues
───────────────────────────────────────────────────────────────────
```

---

## Technical Implementation Requirements

### File Structure

```
internal/
├── tui/
│   ├── components/
│   │   ├── banner.go        # ASCII art banner component
│   │   ├── card.go          # Card container component
│   │   ├── commands.go      # Command list component
│   │   ├── flags.go         # Flags display component
│   │   ├── footer.go        # Footer component
│   │   └── quickstart.go    # Quick start CTA component
│   ├── styles/
│   │   ├── colors.go        # Color palette definitions
│   │   ├── theme.go         # Theme system
│   │   └── typography.go    # Text styling utilities
│   ├── layout/
│   │   ├── flex.go          # Flexbox-like layout
│   │   ├── grid.go          # Grid layout system
│   │   └── responsive.go    # Responsive utilities
│   └── render.go            # Main render function
├── cmd/
│   └── root.go              # Update to use TUI
```

### Key Libraries

```go
import (
    "github.com/charmbracelet/lipgloss"
    "github.com/charmbracelet/bubbletea"
    "github.com/charmbracelet/bubbles/table"
    "github.com/charmbracelet/bubbles/progress"
    "github.com/charmbracelet/bubbles/spinner"
    "github.com/charmbracelet/bubbles/textinput"
    "github.com/charmbracelet/bubbles/list"
    "github.com/charmbracelet/lipgloss/table"
)
```

### Style Definitions (lipgloss)

```go
package styles

import "github.com/charmbracelet/lipgloss"

// Base styles
var (
    // Title style - for main headers
    TitleStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(ClausePrimary).
        Padding(0, 1)
    
    // Command style - for command names
    CommandStyle = lipgloss.NewStyle().
        Bold(true).
        Foreground(ClauseAccent).
        Padding(0, 2, 0, 0)
    
    // Description style - for descriptions
    DescStyle = lipgloss.NewStyle().
        Foreground(DimmedColor)
    
    // Card style - for containers
    CardStyle = lipgloss.NewStyle().
        Border(lipgloss.RoundedBorder()).
        BorderForeground(BorderColor).
        Padding(1, 2).
        Margin(1, 0)
    
    // Section header style
    SectionHeader = lipgloss.NewStyle().
        Bold(true).
        Foreground(TextColor).
        Padding(1, 0, 0, 0).
        MarginBottom(1)
    
    // Divider style
    DividerStyle = lipgloss.NewStyle().
        Foreground(BorderColor).
        SetString("─")
)

// Layout utilities
func JoinVertical(parts ...string) string {
    return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

func JoinHorizontal(parts ...string) string {
    return lipgloss.JoinHorizontal(lipgloss.Top, parts...)
}

func Center(width int, content string) string {
    return lipgloss.NewStyle().
        Width(width).
        Align(lipgloss.Center).
        Render(content)
}
```

### Responsive Design

```go
package layout

import (
    "os"
    "golang.org/x/term"
)

type TerminalSize struct {
    Width  int
    Height int
}

func GetTerminalSize() TerminalSize {
    width, height, err := term.GetSize(int(os.Stdout.Fd()))
    if err != nil {
        return TerminalSize{Width: 80, Height: 24} // defaults
    }
    return TerminalSize{Width: width, Height: height}
}

func IsNarrowTerminal() bool {
    size := GetTerminalSize()
    return size.Width < 60
}

// Adjust content based on terminal width
func ResponsiveLayout(content string) string {
    size := GetTerminalSize()
    // Return appropriate layout based on size
    if size.Width < 40 {
        return CompactLayout(content)
    } else if size.Width < 80 {
        return StandardLayout(content)
    }
    return WideLayout(content)
}
```

---

## Complete Example Output

When user runs `clause` without arguments, they should see:

```
╔═══════════════════════════════════════════════════════════════════════╗
║                                                                       ║
║    ██████╗██╗      █████╗ ██╗   ██╗██████╗ ███████╗                  ║
║   ██╔════╝██║     ██╔══██╗██║   ██║██╔══██╗██╔════╝                  ║
║   ██║     ██║     ███████║██║   ██║██║  ██║█████╗                    ║
║   ██║     ██║     ██╔══██║██║   ██║██║  ██║██╔══╝                    ║
║   ╚██████╗███████╗██║  ██║╚██████╔╝██████╔╝███████╗                  ║
║    ╚═════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝                  ║
║                                                                       ║
║        Framework for Organized, Reproducible,                        ║
║              and Guided Engineering                                  ║
║                          v1.0.0                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  ✨ Create AI-Ready Projects                                        │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Clause generates complete project scaffolding with built-in AI     │
│  governance, context files, and best practices for AI assistants.   │
│                                                                     │
│  • Complete project structure generation                            │
│  • AI governance guidelines included                                │
│  • Context files for AI assistants                                  │
│  • Support for Next.js, FastAPI, Go, and more                       │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  🚀 Quick Start                                                     │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│      clause init                                                    │
│      ────────────────                                               │
│      Launch an interactive wizard to create a new project           │
│                                                                     │
│      clause init --quick                                            │
│      ─────────────────────                                          │
│      Skip the wizard and use sensible defaults                      │
│                                                                     │
│      clause init --template nextjs                                  │
│      ────────────────────────────                                   │
│      Start with a specific template                                 │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  📋 Available Commands                                              │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  PROJECT                                                            │
│  ─────────────────────────────────────────────────────────────────  │
│    init         Initialize a new AI-ready project                  │
│    add          Add components to existing project                 │
│    validate     Check governance compliance                        │
│                                                                     │
│  CONFIGURATION                                                      │
│  ─────────────────────────────────────────────────────────────────  │
│    config       Manage Clause settings                             │
│    update       Update to latest version                           │
│                                                                     │
│  UTILITY                                                            │
│  ─────────────────────────────────────────────────────────────────  │
│    version      Show version info                                  │
│    help         Get help for any command                           │
│    completion   Generate shell completion                          │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  ⚙️  Global Flags                                                   │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│      --config <path>    Use specific config file                   │
│                         Default: ~/.clause/config.yaml              │
│                                                                     │
│  -h, --help             Show help information                      │
│      --no-color         Disable colored output                     │
│  -q, --quiet            Minimal output mode                        │
│  -v, --verbose          Detailed output mode                       │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

───────────────────────────────────────────────────────────────────────
   📚 Docs: clause.dev  │  💻 GitHub: clause-cli/clause  │  v1.0.0
───────────────────────────────────────────────────────────────────────
```

---

## Animation & Interaction Enhancements

### Loading States

```go
// Spinner for async operations
spinnerModel := spinner.New()
spinnerModel.Spinner = spinner.Dot
spinnerModel.Style = lipgloss.NewStyle().Foreground(ClausePrimary)

// Progress bar for file generation
progressModel := progress.New(progress.WithDefaultGradient())
```

### Interactive Mode (for `clause init`)

```
╔═══════════════════════════════════════════════════════════════════════╗
║                    🎯 Clause Project Setup                            ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  Project Name                                                       │
│  ─────────────                                                      │
│  > my-awesome-project_                                              │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  Project Type                                                       │
│  ─────────────                                                      │
│                                                                     │
│    ○ Frontend (Next.js, React, Vue)                                │
│    ● Backend (FastAPI, Express, Go)                                │
│    ○ Full-Stack (Next.js + FastAPI)                                │
│    ○ CLI Tool (Go, Rust, Python)                                   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  Additional Options                                                 │
│  ────────────────────                                               │
│                                                                     │
│    [✓] Include AI governance files                                 │
│    [✓] Add example components                                      │
│    [ ] Set up Docker configuration                                 │
│    [ ] Initialize Git repository                                   │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

  Use ↑/↓ to navigate, Space to select, Enter to continue
```

### Success Output

```
╔═══════════════════════════════════════════════════════════════════════╗
║                         ✅ Project Created!                           ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│  📁 my-awesome-project                                              │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│  Created the following structure:                                   │
│                                                                     │
│  my-awesome-project/                                                │
│  ├── 📁 .clause/                     AI governance config          │
│  │   ├── system_prompt.md           AI behavior rules             │
│  │   ├── architecture.md            Architecture decisions        │
│  │   └── registry.json              Component registry             │
│  ├── 📁 src/                        Source code                    │
│  ├── 📁 templates/                  Project templates              │
│  ├── 📄 .env.example               Environment template            │
│  ├── 📄 README.md                  Project documentation          │
│  └── 📄 docker-compose.yml         Docker setup                   │
│                                                                     │
│  ─────────────────────────────────────────────────────────────────  │
│  📊 Stats: 12 files created, 3 directories                         │
│  ⏱️  Time: 0.23s                                                    │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  🚀 Next Steps                                                      │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│    cd my-awesome-project                                            │
│    code .                        Open in VS Code                   │
│    clause add component          Add new components                │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## Error Display

```
╔═══════════════════════════════════════════════════════════════════════╗
║                         ❌ Error                                      ║
╚═══════════════════════════════════════════════════════════════════════╝

┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│  Failed to create project: directory already exists                │
│                                                                     │
│  Path: ./my-awesome-project                                        │
│                                                                     │
│  💡 Tip: Use --force to overwrite existing files                   │
│     or specify a different name with --name                        │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  🔧 Suggested Commands                                              │
├─────────────────────────────────────────────────────────────────────┤
│                                                                     │
│    clause init --force              Overwrite existing             │
│    clause init --name new-project   Use different name            │
│    clause init --help               Show all options              │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

---

## Implementation Checklist

### Phase 1: Core Styles
- [ ] Create `internal/tui/styles/colors.go` with color palette
- [ ] Create `internal/tui/styles/typography.go` with text styles
- [ ] Create `internal/tui/styles/theme.go` with theme management
- [ ] Add responsive terminal detection

### Phase 2: Components
- [ ] Create `internal/tui/components/banner.go` with ASCII logo
- [ ] Create `internal/tui/components/card.go` for container boxes
- [ ] Create `internal/tui/components/commands.go` for command list
- [ ] Create `internal/tui/components/flags.go` for flags display
- [ ] Create `internal/tui/components/footer.go` for footer
- [ ] Create `internal/tui/components/quickstart.go` for CTA

### Phase 3: Interactive Elements
- [ ] Create `internal/tui/components/form.go` for interactive forms
- [ ] Add spinner for loading states
- [ ] Add progress bar for file generation
- [ ] Add list selection with bubbletea

### Phase 4: Integration
- [ ] Update `internal/cmd/root.go` to use new TUI
- [ ] Update `internal/cmd/init.go` for interactive wizard
- [ ] Add `--no-color` flag support throughout
- [ ] Test on different terminal sizes
- [ ] Test on Windows PowerShell, macOS Terminal, Linux terminals

---

## Testing Requirements

1. **Terminal Compatibility**
   - Test on Windows PowerShell, CMD, Windows Terminal
   - Test on macOS Terminal, iTerm2
   - Test on Linux GNOME Terminal, Konsole, xterm
   - Test in VS Code integrated terminal
   - Test in JetBrains IDE terminals

2. **Responsive Testing**
   - Test at 40 columns (narrow)
   - Test at 80 columns (standard)
   - Test at 120+ columns (wide)
   - Verify no text wrapping issues

3. **Color Support**
   - Test with true color support (24-bit)
   - Test with 256 color support
   - Test with 16 color support
   - Test with `--no-color` flag
   - Test with NO_COLOR environment variable

4. **Accessibility**
   - Verify good contrast ratios
   - Test with screen readers
   - Ensure keyboard navigation works

---

## Final Notes

The goal is to make Clause CLI feel like a **premium, professional tool** from the moment users run it. Every interaction should be polished and delightful. The terminal should feel modern, not like a relic from the 1980s.

**Key differentiators from current design:**
- Visual branding with custom ASCII art
- Color-coded sections and hierarchy
- Box-based layouts with proper spacing
- Grouped commands by category
- Clear call-to-action for new users
- Professional footer with links
- Responsive to terminal width
- Graceful degradation for limited terminals

When implementing, start with the simplest improvements first (colors, spacing, borders) and progressively enhance with more sophisticated features (animations, interactive elements, responsive layouts).
