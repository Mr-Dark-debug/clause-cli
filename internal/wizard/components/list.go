package components

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
	"github.com/clause-cli/clause/pkg/utils"
)

// ListItem represents an item in a list.
type ListItem struct {
	// ID is the item identifier
	ID string

	// Title is the main title
	Title string

	// Description is an optional description
	Description string

	// Filterable text (searches this)
	FilterValue string
}

// ListModel is a scrollable list with filtering.
type ListModel struct {
	// Items is the list of items
	Items []ListItem

	// Filtered items (after applying filter)
	filtered []ListItem

	// Selected index (in filtered list)
	Selected int

	// Filter is the current filter text
	Filter string

	// Width is the list width
	Width int

	// Height is the visible height
	Height int

	// Focused indicates if the list has focus
	Focused bool

	// Theme is the current theme
	Theme *styles.Theme

	// Title for the list
	Title string

	// ShowFilter indicates if filter input should be shown
	ShowFilter bool

	// ShowHelp indicates if help should be shown
	ShowHelp bool

	// Scroll offset
	offset int
}

// NewList creates a new list component.
func NewList(items []ListItem) ListModel {
	l := ListModel{
		Items:      items,
		filtered:   items,
		Selected:   0,
		Width:      60,
		Height:     10,
		Focused:    true,
		ShowFilter: true,
		ShowHelp:   true,
		offset:     0,
	}
	return l
}

// Init initializes the list.
func (m ListModel) Init() tea.Cmd {
	return nil
}

// Update handles updates for the list.
func (m ListModel) Update(msg tea.Msg) (ListModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if !m.Focused {
			return m, nil
		}

		switch msg.String() {
		case "up", "k":
			if m.Selected > 0 {
				m.Selected--
				m.ensureVisible()
			}
		case "down", "j":
			if m.Selected < len(m.filtered)-1 {
				m.Selected++
				m.ensureVisible()
			}
		case "pgup":
			m.Selected -= m.Height
			if m.Selected < 0 {
				m.Selected = 0
			}
			m.ensureVisible()
		case "pgdown":
			m.Selected += m.Height
			if m.Selected >= len(m.filtered) {
				m.Selected = len(m.filtered) - 1
			}
			m.ensureVisible()
		case "home", "g":
			m.Selected = 0
			m.offset = 0
		case "end", "G":
			m.Selected = len(m.filtered) - 1
			m.ensureVisible()
		case "/":
			if m.ShowFilter {
				// Focus filter input
			}
		case "backspace":
			if m.Filter != "" {
				m.Filter = m.Filter[:len(m.Filter)-1]
				m.applyFilter()
			}
		case "esc":
			if m.Filter != "" {
				m.Filter = ""
				m.applyFilter()
			}
		default:
			// Handle text input for filter
			if m.ShowFilter && len(msg.Runes) > 0 {
				char := string(msg.Runes)
				if char != "" && len(char) == 1 {
					m.Filter += char
					m.applyFilter()
				}
			}
		}
	}

	return m, nil
}

