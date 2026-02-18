# Installation Guide

## macOS

### Homebrew (Recommended)

```bash
brew install clause-cli/tap/clause
```

### Manual

```bash
curl -fsSL https://clause.dev/install.sh | bash
```

## Linux

### APT (Debian/Ubuntu)

```bash
curl -fsSL https://clause.dev/apt/gpg.key | sudo gpg --dearmor -o /usr/share/keyrings/clause.gpg
echo "deb [signed-by=/usr/share/keyrings/clause.gpg] https://clause.dev/apt stable main" | sudo tee /etc/apt/sources.list.d/clause.list
sudo apt update && sudo apt install clause
```

### Snap

```bash
sudo snap install clause
```

### Script

```bash
curl -fsSL https://clause.dev/install.sh | bash
```

## Windows

### Winget

```powershell
winget install Clause.ClauseCLI
```

### Scoop

```powershell
scoop bucket add clause-cli https://github.com/Mr-Dark-debug/scoop-bucket
scoop install clause
```

### PowerShell

```powershell
irm https://clause.dev/install.ps1 | iex
```

## Building from Source

Requirements: Go 1.21+

```bash
git clone https://github.com/Mr-Dark-debug/clause-cli.git
cd clause-cli
make build
# Binary will be in ./bin/clause
```
