package styles

import (
	"github.com/charmbracelet/lipgloss"
)

// ThemeMode represents the color scheme mode.
type ThemeMode int

const (
	// ModeDark is the default dark theme.
	ModeDark ThemeMode = iota
	// ModeLight is a light theme variant.
	ModeLight
)

// Theme contains all style definitions for the application.
type Theme struct {
	Mode ThemeMode

	// Colors holds all color definitions.
	Colors ColorPalette

	// Typography holds text style definitions.
	Typography TypographyStyles

	// Layout holds layout-related styles.
	Layout LayoutStyles

	// Component holds component-specific styles.
	Component ComponentStyles
}

// ColorPalette contains all color definitions.
type ColorPalette struct {
	// Primary colors
	Primary      string
	PrimaryDim   string
	PrimaryLight string

	// Background colors
	Background      string
	BackgroundAlt   string
	BackgroundCard  string
	BackgroundHover string

	// Text colors
	Text         string
	TextMuted    string
	TextDim      string
	TextInverted string

	// Semantic colors
	Success string
	Warning string
	Error   string
	Info    string

	// Accent colors
	Accent         string
	AccentAlt      string
	AccentTertiary string

	// Border colors
	Border       string
	BorderMuted  string
	BorderAccent string
}

// TypographyStyles contains text-related style definitions.
type TypographyStyles struct {
	// Primary is for primary text.
	Primary lipgloss.Style
	// Title is for main headings.
	Title lipgloss.Style
	// Subtitle is for secondary headings.
	Subtitle lipgloss.Style
	// Header is for section headers.
	Header lipgloss.Style
	// Body is for regular text.
	Body lipgloss.Style
	// Muted is for less important text.
	Muted lipgloss.Style
	// Code is for code snippets.
	Code lipgloss.Style
	// Label is for labels and tags.
	Label lipgloss.Style
	// Error is for error messages.
	Error lipgloss.Style
	// Success is for success messages.
	Success lipgloss.Style
	// Warning is for warning messages.
	Warning lipgloss.Style
	// Info is for info messages.
	Info lipgloss.Style
}

// LayoutStyles contains layout-related style definitions.
type LayoutStyles struct {
	// Container is the main container style.
	Container lipgloss.Style
	// Card is for card-like containers.
	Card lipgloss.Style
	// Section is for section spacing.
	Section lipgloss.Style
	// Spacer is for adding spacing.
	Spacer lipgloss.Style
	// Divider is for horizontal dividers.
	Divider lipgloss.Style
}

// ComponentStyles contains component-specific style definitions.
type ComponentStyles struct {
	// Input is for input fields.
	Input lipgloss.Style
	// InputFocused is for focused input fields.
	InputFocused lipgloss.Style
	// Button is for buttons.
	Button lipgloss.Style
	// ButtonFocused is for focused buttons.
	ButtonFocused lipgloss.Style
	// ButtonSelected is for selected buttons.
	ButtonSelected lipgloss.Style
	// ListItem is for list items.
	ListItem lipgloss.Style
	// ListItemSelected is for selected list items.
	ListItemSelected lipgloss.Style
	// Checkbox is for checkboxes.
	Checkbox lipgloss.Style
	// CheckboxChecked is for checked checkboxes.
	CheckboxChecked lipgloss.Style
	// RadioButton is for radio buttons.
	RadioButton lipgloss.Style
	// RadioButtonSelected is for selected radio buttons.
	RadioButtonSelected lipgloss.Style
	// Progress is for progress bars.
	Progress lipgloss.Style
	// ProgressFilled is for filled progress bar portions.
	ProgressFilled lipgloss.Style
	// Spinner is for spinner indicators.
	Spinner lipgloss.Style
	// Tooltip is for tooltip text.
	Tooltip lipgloss.Style
	// HelpText is for help/keybinding text.
	HelpText lipgloss.Style
	// Header is for screen headers.
	Header lipgloss.Style
	// Footer is for screen footers.
	Footer lipgloss.Style
	// Tab is for tab labels.
	Tab lipgloss.Style
	// TabActive is for active/selected tabs.
	TabActive lipgloss.Style
}

// DefaultTheme is the default dark theme for Clause.
var DefaultTheme = createDarkTheme()

// LightTheme is a light theme variant.
var LightTheme = createLightTheme()

// currentTheme holds the currently active theme.
var currentTheme *Theme

