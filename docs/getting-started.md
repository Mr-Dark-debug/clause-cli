# Getting Started with Clause CLI

Welcome to Clause CLI! This guide will walk you through creating your first AI-ready project.

---

## Prerequisites

Before you begin, ensure you have:

- **Go 1.21+** (for building from source)
- **Node.js 18+** (for frontend projects)
- **Python 3.11+** (for Python backend projects)
- **Git** (for version control)

---

## Installation

### Quick Install

```bash
# macOS/Linux
curl -fsSL https://get.clause.dev | bash

# Windows (PowerShell)
winget install Clause.ClauseCLI
```

### Verify Installation

```bash
clause version
```

You should see:
```
Clause - AI-Native Project Scaffolding

  Version:    x.x.x
  Build Time: ...
  Commit:     ...
  Go Version: go1.21.x
  Platform:   darwin/arm64
```

---

## Your First Project

### 1. Launch the Wizard

```bash
clause init
```

You'll see an interactive wizard:

```
   ██████╗██╗      █████╗ ██╗   ██╗██████╗ ███████╗
  ██╔════╝██║     ██╔══██╗██║   ██║██╔══██╗██╔════╝
  ██║     ██║     ███████║██║   ██║██║  ██║█████╗
  ██║     ██║     ██╔══██║██║   ██║██║  ██║██╔══╝
  ╚██████╗███████╗██║  ██║╚██████╔╝██████╔╝███████╗
   ╚═════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝

  AI-Native Project Scaffolding
```

### 2. Choose a Preset (Optional)

On the first screen, you can select a preset:

| Preset | Best For |
|--------|----------|
| **Minimal** | Quick prototypes |
| **Standard** | General full-stack apps |
| **SaaS** | Multi-tenant applications |
| **API Only** | Backend services |
| **Frontend Only** | Static sites, SPAs |
| **Enterprise** | Large-scale applications |

### 3. Configure Your Project

Follow the wizard through each screen:

1. **Project Info** - Name, description, author
2. **Frontend** - Framework, TypeScript, styling
3. **Backend** - Language, database, API style
4. **Infrastructure** - Docker, CI/CD, hosting
5. **Governance** - AI context level, documentation

### 4. Review and Create

The final screen shows a summary. Review your choices and press Enter to create the project.

---

## Quick Start with Presets

Skip the wizard by using a preset:

```bash
# Create a SaaS application
clause init my-saas --preset saas

# Create an API-only project
clause init my-api --preset api-only

# Create a frontend-only project
clause init my-portfolio --preset frontend-only

# Non-interactive with defaults
clause init my-project --non-interactive
```

---

## Project Structure

After creation, your project will have:

```
my-project/
├── .clause/                    # AI Governance
│   ├── config.yaml            # Project configuration
│   ├── context.yaml           # AI context
│   ├── prompt-guidelines.md   # AI guidelines
│   └── registry.yaml          # Component registry
├── src/                       # Frontend source
│   ├── components/
│   ├── pages/
│   ├── hooks/
│   └── utils/
├── backend/                   # Backend source
│   ├── app/
│   ├── main.py (or main.go / index.js)
│   └── requirements.txt
├── infrastructure/
│   ├── docker-compose.yml
│   └── .github/workflows/
├── Brainstorm.md             # AI brainstorming
├── README.md
└── package.json
```

---

## Next Steps

### 1. Navigate to Your Project

```bash
cd my-project
```

### 2. Install Dependencies

```bash
# Frontend
npm install

# Backend (Python)
cd backend
pip install -r requirements.txt
cd ..

# Backend (Go)
cd backend
go mod download
cd ..
```

### 3. Start Development

```bash
# Frontend
npm run dev

# Backend (Python)
python backend/main.py

# Backend (Go)
go run backend/main.go
```

### 4. Use with AI Assistants

Point your AI assistant to the governance files:

```
Please read .clause/prompt-guidelines.md and .clause/context.yaml
before making changes to this project.
```

---

## Common Commands

```bash
# Create project with wizard
clause init

# Create with preset
clause init my-project --preset saas

# Create non-interactively
clause init my-project --non-interactive

# Preview without creating files
clause init my-project --dry-run

# Add a component to existing project
clause add frontend component Button

# Validate project
clause validate

# Check CLI version
clause version
```

---

## Working with AI Assistants

### Context Files

Clause generates files that help AI assistants understand your project:

| File | Purpose |
|------|---------|
| `.clause/context.yaml` | Machine-readable project context |
| `.clause/prompt-guidelines.md` | Coding standards and patterns |
| `.clause/registry.yaml` | Component documentation |
| `Brainstorm.md` | AI-assisted brainstorming |

### Best Practices

1. **Keep context updated**: Run `clause add` when creating components
2. **Reference guidelines**: Ask AI to read guidelines before making changes
3. **Use Brainstorm.md**: Document ideas and decisions for AI context

---

## Getting Help

- **Documentation**: [docs.clause.dev](https://docs.clause.dev)
- **GitHub**: [github.com/clause-cli/clause](https://github.com/clause-cli/clause)
- **Discord**: [discord.gg/clause](https://discord.gg/clause)

---

## What's Next?

- Read the [Configuration Guide](configuration.md)
- Learn about the [Governance System](governance.md)
- Explore [Templates](templates.md)
