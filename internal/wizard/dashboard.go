package wizard

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
	"github.com/spf13/cobra"
)

// Dashboard is the main interactive menu for Clause.
type Dashboard struct {
	renderer    *tui.Renderer
	rootCmd     *cobra.Command
	version     string
	cursor      int
	width       int
	height      int
	choices     []MenuChoice
	quitting    bool
	selectedCmd string
	showingHelp bool
}

// MenuChoice represents a selectable item in the dashboard.
type MenuChoice struct {
	label       string
	description string
	command     string
	icon        string
	category    string
}

// NewDashboard creates a new interactive dashboard.
func NewDashboard(rootCmd *cobra.Command, version string) *Dashboard {
	renderer := tui.NewRenderer(nil, 0, 0)

	return &Dashboard{
		renderer: renderer,
		rootCmd:  rootCmd,
		version:  version,
		choices: []MenuChoice{
			{"Initialize", "Start a new AI-ready project", "init", "🚀", "Project"},
			{"Add Component", "Add features to existing project", "add", "📦", "Project"},
			{"Validate", "Check governance compliance", "validate", "✓", "Project"},
			{"Configuration", "Manage Clause settings", "config", "⚙", "Settings"},
			{"Update", "Check for CLI updates", "update", "↑", "Settings"},
			{"Help", "View detailed command reference", "help", "?", "Utility"},
			{"Exit", "Quit the Clause CLI", "exit", "✕", "Utility"},
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
		// If showing help, any key goes back to menu
		if d.showingHelp {
			d.showingHelp = false
			return d, nil
		}

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
			d.selectedCmd = choice.command

			switch choice.command {
			case "exit":
				d.quitting = true
				return d, tea.Quit
			case "init":
				// Transition to wizard
				w := New()
				return w, w.Init()
			case "help":
				d.showingHelp = true
				return d, nil
			default:
				// For other commands, show info and quit
				d.quitting = true
				return d, tea.Quit
			}
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
		return d.renderExitMessage()
	}

	if d.showingHelp {
		return d.renderHelpScreen()
	}

	// Build the UI components
	banner := d.renderBanner()
	descCard := d.renderDescriptionCard()
	menuCard := d.renderMenuCard()
	helpBar := d.renderHelpBar()
	footer := d.renderFooter()

	// Assemble the full UI
	ui := lipgloss.JoinVertical(
		lipgloss.Center,
		banner,
		"",
		descCard,
		"",
		menuCard,
		"",
		helpBar,
		footer,
	)

	// Center the UI in the terminal
	if d.width > 0 && d.height > 0 {
		return lipgloss.Place(d.width, d.height, lipgloss.Center, lipgloss.Center, ui)
	}

	return ui
}

// renderBanner renders the ASCII art banner with gradient effect.
func (d *Dashboard) renderBanner() string {
	theme := d.renderer.Theme()

	logo := `   ██████╗██╗      █████╗ ██╗   ██╗███████╗███████╗
  ██╔════╝██║     ██╔══██╗██║   ██║██╔════╝██╔════╝
  ██║     ██║     ███████║██║   ██║███████╗█████╗
  ██║     ██║     ██╔══██║██║   ██║╚════██║██╔══╝
  ╚██████╗███████╗██║  ██║╚██████╔╝███████║███████╗
   ╚═════╝╚══════╝╚═╝  ╚═╝ ╚═════╝ ╚══════╝╚══════╝`

	// Apply gradient to logo lines
	lines := strings.Split(logo, "\n")
	gradient := styles.Gradient(theme.Colors.Primary, theme.Colors.AccentTertiary, len(lines))
	var styledLogo []string
	for i, line := range lines {
		styledLogo = append(styledLogo, lipgloss.NewStyle().
			Foreground(lipgloss.Color(gradient[i])).
			Render(line))
	}
	logoStr := strings.Join(styledLogo, "\n")

	tagline := "Framework for Organized, Reproducible, and Guided Engineering"
	versionStr := fmt.Sprintf("Version %s", d.version)

	content := lipgloss.JoinVertical(
		lipgloss.Center,
		logoStr,
		"",
		lipgloss.NewStyle().
			Italic(true).
			Foreground(lipgloss.Color(theme.Colors.TextMuted)).
			Render(tagline),
		lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Colors.Primary)).
			Bold(true).
			Render(versionStr),
	)

	// Create banner container with double border
	bannerStyle := lipgloss.NewStyle().
		Border(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color(theme.Colors.Primary)).
		Padding(1, 4).
		Margin(0, 2)

	return bannerStyle.Render(content)
}

