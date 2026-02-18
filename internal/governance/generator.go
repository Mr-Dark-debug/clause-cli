package governance

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/output"
)

// Generator generates governance files for a project.
type Generator struct {
	// ProjectPath is the root path of the project
	ProjectPath string

	// Config is the project configuration
	Config *config.ProjectConfig

	// Logger for output
	Logger *output.Logger
}

// NewGenerator creates a new governance generator.
func NewGenerator(projectPath string, cfg *config.ProjectConfig) *Generator {
	return &Generator{
		ProjectPath: projectPath,
		Config:      cfg,
		Logger:      output.DefaultLogger,
	}
}

// Generate generates all governance files.
func (g *Generator) Generate() error {
	// Create .clause directory
	clauseDir := filepath.Join(g.ProjectPath, ".clause")
	if err := os.MkdirAll(clauseDir, 0755); err != nil {
		return fmt.Errorf("failed to create .clause directory: %w", err)
	}

	// Generate context.yaml
	if err := g.generateContextFile(clauseDir); err != nil {
		return err
	}

	// Generate prompt-guidelines.md if enabled
	if g.Config.Governance.PromptGuidelines {
		if err := g.generatePromptGuidelines(clauseDir); err != nil {
			return err
		}
	}

	// Generate component-registry.yaml if enabled
	if g.Config.Governance.ComponentRegistry {
		if err := g.generateComponentRegistry(clauseDir); err != nil {
			return err
		}
	}

	// Generate Brainstorm.md in project root if enabled
	if g.Config.Governance.BrainstormMd {
		if err := g.generateBrainstormMd(); err != nil {
			return err
		}
	}

	return nil
}

// generateContextFile generates the context.yaml file.
func (g *Generator) generateContextFile(clauseDir string) error {
	contextFile := filepath.Join(clauseDir, "context.yaml")

	// Build context content
	var content strings.Builder

	content.WriteString("# AI Context\n")
	content.WriteString("# This file provides context for AI assistants working on this project.\n\n")

	// Project metadata
	content.WriteString("project_name: \"")
	content.WriteString(g.Config.Metadata.Name)
	content.WriteString("\"\n")

	if g.Config.Metadata.Description != "" {
		content.WriteString("description: \"")
		content.WriteString(g.Config.Metadata.Description)
		content.WriteString("\"\n")
	}

	// Tech Stack
	content.WriteString("\ntech_stack:\n")

	// Frontend tech
	if g.Config.Frontend.Enabled {
		content.WriteString(fmt.Sprintf("  - %s (%s)\n", g.Config.Frontend.Framework, "frontend"))

		if g.Config.Frontend.TypeScript {
			content.WriteString("  - TypeScript\n")
		}

		if g.Config.Frontend.Styling != "" {
			content.WriteString(fmt.Sprintf("  - %s (styling)\n", g.Config.Frontend.Styling))
		}

		if g.Config.Frontend.TestFramework != "" {
			content.WriteString(fmt.Sprintf("  - %s (testing)\n", g.Config.Frontend.TestFramework))
		}
	}

	// Backend tech
	if g.Config.Backend.Enabled {
		lang := g.Config.Backend.Language
		if lang == "" {
			lang = "unknown"
		}
		content.WriteString(fmt.Sprintf("  - %s (backend)\n", g.Config.Backend.Framework))
		content.WriteString(fmt.Sprintf("  - %s\n", strings.Title(lang)))

		if g.Config.Backend.Database.Primary != "" {
			content.WriteString(fmt.Sprintf("  - %s (database)\n", g.Config.Backend.Database.Primary))
		}

		if g.Config.Backend.Database.ORM != "" {
			content.WriteString(fmt.Sprintf("  - %s (ORM)\n", g.Config.Backend.Database.ORM))
		}
	}

	// Infrastructure tech
	if g.Config.Infrastructure.Docker {
		content.WriteString("  - Docker\n")
	}
	if g.Config.Infrastructure.DockerCompose {
		content.WriteString("  - Docker Compose\n")
	}
	if g.Config.Infrastructure.Kubernetes {
		content.WriteString("  - Kubernetes\n")
	}
	if g.Config.Infrastructure.CI != "" {
		content.WriteString(fmt.Sprintf("  - %s (CI/CD)\n", g.Config.Infrastructure.CI))
	}

	// Architecture
	content.WriteString("\narchitecture:\n")
	content.WriteString("  style: \"")
	if g.Config.Frontend.Enabled && g.Config.Backend.Enabled {
		content.WriteString("full-stack")
	} else if g.Config.Frontend.Enabled {
		content.WriteString("frontend")
	} else if g.Config.Backend.Enabled {
		content.WriteString("backend")
	}
	content.WriteString("\"\n")

	if g.Config.Frontend.Enabled {
		content.WriteString(fmt.Sprintf("  frontend: \"%s\"\n", g.Config.Frontend.Framework))
	}
	if g.Config.Backend.Enabled {
		content.WriteString(fmt.Sprintf("  backend: \"%s (%s)\"\n", g.Config.Backend.Framework, g.Config.Backend.Language))
	}
	if g.Config.Backend.Database.Primary != "" {
		content.WriteString(fmt.Sprintf("  database: \"%s\"\n", g.Config.Backend.Database.Primary))
	}

	// Patterns
	content.WriteString("\npatterns:\n")
	if g.Config.Frontend.Enabled {
		content.WriteString("  - component-based architecture\n")
	}
	if g.Config.Backend.Enabled {
		content.WriteString("  - RESTful API\n")
		if g.Config.Backend.API.Style == "graphql" {
			content.WriteString("  - GraphQL\n")
		}
	}

	// Best practices
	content.WriteString("\nbest_practices:\n")
	content.WriteString("  - Write clear, self-documenting code\n")
	content.WriteString("  - Follow the existing code patterns in the project\n")
	content.WriteString("  - Keep functions small and focused\n")
	content.WriteString("  - Handle errors gracefully\n")
	content.WriteString("  - Write tests for new features\n")

	// Key files
	content.WriteString("\nkey_files:\n")
	content.WriteString("  - path: \".clause/config.yaml\"\n")
	content.WriteString("    purpose: \"Project configuration\"\n")
	content.WriteString("  - path: \".clause/context.yaml\"\n")
	content.WriteString("    purpose: \"AI context (this file)\"\n")

	// Components placeholder
	content.WriteString("\ncomponents: []\n")

	// Conventions placeholder
	content.WriteString("\nconventions: []\n")

	return os.WriteFile(contextFile, []byte(content.String()), 0644)
}

