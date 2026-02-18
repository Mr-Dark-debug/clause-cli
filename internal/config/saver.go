package config

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"

	"github.com/clause-cli/clause/pkg/utils"
)

// Saver handles saving configuration to files.
type Saver struct {
	// format is the output format (yaml or json)
	format string

	// indent is the indentation string for output
	indent string

	// backup indicates whether to create backups before overwriting
	backup bool
}

// SaverOption is a functional option for configuring the Saver.
type SaverOption func(*Saver)

// WithFormat sets the output format.
func WithFormat(format string) SaverOption {
	return func(s *Saver) {
		s.format = format
	}
}

// WithIndent sets the indentation string.
func WithIndent(indent string) SaverOption {
	return func(s *Saver) {
		s.indent = indent
	}
}

// WithBackup enables or disables backup creation.
func WithBackup(backup bool) SaverOption {
	return func(s *Saver) {
		s.backup = backup
	}
}

// NewSaver creates a new configuration saver with the given options.
func NewSaver(opts ...SaverOption) *Saver {
	s := &Saver{
		format:  "yaml",
		indent:  "  ",
		backup:  true,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

// Save saves the configuration to the specified file path.
func (s *Saver) Save(config *ProjectConfig, path string) error {
	// Update the timestamp
	config.Metadata.UpdatedAt = time.Now()

	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := utils.EnsureDirectory(dir); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Create backup if file exists and backup is enabled
	if s.backup && utils.FileExists(path) {
		if err := s.createBackup(path); err != nil {
			return fmt.Errorf("failed to create backup: %w", err)
		}
	}

	// Marshal configuration
	var data []byte
	var err error

	switch strings.ToLower(s.format) {
	case "yaml", "yml":
		data, err = yaml.Marshal(config)
	case "json":
		data, err = json.MarshalIndent(config, "", s.indent)
	default:
		return fmt.Errorf("unsupported format: %s", s.format)
	}

	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	// Write atomically
	if err := utils.AtomicWrite(path, data); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}

	return nil
}

// SaveToProject saves the configuration to a project directory.
// The config will be saved to .clause/config.yaml within the project.
func (s *Saver) SaveToProject(config *ProjectConfig, projectDir string) error {
	configPath := filepath.Join(projectDir, ".clause", "config.yaml")
	return s.Save(config, configPath)
}

// SaveToGlobal saves the configuration to the global configuration directory.
func (s *Saver) SaveToGlobal(config *ProjectConfig) error {
	home := utils.GetHomeDirectory()

	configPath := filepath.Join(home, ".clause", "config.yaml")
	return s.Save(config, configPath)
}

// createBackup creates a backup of the existing configuration file.
func (s *Saver) createBackup(path string) error {
	backupPath := path + ".backup"
	return utils.CopyFile(path, backupPath)
}

// Save saves a project configuration to the specified path using default options.
func Save(config *ProjectConfig, path string) error {
	return NewSaver().Save(config, path)
}

// SaveToProject saves a project configuration using default options.
func SaveToProject(config *ProjectConfig, projectDir string) error {
	return NewSaver().SaveToProject(config, projectDir)
}

// SaveToGlobal saves a configuration to the global directory using default options.
func SaveToGlobal(config *ProjectConfig) error {
	return NewSaver().SaveToGlobal(config)
}

// QuickSave saves configuration to a specific path without backups.
func QuickSave(config *ProjectConfig, path string) error {
	return NewSaver(WithBackup(false)).Save(config, path)
}

// InitProjectConfig creates a new project configuration file.
func InitProjectConfig(projectDir string, projectName string) (*ProjectConfig, error) {
	// Create default config
	config := NewProjectConfig()
	config.Metadata.Name = projectName

	// Set clause version
	config.Metadata.ClauseVersion = "1.0.0" // This should be set from build info

	// Save the configuration
	saver := NewSaver(WithBackup(false))
	if err := saver.SaveToProject(config, projectDir); err != nil {
		return nil, fmt.Errorf("failed to initialize project config: %w", err)
	}

	return config, nil
}

// InitGlobalConfig creates the global configuration file.
func InitGlobalConfig() error {
	home := utils.GetHomeDirectory()

	configPath := filepath.Join(home, ".clause", "config.yaml")

	// Check if config already exists
	if utils.FileExists(configPath) {
		return nil // Already exists, don't overwrite
	}

	// Create default global config
	config := NewProjectConfig()

	// Save without backup since it's new
	saver := NewSaver(WithBackup(false))
	return saver.Save(config, configPath)
}

// Export exports the configuration to a different format.
func (s *Saver) Export(config *ProjectConfig, path string, format string) error {
	tempSaver := NewSaver(WithFormat(format), WithBackup(false))
	return tempSaver.Save(config, path)
}

// ExportJSON exports the configuration as JSON.
func ExportJSON(config *ProjectConfig, path string) error {
	return NewSaver(WithFormat("json"), WithBackup(false)).Save(config, path)
}

// ExportYAML exports the configuration as YAML.
func ExportYAML(config *ProjectConfig, path string) error {
	return NewSaver(WithFormat("yaml"), WithBackup(false)).Save(config, path)
}

// UpdateProjectConfig loads, modifies, and saves a project configuration.
func UpdateProjectConfig(projectDir string, modifier func(*ProjectConfig)) error {
	loader := NewLoader(WithProjectDir(projectDir))
	config, err := loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load project config: %w", err)
	}

	modifier(config)

	saver := NewSaver()
	return saver.SaveToProject(config, projectDir)
}

