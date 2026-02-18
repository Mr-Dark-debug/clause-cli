package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/internal/generator"
	"github.com/clause-cli/clause/internal/governance"
	"github.com/clause-cli/clause/internal/wizard"
	"github.com/clause-cli/clause/pkg/output"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/spf13/cobra"
)

// initCmd represents the init command.
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new AI-ready project",
	Long: `Initialize a new project with AI governance built in from day one.

This command creates a complete project structure including:
- Frontend and backend scaffolding based on your choices
- AI governance files (.clause/ directory)
- Component registry
- Brainstorm.md for AI self-reflection
- Documentation standards

If run without a project name, an interactive wizard will guide you
through all configuration options.

Examples:
  clause init                    # Launch interactive wizard
  clause init my-project         # Create project with default settings
  clause init my-project --preset saas  # Use a preset`,
	Args: cobra.MaximumNArgs(1),
	RunE: runInit,
}

var (
	initNonInteractive bool
	initPreset         string
	initDryRun         bool
	initPath           string
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolVarP(&initNonInteractive, "non-interactive", "n", false, "non-interactive mode (use defaults)")
	initCmd.Flags().StringVarP(&initPreset, "preset", "p", "", "use a preset configuration (minimal, standard, saas, api-only, frontend-only, enterprise)")
	initCmd.Flags().BoolVar(&initDryRun, "dry-run", false, "show what would be created without creating files")
	initCmd.Flags().StringVar(&initPath, "path", "", "project creation path (default: current directory)")
}

func runInit(cmd *cobra.Command, args []string) error {
	// Get project name from args
	projectName := ""
	if len(args) > 0 {
		projectName = args[0]
	}

	// Determine mode
	if initNonInteractive || initPreset != "" {
		return runNonInteractiveInit(projectName)
	}

	return runInteractiveInit(projectName)
}

func runInteractiveInit(projectName string) error {
	theme := styles.GetTheme()
	printer := output.NewPrinter(nil, os.Stderr)

	// Print welcome banner
	printBanner(printer, theme)

	// Create and run the wizard
	w := wizard.New(
		wizard.WithProjectName(projectName),
		wizard.WithTheme(theme),
	)

	p := tea.NewProgram(w, tea.WithAltScreen())

	finalModel, err := p.Run()
	if err != nil {
		return fmt.Errorf("wizard failed: %w", err)
	}

	// Get the final configuration
	wiz, ok := finalModel.(*wizard.Wizard)
	if !ok {
		return fmt.Errorf("unexpected model type")
	}

	// Check if user cancelled
	if wiz.IsQuitting() {
		printer.PrintInfo("Project creation cancelled")
		return nil
	}

	// Get the configuration
	cfg := wiz.Config()
	if cfg == nil {
		return fmt.Errorf("no configuration available")
	}

	// Determine project path
	projectPath := initPath
	if projectPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %w", err)
		}
		projectPath = filepath.Join(cwd, cfg.Metadata.Name)
	}

	// Generate the project
	return generateProject(cfg, projectPath, printer)
}

func runNonInteractiveInit(projectName string) error {
	printer := output.NewPrinter(nil, os.Stderr)

	// Load preset or create default config
	var cfg *config.ProjectConfig
	var err error

	if initPreset != "" {
		cfg, err = config.LoadPreset(initPreset)
		if err != nil {
			return fmt.Errorf("failed to load preset: %w", err)
		}
	} else {
		cfg = config.DefaultConfig()
	}

	// Override project name if provided
	if projectName != "" {
		cfg.Metadata.Name = projectName
	}

	// Validate project name
	if cfg.Metadata.Name == "" {
		return fmt.Errorf("project name is required")
	}

	// Determine project path
	projectPath := initPath
	if projectPath == "" {
		cwd, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("failed to get current directory: %w", err)
		}
		projectPath = filepath.Join(cwd, cfg.Metadata.Name)
	}

	// Generate the project
	return generateProject(cfg, projectPath, printer)
}

