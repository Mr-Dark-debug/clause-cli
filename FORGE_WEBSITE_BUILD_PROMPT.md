# 🌐 Clause CLI - Professional Landing Website Build Instructions

---

## 📋 PROJECT OVERVIEW

You are building a **professional, production-ready landing website** for Clause CLI - an AI-native project scaffolding tool. This website will serve as the primary online presence for the project, hosting documentation, installation guides, and community resources.

**GitHub Repository:** https://github.com/Mr-Dark-debug/clause-cli
**Project Name:** Clause (Framework for Organized, Reproducible, and Guided Engineering)
**CLI Command:** `Clause`

---

## 🎯 WEBSITE REQUIREMENTS

### Core Requirements

1. **Fully Responsive** - Must work perfectly on all devices:
   - Mobile phones (320px - 480px)
   - Tablets (481px - 1024px)
   - Laptops (1025px - 1440px)
   - Desktops (1441px+)
   - 4K displays

2. **Professional UI/UX** - Modern, clean, developer-focused design
3. **Fast Loading** - Optimized assets, minimal dependencies
4. **Accessible** - WCAG 2.1 AA compliance
5. **SEO Optimized** - Proper meta tags, semantic HTML
6. **Dark/Light Mode** - Theme toggle with system preference detection

---

## 📁 PROJECT STRUCTURE

Create this EXACT structure in the repository:

```
clause-cli/
├── index.html                  # Main landing page (ROOT LEVEL)
├── website/
│   ├── docs/
│   ├── index.html               # Documentation home
│   ├── getting-started.html     # Getting started guide
│   ├── installation.html        # Installation instructions
│   ├── configuration.html       # Configuration reference
│   ├── templates.html           # Templates documentation
│   ├── governance.html          # AI Governance documentation
│   ├── cli-reference.html       # CLI commands reference
│   └── api.html                 # API documentation (future)
├── pages/
│   ├── changelog.html           # Version history & changes
│   ├── roadmap.html             # Project roadmap
│   ├── contributing.html        # Contribution guidelines
│   ├── community.html           # Community resources
│   ├── sponsors.html            # Sponsors & backers
│   └── license.html             # License information
│---├── assets/
│   │   ├── css/
│   │   │   ├── main.css         # Main stylesheet
│   │   │   ├── variables.css    # CSS custom properties
│   │   │   ├── components.css   # Reusable components
│   │   │   ├── layouts.css      # Layout styles
│   │   │   ├── utilities.css    # Utility classes
│   │   │   ├── animations.css   # Animation definitions
│   │   │   └── prism.css        # Code syntax highlighting
│   │   │
│   │   ├── js/
│   │   │   ├── main.js          # Main JavaScript
│   │   │   ├── navigation.js    # Navigation functionality
│   │   │   ├── theme.js         # Theme toggle logic
│   │   │   ├── copy-code.js     # Copy code functionality
│   │   │   ├── tabs.js          # Tab components
│   │   │   ├── search.js        # Search functionality
│   │   │   └── animations.js    # Scroll animations
│   │   │
│   │   ├── images/
│   │   │   ├── logo/
│   │   │   │   ├── Clause-logo.svg        # Main logo
│   │   │   │   ├── Clause-logo-dark.svg   # Dark mode logo
│   │   │   │   ├── Clause-icon.svg        # Icon only
│   │   │   │   └── Clause-wordmark.svg    # Wordmark only
│   │   │   │
│   │   │   ├── og/
│   │   │   │   ├── og-image.png          # Open Graph image (1200x630)
│   │   │   │   ├── og-image-dark.png     # Dark mode OG image
│   │   │   │   └── twitter-card.png      # Twitter card image
│   │   │   │
│   │   │   ├── icons/
│   │   │   │   ├── favicon.ico           # Favicon
│   │   │   │   ├── favicon-16x16.png
│   │   │   │   ├── favicon-32x32.png
│   │   │   │   ├── apple-touch-icon.png
│   │   │   │   ├── android-chrome-192x192.png
│   │   │   │   ├── android-chrome-512x512.png
│   │   │   │   └── safari-pinned-tab.svg
│   │   │   │
│   │   │   ├── screenshots/
│   │   │   │   ├── hero-demo.png         # Hero section demo
│   │   │   │   ├── wizard-preview.png    # Wizard screenshots
│   │   │   │   ├── terminal-demo.png     # Terminal output demo
│   │   │   │   └── governance-preview.png
│   │   │   │
│   │   │   └── illustrations/
│   │   │       ├── architecture-diagram.svg
│   │   │       ├── workflow-diagram.svg
│   │   │       └── features-illustration.svg
│   │   │
│   │   └── fonts/
│   │       ├── Inter-Variable.woff2
│   │       ├── JetBrainsMono-Variable.woff2
│   │       └── fonts.css
│   │
│   └── components/
│       ├── header.html           # Navigation header
│       ├── footer.html           # Site footer
│       ├── sidebar.html          # Documentation sidebar
│       └── toc.html              # Table of contents
│
├── _headers                      # Cloudflare Pages headers
├── _redirects                    # Netlify redirects
├── robots.txt                    # SEO robots file
├── sitemap.xml                   # SEO sitemap
├── manifest.json                 # PWA manifest
└── sw.js                         # Service worker (optional)
```

