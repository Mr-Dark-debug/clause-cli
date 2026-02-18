package generator

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

// createFrontend creates the frontend structure.
func (g *Generator) createFrontend(projectPath string) error {
	frontendDir := filepath.Join(projectPath, g.Config.Frontend.Directory)

	// Create frontend directory
	if err := g.createDirectory(frontendDir); err != nil {
		return err
	}

	// Create src directory
	srcDir := filepath.Join(frontendDir, "src")
	if err := g.createDirectory(srcDir); err != nil {
		return err
	}

	// Create public directory
	publicDir := filepath.Join(frontendDir, "public")
	if err := g.createDirectory(publicDir); err != nil {
		return err
	}

	// Create package.json
	packageJSON := g.generatePackageJSON()
	if err := g.writeFile(filepath.Join(frontendDir, "package.json"), packageJSON); err != nil {
		return err
	}

	// Create tsconfig.json if TypeScript
	if g.Config.Frontend.TypeScript {
		tsconfig := g.generateTSConfig()
		if err := g.writeFile(filepath.Join(frontendDir, "tsconfig.json"), tsconfig); err != nil {
			return err
		}
	}

	// Create main entry file
	mainContent := g.generateFrontendMain()
	mainFile := "index.tsx"
	if !g.Config.Frontend.TypeScript {
		mainFile = "index.js"
	}
	if err := g.writeFile(filepath.Join(srcDir, mainFile), mainContent); err != nil {
		return err
	}

	// Create App component
	appContent := g.generateAppComponent()
	appFile := "App.tsx"
	if !g.Config.Frontend.TypeScript {
		appFile = "App.js"
	}
	if err := g.writeFile(filepath.Join(srcDir, appFile), appContent); err != nil {
		return err
	}

	return nil
}

// createBackend creates the backend structure.
func (g *Generator) createBackend(projectPath string) error {
	backendDir := filepath.Join(projectPath, g.Config.Backend.Directory)

	// Create backend directory
	if err := g.createDirectory(backendDir); err != nil {
		return err
	}

	// Create structure based on language/framework
	switch g.Config.Backend.Language {
	case "python":
		return g.createPythonBackend(backendDir)
	case "node", "typescript":
		return g.createNodeBackend(backendDir)
	case "go":
		return g.createGoBackend(backendDir)
	default:
		return g.createGenericBackend(backendDir)
	}
}

// createPythonBackend creates Python backend structure.
func (g *Generator) createPythonBackend(backendDir string) error {
	// Create main.py
	mainContent := `"""
Main entry point for the application.
"""

def main():
    print("Hello from {{.Project.Name}}!")

if __name__ == "__main__":
    main()
`
	if err := g.writeTemplate(filepath.Join(backendDir, "main.py"), mainContent); err != nil {
		return err
	}

	// Create requirements.txt
	requirements := `fastapi>=0.100.0
uvicorn>=0.22.0
pydantic>=2.0.0
`
	if err := g.writeFile(filepath.Join(backendDir, "requirements.txt"), requirements); err != nil {
		return err
	}

	// Create pyproject.toml
	pyproject := g.generatePyproject()
	if err := g.writeFile(filepath.Join(backendDir, "pyproject.toml"), pyproject); err != nil {
		return err
	}

	// Create app directory
	appDir := filepath.Join(backendDir, "app")
	if err := g.createDirectory(appDir); err != nil {
		return err
	}

	// Create __init__.py
	if err := g.writeFile(filepath.Join(appDir, "__init__.py"), ""); err != nil {
		return err
	}

	return nil
}

