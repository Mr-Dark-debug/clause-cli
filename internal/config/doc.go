// Package config provides comprehensive configuration management for Clause CLI.
//
// This package handles all aspects of configuration including:
//   - Configuration data structures for projects
//   - Loading configuration from multiple sources
//   - Saving and persisting configuration
//   - Validating configuration values
//   - Applying defaults and presets
//
// # Configuration Structure
//
// The main configuration type is ProjectConfig, which contains:
//   - Metadata: Project identification (name, version, author, etc.)
//   - Frontend: Frontend framework and tooling configuration
//   - Backend: Backend framework and service configuration
//   - Infrastructure: Deployment and infrastructure settings
//   - Governance: AI governance and compliance settings
//   - Development: Development workflow settings
//
// # Loading Configuration
//
// Configuration is loaded with priority-based merging:
//
//	1. Explicit flags/options (highest priority)
//	2. Environment variables (CLAUSE_*)
//	3. Project configuration (.clause/config.yaml)
//	4. Global configuration (~/.clause/config.yaml)
//	5. Default values (lowest priority)
//
// Example usage:
//
//	loader := config.NewLoader(
//	    config.WithProjectDir("/path/to/project"),
//	)
//	cfg, err := loader.Load()
//	if err != nil {
//	    log.Fatal(err)
//	}
//
// # Saving Configuration
//
// Configuration can be saved to files with automatic backup support:
//
//	saver := config.NewSaver(
//	    config.WithFormat("yaml"),
//	    config.WithBackup(true),
//	)
//	if err := saver.SaveToProject(cfg, "/path/to/project"); err != nil {
//	    log.Fatal(err)
//	}
//
// # Validation
//
// Configuration can be validated to ensure correctness:
//
//	validator := config.NewValidator()
//	errors := validator.Validate(cfg)
//	if errors.HasErrors() {
//	    for _, err := range errors {
//	        fmt.Printf("%s: %s\n", err.Field, err.Message)
//	    }
//	}
//
// # Presets
//
// Presets provide pre-configured setups for common use cases:
//
//	- minimal: Basic configuration with essentials
//	- standard: Full-stack configuration (default)
//	- saas: SaaS application with auth, payments, multi-tenancy
//	- api-only: Backend-only API service
//	- frontend-only: Frontend-only static site
//	- enterprise: Full enterprise configuration with strict governance
//
// Example with preset:
//
//	preset, _ := config.GetPreset("saas")
//	cfg := config.NewProjectConfig()
//	preset.Apply(cfg)
//
// # Environment Variables
//
// Configuration can be overridden via environment variables with the CLAUSE_ prefix:
//
//	CLAUSE_FRONTEND_FRAMEWORK=nextjs
//	CLAUSE_BACKEND_FRAMEWORK=fastapi
//	CLAUSE_BACKEND_DATABASE=postgresql
//	CLAUSE_GOVERNANCE_ENABLED=true
//
// # File Format
//
// Configuration files use YAML format by default:
//
//	metadata:
//	  name: my-project
//	  version: "0.1.0"
//	frontend:
//	  enabled: true
//	  framework: react
//	  typescript: true
//	  styling: tailwind
//	backend:
//	  enabled: true
//	  framework: fastapi
//	  database:
//	    primary: postgresql
//	    orm: sqlalchemy
package config
