package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/utils"
)

// Renderer provides common rendering utilities.
type Renderer struct {
	theme  *styles.Theme
	width  int
	height int
}

// NewRenderer creates a new renderer.
func NewRenderer(theme *styles.Theme, width, height int) *Renderer {
	if theme == nil {
		theme = styles.GetTheme()
	}
	return &Renderer{
		theme:  theme,
		width:  width,
		height: height,
	}
}

// Theme returns the current theme.
func (r *Renderer) Theme() *styles.Theme {
	return r.theme
}

// SetSize sets the dimensions.
func (r *Renderer) SetSize(width, height int) {
	r.width = width
	r.height = height
}

// Width returns the current width.
func (r *Renderer) Width() int {
	return r.width
}

// Height returns the current height.
func (r *Renderer) Height() int {
	return r.height
}

// Title renders a title.
func (r *Renderer) Title(text string) string {
	return r.theme.Typography.Title.Render(text)
}

// Subtitle renders a subtitle.
func (r *Renderer) Subtitle(text string) string {
	return r.theme.Typography.Subtitle.Render(text)
}

// Header renders a header.
func (r *Renderer) Header(text string) string {
	return r.theme.Typography.Header.Render(text)
}

// Body renders body text.
func (r *Renderer) Body(text string) string {
	return r.theme.Typography.Body.Render(text)
}

// Muted renders muted text.
func (r *Renderer) Muted(text string) string {
	return r.theme.Typography.Muted.Render(text)
}

// Error renders an error message.
func (r *Renderer) Error(text string) string {
	return r.theme.Typography.Error.Render(text)
}

// Success renders a success message.
func (r *Renderer) Success(text string) string {
	return r.theme.Typography.Success.Render(text)
}

// Warning renders a warning message.
func (r *Renderer) Warning(text string) string {
	return r.theme.Typography.Warning.Render(text)
}

// Info renders an info message.
func (r *Renderer) Info(text string) string {
	return r.theme.Typography.Info.Render(text)
}

// Card renders a card with content.
func (r *Renderer) Card(content string, width int) string {
	style := r.theme.Layout.Card
	if width > 0 {
		style = style.Width(width)
	}
	return style.Render(content)
}

// Box renders a bordered box.
func (r *Renderer) Box(title, content string, width int) string {
	var parts []string

	if title != "" {
		titleStyle := r.theme.Component.Header.Width(width)
		parts = append(parts, titleStyle.Render(title))
	}

	contentStyle := r.theme.Layout.Card.Width(width)
	parts = append(parts, contentStyle.Render(content))

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

// ListItem renders a list item.
func (r *Renderer) ListItem(text string, selected bool) string {
	if selected {
		return r.theme.Component.ListItemSelected.Render("▸ " + text)
	}
	return r.theme.Component.ListItem.Render("  " + text)
}

// Checkbox renders a checkbox.
func (r *Renderer) Checkbox(text string, checked bool) string {
	var box string
	if checked {
		box = r.theme.Component.CheckboxChecked.Render("✓")
	} else {
		box = r.theme.Component.Checkbox.Render("○")
	}
	return box + " " + text
}

// RadioButton renders a radio button.
func (r *Renderer) RadioButton(text string, selected bool) string {
	var btn string
	if selected {
		btn = r.theme.Component.RadioButtonSelected.Render("●")
	} else {
		btn = r.theme.Component.RadioButton.Render("○")
	}
	return btn + " " + text
}

// Button renders a button.
func (r *Renderer) Button(text string, focused, selected bool) string {
	var style lipgloss.Style
	switch {
	case selected:
		style = r.theme.Component.ButtonSelected
	case focused:
		style = r.theme.Component.ButtonFocused
	default:
		style = r.theme.Component.Button
	}
	return style.Render(text)
}

// Divider renders a horizontal divider.
func (r *Renderer) Divider(width int) string {
	if width <= 0 {
		width = r.width
	}
	return styles.HorizontalLine("─", width)
}

// KeyValue renders a key-value pair.
func (r *Renderer) KeyValue(key, value string, keyWidth int) string {
	keyStyle := r.theme.Typography.Muted.Width(keyWidth)
	valueStyle := r.theme.Typography.Body

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		keyStyle.Render(key),
		valueStyle.Render(value),
	)
}

// StatusBadge renders a status badge.
func (r *Renderer) StatusBadge(status string) string {
	typo := styles.NewTypography(r.theme)
	return typo.StatusBadge(status)
}

// BulletList renders a bullet list.
func (r *Renderer) BulletList(items []string) string {
	typo := styles.NewTypography(r.theme)
	var lines []string
	for _, item := range items {
		lines = append(lines, typo.Bullet(item))
	}
	return strings.Join(lines, "\n")
}

// CheckmarkList renders a list with checkmarks.
func (r *Renderer) CheckmarkList(items []string, checked []bool) string {
	typo := styles.NewTypography(r.theme)
	var lines []string
	for i, item := range items {
		if i < len(checked) && checked[i] {
			lines = append(lines, typo.Checkmark(item))
		} else {
			lines = append(lines, typo.Crossmark(item))
		}
	}
	return strings.Join(lines, "\n")
}

// NumberedList renders a numbered list.
func (r *Renderer) NumberedList(items []string, startWidth int) string {
	typo := styles.NewTypography(r.theme)
	var lines []string
	for i, item := range items {
		lines = append(lines, typo.NumberedList(i+1, item, startWidth))
	}
	return strings.Join(lines, "\n")
}

