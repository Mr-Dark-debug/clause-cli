// Package config provides configuration management for Clause CLI.
// It supports multi-source configuration loading with priority-based merging
// and provides validation for all configuration options.
package config

import (
	"fmt"
	"time"
)

// ProjectConfig represents the complete configuration for a Clause project.
// It contains all settings for frontend, backend, infrastructure, and governance.
type ProjectConfig struct {
	// Metadata contains project identification information
	Metadata ProjectMetadata `yaml:"metadata" json:"metadata"`

	// Frontend contains frontend framework and tooling configuration
	Frontend FrontendConfig `yaml:"frontend" json:"frontend"`

	// Backend contains backend framework and service configuration
	Backend BackendConfig `yaml:"backend" json:"backend"`

	// Infrastructure contains deployment and infrastructure configuration
	Infrastructure InfrastructureConfig `yaml:"infrastructure" json:"infrastructure"`

	// Governance contains AI governance and compliance settings
	Governance GovernanceConfig `yaml:"governance" json:"governance"`

	// Development contains development workflow settings
	Development DevelopmentConfig `yaml:"development" json:"development"`

	// Version is the configuration schema version
	Version string `yaml:"version" json:"version"`
}

// ProjectMetadata contains basic project identification information.
type ProjectMetadata struct {
	// Name is the project name
	Name string `yaml:"name" json:"name"`

	// Description is a brief project description
	Description string `yaml:"description,omitempty" json:"description,omitempty"`

	// Version is the current project version
	Version string `yaml:"version" json:"version"`

	// Author is the project author or team
	Author string `yaml:"author,omitempty" json:"author,omitempty"`

	// License is the project license
	License string `yaml:"license,omitempty" json:"license,omitempty"`

	// Repository is the git repository URL
	Repository string `yaml:"repository,omitempty" json:"repository,omitempty"`

	// Keywords are searchable project keywords
	Keywords []string `yaml:"keywords,omitempty" json:"keywords,omitempty"`

	// CreatedAt is when the project was created
	CreatedAt time.Time `yaml:"created_at" json:"created_at"`

	// UpdatedAt is when the configuration was last modified
	UpdatedAt time.Time `yaml:"updated_at" json:"updated_at"`

	// ClauseVersion is the version of Clause used to create the project
	ClauseVersion string `yaml:"clause_version" json:"clause_version"`
}

// FrontendConfig contains frontend framework and tooling configuration.
type FrontendConfig struct {
	// Enabled indicates if the project has a frontend
	Enabled bool `yaml:"enabled" json:"enabled"`

	// Framework is the frontend framework (react, vue, svelte, angular, nextjs)
	Framework string `yaml:"framework" json:"framework"`

	// FrameworkVersion is the framework version
	FrameworkVersion string `yaml:"framework_version,omitempty" json:"framework_version,omitempty"`

	// TypeScript indicates if TypeScript is used
	TypeScript bool `yaml:"typescript" json:"typescript"`

	// Styling is the styling approach (tailwind, css-modules, styled-components, scss)
	Styling string `yaml:"styling" json:"styling"`

	// PackageManager is the package manager (npm, yarn, pnpm, bun)
	PackageManager string `yaml:"package_manager" json:"package_manager"`

	// BuildTool is the build tool (vite, webpack, esbuild, rollup)
	BuildTool string `yaml:"build_tool" json:"build_tool"`

	// TestFramework is the testing framework (jest, vitest, playwright, cypress)
	TestFramework string `yaml:"test_framework,omitempty" json:"test_framework,omitempty"`

	// Linter is the linting tool (eslint, biome)
	Linter string `yaml:"linter,omitempty" json:"linter,omitempty"`

	// Formatter is the code formatter (prettier, biome)
	Formatter string `yaml:"formatter,omitempty" json:"formatter,omitempty"`

	// Features contains optional frontend features
	Features FrontendFeatures `yaml:"features" json:"features"`

	// Directory is the frontend source directory
	Directory string `yaml:"directory" json:"directory"`
}

