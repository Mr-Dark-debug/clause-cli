package template

import (
	"fmt"
	"os"
	"time"

	"github.com/clause-cli/clause/internal/config"
)

// TemplateData provides data available to templates.
type TemplateData struct {
	// Project configuration
	*config.ProjectConfig

	// Additional computed values
	Project  ProjectData
	Frontend FrontendData
	Backend  BackendData
	Infra    InfraData
	Govern   GovernData

	// Current timestamp
	Now time.Time

	// Environment variables
	Env map[string]string

	// Custom variables
	Vars map[string]interface{}
}

// ProjectData contains project-level template data.
type ProjectData struct {
	// Name is the project name
	Name string

	// NameLower is the lowercase project name
	NameLower string

	// NameUpper is the uppercase project name
	NameUpper string

	// NameCamel is the camelCase project name
	NameCamel string

	// NamePascal is the PascalCase project name
	NamePascal string

	// NameSnake is the snake_case project name
	NameSnake string

	// NameKebab is the kebab-case project name
	NameKebab string

	// Description is the project description
	Description string

	// Version is the project version
	Version string

	// Author is the project author
	Author string

	// License is the project license
	License string

	// Repository is the project repository URL
	Repository string

	// GoModule is the Go module path
	GoModule string

	// PackageJSONName is the npm package name
	PackageJSONName string
}

// FrontendData contains frontend template data.
type FrontendData struct {
	// Enabled indicates if frontend is enabled
	Enabled bool

	// Framework is the frontend framework
	Framework string

	// TypeScript indicates if TypeScript is used
	TypeScript bool

	// Styling is the CSS approach
	Styling string

	// PackageManager is the package manager
	PackageManager string

	// BuildTool is the build tool
	BuildTool string

	// TestFramework is the test framework
	TestFramework string

	// Linter is the linter
	Linter string

	// Formatter is the formatter
	Formatter string

	// Directory is the frontend directory
	Directory string

	// Features contains feature flags
	Features FrontendFeaturesData
}

// FrontendFeaturesData contains frontend feature flags.
type FrontendFeaturesData struct {
	SSR      bool
	SSG      bool
	PWA      bool
	I18n     bool
	DarkMode bool
	Storybook bool
}

// BackendData contains backend template data.
type BackendData struct {
	// Enabled indicates if backend is enabled
	Enabled bool

	// Framework is the backend framework
	Framework string

	// Language is the programming language
	Language string

	// LanguageVersion is the language version
	LanguageVersion string

	// Directory is the backend directory
	Directory string

	// Database contains database configuration
	Database DatabaseData

	// Auth contains auth configuration
	Auth AuthData

	// API contains API configuration
	APIData APIData

	// Features contains feature flags
	Features BackendFeaturesData
}

// DatabaseData contains database template data.
type DatabaseData struct {
	Primary    string
	ORM        string
	Migrations bool
	Redis      bool
}

// AuthData contains auth template data.
type AuthData struct {
	Provider        string
	Methods         []string
	SessionDuration int
}

// APIData contains API template data.
type APIData struct {
	Style        string
	Versioning   string
	Documentation bool
}

// BackendFeaturesData contains backend feature flags.
type BackendFeaturesData struct {
	WebSocket      bool
	BackgroundJobs bool
	FileUpload     bool
	Email          bool
	RateLimiting   bool
	Logging        bool
	Metrics        bool
}

// InfraData contains infrastructure template data.
type InfraData struct {
	Docker        bool
	DockerCompose bool
	Kubernetes    bool
	CI            string
	Hosting       string
	CDN           bool
	Monitoring    MonitoringData
}

// MonitoringData contains monitoring template data.
type MonitoringData struct {
	Enabled       bool
	ErrorTracking bool
}

// GovernData contains governance template data.
type GovernData struct {
	Enabled           bool
	ContextLevel      string
	ComponentRegistry bool
	BrainstormMd      bool
	PromptGuidelines  bool
}

