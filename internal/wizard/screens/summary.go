package screens

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// SummaryScreen shows a summary of the configuration.
type SummaryScreen struct {
	BaseScreen
	confirmed bool
	cursor    int
}

// NewSummaryScreen creates a new summary screen.
func NewSummaryScreen() *SummaryScreen {
	s := &SummaryScreen{
		BaseScreen: *NewBaseScreen("Summary", "summary"),
		confirmed:  false,
		cursor:     0,
	}
	s.complete = true
	return s
}

// Init initializes the screen.
func (s *SummaryScreen) Init() tea.Cmd {
	return nil
}

// Update handles updates.
func (s *SummaryScreen) Update(msg tea.Msg) tea.Cmd {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.String() {
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			if s.cursor < 1 {
				s.cursor++
			}
		case "enter", " ":
			if s.cursor == 0 {
				s.confirmed = true
				// Apply all settings before finishing
				s.applyAllSettings()
			}
			// Navigation is handled by wizard
		}
	}

	return nil
}

// View renders the screen.
func (s *SummaryScreen) View() string {
	var b strings.Builder

	b.WriteString(s.Renderer().Title("Configuration Summary"))
	b.WriteString("\n\n")

	b.WriteString(s.Renderer().Body("Review your project configuration before creating."))
	b.WriteString("\n\n")

	// Project info
	b.WriteString(s.renderSection("Project", s.renderProjectSummary()))
	b.WriteString("\n")

	// Frontend
	if s.Config() != nil && s.Config().Frontend.Enabled {
		b.WriteString(s.renderSection("Frontend", s.renderFrontendSummary()))
		b.WriteString("\n")
	}

	// Backend
	if s.Config() != nil && s.Config().Backend.Enabled {
		b.WriteString(s.renderSection("Backend", s.renderBackendSummary()))
		b.WriteString("\n")
	}

	// Infrastructure
	b.WriteString(s.renderSection("Infrastructure", s.renderInfrastructureSummary()))
	b.WriteString("\n")

	// Governance
	b.WriteString(s.renderSection("AI Governance", s.renderGovernanceSummary()))
	b.WriteString("\n")

	// Confirmation
	b.WriteString(s.Renderer().Divider(s.Width() - 4))
	b.WriteString("\n\n")

	options := []string{
		"Create project with this configuration",
		"Go back to edit",
	}

	for i, opt := range options {
		if i == s.cursor {
			b.WriteString(s.Renderer().ListItem("▸ "+opt, true))
		} else {
			b.WriteString(s.Renderer().ListItem("  "+opt, false))
		}
		b.WriteString("\n")
	}

	b.WriteString("\n")

	kb := tui.NewKeyBindings()
	kb.Add("↑/↓", "Navigate")
	kb.Add("Enter", "Select")
	b.WriteString(s.Renderer().HelpText(kb))

	return b.String()
}

func (s *SummaryScreen) renderSection(title, content string) string {
	return s.Renderer().Box(title, content, s.Width()-4)
}

func (s *SummaryScreen) renderProjectSummary() string {
	if s.Config() == nil {
		return "No configuration"
	}

	var items []string
	cfg := s.Config()

	items = append(items, fmt.Sprintf("Name: %s", cfg.Metadata.Name))
	if cfg.Metadata.Description != "" {
		items = append(items, fmt.Sprintf("Description: %s", cfg.Metadata.Description))
	}
	if cfg.Metadata.Author != "" {
		items = append(items, fmt.Sprintf("Author: %s", cfg.Metadata.Author))
	}
	if cfg.Metadata.License != "" {
		items = append(items, fmt.Sprintf("License: %s", cfg.Metadata.License))
	}

	return strings.Join(items, "\n")
}

func (s *SummaryScreen) renderFrontendSummary() string {
	if s.Config() == nil {
		return "No configuration"
	}

	var items []string
	cfg := s.Config()

	items = append(items, fmt.Sprintf("Framework: %s", cfg.Frontend.Framework))
	items = append(items, fmt.Sprintf("TypeScript: %v", cfg.Frontend.TypeScript))
	items = append(items, fmt.Sprintf("Styling: %s", cfg.Frontend.Styling))

	var features []string
	if cfg.Frontend.Features.SSR {
		features = append(features, "SSR")
	}
	if cfg.Frontend.Features.SSG {
		features = append(features, "SSG")
	}
	if cfg.Frontend.Features.PWA {
		features = append(features, "PWA")
	}
	if cfg.Frontend.Features.I18n {
		features = append(features, "i18n")
	}
	if cfg.Frontend.Features.DarkMode {
		features = append(features, "Dark Mode")
	}
	if cfg.Frontend.Features.Storybook {
		features = append(features, "Storybook")
	}

	if len(features) > 0 {
		items = append(items, fmt.Sprintf("Features: %s", strings.Join(features, ", ")))
	}

	return strings.Join(items, "\n")
}

