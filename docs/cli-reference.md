# CLI Reference

Complete command reference for Clause CLI.

## Global Flags

```bash
clause [global flags] <command> [command flags] [arguments]
```

| Flag | Short | Description |
|------|-------|-------------|
| `--config` | `-c` | Path to configuration file |
| `--verbose` | `-v` | Enable verbose output |
| `--quiet` | `-q` | Suppress non-essential output |
| `--no-color` | | Disable colored output |
| `--help` | `-h` | Show help |
| `--version` | | Show version |

---

## clause init

Initialize a new project with AI governance.

### Usage

```bash
clause init [project-name] [flags]
```

### Arguments

| Argument | Description | Required |
|----------|-------------|----------|
| `project-name` | Name of the project to create | Yes |

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--preset` | | Use a preset configuration |
| `--non-interactive` | `false` | Skip interactive wizard |
| `--frontend` | | Frontend framework (nextjs, react, vue, svelte, none) |
| `--backend` | | Backend framework (fastapi, express, django, none) |
| `--database` | | Database type (postgresql, mongodb, mysql, none) |
| `--output` | `.` | Output directory |

### Examples

```bash
# Interactive mode
clause init my-project

# With preset
clause init my-saas --preset saas

# Non-interactive with all options
clause init my-project --non-interactive \
  --frontend nextjs \
  --backend fastapi \
  --database postgresql

# In current directory
clause init . --frontend react --backend express
```

### Presets

| Preset | Frontend | Backend | Database |
|--------|----------|---------|----------|
| `saas` | Next.js | FastAPI | PostgreSQL |
| `startup` | React | Express | MongoDB |
| `enterprise` | Next.js | Django | PostgreSQL |
| `minimal` | React | FastAPI | None |

---

## clause add

Add components to an existing project.

### Usage

```bash
clause add <component-type> [flags]
```

### Component Types

| Type | Description |
|------|-------------|
| `frontend` | Add frontend components |
| `backend` | Add backend modules |
| `governance` | Add governance rules |
| `infrastructure` | Add infrastructure configs |

### Examples

```bash
# Add a frontend component
clause add frontend --component auth

# Add backend module
clause add backend --module users

# Add governance rule
clause add governance --rule no-any-type

# Add Docker configuration
clause add infrastructure --docker
```

---

## clause validate

Validate project for governance compliance.

### Usage

```bash
clause validate [path] [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--strict` | `false` | Enable strict validation |
| `--fix` | `false` | Auto-fix issues where possible |
| `--output` | `text` | Output format (text, json) |

### Examples

```bash
# Validate current directory
clause validate

# Validate specific project
clause validate ./my-project

# Strict validation with JSON output
clause validate --strict --output json

# Auto-fix issues
clause validate --fix
```

### Exit Codes

| Code | Meaning |
|------|---------|
| 0 | All checks passed |
| 1 | Validation errors found |
| 2 | Configuration error |

---

## clause config

Manage Clause configuration.

### Usage

```bash
clause config <command> [flags]
```

### Subcommands

| Command | Description |
|---------|-------------|
| `list` | List all configuration values |
| `get` | Get a specific configuration value |
| `set` | Set a configuration value |
| `init` | Initialize configuration file |

### Examples

```bash
# List all config
clause config list

# Get specific value
clause config get default.frontend

# Set a value
clause config set default.frontend nextjs

# Initialize config file
clause config init
```

---

## clause update

Update Clause CLI to the latest version.

### Usage

```bash
clause update [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--channel` | `stable` | Update channel (stable, beta, nightly) |
| `--check` | `false` | Only check for updates, don't install |
| `--force` | `false` | Force update even if already latest |

### Examples

```bash
# Update to latest stable
clause update

# Check for updates
clause update --check

# Update to beta channel
clause update --channel beta
```

---

## clause version

Display version information.

### Usage

```bash
clause version [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--json` | `false` | Output in JSON format |

### Examples

```bash
clause version
# Output: clause version 1.0.0 (built 2025-01-15)

clause version --json
# Output: {"version": "1.0.0", "buildDate": "2025-01-15", ...}
```

---

## clause completion

Generate shell completion scripts.

### Usage

```bash
clause completion <shell>
```

### Shells

| Shell | Description |
|-------|-------------|
| `bash` | Bash completion |
| `zsh` | Zsh completion |
| `fish` | Fish completion |
| `powershell` | PowerShell completion |

### Examples

```bash
# Generate bash completion
clause completion bash > /etc/bash_completion.d/clause

# Generate zsh completion
clause completion zsh > "${fpath[1]}/_clause"

# Add to PowerShell profile
clause completion powershell >> $PROFILE
```

---

## Environment Variables

| Variable | Description |
|----------|-------------|
| `CLAUSE_CONFIG` | Path to configuration file |
| `CLAUSE_NO_COLOR` | Disable colored output |
| `CLAUSE_CACHE_DIR` | Cache directory location |

---

## Configuration File

Configuration is stored in YAML format:

```yaml
# ~/.clause/config.yaml

defaults:
  frontend: nextjs
  backend: fastapi
  database: postgresql

presets:
  my-custom:
    frontend: react
    backend: express
    styling: tailwind

updates:
  channel: stable
  checkOnStartup: true
```

---

## Exit Codes

All commands use standard exit codes:

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Error |
| 2 | Configuration error |
| 130 | Interrupted (Ctrl+C) |
