<p align="center">
  <img src="website/assets/images/logo/forge-icon.svg" alt="Clause Logo" width="120" />
</p>

<h1 align="center">Clause CLI</h1>

<p align="center">
  <strong>Structure Your AI's Intelligence</strong>
</p>

<p align="center">
  The AI-native project scaffolding tool that guides your AI coding assistant
  to produce consistent, maintainable, architecturally sound code.
</p>

<p align="center">
  <a href="#-features">Features</a> •
  <a href="#-installation">Installation</a> •
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-documentation">Documentation</a> •
  <a href="#-contributing">Contributing</a>
</p>

<p align="center">
  <a href="https://github.com/Mr-Dark-debug/clause-cli/actions">
    <img src="https://github.com/Mr-Dark-debug/clause-cli/workflows/CI/badge.svg" alt="CI Status" />
  </a>
  <a href="https://github.com/Mr-Dark-debug/clause-cli/releases">
    <img src="https://img.shields.io/github/v/release/Mr-Dark-debug/clause-cli" alt="Release" />
  </a>
  <a href="https://github.com/Mr-Dark-debug/clause-cli/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/Mr-Dark-debug/clause-cli" alt="License" />
  </a>
  <a href="https://goreportcard.com/report/github.com/Mr-Dark-debug/clause-cli">
    <img src="https://goreportcard.com/badge/github.com/Mr-Dark-debug/clause-cli" alt="Go Report Card" />
  </a>
</p>

---

## 🎯 Why Clause?

Modern AI coding assistants are powerful, but their power is undirected. They can write virtually any code, but without proper guidance, that code may violate project standards, introduce inconsistencies, or create technical debt.

**Clause provides the direction that transforms raw AI capability into disciplined, project-aligned engineering output.**

### Key Differentiators

- 🤖 **AI-Native Design** - Built from the ground up for AI-assisted development
- 📋 **Behavioral Governance** - Rules that guide AI behavior within your project
- 🧠 **Brainstorm.md** - Novel self-reflection mechanism for autonomous AI problem-solving
- 📦 **Component Registry** - Living inventory of created components for reusability
- 🎨 **Beautiful TUI** - Modern terminal interface with responsive design

---

## ✨ Features

### Interactive Project Wizard

Clause guides you through project configuration with a beautiful terminal UI:

```bash
$ clause init my-project
```

- Choose frontend framework (Next.js, React, Vue, Svelte)
- Choose backend framework (FastAPI, Express, Django)
- Configure database, authentication, and infrastructure
- Set AI governance strictness level

### AI Governance System

Every Clause-generated project includes an `ai_prompt_guidelines/` directory that:

- Defines how AI should behave in your project
- Specifies approved technologies and forbidden patterns
- Provides architectural constraints
- Includes documentation standards
- Offers a Brainstorm.md file for AI self-reflection

### Cross-Platform Support

Works on macOS, Linux, and Windows with native package managers:

- **macOS**: Homebrew
- **Linux**: APT, Snap, AUR
- **Windows**: Winget, Scoop

---

## 📦 Installation

### macOS

```bash
# Homebrew (recommended)
brew install clause-cli/tap/clause

# Or using curl
curl -fsSL https://clause.dev/install.sh | bash
```

### Linux

```bash
# curl | bash
curl -fsSL https://clause.dev/install.sh | bash

# APT (Debian/Ubuntu)
curl -fsSL https://clause.dev/apt/gpg.key | sudo gpg --dearmor -o /usr/share/keyrings/clause.gpg
echo "deb [signed-by=/usr/share/keyrings/clause.gpg] https://clause.dev/apt stable main" | sudo tee /etc/apt/sources.list.d/clause.list
sudo apt update && sudo apt install clause
```

### Windows

```powershell
# Winget
winget install Clause.ClauseCLI

# Scoop
scoop bucket add clause-cli https://github.com/Mr-Dark-debug/scoop-bucket
scoop install clause

# PowerShell
irm https://clause.dev/install.ps1 | iex
```

### Verify Installation

```bash
clause --version
# Output: clause version 1.0.0
```

---

## 🚀 Quick Start

### Create a New Project

```bash
# Interactive mode (recommended)
clause init my-awesome-project

# With a preset
clause init my-saas --preset saas

# Non-interactive mode
clause init my-project --non-interactive \
  --frontend nextjs \
  --backend fastapi \
  --database postgresql
```

### What Clause Creates

```
my-awesome-project/
├── frontend/                    # Your frontend application
│   ├── src/
│   │   ├── app/
│   │   ├── components/
│   │   └── lib/
│   └── package.json
│
├── backend/                     # Your backend application
│   ├── app/
│   │   ├── api/
│   │   ├── models/
│   │   └── services/
│   └── requirements.txt
│
├── ai_prompt_guidelines/        # AI reads this!
│   ├── system_prompt.md         # Core AI behavior rules
│   ├── architecture.md          # Architectural constraints
│   ├── technologies.md          # Approved tech stack
│   ├── documentation.md         # Documentation standards
│   ├── brainstorm.md            # AI working memory
│   ├── registry.json            # Component registry
│   └── context.yaml             # Structured context
│
├── infrastructure/              # Docker, CI/CD configs
│   ├── docker/
│   └── .github/
│
└── README.md                    # Project documentation
```

### Work with AI Assistants

When you open your Clause-generated project in an AI coding assistant (Claude Code, Cursor, Windsurf, etc.), the AI will automatically find the `ai_prompt_guidelines/` directory and follow the rules defined there.

---

## 📚 Documentation

Full documentation is available at [clause.dev/docs](https://clause.dev/docs)

- [Getting Started](https://clause.dev/docs/getting-started)
- [Installation Guide](https://clause.dev/docs/installation)
- [CLI Reference](https://clause.dev/docs/cli-reference)
- [AI Governance](https://clause.dev/docs/governance)
- [Templates](https://clause.dev/docs/templates)
- [Contributing](https://clause.dev/pages/contributing)

---

## 🛠️ CLI Commands

```bash
# Initialize a new project
clause init [project-name]

# Add components to existing project
clause add [component-type]

# Validate project compliance
clause validate

# Manage configuration
clause config [command]

# Update Clause CLI
clause update

# Show version
clause version
```

Run `clause --help` for more information.

---

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](https://clause.dev/pages/contributing) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/Mr-Dark-debug/clause-cli.git
cd clause-cli

# Build
make build

# Run tests
make test

# Run locally
./bin/clause init test-project
```

### Project Structure

```
clause-cli/
├── cmd/forge/              # Entry point
├── internal/               # Private packages
│   ├── cmd/                # CLI commands
│   ├── wizard/             # Interactive wizard
│   ├── config/             # Configuration management
│   ├── template/           # Template engine
│   ├── generator/          # Project generation
│   └── governance/         # AI governance system
├── pkg/                    # Public packages
│   ├── styles/             # Theme and styling
│   ├── tui/                # Terminal UI components
│   └── utils/              # Utilities
└── templates/              # Project templates
```

---

## 📝 License

Clause CLI is released under the [MIT License](LICENSE).

---

## 🙏 Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling

---

## 📬 Connect

- **GitHub**: [github.com/Mr-Dark-debug/clause-cli](https://github.com/Mr-Dark-debug/clause-cli)
- **Discord**: [Join our community](https://discord.gg/clause)
- **Twitter**: [@ClauseCLI](https://twitter.com/ClauseCLI)

---

<p align="center">
  Made with ❤️ for the developer community
</p>