---

## 🎨 DESIGN SYSTEM

### Color Palette

```css
/* Light Mode */
--color-primary: #FF6B35;           /* Brand orange */
--color-primary-hover: #E55A2B;     /* Darker orange for hover */
--color-primary-light: #FFF4F0;     /* Light orange background */

--color-secondary: #1A1B26;         /* Deep navy */
--color-secondary-light: #2D2E3E;   /* Lighter navy */

--color-background: #FFFFFF;        /* Main background */
--color-surface: #F8F9FA;           /* Card backgrounds */
--color-surface-elevated: #FFFFFF;  /* Elevated surfaces */

--color-text-primary: #1A1B26;      /* Primary text */
--color-text-secondary: #4B5563;    /* Secondary text */
--color-text-muted: #8B949E;        /* Muted text */

--color-border: #E5E7EB;            /* Borders */
--color-border-light: #F3F4F6;      /* Light borders */

--color-success: #10B981;           /* Success */
--color-warning: #F59E0B;           /* Warning */
--color-error: #EF4444;             /* Error */
--color-info: #3B82F6;              /* Info */

/* Dark Mode */
--color-background-dark: #0D1117;   /* Main background */
--color-surface-dark: #161B22;      /* Card backgrounds */
--color-surface-elevated-dark: #21262D; /* Elevated surfaces */

--color-text-primary-dark: #F0F6FC; /* Primary text */
--color-text-secondary-dark: #8B949E; /* Secondary text */
--color-text-muted-dark: #484F58;   /* Muted text */

--color-border-dark: #30363D;       /* Borders */
--color-border-light-dark: #21262D; /* Light borders */
```

### Typography

```css
/* Font Families */
--font-sans: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', sans-serif;
--font-mono: 'JetBrains Mono', 'Fira Code', 'Consolas', monospace;

/* Font Sizes */
--text-xs: 0.75rem;      /* 12px */
--text-sm: 0.875rem;     /* 14px */
--text-base: 1rem;       /* 16px */
--text-lg: 1.125rem;     /* 18px */
--text-xl: 1.25rem;      /* 20px */
--text-2xl: 1.5rem;      /* 24px */
--text-3xl: 1.875rem;    /* 30px */
--text-4xl: 2.25rem;     /* 36px */
--text-5xl: 3rem;        /* 48px */
--text-6xl: 3.75rem;     /* 60px */

/* Font Weights */
--font-normal: 400;
--font-medium: 500;
--font-semibold: 600;
--font-bold: 700;

/* Line Heights */
--leading-tight: 1.25;
--leading-normal: 1.5;
--leading-relaxed: 1.625;
```

### Spacing Scale

```css
--space-0: 0;
--space-1: 0.25rem;    /* 4px */
--space-2: 0.5rem;     /* 8px */
--space-3: 0.75rem;    /* 12px */
--space-4: 1rem;       /* 16px */
--space-5: 1.25rem;    /* 20px */
--space-6: 1.5rem;     /* 24px */
--space-8: 2rem;       /* 32px */
--space-10: 2.5rem;    /* 40px */
--space-12: 3rem;      /* 48px */
--space-16: 4rem;      /* 64px */
--space-20: 5rem;      /* 80px */
--space-24: 6rem;      /* 96px */
```

