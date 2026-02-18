package config

import (
	"fmt"
	"strings"
)

// DefaultValues contains all default configuration values.
// These are used when a new project is created or when values are not specified.
var DefaultValues = &defaults{
	// Frontend defaults
	Frontend: frontendDefaults{
		Framework:       "react",
		FrameworkVersion: "18",
		TypeScript:      true,
		Styling:         "tailwind",
		PackageManager:  "npm",
		BuildTool:       "vite",
		TestFramework:   "vitest",
		Linter:          "eslint",
		Formatter:       "prettier",
		Directory:       "src",
	},

	// Backend defaults
	Backend: backendDefaults{
		Framework:       "fastapi",
		Language:        "python",
		LanguageVersion: "3.11",
		Directory:       "backend",
		Database: databaseDefaults{
			Primary:    "postgresql",
			ORM:        "sqlalchemy",
			Migrations: true,
			Redis:      false,
		},
		Auth: authDefaults{
			Provider:        "jwt",
			Methods:         []string{"email", "password"},
			SessionDuration: 24,
		},
		API: apiDefaults{
			Style:        "rest",
			Versioning:   "url",
			Documentation: true,
		},
	},

	// Infrastructure defaults
	Infrastructure: infrastructureDefaults{
		Docker:         true,
		DockerCompose:  true,
		Kubernetes:     false,
		CI:             "github-actions",
		Hosting:        "vercel",
		CDN:            true,
		Monitoring: monitoringDefaults{
			Enabled:       true,
			ErrorTracking: true,
			Logging: loggingDefaults{
				Level:   "info",
				Format:  "json",
			},
		},
	},

	// Governance defaults
	Governance: governanceDefaults{
		Enabled:           true,
		ContextLevel:      "comprehensive",
		ComponentRegistry: true,
		BrainstormMd:      true,
		PromptGuidelines:  true,
		Documentation: documentationDefaults{
			README:       true,
			Contributing: true,
			Changelog:    true,
			API:          true,
			Inline:       true,
			Format:       "markdown",
		},
	},

	// Development defaults
	Development: developmentDefaults{
		Git: true,
		Hooks: gitHooksDefaults{
			PreCommit:  true,
			CommitMsg:  true,
			PrePush:    false,
			LintStaged: true,
		},
		Editor: editorDefaults{
			Config: true,
			VSCode: true,
		},
	},
}

// defaults structure
type defaults struct {
	Frontend      frontendDefaults
	Backend       backendDefaults
	Infrastructure infrastructureDefaults
	Governance    governanceDefaults
	Development   developmentDefaults
}

type frontendDefaults struct {
	Framework       string
	FrameworkVersion string
	TypeScript      bool
	Styling         string
	PackageManager  string
	BuildTool       string
	TestFramework   string
	Linter          string
	Formatter       string
	Directory       string
}

type backendDefaults struct {
	Framework       string
	Language        string
	LanguageVersion string
	Directory       string
	Database        databaseDefaults
	Auth            authDefaults
	API             apiDefaults
}

type databaseDefaults struct {
	Primary    string
	ORM        string
	Migrations bool
	Redis      bool
}

type authDefaults struct {
	Provider        string
	Methods         []string
	SessionDuration int
}

type apiDefaults struct {
	Style         string
	Versioning    string
	Documentation bool
}

type infrastructureDefaults struct {
	Docker         bool
	DockerCompose  bool
	Kubernetes     bool
	CI             string
	Hosting        string
	CDN            bool
	Monitoring     monitoringDefaults
}

type monitoringDefaults struct {
	Enabled       bool
	ErrorTracking bool
	Logging       loggingDefaults
}

type loggingDefaults struct {
	Level  string
	Format string
}

type governanceDefaults struct {
	Enabled           bool
	ContextLevel      string
	ComponentRegistry bool
	BrainstormMd      bool
	PromptGuidelines  bool
	Documentation     documentationDefaults
}

type documentationDefaults struct {
	README       bool
	Contributing bool
	Changelog    bool
	API          bool
	Inline       bool
	Format       string
}

