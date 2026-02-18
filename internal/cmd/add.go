package cmd

import (
	"fmt"

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

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) error {
	componentType := args[0]
	var name string
	if len(args) > 1 {
		name = args[1]
	}

	// TODO: Implement component addition
	fmt.Printf("Adding %s component", componentType)
	if name != "" {
		fmt.Printf(": %s", name)
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Component addition is not yet implemented.")

	return nil
}
