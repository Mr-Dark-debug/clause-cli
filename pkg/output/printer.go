// Package output provides utilities for formatted console output.
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

// Printer provides styled output functions.
type Printer struct {
	theme  *styles.Theme
	writer io.Writer
}

// NewPrinter creates a new printer.
func NewPrinter(theme *styles.Theme, writer io.Writer) *Printer {
	if theme == nil {
		theme = styles.GetTheme()
	}
	if writer == nil {
		writer = os.Stdout
	}
	return &Printer{theme: theme, writer: writer}
}

// DefaultPrinter is the default printer using stdout.
var DefaultPrinter = NewPrinter(nil, nil)

// Print prints a message.
func (p *Printer) Print(args ...interface{}) {
	fmt.Fprint(p.writer, args...)
}

// Println prints a message with a newline.
func (p *Printer) Println(args ...interface{}) {
	fmt.Fprintln(p.writer, args...)
}

// Printf prints a formatted message.
func (p *Printer) Printf(format string, args ...interface{}) {
	fmt.Fprintf(p.writer, format, args...)
}

// PrintStyled prints styled text.
func (p *Printer) PrintStyled(style lipgloss.Style, text string) {
	fmt.Fprint(p.writer, style.Render(text))
}

// PrintStyledln prints styled text with a newline.
func (p *Printer) PrintStyledln(style lipgloss.Style, text string) {
	fmt.Fprintln(p.writer, style.Render(text))
}

// PrintSuccess prints a success message.
func (p *Printer) PrintSuccess(format string, args ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Success))
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintln(p.writer, style.Render("✓ "+msg))
}

// PrintError prints an error message.
func (p *Printer) PrintError(format string, args ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Error))
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintln(p.writer, style.Render("✗ "+msg))
}

// PrintWarning prints a warning message.
func (p *Printer) PrintWarning(format string, args ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Warning))
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintln(p.writer, style.Render("⚠ "+msg))
}

// PrintInfo prints an info message.
func (p *Printer) PrintInfo(format string, args ...interface{}) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Info))
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintln(p.writer, style.Render("ℹ "+msg))
}

// PrintDim prints dimmed text.
func (p *Printer) PrintDim(format string, args ...interface{}) {
	style := lipgloss.NewStyle().Faint(true)
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintln(p.writer, style.Render(msg))
}

// PrintHeader prints a header.
func (p *Printer) PrintHeader(text string) {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(p.theme.Colors.Primary)).
		Padding(1, 0)
	fmt.Fprintln(p.writer, style.Render(text))
}

// PrintSubheader prints a subheader.
func (p *Printer) PrintSubheader(text string) {
	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(p.theme.Colors.Text)).
		PaddingBottom(1)
	fmt.Fprintln(p.writer, style.Render(text))
}

// PrintBullet prints a bullet point.
func (p *Printer) PrintBullet(text string) {
	bulletStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Primary))
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Text))
	fmt.Fprintln(p.writer, bulletStyle.Render("• ")+textStyle.Render(text))
}

// PrintCheckmark prints a checkmark item.
func (p *Printer) PrintCheckmark(text string) {
	checkStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Success))
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Text))
	fmt.Fprintln(p.writer, checkStyle.Render("✓ ")+textStyle.Render(text))
}

// PrintCrossmark prints a crossmark item.
func (p *Printer) PrintCrossmark(text string) {
	crossStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Error))
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Text))
	fmt.Fprintln(p.writer, crossStyle.Render("✗ ")+textStyle.Render(text))
}

// PrintKeyValue prints a key-value pair.
func (p *Printer) PrintKeyValue(key, value string, keyWidth int) {
	keyStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(p.theme.Colors.TextMuted)).
		Width(keyWidth)
	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(p.theme.Colors.Text))
	fmt.Fprintln(p.writer, keyStyle.Render(key)+valueStyle.Render(value))
}

// PrintDivider prints a horizontal divider.
func (p *Printer) PrintDivider(width int) {
	style := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.BorderMuted))
	fmt.Fprintln(p.writer, style.Render(strings.Repeat("─", width)))
}

// Box prints content in a styled box.
func (p *Printer) Box(title, content string, width int) string {
	boxStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color(p.theme.Colors.Border)).
		Padding(0, 1).
		Width(width)

	if title != "" {
		titleStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(p.theme.Colors.Primary))
		return titleStyle.Render(title) + "\n" + boxStyle.Render(content)
	}

	return boxStyle.Render(content)
}

// PrintBox prints content in a box.
func (p *Printer) PrintBox(title, content string, width int) {
	fmt.Fprintln(p.writer, p.Box(title, content, width))
}

