package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
)

// InputModel is a text input component.
type InputModel struct {
	// Value is the current input value
	Value string

	// Placeholder is shown when value is empty
	Placeholder string

	// Prompt is the input prompt character
	Prompt string

	// Width is the input width
	Width int

	// Focused indicates if the input has focus
	Focused bool

	// CharLimit is the maximum character limit (0 = no limit)
	CharLimit int

	// Theme is the current theme
	Theme *styles.Theme

	// Cursor position
	cursorPos int

	// Error message to display
	Error string

	// Label for the input
	Label string
}

// NewInput creates a new input component.
func NewInput() InputModel {
	return InputModel{
		Prompt:     "> ",
		Width:      40,
		Focused:    true,
		cursorPos:  0,
	}
}

// Init initializes the input.
func (m InputModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the input.
func (m InputModel) Update(msg tea.Msg) (InputModel, tea.Cmd) {
	if !m.Focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyBackspace:
			if m.cursorPos > 0 && len(m.Value) > 0 {
				m.Value = m.Value[:m.cursorPos-1] + m.Value[m.cursorPos:]
				m.cursorPos--
			}
		case tea.KeyDelete:
			if m.cursorPos < len(m.Value) {
				m.Value = m.Value[:m.cursorPos] + m.Value[m.cursorPos+1:]
			}
		case tea.KeyLeft:
			if m.cursorPos > 0 {
				m.cursorPos--
			}
		case tea.KeyRight:
			if m.cursorPos < len(m.Value) {
				m.cursorPos++
			}
		case tea.KeyHome:
			m.cursorPos = 0
		case tea.KeyEnd:
			m.cursorPos = len(m.Value)
		case tea.KeyRunes:
			if m.CharLimit == 0 || len(m.Value) < m.CharLimit {
				runes := string(msg.Runes)
				m.Value = m.Value[:m.cursorPos] + runes + m.Value[m.cursorPos:]
				m.cursorPos += len(runes)
			}
		}
	}

	return m, nil
}

// View renders the input.
func (m InputModel) View() string {
	var b strings.Builder

	// Label
	if m.Label != "" {
		labelStyle := lipgloss.NewStyle().Bold(true)
		if m.Theme != nil {
			labelStyle = labelStyle.Foreground(lipgloss.Color(m.Theme.Colors.Text))
		}
		b.WriteString(labelStyle.Render(m.Label))
		b.WriteString("\n")
	}

	// Input box
	var inputStyle lipgloss.Style
	if m.Theme != nil {
		if m.Focused {
			inputStyle = m.Theme.Component.InputFocused
		} else {
			inputStyle = m.Theme.Component.Input
		}
	} else {
		inputStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1)
	}

	// Build the display value
	display := m.Value
	if display == "" && !m.Focused {
		if m.Theme != nil {
			display = m.Theme.Typography.Muted.Render(m.Placeholder)
		} else {
			display = m.Placeholder
		}
	}

	// Add cursor
	if m.Focused {
		if m.cursorPos < len(m.Value) {
			// Cursor in middle
			before := m.Value[:m.cursorPos]
			at := string(m.Value[m.cursorPos])
			after := m.Value[m.cursorPos+1:]
			cursorStyle := lipgloss.NewStyle().Reverse(true)
			display = before + cursorStyle.Render(at) + after
		} else {
			// Cursor at end
			cursorStyle := lipgloss.NewStyle().Reverse(true)
			display = m.Value + cursorStyle.Render(" ")
		}
	}

	b.WriteString(inputStyle.Width(m.Width).Render(m.Prompt + display))

	// Error message
	if m.Error != "" {
		b.WriteString("\n")
		if m.Theme != nil {
			b.WriteString(m.Theme.Typography.Error.Render(m.Error))
		} else {
			b.WriteString("Error: " + m.Error)
		}
	}

	return b.String()
}

// SetValue sets the input value.
func (m *InputModel) SetValue(value string) {
	m.Value = value
	m.cursorPos = len(value)
}

// SetTheme sets the theme.
func (m *InputModel) SetTheme(theme *styles.Theme) {
	m.Theme = theme
}

// SetFocused sets the focus state.
func (m *InputModel) SetFocused(focused bool) {
	m.Focused = focused
}

// SetWidth sets the width.
func (m *InputModel) SetWidth(width int) {
	m.Width = width
}

// SetPlaceholder sets the placeholder.
func (m *InputModel) SetPlaceholder(placeholder string) {
	m.Placeholder = placeholder
}

// SetLabel sets the label.
func (m *InputModel) SetLabel(label string) {
	m.Label = label
}

// SetError sets an error message.
func (m *InputModel) SetError(err string) {
	m.Error = err
}

// ClearError clears any error message.
func (m *InputModel) ClearError() {
	m.Error = ""
}

// Reset resets the input to its initial state.
func (m *InputModel) Reset() {
	m.Value = ""
	m.cursorPos = 0
	m.Error = ""
}

// Blur removes focus from the input.
func (m *InputModel) Blur() {
	m.Focused = false
}

// Focus gives focus to the input.
func (m *InputModel) Focus() {
	m.Focused = true
}