### Border Radius

```css
--radius-sm: 0.25rem;    /* 4px */
--radius-md: 0.5rem;     /* 8px */
--radius-lg: 0.75rem;    /* 12px */
--radius-xl: 1rem;       /* 16px */
--radius-2xl: 1.5rem;    /* 24px */
--radius-full: 9999px;
```

### Shadows

```css
/* Light Mode Shadows */
--shadow-sm: 0 1px 2px 0 rgba(0, 0, 0, 0.05);
--shadow-md: 0 4px 6px -1px rgba(0, 0, 0, 0.1);
--shadow-lg: 0 10px 15px -3px rgba(0, 0, 0, 0.1);
--shadow-xl: 0 20px 25px -5px rgba(0, 0, 0, 0.1);

/* Dark Mode Shadows */
--shadow-sm-dark: 0 1px 2px 0 rgba(0, 0, 0, 0.3);
--shadow-md-dark: 0 4px 6px -1px rgba(0, 0, 0, 0.4);
--shadow-lg-dark: 0 10px 15px -3px rgba(0, 0, 0, 0.5);
```

---

## 📄 PAGE SPECIFICATIONS

### 1. index.html - Main Landing Page

#### Hero Section
```
┌─────────────────────────────────────────────────────────────────────────┐
│  [Logo] Clause              Docs  Installation  Community  [GitHub]     │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│                                                                         │
│                 🔥 Clause                                                │
│                                                                         │
│         Structure Your AI's Intelligence                                │
│                                                                         │
│   The AI-native project scaffolding tool that guides your AI           │
│   coding assistant to produce consistent, maintainable code            │
│                                                                         │
│   [Get Started]  [View on GitHub]                                      │
│                                                                         │
│   ┌─────────────────────────────────────────────────────────────┐     │
│   │  $ Clause init my-project                                     │     │
│   │  ✓ Creating project structure...                             │     │
│   │  ✓ Setting up AI governance...                               │     │
│   │  ✓ Your project is ready!                                    │     │
│   └─────────────────────────────────────────────────────────────┘     │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

**Hero Section Requirements:**
- Large, prominent logo with gradient or glow effect
- Clear tagline: "Structure Your AI's Intelligence"
- Two CTAs: Primary (Get Started) and Secondary (View on GitHub)
- Animated terminal preview showing `Clause init` command
- Subtle background gradient or pattern
- Particle/float animation (optional, subtle)

#### Features Section

Display 4-6 key features in a grid:

```
┌─────────────────────────────────────────────────────────────────────────┐
│                       Why Choose Clause?                                 │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐      │
│  │ 🤖 AI-Native     │  │ 📋 Governance    │  │ 🧠 Brainstorm   │      │
│  │                  │  │                  │  │                  │      │
│  │ Built for AI-    │  │ Rules that guide │  │ Self-reflection  │      │
│  │ assisted coding  │  │ AI behavior      │  │ for AI agents    │      │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘      │
│                                                                         │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐      │
│  │ ⚡ Cross-Platform│  │ 🎨 Beautiful TUI │  │ 🔧 Extensible    │      │
│  │                  │  │                  │  │                  │      │
│  │ macOS, Linux,    │  │ Modern terminal  │  │ Plugin system    │      │
│  │ Windows support  │  │ interface        │  │ for custom rules │      │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘      │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

**Feature Cards Design:**
- Icon/emoji at top (or use SVG icons)
- Feature title in bold
- Brief description (1-2 lines)
- Hover effect: subtle lift and glow
- Click to navigate to detailed doc page

#### Quick Start Section

