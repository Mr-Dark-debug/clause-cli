package cmd

import (
	"fmt"
	"runtime"

	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command.
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Display the version, build time, commit, and Go version for Clause.`,
	Run:   runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(cmd *cobra.Command, args []string) {
	theme := styles.GetTheme()

	// Create styled box
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(theme.Colors.Primary)).
		Padding(1, 2).
		Margin(1, 0)

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(theme.Colors.Primary)).
		MarginBottom(1)

	labelStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.TextMuted)).
		Width(12)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Text))

	// Build version info
	fmt.Println()
	fmt.Println(titleStyle.Render("Clause - AI-Native Project Scaffolding"))
	fmt.Println()

	// Version details
	details := []struct {
		label, value string
	}{
		{"Version:", GetVersion()},
		{"Build Time:", GetBuildTime()},
		{"Commit:", GetCommit()},
		{"Go Version:", runtime.Version()},
		{"Platform:", runtime.GOOS + "/" + runtime.GOARCH},
	}

	for _, d := range details {
		fmt.Printf("  %s%s\n",
			labelStyle.Render(d.label),
			valueStyle.Render(d.value),
		)
	}

	fmt.Println()

	// Check for updates (in a real implementation)
	// updateAvailable := checkForUpdates()
	// if updateAvailable {
	//     fmt.Println(styles.Warning("  A new version is available! Run 'clause update' to upgrade."))
	// }

	// Show quick tip
	tipStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.TextMuted)).
		Italic(true)

	fmt.Println(boxStyle.Render(
		tipStyle.Render("Run 'clause init' to create a new project"),
	))
	fmt.Println()
}
