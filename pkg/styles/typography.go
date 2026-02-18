package styles

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mattn/go-runewidth"
)

// Typography provides text styling and formatting utilities.
type Typography struct {
	theme *Theme
}

// NewTypography creates a new Typography instance with the given theme.
func NewTypography(theme *Theme) *Typography {
	if theme == nil {
		theme = DefaultTheme
	}
	return &Typography{theme: theme}
}

// Title creates a title-styled text.
func (t *Typography) Title(text string) string {
	return t.theme.Typography.Title.Render(text)
}

// Subtitle creates a subtitle-styled text.
func (t *Typography) Subtitle(text string) string {
	return t.theme.Typography.Subtitle.Render(text)
}

// Header creates a header-styled text.
func (t *Typography) Header(text string) string {
	return t.theme.Typography.Header.Render(text)
}

// Body creates body-styled text.
func (t *Typography) Body(text string) string {
	return t.theme.Typography.Body.Render(text)
}

// Muted creates muted/styled text.
func (t *Typography) Muted(text string) string {
	return t.theme.Typography.Muted.Render(text)
}

// Code creates code-styled text.
func (t *Typography) Code(text string) string {
	return t.theme.Typography.Code.Render(text)
}

// Label creates label-styled text.
func (t *Typography) Label(text string) string {
	return t.theme.Typography.Label.Render(text)
}

// Error creates error-styled text.
func (t *Typography) Error(text string) string {
	return t.theme.Typography.Error.Render(text)
}

// Success creates success-styled text.
func (t *Typography) Success(text string) string {
	return t.theme.Typography.Success.Render(text)
}

// Warning creates warning-styled text.
func (t *Typography) Warning(text string) string {
	return t.theme.Typography.Warning.Render(text)
}

// Info creates info-styled text.
func (t *Typography) Info(text string) string {
	return t.theme.Typography.Info.Render(text)
}

// Heading creates a heading with an optional level (1-6).
func (t *Typography) Heading(text string, level int) string {
	switch level {
	case 1:
		return lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.theme.Colors.Primary)).
			Padding(1, 0).
			Render(text)
	case 2:
		return lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.theme.Colors.Text)).
			Padding(0, 1).
			Render(text)
	case 3:
		return lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.theme.Colors.TextMuted)).
			Render(text)
	default:
		return lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.theme.Colors.Text)).
			Bold(true).
			Render(text)
	}
}

// KeyValue formats a key-value pair with styling.
func (t *Typography) KeyValue(key, value string, width int) string {
	keyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.TextMuted)).
		Width(width)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Text))

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		keyStyle.Render(key),
		valueStyle.Render(value),
	)
}

// Bullet creates a bullet point with text.
func (t *Typography) Bullet(text string) string {
	bulletStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Primary)).
		PaddingRight(1)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Text))

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		bulletStyle.Render("•"),
		textStyle.Render(text),
	)
}

// Checkmark creates a checkmark with text.
func (t *Typography) Checkmark(text string) string {
	checkStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Success)).
		PaddingRight(1)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Text))

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		checkStyle.Render("✓"),
		textStyle.Render(text),
	)
}

// Crossmark creates a crossmark with text.
func (t *Typography) Crossmark(text string) string {
	crossStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Error)).
		PaddingRight(1)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Text))

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		crossStyle.Render("✗"),
		textStyle.Render(text),
	)
}

// Arrow creates an arrow indicator with text.
func (t *Typography) Arrow(text string) string {
	arrowStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Primary)).
		PaddingRight(1)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Text))

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		arrowStyle.Render("→"),
		textStyle.Render(text),
	)
}

// NumberedList creates a numbered list item.
func (t *Typography) NumberedList(number int, text string, width int) string {
	numStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Primary)).
		Width(width).
		Align(lipgloss.Right)

	textStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Text)).
		PaddingLeft(1)

	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		numStyle.Render(FormatNumber(number)+"."),
		textStyle.Render(text),
	)
}

// Badge creates a small badge/tag.
func (t *Typography) Badge(text string, color string) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.TextInverted)).
		Background(lipgloss.Color(color)).
		Padding(0, 1).
		Bold(true).
		Render(text)
}

// StatusBadge creates a status badge with semantic coloring.
func (t *Typography) StatusBadge(status string) string {
	var color string
	switch strings.ToLower(status) {
	case "success", "done", "complete", "passed", "ok":
		color = t.theme.Colors.Success
	case "error", "failed", "fail":
		color = t.theme.Colors.Error
	case "warning", "pending", "in_progress":
		color = t.theme.Colors.Warning
	case "info", "running":
		color = t.theme.Colors.Info
	default:
		color = t.theme.Colors.TextMuted
	}

	return t.Badge(strings.ToUpper(status), color)
}

// Link creates a styled link appearance.
func (t *Typography) Link(text, url string) string {
	linkStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(t.theme.Colors.Info)).
		Underline(true)

	return linkStyle.Render(text) + " " + t.Muted("("+url+")")
}

// Quote creates a quoted/blockquote text.
func (t *Typography) Quote(text string) string {
	quoteStyle := lipgloss.NewStyle().
		BorderLeft(true).
		BorderStyle(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color(t.theme.Colors.Primary)).
		PaddingLeft(1).
		Foreground(lipgloss.Color(t.theme.Colors.TextMuted)).
		Italic(true)

	return quoteStyle.Render(text)
}