// NewTemplateData creates template data from a project configuration.
func NewTemplateData(cfg *config.ProjectConfig) *TemplateData {
	data := &TemplateData{
		ProjectConfig: cfg,
		Now:           time.Now(),
		Env:           getEnvMap(),
		Vars:          make(map[string]interface{}),
	}

	// Populate computed project data
	data.Project = ProjectData{
		Name:           cfg.Metadata.Name,
		NameLower:      cfg.Metadata.Name,
		NameUpper:      cfg.Metadata.Name,
		NameCamel:      cfg.Metadata.Name,
		NamePascal:     cfg.Metadata.Name,
		NameSnake:      cfg.Metadata.Name,
		NameKebab:      cfg.Metadata.Name,
		Description:    cfg.Metadata.Description,
		Version:        cfg.Metadata.Version,
		Author:         cfg.Metadata.Author,
		License:        cfg.Metadata.License,
		Repository:     cfg.Metadata.Repository,
		GoModule:       cfg.Metadata.Name,
		PackageJSONName: cfg.Metadata.Name,
	}

	// Populate frontend data
	data.Frontend = FrontendData{
		Enabled:       cfg.Frontend.Enabled,
		Framework:     cfg.Frontend.Framework,
		TypeScript:    cfg.Frontend.TypeScript,
		Styling:       cfg.Frontend.Styling,
		PackageManager: cfg.Frontend.PackageManager,
		BuildTool:     cfg.Frontend.BuildTool,
		TestFramework: cfg.Frontend.TestFramework,
		Linter:        cfg.Frontend.Linter,
		Formatter:     cfg.Frontend.Formatter,
		Directory:     cfg.Frontend.Directory,
		Features: FrontendFeaturesData{
			SSR:       cfg.Frontend.Features.SSR,
			SSG:       cfg.Frontend.Features.SSG,
			PWA:       cfg.Frontend.Features.PWA,
			I18n:      cfg.Frontend.Features.I18n,
			DarkMode:  cfg.Frontend.Features.DarkMode,
			Storybook: cfg.Frontend.Features.Storybook,
		},
	}

	// Populate backend data
	data.Backend = BackendData{
		Enabled:         cfg.Backend.Enabled,
		Framework:       cfg.Backend.Framework,
		Language:        cfg.Backend.Language,
		LanguageVersion: cfg.Backend.LanguageVersion,
		Directory:       cfg.Backend.Directory,
		Database: DatabaseData{
			Primary:    cfg.Backend.Database.Primary,
			ORM:        cfg.Backend.Database.ORM,
			Migrations: cfg.Backend.Database.Migrations,
			Redis:      cfg.Backend.Database.Redis,
		},
		Auth: AuthData{
			Provider:        cfg.Backend.Auth.Provider,
			Methods:         cfg.Backend.Auth.Methods,
			SessionDuration: cfg.Backend.Auth.SessionDuration,
		},
		APIData: APIData{
			Style:        cfg.Backend.API.Style,
			Versioning:   cfg.Backend.API.Versioning,
			Documentation: cfg.Backend.API.Documentation,
		},
		Features: BackendFeaturesData{
			WebSocket:      cfg.Backend.Features.WebSocket,
			BackgroundJobs: cfg.Backend.Features.BackgroundJobs,
			FileUpload:     cfg.Backend.Features.FileUpload,
			Email:          cfg.Backend.Features.Email,
			RateLimiting:   cfg.Backend.Features.RateLimiting,
			Logging:        cfg.Backend.Features.Logging,
			Metrics:        cfg.Backend.Features.Metrics,
		},
	}

	// Populate infrastructure data
	data.Infra = InfraData{
		Docker:        cfg.Infrastructure.Docker,
		DockerCompose: cfg.Infrastructure.DockerCompose,
		Kubernetes:    cfg.Infrastructure.Kubernetes,
		CI:           cfg.Infrastructure.CI,
		Hosting:      cfg.Infrastructure.Hosting,
		CDN:          cfg.Infrastructure.CDN,
		Monitoring: MonitoringData{
			Enabled:       cfg.Infrastructure.Monitoring.Enabled,
			ErrorTracking: cfg.Infrastructure.Monitoring.ErrorTracking,
		},
	}

	// Populate governance data
	data.Govern = GovernData{
		Enabled:           cfg.Governance.Enabled,
		ContextLevel:      cfg.Governance.ContextLevel,
		ComponentRegistry: cfg.Governance.ComponentRegistry,
		BrainstormMd:      cfg.Governance.BrainstormMd,
		PromptGuidelines:  cfg.Governance.PromptGuidelines,
	}

	return data
}

// getEnvMap returns a map of environment variables.
func getEnvMap() map[string]string {
	env := make(map[string]string)
	for _, pair := range os.Environ() {
		parts := splitEnv(pair)
		if len(parts) == 2 {
			env[parts[0]] = parts[1]
		}
	}
	return env
}

// splitEnv splits an environment variable string.
func splitEnv(s string) []string {
	for i := 0; i < len(s); i++ {
		if s[i] == '=' {
			return []string{s[:i], s[i+1:]}
		}
	}
	return []string{s}
}

// SetVar sets a custom variable.
func (d *TemplateData) SetVar(key string, value interface{}) {
	if d.Vars == nil {
		d.Vars = make(map[string]interface{})
	}
	d.Vars[key] = value
}

// GetVar gets a custom variable.
func (d *TemplateData) GetVar(key string) (interface{}, bool) {
	if d.Vars == nil {
		return nil, false
	}
	v, ok := d.Vars[key]
	return v, ok
}

// HasFrontend returns true if frontend is enabled.
func (d *TemplateData) HasFrontend() bool {
	return d.Frontend.Enabled
}

// HasBackend returns true if backend is enabled.
func (d *TemplateData) HasBackend() bool {
	return d.Backend.Enabled
}

// IsFramework returns true if the framework matches.
func (d *TemplateData) IsFramework(framework string) bool {
	return d.Frontend.Framework == framework || d.Backend.Framework == framework
}

// IsDatabase returns true if the database matches.
func (d *TemplateData) IsDatabase(db string) bool {
	return d.Backend.Database.Primary == db
}

// IsLanguage returns true if the language matches.
func (d *TemplateData) IsLanguage(lang string) bool {
	return d.Backend.Language == lang
}

// FormatDate formats a date.
func (d *TemplateData) FormatDate(format string) string {
	return d.Now.Format(format)
}

// Year returns the current year.
func (d *TemplateData) Year() int {
	return d.Now.Year()
}

// String returns a string representation.
func (d *TemplateData) String() string {
	return fmt.Sprintf("TemplateData{Project: %s}", d.Project.Name)
}
