package styles

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Spacing constants for consistent layouts.
const (
	// Spacing units based on 4px grid.
	SpacingNone  = 0
	SpacingXs    = 1
	SpacingSm    = 2
	SpacingMd    = 4
	SpacingLg    = 8
	SpacingXl    = 12
	Spacing2xl   = 16
	Spacing3xl   = 24
	Spacing4xl   = 32

	// Common padding values.
	PaddingNone = 0
	PaddingSm   = 1
	PaddingMd   = 2
	PaddingLg   = 4

	// Common margin values.
	MarginNone = 0
	MarginSm   = 1
	MarginMd   = 2
	MarginLg   = 4
)

// Layout provides layout utilities for terminal UI.
type Layout struct {
	theme   *Theme
	width   int
	height  int
	breakpoint Breakpoint
}

// Breakpoint represents a responsive breakpoint.
type Breakpoint int

const (
	// BreakpointCompact is for narrow terminals (<80 columns).
	BreakpointCompact Breakpoint = iota
	// BreakpointStandard is for normal terminals (80-120 columns).
	BreakpointStandard
	// BreakpointWide is for wide terminals (>120 columns).
	BreakpointWide
)

// NewLayout creates a new Layout instance.
func NewLayout(theme *Theme, width, height int) *Layout {
	if theme == nil {
		theme = DefaultTheme
	}

	return &Layout{
		theme:      theme,
		width:      width,
		height:     height,
		breakpoint: CalculateBreakpoint(width),
	}
}

// CalculateBreakpoint determines the breakpoint based on width.
func CalculateBreakpoint(width int) Breakpoint {
	switch {
	case width < 80:
		return BreakpointCompact
	case width < 120:
		return BreakpointStandard
	default:
		return BreakpointWide
	}
}

// Width returns the current layout width.
func (l *Layout) Width() int {
	return l.width
}

// Height returns the current layout height.
func (l *Layout) Height() int {
	return l.height
}

// Breakpoint returns the current breakpoint.
func (l *Layout) Breakpoint() Breakpoint {
	return l.breakpoint
}

// IsCompact returns true if the layout is in compact mode.
func (l *Layout) IsCompact() bool {
	return l.breakpoint == BreakpointCompact
}

// IsWide returns true if the layout is in wide mode.
func (l *Layout) IsWide() bool {
	return l.breakpoint == BreakpointWide
}

// Container creates a container style with appropriate padding.
func (l *Layout) Container() lipgloss.Style {
	padding := PaddingMd
	if l.IsCompact() {
		padding = PaddingSm
	}

	return lipgloss.NewStyle().
		Width(l.width).
		Padding(padding, padding)
}

// Card creates a card style with border and background.
func (l *Layout) Card() lipgloss.Style {
	return l.theme.Layout.Card
}

// Section creates spacing between sections.
func (l *Layout) Section() lipgloss.Style {
	return lipgloss.NewStyle().Margin(SpacingMd, 0)
}

// Spacer creates a vertical spacer.
func (l *Layout) Spacer(height int) lipgloss.Style {
	return lipgloss.NewStyle().Height(height)
}

// Divider creates a horizontal divider line.
func (l *Layout) Divider() string {
	divChar := "─"
	if l.IsCompact() {
		divChar = "-"
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(l.theme.Colors.BorderMuted)).
		Render(strings.Repeat(divChar, l.width))
}

// DoubleDivider creates a double horizontal divider line.
func (l *Layout) DoubleDivider() string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(l.theme.Colors.Border)).
		Render(strings.Repeat("═", l.width))
}

// BorderBox creates a bordered box with title and content.
func (l *Layout) BorderBox(title, content string, width int) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(l.theme.Colors.Border)).
		Padding(PaddingSm, PaddingMd).
		Width(width)

	if title != "" {
		titleStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(l.theme.Colors.Primary)).
			PaddingBottom(PaddingSm)

		return lipgloss.JoinVertical(
			lipgloss.Left,
			titleStyle.Render(title),
			boxStyle.Render(content),
		)
	}

	return boxStyle.Render(content)
}

