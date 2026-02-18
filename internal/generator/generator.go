package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/internal/template"
	"github.com/clause-cli/clause/pkg/output"
	"github.com/clause-cli/clause/pkg/utils"
)

// Generator creates project scaffolding.
type Generator struct {
	// Config is the project configuration
	Config *config.ProjectConfig

	// TemplateEngine is the template renderer
	TemplateEngine *template.Engine

	// DryRun indicates if files should be created
	DryRun bool

	// Verbose enables verbose output
	Verbose bool

	// Logger for output
	Logger *output.Logger

	// Progress callback
	OnProgress func(message string)
}

// GeneratorOption is a functional option for configuring the generator.
type GeneratorOption func(*Generator)

// NewGenerator creates a new generator.
func NewGenerator(cfg *config.ProjectConfig, opts ...GeneratorOption) *Generator {
	g := &Generator{
		Config:         cfg,
		TemplateEngine: template.NewEngine(),
		Logger:         output.DefaultLogger,
	}

	for _, opt := range opts {
		opt(g)
	}

	return g
}

// WithDryRun sets dry run mode.
func WithDryRun(dryRun bool) GeneratorOption {
	return func(g *Generator) {
		g.DryRun = dryRun
	}
}

// WithVerbose sets verbose mode.
func WithVerbose(verbose bool) GeneratorOption {
	return func(g *Generator) {
		g.Verbose = verbose
	}
}

// WithLogger sets the logger.
func WithLogger(logger *output.Logger) GeneratorOption {
	return func(g *Generator) {
		g.Logger = logger
	}
}

// WithProgress sets the progress callback.
func WithProgress(callback func(message string)) GeneratorOption {
	return func(g *Generator) {
		g.OnProgress = callback
	}
}

// Generate generates the project at the specified path.
func (g *Generator) Generate(projectPath string) error {
	g.progress("Creating project directory structure...")

	// Validate configuration
	if err := g.validateConfig(); err != nil {
		return fmt.Errorf("invalid configuration: %w", err)
	}

	// Create root directory
	if err := g.createDirectory(projectPath); err != nil {
		return fmt.Errorf("failed to create project directory: %w", err)
	}

	// Create .clause directory with config
	g.progress("Creating Clause configuration...")
	if err := g.createClauseConfig(projectPath); err != nil {
		return err
	}

	// Create common files
	g.progress("Creating common files...")
	if err := g.createCommonFiles(projectPath); err != nil {
		return err
	}

	// Create frontend if enabled
	if g.Config.Frontend.Enabled {
		g.progress("Creating frontend structure...")
		if err := g.createFrontend(projectPath); err != nil {
			return err
		}
	}

	// Create backend if enabled
	if g.Config.Backend.Enabled {
		g.progress("Creating backend structure...")
		if err := g.createBackend(projectPath); err != nil {
			return err
		}
	}

	// Create infrastructure files
	g.progress("Creating infrastructure files...")
	if err := g.createInfrastructure(projectPath); err != nil {
		return err
	}

	// Create governance files
	if g.Config.Governance.Enabled {
		g.progress("Creating governance files...")
		if err := g.createGovernance(projectPath); err != nil {
			return err
		}
	}

	// Initialize git if enabled
	if g.Config.Development.Git {
		g.progress("Initializing git repository...")
		if err := g.initGit(projectPath); err != nil {
			g.Logger.Warn("Failed to initialize git: %v", err)
		}
	}

	g.progress("Project generation complete!")
	return nil
}

// validateConfig validates the configuration before generation.
func (g *Generator) validateConfig() error {
	errors := config.Validate(g.Config)
	if errors.HasErrors() {
		return fmt.Errorf("configuration has errors: %v", errors)
	}
	return nil
}

// createDirectory creates a directory.
func (g *Generator) createDirectory(path string) error {
	if g.DryRun {
		g.Logger.Info("[DRY RUN] Would create directory: %s", path)
		return nil
	}
	return utils.EnsureDirectory(path)
}