// generatePromptGuidelines generates the prompt-guidelines.md file.
func (g *Generator) generatePromptGuidelines(clauseDir string) error {
	guidelinesFile := filepath.Join(clauseDir, "prompt-guidelines.md")

	var content strings.Builder

	content.WriteString("# AI Prompt Guidelines\n\n")
	content.WriteString("This document provides guidelines for working with AI assistants on this project.\n\n")

	// Project Context
	content.WriteString("## Project Context\n\n")

	content.WriteString(fmt.Sprintf("- **Name**: %s\n", g.Config.Metadata.Name))
	if g.Config.Metadata.Description != "" {
		content.WriteString(fmt.Sprintf("- **Description**: %s\n", g.Config.Metadata.Description))
	}
	if g.Config.Metadata.Author != "" {
		content.WriteString(fmt.Sprintf("- **Author**: %s\n", g.Config.Metadata.Author))
	}
	if g.Config.Metadata.License != "" {
		content.WriteString(fmt.Sprintf("- **License**: %s\n", g.Config.Metadata.License))
	}

	content.WriteString("\n")

	// Technology Stack
	content.WriteString("## Technology Stack\n\n")

	if g.Config.Frontend.Enabled {
		content.WriteString("### Frontend\n\n")
		content.WriteString(fmt.Sprintf("- **Framework**: %s\n", g.Config.Frontend.Framework))
		if g.Config.Frontend.TypeScript {
			content.WriteString("- **Language**: TypeScript\n")
		}
		if g.Config.Frontend.Styling != "" {
			content.WriteString(fmt.Sprintf("- **Styling**: %s\n", g.Config.Frontend.Styling))
		}
		if g.Config.Frontend.TestFramework != "" {
			content.WriteString(fmt.Sprintf("- **Testing**: %s\n", g.Config.Frontend.TestFramework))
		}
		if g.Config.Frontend.PackageManager != "" {
			content.WriteString(fmt.Sprintf("- **Package Manager**: %s\n", g.Config.Frontend.PackageManager))
		}
		content.WriteString("\n")
	}

	if g.Config.Backend.Enabled {
		content.WriteString("### Backend\n\n")
		content.WriteString(fmt.Sprintf("- **Framework**: %s\n", g.Config.Backend.Framework))
		if g.Config.Backend.Language != "" {
			content.WriteString(fmt.Sprintf("- **Language**: %s\n", g.Config.Backend.Language))
		}
		if g.Config.Backend.Database.Primary != "" {
			content.WriteString(fmt.Sprintf("- **Database**: %s\n", g.Config.Backend.Database.Primary))
		}
		if g.Config.Backend.Database.ORM != "" {
			content.WriteString(fmt.Sprintf("- **ORM**: %s\n", g.Config.Backend.Database.ORM))
		}
		content.WriteString("\n")
	}

	// Code Style
	content.WriteString("## Code Style\n\n")
	content.WriteString("- Follow the existing code patterns in the project\n")
	content.WriteString("- Keep functions small and focused\n")
	content.WriteString("- Write clear, descriptive variable names\n")
	content.WriteString("- Add comments for complex logic\n")
	content.WriteString("- Use consistent formatting\n")
	content.WriteString("\n")

	// Best Practices
	content.WriteString("## Best Practices\n\n")
	content.WriteString("1. Always validate inputs at system boundaries\n")
	content.WriteString("2. Handle errors gracefully with meaningful messages\n")
	content.WriteString("3. Write tests for new features\n")
	content.WriteString("4. Update documentation when needed\n")
	content.WriteString("5. Follow the principle of least surprise\n")
	content.WriteString("\n")

	// When asking for help
	content.WriteString("## When asking for help\n\n")
	content.WriteString("1. Provide context about what you're trying to achieve\n")
	content.WriteString("2. Share relevant code snippets\n")
	content.WriteString("3. Describe any errors you're seeing\n")
	content.WriteString("4. Explain what you've already tried\n")
	content.WriteString("\n")

	// Component Guidelines
	if g.Config.Governance.ComponentRegistry {
		content.WriteString("## Component Guidelines\n\n")
		content.WriteString("- Register new components in `.clause/registry.yaml`\n")
		content.WriteString("- Document component dependencies\n")
		content.WriteString("- Keep component interfaces stable\n")
		content.WriteString("\n")
	}

	return os.WriteFile(guidelinesFile, []byte(content.String()), 0644)
}