// MergeConfig merges partial configuration into an existing project configuration.
func MergeConfig(projectDir string, partial map[string]interface{}) error {
	loader := NewLoader(WithProjectDir(projectDir))
	config, err := loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load project config: %w", err)
	}

	if err := mergeMapIntoConfig(config, partial); err != nil {
		return fmt.Errorf("failed to merge config: %w", err)
	}

	saver := NewSaver()
	return saver.SaveToProject(config, projectDir)
}

// SetConfigValue sets a specific configuration value by key path.
// Key paths use dot notation (e.g., "frontend.framework", "backend.database.primary").
func SetConfigValue(projectDir string, keyPath string, value interface{}) error {
	loader := NewLoader(WithProjectDir(projectDir))
	config, err := loader.Load()
	if err != nil {
		return fmt.Errorf("failed to load project config: %w", err)
	}

	if err := setNestedValue(config, keyPath, value); err != nil {
		return fmt.Errorf("failed to set config value: %w", err)
	}

	saver := NewSaver()
	return saver.SaveToProject(config, projectDir)
}

// setNestedValue sets a value in the config using dot notation path.
func setNestedValue(config *ProjectConfig, path string, value interface{}) error {
	parts := strings.Split(path, ".")
	if len(parts) == 0 {
		return fmt.Errorf("empty path")
	}

	// Handle top-level fields
	switch parts[0] {
	case "metadata":
		return setMetadataValue(&config.Metadata, parts[1:], value)
	case "frontend":
		return setFrontendValue(&config.Frontend, parts[1:], value)
	case "backend":
		return setBackendValue(&config.Backend, parts[1:], value)
	case "infrastructure":
		return setInfrastructureValue(&config.Infrastructure, parts[1:], value)
	case "governance":
		return setGovernanceValue(&config.Governance, parts[1:], value)
	case "development":
		return setDevelopmentValue(&config.Development, parts[1:], value)
	default:
		return fmt.Errorf("unknown top-level field: %s", parts[0])
	}
}

func setMetadataValue(m *ProjectMetadata, parts []string, value interface{}) error {
	if len(parts) != 1 {
		return fmt.Errorf("invalid metadata path")
	}

	switch parts[0] {
	case "name":
		if v, ok := value.(string); ok {
			m.Name = v
			return nil
		}
	case "description":
		if v, ok := value.(string); ok {
			m.Description = v
			return nil
		}
	case "version":
		if v, ok := value.(string); ok {
			m.Version = v
			return nil
		}
	case "author":
		if v, ok := value.(string); ok {
			m.Author = v
			return nil
		}
	case "license":
		if v, ok := value.(string); ok {
			m.License = v
			return nil
		}
	case "repository":
		if v, ok := value.(string); ok {
			m.Repository = v
			return nil
		}
	}
	return fmt.Errorf("unknown metadata field: %s", parts[0])
}

