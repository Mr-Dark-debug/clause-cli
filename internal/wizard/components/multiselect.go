package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
)

// MultiSelectItem represents an item in a multi-select list.
type MultiSelectItem struct {
	// Label is the displayed text
	Label string

	// Value is the underlying value
	Value string

	// Selected indicates if the item is selected
	Selected bool

	// Description is an optional description
	Description string

	// Disabled indicates if the item can be toggled
	Disabled bool
}

// MultiSelectModel is a multiple-selection list component.
type MultiSelectModel struct {
	// Items is the list of selectable items
	Items []MultiSelectItem

	// Cursor is the current cursor position
	Cursor int

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

	// Min selections required
	MinSelections int

	// Max selections allowed (0 = unlimited)
	MaxSelections int
}

// NewMultiSelect creates a new multi-select component.
func NewMultiSelect(items []MultiSelectItem) MultiSelectModel {
	return MultiSelectModel{
		Items:    items,
		Cursor:   0,
		Width:    40,
		Height:   10,
		Focused:  true,
		offset:   0,
	}
}

// Init initializes the multi-select.
func (m MultiSelectModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the multi-select.
func (m MultiSelectModel) Update(msg tea.Msg) (MultiSelectModel, tea.Cmd) {
	if !m.Focused {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
				m.ensureVisible()
			}
		case "down", "j":
			if m.Cursor < len(m.Items)-1 {
				m.Cursor++
				m.ensureVisible()
			}
		case " ", "x":
			m.toggleCurrent()
		case "a":
			m.toggleAll()
		case "home", "g":
			m.Cursor = 0
			m.offset = 0
		case "end", "G":
			m.Cursor = len(m.Items) - 1
			m.ensureVisible()
		}
	}

	return m, nil
}

// View renders the multi-select.
func (m MultiSelectModel) View() string {
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
		isCursor := i == m.Cursor

		// Checkbox
		var checkbox string
		if item.Selected {
			if m.Theme != nil {
				checkbox = m.Theme.Component.CheckboxChecked.Render("✓")
			} else {
				checkbox = "[✓]"
			}
		} else {
			if m.Theme != nil {
				checkbox = m.Theme.Component.Checkbox.Render("○")
			} else {
				checkbox = "[ ]"
			}
		}

		// Build line
		var line strings.Builder
		if isCursor {
			line.WriteString("▸ ")
		} else {
			line.WriteString("  ")
		}
		line.WriteString(checkbox)
		line.WriteString(" ")
		line.WriteString(item.Label)

		if item.Description != "" {
			line.WriteString(" ")
			if m.Theme != nil {
				line.WriteString(m.Theme.Typography.Muted.Render("- " + item.Description))
			} else {
				line.WriteString("- " + item.Description)
			}
		}

		rendered := line.String()
		if m.Theme != nil && isCursor {
			rendered = m.Theme.Component.ListItemSelected.Render(rendered)
		}

		b.WriteString(rendered)
		b.WriteString("\n")
	}

	// Help text
	if m.Theme != nil {
		helpStyle := m.Theme.Typography.Muted
		b.WriteString("\n")
		b.WriteString(helpStyle.Render("[Space] Toggle  [a] Toggle all"))
	}

	return b.String()
}

// ensureVisible ensures the cursor is visible.
func (m *MultiSelectModel) ensureVisible() {
	if m.Cursor < m.offset {
		m.offset = m.Cursor
	} else if m.Cursor >= m.offset+m.Height {
		m.offset = m.Cursor - m.Height + 1
	}
}

// toggleCurrent toggles the current item.
func (m *MultiSelectModel) toggleCurrent() {
	if m.Cursor < 0 || m.Cursor >= len(m.Items) {
		return
	}

	item := &m.Items[m.Cursor]
	if item.Disabled {
		return
	}

	// Check max selections
	if !item.Selected && m.MaxSelections > 0 {
		if m.countSelected() >= m.MaxSelections {
			return
		}
	}

	item.Selected = !item.Selected
}

// toggleAll toggles all items.
func (m *MultiSelectModel) toggleAll() {
	allSelected := true
	for _, item := range m.Items {
		if !item.Disabled && !item.Selected {
			allSelected = false
			break
		}
	}

	// Toggle based on current state
	for i := range m.Items {
		if !m.Items[i].Disabled {
			if allSelected {
				m.Items[i].Selected = false
			} else {
				if m.MaxSelections == 0 || m.countSelected() < m.MaxSelections {
					m.Items[i].Selected = true
				}
			}
		}
	}
}

// countSelected returns the number of selected items.
func (m MultiSelectModel) countSelected() int {
	count := 0
	for _, item := range m.Items {
		if item.Selected {
			count++
		}
	}
	return count
}

// SelectedItems returns all selected items.
func (m MultiSelectModel) SelectedItems() []MultiSelectItem {
	var selected []MultiSelectItem
	for _, item := range m.Items {
		if item.Selected {
			selected = append(selected, item)
		}
	}
	return selected
}

// SelectedValues returns the values of all selected items.
func (m MultiSelectModel) SelectedValues() []string {
	var values []string
	for _, item := range m.Items {
		if item.Selected {
			values = append(values, item.Value)
		}
	}
	return values
}

// SelectedLabels returns the labels of all selected items.
func (m MultiSelectModel) SelectedLabels() []string {
	var labels []string
	for _, item := range m.Items {
		if item.Selected {
			labels = append(labels, item.Label)
		}
	}
	return labels
}

// SetItems sets the items list.
func (m *MultiSelectModel) SetItems(items []MultiSelectItem) {
	m.Items = items
	if m.Cursor >= len(items) {
		m.Cursor = len(items) - 1
	}
	if m.Cursor < 0 {
		m.Cursor = 0
	}
}

// SetTheme sets the theme.
func (m *MultiSelectModel) SetTheme(theme *styles.Theme) {
	m.Theme = theme
}

// SetWidth sets the width.
func (m *MultiSelectModel) SetWidth(width int) {
	m.Width = width
}

// SetHeight sets the visible height.
func (m *MultiSelectModel) SetHeight(height int) {
	m.Height = height
}

// SetFocused sets the focus state.
func (m *MultiSelectModel) SetFocused(focused bool) {
	m.Focused = focused
}

// SetLabel sets the label.
func (m *MultiSelectModel) SetLabel(label string) {
	m.Label = label
}

// Blur removes focus.
func (m *MultiSelectModel) Blur() {
	m.Focused = false
}

// Focus gives focus.
func (m *MultiSelectModel) Focus() {
	m.Focused = true
}

// IsValid returns true if selection constraints are met.
func (m MultiSelectModel) IsValid() bool {
	count := m.countSelected()
	return count >= m.MinSelections && (m.MaxSelections == 0 || count <= m.MaxSelections)
}