// FrontendFeatures contains optional frontend feature flags.
type FrontendFeatures struct {
	// SSR enables server-side rendering
	SSR bool `yaml:"ssr" json:"ssr"`

	// SSG enables static site generation
	SSG bool `yaml:"ssg" json:"ssg"`

	// PWA enables progressive web app features
	PWA bool `yaml:"pwa" json:"pwa"`

	// I18n enables internationalization
	I18n bool `yaml:"i18n" json:"i18n"`

	// DarkMode enables dark mode support
	DarkMode bool `yaml:"dark_mode" json:"dark_mode"`

	// Storybook enables Storybook for component development
	Storybook bool `yaml:"storybook" json:"storybook"`
}

// BackendConfig contains backend framework and service configuration.
type BackendConfig struct {
	// Enabled indicates if the project has a backend
	Enabled bool `yaml:"enabled" json:"enabled"`

	// Framework is the backend framework (fastapi, express, nestjs, go-gin, rust-axum)
	Framework string `yaml:"framework" json:"framework"`

	// FrameworkVersion is the framework version
	FrameworkVersion string `yaml:"framework_version,omitempty" json:"framework_version,omitempty"`

	// Language is the backend programming language
	Language string `yaml:"language" json:"language"`

	// LanguageVersion is the language version
	LanguageVersion string `yaml:"language_version,omitempty" json:"language_version,omitempty"`

	// Database contains database configuration
	Database DatabaseConfig `yaml:"database" json:"database"`

	// Auth contains authentication configuration
	Auth AuthConfig `yaml:"auth" json:"auth"`

	// API contains API configuration
	API APIConfig `yaml:"api" json:"api"`

	// Features contains optional backend features
	Features BackendFeatures `yaml:"features" json:"features"`

	// Directory is the backend source directory
	Directory string `yaml:"directory" json:"directory"`
}

// DatabaseConfig contains database configuration.
type DatabaseConfig struct {
	// Primary is the primary database type (postgresql, mysql, sqlite, mongodb)
	Primary string `yaml:"primary" json:"primary"`

	// PrimaryVersion is the database version
	PrimaryVersion string `yaml:"primary_version,omitempty" json:"primary_version,omitempty"`

	// ORM is the ORM/tool to use (prisma, sqlalchemy, gorm, mongoose)
	ORM string `yaml:"orm" json:"orm"`

	// Migrations indicates if database migrations are enabled
	Migrations bool `yaml:"migrations" json:"migrations"`

	// Redis indicates if Redis is used for caching
	Redis bool `yaml:"redis" json:"redis"`

	// RedisVersion is the Redis version
	RedisVersion string `yaml:"redis_version,omitempty" json:"redis_version,omitempty"`
}

// AuthConfig contains authentication configuration.
type AuthConfig struct {
	// Provider is the authentication provider (jwt, oauth, clerk, auth0, firebase)
	Provider string `yaml:"provider" json:"provider"`

	// Methods contains enabled authentication methods
	Methods []string `yaml:"methods,omitempty" json:"methods,omitempty"`

	// SessionDuration is the session duration in hours
	SessionDuration int `yaml:"session_duration" json:"session_duration"`
}

// APIConfig contains API configuration.
type APIConfig struct {
	// Style is the API style (rest, graphql, grpc, trpc)
	Style string `yaml:"style" json:"style"`

	// Versioning is the API versioning strategy (url, header, none)
	Versioning string `yaml:"versioning" json:"versioning"`

	// Documentation indicates if API documentation is generated
	Documentation bool `yaml:"documentation" json:"documentation"`

	// CORS contains CORS configuration
	CORS CORSConfig `yaml:"cors" json:"cors"`
}

// CORSConfig contains CORS configuration.
type CORSConfig struct {
	// Enabled indicates if CORS is enabled
	Enabled bool `yaml:"enabled" json:"enabled"`

	// Origins contains allowed origins
	Origins []string `yaml:"origins,omitempty" json:"origins,omitempty"`

	// Methods contains allowed HTTP methods
	Methods []string `yaml:"methods,omitempty" json:"methods,omitempty"`

	// Credentials indicates if credentials are allowed
	Credentials bool `yaml:"credentials" json:"credentials"`
}