func setFrontendValue(f *FrontendConfig, parts []string, value interface{}) error {
	if len(parts) == 0 {
		return fmt.Errorf("empty frontend path")
	}

	if len(parts) == 2 && parts[0] == "features" {
		return setFrontendFeaturesValue(&f.Features, parts[1], value)
	}

	if len(parts) != 1 {
		return fmt.Errorf("invalid frontend path")
	}

	switch parts[0] {
	case "enabled":
		if v, ok := value.(bool); ok {
			f.Enabled = v
			return nil
		}
	case "framework":
		if v, ok := value.(string); ok {
			f.Framework = v
			return nil
		}
	case "typescript":
		if v, ok := value.(bool); ok {
			f.TypeScript = v
			return nil
		}
	case "styling":
		if v, ok := value.(string); ok {
			f.Styling = v
			return nil
		}
	case "package_manager":
		if v, ok := value.(string); ok {
			f.PackageManager = v
			return nil
		}
	case "build_tool":
		if v, ok := value.(string); ok {
			f.BuildTool = v
			return nil
		}
	case "test_framework":
		if v, ok := value.(string); ok {
			f.TestFramework = v
			return nil
		}
	case "linter":
		if v, ok := value.(string); ok {
			f.Linter = v
			return nil
		}
	case "formatter":
		if v, ok := value.(string); ok {
			f.Formatter = v
			return nil
		}
	case "directory":
		if v, ok := value.(string); ok {
			f.Directory = v
			return nil
		}
	}
	return fmt.Errorf("unknown frontend field: %s", parts[0])
}

func setFrontendFeaturesValue(f *FrontendFeatures, field string, value interface{}) error {
	v, ok := value.(bool)
	if !ok {
		return fmt.Errorf("feature value must be boolean")
	}

	switch field {
	case "ssr":
		f.SSR = v
	case "ssg":
		f.SSG = v
	case "pwa":
		f.PWA = v
	case "i18n":
		f.I18n = v
	case "dark_mode":
		f.DarkMode = v
	case "storybook":
		f.Storybook = v
	default:
		return fmt.Errorf("unknown frontend feature: %s", field)
	}
	return nil
}

func setBackendValue(b *BackendConfig, parts []string, value interface{}) error {
	if len(parts) == 0 {
		return fmt.Errorf("empty backend path")
	}

	// Handle nested configs
	if len(parts) >= 2 {
		switch parts[0] {
		case "database":
			return setDatabaseValue(&b.Database, parts[1:], value)
		case "auth":
			return setAuthValue(&b.Auth, parts[1:], value)
		case "api":
			return setAPIValue(&b.API, parts[1:], value)
		case "features":
			return setBackendFeaturesValue(&b.Features, parts[1], value)
		}
	}

	if len(parts) != 1 {
		return fmt.Errorf("invalid backend path")
	}

	switch parts[0] {
	case "enabled":
		if v, ok := value.(bool); ok {
			b.Enabled = v
			return nil
		}
	case "framework":
		if v, ok := value.(string); ok {
			b.Framework = v
			return nil
		}
	case "language":
		if v, ok := value.(string); ok {
			b.Language = v
			return nil
		}
	case "directory":
		if v, ok := value.(string); ok {
			b.Directory = v
			return nil
		}
	}
	return fmt.Errorf("unknown backend field: %s", parts[0])
}

func setDatabaseValue(d *DatabaseConfig, parts []string, value interface{}) error {
	if len(parts) != 1 {
		return fmt.Errorf("invalid database path")
	}

	switch parts[0] {
	case "primary":
		if v, ok := value.(string); ok {
			d.Primary = v
			return nil
		}
	case "orm":
		if v, ok := value.(string); ok {
			d.ORM = v
			return nil
		}
	case "migrations":
		if v, ok := value.(bool); ok {
			d.Migrations = v
			return nil
		}
	case "redis":
		if v, ok := value.(bool); ok {
			d.Redis = v
			return nil
		}
	}
	return fmt.Errorf("unknown database field: %s", parts[0])
}

