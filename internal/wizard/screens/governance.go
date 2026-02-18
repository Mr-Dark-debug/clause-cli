package screens

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// GovernanceScreen configures AI governance options.
type GovernanceScreen struct {
	BaseScreen
	cursor          int
	contextLevelIdx int
	features        map[string]bool
}

// Context level options
var contextLevels = []struct {
	name        string
	key         string
	description string
}{
	{"Minimal", "minimal", "Essential context only"},
	{"Standard", "standard", "Balanced context for most projects"},
	{"Comprehensive", "comprehensive", "Full context for complex projects"},
}

// Governance feature options
var governanceFeatureOptions = []struct {
	key         string
	name        string
	description string
}{
	{"brainstorm_md", "Brainstorm.md", "AI brainstorming workspace"},
	{"prompt_guidelines", "Prompt Guidelines", "AI prompt templates"},
	{"component_registry", "Component Registry", "Reusable component tracking"},
	{"layered_prompts", "Layered Prompts", "Multi-level AI context"},
	{"code_standards", "Code Standards", "Project coding standards"},
}

// NewGovernanceScreen creates a new governance screen.
func NewGovernanceScreen() *GovernanceScreen {
	return &GovernanceScreen{
		BaseScreen:     *NewBaseScreen("AI Governance", "governance"),
		contextLevelIdx: 1, // Default to Standard
		features: map[string]bool{
			"brainstorm_md":      true,
			"prompt_guidelines":  true,
			"component_registry": true,
			"layered_prompts":    true,
			"code_standards":     true,
		},
		cursor: 0,
	}
}

// Init initializes the screen.
func (s *GovernanceScreen) Init() tea.Cmd {
	return nil
}

// Update handles updates.
func (s *GovernanceScreen) Update(msg tea.Msg) tea.Cmd {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.String() {
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			max := len(contextLevels) + len(governanceFeatureOptions)
			if s.cursor < max-1 {
				s.cursor++
			}
		case "enter", " ":
			s.toggle()
		}
	}

	s.complete = true
	return nil
}

func (s *GovernanceScreen) toggle() {
	// Check if selecting context level
	if s.cursor < len(contextLevels) {
		s.contextLevelIdx = s.cursor
	} else {
		// Toggle feature
		featureIdx := s.cursor - len(contextLevels)
		if featureIdx >= 0 && featureIdx < len(governanceFeatureOptions) {
			key := governanceFeatureOptions[featureIdx].key
			s.features[key] = !s.features[key]
		}
	}
}

// View renders the screen.
func (s *GovernanceScreen) View() string {
	var b strings.Builder

	b.WriteString(s.Renderer().Title("AI Governance Configuration"))
	b.WriteString("\n\n")

	// Description
	b.WriteString(s.Renderer().Body("Configure how Clause helps you manage AI-assisted development."))
	b.WriteString("\n\n")

	// Context level
	b.WriteString(s.Renderer().Header("Context Level"))
	b.WriteString("\n")
	b.WriteString(s.Renderer().Muted("Choose how much context to provide to AI assistants"))
	b.WriteString("\n\n")

	for i, level := range contextLevels {
		if i == s.cursor {
			b.WriteString(s.Renderer().ListItem("▸ "+level.name, true))
		} else if i == s.contextLevelIdx {
			b.WriteString(s.Renderer().ListItem("● "+level.name, false))
		} else {
			b.WriteString(s.Renderer().ListItem("○ "+level.name, false))
		}
		b.WriteString(s.Renderer().Muted(" - "+level.description))
		b.WriteString("\n")
	}

	b.WriteString("\n")

	// Features
	b.WriteString(s.Renderer().Header("Governance Features"))
	b.WriteString("\n\n")

	for i, feat := range governanceFeatureOptions {
		cursorIdx := i + len(contextLevels)
		checked := s.features[feat.key]
		line := s.Renderer().Checkbox(feat.name, checked)

		if s.cursor == cursorIdx {
			b.WriteString("▸ " + line + "\n")
		} else {
			b.WriteString("  " + line + "\n")
		}
	}

	b.WriteString("\n")

	// Help
	kb := tui.NewKeyBindings()
	kb.Add("↑/↓", "Navigate")
	kb.Add("Enter/Space", "Select/Toggle")
	b.WriteString(s.Renderer().HelpText(kb))

	return b.String()
}

// ApplyToConfig applies settings to config.
func (s *GovernanceScreen) ApplyToConfig() {
	if s.config == nil {
		return
	}

	if s.contextLevelIdx < len(contextLevels) {
		s.config.Governance.ContextLevel = contextLevels[s.contextLevelIdx].key
	}

	s.config.Governance.Enabled = true
	s.config.Governance.BrainstormMd = s.features["brainstorm_md"]
	s.config.Governance.PromptGuidelines = s.features["prompt_guidelines"]
	s.config.Governance.ComponentRegistry = s.features["component_registry"]
}

// SetTheme sets the theme.
func (s *GovernanceScreen) SetTheme(theme *styles.Theme) {
	s.BaseScreen.SetTheme(theme)
}

// SetConfig sets the config.
func (s *GovernanceScreen) SetConfig(config *config.ProjectConfig) {
	s.BaseScreen.SetConfig(config)
}

// SetSize sets the size.
func (s *GovernanceScreen) SetSize(width, height int) {
	s.BaseScreen.SetSize(width, height)
}