// TextWrap wraps text to a maximum width, preserving words.
func TextWrap(text string, maxWidth int) string {
	if maxWidth <= 0 {
		return text
	}

	words := strings.Fields(text)
	if len(words) == 0 {
		return text
	}

	var result strings.Builder
	var lineLen int

	for _, word := range words {
		wordLen := runewidth.StringWidth(word)

		if lineLen > 0 && lineLen+1+wordLen > maxWidth {
			result.WriteString("\n")
			lineLen = 0
		}

		if lineLen > 0 {
			result.WriteString(" ")
			lineLen++
		}

		result.WriteString(word)
		lineLen += wordLen
	}

	return result.String()
}

// TruncateText truncates text to maxLen characters with ellipsis.
func TruncateText(text string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}

	textLen := runewidth.StringWidth(text)
	if textLen <= maxLen {
		return text
	}

	// Account for ellipsis
	if maxLen <= 3 {
		return "..."
	}

	truncated := make([]rune, 0, maxLen)
	currentWidth := 0

	for _, r := range text {
		rw := runewidth.RuneWidth(r)
		if currentWidth+rw > maxLen-3 {
			break
		}
		truncated = append(truncated, r)
		currentWidth += rw
	}

	return string(truncated) + "..."
}

// TruncateMiddle truncates text in the middle, preserving start and end.
func TruncateMiddle(text string, maxLen int) string {
	if maxLen <= 0 {
		return ""
	}

	textLen := runewidth.StringWidth(text)
	if textLen <= maxLen {
		return text
	}

	if maxLen <= 3 {
		return "..."
	}

	// Calculate how much to keep from each end
	ellipsis := "..."
	available := maxLen - len(ellipsis)
	startLen := available / 2
	endLen := available - startLen

	// Find the cut points
	runes := []rune(text)
	startEnd := 0
	startWidth := 0
	for i, r := range runes {
		rw := runewidth.RuneWidth(r)
		if startWidth+rw > startLen {
			break
		}
		startEnd = i + 1
		startWidth += rw
	}

	endStart := len(runes)
	endWidth := 0
	for i := len(runes) - 1; i >= 0; i-- {
		rw := runewidth.RuneWidth(runes[i])
		if endWidth+rw > endLen {
			break
		}
		endStart = i
		endWidth += rw
	}

	return string(runes[:startEnd]) + ellipsis + string(runes[endStart:])
}

// PadLeft pads text on the left to reach the specified width.
func PadLeft(text string, width int) string {
	textLen := runewidth.StringWidth(text)
	if textLen >= width {
		return text
	}

	padding := width - textLen
	return strings.Repeat(" ", padding) + text
}

// PadRight pads text on the right to reach the specified width.
func PadRight(text string, width int) string {
	textLen := runewidth.StringWidth(text)
	if textLen >= width {
		return text
	}

	padding := width - textLen
	return text + strings.Repeat(" ", padding)
}

// Center centers text within the specified width.
func Center(text string, width int) string {
	textLen := runewidth.StringWidth(text)
	if textLen >= width {
		return text
	}

	leftPadding := (width - textLen) / 2
	rightPadding := width - textLen - leftPadding

	return strings.Repeat(" ", leftPadding) + text + strings.Repeat(" ", rightPadding)
}

// RepeatChar repeats a character to fill the specified width.
func RepeatChar(char string, width int) string {
	if width <= 0 {
		return ""
	}
	return strings.Repeat(char, width)
}

// HorizontalLine creates a horizontal line with the specified character.
func HorizontalLine(char string, width int) string {
	return RepeatChar(char, width)
}

// FormatNumber formats a number with thousand separators.
func FormatNumber(n int) string {
	if n < 1000 {
		return string(rune('0'+n%10))
	}

	var result []byte
	for n > 0 {
		if len(result) > 0 && len(result)%3 == 0 {
			result = append(result, ',')
		}
		result = append(result, byte('0'+n%10))
		n /= 10
	}

	// Reverse
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return string(result)
}

// Pluralize returns the plural form of a word based on count.
func Pluralize(count int, singular, plural string) string {
	if count == 1 {
		return singular
	}
	return plural
}

// Indent indents each line of text by the specified number of spaces.
func Indent(text string, spaces int) string {
	if spaces <= 0 {
		return text
	}

	indent := strings.Repeat(" ", spaces)
	lines := strings.Split(text, "\n")

	for i, line := range lines {
		if line != "" {
			lines[i] = indent + line
		}
	}

	return strings.Join(lines, "\n")
}

// Bold applies bold styling to text.
func Bold(text string) string {
	return lipgloss.NewStyle().Bold(true).Render(text)
}

// Italic applies italic styling to text.
func Italic(text string) string {
	return lipgloss.NewStyle().Italic(true).Render(text)
}

// Underline applies underline styling to text.
func Underline(text string) string {
	return lipgloss.NewStyle().Underline(true).Render(text)
}

// StrikeThrough applies strikethrough styling to text.
func StrikeThrough(text string) string {
	return lipgloss.NewStyle().Strikethrough(true).Render(text)
}

// DimText applies dim styling to text.
func DimText(text string) string {
	return lipgloss.NewStyle().Faint(true).Render(text)
}

// Blink applies blink styling to text.
func Blink(text string) string {
	return lipgloss.NewStyle().Blink(true).Render(text)
}

// Reverse reverses foreground and background colors.
func Reverse(text string) string {
	return lipgloss.NewStyle().Reverse(true).Render(text)
}
