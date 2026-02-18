package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/internal/governance"
	"github.com/clause-cli/clause/pkg/output"
	"github.com/spf13/cobra"
)

// addCmd represents the add command.
var addCmd = &cobra.Command{
	Use:   "add <component-type> [name]",
	Short: "Add a new component to an existing project",
	Long: `Add new components to an existing Clause project.

Component types:
  frontend    Add frontend components (pages, components, utilities)
  backend     Add backend components (routes, models, services)
  governance  Add governance rules and AI context

Examples:
  clause add frontend component Button
  clause add backend route users
  clause add governance rule no-any-type`,
	Args: cobra.MinimumNArgs(1),
	RunE: runAdd,
}

var (
	addDescription string
	addPath        string
	addDeps        []string
	addTags        []string
)

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringVarP(&addDescription, "description", "d", "", "component description")
	addCmd.Flags().StringVarP(&addPath, "path", "p", "", "component path")
	addCmd.Flags().StringSliceVar(&addDeps, "deps", []string{}, "component dependencies")
	addCmd.Flags().StringSliceVar(&addTags, "tags", []string{}, "component tags")
}

func runAdd(cmd *cobra.Command, args []string) error {
	printer := output.NewPrinter(nil, os.Stderr)

	// Find project root (look for .clause directory)
	projectPath, err := findProjectRoot()
	if err != nil {
		return fmt.Errorf("not in a Clause project: %w", err)
	}

	componentType := args[0]
	var name string
	if len(args) > 1 {
		name = args[1]
	}

	// Check if governance is enabled
	cfg, err := loadProjectConfig(projectPath)
	if err != nil {
		return fmt.Errorf("failed to load project config: %w", err)
	}

	if !cfg.Governance.Enabled || !cfg.Governance.ComponentRegistry {
		printer.PrintWarning("Component registry is not enabled for this project")
		printer.PrintInfo("Enable it in .clause/config.yaml to use component tracking")
		return nil
	}

	// Create governance manager
	gov := governance.New(projectPath, governance.WithConfig(cfg))

	// Handle different component types
	switch componentType {
	case "frontend":
		return addFrontendComponent(gov, name, printer)
	case "backend":
		return addBackendComponent(gov, name, printer)
	case "governance":
		return addGovernanceRule(gov, name, printer)
	case "component":
		// Generic component registration
		return addGenericComponent(gov, name, componentType, printer)
	default:
		return fmt.Errorf("unknown component type: %s", componentType)
	}
}

func addFrontendComponent(gov *governance.Governance, name string, printer *output.Printer) error {
	if name == "" {
		return fmt.Errorf("component name is required")
	}

	comp := governance.Component{
		Name:        name,
		Type:        "frontend",
		Path:        addPath,
		Description: addDescription,
		Dependencies: addDeps,
		Tags:        addTags,
		TechStack:   []string{"react", "typescript", "tailwind"},
	}

	if comp.Path == "" {
		comp.Path = fmt.Sprintf("src/components/%s", name)
	}

	if err := gov.RegisterComponent(comp); err != nil {
		return err
	}

	printer.PrintSuccess("Registered frontend component: %s", name)
	printer.PrintDim("  Path: %s", comp.Path)
	if len(comp.Dependencies) > 0 {
		printer.PrintDim("  Dependencies: %v", comp.Dependencies)
	}

	return nil
}

func addBackendComponent(gov *governance.Governance, name string, printer *output.Printer) error {
	if name == "" {
		return fmt.Errorf("component name is required")
	}

	comp := governance.Component{
		Name:        name,
		Type:        "backend",
		Path:        addPath,
		Description: addDescription,
		Dependencies: addDeps,
		Tags:        addTags,
		TechStack:   []string{"python", "fastapi", "sqlalchemy"},
	}

	if comp.Path == "" {
		comp.Path = fmt.Sprintf("backend/services/%s", name)
	}

	if err := gov.RegisterComponent(comp); err != nil {
		return err
	}

	printer.PrintSuccess("Registered backend component: %s", name)
	printer.PrintDim("  Path: %s", comp.Path)
	if len(comp.Dependencies) > 0 {
		printer.PrintDim("  Dependencies: %v", comp.Dependencies)
	}

	return nil
}

func addGovernanceRule(gov *governance.Governance, name string, printer *output.Printer) error {
	if name == "" {
		return fmt.Errorf("rule name is required")
	}

	printer.PrintInfo("Governance rules are defined in .clause/rules/")
	printer.PrintDim("To add a rule, edit the rules configuration file")

	return nil
}

func addGenericComponent(gov *governance.Governance, name, compType string, printer *output.Printer) error {
	if name == "" {
		return fmt.Errorf("component name is required")
	}

	comp := governance.Component{
		Name:        name,
		Type:        compType,
		Path:        addPath,
		Description: addDescription,
		Dependencies: addDeps,
		Tags:        addTags,
	}

	if comp.Path == "" {
		comp.Path = name
	}

	if err := gov.RegisterComponent(comp); err != nil {
		return err
	}

	printer.PrintSuccess("Registered component: %s", name)
	printer.PrintDim("  Type: %s", comp.Type)
	printer.PrintDim("  Path: %s", comp.Path)

	return nil
}

func findProjectRoot() (string, error) {
	// Start from current directory
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Walk up the directory tree looking for .clause
	for {
		clauseDir := filepath.Join(dir, ".clause")
		if _, err := os.Stat(clauseDir); err == nil {
			return dir, nil
		}

		// Move to parent directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root without finding .clause
			return "", fmt.Errorf("could not find .clause directory")
		}
		dir = parent
	}
}

func loadProjectConfig(projectPath string) (*config.ProjectConfig, error) {
	loader := config.NewLoader()
	configPath := filepath.Join(projectPath, ".clause", "config.yaml")

	cfg, err := loader.LoadFromPath(configPath)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