type developmentDefaults struct {
	Git    bool
	Hooks  gitHooksDefaults
	Editor editorDefaults
}

type gitHooksDefaults struct {
	PreCommit  bool
	CommitMsg  bool
	PrePush    bool
	LintStaged bool
}

type editorDefaults struct {
	Config bool
	VSCode bool
}

// ApplyDefaults applies default values to a configuration.
// This fills in any missing values with their defaults.
func ApplyDefaults(config *ProjectConfig) {
	// Apply frontend defaults
	if config.Frontend.Framework == "" {
		config.Frontend.Framework = DefaultValues.Frontend.Framework
	}
	if config.Frontend.Styling == "" {
		config.Frontend.Styling = DefaultValues.Frontend.Styling
	}
	if config.Frontend.PackageManager == "" {
		config.Frontend.PackageManager = DefaultValues.Frontend.PackageManager
	}
	if config.Frontend.BuildTool == "" {
		config.Frontend.BuildTool = DefaultValues.Frontend.BuildTool
	}
	if config.Frontend.TestFramework == "" {
		config.Frontend.TestFramework = DefaultValues.Frontend.TestFramework
	}
	if config.Frontend.Linter == "" {
		config.Frontend.Linter = DefaultValues.Frontend.Linter
	}
	if config.Frontend.Formatter == "" {
		config.Frontend.Formatter = DefaultValues.Frontend.Formatter
	}
	if config.Frontend.Directory == "" {
		config.Frontend.Directory = DefaultValues.Frontend.Directory
	}

	// Apply backend defaults
	if config.Backend.Framework == "" {
		config.Backend.Framework = DefaultValues.Backend.Framework
	}
	if config.Backend.Language == "" {
		config.Backend.Language = DefaultValues.Backend.Language
	}
	if config.Backend.Directory == "" {
		config.Backend.Directory = DefaultValues.Backend.Directory
	}
	if config.Backend.Database.Primary == "" {
		config.Backend.Database.Primary = DefaultValues.Backend.Database.Primary
	}
	if config.Backend.Database.ORM == "" {
		config.Backend.Database.ORM = DefaultValues.Backend.Database.ORM
	}
	if config.Backend.Auth.Provider == "" {
		config.Backend.Auth.Provider = DefaultValues.Backend.Auth.Provider
	}
	if len(config.Backend.Auth.Methods) == 0 {
		config.Backend.Auth.Methods = DefaultValues.Backend.Auth.Methods
	}
	if config.Backend.Auth.SessionDuration == 0 {
		config.Backend.Auth.SessionDuration = DefaultValues.Backend.Auth.SessionDuration
	}
	if config.Backend.API.Style == "" {
		config.Backend.API.Style = DefaultValues.Backend.API.Style
	}
	if config.Backend.API.Versioning == "" {
		config.Backend.API.Versioning = DefaultValues.Backend.API.Versioning
	}

	// Apply infrastructure defaults
	if config.Infrastructure.CI == "" {
		config.Infrastructure.CI = DefaultValues.Infrastructure.CI
	}
	if config.Infrastructure.Hosting == "" {
		config.Infrastructure.Hosting = DefaultValues.Infrastructure.Hosting
	}
	if config.Infrastructure.Monitoring.Logging.Level == "" {
		config.Infrastructure.Monitoring.Logging.Level = DefaultValues.Infrastructure.Monitoring.Logging.Level
	}
	if config.Infrastructure.Monitoring.Logging.Format == "" {
		config.Infrastructure.Monitoring.Logging.Format = DefaultValues.Infrastructure.Monitoring.Logging.Format
	}

	// Apply governance defaults
	if config.Governance.ContextLevel == "" {
		config.Governance.ContextLevel = DefaultValues.Governance.ContextLevel
	}
	if config.Governance.Documentation.Format == "" {
		config.Governance.Documentation.Format = DefaultValues.Governance.Documentation.Format
	}
}

// Preset represents a configuration preset.
type Preset struct {
	// Name is the preset name
	Name string

	// Description describes what the preset is for
	Description string

	// Apply applies the preset to a configuration
	Apply func(*ProjectConfig)
}

