package screens

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// Screen represents a wizard screen interface.
// Screens implement this interface to be used in the wizard.
type Screen interface {
	// Name returns the screen name for display
	Name() string

	// ID returns a unique identifier for the screen
	ID() string

	// Init initializes the screen
	Init() tea.Cmd

	// Update handles updates for the screen and returns a tea.Cmd
	Update(tea.Msg) tea.Cmd

	// View renders the screen
	View() string

	// SetSize sets the screen dimensions
	SetSize(width, height int)

	// SetTheme sets the screen theme
	SetTheme(theme *styles.Theme)

	// SetConfig sets the project configuration
	SetConfig(config *config.ProjectConfig)

	// CanGoBack returns true if user can go back from this screen
	CanGoBack() bool

	// CanGoNext returns true if user can proceed to next screen
	CanGoNext() bool

	// IsComplete returns true if the screen's data entry is complete
	IsComplete() bool
}

// BaseScreen provides common functionality for all screens.
// Embed this struct in your screen implementations to get basic
// screen functionality for free.
type BaseScreen struct {
	name     string
	id       string
	theme    *styles.Theme
	renderer *tui.Renderer
	config   *config.ProjectConfig
	width    int
	height   int
	complete bool
	focused  bool
}

// NewBaseScreen creates a new base screen with the given name and ID.
func NewBaseScreen(name, id string) *BaseScreen {
	return &BaseScreen{
		name:     name,
		id:       id,
		complete: false,
		focused:  true,
	}
}

// Name returns the screen name.
func (s *BaseScreen) Name() string {
	return s.name
}

// ID returns the screen ID.
func (s *BaseScreen) ID() string {
	return s.id
}

// SetSize sets the screen dimensions.
func (s *BaseScreen) SetSize(width, height int) {
	s.width = width
	s.height = height
	if s.renderer != nil {
		s.renderer.SetSize(width, height)
	}
}

// SetTheme sets the screen theme.
func (s *BaseScreen) SetTheme(theme *styles.Theme) {
	s.theme = theme
	s.renderer = tui.NewRenderer(theme, s.width, s.height)
}

// SetConfig sets the project configuration.
func (s *BaseScreen) SetConfig(config *config.ProjectConfig) {
	s.config = config
}

// CanGoBack returns true by default.
func (s *BaseScreen) CanGoBack() bool {
	return true
}

// CanGoNext returns true by default.
func (s *BaseScreen) CanGoNext() bool {
	return true
}

// IsComplete returns the completion state.
func (s *BaseScreen) IsComplete() bool {
	return s.complete
}

// SetComplete sets the completion state.
func (s *BaseScreen) SetComplete(complete bool) {
	s.complete = complete
}

// Focused returns whether the screen is focused.
func (s *BaseScreen) Focused() bool {
	return s.focused
}

// SetFocused sets the focus state.
func (s *BaseScreen) SetFocused(focused bool) {
	s.focused = focused
}

// Theme returns the current theme.
func (s *BaseScreen) Theme() *styles.Theme {
	return s.theme
}

// Renderer returns the current renderer.
func (s *BaseScreen) Renderer() *tui.Renderer {
	return s.renderer
}

// Config returns the current configuration.
func (s *BaseScreen) Config() *config.ProjectConfig {
	return s.config
}

// Width returns the screen width.
func (s *BaseScreen) Width() int {
	return s.width
}

// Height returns the screen height.
func (s *BaseScreen) Height() int {
	return s.height
}

// BaseInit provides a default no-op init.
func (s *BaseScreen) BaseInit() tea.Cmd {
	return nil
}

// BaseUpdate provides a default no-op update.
func (s *BaseScreen) BaseUpdate(msg tea.Msg) tea.Cmd {
	return nil
}

// BaseView provides a default no-op view.
func (s *BaseScreen) BaseView() string {
	return ""
}

// ScreenHeader renders a common header for screens.
func (s *BaseScreen) ScreenHeader() string {
	if s.renderer == nil {
		return s.name
	}
	return s.renderer.Title(s.name)
}

// ScreenHelp renders help text for the screen.
func (s *BaseScreen) ScreenHelp(bindings tui.KeyBindings) string {
	if s.renderer == nil {
		return bindings.Help()
	}
	return s.renderer.HelpText(bindings)
}

// Validate validates screen input - override in specific screens.
func (s *BaseScreen) Validate() error {
	return nil
}
