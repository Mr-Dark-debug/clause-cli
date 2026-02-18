// Package cmd provides the CLI commands for Clause.
package cmd

import (
	"fmt"
	"os"

	"github.com/clause-cli/clause/internal/wizard"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Build information (set via ldflags during build).
var (
	version   = "dev"
	buildTime = "unknown"
	commit    = "unknown"
)

// Global flags.
var (
	cfgFile string
	verbose bool
	quiet   bool
	noColor bool
)

// rootCmd represents the base command when called without any subcommands.
var rootCmd = &cobra.Command{
	Use:   "clause",
	Short: "AI-native project scaffolding tool",
	Long: `Clause (Framework for Organized, Reproducible, and Guided Engineering)
is a cross-platform CLI tool for creating AI-ready project structures.`,
	SilenceUsage:      true,
	SilenceErrors:     true,
	PersistentPreRunE: preRun,
	Run: func(cmd *cobra.Command, args []string) {
		// If no arguments, launch interactive dashboard
		if len(args) == 0 {
			if err := wizard.StartDashboard(cmd, version); err != nil {
				fmt.Fprintf(os.Stderr, "Error launching dashboard: %v\n", err)
				os.Exit(1)
			}
			return
		}
	},
}

func init() {
	cobra.OnInitialize(initConfig)

	// Set custom help function
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		renderer := tui.NewRenderer(nil, 0, 0)
		if width, _, err := styles.GetTerminalSize(); err == nil {
			renderer.SetSize(width, 0)
		}

		fmt.Println(renderer.WelcomeScreen(cmd, version))
	})

	// Persistent flags (available to all subcommands)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.clause/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress non-essential output")
	rootCmd.PersistentFlags().BoolVar(&noColor, "no-color", false, "disable colored output")

	// Bind flags to viper
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
	viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet"))
	viper.BindPFlag("no_color", rootCmd.PersistentFlags().Lookup("no-color"))
}

func preRun(cmd *cobra.Command, args []string) error {
	// Handle color settings
	if noColor || viper.GetBool("no_color") {
		// Disable colors in output
		os.Setenv("NO_COLOR", "1")
	}

	return nil
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// Search for config in home directory
		viper.AddConfigPath(home + "/.clause")
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	// Read in environment variables that match
	viper.SetEnvPrefix("CLAUSE")
	viper.AutomaticEnv()

	// If a config file is found, read it in
	viper.ReadInConfig()
}

// GetVersion returns the current version.
func GetVersion() string {
	return version
}

// GetBuildTime returns the build time.
func GetBuildTime() string {
	return buildTime
}

// GetCommit returns the git commit.
func GetCommit() string {
	return commit
}

// IsVerbose returns true if verbose mode is enabled.
func IsVerbose() bool {
	return verbose || viper.GetBool("verbose")
}

// IsQuiet returns true if quiet mode is enabled.
func IsQuiet() bool {
	return quiet || viper.GetBool("quiet")
}

// ExecuteWithError adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
// Returns the exit code (0 for success, non-zero for error).
func ExecuteWithError() int {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		return 1
	}
	return 0
}