// AvailablePresets contains all available configuration presets.
var AvailablePresets = []Preset{
	{
		Name:        "minimal",
		Description: "Minimal configuration with only essentials",
		Apply:       applyMinimalPreset,
	},
	{
		Name:        "standard",
		Description: "Standard full-stack configuration",
		Apply:       applyStandardPreset,
	},
	{
		Name:        "saas",
		Description: "SaaS application with auth, payments, and multi-tenancy",
		Apply:       applySaaSPreset,
	},
	{
		Name:        "api-only",
		Description: "API-only backend service",
		Apply:       applyAPIOnlyPreset,
	},
	{
		Name:        "frontend-only",
		Description: "Frontend-only static site",
		Apply:       applyFrontendOnlyPreset,
	},
	{
		Name:        "enterprise",
		Description: "Enterprise configuration with full governance",
		Apply:       applyEnterprisePreset,
	},
}

// GetPreset returns a preset by name.
func GetPreset(name string) (*Preset, error) {
	for _, preset := range AvailablePresets {
		if preset.Name == name {
			return &preset, nil
		}
	}
	return nil, fmt.Errorf("preset not found: %s", name)
}

// PresetNames returns all available preset names.
func PresetNames() []string {
	names := make([]string, len(AvailablePresets))
	for i, preset := range AvailablePresets {
		names[i] = preset.Name
	}
	return names
}

// Preset implementations

func applyMinimalPreset(c *ProjectConfig) {
	// Minimal frontend
	c.Frontend.Enabled = true
	c.Frontend.Framework = "react"
	c.Frontend.TypeScript = true
	c.Frontend.Styling = "tailwind"
	c.Frontend.TestFramework = "vitest"
	c.Frontend.Features = FrontendFeatures{}

	// Minimal backend
	c.Backend.Enabled = true
	c.Backend.Framework = "fastapi"
	c.Backend.Database = DatabaseConfig{
		Primary: "sqlite",
		ORM:     "sqlalchemy",
	}
	c.Backend.API = APIConfig{
		Style:        "rest",
		Documentation: false,
	}
	c.Backend.Features = BackendFeatures{
		Logging: true,
	}

	// Minimal infrastructure
	c.Infrastructure.Docker = false
	c.Infrastructure.DockerCompose = true
	c.Infrastructure.CI = "github-actions"
	c.Infrastructure.Hosting = ""

	// Minimal governance
	c.Governance.Enabled = true
	c.Governance.ContextLevel = "minimal"
	c.Governance.ComponentRegistry = false
	c.Governance.BrainstormMd = false
	c.Governance.Documentation = DocumentationConfig{
		README: true,
	}
}

func applyStandardPreset(c *ProjectConfig) {
	// Standard is the default
	ApplyDefaults(c)
}

func applySaaSPreset(c *ProjectConfig) {
	// Full-featured frontend
	c.Frontend.Enabled = true
	c.Frontend.Framework = "nextjs"
	c.Frontend.TypeScript = true
	c.Frontend.Styling = "tailwind"
	c.Frontend.Features = FrontendFeatures{
		SSR:      true,
		DarkMode: true,
		I18n:     true,
	}

	// Full-featured backend
	c.Backend.Enabled = true
	c.Backend.Framework = "fastapi"
	c.Backend.Database = DatabaseConfig{
		Primary:    "postgresql",
		ORM:        "sqlalchemy",
		Migrations: true,
		Redis:      true,
	}
	c.Backend.Auth = AuthConfig{
		Provider:        "clerk",
		Methods:         []string{"email", "google", "github"},
		SessionDuration: 168, // 1 week
	}
	c.Backend.API = APIConfig{
		Style:         "rest",
		Versioning:    "url",
		Documentation: true,
	}
	c.Backend.Features = BackendFeatures{
		WebSocket:      true,
		BackgroundJobs: true,
		FileUpload:     true,
		Email:          true,
		Logging:        true,
		Metrics:        true,
		RateLimiting:   true,
	}

	// Full infrastructure
	c.Infrastructure.Docker = true
	c.Infrastructure.DockerCompose = true
	c.Infrastructure.Kubernetes = true
	c.Infrastructure.CI = "github-actions"
	c.Infrastructure.Hosting = "aws"
	c.Infrastructure.CDN = true
	c.Infrastructure.Monitoring = MonitoringConfig{
		Enabled:               true,
		Provider:              "datadog",
		ErrorTracking:         true,
		ErrorTrackingProvider: "sentry",
	}

	// Full governance
	c.Governance.Enabled = true
	c.Governance.ContextLevel = "comprehensive"
	c.Governance.ComponentRegistry = true
	c.Governance.BrainstormMd = true
	c.Governance.PromptGuidelines = true
}