func (s *SummaryScreen) renderBackendSummary() string {
	if s.Config() == nil {
		return "No configuration"
	}

	var items []string
	cfg := s.Config()

	items = append(items, fmt.Sprintf("Framework: %s", cfg.Backend.Framework))
	items = append(items, fmt.Sprintf("Language: %s", cfg.Backend.Language))
	items = append(items, fmt.Sprintf("Database: %s", cfg.Backend.Database.Primary))
	items = append(items, fmt.Sprintf("API Style: %s", cfg.Backend.API.Style))

	if cfg.Backend.Auth.Provider != "" {
		items = append(items, fmt.Sprintf("Auth: %s", cfg.Backend.Auth.Provider))
	}

	var features []string
	if cfg.Backend.Features.WebSocket {
		features = append(features, "WebSocket")
	}
	if cfg.Backend.Features.BackgroundJobs {
		features = append(features, "Background Jobs")
	}
	if cfg.Backend.Features.FileUpload {
		features = append(features, "File Upload")
	}
	if cfg.Backend.Features.Email {
		features = append(features, "Email")
	}
	if cfg.Backend.Features.RateLimiting {
		features = append(features, "Rate Limiting")
	}

	if len(features) > 0 {
		items = append(items, fmt.Sprintf("Features: %s", strings.Join(features, ", ")))
	}

	return strings.Join(items, "\n")
}

func (s *SummaryScreen) renderInfrastructureSummary() string {
	if s.Config() == nil {
		return "No configuration"
	}

	var items []string
	cfg := s.Config()

	items = append(items, fmt.Sprintf("Hosting: %s", cfg.Infrastructure.Hosting))
	items = append(items, fmt.Sprintf("CI/CD: %s", cfg.Infrastructure.CI))

	if cfg.Infrastructure.Docker {
		items = append(items, "Docker: Enabled")
	}
	if cfg.Infrastructure.DockerCompose {
		items = append(items, "Docker Compose: Enabled")
	}
	if cfg.Infrastructure.Kubernetes {
		items = append(items, "Kubernetes: Enabled")
	}
	if cfg.Infrastructure.CDN {
		items = append(items, "CDN: Enabled")
	}
	if cfg.Infrastructure.Monitoring.Enabled {
		items = append(items, "Monitoring: Enabled")
	}

	return strings.Join(items, "\n")
}

func (s *SummaryScreen) renderGovernanceSummary() string {
	if s.Config() == nil {
		return "No configuration"
	}

	var items []string
	cfg := s.Config()

	items = append(items, fmt.Sprintf("Context Level: %s", cfg.Governance.ContextLevel))

	if cfg.Governance.BrainstormMd {
		items = append(items, "Brainstorm.md: Enabled")
	}
	if cfg.Governance.PromptGuidelines {
		items = append(items, "Prompt Guidelines: Enabled")
	}
	if cfg.Governance.ComponentRegistry {
		items = append(items, "Component Registry: Enabled")
	}

	return strings.Join(items, "\n")
}

// applyAllSettings applies all screen settings to config.
// This is called when the user confirms the configuration.
func (s *SummaryScreen) applyAllSettings() {
	// Settings are applied as users navigate through screens
	// This method can be used for final validation/transformation
}

// IsConfirmed returns true if user confirmed the configuration.
func (s *SummaryScreen) IsConfirmed() bool {
	return s.confirmed
}

// CanGoBack returns true.
func (s *SummaryScreen) CanGoBack() bool {
	return true
}

// SetTheme sets the theme.
func (s *SummaryScreen) SetTheme(theme *styles.Theme) {
	s.BaseScreen.SetTheme(theme)
}

// SetConfig sets the config.
func (s *SummaryScreen) SetConfig(config *config.ProjectConfig) {
	s.BaseScreen.SetConfig(config)
}

// SetSize sets the size.
func (s *SummaryScreen) SetSize(width, height int) {
	s.BaseScreen.SetSize(width, height)
}