// renderDescriptionCard renders the feature description card.
func (d *Dashboard) renderDescriptionCard() string {
	theme := d.renderer.Theme()

	description := "Clause generates complete project scaffolding with built-in AI governance, context files, and best practices for AI assistants."

	features := []struct {
		icon string
		text string
	}{
		{"📁", "Complete project structure generation"},
		{"🤖", "AI governance guidelines included"},
		{"📝", "Context files for AI assistants"},
		{"⚡", "Support for Next.js, FastAPI, Go, and more"},
	}

	var featureLines []string
	for _, f := range features {
		iconStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colors.Primary)).PaddingRight(1)
		textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colors.Text))
		featureLines = append(featureLines,
			lipgloss.JoinHorizontal(lipgloss.Top,
				iconStyle.Render(f.icon),
				textStyle.Render(f.text),
			),
		)
	}

	// Title
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(theme.Colors.Primary)).
		PaddingBottom(1)

	// Card style
	cardWidth := 70
	if d.width > 20 {
		cardWidth = min(d.width-10, 75)
	}
	cardStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(theme.Colors.Border)).
		Padding(1, 2).
		Width(cardWidth)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		lipgloss.NewStyle().Foreground(lipgloss.Color(theme.Colors.Text)).Render(description),
		"",
		strings.Join(featureLines, "\n"),
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("✨ Create AI-Ready Projects"),
		cardStyle.Render(content),
	)
}

// renderMenuCard renders the interactive menu.
func (d *Dashboard) renderMenuCard() string {
	theme := d.renderer.Theme()

	// Group choices by category
	categories := map[string][]MenuChoice{
		"Project":  {},
		"Settings": {},
		"Utility":  {},
	}
	categoryOrder := []string{"Project", "Settings", "Utility"}

	for _, choice := range d.choices {
		if cat, ok := categories[choice.category]; ok {
			categories[choice.category] = append(cat, choice)
		}
	}

	var sections []string

	for _, catName := range categoryOrder {
		choices := categories[catName]
		if len(choices) == 0 {
			continue
		}

		// Category header
		headerStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Colors.TextMuted)).
			Bold(true).
			PaddingBottom(0)

		dividerStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Colors.BorderMuted))

		var menuItems []string
		for _, choice := range choices {
			menuItems = append(menuItems, d.renderMenuItem(choice))
		}

		dividerWidth := 50
		if d.width > 20 {
			dividerWidth = min(d.width-25, 55)
		}

		section := lipgloss.JoinVertical(
			lipgloss.Left,
			headerStyle.Render("  "+strings.ToUpper(catName)),
			dividerStyle.Render("  "+strings.Repeat("─", dividerWidth)),
			strings.Join(menuItems, "\n"),
		)
		sections = append(sections, section)
	}

	// Title
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(theme.Colors.Primary)).
		PaddingBottom(1)

	// Card style
	cardWidth := 65
	if d.width > 20 {
		cardWidth = min(d.width-10, 70)
	}
	cardStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(theme.Colors.Border)).
		Padding(1, 0).
		Width(cardWidth)

	content := strings.Join(sections, "\n\n")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("🎯 Main Menu"),
		cardStyle.Render(content),
	)
}

// renderMenuItem renders a single menu item.
func (d *Dashboard) renderMenuItem(choice MenuChoice) string {
	theme := d.renderer.Theme()
	isSelected := d.choices[d.cursor].command == choice.command

	// Cursor indicator
	var cursor string
	if isSelected {
		cursor = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Colors.Success)).
			Bold(true).
			Render("▸")
	} else {
		cursor = lipgloss.NewStyle().
			Foreground(lipgloss.Color(theme.Colors.BorderMuted)).
			Render(" ")
	}

	// Icon
	iconStyle := lipgloss.NewStyle().Padding(0, 1)
	if isSelected {
		iconStyle = iconStyle.Foreground(lipgloss.Color(theme.Colors.Primary))
	} else {
		iconStyle = iconStyle.Foreground(lipgloss.Color(theme.Colors.TextMuted))
	}

	// Label
	labelStyle := lipgloss.NewStyle()
	if isSelected {
		labelStyle = labelStyle.
			Foreground(lipgloss.Color(theme.Colors.Text)).
			Bold(true)
	} else {
		labelStyle = labelStyle.Foreground(lipgloss.Color(theme.Colors.Text))
	}

	// Description
	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.TextMuted)).
		PaddingLeft(1)

	// Background highlight for selected item
	itemContent := lipgloss.JoinHorizontal(
		lipgloss.Top,
		cursor+" ",
		iconStyle.Render(choice.icon),
		labelStyle.Render(choice.label),
		descStyle.Render("− "+choice.description),
	)

	if isSelected {
		// Add subtle background highlight
		highlightWidth := 55
		if d.width > 20 {
			highlightWidth = min(d.width-15, 60)
		}
		highlightStyle := lipgloss.NewStyle().
			Background(lipgloss.Color(theme.Colors.BackgroundHover)).
			Width(highlightWidth).
			Padding(0, 1)
		return "  " + highlightStyle.Render(itemContent)
	}

	return "  " + itemContent
}