func applyAPIOnlyPreset(c *ProjectConfig) {
	// No frontend
	c.Frontend.Enabled = false
	c.Frontend = FrontendConfig{}

	// API-focused backend
	c.Backend.Enabled = true
	c.Backend.Framework = "fastapi"
	c.Backend.Database = DatabaseConfig{
		Primary:    "postgresql",
		ORM:        "sqlalchemy",
		Migrations: true,
		Redis:      true,
	}
	c.Backend.API = APIConfig{
		Style:         "rest",
		Versioning:    "url",
		Documentation: true,
	}
	c.Backend.Features = BackendFeatures{
		Logging:      true,
		Metrics:      true,
		RateLimiting: true,
	}

	// API-focused infrastructure
	c.Infrastructure.Docker = true
	c.Infrastructure.DockerCompose = true
	c.Infrastructure.Kubernetes = true
	c.Infrastructure.CI = "github-actions"
	c.Infrastructure.Hosting = "aws"
	c.Infrastructure.Monitoring = MonitoringConfig{
		Enabled:       true,
		ErrorTracking: true,
	}
}

func applyFrontendOnlyPreset(c *ProjectConfig) {
	// Frontend-focused
	c.Frontend.Enabled = true
	c.Frontend.Framework = "nextjs"
	c.Frontend.TypeScript = true
	c.Frontend.Styling = "tailwind"
	c.Frontend.Features = FrontendFeatures{
		SSR:      true,
		SSG:      true,
		DarkMode: true,
		I18n:     true,
		PWA:      true,
	}

	// No backend
	c.Backend.Enabled = false
	c.Backend = BackendConfig{}

	// Frontend-focused infrastructure
	c.Infrastructure.Docker = false
	c.Infrastructure.DockerCompose = false
	c.Infrastructure.Kubernetes = false
	c.Infrastructure.CI = "github-actions"
	c.Infrastructure.Hosting = "vercel"
	c.Infrastructure.CDN = true
}

func applyEnterprisePreset(c *ProjectConfig) {
	// Apply SaaS preset as base
	applySaaSPreset(c)

	// Additional enterprise features
	c.Governance.Rules = GovernanceRules{
		Enabled:    true,
		StrictMode: true,
	}

	c.Infrastructure.Monitoring.Logging = LoggingConfig{
		Level:   "warn",
		Format:  "json",
		Provider: "datadog",
	}

	c.Development.Hooks = GitHooksConfig{
		PreCommit:  true,
		CommitMsg:  true,
		PrePush:    true,
		LintStaged: true,
	}
}

// GetDefaultFor returns the default value for a configuration key path.
func GetDefaultFor(keyPath string) (interface{}, error) {
	parts := strings.Split(keyPath, ".")
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty key path")
	}

	switch parts[0] {
	case "frontend":
		return getFrontendDefault(parts[1:])
	case "backend":
		return getBackendDefault(parts[1:])
	case "infrastructure":
		return getInfrastructureDefault(parts[1:])
	case "governance":
		return getGovernanceDefault(parts[1:])
	case "development":
		return getDevelopmentDefault(parts[1:])
	default:
		return nil, fmt.Errorf("unknown configuration section: %s", parts[0])
	}
}

