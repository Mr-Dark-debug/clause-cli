package components

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
)

// ProgressModel is a progress bar component.
type ProgressModel struct {
	// Percent is the completion percentage (0.0 to 1.0)
	Percent float64

	// Width is the progress bar width
	Width int

	// ShowPercent indicates if percentage should be shown
	ShowPercent bool

	// ShowCount indicates if count should be shown (e.g., "5/10")
	ShowCount bool

	// Current count
	Current int

	// Total count
	Total int

	// Theme is the current theme
	Theme *styles.Theme

	// Label for the progress bar
	Label string

	// Filled character
	FilledChar string

	// Empty character
	EmptyChar string

	// Animated indicates if the bar should animate
	Animated bool

	// animation state
	animPhase float64
}

// NewProgress creates a new progress component.
func NewProgress() ProgressModel {
	return ProgressModel{
		Width:       40,
		ShowPercent: true,
		FilledChar:  "█",
		EmptyChar:   "░",
		Animated:    false,
	}
}

// Init initializes the progress.
func (m ProgressModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the progress.
func (m ProgressModel) Update(msg tea.Msg) (ProgressModel, tea.Cmd) {
	if m.Animated {
		m.animPhase += 0.1
		if m.animPhase > 1 {
			m.animPhase = 0
		}
	}
	return m, nil
}

// View renders the progress bar.
func (m ProgressModel) View() string {
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

	// Calculate filled width
	filledWidth := int(float64(m.Width) * m.Percent)
	if filledWidth > m.Width {
		filledWidth = m.Width
	}
	if filledWidth < 0 {
		filledWidth = 0
	}
	emptyWidth := m.Width - filledWidth

	// Build bar
	var bar strings.Builder
	if m.Theme != nil {
		filledStyle := m.Theme.Component.ProgressFilled
		emptyStyle := m.Theme.Component.Progress

		bar.WriteString(filledStyle.Render(strings.Repeat(m.FilledChar, filledWidth)))
		bar.WriteString(emptyStyle.Render(strings.Repeat(m.EmptyChar, emptyWidth)))
	} else {
		bar.WriteString(strings.Repeat(m.FilledChar, filledWidth))
		bar.WriteString(strings.Repeat(m.EmptyChar, emptyWidth))
	}

	b.WriteString(bar.String())

	// Percentage or count
	if m.ShowPercent {
		percentText := fmt.Sprintf(" %.0f%%", m.Percent*100)
		if m.Theme != nil {
			b.WriteString(m.Theme.Typography.Body.Render(percentText))
		} else {
			b.WriteString(percentText)
		}
	} else if m.ShowCount && m.Total > 0 {
		countText := fmt.Sprintf(" %d/%d", m.Current, m.Total)
		if m.Theme != nil {
			b.WriteString(m.Theme.Typography.Body.Render(countText))
		} else {
			b.WriteString(countText)
		}
	}

	return b.String()
}

// SetPercent sets the progress percentage.
func (m *ProgressModel) SetPercent(percent float64) {
	m.Percent = percent
	if m.Percent < 0 {
		m.Percent = 0
	}
	if m.Percent > 1 {
		m.Percent = 1
	}
}

// SetCount sets the current and total counts.
func (m *ProgressModel) SetCount(current, total int) {
	m.Current = current
	m.Total = total
	m.ShowCount = true
	m.ShowPercent = false

	if total > 0 {
		m.Percent = float64(current) / float64(total)
	}
}

// SetTheme sets the theme.
func (m *ProgressModel) SetTheme(theme *styles.Theme) {
	m.Theme = theme
}

// SetWidth sets the width.
func (m *ProgressModel) SetWidth(width int) {
	m.Width = width
}

// SetLabel sets the label.
func (m *ProgressModel) SetLabel(label string) {
	m.Label = label
}

// Increment increments the progress by a given amount.
func (m *ProgressModel) Increment(amount float64) {
	m.SetPercent(m.Percent + amount)
}

// IncrementCount increments the current count.
func (m *ProgressModel) IncrementCount() {
	m.Current++
	if m.Total > 0 && m.Current <= m.Total {
		m.Percent = float64(m.Current) / float64(m.Total)
	}
}

// Complete sets the progress to 100%.
func (m *ProgressModel) Complete() {
	m.Percent = 1.0
	m.Current = m.Total
}

// Reset resets the progress to 0.
func (m *ProgressModel) Reset() {
	m.Percent = 0
	m.Current = 0
}

// IsComplete returns true if progress is at 100%.
func (m ProgressModel) IsComplete() bool {
	return m.Percent >= 1.0
}

// FullProgress returns a full-width progress bar for the given width.
func FullProgress(percent float64, width int, theme *styles.Theme) string {
	m := NewProgress()
	m.SetPercent(percent)
	m.SetWidth(width - 10) // Leave room for percentage
	m.SetTheme(theme)
	m.ShowPercent = true
	return m.View()
}