// Indent prints indented content.
func (p *Printer) Indent(spaces int, text string) {
	indent := strings.Repeat(" ", spaces)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		fmt.Fprintln(p.writer, indent+line)
	}
}

// PrintIndented prints indented content.
func (p *Printer) PrintIndented(spaces int, format string, args ...interface{}) {
	p.Indent(spaces, fmt.Sprintf(format, args...))
}

// Banner prints a styled banner.
func (p *Printer) Banner(title, subtitle string, width int) {
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(p.theme.Colors.Primary)).
		Align(lipgloss.Center).
		Width(width)

	subtitleStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color(p.theme.Colors.TextMuted)).
		Align(lipgloss.Center).
		Width(width)

	fmt.Fprintln(p.writer)
	fmt.Fprintln(p.writer, titleStyle.Render(title))
	if subtitle != "" {
		fmt.Fprintln(p.writer, subtitleStyle.Render(subtitle))
	}
	fmt.Fprintln(p.writer)
}

// ProgressBar prints a progress bar.
func (p *Printer) ProgressBar(percent float64, width int, label string) {
	filled := int(float64(width) * percent)
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}

	filledStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Primary))
	emptyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.BorderMuted))

	bar := filledStyle.Render(strings.Repeat("█", filled)) +
		emptyStyle.Render(strings.Repeat("░", width-filled))

	if label != "" {
		percentText := fmt.Sprintf(" %3.0f%%", percent*100)
		fmt.Fprintln(p.writer, bar+percentText+" "+label)
	} else {
		percentText := fmt.Sprintf(" %3.0f%%", percent*100)
		fmt.Fprintln(p.writer, bar+percentText)
	}
}

// Spinner prints a spinner animation frame.
func (p *Printer) Spinner(frame, text string) {
	// Clear line and print spinner
	spinnerStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Primary))
	textStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(p.theme.Colors.Text))

	fmt.Fprint(p.writer, "\r")
	fmt.Fprint(p.writer, spinnerStyle.Render(frame))
	fmt.Fprint(p.writer, " ")
	fmt.Fprint(p.writer, textStyle.Render(text))
}

// ClearLine clears the current line.
func (p *Printer) ClearLine() {
	fmt.Fprint(p.writer, "\r"+strings.Repeat(" ", 80)+"\r")
}

// StatusBadge creates a styled status badge.
func (p *Printer) StatusBadge(status string) string {
	var color string
	switch strings.ToLower(status) {
	case "success", "done", "complete", "passed", "ok":
		color = p.theme.Colors.Success
	case "error", "failed", "fail":
		color = p.theme.Colors.Error
	case "warning", "pending", "in_progress":
		color = p.theme.Colors.Warning
	case "info", "running":
		color = p.theme.Colors.Info
	default:
		color = p.theme.Colors.TextMuted
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(p.theme.Colors.TextInverted)).
		Background(lipgloss.Color(color)).
		Padding(0, 1).
		Bold(true).
		Render(strings.ToUpper(status))
}

// Wrap wraps text to the specified width.
func (p *Printer) Wrap(text string, width int) string {
	return styles.TextWrap(text, width)
}

// Truncate truncates text to the specified length.
func (p *Printer) Truncate(text string, maxLen int) string {
	return utils.TruncateText(text, maxLen)
}

// Package-level convenience functions

// Print prints to the default printer.
func Print(args ...interface{}) {
	DefaultPrinter.Print(args...)
}

// Println prints with newline to the default printer.
func Println(args ...interface{}) {
	DefaultPrinter.Println(args...)
}

// Printf prints formatted to the default printer.
func Printf(format string, args ...interface{}) {
	DefaultPrinter.Printf(format, args...)
}

// PrintSuccess prints a success message.
func PrintSuccess(format string, args ...interface{}) {
	DefaultPrinter.PrintSuccess(format, args...)
}

// PrintError prints an error message.
func PrintError(format string, args ...interface{}) {
	DefaultPrinter.PrintError(format, args...)
}

// PrintWarning prints a warning message.
func PrintWarning(format string, args ...interface{}) {
	DefaultPrinter.PrintWarning(format, args...)
}

// PrintInfo prints an info message.
func PrintInfo(format string, args ...interface{}) {
	DefaultPrinter.PrintInfo(format, args...)
}

// PrintHeader prints a header.
func PrintHeader(text string) {
	DefaultPrinter.PrintHeader(text)
}

// PrintBullet prints a bullet point.
func PrintBullet(text string) {
	DefaultPrinter.PrintBullet(text)
}

// PrintCheckmark prints a checkmark item.
func PrintCheckmark(text string) {
	DefaultPrinter.PrintCheckmark(text)
}