// BackendFeatures contains optional backend feature flags.
type BackendFeatures struct {
	// WebSocket enables WebSocket support
	WebSocket bool `yaml:"websocket" json:"websocket"`

	// BackgroundJobs enables background job processing
	BackgroundJobs bool `yaml:"background_jobs" json:"background_jobs"`

	// FileUpload enables file upload handling
	FileUpload bool `yaml:"file_upload" json:"file_upload"`

	// Email enables email sending capabilities
	Email bool `yaml:"email" json:"email"`

	// RateLimiting enables API rate limiting
	RateLimiting bool `yaml:"rate_limiting" json:"rate_limiting"`

	// Logging enables structured logging
	Logging bool `yaml:"logging" json:"logging"`

	// Metrics enables metrics collection
	Metrics bool `yaml:"metrics" json:"metrics"`
}

// InfrastructureConfig contains deployment and infrastructure configuration.
type InfrastructureConfig struct {
	// Docker indicates if Docker is used
	Docker bool `yaml:"docker" json:"docker"`

	// DockerCompose indicates if Docker Compose is used for local development
	DockerCompose bool `yaml:"docker_compose" json:"docker_compose"`

	// Kubernetes indicates if Kubernetes manifests are generated
	Kubernetes bool `yaml:"kubernetes" json:"kubernetes"`

	// CI is the CI/CD platform (github-actions, gitlab-ci, circleci, jenkins)
	CI string `yaml:"ci,omitempty" json:"ci,omitempty"`

	// Hosting is the hosting platform (vercel, netlify, aws, gcp, azure, self-hosted)
	Hosting string `yaml:"hosting,omitempty" json:"hosting,omitempty"`

	// CDN indicates if a CDN is used
	CDN bool `yaml:"cdn" json:"cdn"`

	// Monitoring contains monitoring configuration
	Monitoring MonitoringConfig `yaml:"monitoring" json:"monitoring"`
}

// MonitoringConfig contains monitoring and observability configuration.
type MonitoringConfig struct {
	// Enabled indicates if monitoring is enabled
	Enabled bool `yaml:"enabled" json:"enabled"`

	// Provider is the monitoring provider (datadog, newrelic, prometheus, grafana)
	Provider string `yaml:"provider,omitempty" json:"provider,omitempty"`

	// ErrorTracking indicates if error tracking is enabled
	ErrorTracking bool `yaml:"error_tracking" json:"error_tracking"`

	// ErrorTrackingProvider is the error tracking provider (sentry, rollbar)
	ErrorTrackingProvider string `yaml:"error_tracking_provider,omitempty" json:"error_tracking_provider,omitempty"`

	// Logging contains logging configuration
	Logging LoggingConfig `yaml:"logging" json:"logging"`
}

// LoggingConfig contains logging configuration.
type LoggingConfig struct {
	// Level is the log level (debug, info, warn, error)
	Level string `yaml:"level" json:"level"`

	// Format is the log format (json, text)
	Format string `yaml:"format" json:"format"`

	// Provider is the logging provider (none, datadog, cloudwatch, stackdriver)
	Provider string `yaml:"provider,omitempty" json:"provider,omitempty"`
}

// GovernanceConfig contains AI governance and compliance settings.
type GovernanceConfig struct {
	// Enabled indicates if governance features are enabled
	Enabled bool `yaml:"enabled" json:"enabled"`

	// ContextLevel is the AI context detail level (minimal, standard, comprehensive)
	ContextLevel string `yaml:"context_level" json:"context_level"`

	// ComponentRegistry indicates if component registry is maintained
	ComponentRegistry bool `yaml:"component_registry" json:"component_registry"`

	// BrainstormMd indicates if Brainstorm.md is generated
	BrainstormMd bool `yaml:"brainstorm_md" json:"brainstorm_md"`

	// PromptGuidelines indicates if AI prompt guidelines are generated
	PromptGuidelines bool `yaml:"prompt_guidelines" json:"prompt_guidelines"`

	// Rules contains governance rules configuration
	Rules GovernanceRules `yaml:"rules" json:"rules"`

	// Documentation contains documentation standards
	Documentation DocumentationConfig `yaml:"documentation" json:"documentation"`
}