// createDarkTheme creates and returns the default dark theme.
func createDarkTheme() *Theme {
	t := &Theme{
		Mode: ModeDark,
		Colors: ColorPalette{
			Primary:         PrimaryPurple,
			PrimaryDim:      PrimaryPurple, // Using PrimaryPurple for consistency
			PrimaryLight:    PrimaryPurpleLight,
			Background:      BackgroundNavy,
			BackgroundAlt:   BackgroundDarker,
			BackgroundCard:  BackgroundCard,
			BackgroundHover: BackgroundHover,
			Text:            TextPrimary,
			TextMuted:       TextSecondary,
			TextDim:         TextDim,
			TextInverted:    BackgroundNavy,
			Success:         SuccessGreen,
			Warning:         WarningAmber,
			Error:           ErrorRed,
			Info:            InfoBlue,
			Accent:          AccentPurple,
			AccentAlt:       AccentCyan,
			AccentTertiary:  AccentPink,
			Border:          BorderDefault,
			BorderMuted:     BorderMuted,
			BorderAccent:    BorderAccent,
		},
	}

	t.initStyles()
	return t
}

// createLightTheme creates and returns a light theme.
func createLightTheme() *Theme {
	t := &Theme{
		Mode: ModeLight,
		Colors: ColorPalette{
			Primary:         PrimaryPurple,
			PrimaryDim:      PrimaryPurple,
			PrimaryLight:    PrimaryPurpleLight,
			Background:      "#FFFFFF",
			BackgroundAlt:   "#F6F8FA",
			BackgroundCard:  "#F6F8FA",
			BackgroundHover: "#E8ECF0",
			Text:            "#1F2328",
			TextMuted:       "#656D76",
			TextDim:         "#8C959F",
			TextInverted:    "#FFFFFF",
			Success:         "#059669",
			Warning:         "#D97706",
			Error:           "#DC2626",
			Info:            "#2563EB",
			Accent:          "#7C3AED",
			AccentAlt:       "#2563EB",
			AccentTertiary:  "#DB2777",
			Border:          "#D0D7DE",
			BorderMuted:     "#E8ECF0",
			BorderAccent:    PrimaryPurple,
		},
	}

	t.initStyles()
	return t
}

// initStyles initializes all style definitions based on the color palette.
func (t *Theme) initStyles() {
	// Typography styles
	t.Typography = TypographyStyles{
		Primary: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Primary)),

		Title: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.Colors.Primary)).
			Padding(0, 1),

		Subtitle: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.Colors.Text)).
			Padding(0, 1),

		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.Colors.Primary)).
			PaddingBottom(1),

		Body: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Text)),

		Muted: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.TextMuted)),

		Code: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Accent)).
			Background(lipgloss.Color(t.Colors.BackgroundAlt)).
			Padding(0, 1),

		Label: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.Colors.Text)).
			PaddingRight(1),

		Error: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Error)),

		Success: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Success)),

		Warning: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Warning)),

		Info: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Info)),
	}

	// Layout styles
	t.Layout = LayoutStyles{
		Container: lipgloss.NewStyle().
			Padding(1, 2),

		Card: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(t.Colors.Border)).
			Padding(1, 2).
			Background(lipgloss.Color(t.Colors.BackgroundCard)),

		Section: lipgloss.NewStyle().
			Margin(1, 0),

		Spacer: lipgloss.NewStyle().
			Height(1),

		Divider: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.BorderMuted)).
			Padding(0, 1),
	}

	// Component styles
	t.Component = ComponentStyles{
		Input: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(t.Colors.Border)).
			Padding(0, 1).
			Width(40),

		InputFocused: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color(t.Colors.Primary)).
			Padding(0, 1).
			Width(40),

		Button: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(t.Colors.Border)).
			Padding(0, 2).
			Margin(0, 1),

		ButtonFocused: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(t.Colors.Primary)).
			Padding(0, 2).
			Margin(0, 1).
			Foreground(lipgloss.Color(t.Colors.Primary)),

		ButtonSelected: lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color(t.Colors.Primary)).
			Background(lipgloss.Color(t.Colors.Primary)).
			Foreground(lipgloss.Color(t.Colors.TextInverted)).
			Padding(0, 2).
			Margin(0, 1),

		ListItem: lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(lipgloss.Color(t.Colors.Text)),

		ListItemSelected: lipgloss.NewStyle().
			Padding(0, 1).
			Foreground(lipgloss.Color(t.Colors.Primary)).
			Bold(true),

		Checkbox: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.TextMuted)),

		CheckboxChecked: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Success)),

		RadioButton: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.TextMuted)),

		RadioButtonSelected: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Primary)),

		Progress: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.BorderMuted)),

		ProgressFilled: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Primary)),

		Spinner: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Primary)),

		Tooltip: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.TextDim)).
			Italic(true),

		HelpText: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.TextMuted)).
			Padding(1, 0),

		Header: lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color(t.Colors.Primary)).
			Padding(0, 1).
			MarginBottom(1),

		Footer: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.TextMuted)).
			Padding(1, 0).
			MarginTop(1),

		Tab: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.TextMuted)).
			Background(lipgloss.Color(t.Colors.Background)).
			Padding(0, 2).
			MarginRight(1),

		TabActive: lipgloss.NewStyle().
			Foreground(lipgloss.Color(t.Colors.Text)).
			Background(lipgloss.Color(t.Colors.Primary)).
			Bold(true).
			Padding(0, 2).
			MarginRight(1),
	}
}

