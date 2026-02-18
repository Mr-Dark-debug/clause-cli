package cmd

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd represents the config command.
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage Clause configuration",
	Long: `Manage global and project-level Clause configuration.

Configuration is stored in YAML format and can be set at two levels:
- Global: ~/.clause/config.yaml
- Project: .clause/config.yaml (in your project directory)

Project configuration overrides global configuration.

Examples:
  clause config list              # Show all configuration
  clause config get <key>         # Get a specific value
  clause config set <key> <value> # Set a value
  clause config init              # Initialize configuration`,
}

var (
	configGlobal bool
	configLocal  bool
)

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().BoolVar(&configGlobal, "global", false, "operate on global config")
	configCmd.PersistentFlags().BoolVar(&configLocal, "local", false, "operate on project config")

	// Add subcommands
	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configInitCmd)
}

// configListCmd lists all configuration.
var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all configuration values",
	Run:   runConfigList,
}

func runConfigList(cmd *cobra.Command, args []string) {
	theme := styles.GetTheme()

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(theme.Colors.Primary))

	keyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.TextMuted)).
		Width(25)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Text))

	fmt.Println()
	fmt.Println(titleStyle.Render("Clause Configuration"))
	fmt.Println()

	// Show all settings
	settings := []struct {
		key, defaultValue string
	}{
		{"verbose", "false"},
		{"quiet", "false"},
		{"no_color", "false"},
		{"defaults.frontend", "react"},
		{"defaults.backend", "fastapi"},
		{"defaults.styling", "tailwind"},
		{"defaults.database", "postgresql"},
		{"defaults.ai_context", "comprehensive"},
		{"templates.registry", "https://registry.clause.dev"},
		{"updates.check_enabled", "true"},
		{"updates.channel", "stable"},
	}

	for _, s := range settings {
		value := viper.GetString(s.key)
		if value == "" {
			value = s.defaultValue
		}
		fmt.Printf("  %s%s\n",
			keyStyle.Render(s.key),
			valueStyle.Render(value),
		)
	}

	fmt.Println()
	fmt.Println("Config file:", viper.ConfigFileUsed())
	fmt.Println()
}

// configGetCmd gets a configuration value.
var configGetCmd = &cobra.Command{
	Use:   "get <key>",
	Short: "Get a configuration value",
	Args:  cobra.ExactArgs(1),
	Run:   runConfigGet,
}

func runConfigGet(cmd *cobra.Command, args []string) {
	key := args[0]
	value := viper.Get(key)

	if value == nil {
		fmt.Println("(not set)")
		return
	}

	fmt.Println(value)
}

// configSetCmd sets a configuration value.
var configSetCmd = &cobra.Command{
	Use:   "set <key> <value>",
	Short: "Set a configuration value",
	Args:  cobra.ExactArgs(2),
	Run:   runConfigSet,
}

func runConfigSet(cmd *cobra.Command, args []string) {
	key := args[0]
	value := args[1]

	// Set the value
	viper.Set(key, value)

	// Determine config file
	configFile := cfgFile
	if configFile == "" {
		home, _ := homeDir()
		configFile = home + "/.clause/config.yaml"
	}

	// Write config
	if err := viper.WriteConfigAs(configFile); err != nil {
		fmt.Printf("Error writing config: %v\n", err)
		return
	}

	fmt.Printf("Set %s = %s\n", key, value)
}

// configInitCmd initializes configuration.
var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize configuration with defaults",
	Run:   runConfigInit,
}

func runConfigInit(cmd *cobra.Command, args []string) {
	theme := styles.GetTheme()

	// Get config file path
	home, err := homeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		return
	}

	configFile := home + "/.clause/config.yaml"

	// Check if config already exists
	if err := viper.SafeWriteConfigAs(configFile); err != nil {
		successStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Colors.Success))
		fmt.Println(successStyle.Render("✓ Configuration already exists at: " + configFile))
		return
	}

	successStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Success))
	fmt.Println(successStyle.Render("✓ Configuration initialized at: " + configFile))
}

func homeDir() (string, error) {
	home, err := viper.GetString("home"), error(nil)
	if home == "" {
		home, err = os.UserHomeDir()
	}
	return home, err
}
