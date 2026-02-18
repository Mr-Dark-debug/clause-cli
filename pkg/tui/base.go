// Package tui provides terminal UI utilities for the Clause CLI.
package tui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/pkg/styles"
)

// Common message types used across the application.

// TickMsg is sent on timer ticks for animations.
type TickMsg struct {
	Time time.Time
	ID   int
}

// Tick creates a command that sends a TickMsg after the duration.
func Tick(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg{Time: t}
	})
}

// TickWithID creates a tick command with an identifier.
func TickWithID(d time.Duration, id int) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return TickMsg{Time: t, ID: id}
	})
}

// FrameMsg is sent for frame-based animations (60fps).
type FrameMsg struct {
	Time time.Time
}

// Frame creates a command for 60fps animations.
func Frame() tea.Cmd {
	return tea.Tick(time.Second/60, func(t time.Time) tea.Msg {
		return FrameMsg{Time: t}
	})
}

// WindowSizeMsg is sent when the terminal is resized.
type WindowSizeMsg struct {
	Width  int
	Height int
}

// KeyMsg wraps a key event.
type KeyMsg struct {
	Key string
}

// FocusMsg indicates focus was gained.
type FocusMsg struct{}

// BlurMsg indicates focus was lost.
type BlurMsg struct{}

// CompleteMsg indicates an operation completed.
type CompleteMsg struct {
	ID    string
	Error error
}

// ErrorMsg contains an error to display.
type ErrorMsg struct {
	Error error
}

// BaseModel provides common functionality for all models.
type BaseModel struct {
	width      int
	height     int
	breakpoint styles.Breakpoint
	focused    bool
	theme      *styles.Theme
	layout     *styles.Layout
}

// NewBaseModel creates a new base model.
func NewBaseModel() BaseModel {
	return BaseModel{
		theme: styles.GetTheme(),
	}
}

// SetSize sets the dimensions and updates the layout.
func (b *BaseModel) SetSize(width, height int) {
	b.width = width
	b.height = height
	b.breakpoint = styles.CalculateBreakpoint(width)
	b.layout = styles.NewLayout(b.theme, width, height)
}

// Width returns the current width.
func (b *BaseModel) Width() int {
	return b.width
}

// Height returns the current height.
func (b *BaseModel) Height() int {
	return b.height
}

// Breakpoint returns the current responsive breakpoint.
func (b *BaseModel) Breakpoint() styles.Breakpoint {
	return b.breakpoint
}

// IsCompact returns true if in compact mode.
func (b *BaseModel) IsCompact() bool {
	return b.breakpoint == styles.BreakpointCompact
}

// IsWide returns true if in wide mode.
func (b *BaseModel) IsWide() bool {
	return b.breakpoint == styles.BreakpointWide
}

// Focused returns whether the model is focused.
func (b *BaseModel) Focused() bool {
	return b.focused
}

// SetFocused sets the focus state.
func (b *BaseModel) SetFocused(focused bool) {
	b.focused = focused
}

// Theme returns the current theme.
func (b *BaseModel) Theme() *styles.Theme {
	return b.theme
}

// SetTheme sets the theme.
func (b *BaseModel) SetTheme(theme *styles.Theme) {
	b.theme = theme
	if b.width > 0 {
		b.layout = styles.NewLayout(theme, b.width, b.height)
	}
}

// Layout returns the current layout.
func (b *BaseModel) Layout() *styles.Layout {
	return b.layout
}

// HandleResize handles window resize messages.
func (b *BaseModel) HandleResize(msg tea.WindowSizeMsg) {
	b.SetSize(msg.Width, msg.Height)
}

// KeyBinding represents a keyboard shortcut.
type KeyBinding struct {
	Key         string
	Description string
	Action      func() tea.Cmd
}

// KeyBindings is a collection of key bindings.
type KeyBindings []KeyBinding

// NewKeyBindings creates a new empty KeyBindings.
func NewKeyBindings() KeyBindings {
	return KeyBindings{}
}

// Add adds a new key binding to the collection.
func (kb *KeyBindings) Add(key, description string) {
	*kb = append(*kb, KeyBinding{Key: key, Description: description})
}

// Help returns formatted help text for the key bindings.
func (kb KeyBindings) Help() string {
	if len(kb) == 0 {
		return ""
	}

	theme := styles.GetTheme()
	keyStyle := theme.Typography.Code
	descStyle := theme.Typography.Muted

	var help string
	for _, binding := range kb {
		if help != "" {
			help += "  "
		}
		help += keyStyle.Render("[" + binding.Key + "]") + " " + descStyle.Render(binding.Description)
	}

	return help
}

// Common key bindings used across the application.
var (
	KeyQuit   = KeyBinding{Key: "q", Description: "Quit"}
	KeyUp     = KeyBinding{Key: "↑/k", Description: "Up"}
	KeyDown   = KeyBinding{Key: "↓/j", Description: "Down"}
	KeyLeft   = KeyBinding{Key: "←/h", Description: "Left"}
	KeyRight  = KeyBinding{Key: "→/l", Description: "Right"}
	KeyEnter  = KeyBinding{Key: "Enter", Description: "Select"}
	KeyBack   = KeyBinding{Key: "Esc", Description: "Back"}
	KeyHelp   = KeyBinding{Key: "?", Description: "Help"}
	KeyTab    = KeyBinding{Key: "Tab", Description: "Next"}
	KeyShiftTab = KeyBinding{Key: "Shift+Tab", Description: "Previous"}
	KeySpace  = KeyBinding{Key: "Space", Description: "Toggle"}
)

