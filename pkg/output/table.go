package output

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/utils"
)

// TableColumn represents a table column.
type TableColumn struct {
	Title     string
	Width     int
	Alignment lipgloss.Position
	MaxWidth  int
}

// TableRow represents a table row.
type TableRow struct {
	Cells   []string
	IsTitle bool
}

// Table provides styled table rendering.
type Table struct {
	columns   []TableColumn
	rows      []TableRow
	theme     *styles.Theme
	writer    io.Writer
	style     TableStyle
	width     int
	showBorder bool
	showHeader bool
	compact   bool
}

// TableStyle holds table styling options.
type TableStyle struct {
	BorderStyle     lipgloss.Border
	BorderColor     string
	HeaderStyle     lipgloss.Style
	RowStyle        lipgloss.Style
	AltRowStyle     lipgloss.Style
	CellStyle       lipgloss.Style
	TitleStyle      lipgloss.Style
	SeparatorStyle  lipgloss.Style
}

// NewTable creates a new table.
func NewTable(columns []TableColumn, opts ...TableOption) *Table {
	t := &Table{
		columns:    columns,
		theme:      styles.GetTheme(),
		writer:     os.Stdout,
		showBorder: true,
		showHeader: true,
	}

	// Apply default styles
	t.style = TableStyle{
		BorderColor: t.theme.Colors.Border,
		HeaderStyle: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.theme.Colors.Text)),
		RowStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.theme.Colors.Text)),
		AltRowStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.theme.Colors.Text)),
		CellStyle: lipgloss.NewStyle().
			Padding(0, 1),
		TitleStyle: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.theme.Colors.Primary)).
			Padding(0, 1),
		SeparatorStyle: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.theme.Colors.BorderMuted)),
	}

	for _, opt := range opts {
		opt(t)
	}

	return t
}

// TableOption is a functional option for table configuration.
type TableOption func(*Table)

// WithTableTheme sets the table theme.
func WithTableTheme(theme *styles.Theme) TableOption {
	return func(t *Table) {
		t.theme = theme
	}
}

// WithTableWriter sets the output writer.
func WithTableWriter(w io.Writer) TableOption {
	return func(t *Table) {
		t.writer = w
	}
}

// WithTableWidth sets the table width.
func WithTableWidth(width int) TableOption {
	return func(t *Table) {
		t.width = width
	}
}

// WithTableBorder enables or disables borders.
func WithTableBorder(show bool) TableOption {
	return func(t *Table) {
		t.showBorder = show
	}
}

// WithTableHeader enables or disables the header.
func WithTableHeader(show bool) TableOption {
	return func(t *Table) {
		t.showHeader = show
	}
}

// WithCompact enables compact mode.
func WithCompact(compact bool) TableOption {
	return func(t *Table) {
		t.compact = compact
	}
}

// AddRow adds a row to the table.
func (t *Table) AddRow(cells ...string) *Table {
	t.rows = append(t.rows, TableRow{Cells: cells})
	return t
}

// AddTitle adds a title row.
func (t *Table) AddTitle(title string) *Table {
	t.rows = append(t.rows, TableRow{Cells: []string{title}, IsTitle: true})
	return t
}

// calculateWidths calculates column widths.
func (t *Table) calculateWidths() []int {
	widths := make([]int, len(t.columns))

	// Start with column specifications
	for i, col := range t.columns {
		widths[i] = col.Width
	}

	// Expand for content
	for _, row := range t.rows {
		for i, cell := range row.Cells {
			if i < len(widths) {
				cellWidth := lipgloss.Width(cell)
				if cellWidth > widths[i] {
					maxWidth := t.columns[i].MaxWidth
					if maxWidth > 0 && cellWidth > maxWidth {
						cellWidth = maxWidth
					}
					if cellWidth > widths[i] {
						widths[i] = cellWidth
					}
				}
			}
		}
	}

	// Respect total width if specified
	if t.width > 0 {
		totalWidth := 0
		for _, w := range widths {
			totalWidth += w
		}
		// Add padding and separators
		totalWidth += len(widths) * 2 + len(widths) - 1

		if totalWidth > t.width {
			// Scale down proportionally
			scale := float64(t.width-len(widths)-1) / float64(totalWidth-len(widths)-1)
			for i := range widths {
				widths[i] = int(float64(widths[i]) * scale)
			}
		}
	}

	return widths
}

