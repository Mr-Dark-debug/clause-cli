package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// FormField represents a field in a form.
type FormField struct {
	// Key is the field identifier
	Key string

	// Label is the field label
	Label string

	// Placeholder is the input placeholder
	Placeholder string

	// Value is the current value
	Value string

	// Error is any validation error
	Error string

	// Required indicates if the field is required
	Required bool

	// Validator is a validation function
	Validator func(string) error

	// Type is the field type (text, password, etc.)
	Type string
}

// FormModel is a form component with multiple fields.
type FormModel struct {
	// Fields is the list of form fields
	Fields []FormField

	// FocusIndex is the currently focused field
	FocusIndex int

	// Width is the form width
	Width int

	// Theme is the current theme
	Theme *styles.Theme

	// Label for the form
	Label string

	// Focused indicates if the form has focus
	Focused bool

	// ShowHelp indicates if help should be shown
	ShowHelp bool
}

// NewForm creates a new form component.
func NewForm(fields []FormField) FormModel {
	return FormModel{
		Fields:     fields,
		FocusIndex: 0,
		Width:      60,
		Focused:    true,
		ShowHelp:   true,
	}
}

// Init initializes the form.
func (m FormModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the form.
func (m FormModel) Update(msg tea.Msg) (FormModel, tea.Cmd) {
	if !m.Focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyUp:
			if m.FocusIndex > 0 {
				m.FocusIndex--
			}
		case tea.KeyDown, tea.KeyTab:
			if m.FocusIndex < len(m.Fields)-1 {
				m.FocusIndex++
			}
		case tea.KeyShiftTab:
			if m.FocusIndex > 0 {
				m.FocusIndex--
			}
		case tea.KeyBackspace:
			field := &m.Fields[m.FocusIndex]
			if len(field.Value) > 0 {
				field.Value = field.Value[:len(field.Value)-1]
			}
			m.validateField(m.FocusIndex)
		case tea.KeyEnter:
			// Move to next field or finish
			if m.FocusIndex < len(m.Fields)-1 {
				m.FocusIndex++
			}
		case tea.KeyRunes:
			field := &m.Fields[m.FocusIndex]
			field.Value += string(msg.Runes)
			m.validateField(m.FocusIndex)
		}
	}

	return m, nil
}

// View renders the form.
func (m FormModel) View() string {
	var b strings.Builder

	// Form label
	if m.Label != "" {
		labelStyle := lipgloss.NewStyle().Bold(true)
		if m.Theme != nil {
			labelStyle = labelStyle.Foreground(lipgloss.Color(m.Theme.Colors.Primary))
		}
		b.WriteString(labelStyle.Render(m.Label))
		b.WriteString("\n\n")
	}

	// Fields
	for i, field := range m.Fields {
		isFocused := i == m.FocusIndex

		// Label
		labelStyle := lipgloss.NewStyle().Bold(true)
		if m.Theme != nil {
			if isFocused {
				labelStyle = labelStyle.Foreground(lipgloss.Color(m.Theme.Colors.Primary))
			} else {
				labelStyle = labelStyle.Foreground(lipgloss.Color(m.Theme.Colors.Text))
			}
		}
		b.WriteString(labelStyle.Render(field.Label))
		if field.Required {
			b.WriteString(" *")
		}
		b.WriteString("\n")

		// Input field
		input := m.renderInput(field, isFocused)
		b.WriteString(input)
		b.WriteString("\n")

		// Error message
		if field.Error != "" {
			if m.Theme != nil {
				b.WriteString(m.Theme.Typography.Error.Render("  " + field.Error))
			} else {
				b.WriteString("  " + field.Error)
			}
			b.WriteString("\n")
		}

		b.WriteString("\n")
	}

	// Help
	if m.ShowHelp {
		kb := tui.NewKeyBindings()
		kb.Add("↑/↓/Tab", "Navigate")
		kb.Add("Enter", "Next field")
		if m.Theme != nil {
			b.WriteString(m.Theme.Typography.Muted.Render(kb.Help()))
		} else {
			b.WriteString(kb.Help())
		}
	}

	return b.String()
}

// renderInput renders a single input field.
func (m FormModel) renderInput(field FormField, focused bool) string {
	var style lipgloss.Style
	if m.Theme != nil {
		if focused {
			style = m.Theme.Component.InputFocused
		} else {
			style = m.Theme.Component.Input
		}
	} else {
		style = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(0, 1)
	}

	style = style.Width(m.Width)

	display := field.Value
	if display == "" && !focused {
		if m.Theme != nil {
			display = m.Theme.Typography.Muted.Render(field.Placeholder)
		} else {
			display = field.Placeholder
		}
	}

	// Handle password type
	if field.Type == "password" && field.Value != "" {
		display = strings.Repeat("•", len(field.Value))
	}

	// Add cursor
	if focused {
		cursorStyle := lipgloss.NewStyle().Reverse(true)
		display = display + cursorStyle.Render(" ")
	}

	return style.Render(display)
}

// validateField validates a specific field.
func (m *FormModel) validateField(index int) {
	if index < 0 || index >= len(m.Fields) {
		return
	}

	field := &m.Fields[index]
	field.Error = ""

	if field.Required && field.Value == "" {
		field.Error = "This field is required"
		return
	}

	if field.Validator != nil && field.Value != "" {
		if err := field.Validator(field.Value); err != nil {
			field.Error = err.Error()
		}
	}
}

// Validate validates all fields.
func (m *FormModel) Validate() bool {
	valid := true
	for i := range m.Fields {
		m.validateField(i)
		if m.Fields[i].Error != "" {
			valid = false
		}
	}
	return valid
}

// Values returns a map of field keys to values.
func (m FormModel) Values() map[string]string {
	values := make(map[string]string)
	for _, field := range m.Fields {
		values[field.Key] = field.Value
	}
	return values
}

// SetValue sets a field value by key.
func (m *FormModel) SetValue(key, value string) {
	for i := range m.Fields {
		if m.Fields[i].Key == key {
			m.Fields[i].Value = value
			m.validateField(i)
			return
		}
	}
}

// GetValue gets a field value by key.
func (m FormModel) GetValue(key string) string {
	for _, field := range m.Fields {
		if field.Key == key {
			return field.Value
		}
	}
	return ""
}

// SetTheme sets the theme.
func (m *FormModel) SetTheme(theme *styles.Theme) {
	m.Theme = theme
}

// SetWidth sets the width.
func (m *FormModel) SetWidth(width int) {
	m.Width = width
}

// SetLabel sets the label.
func (m *FormModel) SetLabel(label string) {
	m.Label = label
}

// SetFocused sets the focus state.
func (m *FormModel) SetFocused(focused bool) {
	m.Focused = focused
}

// Blur removes focus.
func (m *FormModel) Blur() {
	m.Focused = false
}

// Focus gives focus.
func (m *FormModel) Focus() {
	m.Focused = true
}

// IsValid returns true if all fields are valid.
func (m FormModel) IsValid() bool {
	for _, field := range m.Fields {
		if field.Error != "" {
			return false
		}
		if field.Required && field.Value == "" {
			return false
		}
	}
	return true
}

// Reset resets all fields to empty.
func (m *FormModel) Reset() {
	for i := range m.Fields {
		m.Fields[i].Value = ""
		m.Fields[i].Error = ""
	}
	m.FocusIndex = 0
}