// ProgressBar renders a progress bar.
func (r *Renderer) ProgressBar(percent float64, width int) string {
	if width <= 0 {
		width = r.width - 4
	}

	filled := int(float64(width) * percent)
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}

	filledStyle := r.theme.Component.ProgressFilled
	emptyStyle := r.theme.Component.Progress

	bar := filledStyle.Render(strings.Repeat("█", filled)) +
		emptyStyle.Render(strings.Repeat("░", width-filled))

	return bar
}

// PercentText renders percentage text.
func (r *Renderer) PercentText(percent float64) string {
	return r.theme.Typography.Body.Render(
		utils.PadLeft(fmt.Sprintf("%.0f%%", percent*100), " ", 5),
	)
}

// InputField renders an input field.
func (r *Renderer) InputField(value, placeholder string, focused bool, width int) string {
	var style lipgloss.Style
	if focused {
		style = r.theme.Component.InputFocused
	} else {
		style = r.theme.Component.Input
	}

	if width > 0 {
		style = style.Width(width)
	}

	display := value
	if display == "" && !focused {
		display = r.theme.Typography.Muted.Render(placeholder)
	}

	return style.Render(display)
}

// HelpText renders help text.
func (r *Renderer) HelpText(bindings KeyBindings) string {
	return r.theme.Component.HelpText.Render(bindings.Help())
}

// HeaderBar renders a header bar with title.
func (r *Renderer) HeaderBar(title string, width int) string {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(r.theme.Colors.Primary)).
		Padding(0, 1)

	return titleStyle.Width(width).Render(title)
}

// FooterBar renders a footer bar with content.
func (r *Renderer) FooterBar(content string, width int) string {
	return r.theme.Component.Footer.Width(width).Render(content)
}

// Screen renders a complete screen with header, content, and footer.
func (r *Renderer) Screen(header, content, footer string) string {
	var parts []string

	if header != "" {
		parts = append(parts, header)
	}

	if content != "" {
		parts = append(parts, content)
	}

	if footer != "" {
		parts = append(parts, footer)
	}

	return lipgloss.JoinVertical(lipgloss.Left, parts...)
}

// Center centers content in available space.
func (r *Renderer) Center(content string, width, height int) string {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(content)
}

// Padding adds padding around content.
func (r *Renderer) Padding(content string, top, right, bottom, left int) string {
	return lipgloss.NewStyle().
		Padding(top, right, bottom, left).
		Render(content)
}

// Margin adds margin around content.
func (r *Renderer) Margin(content string, top, right, bottom, left int) string {
	return lipgloss.NewStyle().
		Margin(top, right, bottom, left).
		Render(content)
}

// JoinVertical joins strings vertically.
func JoinVertical(items ...string) string {
	return lipgloss.JoinVertical(lipgloss.Left, items...)
}

// JoinHorizontal joins strings horizontally.
func JoinHorizontal(items ...string) string {
	return lipgloss.JoinHorizontal(lipgloss.Top, items...)
}

// LimitLines limits content to a maximum number of lines.
func LimitLines(content string, maxLines int) string {
	lines := strings.Split(content, "\n")
	if len(lines) <= maxLines {
		return content
	}
	return strings.Join(lines[:maxLines], "\n")
}

// TruncateLines truncates content to fit within a height.
func TruncateLines(content string, maxLines int, indicator string) string {
	lines := strings.Split(content, "\n")
	if len(lines) <= maxLines {
		return content
	}

	if maxLines > 1 {
		return strings.Join(lines[:maxLines-1], "\n") + "\n" + indicator
	}
	return indicator
}

// EnsureHeight pads or truncates content to exact height.
func EnsureHeight(content string, height int, truncate bool) string {
	lines := strings.Split(content, "\n")

	if len(lines) > height {
		if truncate {
			return strings.Join(lines[:height], "\n")
		}
		return content
	}

	// Pad with empty lines
	for len(lines) < height {
		lines = append(lines, "")
	}

	return strings.Join(lines, "\n")
}

// RenderList renders a scrollable list with indicator.
func RenderList(items []string, selected, offset, visible int, width int) string {
	if len(items) == 0 {
		return ""
	}

	// Ensure bounds
	if selected < 0 {
		selected = 0
	}
	if selected >= len(items) {
		selected = len(items) - 1
	}

	start, end := CalculateVisibleRange(len(items), selected, visible, offset)

	var lines []string
	for i := start; i < end; i++ {
		item := items[i]
		if len(item) > width-2 {
			item = utils.TruncateText(item, width-2)
		}

		if i == selected {
			lines = append(lines, "▸ "+item)
		} else {
			lines = append(lines, "  "+item)
		}
	}

	return strings.Join(lines, "\n")
}

// ScrollIndicator renders a scroll position indicator.
func ScrollIndicator(total, visible, offset int) string {
	if total <= visible {
		return ""
	}

	// Calculate position
	position := float64(offset) / float64(total-visible)
	if position < 0 {
		position = 0
	}
	if position > 1 {
		position = 1
	}

	// Create indicator
	indicatorHeight := 5
	pos := int(position * float64(indicatorHeight-1))

	var lines []string
	for i := 0; i < indicatorHeight; i++ {
		if i == pos {
			lines = append(lines, "█")
		} else {
			lines = append(lines, "│")
		}
	}

	return strings.Join(lines, "\n")
}