// GovernanceRules contains governance rule configuration.
type GovernanceRules struct {
	// Enabled indicates if rules enforcement is enabled
	Enabled bool `yaml:"enabled" json:"enabled"`

	// StrictMode enables strict rule enforcement (fails on warnings)
	StrictMode bool `yaml:"strict_mode" json:"strict_mode"`

	// CustomRulesPath is the path to custom rules file
	CustomRulesPath string `yaml:"custom_rules_path,omitempty" json:"custom_rules_path,omitempty"`

	// ExcludePatterns contains glob patterns for files to exclude from governance
	ExcludePatterns []string `yaml:"exclude_patterns,omitempty" json:"exclude_patterns,omitempty"`

	// Rules contains specific rule configurations
	Rules map[string]RuleConfig `yaml:"rules,omitempty" json:"rules,omitempty"`
}

// RuleConfig contains configuration for a specific rule.
type RuleConfig struct {
	// Enabled indicates if the rule is enabled
	Enabled bool `yaml:"enabled" json:"enabled"`

	// Severity is the rule severity (error, warning, info)
	Severity string `yaml:"severity" json:"severity"`

	// Options contains rule-specific options
	Options map[string]interface{} `yaml:"options,omitempty" json:"options,omitempty"`
}

// DocumentationConfig contains documentation standards configuration.
type DocumentationConfig struct {
	// README indicates if README.md is generated
	README bool `yaml:"readme" json:"readme"`

	// Contributing indicates if CONTRIBUTING.md is generated
	Contributing bool `yaml:"contributing" json:"contributing"`

	// Changelog indicates if CHANGELOG.md is generated
	Changelog bool `yaml:"changelog" json:"changelog"`

	// API indicates if API documentation is generated
	API bool `yaml:"api" json:"api"`

	// Inline indicates if inline code documentation is enforced
	Inline bool `yaml:"inline" json:"inline"`

	// Format is the documentation format (markdown, restructuredtext)
	Format string `yaml:"format" json:"format"`
}

// DevelopmentConfig contains development workflow settings.
type DevelopmentConfig struct {
	// Git indicates if git is initialized
	Git bool `yaml:"git" json:"git"`

	// Hooks contains git hooks configuration
	Hooks GitHooksConfig `yaml:"hooks" json:"hooks"`

	// Editor contains editor configuration
	Editor EditorConfig `yaml:"editor" json:"editor"`

	// Scripts contains custom npm/make scripts
	Scripts map[string]string `yaml:"scripts,omitempty" json:"scripts,omitempty"`
}

// GitHooksConfig contains git hooks configuration.
type GitHooksConfig struct {
	// PreCommit enables pre-commit hooks
	PreCommit bool `yaml:"pre_commit" json:"pre_commit"`

	// CommitMsg enables commit message validation
	CommitMsg bool `yaml:"commit_msg" json:"commit_msg"`

	// PrePush enables pre-push hooks
	PrePush bool `yaml:"pre_push" json:"pre_push"`

	// LintStaged enables linting of staged files
	LintStaged bool `yaml:"lint_staged" json:"lint_staged"`
}

// EditorConfig contains editor configuration.
type EditorConfig struct {
	// Config indicates if .editorconfig is generated
	Config bool `yaml:"config" json:"config"`

	// VSCode indicates if VS Code settings are generated
	VSCode bool `yaml:"vscode" json:"vscode"`

	// Extensions contains recommended VS Code extensions
	Extensions []string `yaml:"extensions,omitempty" json:"extensions,omitempty"`
}

// ConfigVersion is the current configuration schema version.
const ConfigVersion = "1.0.0"

