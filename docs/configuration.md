# Configuration

Clause CLI can be configured via a YAML file or environment variables.

## Configuration File

The default location is `~/.clause/config.yaml`.

```yaml
defaults:
  frontend: nextjs       # Default frontend framework
  backend: fastapi       # Default backend framework
  database: postgresql   # Default database
  license: MIT           # Default license

telemetry:
  enabled: true          # Enable anonymous usage stats

updates:
  channel: stable        # specific update channel (stable/beta)
  check_frequency: 24h   # How often to check for updates
```

## Environment Variables

| Variable | Description |
|----------|-------------|
| `CLAUSE_CONFIG` | Path to custom config file |
| `CLAUSE_NO_COLOR` | Disable color output |
| `CLAUSE_DEBUG` | Enable debug logging |

## CLI Commands

You can manage configuration directly from the CLI:

```bash
# View current config
clause config list

# Set a value
clause config set defaults.frontend react

# Reset to defaults
clause config reset
```