// generateComponentRegistry generates the component registry file.
func (g *Generator) generateComponentRegistry(clauseDir string) error {
	registryFile := filepath.Join(clauseDir, "registry.yaml")

	var content strings.Builder

	content.WriteString("# Component Registry\n")
	content.WriteString("# This file tracks all components in the project.\n\n")

	content.WriteString("components: []\n")
	content.WriteString("\n")
	content.WriteString("# Example component:\n")
	content.WriteString("# - name: \"user-auth\"\n")
	content.WriteString("#   type: \"service\"\n")
	content.WriteString("#   path: \"backend/services/auth\"\n")
	content.WriteString("#   description: \"Authentication service\"\n")
	content.WriteString("#   dependencies: []\n")
	content.WriteString("#   tags: [\"auth\", \"security\"]\n")
	content.WriteString("#   tech_stack: [\"go\", \"jwt\"]\n")

	return os.WriteFile(registryFile, []byte(content.String()), 0644)
}

// generateBrainstormMd generates the Brainstorm.md file in project root.
func (g *Generator) generateBrainstormMd() error {
	brainstormFile := filepath.Join(g.ProjectPath, "Brainstorm.md")

	var content strings.Builder

	content.WriteString(fmt.Sprintf("# Brainstorm: %s\n\n", g.Config.Metadata.Name))

	if g.Config.Metadata.Description != "" {
		content.WriteString(fmt.Sprintf("%s\n\n", g.Config.Metadata.Description))
	}

	content.WriteString("Welcome to the brainstorming workspace for this project.\n\n")

	// Ideas section
	content.WriteString("## Ideas\n\n")
	content.WriteString("*Use this space to explore ideas and concepts.*\n\n")
	content.WriteString("- \n\n")

	// Questions section
	content.WriteString("## Questions\n\n")
	content.WriteString("*What questions do you need to answer?*\n\n")
	content.WriteString("- \n\n")

	// Research section
	content.WriteString("## Research\n\n")
	content.WriteString("*Links and notes from research.*\n\n")
	content.WriteString("- \n\n")

	// Decisions section
	content.WriteString("## Decisions\n\n")
	content.WriteString("*Key decisions and their rationale.*\n\n")
	content.WriteString("- \n\n")

	// Tasks section
	content.WriteString("## Tasks\n\n")
	content.WriteString("*Things to do.*\n\n")
	content.WriteString("- [ ] \n\n")

	// Notes section
	content.WriteString("## Notes\n\n")
	content.WriteString("*Free-form notes and thoughts.*\n\n")

	return os.WriteFile(brainstormFile, []byte(content.String()), 0644)
}
