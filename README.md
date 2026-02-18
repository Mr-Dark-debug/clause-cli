<h1 align="center">Clause CLI</h1>

<p align="center">
  <strong>Framework for Organized, Reproducible, and Guided Engineering</strong>
</p>

<p align="center">
  AI-Native Project Scaffolding for Modern Development Teams
</p>

<p align="center">
  <a href="https://github.com/clause-cli/clause/releases">
    <img src="https://img.shields.io/github/v/release/clause-cli/clause?style=flat-square" alt="Release">
  </a>
  <a href="https://github.com/clause-cli/clause/actions">
    <img src="https://img.shields.io/github/actions/workflow/status/clause-cli/clause/release.yml?style=flat-square" alt="Build Status">
  </a>
  <a href="https://goreportcard.com/report/github.com/clause-cli/clause">
    <img src="https://goreportcard.com/badge/github.com/clause-cli/clause?style=flat-square" alt="Go Report Card">
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/github/license/clause-cli/clause?style=flat-square" alt="License">
  </a>
</p>

---

## What is Clause?

Clause is a **cross-platform CLI tool** that creates **AI-ready project structures** with built-in governance systems. Unlike traditional scaffolding tools that just create folder structures, Clause generates comprehensive AI context systems that guide AI coding assistants toward consistent, maintainable code.

### Why Clause?

- **AI-First Design**: Built from the ground up to work seamlessly with AI coding assistants like Claude, GitHub Copilot, and ChatGPT
- **Governance Built-In**: Every project includes AI context files, component registries, and prompt guidelines
- **Consistent Patterns**: Enforces consistent coding patterns across your team and AI assistants
- **Self-Documenting**: Projects document themselves through the governance system
- **Framework Agnostic**: Supports React, Vue, Svelte, Next.js, FastAPI, Express, Go, and more

---

## Installation

### macOS

```bash
# Using Homebrew (recommended)
brew install clause-cli/tap/clause

# Using curl
curl -fsSL https://get.clause.dev | bash
```

### Linux

```bash
# Using curl
curl -fsSL https://get.clause.dev | bash

# Using AUR (Arch Linux)
yay -S clause-bin

# Using Snap
snap install clause
```

### Windows

```powershell
# Using Winget
winget install Clause.ClauseCLI

# Using Scoop
scoop bucket add clause-cli https://github.com/clause-cli/scoop-bucket
scoop install clause

# Using Chocolatey
choco install clause-cli
```

### From Source

```bash
# Clone the repository
git clone https://github.com/clause-cli/clause.git
cd clause

# Build and install
make install
```

### Verify Installation

```bash
clause version
```

---

## Quick Start

### Create a New Project

```bash
# Launch interactive wizard
clause init

# Create with a specific name
clause init my-awesome-project

# Use a preset
clause init my-saas --preset saas

# Non-interactive mode with defaults
clause init my-api --non-interactive --preset api-only
```

### Available Presets

| Preset | Description |
|--------|-------------|
| `minimal` | Minimal setup with only essentials |
| `standard` | Standard full-stack configuration (default) |
| `saas` | SaaS application with auth, payments, multi-tenancy |
| `api-only` | API-only backend service |
| `frontend-only` | Frontend-only static site |
| `enterprise` | Enterprise configuration with full governance |

### Project Structure

After running `clause init`, your project will have:

```
my-project/
├── .clause/                    # AI Governance
│   ├── config.yaml            # Project configuration
│   ├── context.yaml           # AI context file
│   ├── prompt-guidelines.md   # AI prompt guidelines
│   └── registry.yaml          # Component registry
├── src/                       # Frontend source
│   ├── components/
│   ├── pages/
│   ├── hooks/
│   └── utils/
├── backend/                   # Backend source
│   ├── app/
│   ├── main.py
│   └── requirements.txt
├── infrastructure/            # Infrastructure configs
│   ├── docker-compose.yml
│   └── .github/workflows/
├── Brainstorm.md             # AI brainstorming workspace
├── README.md
└── package.json
```

