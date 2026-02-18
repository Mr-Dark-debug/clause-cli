// Package wizard provides the interactive project creation wizard.
//
// The wizard guides users through project configuration with a series of
// interactive screens. Each screen collects specific configuration options
// and the final configuration is used to scaffold the project.
//
// Usage:
//
//	w := wizard.New()
//	p := tea.NewProgram(w, tea.WithAltScreen())
//	if _, err := p.Run(); err != nil {
//	    log.Fatal(err)
//	}
//	config := w.Config()
package wizard
