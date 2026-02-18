package wizard

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/tui"
	"github.com/spf13/cobra"
)

// Dashboard is the main interactive menu for Clause.
type Dashboard struct {
	renderer *tui.Renderer
	rootCmd  *cobra.Command
	version  string
	cursor   int
	width    int
	height   int
	choices  []MenuChoice
	quitting bool
}

// MenuChoice represents a selectable item in the dashboard.
type MenuChoice struct {
	label       string
	description string
	command     string
}

// NewDashboard creates a new interactive dashboard.
func NewDashboard(rootCmd *cobra.Command, version string) *Dashboard {
	renderer := tui.NewRenderer(nil, 0, 0)

	return &Dashboard{
		renderer: renderer,
		rootCmd:  rootCmd,
		version:  version,
		choices: []MenuChoice{
			{"Initialize", "Start a new AI-ready project", "init"},
			{"Add Component", "Add features to existing project", "add"},
			{"Validate", "Check governance compliance", "validate"},
			{"Configuration", "Manage Clause settings", "config"},
			{"Update", "Check for CLI updates", "update"},
			{"Help", "View detailed command reference", "help"},
			{"Exit", "Quit the Clause CLI", "exit"},
		},
	}
}

// Init initializes the dashboard.
func (d *Dashboard) Init() tea.Cmd {
	return nil
}

// Update handles interactive messages.
func (d *Dashboard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch m := msg.(type) {
	case tea.WindowSizeMsg:
		d.width = m.Width
		d.height = m.Height
		d.renderer.SetSize(m.Width, m.Height)

	case tea.KeyMsg:
		switch m.String() {
		case "up", "k":
			if d.cursor > 0 {
				d.cursor--
			}
		case "down", "j":
			if d.cursor < len(d.choices)-1 {
				d.cursor++
			}
		case "enter":
			choice := d.choices[d.cursor]
			if choice.command == "exit" {
				d.quitting = true
				return d, tea.Quit
			}

			// For "Initialize", we can transition directly to the wizard
			if choice.command == "init" {
				w := New()
				return w, w.Init()
			}

			// For others, we quit and explain how to run
			d.quitting = true
			return d, tea.Quit
		case "q", "ctrl+c", "esc":
			d.quitting = true
			return d, tea.Quit
		}
	}
	return d, nil
}

// View renders the interactive dashboard.
func (d *Dashboard) View() string {
	if d.quitting {
		if d.cursor < len(d.choices) && d.choices[d.cursor].command != "exit" {
			return d.renderer.Success(fmt.Sprintf("\nSelected: %s\nRun 'clause %s' or explore it in the interactive wizard soon!\n",
				d.choices[d.cursor].label, d.choices[d.cursor].command))
		}
		return "\n  See you later!\n"
	}

	banner := d.renderer.Banner(d.version)

	description := `Clause generates complete project scaffolding with built-in AI
governance, context files, and best practices for AI assistants.`

	features := []string{
		"Complete project structure generation",
		"AI governance guidelines included",
		"Context files for AI assistants",
		"Support for Next.js, FastAPI, Go, and more",
	}

	var featureLines []string
	for _, f := range features {
		featureLines = append(featureLines, d.renderer.Theme().Typography.Primary.Render(" • ")+d.renderer.Theme().Typography.Body.Render(f))
	}

	titleDesc := d.renderer.Theme().Typography.Header.Render("✨ Create AI-Ready Projects")
	descContent := lipgloss.JoinVertical(
		lipgloss.Left,
		d.renderer.Theme().Typography.Body.Render(description),
		"",
		strings.Join(featureLines, "\n"),
	)

	descCard := d.renderer.Theme().Layout.Card.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(d.renderer.Theme().Colors.Border)).
		Padding(1, 2).
		Width(lipgloss.Width(banner)).
		Render(descContent)

	// Create the menu list
	var menuItems []string
	for i, choice := range d.choices {
		cursor := "  "
		if d.cursor == i {
			cursor = d.renderer.Theme().Typography.Success.Render("▸ ")
		}

		label := d.renderer.Theme().Typography.Body.Copy().Bold(true).Render(choice.label)
		if d.cursor == i {
			label = d.renderer.Theme().Typography.Primary.Copy().Bold(true).Render(choice.label)
		}

		desc := d.renderer.Theme().Typography.Muted.Render(" - " + choice.description)
		menuItems = append(menuItems, cursor+label+desc)
	}

	menuTitle := d.renderer.Theme().Typography.Header.Render("🚀 Main Menu")
	menuCard := d.renderer.Theme().Layout.Card.
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(d.renderer.Theme().Colors.Border)).
		Padding(1, 2).
		Width(lipgloss.Width(banner)).
		Render(tui.JoinVertical(menuItems...))

	help := d.renderer.Theme().Typography.Muted.Render("\n  Use ↑/↓ to navigate • Enter to select • q to quit")

	return tui.JoinVertical(
		banner,
		"",
		lipgloss.JoinVertical(lipgloss.Left, titleDesc, descCard),
		"",
		lipgloss.JoinVertical(lipgloss.Left, menuTitle, menuCard),
		help,
	)
}

// StartDashboard launches the interactive dashboard.
func StartDashboard(rootCmd *cobra.Command, version string) error {
	p := tea.NewProgram(NewDashboard(rootCmd, version), tea.WithAltScreen())
	_, err := p.Run()
	return err
}