// NewProjectConfig creates a new ProjectConfig with default values.
func NewProjectConfig() *ProjectConfig {
	now := time.Now()
	return &ProjectConfig{
		Version: ConfigVersion,
		Metadata: ProjectMetadata{
			Version:       "0.1.0",
			CreatedAt:     now,
			UpdatedAt:     now,
			ClauseVersion: "unknown",
		},
		Frontend: FrontendConfig{
			Enabled:       true,
			Framework:     "react",
			TypeScript:    true,
			Styling:       "tailwind",
			PackageManager: "npm",
			BuildTool:     "vite",
			TestFramework: "vitest",
			Linter:        "eslint",
			Formatter:     "prettier",
			Directory:     "src",
			Features: FrontendFeatures{
				DarkMode: true,
			},
		},
		Backend: BackendConfig{
			Enabled:   true,
			Framework: "fastapi",
			Language:  "python",
			Directory: "backend",
			Database: DatabaseConfig{
				Primary:    "postgresql",
				ORM:        "sqlalchemy",
				Migrations: true,
			},
			Auth: AuthConfig{
				Provider:        "jwt",
				Methods:         []string{"email", "password"},
				SessionDuration: 24,
			},
			API: APIConfig{
				Style:        "rest",
				Versioning:   "url",
				Documentation: true,
				CORS: CORSConfig{
					Enabled:     true,
					Origins:     []string{"http://localhost:3000"},
					Methods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
					Credentials: true,
				},
			},
			Features: BackendFeatures{
				Logging:  true,
				Metrics:  true,
				FileUpload: true,
			},
		},
		Infrastructure: InfrastructureConfig{
			Docker:         true,
			DockerCompose:  true,
			CI:             "github-actions",
			Hosting:        "vercel",
			Monitoring: MonitoringConfig{
				Enabled:         true,
				ErrorTracking:   true,
				ErrorTrackingProvider: "sentry",
				Logging: LoggingConfig{
					Level:   "info",
					Format:  "json",
					Provider: "none",
				},
			},
		},
		Governance: GovernanceConfig{
			Enabled:           true,
			ContextLevel:      "comprehensive",
			ComponentRegistry: true,
			BrainstormMd:      true,
			PromptGuidelines:  true,
			Rules: GovernanceRules{
				Enabled: true,
			},
			Documentation: DocumentationConfig{
				README:       true,
				Contributing: true,
				Changelog:    true,
				API:          true,
				Inline:       true,
				Format:       "markdown",
			},
		},
		Development: DevelopmentConfig{
			Git: true,
			Hooks: GitHooksConfig{
				PreCommit:  true,
				CommitMsg:  true,
				PrePush:    false,
				LintStaged: true,
			},
			Editor: EditorConfig{
				Config:  true,
				VSCode:  true,
			},
		},
	}
}

// Clone creates a deep copy of the ProjectConfig.
func (c *ProjectConfig) Clone() *ProjectConfig {
	cloned := *c

	// Clone slices
	if c.Metadata.Keywords != nil {
		cloned.Metadata.Keywords = make([]string, len(c.Metadata.Keywords))
		copy(cloned.Metadata.Keywords, c.Metadata.Keywords)
	}

	if c.Backend.Auth.Methods != nil {
		cloned.Backend.Auth.Methods = make([]string, len(c.Backend.Auth.Methods))
		copy(cloned.Backend.Auth.Methods, c.Backend.Auth.Methods)
	}

	if c.Backend.API.CORS.Origins != nil {
		cloned.Backend.API.CORS.Origins = make([]string, len(c.Backend.API.CORS.Origins))
		copy(cloned.Backend.API.CORS.Origins, c.Backend.API.CORS.Origins)
	}

	if c.Backend.API.CORS.Methods != nil {
		cloned.Backend.API.CORS.Methods = make([]string, len(c.Backend.API.CORS.Methods))
		copy(cloned.Backend.API.CORS.Methods, c.Backend.API.CORS.Methods)
	}

	if c.Governance.Rules.ExcludePatterns != nil {
		cloned.Governance.Rules.ExcludePatterns = make([]string, len(c.Governance.Rules.ExcludePatterns))
		copy(cloned.Governance.Rules.ExcludePatterns, c.Governance.Rules.ExcludePatterns)
	}

	if c.Governance.Rules.Rules != nil {
		cloned.Governance.Rules.Rules = make(map[string]RuleConfig)
		for k, v := range c.Governance.Rules.Rules {
			cloned.Governance.Rules.Rules[k] = v
		}
	}

	if c.Development.Scripts != nil {
		cloned.Development.Scripts = make(map[string]string)
		for k, v := range c.Development.Scripts {
			cloned.Development.Scripts[k] = v
		}
	}

	if c.Development.Editor.Extensions != nil {
		cloned.Development.Editor.Extensions = make([]string, len(c.Development.Editor.Extensions))
		copy(cloned.Development.Editor.Extensions, c.Development.Editor.Extensions)
	}

	return &cloned
}

// String returns a human-readable summary of the configuration.
func (c *ProjectConfig) String() string {
	return fmt.Sprintf("ProjectConfig{name=%s, frontend=%s, backend=%s}",
		c.Metadata.Name,
		c.Frontend.Framework,
		c.Backend.Framework,
	)
}
