// Package wizard provides the interactive project creation wizard.
package wizard

import (
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/internal/wizard/screens"
)

// Screen is an alias to screens.Screen for convenience
type Screen = screens.Screen

// WizardState represents the current state of the wizard.
type WizardState struct {
	CurrentScreen int
	TotalScreens  int
	Config        *config.ProjectConfig
	Preset        string
}

// Messages for wizard events
type (
	// NextScreenMsg signals to move to the next screen
	NextScreenMsg struct{}

	// PrevScreenMsg signals to move to the previous screen
	PrevScreenMsg struct{}

	// FinishMsg signals that the wizard is complete
	FinishMsg struct{}

	// QuitMsg signals to quit the wizard
	QuitMsg struct{}

	// ErrorMsg contains an error
	ErrorMsg struct {
		Error error
	}

	// FadeInMsg signals a fade-in animation tick
	FadeInMsg struct{}
)

// Animation constants
const (
	FadeInterval = 16 // milliseconds (60fps)
	FadeStep     = 0.1
)