// CommonKeyBindings returns the standard navigation bindings.
func CommonKeyBindings() KeyBindings {
	return KeyBindings{
		KeyUp,
		KeyDown,
		KeyEnter,
		KeyBack,
		KeyHelp,
	}
}

// NavigationKeyBindings returns just navigation bindings.
func NavigationKeyBindings() KeyBindings {
	return KeyBindings{
		KeyUp,
		KeyDown,
		KeyEnter,
		KeyBack,
	}
}

// SelectionKeyBindings returns bindings for selection screens.
func SelectionKeyBindings() KeyBindings {
	return KeyBindings{
		KeyUp,
		KeyDown,
		KeyEnter,
		KeyBack,
		KeyHelp,
	}
}

// FocusManager manages focus between multiple focusable elements.
type FocusManager struct {
	elements []string
	current  int
}

// NewFocusManager creates a new focus manager.
func NewFocusManager(elements ...string) *FocusManager {
	return &FocusManager{
		elements: elements,
		current:  0,
	}
}

// Current returns the currently focused element.
func (fm *FocusManager) Current() string {
	if len(fm.elements) == 0 {
		return ""
	}
	return fm.elements[fm.current]
}

// CurrentIndex returns the index of the focused element.
func (fm *FocusManager) CurrentIndex() int {
	return fm.current
}

// Next moves focus to the next element.
func (fm *FocusManager) Next() string {
	if len(fm.elements) == 0 {
		return ""
	}
	fm.current = (fm.current + 1) % len(fm.elements)
	return fm.elements[fm.current]
}

// Prev moves focus to the previous element.
func (fm *FocusManager) Prev() string {
	if len(fm.elements) == 0 {
		return ""
	}
	fm.current = (fm.current - 1 + len(fm.elements)) % len(fm.elements)
	return fm.elements[fm.current]
}

// Set sets focus to a specific element.
func (fm *FocusManager) Set(element string) bool {
	for i, e := range fm.elements {
		if e == element {
			fm.current = i
			return true
		}
	}
	return false
}

// SetIndex sets focus by index.
func (fm *FocusManager) SetIndex(index int) {
	if index >= 0 && index < len(fm.elements) {
		fm.current = index
	}
}

// IsFocused checks if an element is focused.
func (fm *FocusManager) IsFocused(element string) bool {
	return fm.Current() == element
}

// Count returns the number of focusable elements.
func (fm *FocusManager) Count() int {
	return len(fm.elements)
}

// Direction represents a navigation direction.
type Direction int

const (
	DirectionNone Direction = iota
	DirectionUp
	DirectionDown
	DirectionLeft
	DirectionRight
)

// ParseDirection parses a key message to a direction.
func ParseDirection(msg tea.KeyMsg) Direction {
	switch msg.String() {
	case "up", "k":
		return DirectionUp
	case "down", "j":
		return DirectionDown
	case "left", "h":
		return DirectionLeft
	case "right", "l":
		return DirectionRight
	default:
		return DirectionNone
	}
}

// State represents common UI states.
type State int

const (
	StateNormal State = iota
	StateFocused
	StateSelected
	StateDisabled
	StateLoading
	StateError
	StateSuccess
)

// String returns the state name.
func (s State) String() string {
	switch s {
	case StateNormal:
		return "normal"
	case StateFocused:
		return "focused"
	case StateSelected:
		return "selected"
	case StateDisabled:
		return "disabled"
	case StateLoading:
		return "loading"
	case StateError:
		return "error"
	case StateSuccess:
		return "success"
	default:
		return "unknown"
	}
}

// Color returns the color for the state.
func (s State) Color(theme *styles.Theme) string {
	switch s {
	case StateFocused:
		return theme.Colors.Primary
	case StateSelected:
		return theme.Colors.Success
	case StateDisabled:
		return theme.Colors.TextDim
	case StateLoading:
		return theme.Colors.Info
	case StateError:
		return theme.Colors.Error
	case StateSuccess:
		return theme.Colors.Success
	default:
		return theme.Colors.Text
	}
}

// Box draws a box with content.
func Box(title, content string, width int, theme *styles.Theme) string {
	if theme == nil {
		theme = styles.GetTheme()
	}

	// Create title style
	titleStyle := theme.Component.Header

	// Create content style
	contentStyle := theme.Layout.Card.Width(width)

	if title != "" {
		return titleStyle.Render(title) + "\n" + contentStyle.Render(content)
	}

	return contentStyle.Render(content)
}

// Divider draws a horizontal divider.
func Divider(width int, theme *styles.Theme) string {
	if theme == nil {
		theme = styles.GetTheme()
	}
	return styles.NewLayout(theme, width, 1).Divider()
}

// CenteredText centers text within a width.
func CenteredText(text string, width int) string {
	return styles.Center(text, width)
}