// createNodeBackend creates Node.js backend structure.
func (g *Generator) createNodeBackend(backendDir string) error {
	// Create package.json
	packageJSON := g.generateBackendPackageJSON()
	if err := g.writeFile(filepath.Join(backendDir, "package.json"), packageJSON); err != nil {
		return err
	}

	// Create src directory
	srcDir := filepath.Join(backendDir, "src")
	if err := g.createDirectory(srcDir); err != nil {
		return err
	}

	// Create index.js
	indexContent := "const express = require('express');\n\n" +
		"const app = express();\n" +
		"const port = process.env.PORT || 3000;\n\n" +
		"app.get('/', (req, res) => {\n" +
		"  res.json({ message: 'Hello from {{.Project.Name}}!' });\n" +
		"});\n\n" +
		"app.listen(port, () => {\n" +
		"  console.log(`Server running on port ${port}`);\n" +
		"});\n"
	if err := g.writeTemplate(filepath.Join(srcDir, "index.js"), indexContent); err != nil {
		return err
	}

	return nil
}

// createGoBackend creates Go backend structure.
func (g *Generator) createGoBackend(backendDir string) error {
	// Create main.go
	mainContent := `package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from {{.Project.Name}}!")
	})

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}
`
	if err := g.writeTemplate(filepath.Join(backendDir, "main.go"), mainContent); err != nil {
		return err
	}

	// Create go.mod
	goMod := fmt.Sprintf(`module %s

go 1.21
`, g.Config.Metadata.Name)
	if err := g.writeFile(filepath.Join(backendDir, "go.mod"), goMod); err != nil {
		return err
	}

	return nil
}

// createGenericBackend creates a generic backend structure.
func (g *Generator) createGenericBackend(backendDir string) error {
	// Create a basic README
	readme := fmt.Sprintf("# %s Backend\n\nBackend implementation for %s.\n",
		g.Config.Metadata.Name, g.Config.Metadata.Name)
	return g.writeFile(filepath.Join(backendDir, "README.md"), readme)
}

// createInfrastructure creates infrastructure files.
func (g *Generator) createInfrastructure(projectPath string) error {
	// Create Dockerfile if enabled
	if g.Config.Infrastructure.Docker {
		dockerfile := g.generateDockerfile()
		if err := g.writeFile(filepath.Join(projectPath, "Dockerfile"), dockerfile); err != nil {
			return err
		}
	}

	// Create docker-compose.yml if enabled
	if g.Config.Infrastructure.DockerCompose {
		dockerCompose := g.generateDockerCompose()
		if err := g.writeFile(filepath.Join(projectPath, "docker-compose.yml"), dockerCompose); err != nil {
			return err
		}
	}

	// Create CI configuration
	if g.Config.Infrastructure.CI != "" {
		if err := g.createCIConfig(projectPath); err != nil {
			return err
		}
	}

	return nil
}

// createCIConfig creates CI configuration files.
func (g *Generator) createCIConfig(projectPath string) error {
	switch g.Config.Infrastructure.CI {
	case "github-actions":
		// Create .github/workflows directory
		workflowsDir := filepath.Join(projectPath, ".github", "workflows")
		if err := g.createDirectory(workflowsDir); err != nil {
			return err
		}

		// Create main workflow
		workflow := g.generateGitHubActionsWorkflow()
		if err := g.writeFile(filepath.Join(workflowsDir, "main.yml"), workflow); err != nil {
			return err
		}
	}

	return nil
}

// createGovernance creates governance files.
func (g *Generator) createGovernance(projectPath string) error {
	// Create ai_prompt_guidelines directory
	guidelinesDir := filepath.Join(projectPath, "ai_prompt_guidelines")
	if err := g.createDirectory(guidelinesDir); err != nil {
		return err
	}

	// Create Brainstorm.md if enabled
	if g.Config.Governance.BrainstormMd {
		brainstorm := g.generateBrainstormMd()
		if err := g.writeFile(filepath.Join(guidelinesDir, "brainstorm.md"), brainstorm); err != nil {
			return err
		}
	}

	// Create system_prompt.md if enabled (via PromptGuidelines flag)
	if g.Config.Governance.PromptGuidelines {
		systemPrompt := g.generateSystemPromptMd()
		if err := g.writeFile(filepath.Join(guidelinesDir, "system_prompt.md"), systemPrompt); err != nil {
			return err
		}
	}

	// Create architecture.md (standard governance file)
	architecture := g.generateArchitectureMd()
	if err := g.writeFile(filepath.Join(guidelinesDir, "architecture.md"), architecture); err != nil {
		return err
	}

	// Create component_registry.json if enabled
	if g.Config.Governance.ComponentRegistry {
		registry := g.generateComponentRegistryJson()
		if err := g.writeFile(filepath.Join(guidelinesDir, "component_registry.json"), registry); err != nil {
			return err
		}
	}

	return nil
}

