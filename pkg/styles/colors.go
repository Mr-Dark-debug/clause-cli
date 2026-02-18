// Package styles provides styling utilities for the Clause CLI.
// This package defines the color palette, theme system, and text styling
// helpers used throughout the application.
package styles

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

// Primary Brand Colors
const (
	// PrimaryOrange is the main brand color used for CTAs and highlights.
	PrimaryOrange = "#FF6B35"

	// PrimaryOrangeDim is the dimmed version for secondary elements.
	PrimaryOrangeDim = "#CC5529"
)

// Background Colors
const (
	// BackgroundNavy is the main background color for dark mode.
	BackgroundNavy = "#0D1117"

	// BackgroundCard is used for card and component backgrounds.
	BackgroundCard = "#161B22"

	// BackgroundHover is used for hover states.
	BackgroundHover = "#21262D"
)

// Text Colors
const (
	// TextPrimary is the primary text color.
	TextPrimary = "#F0F6FC"

	// TextSecondary is used for secondary or muted text.
	TextSecondary = "#8B949E"

	// TextDim is used for very dim text like hints and placeholders.
	TextDim = "#484F58"
)

// Semantic Colors
const (
	// SuccessGreen is used for success states.
	SuccessGreen = "#3FB950"

	// WarningAmber is used for warnings.
	WarningAmber = "#D29922"

	// ErrorRed is used for errors.
	ErrorRed = "#F85149"

	// InfoBlue is used for informational states.
	InfoBlue = "#58A6FF"
)

// Accent Colors
const (
	// AccentPurple is used for special highlights.
	AccentPurple = "#A371F7"

	// AccentCyan is used for secondary accents.
	AccentCyan = "#39C5CF"

	// AccentPink is used for tertiary accents.
	AccentPink = "#DB61A2"
)

// Border Colors
const (
	// BorderDefault is the default border color.
	BorderDefault = "#30363D"

	// BorderMuted is used for subtle borders.
	BorderMuted = "#21262D"

	// BorderAccent is used for accent borders.
	BorderAccent = "#FF6B35"
)

// Color represents a color that can be used with lipgloss.
type Color lipgloss.Color

// AsLipgloss returns the color as a lipgloss.Color.
func (c Color) AsLipgloss() lipgloss.Color {
	return lipgloss.Color(c)
}

// TerminalColorCapability represents the color support level of a terminal.
type TerminalColorCapability int

const (
	// ColorNone indicates no color support.
	ColorNone TerminalColorCapability = iota
	// Color16 indicates 16-color support.
	Color16
	// Color256 indicates 256-color support.
	Color256
	// ColorTrue indicates 24-bit true color support.
	ColorTrue
)

// colorMap maps hex colors to their 256-color approximations.
// Note: Colors with same hex values are mapped only once.
var colorMap = map[string]string{
	PrimaryOrange:    "208",
	PrimaryOrangeDim: "166",
	BackgroundNavy:   "234",
	BackgroundCard:   "235",
	BackgroundHover:  "236",
	TextPrimary:      "255",
	TextSecondary:    "246",
	TextDim:          "240",
	SuccessGreen:     "71",
	WarningAmber:     "178",
	ErrorRed:         "203",
	InfoBlue:         "39",
	AccentPurple:     "140",
	AccentCyan:       "44",
	AccentPink:       "168",
	BorderDefault:    "238",
	// BorderMuted and BorderAccent use same colors as BackgroundHover and PrimaryOrange
}

// TerminalInfo holds information about terminal capabilities.
type TerminalInfo struct {
	Width       int
	Height      int
	ColorLevel  TerminalColorCapability
	IsTerminal  bool
	SupportsUTF bool
}

// globalTerminalInfo caches terminal information.
var globalTerminalInfo *TerminalInfo

// GetTerminalInfo returns cached terminal information.
func GetTerminalInfo() TerminalInfo {
	if globalTerminalInfo != nil {
		return *globalTerminalInfo
	}

	info := TerminalInfo{
		IsTerminal:  IsTerminal(),
		ColorLevel:  DetectColorCapability(),
		SupportsUTF: true, // Assume UTF-8 support by default
	}

	if width, height, err := GetTerminalSize(); err == nil {
		info.Width = width
		info.Height = height
	}

	globalTerminalInfo = &info
	return info
}