---

## Features

### 1. Interactive Wizard

The interactive wizard guides you through all configuration options:

- **Project Info**: Name, description, author, license
- **Frontend**: Framework, TypeScript, styling, testing
- **Backend**: Language, framework, database, auth
- **Infrastructure**: Docker, CI/CD, hosting, monitoring
- **Governance**: AI context level, component registry, documentation

### 2. AI Governance System

Every Clause project includes a comprehensive governance system:

#### Context Files

- **`.clause/context.yaml`**: Machine-readable AI context with tech stack, patterns, and conventions
- **`.clause/prompt-guidelines.md`**: Guidelines for working with AI assistants
- **`Brainstorm.md`**: Workspace for AI-assisted brainstorming

#### Component Registry

Track and document your project's components:

```bash
# Register a new component
clause add frontend component Button --description "Reusable button component"

# Register a backend service
clause add backend service UserService --description "User management service"
```

### 3. Multi-Framework Support

#### Frontend Frameworks

| Framework | TypeScript | Styling | Testing |
|-----------|------------|---------|---------|
| React | ✅ | Tailwind, CSS Modules, Styled Components | Vitest, Jest |
| Vue | ✅ | Tailwind, CSS Modules | Vitest |
| Svelte | ✅ | Tailwind | Vitest |
| Angular | ✅ | Tailwind, SCSS | Jest |
| Next.js | ✅ | Tailwind, CSS Modules | Vitest, Jest |

#### Backend Frameworks

| Language | Frameworks | Database | ORM |
|----------|------------|----------|-----|
| Python | FastAPI, Flask, Django | PostgreSQL, MySQL, MongoDB | SQLAlchemy, Prisma |
| Node.js | Express, Fastify, NestJS | PostgreSQL, MySQL, MongoDB | Prisma, TypeORM |
| Go | Gin, Echo, Fiber | PostgreSQL, MySQL | GORM, sqlx |

### 4. Infrastructure Templates

- **Docker & Docker Compose**: Container configuration
- **CI/CD**: GitHub Actions, GitLab CI, CircleCI
- **Hosting**: Vercel, Netlify, AWS, GCP, Azure
- **Monitoring**: Datadog, Sentry, Prometheus

### 5. Project Validation

Validate your project's governance compliance:

```bash
# Run validation
clause validate

# Output as JSON
clause validate --json

# Attempt to fix issues
clause validate --fix
```

---

## Usage Examples

### Example 1: Full-Stack SaaS Application

```bash
# Create SaaS project
clause init my-saas --preset saas

# Navigate to project
cd my-saas

# Install dependencies
npm install
cd backend && pip install -r requirements.txt

# Start development
npm run dev  # Frontend
python main.py  # Backend
```

### Example 2: API-Only Service

```bash
# Create API project
clause init my-api --preset api-only

cd my-api

# Install dependencies
pip install -r requirements.txt

# Start server
python main.py
```

### Example 3: Frontend Application

```bash
# Create frontend-only project
clause init my-portfolio --preset frontend-only

cd my-portfolio

# Install and start
npm install
npm run dev
```

### Example 4: Adding Components

```bash
# Register a new frontend component
clause add frontend component UserProfile \
  --description "User profile display component" \
  --path "src/components/UserProfile" \
  --tags "user,profile,display"

# Register a backend service
clause add backend service EmailService \
  --description "Email sending service" \
  --deps "UserService,TemplateService" \
  --tags "email,notification"
```

---

## Configuration

### Project Configuration (`.clause/config.yaml`)

```yaml
metadata:
  name: my-project
  description: My awesome project
  version: 1.0.0
  author: Your Name

frontend:
  enabled: true
  framework: react
  typescript: true
  styling: tailwind
  package_manager: npm
  test_framework: vitest

backend:
  enabled: true
  language: python
  framework: fastapi
  database:
    primary: postgresql
    orm: sqlalchemy

governance:
  enabled: true
  context_level: comprehensive
  component_registry: true
  brainstorm_md: true
  prompt_guidelines: true
```