// writeFile writes a file with content.
func (g *Generator) writeFile(path, content string) error {
	if g.DryRun {
		g.Logger.Info("[DRY RUN] Would create file: %s", path)
		return nil
	}

	// Ensure directory exists
	dir := filepath.Dir(path)
	if err := utils.EnsureDirectory(dir); err != nil {
		return err
	}

	return os.WriteFile(path, []byte(content), 0644)
}

// writeTemplate writes a templated file.
func (g *Generator) writeTemplate(path, tmpl string) error {
	data := template.NewTemplateData(g.Config)
	content, err := g.TemplateEngine.Render(tmpl, data)
	if err != nil {
		return fmt.Errorf("failed to render template for %s: %w", path, err)
	}
	return g.writeFile(path, content)
}

// progress reports progress.
func (g *Generator) progress(message string) {
	if g.OnProgress != nil {
		g.OnProgress(message)
	}
	if g.Verbose {
		g.Logger.Info(message)
	}
}

// createClauseConfig creates the .clause configuration directory.
func (g *Generator) createClauseConfig(projectPath string) error {
	clauseDir := filepath.Join(projectPath, ".clause")
	if err := g.createDirectory(clauseDir); err != nil {
		return err
	}

	// Save configuration
	saver := config.NewSaver()
	if !g.DryRun {
		if err := saver.Save(g.Config, filepath.Join(clauseDir, "config.yaml")); err != nil {
			return fmt.Errorf("failed to save config: %w", err)
		}
	}

	return nil
}

// createCommonFiles creates common project files.
func (g *Generator) createCommonFiles(projectPath string) error {
	// Create README.md
	readmeContent := g.generateREADME()
	if err := g.writeFile(filepath.Join(projectPath, "README.md"), readmeContent); err != nil {
		return err
	}

	// Create .gitignore
	gitignoreContent := g.generateGitignore()
	if err := g.writeFile(filepath.Join(projectPath, ".gitignore"), gitignoreContent); err != nil {
		return err
	}

	// Create .editorconfig
	editorconfigContent := g.generateEditorconfig()
	if err := g.writeFile(filepath.Join(projectPath, ".editorconfig"), editorconfigContent); err != nil {
		return err
	}

	return nil
}

// generateREADME generates README.md content.
func (g *Generator) generateREADME() string {
	return fmt.Sprintf(`# %s

%s

## Getting Started

### Prerequisites

`, g.Config.Metadata.Name, g.Config.Metadata.Description)
}

// generateGitignore generates .gitignore content.
func (g *Generator) generateGitignore() string {
	var content strings.Builder

	content.WriteString("# Dependencies\n")
	content.WriteString("node_modules/\n")
	content.WriteString("vendor/\n")
	content.WriteString("__pycache__/\n")
	content.WriteString("*.pyc\n\n")

	content.WriteString("# Build outputs\n")
	content.WriteString("dist/\n")
	content.WriteString("build/\n")
	content.WriteString("*.exe\n\n")

	content.WriteString("# Environment\n")
	content.WriteString(".env\n")
	content.WriteString(".env.local\n\n")

	content.WriteString("# IDE\n")
	content.WriteString(".idea/\n")
	content.WriteString(".vscode/\n")
	content.WriteString("*.swp\n\n")

	content.WriteString("# OS\n")
	content.WriteString(".DS_Store\n")
	content.WriteString("Thumbs.db\n")

	return content.String()
}

// generateEditorconfig generates .editorconfig content.
func (g *Generator) generateEditorconfig() string {
	return `root = true

[*]
charset = utf-8
end_of_line = lf
indent_style = tab
indent_size = 4
insert_final_newline = true
trim_trailing_whitespace = true

[*.md]
trim_trailing_whitespace = false

[*.{yml,yaml}]
indent_style = space
indent_size = 2
`
}
