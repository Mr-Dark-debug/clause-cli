package screens

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
	"github.com/clause-cli/clause/pkg/utils"
)

// ProjectScreen collects basic project information.
type ProjectScreen struct {
	BaseScreen
	fields      []projectField
	activeField int
}

type projectField struct {
	name        string
	key         string
	value       string
	placeholder string
	validator   func(string) bool
}

// NewProjectScreen creates a new project screen.
func NewProjectScreen() *ProjectScreen {
	s := &ProjectScreen{
		BaseScreen: *NewBaseScreen("Project Info", "project"),
		fields: []projectField{
			{
				name:        "Project Name",
				key:         "name",
				placeholder: "my-awesome-project",
				validator:   isValidProjectName,
			},
			{
				name:        "Description",
				key:         "description",
				placeholder: "A brief description of your project",
				validator:   func(s string) bool { return true },
			},
			{
				name:        "Author",
				key:         "author",
				placeholder: "Your name or organization",
				validator:   func(s string) bool { return true },
			},
			{
				name:        "License",
				key:         "license",
				placeholder: "MIT",
				validator:   func(s string) bool { return true },
			},
		},
		activeField: 0,
	}

	return s
}

// Init initializes the screen.
func (s *ProjectScreen) Init() tea.Cmd {
	return nil
}

// Update handles updates for the screen.
func (s *ProjectScreen) Update(msg tea.Msg) tea.Cmd {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.Type {
		case tea.KeyUp:
			if s.activeField > 0 {
				s.activeField--
			}
		case tea.KeyDown, tea.KeyTab:
			if s.activeField < len(s.fields)-1 {
				s.activeField++
			}
		case tea.KeyBackspace:
			if len(s.fields[s.activeField].value) > 0 {
				s.fields[s.activeField].value = s.fields[s.activeField].value[:len(s.fields[s.activeField].value)-1]
			}
		case tea.KeyEnter:
			// Move to next field or finish
			if s.activeField < len(s.fields)-1 {
				s.activeField++
			} else if s.IsComplete() {
				s.applyValues()
				return nil // Navigation handled by wizard
			}
		default:
			// Handle text input
			if m.Type == tea.KeyRunes {
				s.fields[s.activeField].value += string(m.Runes)
			}
		}
	}

	// Update completion status
	s.complete = s.validateAll()

	return nil
}

// View renders the screen.
func (s *ProjectScreen) View() string {
	var b strings.Builder

	// Title
	b.WriteString(s.Renderer().Title("Project Information"))
	b.WriteString("\n\n")

	// Description
	b.WriteString(s.Renderer().Body("Let's start with some basic information about your project."))
	b.WriteString("\n\n")

	// Fields
	for i, f := range s.fields {
		label := s.Renderer().Header(f.name)
		b.WriteString(label)
		b.WriteString("\n")

		focused := i == s.activeField
		input := s.Renderer().InputField(f.value, f.placeholder, focused, s.Width()-10)
		b.WriteString(input)

		// Validation indicator
		if f.value != "" {
			if f.validator(f.value) {
				b.WriteString(" ")
				b.WriteString(s.Renderer().Success("✓"))
			} else {
				b.WriteString(" ")
				b.WriteString(s.Renderer().Warning("⚠"))
			}
		}

		b.WriteString("\n\n")
	}

	// Help
	kb := tui.NewKeyBindings()
	kb.Add("↑/↓", "Navigate fields")
	kb.Add("Tab", "Next field")
	if s.IsComplete() {
		kb.Add("Enter", "Continue")
	}
	b.WriteString(s.Renderer().HelpText(kb))

	return b.String()
}

// validateAll checks if all required fields are valid.
func (s *ProjectScreen) validateAll() bool {
	// Project name is required
	if s.fields[0].value == "" || !s.fields[0].validator(s.fields[0].value) {
		return false
	}
	return true
}

// applyValues applies the field values to the config.
func (s *ProjectScreen) applyValues() {
	if s.config == nil {
		return
	}
	for _, f := range s.fields {
		switch f.key {
		case "name":
			s.config.Metadata.Name = f.value
		case "description":
			s.config.Metadata.Description = f.value
		case "author":
			s.config.Metadata.Author = f.value
		case "license":
			s.config.Metadata.License = f.value
		}
	}
}

// isValidProjectName validates a project name.
func isValidProjectName(name string) bool {
	if len(name) == 0 || len(name) > 100 {
		return false
	}
	// Must start with letter, contain only lowercase letters, numbers, and hyphens (uses helper from utils)
	return utils.MatchesRegex(name, `^[a-z][a-z0-9-]*$`)
}

// SetTheme sets the theme.
func (s *ProjectScreen) SetTheme(theme *styles.Theme) {
	s.BaseScreen.SetTheme(theme)
}

// SetConfig sets the config.
func (s *ProjectScreen) SetConfig(config *config.ProjectConfig) {
	s.BaseScreen.SetConfig(config)
}

// SetSize sets the size.
func (s *ProjectScreen) SetSize(width, height int) {
	s.BaseScreen.SetSize(width, height)
}