func getFrontendDefault(parts []string) (interface{}, error) {
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty frontend path")
	}

	switch parts[0] {
	case "framework":
		return DefaultValues.Frontend.Framework, nil
	case "typescript":
		return DefaultValues.Frontend.TypeScript, nil
	case "styling":
		return DefaultValues.Frontend.Styling, nil
	case "package_manager":
		return DefaultValues.Frontend.PackageManager, nil
	case "build_tool":
		return DefaultValues.Frontend.BuildTool, nil
	case "test_framework":
		return DefaultValues.Frontend.TestFramework, nil
	case "linter":
		return DefaultValues.Frontend.Linter, nil
	case "formatter":
		return DefaultValues.Frontend.Formatter, nil
	case "directory":
		return DefaultValues.Frontend.Directory, nil
	default:
		return nil, fmt.Errorf("unknown frontend field: %s", parts[0])
	}
}

func getBackendDefault(parts []string) (interface{}, error) {
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty backend path")
	}

	switch parts[0] {
	case "framework":
		return DefaultValues.Backend.Framework, nil
	case "language":
		return DefaultValues.Backend.Language, nil
	case "directory":
		return DefaultValues.Backend.Directory, nil
	case "database":
		if len(parts) < 2 {
			return nil, fmt.Errorf("incomplete database path")
		}
		switch parts[1] {
		case "primary":
			return DefaultValues.Backend.Database.Primary, nil
		case "orm":
			return DefaultValues.Backend.Database.ORM, nil
		case "migrations":
			return DefaultValues.Backend.Database.Migrations, nil
		default:
			return nil, fmt.Errorf("unknown database field: %s", parts[1])
		}
	case "auth":
		if len(parts) < 2 {
			return nil, fmt.Errorf("incomplete auth path")
		}
		switch parts[1] {
		case "provider":
			return DefaultValues.Backend.Auth.Provider, nil
		case "session_duration":
			return DefaultValues.Backend.Auth.SessionDuration, nil
		default:
			return nil, fmt.Errorf("unknown auth field: %s", parts[1])
		}
	case "api":
		if len(parts) < 2 {
			return nil, fmt.Errorf("incomplete api path")
		}
		switch parts[1] {
		case "style":
			return DefaultValues.Backend.API.Style, nil
		case "versioning":
			return DefaultValues.Backend.API.Versioning, nil
		default:
			return nil, fmt.Errorf("unknown api field: %s", parts[1])
		}
	default:
		return nil, fmt.Errorf("unknown backend field: %s", parts[0])
	}
}

func getInfrastructureDefault(parts []string) (interface{}, error) {
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty infrastructure path")
	}

	switch parts[0] {
	case "docker":
		return DefaultValues.Infrastructure.Docker, nil
	case "docker_compose":
		return DefaultValues.Infrastructure.DockerCompose, nil
	case "ci":
		return DefaultValues.Infrastructure.CI, nil
	case "hosting":
		return DefaultValues.Infrastructure.Hosting, nil
	case "cdn":
		return DefaultValues.Infrastructure.CDN, nil
	default:
		return nil, fmt.Errorf("unknown infrastructure field: %s", parts[0])
	}
}

func getGovernanceDefault(parts []string) (interface{}, error) {
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty governance path")
	}

	switch parts[0] {
	case "context_level":
		return DefaultValues.Governance.ContextLevel, nil
	case "component_registry":
		return DefaultValues.Governance.ComponentRegistry, nil
	case "brainstorm_md":
		return DefaultValues.Governance.BrainstormMd, nil
	case "prompt_guidelines":
		return DefaultValues.Governance.PromptGuidelines, nil
	default:
		return nil, fmt.Errorf("unknown governance field: %s", parts[0])
	}
}

func getDevelopmentDefault(parts []string) (interface{}, error) {
	if len(parts) == 0 {
		return nil, fmt.Errorf("empty development path")
	}

	switch parts[0] {
	case "git":
		return DefaultValues.Development.Git, nil
	default:
		return nil, fmt.Errorf("unknown development field: %s", parts[0])
	}
}

// LoadPreset creates a new ProjectConfig with the specified preset applied.
func LoadPreset(name string) (*ProjectConfig, error) {
	preset, err := GetPreset(name)
	if err != nil {
		return nil, err
	}

	config := NewProjectConfig()
	preset.Apply(config)

	return config, nil
}