### Global Configuration (`~/.clause/config.yaml`)

```yaml
# Default settings for new projects
defaults:
  frontend: react
  backend: fastapi
  styling: tailwind
  database: postgresql

# Update settings
updates:
  check_enabled: true
  channel: stable

# Template registry
templates:
  registry: https://registry.clause.dev
```

---

## CLI Reference

### Global Flags

| Flag | Description |
|------|-------------|
| `--config, -c` | Config file path |
| `--verbose, -v` | Verbose output |
| `--quiet, -q` | Suppress non-essential output |
| `--no-color` | Disable colored output |

### Commands

#### `clause init [project-name]`

Initialize a new project.

```bash
clause init                    # Interactive wizard
clause init my-project         # With name
clause init my-project -p saas # With preset
clause init my-project -n      # Non-interactive
clause init my-project --dry-run # Preview only
```

#### `clause add <type> [name]`

Add a component to an existing project.

```bash
clause add frontend component Button
clause add backend service UserService
clause add governance rule no-any-type
```

#### `clause validate`

Validate project governance compliance.

```bash
clause validate         # Run validation
clause validate --fix   # Fix issues
clause validate --json  # JSON output
```

#### `clause config`

Manage configuration.

```bash
clause config list              # Show all config
clause config get frontend.framework
clause config set defaults.frontend vue
clause config init              # Initialize config
```

#### `clause version`

Display version information.

```bash
clause version
clause version --json
```

#### `clause update`

Update Clause CLI.

```bash
clause update           # Update to latest
clause update --check   # Check only
clause update --channel beta
```

---

## Working with AI Assistants

### Using the Context Files

When working with AI assistants, point them to your governance files:

```
Please read .clause/prompt-guidelines.md and .clause/context.yaml
to understand this project's conventions before making changes.
```

### Brainstorm.md

Use the Brainstorm.md file for AI-assisted exploration:

```markdown
## Ideas

- Add real-time collaboration features
- Implement dark mode toggle

## Questions

- Should we use WebSockets or Server-Sent Events?
- What's the best approach for state management?
```

### Component Registry

Keep your component registry updated so AI assistants understand your architecture:

```yaml
components:
  - name: UserProfile
    type: frontend
    path: src/components/UserProfile
    description: Displays user profile information
    dependencies: [Avatar, UserInfo]
    tags: [user, profile, display]
```

---

## Documentation

- **[Getting Started Guide](docs/getting-started.md)** - Detailed setup and first project
- **[Configuration Reference](docs/configuration.md)** - All configuration options
- **[Governance System](docs/governance.md)** - AI governance in depth
- **[Templates](docs/templates.md)** - Creating custom templates
- **[Contributing](CONTRIBUTING.md)** - How to contribute

---

## Roadmap

### v1.0 (Current)

- [x] Interactive wizard
- [x] Multi-framework support
- [x] AI governance system
- [x] Component registry
- [x] Project validation

### v1.1 (Planned)

- [ ] Custom template creation
- [ ] Plugin system
- [ ] Team collaboration features
- [ ] Template marketplace

### v1.2 (Future)

- [ ] AI-powered code generation
- [ ] Integration with more AI tools
- [ ] Cloud sync for governance files

---

## Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/clause-cli/clause.git
cd clause

# Install dependencies
make deps

# Build
make build

# Run tests
make test

# Run in development mode
make dev
```

---

## License

Clause CLI is released under the [MIT License](LICENSE).

---

## Community

- **GitHub Discussions**: [github.com/clause-cli/clause/discussions](https://github.com/clause-cli/clause/discussions)
- **Twitter**: [@clausecli](https://twitter.com/clausecli)
- **Discord**: [Join our server](https://discord.gg/clause)

---

## Acknowledgments

Built with ❤️ by the Clause team and contributors.

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling

---

<p align="center">
  <a href="https://clause.dev">clause.dev</a>
</p>