// Panel creates a panel with a header and body.
func (l *Layout) Panel(header, body string, width int) string {
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(l.theme.Colors.Text)).
		Background(lipgloss.Color(l.theme.Colors.BackgroundCard)).
		Padding(PaddingSm, PaddingMd).
		Width(width)

	bodyStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder(), false, true, false, true).
		BorderForeground(lipgloss.Color(l.theme.Colors.Border)).
		Padding(PaddingMd).
		Width(width)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		headerStyle.Render(header),
		bodyStyle.Render(body),
	)
}

// Columns creates a multi-column layout.
func (l *Layout) Columns(items []string, gap int) string {
	if len(items) == 0 {
		return ""
	}

	availableWidth := l.width - (gap * (len(items) - 1))
	colWidth := availableWidth / len(items)

	columns := make([]string, len(items))
	for i, item := range items {
		columns[i] = lipgloss.NewStyle().
			Width(colWidth).
			Render(item)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, columns...)
}

// TwoColumn creates a two-column layout with specified widths.
func (l *Layout) TwoColumn(left, right string, leftWidth int) string {
	rightWidth := l.width - leftWidth - SpacingMd

	leftStyle := lipgloss.NewStyle().Width(leftWidth)
	rightStyle := lipgloss.NewStyle().Width(rightWidth)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		leftStyle.Render(left),
		rightStyle.Render(right),
	)
}

// ThreeColumn creates a three-column layout.
func (l *Layout) ThreeColumn(left, center, right string) string {
	availableWidth := l.width - (SpacingMd * 2)
	colWidth := availableWidth / 3

	style := lipgloss.NewStyle().Width(colWidth)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		style.Render(left),
		style.Render(center),
		style.Render(right),
	)
}

// Grid creates a grid layout with specified columns.
func (l *Layout) Grid(items []string, columns int, gap int) string {
	if len(items) == 0 || columns <= 0 {
		return ""
	}

	availableWidth := l.width - (gap * (columns - 1))
	colWidth := availableWidth / columns

	colStyle := lipgloss.NewStyle().Width(colWidth)

	rows := make([]string, 0)
	currentRow := make([]string, 0, columns)

	for _, item := range items {
		currentRow = append(currentRow, colStyle.Render(item))

		if len(currentRow) == columns {
			rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, currentRow...))
			currentRow = currentRow[:0]
		}
	}

	// Handle remaining items
	if len(currentRow) > 0 {
		// Pad remaining cells
		for len(currentRow) < columns {
			currentRow = append(currentRow, colStyle.Render(""))
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Top, currentRow...))
	}

	return lipgloss.JoinVertical(lipgloss.Left, rows...)
}

// List creates a vertical list with bullets.
func (l *Layout) List(items []string, bullet string) string {
	if len(items) == 0 {
		return ""
	}

	bulletStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(l.theme.Colors.Primary)).
		PaddingRight(PaddingSm)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(l.theme.Colors.Text))

	var lines []string
	for _, item := range items {
		line := lipgloss.JoinHorizontal(
			lipgloss.Top,
			bulletStyle.Render(bullet),
			textStyle.Render(item),
		)
		lines = append(lines, line)
	}

	return lipgloss.JoinVertical(lipgloss.Left, lines...)
}

// KeyedListItem creates a list item with a key/value format.
func (l *Layout) KeyedListItem(key, value string, keyWidth int) string {
	keyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(l.theme.Colors.TextMuted)).
		Width(keyWidth).
		PaddingRight(PaddingMd)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(l.theme.Colors.Text))

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		keyStyle.Render(key),
		valueStyle.Render(value),
	)
}

// Box creates a simple box with content.
func Box(content string, width, height int) string {
	return lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(PaddingMd).
		Width(width).
		Height(height).
		Render(content)
}

// HorizontalBox creates a horizontal arrangement of boxes.
func HorizontalBox(items []string, gap int) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, items...)
}

