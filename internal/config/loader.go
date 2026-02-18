package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"

	"github.com/clause-cli/clause/pkg/utils"
)

// Loader handles loading configuration from multiple sources with priority.
// Priority order (highest to lowest):
// 1. Explicit flags/options
// 2. Environment variables
// 3. Project configuration (.clause/config.yaml)
// 4. Global configuration (~/.clause/config.yaml)
// 5. Default values
type Loader struct {
	// projectDir is the project directory path
	projectDir string

	// globalDir is the global Clause configuration directory
	globalDir string

	// envPrefix is the prefix for environment variables
	envPrefix string

	// overrides contains explicit flag/option overrides
	overrides map[string]interface{}
}

// LoaderOption is a functional option for configuring the Loader.
type LoaderOption func(*Loader)

// WithProjectDir sets the project directory for the loader.
func WithProjectDir(dir string) LoaderOption {
	return func(l *Loader) {
		l.projectDir = dir
	}
}

// WithGlobalDir sets the global configuration directory.
func WithGlobalDir(dir string) LoaderOption {
	return func(l *Loader) {
		l.globalDir = dir
	}
}

// WithEnvPrefix sets the environment variable prefix.
func WithEnvPrefix(prefix string) LoaderOption {
	return func(l *Loader) {
		l.envPrefix = prefix
	}
}

// WithOverrides sets explicit configuration overrides.
func WithOverrides(overrides map[string]interface{}) LoaderOption {
	return func(l *Loader) {
		l.overrides = overrides
	}
}

// NewLoader creates a new configuration loader with the given options.
func NewLoader(opts ...LoaderOption) *Loader {
	home := utils.GetHomeDirectory()

	l := &Loader{
		globalDir: filepath.Join(home, ".clause"),
		envPrefix: "CLAUSE_",
		overrides: make(map[string]interface{}),
	}

	for _, opt := range opts {
		opt(l)
	}

	return l
}

// Load loads configuration from all sources and merges them.
func (l *Loader) Load() (*ProjectConfig, error) {
	// Start with defaults
	config := NewProjectConfig()

	// Load global configuration (lowest priority)
	if err := l.loadGlobalConfig(config); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load global config: %w", err)
	}

	// Load project configuration
	if err := l.loadProjectConfig(config); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to load project config: %w", err)
	}

	// Apply environment variables
	l.applyEnvVars(config)

	// Apply explicit overrides (highest priority)
	l.applyOverrides(config)

	return config, nil
}

// LoadFromPath loads configuration from a specific file path.
func (l *Loader) LoadFromPath(path string) (*ProjectConfig, error) {
	config := NewProjectConfig()

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".yaml", ".yml":
		if err := yaml.Unmarshal(data, config); err != nil {
			return nil, fmt.Errorf("failed to parse YAML config: %w", err)
		}
	case ".json":
		if err := json.Unmarshal(data, config); err != nil {
			return nil, fmt.Errorf("failed to parse JSON config: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported config format: %s", ext)
	}

	return config, nil
}

// loadGlobalConfig loads configuration from the global config directory.
func (l *Loader) loadGlobalConfig(config *ProjectConfig) error {
	configPath := filepath.Join(l.globalDir, "config.yaml")
	return l.mergeConfigFile(config, configPath)
}

// loadProjectConfig loads configuration from the project directory.
func (l *Loader) loadProjectConfig(config *ProjectConfig) error {
	if l.projectDir == "" {
		return nil
	}

	// Check multiple possible config locations
	locations := []string{
		filepath.Join(l.projectDir, ".clause", "config.yaml"),
		filepath.Join(l.projectDir, ".clause", "config.yml"),
		filepath.Join(l.projectDir, "clause.yaml"),
		filepath.Join(l.projectDir, "clause.yml"),
	}

	for _, path := range locations {
		if utils.FileExists(path) {
			return l.mergeConfigFile(config, path)
		}
	}

	return os.ErrNotExist
}

