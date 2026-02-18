package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// initCmd represents the init command.
var initCmd = &cobra.Command{
	Use:   "init [project-name]",
	Short: "Initialize a new AI-ready project",
	Long: `Initialize a new project with AI governance built in from day one.

This command creates a complete project structure including:
- Frontend and backend scaffolding based on your choices
- AI governance files (ai_prompt_guidelines/)
- Component registry
- Brainstorm.md for AI self-reflection
- Documentation standards

If run without a project name, an interactive wizard will guide you
through all configuration options.

Examples:
  clause init                    # Launch interactive wizard
  clause init my-project         # Create project with default settings
  clause init my-project --template saas-starter  # Use a template`,
	Args: cobra.MaximumNArgs(1),
	RunE: runInit,
}

var (
	initNonInteractive bool
	initTemplate       string
	initDefaults       bool
)

func init() {
	rootCmd.AddCommand(initCmd)

	initCmd.Flags().BoolVarP(&initNonInteractive, "non-interactive", "n", false, "non-interactive mode (use defaults)")
	initCmd.Flags().StringVarP(&initTemplate, "template", "t", "", "use a specific template")
	initCmd.Flags().BoolVarP(&initDefaults, "defaults", "d", false, "use all default values")
}

func runInit(cmd *cobra.Command, args []string) error {
	// Get project name from args or prompt
	projectName := ""
	if len(args) > 0 {
		projectName = args[0]
	}

	// Determine mode
	if initNonInteractive || initDefaults {
		return runNonInteractiveInit(projectName)
	}

	return runInteractiveInit(projectName)
}

func runInteractiveInit(projectName string) error {
	// TODO: Launch interactive wizard
	// For now, show a placeholder message
	fmt.Println("🚀 Clause - AI-Native Project Scaffolding")
	fmt.Println()
	fmt.Println("Interactive wizard coming soon!")
	fmt.Println()
	fmt.Println("For now, use --non-interactive to create a project with defaults:")
	fmt.Println("  clause init my-project --non-interactive")
	fmt.Println()

	return nil
}

func runNonInteractiveInit(projectName string) error {
	// TODO: Implement non-interactive project generation
	// For now, show a placeholder message
	fmt.Println("Creating project:", projectName)
	fmt.Println()
	fmt.Println("Non-interactive mode is not yet fully implemented.")
	fmt.Println()

	return nil
}