// IsTerminal returns true if stdout is connected to a terminal.
func IsTerminal() bool {
	return term.IsTerminal(int(os.Stdout.Fd()))
}

// GetTerminalSize returns the terminal dimensions.
func GetTerminalSize() (width, height int, err error) {
	return term.GetSize(int(os.Stdout.Fd()))
}

// DetectColorCapability determines the color support level of the terminal.
func DetectColorCapability() TerminalColorCapability {
	// Check COLORTERM environment variable for true color support
	if colorterm := os.Getenv("COLORTERM"); colorterm != "" {
		colorterm = strings.ToLower(colorterm)
		if strings.Contains(colorterm, "truecolor") || strings.Contains(colorterm, "24bit") {
			return ColorTrue
		}
	}

	// Check TERM environment variable
	term := os.Getenv("TERM")
	if term == "" {
		return ColorNone
	}

	term = strings.ToLower(term)

	// Check for 256-color support
	if strings.Contains(term, "256color") || strings.Contains(term, "256") {
		return Color256
	}

	// Check for basic color terminals
	colorTerms := []string{
		"xterm", "screen", "vt100", "color", "ansi", "cygwin",
	}
	for _, ct := range colorTerms {
		if strings.Contains(term, ct) {
			return Color16
		}
	}

	// Check TERM_PROGRAM for common terminals
	termProgram := os.Getenv("TERM_PROGRAM")
	switch strings.ToLower(termProgram) {
	case "iterm.app", "apple_terminal", "warpterminal", "vscode":
		return ColorTrue
	case "hyper":
		return ColorTrue
	}

	return Color16
}

// GetColor returns the appropriate color string based on terminal capabilities.
// For true color terminals, it returns the hex color.
// For 256-color terminals, it returns the closest 256-color code.
// For basic terminals, it returns a basic ANSI color.
func GetColor(hexColor string) string {
	info := GetTerminalInfo()

	switch info.ColorLevel {
	case ColorTrue:
		return hexColor
	case Color256:
		if approx, ok := colorMap[hexColor]; ok {
			return approx
		}
		// Fallback to a reasonable default
		return "255"
	case Color16:
		return get16Color(hexColor)
	default:
		return ""
	}
}

// get16Color maps hex colors to the closest 16-color ANSI code.
func get16Color(hexColor string) string {
	// Map our specific colors to 16-color equivalents
	// Note: PrimaryOrange and BorderAccent have same hex value
	switch hexColor {
	case PrimaryOrange:
		return "3" // Yellow (closest to orange in 16-color)
	case SuccessGreen:
		return "2" // Green
	case WarningAmber:
		return "3" // Yellow
	case ErrorRed:
		return "1" // Red
	case InfoBlue, AccentCyan:
		return "4" // Blue
	case AccentPurple, AccentPink:
		return "5" // Magenta
	case TextPrimary:
		return "15" // White
	case TextSecondary:
		return "7" // Light gray
	case TextDim:
		return "8" // Dark gray
	case BackgroundNavy, BackgroundCard:
		return "0" // Black
	default:
		return "7" // Default to light gray
	}
}

// ColorFunc returns a function that creates a lipgloss style with the given color.
func ColorFunc(hexColor string) func(string) string {
	info := GetTerminalInfo()
	var style lipgloss.Style

	switch info.ColorLevel {
	case ColorTrue:
		style = lipgloss.NewStyle().Foreground(lipgloss.Color(hexColor))
	case Color256:
		if approx, ok := colorMap[hexColor]; ok {
			style = lipgloss.NewStyle().Foreground(lipgloss.Color(approx))
		} else {
			style = lipgloss.NewStyle()
		}
	case Color16:
		// For 16-color, we use ANSI codes directly
		ansiCode := get16Color(hexColor)
		style = lipgloss.NewStyle().Foreground(lipgloss.Color(ansiCode))
	default:
		style = lipgloss.NewStyle()
	}

	return func(s string) string {
		return style.Render(s)
	}
}