```
┌─────────────────────────────────────────────────────────────────────────┐
│                       Quick Start                                       │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  Install Clause in seconds:                                              │
│                                                                         │
│  [Tabs: macOS | Linux | Windows]                                        │
│                                                                         │
│  ┌──────────────────────────────────────────────────────────────────┐  │
│  │ # Install via Homebrew                                           │  │
│  │ $ brew install Clause-cli/tap/Clause                               │  │
│  │                                                                  │  │
│  │ # Or using curl                                                  │  │
│  │ $ curl -fsSL https://Clause.dev/install.sh | bash                │  │
│  │                                                                  │  │
│  │ # Initialize a new project                                       │  │
│  │ $ Clause init my-project                                          │  │
│  └──────────────────────────────────────────────────────────────────┘  │
│                                                                         │
│  [Copy] button for each code block                                      │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

#### Code Preview Section

Show what Clause generates:

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    What Clause Creates                                   │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  [Interactive file tree explorer]                                       │
│                                                                         │
│  📁 my-project/                                                         │
│  ├── 📁 frontend/                                                       │
│  │   ├── 📁 src/                                                        │
│  │   ├── 📁 components/                                                 │
│  │   └── 📄 package.json                                                │
│  ├── 📁 backend/                                                        │
│  │   ├── 📁 api/                                                        │
│  │   ├── 📁 models/                                                     │
│  │   └── 📄 requirements.txt                                            │
│  ├── 📁 ai_prompt_guidelines/    ← AI reads this!                       │
│  │   ├── 📄 system_prompt.md                                            │
│  │   ├── 📄 architecture.md                                             │
│  │   ├── 📄 brainstorm.md                                               │
│  │   └── 📄 component_registry.json                                     │
│  └── 📄 README.md                                                       │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

#### Testimonials/Social Proof Section (Optional)

```
┌─────────────────────────────────────────────────────────────────────────┐
│                    Loved by Developers                                  │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  ⭐ 2.5k+ GitHub Stars                                                  │
│  📦 50k+ Downloads                                                      │
│  🌍 Used in 100+ countries                                              │
│                                                                         │
│  ┌──────────────────────┐  ┌──────────────────────┐                    │
│  │ "Clause transformed   │  │ "Finally, AI that    │                    │
│  │  how we use AI in    │  │  follows our rules!" │                    │
│  │  our team..."        │  │                      │                    │
│  │                      │  │                      │                    │
│  │  - @developer        │  │  - @another_dev      │                    │
│  └──────────────────────┘  └──────────────────────┘                    │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

#### CTA Section

```
┌─────────────────────────────────────────────────────────────────────────┐
│                                                                         │
│           Ready to structure your AI's intelligence?                    │
│                                                                         │
│                    [Get Started Now]                                    │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

#### Footer

```
┌─────────────────────────────────────────────────────────────────────────┐
│                                                                         │
│  🔥 Clause                          Resources        Community          │
│                                    Documentation    GitHub              │
│  Structure Your AI's Intelligence  Installation     Discord            │
│                                    Templates        Twitter             │
│                                    Changelog        Contributing        │
│                                                                         │
│  ─────────────────────────────────────────────────────────────────────  │
│                                                                         │
│  © 2025 Clause CLI. Released under the MIT License.                     │
│  Built with ❤️ for the developer community                              │
│                                                                         │
│  [GitHub] [Discord] [Twitter]                                           │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

### 2. docs/index.html - Documentation Home

```
┌─────────────────────────────────────────────────────────────────────────┐
│  [Logo] Clause CLI Documentation                                         │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  [Search documentation...]                                    🔍        │
│                                                                         │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  Getting Started                                                        │
│  ────────────────                                                       │
│  Everything you need to know to get started with Clause                 │
│                                                                         │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐      │
│  │ 📦 Installation  │  │ 🚀 Quick Start   │  │ ⚙️ Configuration │      │
│  │                  │  │                  │  │                  │      │
│  │ Install on any   │  │ Create your      │  │ Customize your   │      │
│  │ platform         │  │ first project    │  │ setup            │      │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘      │
│                                                                         │
│  Core Concepts                                                          │
│  ────────────────                                                       │
│                                                                         │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐      │
│  │ 🤖 AI Governance │  │ 📝 Templates     │  │ 🧩 Components    │      │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘      │
│                                                                         │
│  ┌──────────────────┐  ┌──────────────────┐  ┌──────────────────┐      │
│  │ 🔄 Brainstorm.md │  │ 🔌 Plugins      │  │ ✅ Validation    │      │
│  └──────────────────┘  └──────────────────┘  └──────────────────┘      │
│                                                                         │
│  Reference                                                              │
│  ────────────────                                                       │
│                                                                         │
│  ┌──────────────────┐  ┌──────────────────┐                            │
│  │ 📚 CLI Reference │  │ 🎨 Theming       │                            │
│  └──────────────────┘  └──────────────────┘                            │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

### 3. docs/installation.html - Installation Guide

**Must include:**

1. **Requirements Section**
   - Go 1.21+ (for building from source)
   - Git (optional, for version control)
   - Terminal/Command prompt

2. **Installation Methods** (with tabs)

**macOS:**
```bash
# Homebrew (recommended)
brew install Clause-cli/tap/Clause