// renderHelpBar renders the help/keyboard shortcuts bar.
func (d *Dashboard) renderHelpBar() string {
	theme := d.renderer.Theme()

	keyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Text)).
		Background(lipgloss.Color(theme.Colors.BackgroundAlt)).
		Padding(0, 1).
		Bold(true)

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.TextMuted)).
		PaddingRight(2)

	bindings := []struct {
		key  string
		desc string
	}{
		{"↑/↓", "Navigate"},
		{"Enter", "Select"},
		{"q/Esc", "Quit"},
	}

	var items []string
	for _, b := range bindings {
		items = append(items,
			lipgloss.JoinHorizontal(
				lipgloss.Top,
				keyStyle.Render(b.key),
				" ",
				descStyle.Render(b.desc),
			),
		)
	}

	return lipgloss.JoinHorizontal(lipgloss.Center, items...)
}

// renderFooter renders the footer with version and links.
func (d *Dashboard) renderFooter() string {
	theme := d.renderer.Theme()

	linkStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Info)).
		Underline(true)

	versionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.TextMuted))

	links := []string{
		"📚 " + linkStyle.Render("docs.clause.dev"),
		"💻 " + linkStyle.Render("github.com/clause-cli/clause"),
		versionStyle.Render("v" + d.version),
	}

	content := strings.Join(links, "  │  ")

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.TextMuted)).
		Padding(1, 0).
		Render(content)
}

// renderExitMessage renders the exit message when a command is selected.
func (d *Dashboard) renderExitMessage() string {
	theme := d.renderer.Theme()

	// If selected exit, show goodbye
	if d.selectedCmd == "exit" {
		return d.renderGoodbye()
	}

	// Find the selected choice
	var choice MenuChoice
	for _, c := range d.choices {
		if c.command == d.selectedCmd {
			choice = c
			break
		}
	}

	if choice.command == "" {
		return d.renderGoodbye()
	}

	// Success box
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(theme.Colors.Success)).
		Padding(1, 2).
		Margin(1, 2)

	titleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Success)).
		Bold(true).
		PaddingBottom(1)

	cmdStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Primary)).
		Background(lipgloss.Color(theme.Colors.BackgroundAlt)).
		Padding(0, 1).
		Bold(true)

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.TextMuted)).
		PaddingTop(1)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("✓ Selected: "+choice.label),
		"",
		"Run the following command:",
		"",
		cmdStyle.Render("  clause "+choice.command+"  "),
		descStyle.Render("Run this command in your terminal to execute."),
	)

	return "\n" + boxStyle.Render(content) + "\n"
}

// renderGoodbye renders the goodbye message.
func (d *Dashboard) renderGoodbye() string {
	theme := d.renderer.Theme()

	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.TextMuted)).
		Italic(true).
		Padding(1, 2)

	return "\n" + style.Render("👋 See you later! Run 'clause' anytime to get started.") + "\n"
}

// renderHelpScreen renders the help screen.
func (d *Dashboard) renderHelpScreen() string {
	theme := d.renderer.Theme()

	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(theme.Colors.Primary)).
		PaddingBottom(1)

	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(theme.Colors.TextMuted)).
		PaddingTop(1)

	cmdStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Primary)).
		Bold(true).
		Width(20)

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Text))

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		titleStyle.Render("📚 Help - Clause CLI Reference"),
		"",
		"Clause is a cross-platform CLI tool for creating AI-ready project structures.",
		"",
		headerStyle.Render("COMMANDS"),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top, cmdStyle.Render("init"), descStyle.Render("Initialize a new AI-ready project")),
		lipgloss.JoinHorizontal(lipgloss.Top, cmdStyle.Render("add"), descStyle.Render("Add components to existing project")),
		lipgloss.JoinHorizontal(lipgloss.Top, cmdStyle.Render("validate"), descStyle.Render("Check governance compliance")),
		lipgloss.JoinHorizontal(lipgloss.Top, cmdStyle.Render("config"), descStyle.Render("Manage Clause settings")),
		lipgloss.JoinHorizontal(lipgloss.Top, cmdStyle.Render("update"), descStyle.Render("Update to latest version")),
		"",
		headerStyle.Render("FLAGS"),
		"",
		lipgloss.JoinHorizontal(lipgloss.Top, cmdStyle.Render("--config"), descStyle.Render("Use specific config file")),
		lipgloss.JoinHorizontal(lipgloss.Top, cmdStyle.Render("--no-color"), descStyle.Render("Disable colored output")),
		lipgloss.JoinHorizontal(lipgloss.Top, cmdStyle.Render("-v, --verbose"), descStyle.Render("Enable verbose output")),
		lipgloss.JoinHorizontal(lipgloss.Top, cmdStyle.Render("-q, --quiet"), descStyle.Render("Suppress non-essential output")),
		"",
		"Press any key to go back...",
	)

	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(theme.Colors.Border)).
		Padding(1, 2).
		Width(70)

	return lipgloss.NewStyle().
		Align(lipgloss.Center, lipgloss.Center).
		Render(boxStyle.Render(content))
}

// StartDashboard launches the interactive dashboard.
func StartDashboard(rootCmd *cobra.Command, version string) error {
	p := tea.NewProgram(NewDashboard(rootCmd, version), tea.WithAltScreen())
	_, err := p.Run()
	return err
}

// min returns the minimum of two integers.
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
