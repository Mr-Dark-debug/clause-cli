package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// SelectItem represents an item in a select list.
type SelectItem struct {
	// Label is the displayed text
	Label string

	// Value is the underlying value
	Value string

	// Description is an optional description
	Description string

	// Disabled indicates if the item can be selected
	Disabled bool
}

// SelectModel is a single-selection dropdown component.
type SelectModel struct {
	// Items is the list of selectable items
	Items []SelectItem

	// Selected is the index of the selected item
	Selected int

	// Width is the component width
	Width int

	// Height is the visible height (for scrolling)
	Height int

	// Focused indicates if the component has focus
	Focused bool

	// Theme is the current theme
	Theme *styles.Theme

	// Label for the component
	Label string

	// Scroll offset for long lists
	offset int

	// Show descriptions
	showDescriptions bool
}

// NewSelect creates a new select component.
func NewSelect(items []SelectItem) SelectModel {
	return SelectModel{
		Items:            items,
		Selected:         0,
		Width:            40,
		Height:           10,
		Focused:          true,
		offset:           0,
		showDescriptions: true,
	}
}

// Init initializes the select.
func (m SelectModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the select.
func (m SelectModel) Update(msg tea.Msg) (SelectModel, tea.Cmd) {
	if !m.Focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Selected > 0 {
				m.Selected--
				m.ensureVisible()
			}
		case "down", "j":
			if m.Selected < len(m.Items)-1 {
				m.Selected++
				m.ensureVisible()
			}
		case "home", "g":
			m.Selected = 0
			m.offset = 0
		case "end", "G":
			m.Selected = len(m.Items) - 1
			m.ensureVisible()
		case "pgup":
			m.Selected -= m.Height
			if m.Selected < 0 {
				m.Selected = 0
			}
			m.ensureVisible()
		case "pgdown":
			m.Selected += m.Height
			if m.Selected >= len(m.Items) {
				m.Selected = len(m.Items) - 1
			}
			m.ensureVisible()
		}
	}

	return m, nil
}

// View renders the select.
func (m SelectModel) View() string {
	var b strings.Builder

	// Label
	if m.Label != "" {
		labelStyle := lipgloss.NewStyle().Bold(true)
		if m.Theme != nil {
			labelStyle = labelStyle.Foreground(lipgloss.Color(m.Theme.Colors.Text))
		}
		b.WriteString(labelStyle.Render(m.Label))
		b.WriteString("\n\n")
	}

	// Items
	for i := m.offset; i < len(m.Items) && i < m.offset+m.Height; i++ {
		item := m.Items[i]
		isSelected := i == m.Selected

		// Build item line
		var marker string
		if isSelected {
			marker = "▸ "
		} else {
			marker = "  "
		}

		var line string
		if m.Theme != nil {
			if isSelected {
				line = m.Theme.Component.ListItemSelected.Render(marker + item.Label)
			} else {
				line = m.Theme.Component.ListItem.Render(marker + item.Label)
			}
		} else {
			line = marker + item.Label
		}

		b.WriteString(line)

		// Description
		if m.showDescriptions && item.Description != "" {
			if m.Theme != nil {
				b.WriteString(" ")
				b.WriteString(m.Theme.Typography.Muted.Render("- " + item.Description))
			} else {
				b.WriteString(" - " + item.Description)
			}
		}

		b.WriteString("\n")
	}

	// Scroll indicator
	if len(m.Items) > m.Height {
		b.WriteString("\n")
		indicator := tui.ScrollIndicator(len(m.Items), m.Height, m.offset)
		if m.Theme != nil {
			b.WriteString(m.Theme.Typography.Muted.Render(indicator))
		} else {
			b.WriteString(indicator)
		}
	}

	return b.String()
}

// ensureVisible ensures the selected item is visible.
func (m *SelectModel) ensureVisible() {
	if m.Selected < m.offset {
		m.offset = m.Selected
	} else if m.Selected >= m.offset+m.Height {
		m.offset = m.Selected - m.Height + 1
	}
}

// SelectedItem returns the currently selected item.
func (m SelectModel) SelectedItem() SelectItem {
	if m.Selected >= 0 && m.Selected < len(m.Items) {
		return m.Items[m.Selected]
	}
	return SelectItem{}
}

// SelectedValue returns the value of the selected item.
func (m SelectModel) SelectedValue() string {
	return m.SelectedItem().Value
}

// SelectedLabel returns the label of the selected item.
func (m SelectModel) SelectedLabel() string {
	return m.SelectedItem().Label
}

// SetItems sets the items list.
func (m *SelectModel) SetItems(items []SelectItem) {
	m.Items = items
	if m.Selected >= len(items) {
		m.Selected = len(items) - 1
	}
	if m.Selected < 0 {
		m.Selected = 0
	}
}

// SetSelected sets the selected index.
func (m *SelectModel) SetSelected(index int) {
	if index >= 0 && index < len(m.Items) {
		m.Selected = index
		m.ensureVisible()
	}
}

// SetTheme sets the theme.
func (m *SelectModel) SetTheme(theme *styles.Theme) {
	m.Theme = theme
}

// SetWidth sets the width.
func (m *SelectModel) SetWidth(width int) {
	m.Width = width
}

// SetHeight sets the visible height.
func (m *SelectModel) SetHeight(height int) {
	m.Height = height
}

// SetFocused sets the focus state.
func (m *SelectModel) SetFocused(focused bool) {
	m.Focused = focused
}

// SetLabel sets the label.
func (m *SelectModel) SetLabel(label string) {
	m.Label = label
}

// Blur removes focus.
func (m *SelectModel) Blur() {
	m.Focused = false
}

// Focus gives focus.
func (m *SelectModel) Focus() {
	m.Focused = true
}