# curl | bash
curl -fsSL https://Clause.dev/install.sh | bash

# Manual download
# Download from https://github.com/Mr-Dark-debug/clause-cli/releases
# Move to /usr/local/bin
```

**Linux:**
```bash
# curl | bash
curl -fsSL https://Clause.dev/install.sh | bash

# APT (Debian/Ubuntu)
curl -fsSL https://Clause.dev/apt/gpg.key | sudo gpg --dearmor -o /usr/share/keyrings/Clause.gpg
echo "deb [signed-by=/usr/share/keyrings/Clause.gpg] https://Clause.dev/apt stable main" | sudo tee /etc/apt/sources.list.d/Clause.list
sudo apt update && sudo apt install Clause

# Snap
sudo snap install Clause

# Arch Linux (AUR)
yay -S Clause-cli
```

**Windows:**
```powershell
# Winget
winget install Clause.ClauseCLI

# Scoop
scoop bucket add Clause-cli https://github.com/Mr-Dark-debug/scoop-bucket
scoop install Clause

# Chocolatey
choco install Clause-cli

# PowerShell
irm https://Clause.dev/install.ps1 | iex
```

3. **Verify Installation**
```bash
Clause --version
# Output: Clause version 1.0.0
```

4. **Updating**
```bash
# Self-update
Clause update

# Or via package manager
brew upgrade Clause
```

5. **Uninstallation**
```bash
# Homebrew
brew uninstall Clause

# Manual
rm /usr/local/bin/Clause
```

---

### 4. docs/getting-started.html - Getting Started Guide

**Content Structure:**

1. **Introduction**
   - What Clause does
   - Who should use it
   - Key concepts

2. **Your First Project**
   - Step-by-step tutorial
   - Screenshots of the wizard
   - Expected output

3. **Understanding the Output**
   - File structure explanation
   - Key files and their purposes

4. **Next Steps**
   - Links to advanced topics

---

### 5. docs/governance.html - AI Governance Documentation

**This is a KEY page explaining Clause's unique value:**

1. **What is AI Governance?**
   - Why it matters
   - How Clause approaches it

2. **The ai_prompt_guidelines Directory**
   - Structure overview
   - File explanations

3. **System Prompts**
   - How they work
   - Customization options

4. **Brainstorm.md System**
   - Self-reflection mechanism
   - How AI uses it
   - Best practices

5. **Component Registry**
   - Purpose and usage
   - Maintaining the registry

6. **Custom Governance Rules**
   - Creating custom rules
   - Rule syntax and examples

---

### 6. docs/cli-reference.html - CLI Commands Reference

**Complete command documentation:**

```
Clause [global flags] <command> [command flags] [arguments]

Global Flags:
  -c, --config string   Config file path
  -v, --verbose         Verbose output
  -q, --quiet           Suppress non-essential output
      --no-color        Disable colored output
  -h, --help            Help for Clause

Commands:
  init        Initialize a new project
  add         Add components to existing project
  update      Update Clause to latest version
  validate    Validate project governance compliance
  config      Manage configuration
  version     Print version information
  completion  Generate shell completion
  help        Help about any command
```

Each command should have its own section with:
- Usage syntax
- Flags and their descriptions
- Examples
- Related commands

---

### 7. pages/changelog.html - Version History

**Structure:**

```markdown
# Changelog

All notable changes to Clause are documented here.

## [1.0.0] - 2025-01-15

### Added
- Initial release
- Interactive TUI wizard
- AI governance system
- Brainstorm.md self-reflection
- Cross-platform support (macOS, Linux, Windows)

### Changed
- N/A

### Fixed
- N/A

## [0.9.0] - 2025-01-01