// View renders the list.
func (m ListModel) View() string {
	var b strings.Builder

	// Title
	if m.Title != "" {
		titleStyle := lipgloss.NewStyle().Bold(true)
		if m.Theme != nil {
			titleStyle = titleStyle.Foreground(lipgloss.Color(m.Theme.Colors.Primary))
		}
		b.WriteString(titleStyle.Render(m.Title))
		b.WriteString("\n")
	}

	// Filter
	if m.ShowFilter {
		filterStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			Padding(0, 1)

		filterText := "Filter: " + m.Filter
		if m.Filter == "" {
			filterText = "Filter: (press / to search)"
		}

		b.WriteString(filterStyle.Render(filterText))
		b.WriteString("\n\n")
	}

	// Empty state
	if len(m.filtered) == 0 {
		emptyStyle := lipgloss.NewStyle().Faint(true)
		if m.Theme != nil {
			emptyStyle = m.Theme.Typography.Muted
		}
		b.WriteString(emptyStyle.Render("No items found"))
		return b.String()
	}

	// Items
	for i := m.offset; i < len(m.filtered) && i < m.offset+m.Height; i++ {
		item := m.filtered[i]
		isSelected := i == m.Selected

		// Item marker
		marker := "  "
		if isSelected {
			marker = "▸ "
		}

		// Title
		var line string
		if isSelected {
			if m.Theme != nil {
				line = m.Theme.Component.ListItemSelected.Render(marker + item.Title)
			} else {
				line = marker + item.Title
			}
		} else {
			if m.Theme != nil {
				line = m.Theme.Component.ListItem.Render(marker + item.Title)
			} else {
				line = marker + item.Title
			}
		}

		b.WriteString(line)

		// Description
		if item.Description != "" {
			b.WriteString(" ")
			if m.Theme != nil {
				b.WriteString(m.Theme.Typography.Muted.Render("- " + item.Description))
			} else {
				b.WriteString("- " + item.Description)
			}
		}

		b.WriteString("\n")
	}

	// Scroll indicator
	if len(m.filtered) > m.Height {
		b.WriteString("\n")
		indicator := tui.ScrollIndicator(len(m.filtered), m.Height, m.offset)
		if m.Theme != nil {
			b.WriteString(m.Theme.Typography.Muted.Render(indicator))
		} else {
			b.WriteString(indicator)
		}
		b.WriteString("\n")
	}

	// Help
	if m.ShowHelp {
		b.WriteString("\n")
		kb := tui.NewKeyBindings()
		kb.Add("↑/↓", "Navigate")
		if m.ShowFilter {
			kb.Add("/", "Filter")
			kb.Add("Esc", "Clear filter")
		}
		if m.Theme != nil {
			b.WriteString(m.Theme.Typography.Muted.Render(kb.Help()))
		} else {
			b.WriteString(kb.Help())
		}
	}

	return b.String()
}

// applyFilter applies the current filter to items.
func (m *ListModel) applyFilter() {
	if m.Filter == "" {
		m.filtered = m.Items
		m.Selected = 0
		m.offset = 0
		return
	}

	filter := strings.ToLower(m.Filter)
	m.filtered = nil

	for _, item := range m.Items {
		searchText := strings.ToLower(item.Title + " " + item.Description + " " + item.FilterValue)
		if strings.Contains(searchText, filter) {
			m.filtered = append(m.filtered, item)
		}
	}

	m.Selected = 0
	m.offset = 0
}

// ensureVisible ensures the selected item is visible.
func (m *ListModel) ensureVisible() {
	if m.Selected < m.offset {
		m.offset = m.Selected
	} else if m.Selected >= m.offset+m.Height {
		m.offset = m.Selected - m.Height + 1
	}
}

// SelectedItem returns the currently selected item.
func (m ListModel) SelectedItem() ListItem {
	if m.Selected >= 0 && m.Selected < len(m.filtered) {
		return m.filtered[m.Selected]
	}
	return ListItem{}
}

// SelectedID returns the ID of the selected item.
func (m ListModel) SelectedID() string {
	return m.SelectedItem().ID
}

// SetItems sets the items list.
func (m *ListModel) SetItems(items []ListItem) {
	m.Items = items
	m.applyFilter()
}

// SetFilter sets the filter text.
func (m *ListModel) SetFilter(filter string) {
	m.Filter = filter
	m.applyFilter()
}

// SetTheme sets the theme.
func (m *ListModel) SetTheme(theme *styles.Theme) {
	m.Theme = theme
}

// SetWidth sets the width.
func (m *ListModel) SetWidth(width int) {
	m.Width = width
}

// SetHeight sets the visible height.
func (m *ListModel) SetHeight(height int) {
	m.Height = height
}

// SetFocused sets the focus state.
func (m *ListModel) SetFocused(focused bool) {
	m.Focused = focused
}

// SetTitle sets the title.
func (m *ListModel) SetTitle(title string) {
	m.Title = title
}

// Blur removes focus.
func (m *ListModel) Blur() {
	m.Focused = false
}

// Focus gives focus.
func (m *ListModel) Focus() {
	m.Focused = true
}

// FilteredCount returns the number of items after filtering.
func (m ListModel) FilteredCount() int {
	return len(m.filtered)
}

// TotalCount returns the total number of items.
func (m ListModel) TotalCount() int {
	return len(m.Items)
}

// Truncate text helper
var truncateText = utils.TruncateText