// initGit initializes a git repository.
func (g *Generator) initGit(projectPath string) error {
	if g.DryRun {
		g.Logger.Info("[DRY RUN] Would initialize git repository")
		return nil
	}

	cmd := exec.Command("git", "init")
	cmd.Dir = projectPath
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to initialize git: %w", err)
	}

	// Initial commit
	cmd = exec.Command("git", "add", ".")
	cmd.Dir = projectPath
	cmd.Run()

	cmd = exec.Command("git", "commit", "-m", "Initial commit")
	cmd.Dir = projectPath
	cmd.Run()

	return nil
}

// Helper functions for content generation

func (g *Generator) generatePackageJSON() string {
	return fmt.Sprintf(`{
  "name": "%s",
  "version": "1.0.0",
  "description": "%s",
  "scripts": {
    "dev": "vite",
    "build": "vite build",
    "preview": "vite preview"
  },
  "dependencies": {
    "react": "^18.2.0",
    "react-dom": "^18.2.0"
  },
  "devDependencies": {
    "@types/react": "^18.2.0",
    "@vitejs/plugin-react": "^4.0.0",
    "typescript": "^5.0.0",
    "vite": "^4.4.0"
  }
}
`, g.Config.Metadata.Name, g.Config.Metadata.Description)
}

func (g *Generator) generateBackendPackageJSON() string {
	return fmt.Sprintf(`{
  "name": "%s-backend",
  "version": "1.0.0",
  "description": "%s",
  "main": "src/index.js",
  "scripts": {
    "start": "node src/index.js",
    "dev": "nodemon src/index.js"
  },
  "dependencies": {
    "express": "^4.18.0"
  },
  "devDependencies": {
    "nodemon": "^3.0.0"
  }
}
`, g.Config.Metadata.Name, g.Config.Metadata.Description)
}

func (g *Generator) generateTSConfig() string {
	return `{
  "compilerOptions": {
    "target": "ES2020",
    "useDefineForClassFields": true,
    "lib": ["ES2020", "DOM", "DOM.Iterable"],
    "module": "ESNext",
    "skipLibCheck": true,
    "moduleResolution": "bundler",
    "allowImportingTsExtensions": true,
    "resolveJsonModule": true,
    "isolatedModules": true,
    "noEmit": true,
    "jsx": "react-jsx",
    "strict": true,
    "noUnusedLocals": true,
    "noUnusedParameters": true,
    "noFallthroughCasesInSwitch": true
  },
  "include": ["src"],
  "references": [{ "path": "./tsconfig.node.json" }]
}
`
}

func (g *Generator) generatePyproject() string {
	return fmt.Sprintf(`[project]
name = "%s"
version = "0.1.0"
description = "%s"
requires-python = ">=3.11"

[build-system]
requires = ["setuptools>=61.0"]
build-backend = "setuptools.build_meta"
`, g.Config.Metadata.Name, g.Config.Metadata.Description)
}

func (g *Generator) generateFrontendMain() string {
	return `import React from 'react'
import ReactDOM from 'react-dom/client'
import App from './App'

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
)
`
}