### Added
- Beta release
- Core functionality
...
```

Use semantic versioning and follow [Keep a Changelog](https://keepachangelog.com/) format.

---

### 8. pages/roadmap.html - Project Roadmap

**Visual roadmap with phases:**

```
┌─────────────────────────────────────────────────────────────────────────┐
│                           Clause Roadmap                                 │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  Phase 1: Foundation ✅ COMPLETE                                        │
│  ─────────────────────────────                                          │
│  • Core CLI framework                                                   │
│  • Interactive wizard                                                   │
│  • Basic templates                                                      │
│  • AI governance foundation                                             │
│                                                                         │
│  Phase 2: Enhancement 🔄 IN PROGRESS                                    │
│  ─────────────────────────────                                          │
│  • Plugin system                                                        │
│  • Component registry                                                   │
│  • Validation system                                                    │
│  • Extended templates                                                   │
│                                                                         │
│  Phase 3: Growth 📋 PLANNED                                             │
│  ─────────────────────────────                                          │
│  • Template marketplace                                                 │
│  • Team collaboration                                                   │
│  • Cloud sync                                                           │
│                                                                         │
│  Phase 4: Scale 🔮 FUTURE                                               │
│  ─────────────────────────────                                          │
│  • Enterprise features                                                  │
│  • IDE integrations                                                     │
│  • Advanced governance                                                  │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

### 9. pages/contributing.html - Contribution Guidelines

**Content:**

1. **Code of Conduct**
2. **How to Contribute**
   - Reporting bugs
   - Suggesting features
   - Pull request process
3. **Development Setup**
   - Prerequisites
   - Building from source
   - Running tests
4. **Coding Standards**
   - Code style
   - Documentation requirements
   - Testing requirements
5. **Commit Guidelines**
   - Commit message format
   - Conventional commits
6. **Release Process**

---

### 10. pages/community.html - Community Resources

**Content:**

1. **Join the Community**
   - Discord server link
   - Twitter/X follow
   - GitHub discussions

2. **Getting Help**
   - Documentation
   - FAQ
   - Issue tracker

3. **Stay Updated**
   - Newsletter signup
   - Blog/RSS
   - Changelog

4. **Contributing**
   - Link to contributing page

---

## 🧩 REUSABLE COMPONENTS

### Header Component

```html
<header class="site-header">
  <div class="container">
    <a href="/" class="logo">
      <img src="website/assets/images/logo/Clause-icon.svg" alt="Clause" />
      <span>Clause</span>
    </a>
    
    <nav class="main-nav">
      <a href="/docs">Documentation</a>
      <a href="/docs/installation">Installation</a>
      <a href="/pages/changelog">Changelog</a>
      <a href="/pages/community">Community</a>
    </nav>
    
    <div class="header-actions">
      <button class="theme-toggle" aria-label="Toggle theme">
        <span class="icon-sun">☀️</span>
        <span class="icon-moon">🌙</span>
      </button>
      <a href="https://github.com/Mr-Dark-debug/clause-cli" class="btn btn-github">
        <svg><!-- GitHub icon --></svg>
        GitHub
      </a>
    </div>
    
    <button class="mobile-menu-toggle" aria-label="Toggle menu">
      <span></span>
      <span></span>
      <span></span>
    </button>
  </div>
</header>
```

### Footer Component

```html
<footer class="site-footer">
  <div class="container">
    <div class="footer-grid">
      <div class="footer-brand">
        <img src="website/assets/images/logo/Clause-icon.svg" alt="Clause" />
        <p>Structure Your AI's Intelligence</p>
        <p class="copyright">© 2025 Clause CLI. MIT License.</p>
      </div>
      
      <div class="footer-links">
        <h4>Documentation</h4>
        <ul>
          <li><a href="/docs/getting-started">Getting Started</a></li>
          <li><a href="/docs/installation">Installation</a></li>
          <li><a href="/docs/cli-reference">CLI Reference</a></li>
          <li><a href="/docs/governance">AI Governance</a></li>
        </ul>
      </div>
      
      <div class="footer-links">
        <h4>Community</h4>
        <ul>
          <li><a href="https://github.com/Mr-Dark-debug/clause-cli">GitHub</a></li>
          <li><a href="#">Discord</a></li>
          <li><a href="#">Twitter</a></li>
          <li><a href="/pages/contributing">Contributing</a></li>
        </ul>
      </div>
      
      <div class="footer-links">
        <h4>Resources</h4>
        <ul>
          <li><a href="/pages/changelog">Changelog</a></li>
          <li><a href="/pages/roadmap">Roadmap</a></li>
          <li><a href="/pages/license">License</a></li>
        </ul>
      </div>
    </div>
  </div>
</footer>
```