// VerticalBox creates a vertical arrangement of boxes.
func VerticalBox(items []string, gap int) string {
	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// Padding creates padding around content.
func Padding(content string, top, right, bottom, left int) string {
	return lipgloss.NewStyle().
		Padding(top, right, bottom, left).
		Render(content)
}

// Margin creates margin around content.
func Margin(content string, top, right, bottom, left int) string {
	return lipgloss.NewStyle().
		Margin(top, right, bottom, left).
		Render(content)
}

// BorderStyle represents different border styles.
type BorderStyle int

const (
	// BorderNone means no border.
	BorderNone BorderStyle = iota
	// BorderSingle is a single-line border.
	BorderSingle
	// BorderDouble is a double-line border.
	BorderDouble
	// BorderRounded is a rounded border.
	BorderRounded
	// BorderThick is a thick border.
	BorderThick
)

// ApplyBorder applies a border style to content.
func ApplyBorder(content string, style BorderStyle, color string, width int) string {
	var border lipgloss.Border

	switch style {
	case BorderSingle:
		border = lipgloss.NormalBorder()
	case BorderDouble:
		border = lipgloss.DoubleBorder()
	case BorderRounded:
		border = lipgloss.RoundedBorder()
	case BorderThick:
		border = lipgloss.ThickBorder()
	default:
		return content
	}

	return lipgloss.NewStyle().
		Border(border).
		BorderForeground(lipgloss.Color(color)).
		Padding(PaddingSm, PaddingMd).
		Width(width).
		Render(content)
}

// FlexItem represents an item in a flex layout.
type FlexItem struct {
	Content string
	Grow    float64
	Shrink  float64
	Basis   int
}

// FlexRow creates a horizontal flex layout.
func FlexRow(items []FlexItem, width int, gap int) string {
	if len(items) == 0 {
		return ""
	}

	// Calculate fixed and flexible space
	var fixedSpace int
	var totalGrow float64

	for _, item := range items {
		if item.Basis > 0 {
			fixedSpace += item.Basis
		}
		totalGrow += item.Grow
	}

	availableSpace := width - (gap * (len(items) - 1)) - fixedSpace

	columns := make([]string, len(items))
	for i, item := range items {
		var itemWidth int

		if item.Basis > 0 {
			itemWidth = item.Basis
		} else if totalGrow > 0 {
			itemWidth = int(float64(availableSpace) * (item.Grow / totalGrow))
		} else {
			itemWidth = availableSpace / len(items)
		}

		columns[i] = lipgloss.NewStyle().Width(itemWidth).Render(item.Content)
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, columns...)
}

// CenterContent centers content horizontally and optionally vertically.
func CenterContent(content string, width, height int, vertical bool) string {
	hStyle := lipgloss.NewStyle().Width(width).Align(lipgloss.Center)

	if vertical {
		hStyle = hStyle.Height(height).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center)
	}

	return hStyle.Render(content)
}

// SafeWidth returns a safe width that won't exceed available space.
func SafeWidth(requested, available int) int {
	if requested <= 0 || requested > available {
		return available
	}
	return requested
}

// MinMax constrains a value between min and max.
func MinMax(value, min, max int) int {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// DistributeSpace distributes available space among items.
func DistributeSpace(itemCount, availableSpace, gap int) []int {
	if itemCount <= 0 {
		return nil
	}

	totalGap := gap * (itemCount - 1)
	usableSpace := availableSpace - totalGap

	if usableSpace <= 0 {
		return make([]int, itemCount)
	}

	baseSize := usableSpace / itemCount
	remainder := usableSpace % itemCount

	sizes := make([]int, itemCount)
	for i := 0; i < itemCount; i++ {
		sizes[i] = baseSize
		if i < remainder {
			sizes[i]++
		}
	}

	return sizes
}

// ResponsiveValue returns different values based on breakpoint.
func ResponsiveValue(breakpoint Breakpoint, compact, standard, wide interface{}) interface{} {
	switch breakpoint {
	case BreakpointCompact:
		return compact
	case BreakpointWide:
		return wide
	default:
		return standard
	}
}