// Render renders the table as a string.
func (t *Table) Render() string {
	widths := t.calculateWidths()
	var lines []string

	// Top border
	if t.showBorder {
		lines = append(lines, t.renderBorder("top", widths))
	}

	// Header
	if t.showHeader {
		lines = append(lines, t.renderHeader(widths))
		if t.showBorder {
			lines = append(lines, t.renderBorder("separator", widths))
		}
	}

	// Rows
	for i, row := range t.rows {
		if row.IsTitle {
			lines = append(lines, t.renderTitleRow(row, widths))
		} else {
			lines = append(lines, t.renderRow(row, widths, i))
		}
	}

	// Bottom border
	if t.showBorder {
		lines = append(lines, t.renderBorder("bottom", widths))
	}

	return strings.Join(lines, "\n")
}

// renderBorder renders a border line.
func (t *Table) renderBorder(kind string, widths []int) string {
	var left, middle, right, horiz string

	if t.compact {
		switch kind {
		case "top":
			left, middle, right, horiz = "┌", "┬", "┐", "─"
		case "bottom":
			left, middle, right, horiz = "└", "┴", "┘", "─"
		case "separator":
			left, middle, right, horiz = "├", "┼", "┤", "─"
		}
	} else {
		switch kind {
		case "top":
			left, middle, right, horiz = "╭", "┬", "╮", "─"
		case "bottom":
			left, middle, right, horiz = "╰", "┴", "╯", "─"
		case "separator":
			left, middle, right, horiz = "├", "┼", "┤", "─"
		}
	}

	parts := []string{left}
	for i, w := range widths {
		if i > 0 {
			parts = append(parts, middle)
		}
		parts = append(parts, strings.Repeat(horiz, w+2))
	}
	parts = append(parts, right)

	style := lipgloss.NewStyle().Foreground(lipgloss.Color(t.style.BorderColor))
	return style.Render(strings.Join(parts, ""))
}

// renderHeader renders the header row.
func (t *Table) renderHeader(widths []int) string {
	cells := make([]string, len(t.columns))
	for i, col := range t.columns {
		cell := t.style.CellStyle.
			Width(widths[i]).
			Align(col.Alignment).
			Render(col.Title)
		cells[i] = t.style.HeaderStyle.Render(cell)
	}

	vertical := "│"
	if t.compact {
		vertical = "│"
	}

	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(t.style.BorderColor))
	return borderStyle.Render(vertical) +
		strings.Join(cells, borderStyle.Render(vertical)) +
		borderStyle.Render(vertical)
}

// renderRow renders a data row.
func (t *Table) renderRow(row TableRow, widths []int, rowIndex int) string {
	cells := make([]string, len(widths))
	for i := range widths {
		var content string
		if i < len(row.Cells) {
			content = row.Cells[i]
		}

		// Truncate if necessary
		if len(content) > widths[i] {
			content = utils.TruncateText(content, widths[i])
		}

		alignment := lipgloss.Left
		if i < len(t.columns) {
			alignment = t.columns[i].Alignment
		}

		cells[i] = t.style.CellStyle.
			Width(widths[i]).
			Align(alignment).
			Render(content)
	}

	// Apply alternating row style
	var rowStyle lipgloss.Style
	if rowIndex%2 == 1 {
		rowStyle = t.style.AltRowStyle
	} else {
		rowStyle = t.style.RowStyle
	}

	for i, cell := range cells {
		cells[i] = rowStyle.Render(cell)
	}

	vertical := "│"
	if t.compact {
		vertical = "│"
	}

	borderStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(t.style.BorderColor))
	return borderStyle.Render(vertical) +
		strings.Join(cells, borderStyle.Render(vertical)) +
		borderStyle.Render(vertical)
}