### Code Block Component

```html
<div class="code-block">
  <div class="code-header">
    <span class="code-language">bash</span>
    <button class="copy-button" data-copy>
      <svg><!-- Copy icon --></svg>
      <span>Copy</span>
    </button>
  </div>
  <pre><code class="language-bash"># Install Clause
curl -fsSL https://Clause.dev/install.sh | bash</code></pre>
</div>
```

### Feature Card Component

```html
<article class="feature-card">
  <div class="feature-icon">
    <svg><!-- Custom icon --></svg>
  </div>
  <h3 class="feature-title">Feature Name</h3>
  <p class="feature-description">
    Brief description of the feature and its benefits.
  </p>
  <a href="/docs/feature" class="feature-link">
    Learn more →
  </a>
</article>
```

### Tabs Component

```html
<div class="tabs">
  <div class="tab-list" role="tablist">
    <button class="tab active" role="tab" data-tab="macos">macOS</button>
    <button class="tab" role="tab" data-tab="linux">Linux</button>
    <button class="tab" role="tab" data-tab="windows">Windows</button>
  </div>
  
  <div class="tab-content">
    <div class="tab-panel active" data-panel="macos">
      <!-- macOS content -->
    </div>
    <div class="tab-panel" data-panel="linux">
      <!-- Linux content -->
    </div>
    <div class="tab-panel" data-panel="windows">
      <!-- Windows content -->
    </div>
  </div>
</div>
```

---

## 🎬 ANIMATIONS & INTERACTIONS

### Scroll Animations

```css
/* Fade in on scroll */
.animate-on-scroll {
  opacity: 0;
  transform: translateY(20px);
  transition: opacity 0.6s ease, transform 0.6s ease;
}

.animate-on-scroll.visible {
  opacity: 1;
  transform: translateY(0);
}

/* Staggered children */
.animate-stagger > * {
  opacity: 0;
  transform: translateY(20px);
}

.animate-stagger.visible > *:nth-child(1) { transition-delay: 0ms; }
.animate-stagger.visible > *:nth-child(2) { transition-delay: 100ms; }
.animate-stagger.visible > *:nth-child(3) { transition-delay: 200ms; }
/* ... etc */
```

### Hover Effects

```css
/* Card hover */
.feature-card {
  transition: transform 0.3s ease, box-shadow 0.3s ease;
}

.feature-card:hover {
  transform: translateY(-4px);
  box-shadow: var(--shadow-lg);
}

/* Button hover */
.btn-primary {
  transition: background-color 0.2s ease, transform 0.1s ease;
}

.btn-primary:hover {
  background-color: var(--color-primary-hover);
}

.btn-primary:active {
  transform: scale(0.98);
}
```

### Terminal Animation

```css
/* Typing animation */
@keyframes typing {
  from { width: 0; }
  to { width: 100%; }
}

@keyframes blink {
  0%, 50% { border-color: var(--color-text-primary); }
  51%, 100% { border-color: transparent; }
}

.terminal-line {
  overflow: hidden;
  white-space: nowrap;
  border-right: 2px solid var(--color-text-primary);
  animation: typing 2s steps(40) forwards, blink 0.8s step-end infinite;
}
```

---

## 📱 RESPONSIVE BREAKPOINTS

```css
/* Mobile first approach */

/* Extra small devices (phones, 320px and up) */
@media (min-width: 320px) {
  /* Base mobile styles */
}

/* Small devices (phones landscape, 576px and up) */
@media (min-width: 576px) {
  /* Larger mobile styles */
}

/* Medium devices (tablets, 768px and up) */
@media (min-width: 768px) {
  /* Tablet styles */
}

/* Large devices (desktops, 992px and up) */
@media (min-width: 992px) {
  /* Desktop styles */
}

/* Extra large devices (large desktops, 1200px and up) */
@media (min-width: 1200px) {
  /* Large desktop styles */
}

/* Extra extra large devices (1400px and up) */
@media (min-width: 1400px) {
  /* Extra large desktop styles */
}
```

