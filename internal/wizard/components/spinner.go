package components

import (
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// SpinnerModel is a loading spinner component.
type SpinnerModel struct {
	// Style is the spinner style (dots, line, bounce, pulse, points)
	Style string

	// Text is the text to display next to the spinner
	Text string

	// Theme is the current theme
	Theme *styles.Theme

	// Animation is the underlying animation
	animation *tui.Animation

	// Running indicates if the spinner is active
	Running bool
}

// TickMsg is a message for spinner animation.
type TickMsg struct {
	Time time.Time
}

// NewSpinner creates a new spinner component.
func NewSpinner() SpinnerModel {
	return SpinnerModel{
		Style:  "dots",
		Running: true,
	}
}

// NewSpinnerWithStyle creates a new spinner with a specific style.
func NewSpinnerWithStyle(style string) SpinnerModel {
	s := NewSpinner()
	s.Style = style
	return s
}

// Init initializes the spinner.
func (m SpinnerModel) Init() tea.Cmd {
	if m.animation == nil {
		m.animation = tui.NewSpinner(m.Style)
	}
	return tickCmd(80 * time.Millisecond)
}

// Update handles updates for the spinner.
func (m SpinnerModel) Update(msg tea.Msg) (SpinnerModel, tea.Cmd) {
	if !m.Running {
		return m, nil
	}

	switch msg := msg.(type) {
	case TickMsg:
		if m.animation == nil {
			m.animation = tui.NewSpinner(m.Style)
		}
		m.animation.Update(msg.Time)
		return m, tickCmd(80 * time.Millisecond)
	}

	return m, nil
}

// View renders the spinner.
func (m SpinnerModel) View() string {
	if m.animation == nil {
		return ""
	}

	spinner := m.animation.Current()

	var result string
	if m.Theme != nil {
		spinnerStyle := m.Theme.Component.Spinner
		result = spinnerStyle.Render(spinner)
	} else {
		result = spinner
	}

	if m.Text != "" {
		result += " " + m.Text
	}

	return result
}

// SetText sets the spinner text.
func (m *SpinnerModel) SetText(text string) {
	m.Text = text
}

// SetStyle sets the spinner style.
func (m *SpinnerModel) SetStyle(style string) {
	m.Style = style
	m.animation = tui.NewSpinner(style)
}

// SetTheme sets the theme.
func (m *SpinnerModel) SetTheme(theme *styles.Theme) {
	m.Theme = theme
}

// Start starts the spinner.
func (m *SpinnerModel) Start() {
	m.Running = true
}

// Stop stops the spinner.
func (m *SpinnerModel) Stop() {
	m.Running = false
}

// tickCmd returns a command that sends a tick message.
func tickCmd(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg{Time: t}
	})
}

// LoadingState represents a loading state with spinner.
type LoadingState struct {
	Spinner  SpinnerModel
	Message  string
	Progress float64
	Width    int
}

// NewLoadingState creates a new loading state.
func NewLoadingState(message string) LoadingState {
	spinner := NewSpinner()
	spinner.SetText(message)
	return LoadingState{
		Spinner:  spinner,
		Message:  message,
		Progress: 0,
		Width:    60,
	}
}

// Init initializes the loading state.
func (s LoadingState) Init() tea.Cmd {
	return s.Spinner.Init()
}

// Update handles updates for the loading state.
func (s LoadingState) Update(msg tea.Msg) (LoadingState, tea.Cmd) {
	var cmd tea.Cmd
	s.Spinner, cmd = s.Spinner.Update(msg)
	return s, cmd
}

// View renders the loading state.
func (s LoadingState) View() string {
	var b strings.Builder

	// Spinner
	b.WriteString(s.Spinner.View())
	b.WriteString("\n\n")

	// Message
	if s.Message != "" {
		if s.Spinner.Theme != nil {
			b.WriteString(s.Spinner.Theme.Typography.Body.Render(s.Message))
		} else {
			b.WriteString(s.Message)
		}
		b.WriteString("\n\n")
	}

	// Progress bar (if set)
	if s.Progress > 0 {
		progress := NewProgress()
		progress.SetPercent(s.Progress)
		progress.SetWidth(s.Width - 10)
		progress.SetTheme(s.Spinner.Theme)
		b.WriteString(progress.View())
	}

	return b.String()
}

// SetProgress sets the progress (0.0 to 1.0).
func (s *LoadingState) SetProgress(progress float64) {
	s.Progress = progress
}

// SetMessage sets the message.
func (s *LoadingState) SetMessage(message string) {
	s.Message = message
	s.Spinner.SetText(message)
}

// SetWidth sets the width.
func (s *LoadingState) SetWidth(width int) {
	s.Width = width
}

// SetTheme sets the theme.
func (s *LoadingState) SetTheme(theme *styles.Theme) {
	s.Spinner.SetTheme(theme)
}

// CompleteText returns a completion message with checkmark.
func CompleteText(message string, theme *styles.Theme) string {
	if theme != nil {
		return theme.Typography.Success.Render("✓ " + message)
	}
	return "✓ " + message
}

// ErrorText returns an error message.
func ErrorText(message string, theme *styles.Theme) string {
	if theme != nil {
		return theme.Typography.Error.Render("✗ " + message)
	}
	return "✗ " + message
}

// StatusLine returns a status line with icon.
func StatusLine(icon, message string, status string, theme *styles.Theme) string {
	var result string
	switch status {
	case "success":
		if theme != nil {
			result = theme.Typography.Success.Render(icon + " " + message)
		} else {
			result = "\x1b[32m" + icon + " " + message + "\x1b[0m"
		}
	case "error":
		if theme != nil {
			result = theme.Typography.Error.Render(icon + " " + message)
		} else {
			result = "\x1b[31m" + icon + " " + message + "\x1b[0m"
		}
	case "warning":
		if theme != nil {
			result = theme.Typography.Warning.Render(icon + " " + message)
		} else {
			result = "\x1b[33m" + icon + " " + message + "\x1b[0m"
		}
	default:
		if theme != nil {
			result = theme.Typography.Body.Render(icon + " " + message)
		} else {
			result = icon + " " + message
		}
	}

	return result
}