// Primary returns a style function for the primary brand color.
func Primary() func(string) string {
	return ColorFunc(PrimaryOrange)
}

// Success returns a style function for success states.
func Success() func(string) string {
	return ColorFunc(SuccessGreen)
}

// Warning returns a style function for warning states.
func Warning() func(string) string {
	return ColorFunc(WarningAmber)
}

// Error returns a style function for error states.
func Error() func(string) string {
	return ColorFunc(ErrorRed)
}

// Info returns a style function for informational states.
func Info() func(string) string {
	return ColorFunc(InfoBlue)
}

// Dim returns a style function for dimmed text.
func Dim() func(string) string {
	return ColorFunc(TextDim)
}

// Secondary returns a style function for secondary text.
func Secondary() func(string) string {
	return ColorFunc(TextSecondary)
}

// StyleWithColor creates a lipgloss style with the given hex color as foreground.
func StyleWithColor(hexColor string) lipgloss.Style {
	info := GetTerminalInfo()

	switch info.ColorLevel {
	case ColorTrue:
		return lipgloss.NewStyle().Foreground(lipgloss.Color(hexColor))
	case Color256:
		if approx, ok := colorMap[hexColor]; ok {
			return lipgloss.NewStyle().Foreground(lipgloss.Color(approx))
		}
		return lipgloss.NewStyle()
	case Color16:
		ansiCode := get16Color(hexColor)
		return lipgloss.NewStyle().Foreground(lipgloss.Color(ansiCode))
	default:
		return lipgloss.NewStyle()
	}
}

// StyleWithBackground creates a lipgloss style with the given hex color as background.
func StyleWithBackground(hexColor string) lipgloss.Style {
	info := GetTerminalInfo()

	switch info.ColorLevel {
	case ColorTrue:
		return lipgloss.NewStyle().Background(lipgloss.Color(hexColor))
	case Color256:
		if approx, ok := colorMap[hexColor]; ok {
			return lipgloss.NewStyle().Background(lipgloss.Color(approx))
		}
		return lipgloss.NewStyle()
	case Color16:
		ansiCode := get16Color(hexColor)
		return lipgloss.NewStyle().Background(lipgloss.Color(ansiCode))
	default:
		return lipgloss.NewStyle()
	}
}

// Gradient creates a gradient effect between two colors.
// Returns a slice of color strings for the gradient steps.
func Gradient(startColor, endColor string, steps int) []string {
	if steps < 2 {
		return []string{startColor}
	}

	// Parse hex colors
	start := parseHexColor(startColor)
	end := parseHexColor(endColor)

	result := make([]string, steps)
	for i := 0; i < steps; i++ {
		t := float64(i) / float64(steps-1)
		r := interpolate(start.r, end.r, t)
		g := interpolate(start.g, end.g, t)
		b := interpolate(start.b, end.b, t)
		result[i] = rgbToHex(r, g, b)
	}

	return result
}

type rgbColor struct {
	r, g, b int
}

func parseHexColor(hex string) rgbColor {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) != 6 {
		return rgbColor{0, 0, 0}
	}

	r := hexToInt(hex[0:2])
	g := hexToInt(hex[2:4])
	b := hexToInt(hex[4:6])

	return rgbColor{r, g, b}
}

func hexToInt(hex string) int {
	var result int
	for _, c := range hex {
		result *= 16
		switch {
		case c >= '0' && c <= '9':
			result += int(c - '0')
		case c >= 'a' && c <= 'f':
			result += int(c-'a') + 10
		case c >= 'A' && c <= 'F':
			result += int(c-'A') + 10
		}
	}
	return result
}

func interpolate(start, end int, t float64) int {
	return start + int(float64(end-start)*t)
}

func rgbToHex(r, g, b int) string {
	return formatHex(r) + formatHex(g) + formatHex(b)
}

func formatHex(n int) string {
	hex := ""
	for n > 0 {
		rem := n % 16
		if rem < 10 {
			hex = string('0'+rem) + hex
		} else {
			hex = string('a'+rem-10) + hex
		}
		n /= 16
	}
	for len(hex) < 2 {
		hex = "0" + hex
	}
	return hex
}