// mergeConfigFile merges a configuration file into the existing config.
func (l *Loader) mergeConfigFile(config *ProjectConfig, path string) error {
	if !utils.FileExists(path) {
		return os.ErrNotExist
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	// Parse as generic map first for partial updates
	var partial map[string]interface{}
	if err := yaml.Unmarshal(data, &partial); err != nil {
		return fmt.Errorf("failed to parse config file: %w", err)
	}

	// Merge into config
	return mergeMapIntoConfig(config, partial)
}

// applyEnvVars applies environment variable overrides to the config.
func (l *Loader) applyEnvVars(config *ProjectConfig) {
	envMappings := map[string]func(string){
		"CLAUSE_FRONTEND_FRAMEWORK":      func(v string) { config.Frontend.Framework = v },
		"CLAUSE_FRONTEND_STYLING":        func(v string) { config.Frontend.Styling = v },
		"CLAUSE_FRONTEND_TYPESCRIPT":     func(v string) { config.Frontend.TypeScript = parseBool(v) },
		"CLAUSE_FRONTEND_PACKAGE_MANAGER": func(v string) { config.Frontend.PackageManager = v },
		"CLAUSE_BACKEND_FRAMEWORK":       func(v string) { config.Backend.Framework = v },
		"CLAUSE_BACKEND_LANGUAGE":        func(v string) { config.Backend.Language = v },
		"CLAUSE_BACKEND_DATABASE":        func(v string) { config.Backend.Database.Primary = v },
		"CLAUSE_BACKEND_ORM":             func(v string) { config.Backend.Database.ORM = v },
		"CLAUSE_INFRASTRUCTURE_CI":       func(v string) { config.Infrastructure.CI = v },
		"CLAUSE_INFRASTRUCTURE_HOSTING":  func(v string) { config.Infrastructure.Hosting = v },
		"CLAUSE_GOVERNANCE_ENABLED":      func(v string) { config.Governance.Enabled = parseBool(v) },
		"CLAUSE_GOVERNANCE_CONTEXT_LEVEL": func(v string) { config.Governance.ContextLevel = v },
		"CLAUSE_DEBUG":                   func(v string) { /* handled elsewhere */ },
		"CLAUSE_QUIET":                   func(v string) { /* handled elsewhere */ },
		"CLAUSE_NO_COLOR":                func(v string) { /* handled elsewhere */ },
	}

	for envKey, setter := range envMappings {
		if value := os.Getenv(envKey); value != "" {
			setter(value)
		}
	}
}

// applyOverrides applies explicit overrides to the config.
func (l *Loader) applyOverrides(config *ProjectConfig) {
	if len(l.overrides) == 0 {
		return
	}

	_ = mergeMapIntoConfig(config, l.overrides)
}

// mergeMapIntoConfig merges a generic map into a ProjectConfig struct.
func mergeMapIntoConfig(config *ProjectConfig, m map[string]interface{}) error {
	// Handle metadata
	if metadata, ok := m["metadata"].(map[string]interface{}); ok {
		if name, ok := metadata["name"].(string); ok {
			config.Metadata.Name = name
		}
		if desc, ok := metadata["description"].(string); ok {
			config.Metadata.Description = desc
		}
		if version, ok := metadata["version"].(string); ok {
			config.Metadata.Version = version
		}
		if author, ok := metadata["author"].(string); ok {
			config.Metadata.Author = author
		}
		if license, ok := metadata["license"].(string); ok {
			config.Metadata.License = license
		}
		if repo, ok := metadata["repository"].(string); ok {
			config.Metadata.Repository = repo
		}
		if keywords, ok := metadata["keywords"].([]interface{}); ok {
			config.Metadata.Keywords = toStringSlice(keywords)
		}
	}

	// Handle frontend
	if frontend, ok := m["frontend"].(map[string]interface{}); ok {
		if enabled, ok := frontend["enabled"].(bool); ok {
			config.Frontend.Enabled = enabled
		}
		if framework, ok := frontend["framework"].(string); ok {
			config.Frontend.Framework = framework
		}
		if ts, ok := frontend["typescript"].(bool); ok {
			config.Frontend.TypeScript = ts
		}
		if styling, ok := frontend["styling"].(string); ok {
			config.Frontend.Styling = styling
		}
		if pm, ok := frontend["package_manager"].(string); ok {
			config.Frontend.PackageManager = pm
		}
		if buildTool, ok := frontend["build_tool"].(string); ok {
			config.Frontend.BuildTool = buildTool
		}
		if testFramework, ok := frontend["test_framework"].(string); ok {
			config.Frontend.TestFramework = testFramework
		}
		if linter, ok := frontend["linter"].(string); ok {
			config.Frontend.Linter = linter
		}
		if formatter, ok := frontend["formatter"].(string); ok {
			config.Frontend.Formatter = formatter
		}
		if dir, ok := frontend["directory"].(string); ok {
			config.Frontend.Directory = dir
		}
		if features, ok := frontend["features"].(map[string]interface{}); ok {
			mergeFrontendFeatures(&config.Frontend.Features, features)
		}
	}

	// Handle backend
	if backend, ok := m["backend"].(map[string]interface{}); ok {
		if enabled, ok := backend["enabled"].(bool); ok {
			config.Backend.Enabled = enabled
		}
		if framework, ok := backend["framework"].(string); ok {
			config.Backend.Framework = framework
		}
		if language, ok := backend["language"].(string); ok {
			config.Backend.Language = language
		}
		if dir, ok := backend["directory"].(string); ok {
			config.Backend.Directory = dir
		}
		if database, ok := backend["database"].(map[string]interface{}); ok {
			mergeDatabaseConfig(&config.Backend.Database, database)
		}
		if auth, ok := backend["auth"].(map[string]interface{}); ok {
			mergeAuthConfig(&config.Backend.Auth, auth)
		}
		if api, ok := backend["api"].(map[string]interface{}); ok {
			mergeAPIConfig(&config.Backend.API, api)
		}
		if features, ok := backend["features"].(map[string]interface{}); ok {
			mergeBackendFeatures(&config.Backend.Features, features)
		}
	}

	// Handle infrastructure
	if infra, ok := m["infrastructure"].(map[string]interface{}); ok {
		if docker, ok := infra["docker"].(bool); ok {
			config.Infrastructure.Docker = docker
		}
		if compose, ok := infra["docker_compose"].(bool); ok {
			config.Infrastructure.DockerCompose = compose
		}
		if ci, ok := infra["ci"].(string); ok {
			config.Infrastructure.CI = ci
		}
		if hosting, ok := infra["hosting"].(string); ok {
			config.Infrastructure.Hosting = hosting
		}
		if monitoring, ok := infra["monitoring"].(map[string]interface{}); ok {
			mergeMonitoringConfig(&config.Infrastructure.Monitoring, monitoring)
		}
	}

	// Handle governance
	if governance, ok := m["governance"].(map[string]interface{}); ok {
		if enabled, ok := governance["enabled"].(bool); ok {
			config.Governance.Enabled = enabled
		}
		if contextLevel, ok := governance["context_level"].(string); ok {
			config.Governance.ContextLevel = contextLevel
		}
		if registry, ok := governance["component_registry"].(bool); ok {
			config.Governance.ComponentRegistry = registry
		}
		if brainstorm, ok := governance["brainstorm_md"].(bool); ok {
			config.Governance.BrainstormMd = brainstorm
		}
		if guidelines, ok := governance["prompt_guidelines"].(bool); ok {
			config.Governance.PromptGuidelines = guidelines
		}
	}

	// Handle development
	if dev, ok := m["development"].(map[string]interface{}); ok {
		if git, ok := dev["git"].(bool); ok {
			config.Development.Git = git
		}
		if hooks, ok := dev["hooks"].(map[string]interface{}); ok {
			mergeGitHooksConfig(&config.Development.Hooks, hooks)
		}
		if editor, ok := dev["editor"].(map[string]interface{}); ok {
			mergeEditorConfig(&config.Development.Editor, editor)
		}
	}

	return nil
}

// Helper functions for merging nested configs

func mergeFrontendFeatures(f *FrontendFeatures, m map[string]interface{}) {
	if ssr, ok := m["ssr"].(bool); ok {
		f.SSR = ssr
	}
	if ssg, ok := m["ssg"].(bool); ok {
		f.SSG = ssg
	}
	if pwa, ok := m["pwa"].(bool); ok {
		f.PWA = pwa
	}
	if i18n, ok := m["i18n"].(bool); ok {
		f.I18n = i18n
	}
	if darkMode, ok := m["dark_mode"].(bool); ok {
		f.DarkMode = darkMode
	}
	if storybook, ok := m["storybook"].(bool); ok {
		f.Storybook = storybook
	}
}

func mergeDatabaseConfig(d *DatabaseConfig, m map[string]interface{}) {
	if primary, ok := m["primary"].(string); ok {
		d.Primary = primary
	}
	if orm, ok := m["orm"].(string); ok {
		d.ORM = orm
	}
	if migrations, ok := m["migrations"].(bool); ok {
		d.Migrations = migrations
	}
	if redis, ok := m["redis"].(bool); ok {
		d.Redis = redis
	}
}

func mergeAuthConfig(a *AuthConfig, m map[string]interface{}) {
	if provider, ok := m["provider"].(string); ok {
		a.Provider = provider
	}
	if methods, ok := m["methods"].([]interface{}); ok {
		a.Methods = toStringSlice(methods)
	}
	if duration, ok := m["session_duration"].(int); ok {
		a.SessionDuration = duration
	}
}

func mergeAPIConfig(a *APIConfig, m map[string]interface{}) {
	if style, ok := m["style"].(string); ok {
		a.Style = style
	}
	if versioning, ok := m["versioning"].(string); ok {
		a.Versioning = versioning
	}
	if docs, ok := m["documentation"].(bool); ok {
		a.Documentation = docs
	}
	if cors, ok := m["cors"].(map[string]interface{}); ok {
		if enabled, ok := cors["enabled"].(bool); ok {
			a.CORS.Enabled = enabled
		}
		if origins, ok := cors["origins"].([]interface{}); ok {
			a.CORS.Origins = toStringSlice(origins)
		}
		if methods, ok := cors["methods"].([]interface{}); ok {
			a.CORS.Methods = toStringSlice(methods)
		}
		if creds, ok := cors["credentials"].(bool); ok {
			a.CORS.Credentials = creds
		}
	}
}

func mergeBackendFeatures(f *BackendFeatures, m map[string]interface{}) {
	if ws, ok := m["websocket"].(bool); ok {
		f.WebSocket = ws
	}
	if jobs, ok := m["background_jobs"].(bool); ok {
		f.BackgroundJobs = jobs
	}
	if upload, ok := m["file_upload"].(bool); ok {
		f.FileUpload = upload
	}
	if email, ok := m["email"].(bool); ok {
		f.Email = email
	}
	if rateLimit, ok := m["rate_limiting"].(bool); ok {
		f.RateLimiting = rateLimit
	}
	if logging, ok := m["logging"].(bool); ok {
		f.Logging = logging
	}
	if metrics, ok := m["metrics"].(bool); ok {
		f.Metrics = metrics
	}
}

func mergeMonitoringConfig(m *MonitoringConfig, cfg map[string]interface{}) {
	if enabled, ok := cfg["enabled"].(bool); ok {
		m.Enabled = enabled
	}
	if provider, ok := cfg["provider"].(string); ok {
		m.Provider = provider
	}
	if tracking, ok := cfg["error_tracking"].(bool); ok {
		m.ErrorTracking = tracking
	}
	if trackingProvider, ok := cfg["error_tracking_provider"].(string); ok {
		m.ErrorTrackingProvider = trackingProvider
	}
	if logging, ok := cfg["logging"].(map[string]interface{}); ok {
		if level, ok := logging["level"].(string); ok {
			m.Logging.Level = level
		}
		if format, ok := logging["format"].(string); ok {
			m.Logging.Format = format
		}
		if provider, ok := logging["provider"].(string); ok {
			m.Logging.Provider = provider
		}
	}
}

func mergeGitHooksConfig(g *GitHooksConfig, m map[string]interface{}) {
	if preCommit, ok := m["pre_commit"].(bool); ok {
		g.PreCommit = preCommit
	}
	if commitMsg, ok := m["commit_msg"].(bool); ok {
		g.CommitMsg = commitMsg
	}
	if prePush, ok := m["pre_push"].(bool); ok {
		g.PrePush = prePush
	}
	if lintStaged, ok := m["lint_staged"].(bool); ok {
		g.LintStaged = lintStaged
	}
}

func mergeEditorConfig(e *EditorConfig, m map[string]interface{}) {
	if config, ok := m["config"].(bool); ok {
		e.Config = config
	}
	if vscode, ok := m["vscode"].(bool); ok {
		e.VSCode = vscode
	}
	if exts, ok := m["extensions"].([]interface{}); ok {
		e.Extensions = toStringSlice(exts)
	}
}

// toStringSlice converts an []interface{} to []string.
func toStringSlice(slice []interface{}) []string {
	result := make([]string, 0, len(slice))
	for _, v := range slice {
		if s, ok := v.(string); ok {
			result = append(result, s)
		}
	}
	return result
}

// parseBool parses a string to bool with common variations.
func parseBool(s string) bool {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "1", "yes", "on", "enabled":
		return true
	default:
		return false
	}
}

// FindProjectConfig searches for a project configuration file starting from dir
// and walking up the directory tree.
func FindProjectConfig(dir string) (string, error) {
	locations := []string{
		".clause/config.yaml",
		".clause/config.yml",
		"clause.yaml",
		"clause.yml",
	}

	for _, loc := range locations {
		result := utils.FindFileUp(loc, dir)
		if result != "" {
			return result, nil
		}
	}

	return "", fmt.Errorf("no project configuration found in %s or parent directories", dir)
}

// LoadProject loads the configuration for the current project.
// It searches for the configuration file starting from the current directory.
func LoadProject() (*ProjectConfig, string, error) {
	// Get current directory
	cwd, err := os.Getwd()
	if err != nil {
		return nil, "", fmt.Errorf("failed to get current directory: %w", err)
	}

	// Find project config
	configPath, err := FindProjectConfig(cwd)
	if err != nil {
		return nil, "", err
	}

	// Load the config
	loader := NewLoader()
	config, err := loader.LoadFromPath(configPath)
	if err != nil {
		return nil, "", fmt.Errorf("failed to load project config: %w", err)
	}

	return config, configPath, nil
}
