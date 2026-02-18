package screens

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// WelcomeScreen is the initial welcome screen of the wizard.
type WelcomeScreen struct {
	BaseScreen
	cursor    int
	presets   []config.Preset
	presetIdx int
}

// NewWelcomeScreen creates a new welcome screen.
func NewWelcomeScreen() *WelcomeScreen {
	s := &WelcomeScreen{
		BaseScreen: *NewBaseScreen("Welcome", "welcome"),
		presetIdx:  0,
	}
	s.complete = true // Welcome screen is always complete

	// Load available presets
	s.presets = config.AvailablePresets

	return s
}

// Init initializes the screen.
func (s *WelcomeScreen) Init() tea.Cmd {
	return nil
}

// Update handles updates for the screen.
func (s *WelcomeScreen) Update(msg tea.Msg) tea.Cmd {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.String() {
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			if s.cursor < 2 {
				s.cursor++
			}
		case "enter", " ":
			return s.applyPreset()
		}
	}

	return nil
}

// View renders the screen.
func (s *WelcomeScreen) View() string {
	var b strings.Builder

	// Logo and title
	b.WriteString(s.Renderer().Title("Clause"))
	b.WriteString("\n\n")

	// Tagline
	tagline := "AI-Native Project Scaffolding Tool"
	b.WriteString(s.Renderer().Subtitle(tagline))
	b.WriteString("\n\n")

	// Description
	desc := "Create production-ready projects with AI-assisted development workflows."
	b.WriteString(s.Renderer().Body(desc))
	b.WriteString("\n\n")

	// Divider
	b.WriteString(s.Renderer().Divider(s.Width() - 4))
	b.WriteString("\n\n")

	// Preset selection
	b.WriteString(s.Renderer().Header("Choose a starting point:"))
	b.WriteString("\n\n")

	options := []struct {
		title       string
		description string
	}{
		{"Quick Start", "Minimal configuration, get coding fast"},
		{"Standard", "Balanced setup with common features"},
		{"Custom", "Full control over all options"},
	}

	for i, opt := range options {
		if i == s.cursor {
			b.WriteString(s.Renderer().ListItem("▸ "+opt.title+": "+opt.description, true))
		} else {
			b.WriteString(s.Renderer().ListItem("  "+opt.title+": "+opt.description, false))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")

	// Key bindings
	kb := tui.NewKeyBindings()
	kb.Add("↑/↓", "Navigate")
	kb.Add("Enter", "Select")
	b.WriteString(s.Renderer().HelpText(kb))

	return b.String()
}

// CanGoBack returns false - can't go back from welcome.
func (s *WelcomeScreen) CanGoBack() bool {
	return false
}

// applyPreset applies the selected preset.
func (s *WelcomeScreen) applyPreset() tea.Cmd {
	return func() tea.Msg {
		if s.config == nil {
			return nil
		}

		var presetName string
		switch s.cursor {
		case 0:
			presetName = "minimal"
		case 1:
			presetName = "standard"
		default:
			presetName = "" // Custom - use defaults
		}

		if presetName != "" {
			if preset, err := config.LoadPreset(presetName); err == nil {
				s.config.Metadata = preset.Metadata
				s.config.Frontend = preset.Frontend
				s.config.Backend = preset.Backend
				s.config.Infrastructure = preset.Infrastructure
				s.config.Governance = preset.Governance
			}
		}

		return nil
	}
}

// SetTheme sets the theme.
func (s *WelcomeScreen) SetTheme(theme *styles.Theme) {
	s.BaseScreen.SetTheme(theme)
}

// SetConfig sets the config.
func (s *WelcomeScreen) SetConfig(config *config.ProjectConfig) {
	s.BaseScreen.SetConfig(config)
}

// SetSize sets the size.
func (s *WelcomeScreen) SetSize(width, height int) {
	s.BaseScreen.SetSize(width, height)
}
