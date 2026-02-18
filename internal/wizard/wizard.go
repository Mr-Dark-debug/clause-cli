package wizard

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/internal/wizard/screens"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// Wizard is the main project creation wizard.
type Wizard struct {
	// Configuration
	config   *config.ProjectConfig
	theme    *styles.Theme
	renderer *tui.Renderer
	preset   string

	// State
	screenInstances []Screen
	current         int
	width           int
	height          int
	quitting        bool
	finished        bool
	err             error

	// Animation
	fadeIn    bool
	fadeAlpha float64
}

// New creates a new wizard.
func New() *Wizard {
	theme := styles.GetTheme()

	w := &Wizard{
		config:   config.NewProjectConfig(),
		theme:    theme,
		renderer: tui.NewRenderer(theme, 80, 24),
		current:  0,
		fadeIn:   true,
	}

	// Add screens in order
	w.addScreens()

	return w
}

// NewWithPreset creates a new wizard with a preset configuration.
func NewWithPreset(preset string) (*Wizard, error) {
	presetConfig, err := config.LoadPreset(preset)
	if err != nil {
		return nil, fmt.Errorf("failed to load preset: %w", err)
	}

	w := New()
	w.config = presetConfig
	w.preset = preset

	return w, nil
}

// addScreens adds all wizard screens in order.
func (w *Wizard) addScreens() {
	// Order: Welcome -> Project Info -> Frontend -> Backend -> Infrastructure -> Governance -> Summary
	w.screenInstances = []Screen{
		screens.NewWelcomeScreen(),
		screens.NewProjectScreen(),
		screens.NewFrontendScreen(),
		screens.NewBackendScreen(),
		screens.NewInfrastructureScreen(),
		screens.NewGovernanceScreen(),
		screens.NewSummaryScreen(),
	}

	// Initialize all screens with theme and config
	for _, screen := range w.screenInstances {
		screen.SetTheme(w.theme)
		screen.SetConfig(w.config)
	}
}

// Init implements tea.Model.
func (w *Wizard) Init() tea.Cmd {
	var cmds []tea.Cmd

	// Start fade-in animation
	cmds = append(cmds, w.fadeInCmd())

	// Initialize current screen
	if len(w.screenInstances) > 0 {
		cmds = append(cmds, w.screenInstances[w.current].Init())
	}

	return tea.Batch(cmds...)
}

// Update implements tea.Model.
func (w *Wizard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	// Handle fade-in animation
	if w.fadeIn {
		if cmd, done := w.updateFadeIn(msg); done {
			w.fadeIn = false
		} else {
			cmds = append(cmds, cmd)
		}
	}

	switch m := msg.(type) {
	case tea.WindowSizeMsg:
		w.width = m.Width
		w.height = m.Height
		w.renderer.SetSize(m.Width, m.Height)
		for _, screen := range w.screenInstances {
			screen.SetSize(m.Width, m.Height)
		}

	case tea.KeyMsg:
		// Handle global keys
		switch m.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			if w.current == 0 {
				w.quitting = true
				return w, tea.Quit
			}
		}

		// Handle navigation
		if cmd := w.handleNavigation(m); cmd != nil {
			cmds = append(cmds, cmd)
		}

	case NextScreenMsg:
		return w, w.nextScreen()

	case PrevScreenMsg:
		return w, w.prevScreen()

	case FinishMsg:
		w.finished = true
		return w, tea.Quit

	case QuitMsg:
		w.quitting = true
		return w, tea.Quit

	case ErrorMsg:
		w.err = m.Error
	}

	// Update current screen
	if len(w.screenInstances) > 0 {
		cmd := w.screenInstances[w.current].Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return w, tea.Batch(cmds...)
}

// View implements tea.Model.
func (w *Wizard) View() string {
	if w.quitting {
		return w.viewQuit()
	}

	if w.finished {
		return w.viewFinished()
	}

	if w.err != nil {
		return w.viewError()
	}

	// Render current screen
	if len(w.screenInstances) == 0 {
		return "No screens configured"
	}

	content := w.screenInstances[w.current].View()

	// Add progress indicator
	content = w.addProgressIndicator(content)

	// Apply fade effect
	if w.fadeIn && w.fadeAlpha < 1.0 {
		content = tui.ApplyFade(content, w.fadeAlpha)
	}

	return content
}

// viewQuit renders the quit message.
func (w *Wizard) viewQuit() string {
	return w.renderer.Info("Wizard cancelled. Run 'clause init' to start again.")
}

// viewFinished renders the completion message.
func (w *Wizard) viewFinished() string {
	var parts []string

	parts = append(parts, "")
	parts = append(parts, w.renderer.Success("Project configured successfully!"))
	parts = append(parts, "")
	parts = append(parts, w.renderer.Body("Configuration saved to .clause/config.yaml"))
	parts = append(parts, "")
	parts = append(parts, w.renderer.Muted("Next steps:"))
	parts = append(parts, w.renderer.BulletList([]string{
		"Review your configuration",
		"Run 'clause generate' to scaffold your project",
		"Start building with AI assistance",
	}))
	parts = append(parts, "")

	return tui.JoinVertical(parts...)
}