// GetTheme returns the current theme, initializing it if necessary.
func GetTheme() *Theme {
	if currentTheme == nil {
		currentTheme = DefaultTheme
	}
	return currentTheme
}

// SetTheme sets the current theme.
func SetTheme(theme *Theme) {
	currentTheme = theme
}

// SetThemeMode sets the theme by mode.
func SetThemeMode(mode ThemeMode) {
	switch mode {
	case ModeLight:
		currentTheme = LightTheme
	default:
		currentTheme = DefaultTheme
	}
}

// ApplyTheme returns a copy of the theme with custom colors applied.
func ApplyTheme(base *Theme, overrides ColorPalette) *Theme {
	t := &Theme{
		Mode:   base.Mode,
		Colors: base.Colors,
	}

	// Apply overrides
	if overrides.Primary != "" {
		t.Colors.Primary = overrides.Primary
	}
	if overrides.Background != "" {
		t.Colors.Background = overrides.Background
	}
	if overrides.Text != "" {
		t.Colors.Text = overrides.Text
	}
	if overrides.Success != "" {
		t.Colors.Success = overrides.Success
	}
	if overrides.Warning != "" {
		t.Colors.Warning = overrides.Warning
	}
	if overrides.Error != "" {
		t.Colors.Error = overrides.Error
	}
	if overrides.Info != "" {
		t.Colors.Info = overrides.Info
	}

	t.initStyles()
	return t
}

// StyledText is a helper for creating styled text strings.
type StyledText struct {
	style lipgloss.Style
	text  string
}

// NewStyledText creates a new StyledText with the given text.
func NewStyledText(text string) *StyledText {
	return &StyledText{
		style: lipgloss.NewStyle(),
		text:  text,
	}
}

// WithForeground sets the foreground color.
func (s *StyledText) WithForeground(color string) *StyledText {
	s.style = s.style.Foreground(lipgloss.Color(color))
	return s
}

// WithBackground sets the background color.
func (s *StyledText) WithBackground(color string) *StyledText {
	s.style = s.style.Background(lipgloss.Color(color))
	return s
}

// WithBold sets bold styling.
func (s *StyledText) WithBold(bold bool) *StyledText {
	s.style = s.style.Bold(bold)
	return s
}

// WithItalic sets italic styling.
func (s *StyledText) WithItalic(italic bool) *StyledText {
	s.style = s.style.Italic(italic)
	return s
}

// WithUnderline sets underline styling.
func (s *StyledText) WithUnderline(underline bool) *StyledText {
	s.style = s.style.Underline(underline)
	return s
}

// WithPadding sets padding.
func (s *StyledText) WithPadding(top, right int) *StyledText {
	s.style = s.style.Padding(top, right)
	return s
}

// WithMargin sets margin.
func (s *StyledText) WithMargin(top, right int) *StyledText {
	s.style = s.style.Margin(top, right)
	return s
}

// WithWidth sets width.
func (s *StyledText) WithWidth(width int) *StyledText {
	s.style = s.style.Width(width)
	return s
}

// Render renders the styled text.
func (s *StyledText) Render() string {
	return s.style.Render(s.text)
}

// String implements fmt.Stringer.
func (s *StyledText) String() string {
	return s.Render()
}