func (g *Generator) generateAppComponent() string {
	return `import React from 'react'

function App() {
  return (
    <div>
      <h1>Welcome to {{.Project.Name}}</h1>
      <p>{{.Project.Description}}</p>
    </div>
  )
}

export default App
`
}

func (g *Generator) generateDockerfile() string {
	return `# Build stage
FROM node:18-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .
RUN npm run build

# Production stage
FROM nginx:alpine

COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
`
}

func (g *Generator) generateDockerCompose() string {
	return `version: '3.8'

services:
  frontend:
    build:
      context: .
      dockerfile: Dockerfile
    ports:
      - "3000:80"
    environment:
      - NODE_ENV=production

  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    ports:
      - "8000:8000"
    environment:
      - DATABASE_URL=postgresql://postgres:postgres@db:5432/app
    depends_on:
      - db

  db:
    image: postgres:15-alpine
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB=app
    volumes:
      - postgres_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

volumes:
  postgres_data:
`
}

func (g *Generator) generateGitHubActionsWorkflow() string {
	return `name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: Setup Node.js
        uses: actions/setup-node@v4
        with:
          node-version: '18'
          cache: 'npm'

      - name: Install dependencies
        run: npm ci

      - name: Run tests
        run: npm test

      - name: Build
        run: npm run build
`
}

func (g *Generator) generateBrainstormMd() string {
	return fmt.Sprintf(`# Brainstorm

Welcome to the brainstorming workspace for **%s**.

## Ideas

*Use this space to explore ideas and concepts for your project.*

-

## Questions

*What questions do you need to answer?*

-

## Research

*Links and notes from research*

-

## Decisions

*Key decisions and their rationale*

-
`, g.Config.Metadata.Name)
}

func (g *Generator) generateSystemPromptMd() string {
	return fmt.Sprintf(`# System Prompt

You are an AI assistant working on the **%s** project.

## Project Context
%s

## Technology Stack
- Frontend: %s
- Backend: %s
- Database: %s

## Coding Standards
1. Follow the component structure defined in component_registry.json
2. Implement strict type checking
3. Write comprehensive tests
4. Use the architecture defined in architecture.md
`, g.Config.Metadata.Name, g.Config.Metadata.Description, g.Config.Frontend.Framework, g.Config.Backend.Framework, g.Config.Backend.Database.Primary)
}

func (g *Generator) generateArchitectureMd() string {
	return fmt.Sprintf(`# Project Architecture

## Overview
%s

## Frontend
- Framework: %s
- Application Structure: %s

## Backend
- Framework: %s
- Database: %s
- API Style: %s

## Infrastructure
- Deployment: %s
- CI/CD: %s
`, g.Config.Metadata.Description, g.Config.Frontend.Framework, g.Config.Frontend.Directory, g.Config.Backend.Framework, g.Config.Backend.Database.Primary, g.Config.Backend.API.Style, g.Config.Infrastructure.Hosting, g.Config.Infrastructure.CI)
}

func (g *Generator) generateComponentRegistryJson() string {
	return `{
  "components": [],
  "layouts": [],
  "pages": []
}
`
}

func (g *Generator) generatePromptGuidelines() string {
	return `# AI Prompt Guidelines

This document provides guidelines for working with AI assistants on this project.

## Project Context

- **Name**: {{.Project.Name}}
- **Description**: {{.Project.Description}}

## Technology Stack

- **Frontend**: {{.Frontend.Framework}}
- **Backend**: {{.Backend.Framework}}
- **Database**: {{.Backend.Database.Primary}}

## Code Style

- Follow the existing code patterns in the project
- Keep functions small and focused
- Write clear, descriptive variable names
- Add comments for complex logic

## Best Practices

1. Always validate inputs
2. Handle errors gracefully
3. Write tests for new features
4. Update documentation when needed

## When asking for help

1. Provide context about what you're trying to achieve
2. Share relevant code snippets
3. Describe any errors you're seeing
4. Explain what you've already tried
`
}