func generateProject(cfg *config.ProjectConfig, projectPath string, printer *output.Printer) error {
	printer.PrintInfo("Creating project: %s", cfg.Metadata.Name)
	printer.PrintDim("Location: %s", projectPath)

	if initDryRun {
		printer.PrintWarning("Dry run mode - no files will be created")
	}

	// Create the generator
	gen := generator.NewGenerator(cfg,
		generator.WithDryRun(initDryRun),
		generator.WithVerbose(IsVerbose()),
		generator.WithLogger(output.DefaultLogger),
		generator.WithProgress(func(message string) {
			if !IsQuiet() {
				printer.PrintDim("  %s", message)
			}
		}),
	)

	// Generate project files
	if err := gen.Generate(projectPath); err != nil {
		return fmt.Errorf("failed to generate project: %w", err)
	}

	// Initialize governance
	if cfg.Governance.Enabled && !initDryRun {
		gov := governance.New(projectPath, governance.WithConfig(cfg))
		if err := gov.Initialize(); err != nil {
			printer.PrintWarning("Failed to initialize governance: %v", err)
		}
	}

	// Print success message
	printer.Println()
	printer.PrintSuccess("Project created successfully!")
	printer.Println()

	// Print next steps
	printNextSteps(printer, cfg, projectPath)

	return nil
}

func printBanner(printer *output.Printer, theme *styles.Theme) {
	printer.Println()
	banner := `
   ██████╗██╗      █████╗ ██╗   ██╗██████╗ ███████╗
  ██╔════╝██║     ██╔══██╗██║   ██║██╔══██╗██╔════╝
  ██║     ██║     ███████║██║   ██║██║  ██║█████╗
  ██║     ██║     ██╔══██║██║   ██║██║  ██║██╔══╝
  ╚██████╗███████╗██║  ██║╚██████╔╝██████╔╝███████╗
   ╚═════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚═════╝ ╚══════╝
`
	printer.PrintPrimary(banner)
	printer.PrintMuted("  AI-Native Project Scaffolding")
	printer.Println()
}

func printNextSteps(printer *output.Printer, cfg *config.ProjectConfig, projectPath string) {
	printer.PrintInfo("Next steps:")
	printer.Println()

	// Change directory instruction
	printer.PrintMuted("  1. Navigate to your project:")
	printer.Printf("     cd %s\n", projectPath)
	printer.Println()

	// Install dependencies based on configuration
	step := 2
	if cfg.Frontend.Enabled {
		printer.PrintMuted("  %d. Install frontend dependencies:", step)
		pm := cfg.Frontend.PackageManager
		if pm == "" {
			pm = "npm"
		}
		switch pm {
		case "npm":
			printer.Printf("     npm install\n")
		case "yarn":
			printer.Printf("     yarn install\n")
		case "pnpm":
			printer.Printf("     pnpm install\n")
		case "bun":
			printer.Printf("     bun install\n")
		}
		step++
		printer.Println()
	}

	if cfg.Backend.Enabled {
		lang := cfg.Backend.Language
		if lang == "" {
			lang = "node"
		}
		printer.PrintMuted("  %d. Install backend dependencies:", step)
		switch lang {
		case "python":
			printer.Printf("     pip install -r backend/requirements.txt\n")
		case "node", "typescript":
			printer.Printf("     cd backend && npm install\n")
		case "go":
			printer.Printf("     cd backend && go mod download\n")
		}
		step++
		printer.Println()
	}

	// Start development
	printer.PrintMuted("  %d. Start developing:", step)
	if cfg.Frontend.Enabled {
		printer.Printf("     npm run dev")
		if cfg.Backend.Enabled {
			printer.Printf("  (frontend)")
		}
		printer.Println()
	}
	if cfg.Backend.Enabled {
		lang := cfg.Backend.Language
		if lang == "" {
			lang = "node"
		}
		switch lang {
		case "python":
			printer.Printf("     python backend/main.py")
		case "node", "typescript":
			printer.Printf("     npm run dev --prefix backend")
		case "go":
			printer.Printf("     go run backend/main.go")
		}
		printer.Println()
	}
	printer.Println()

	// Governance reminder
	if cfg.Governance.Enabled {
		printer.PrintMuted("  Governance files are in .clause/")
		printer.Println()
	}

	// Happy coding
	printer.PrintSuccess("Happy coding!")
}