func setAuthValue(a *AuthConfig, parts []string, value interface{}) error {
	if len(parts) != 1 {
		return fmt.Errorf("invalid auth path")
	}

	switch parts[0] {
	case "provider":
		if v, ok := value.(string); ok {
			a.Provider = v
			return nil
		}
	case "session_duration":
		if v, ok := value.(int); ok {
			a.SessionDuration = v
			return nil
		}
	}
	return fmt.Errorf("unknown auth field: %s", parts[0])
}

func setAPIValue(a *APIConfig, parts []string, value interface{}) error {
	if len(parts) != 1 {
		return fmt.Errorf("invalid api path")
	}

	switch parts[0] {
	case "style":
		if v, ok := value.(string); ok {
			a.Style = v
			return nil
		}
	case "versioning":
		if v, ok := value.(string); ok {
			a.Versioning = v
			return nil
		}
	case "documentation":
		if v, ok := value.(bool); ok {
			a.Documentation = v
			return nil
		}
	}
	return fmt.Errorf("unknown api field: %s", parts[0])
}

func setBackendFeaturesValue(f *BackendFeatures, field string, value interface{}) error {
	v, ok := value.(bool)
	if !ok {
		return fmt.Errorf("feature value must be boolean")
	}

	switch field {
	case "websocket":
		f.WebSocket = v
	case "background_jobs":
		f.BackgroundJobs = v
	case "file_upload":
		f.FileUpload = v
	case "email":
		f.Email = v
	case "rate_limiting":
		f.RateLimiting = v
	case "logging":
		f.Logging = v
	case "metrics":
		f.Metrics = v
	default:
		return fmt.Errorf("unknown backend feature: %s", field)
	}
	return nil
}

func setInfrastructureValue(i *InfrastructureConfig, parts []string, value interface{}) error {
	if len(parts) == 0 {
		return fmt.Errorf("empty infrastructure path")
	}

	if len(parts) != 1 {
		return fmt.Errorf("invalid infrastructure path")
	}

	switch parts[0] {
	case "docker":
		if v, ok := value.(bool); ok {
			i.Docker = v
			return nil
		}
	case "docker_compose":
		if v, ok := value.(bool); ok {
			i.DockerCompose = v
			return nil
		}
	case "kubernetes":
		if v, ok := value.(bool); ok {
			i.Kubernetes = v
			return nil
		}
	case "ci":
		if v, ok := value.(string); ok {
			i.CI = v
			return nil
		}
	case "hosting":
		if v, ok := value.(string); ok {
			i.Hosting = v
			return nil
		}
	case "cdn":
		if v, ok := value.(bool); ok {
			i.CDN = v
			return nil
		}
	}
	return fmt.Errorf("unknown infrastructure field: %s", parts[0])
}

func setGovernanceValue(g *GovernanceConfig, parts []string, value interface{}) error {
	if len(parts) == 0 {
		return fmt.Errorf("empty governance path")
	}

	if len(parts) != 1 {
		return fmt.Errorf("invalid governance path")
	}

	switch parts[0] {
	case "enabled":
		if v, ok := value.(bool); ok {
			g.Enabled = v
			return nil
		}
	case "context_level":
		if v, ok := value.(string); ok {
			g.ContextLevel = v
			return nil
		}
	case "component_registry":
		if v, ok := value.(bool); ok {
			g.ComponentRegistry = v
			return nil
		}
	case "brainstorm_md":
		if v, ok := value.(bool); ok {
			g.BrainstormMd = v
			return nil
		}
	case "prompt_guidelines":
		if v, ok := value.(bool); ok {
			g.PromptGuidelines = v
			return nil
		}
	}
	return fmt.Errorf("unknown governance field: %s", parts[0])
}

func setDevelopmentValue(d *DevelopmentConfig, parts []string, value interface{}) error {
	if len(parts) == 0 {
		return fmt.Errorf("empty development path")
	}

	if len(parts) != 1 {
		return fmt.Errorf("invalid development path")
	}

	switch parts[0] {
	case "git":
		if v, ok := value.(bool); ok {
			d.Git = v
			return nil
		}
	}
	return fmt.Errorf("unknown development field: %s", parts[0])
}