// renderTitleRow renders a title row.
func (t *Table) renderTitleRow(row TableRow, widths []int) string {
	if len(row.Cells) == 0 {
		return ""
	}

	// Calculate total width
	totalWidth := 0
	for _, w := range widths {
		totalWidth += w + 2
	}
	totalWidth += len(widths) - 1

	return t.style.TitleStyle.
		Width(totalWidth).
		Align(lipgloss.Center).
		Render(row.Cells[0])
}

// Print prints the table.
func (t *Table) Print() {
	fmt.Fprintln(t.writer, t.Render())
}

// String implements fmt.Stringer.
func (t *Table) String() string {
	return t.Render()
}

// SimpleTable creates a simple table with default settings.
func SimpleTable(headers []string, rows [][]string) string {
	columns := make([]TableColumn, len(headers))
	for i, h := range headers {
		columns[i] = TableColumn{
			Title:     h,
			Width:     len(h),
			Alignment: lipgloss.Left,
		}
	}

	table := NewTable(columns, WithTableBorder(false))
	for _, row := range rows {
		table.AddRow(row...)
	}

	return table.Render()
}

// KeyValueTable creates a key-value table.
func KeyValueTable(pairs [][2]string, keyWidth int) string {
	columns := []TableColumn{
		{Title: "Key", Width: keyWidth, Alignment: lipgloss.Right},
		{Title: "Value", Width: 40, Alignment: lipgloss.Left},
	}

	table := NewTable(columns,
		WithTableBorder(false),
		WithTableHeader(false),
	)

	for _, pair := range pairs {
		table.AddRow(pair[0], pair[1])
	}

	return table.Render()
}

// DefinitionList creates a definition list (term: definition).
func DefinitionList(items []struct {
	Term       string
	Definition string
}, termWidth int) string {
	theme := styles.GetTheme()

	var lines []string
	termStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(theme.Colors.Text)).
		Width(termWidth)

	defStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Text))

	for _, item := range items {
		line := termStyle.Render(item.Term) + defStyle.Render(item.Definition)
		lines = append(lines, line)
	}

	return strings.Join(lines, "\n")
}

// ListTable creates a simple list with optional numbering.
func ListTable(items []string, numbered bool) string {
	theme := styles.GetTheme()

	var lines []string
	bulletStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Primary))
	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Text))

	for i, item := range items {
		var prefix string
		if numbered {
			prefix = bulletStyle.Render(fmt.Sprintf("%d. ", i+1))
		} else {
			prefix = bulletStyle.Render("• ")
		}
		lines = append(lines, prefix+textStyle.Render(item))
	}

	return strings.Join(lines, "\n")
}

// ColumnLayout creates a multi-column layout.
func ColumnLayout(items []string, columns int, width int) string {
	if columns < 1 {
		columns = 1
	}

	colWidth := width / columns
	if colWidth < 10 {
		colWidth = 10
	}

	theme := styles.GetTheme()
	style := lipgloss.NewStyle().
		Foreground(lipgloss.Color(theme.Colors.Text)).
		Width(colWidth)

	var lines []string
	var currentLine []string

	for _, item := range items {
		if len(currentLine) >= columns {
			lines = append(lines, strings.Join(currentLine, ""))
			currentLine = nil
		}

		truncated := item
		if len(item) > colWidth {
			truncated = utils.TruncateText(item, colWidth-2)
		}
		currentLine = append(currentLine, style.Render(truncated))
	}

	if len(currentLine) > 0 {
		// Pad remaining columns
		for len(currentLine) < columns {
			currentLine = append(currentLine, style.Render(""))
		}
		lines = append(lines, strings.Join(currentLine, ""))
	}

	return strings.Join(lines, "\n")
}