---

## ⚡ PERFORMANCE REQUIREMENTS

1. **Critical CSS Inlining** - Inline above-the-fold CSS
2. **Lazy Loading Images** - Use `loading="lazy"` for images below fold
3. **Minified Assets** - Minify CSS and JS in production
4. **Optimized Images** - Use WebP with fallbacks, proper sizing
5. **Preload Key Resources** - Preload fonts and critical images
6. **Defer Non-Critical JS** - Use `defer` attribute
7. **Cache Headers** - Proper cache control for static assets

---

## ✅ ACCESSIBILITY REQUIREMENTS

1. **Semantic HTML** - Use proper HTML5 elements
2. **ARIA Labels** - Label interactive elements
3. **Keyboard Navigation** - All interactive elements focusable
4. **Focus Indicators** - Visible focus states
5. **Color Contrast** - WCAG AA minimum (4.5:1 for text)
6. **Alt Text** - All images have descriptive alt text
7. **Skip Links** - Skip to main content link
8. **Form Labels** - All form inputs have labels
9. **Reduced Motion** - Respect `prefers-reduced-motion`

---

## 🔧 TECHNICAL IMPLEMENTATION

### No Build Tools Required

This website should work without any build process:
- Plain HTML files
- Vanilla CSS (with CSS custom properties)
- Vanilla JavaScript (ES6+)
- No framework dependencies

### Optional Enhancements (Future)

- Service Worker for offline support
- Search functionality (client-side)
- Versioned documentation
- Internationalization (i18n)

---

## 📝 CONTENT TO WRITE

The following content must be created:

1. **All page copy** - Professional, clear, developer-focused
2. **Code examples** - Working, tested commands
3. **Image assets** - Logo, icons, illustrations
4. **Meta descriptions** - For each page
5. **Open Graph tags** - For social sharing

---

## 🚀 EXECUTION ORDER

Build in this order:

1. **CSS Variables & Base Styles** (`variables.css`, `main.css`)
2. **Typography & Utilities** (`utilities.css`)
3. **Components** (`components.css`)
4. **Layouts** (`layouts.css`)
5. **Header & Footer** (create components, add to all pages)
6. **Homepage** (`index.html`)
7. **Documentation Home** (`docs/index.html`)
8. **Installation Page** (`docs/installation.html`)
9. **Getting Started** (`docs/getting-started.html`)
10. **Governance Page** (`docs/governance.html`)
11. **CLI Reference** (`docs/cli-reference.html`)
12. **Changelog** (`pages/changelog.html`)
13. **Roadmap** (`pages/roadmap.html`)
14. **Contributing** (`pages/contributing.html`)
15. **Community** (`pages/community.html`)
16. **JavaScript Functionality** (theme toggle, navigation, etc.)
17. **Animations** (`animations.css`, `animations.js`)
18. **Final Polish** (responsive testing, accessibility audit)

---

## 📋 QUALITY CHECKLIST

Before considering the website complete:

- [ ] All pages render correctly on mobile, tablet, desktop
- [ ] Dark/light theme toggle works and persists
- [ ] All internal links work correctly
- [ ] All external links open in new tab
- [ ] Copy buttons on code blocks work
- [ ] Tab components switch content
- [ ] Mobile navigation works
- [ ] Images load with proper fallbacks
- [ ] No console errors
- [ ] Lighthouse score > 90 for all categories
- [ ] All pages have proper meta tags
- [ ] Accessibility audit passes
- [ ] All forms (if any) work correctly
- [ ] Search functionality works (if implemented)

---

## 🎯 FINAL DELIVERABLE

A complete, production-ready website that:
- ✅ Looks professional and modern
- ✅ Works on all devices and browsers
- ✅ Loads fast and performs well
- ✅ Is accessible to all users
- ✅ Has comprehensive documentation
- ✅ Represents Clause CLI professionally
- ✅ Encourages adoption and contribution

---

**Now, build this website following these specifications. Create beautiful, responsive, accessible HTML/CSS/JS. Make something the open source community will be proud of.** 🚀