// viewError renders an error message.
func (w *Wizard) viewError() string {
	var parts []string

	parts = append(parts, "")
	parts = append(parts, w.renderer.Error("An error occurred:"))
	parts = append(parts, "")
	parts = append(parts, w.renderer.Body(w.err.Error()))
	parts = append(parts, "")

	return tui.JoinVertical(parts...)
}

// addProgressIndicator adds a progress bar to the content.
func (w *Wizard) addProgressIndicator(content string) string {
	if len(w.screenInstances) == 0 {
		return content
	}

	progress := float64(w.current+1) / float64(len(w.screenInstances))
	bar := w.renderer.ProgressBar(progress, w.width-10)
	percent := w.renderer.PercentText(progress)

	progressLine := tui.JoinHorizontal(bar, " ", percent)

	// Add screen name
	screenName := ""
	if w.current < len(w.screenInstances) {
		screenName = w.renderer.Muted(fmt.Sprintf("Step %d of %d: %s",
			w.current+1, len(w.screenInstances), w.screenInstances[w.current].Name()))
	}

	return tui.JoinVertical(
		screenName,
		"",
		content,
		"",
		progressLine,
		w.renderer.HelpText(w.getKeyBindings()),
	)
}

// handleNavigation handles navigation key presses.
func (w *Wizard) handleNavigation(msg tea.KeyMsg) tea.Cmd {
	if len(w.screenInstances) == 0 {
		return nil
	}

	currentScreen := w.screenInstances[w.current]

	switch msg.String() {
	case "enter", "right", "l":
		if currentScreen.CanGoNext() && currentScreen.IsComplete() {
			return w.nextScreen
		}

	case "backspace", "left", "h":
		if currentScreen.CanGoBack() {
			return w.prevScreen
		}

	case "ctrl+n":
		if currentScreen.CanGoNext() {
			return w.nextScreen
		}

	case "ctrl+p":
		if currentScreen.CanGoBack() {
			return w.prevScreen
		}
	}

	return nil
}

// nextScreen moves to the next screen.
func (w *Wizard) nextScreen() tea.Cmd {
	if w.current >= len(w.screenInstances)-1 {
		// Last screen, finish
		return tea.Quit
	}

	w.current++
	w.fadeIn = true
	w.fadeAlpha = 0.0

	return tea.Batch(
		w.fadeInCmd(),
		w.screenInstances[w.current].Init(),
	)
}

// prevScreen moves to the previous screen.
func (w *Wizard) prevScreen() tea.Cmd {
	if w.current <= 0 {
		return nil
	}

	w.current--
	w.fadeIn = true
	w.fadeAlpha = 0.0

	return tea.Batch(
		w.fadeInCmd(),
		w.screenInstances[w.current].Init(),
	)
}

// fadeInCmd returns a command for the fade-in animation.
func (w *Wizard) fadeInCmd() tea.Cmd {
	return tea.Tick(time.Duration(FadeInterval)*time.Millisecond, func(t time.Time) tea.Msg {
		return FadeInMsg{}
	})
}

// updateFadeIn updates the fade-in animation.
func (w *Wizard) updateFadeIn(msg tea.Msg) (tea.Cmd, bool) {
	if _, ok := msg.(FadeInMsg); !ok {
		return w.fadeInCmd(), false
	}

	w.fadeAlpha += FadeStep
	if w.fadeAlpha >= 1.0 {
		w.fadeAlpha = 1.0
		return nil, true
	}

	return w.fadeInCmd(), false
}

// getKeyBindings returns the key bindings for the wizard.
func (w *Wizard) getKeyBindings() tui.KeyBindings {
	kb := tui.NewKeyBindings()

	if len(w.screenInstances) > 0 {
		current := w.screenInstances[w.current]

		if current.CanGoBack() {
			kb.Add("←/h", "Back")
		}

		if current.CanGoNext() && current.IsComplete() {
			if w.current == len(w.screenInstances)-1 {
				kb.Add("Enter", "Finish")
			} else {
				kb.Add("Enter/→", "Next")
			}
		}
	}

	kb.Add("Esc/Ctrl+C", "Cancel")

	return kb
}

// Config returns the current configuration.
func (w *Wizard) Config() *config.ProjectConfig {
	return w.config
}

// CurrentScreen returns the current screen.
func (w *Wizard) CurrentScreen() Screen {
	if w.current < len(w.screenInstances) {
		return w.screenInstances[w.current]
	}
	return nil
}

// SetPreset sets the preset configuration.
func (w *Wizard) SetPreset(preset string) error {
	presetConfig, err := config.LoadPreset(preset)
	if err != nil {
		return err
	}

	w.config = presetConfig
	w.preset = preset

	// Update all screens with new config
	for _, screen := range w.screenInstances {
		screen.SetConfig(w.config)
	}

	return nil
}

// IsFinished returns true if the wizard completed successfully.
func (w *Wizard) IsFinished() bool {
	return w.finished
}

// IsQuitting returns true if the user quit the wizard.
func (w *Wizard) IsQuitting() bool {
	return w.quitting
}
