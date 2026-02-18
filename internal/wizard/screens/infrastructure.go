package screens

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// InfrastructureScreen configures infrastructure options.
type InfrastructureScreen struct {
	BaseScreen
	section    int
	cursor     int
	features   map[string]bool
	hostingIdx int
	ciIdx      int
}

// Hosting options
var hostingOptions = []struct {
	name        string
	description string
}{
	{"Vercel", "Optimized for frontend & serverless"},
	{"Netlify", "Static sites & functions"},
	{"AWS", "Amazon Web Services"},
	{"Google Cloud", "Google Cloud Platform"},
	{"Azure", "Microsoft Azure"},
	{"Railway", "Simple infrastructure platform"},
	{"Render", "Cloud platform for apps"},
	{"Fly.io", "Global app deployment"},
	{"DigitalOcean", "Cloud infrastructure"},
	{"Self-hosted", "On your own servers"},
}

// CI/CD options
var ciOptions = []struct {
	name        string
	description string
}{
	{"GitHub Actions", "Native GitHub CI/CD"},
	{"GitLab CI", "GitLab continuous integration"},
	{"CircleCI", "Fast CI/CD platform"},
	{"Jenkins", "Self-hosted automation"},
	{"None", "Skip CI configuration"},
}

// Infrastructure feature options
var infraFeatureOptions = []struct {
	key         string
	name        string
	description string
}{
	{"docker", "Docker", "Container support"},
	{"docker_compose", "Docker Compose", "Multi-container apps"},
	{"kubernetes", "Kubernetes", "Container orchestration"},
	{"terraform", "Terraform", "Infrastructure as code"},
	{"monitoring", "Monitoring", "App monitoring & alerting"},
	{"cdn", "CDN", "Content delivery network"},
}

// NewInfrastructureScreen creates a new infrastructure screen.
func NewInfrastructureScreen() *InfrastructureScreen {
	return &InfrastructureScreen{
		BaseScreen: *NewBaseScreen("Infrastructure", "infrastructure"),
		features: map[string]bool{
			"docker":         true,
			"docker_compose": true,
			"kubernetes":     false,
			"terraform":      false,
			"monitoring":     true,
			"cdn":            true,
		},
		hostingIdx: 0,
		ciIdx:      0,
		section:    0,
		cursor:     0,
	}
}

// Init initializes the screen.
func (s *InfrastructureScreen) Init() tea.Cmd {
	return nil
}

// Update handles updates.
func (s *InfrastructureScreen) Update(msg tea.Msg) tea.Cmd {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.String() {
		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			max := s.getMaxItems()
			if s.cursor < max-1 {
				s.cursor++
			}
		case "left", "h":
			if s.section > 0 {
				s.section--
				s.cursor = 0
			}
		case "right", "l":
			if s.section < 2 {
				s.section++
				s.cursor = 0
			}
		case "enter", " ":
			s.toggle()
		case "tab":
			if s.section < 2 {
				s.section++
				s.cursor = 0
			}
		}
	}

	s.complete = true
	return nil
}

func (s *InfrastructureScreen) getMaxItems() int {
	switch s.section {
	case 0:
		return len(hostingOptions)
	case 1:
		return len(ciOptions)
	case 2:
		return len(infraFeatureOptions)
	}
	return 0
}

func (s *InfrastructureScreen) toggle() {
	switch s.section {
	case 0:
		s.hostingIdx = s.cursor
	case 1:
		s.ciIdx = s.cursor
	case 2:
		if s.cursor < len(infraFeatureOptions) {
			key := infraFeatureOptions[s.cursor].key
			s.features[key] = !s.features[key]
		}
	}
}

// View renders the screen.
func (s *InfrastructureScreen) View() string {
	var b strings.Builder

	b.WriteString(s.Renderer().Title("Infrastructure"))
	b.WriteString("\n\n")

	// Tabs
	tabs := []string{"Hosting", "CI/CD", "Features"}
	b.WriteString(s.renderTabs(tabs, s.section))
	b.WriteString("\n\n")

	switch s.section {
	case 0:
		b.WriteString(s.renderHostingSection())
	case 1:
		b.WriteString(s.renderCISection())
	case 2:
		b.WriteString(s.renderFeaturesSection())
	}

	b.WriteString("\n")

	kb := tui.NewKeyBindings()
	kb.Add("←/→", "Switch sections")
	kb.Add("↑/↓", "Navigate")
	kb.Add("Enter/Space", "Select")
	b.WriteString(s.Renderer().HelpText(kb))

	return b.String()
}

func (s *InfrastructureScreen) renderTabs(tabs []string, active int) string {
	var parts []string
	for i, tab := range tabs {
		if i == active {
			parts = append(parts, s.Theme().Component.TabActive.Render(" "+tab+" "))
		} else {
			parts = append(parts, s.Theme().Component.Tab.Render(" "+tab+" "))
		}
	}
	return tui.JoinHorizontal(parts...)
}

func (s *InfrastructureScreen) renderHostingSection() string {
	var b strings.Builder
	b.WriteString(s.Renderer().Header("Select Hosting Platform"))
	b.WriteString("\n\n")

	for i, h := range hostingOptions {
		if i == s.cursor {
			b.WriteString(s.Renderer().ListItem("▸ "+h.name, true))
		} else {
			b.WriteString(s.Renderer().ListItem("  "+h.name, false))
		}
		b.WriteString(s.Renderer().Muted(" - "+h.description))
		b.WriteString("\n")
	}

	return b.String()
}

func (s *InfrastructureScreen) renderCISection() string {
	var b strings.Builder
	b.WriteString(s.Renderer().Header("Select CI/CD Platform"))
	b.WriteString("\n\n")

	for i, ci := range ciOptions {
		if i == s.cursor {
			b.WriteString(s.Renderer().ListItem("▸ "+ci.name, true))
		} else {
			b.WriteString(s.Renderer().ListItem("  "+ci.name, false))
		}
		b.WriteString(s.Renderer().Muted(" - "+ci.description))
		b.WriteString("\n")
	}

	return b.String()
}

func (s *InfrastructureScreen) renderFeaturesSection() string {
	var b strings.Builder
	b.WriteString(s.Renderer().Header("Infrastructure Features"))
	b.WriteString("\n\n")

	for i, feat := range infraFeatureOptions {
		checked := s.features[feat.key]
		line := s.Renderer().Checkbox(feat.name, checked)
		if i == s.cursor {
			b.WriteString("▸ " + line + "\n")
		} else {
			b.WriteString("  " + line + "\n")
		}
	}

	return b.String()
}

// ApplyToConfig applies settings to config.
func (s *InfrastructureScreen) ApplyToConfig() {
	if s.config == nil {
		return
	}

	if s.hostingIdx < len(hostingOptions) {
		s.config.Infrastructure.Hosting = strings.ToLower(strings.ReplaceAll(hostingOptions[s.hostingIdx].name, ".", ""))
	}

	if s.ciIdx < len(ciOptions) {
		ci := ciOptions[s.ciIdx]
		if ci.name == "GitHub Actions" {
			s.config.Infrastructure.CI = "github-actions"
		} else if ci.name == "GitLab CI" {
			s.config.Infrastructure.CI = "gitlab-ci"
		} else if ci.name != "None" {
			s.config.Infrastructure.CI = strings.ToLower(strings.ReplaceAll(ci.name, " ", "-"))
		}
	}

	s.config.Infrastructure.Docker = s.features["docker"]
	s.config.Infrastructure.DockerCompose = s.features["docker_compose"]
	s.config.Infrastructure.Kubernetes = s.features["kubernetes"]
	s.config.Infrastructure.CDN = s.features["cdn"]
	s.config.Infrastructure.Monitoring.Enabled = s.features["monitoring"]
}

// SetTheme sets the theme.
func (s *InfrastructureScreen) SetTheme(theme *styles.Theme) {
	s.BaseScreen.SetTheme(theme)
}

// SetConfig sets the config.
func (s *InfrastructureScreen) SetConfig(config *config.ProjectConfig) {
	s.BaseScreen.SetConfig(config)
}

// SetSize sets the size.
func (s *InfrastructureScreen) SetSize(width, height int) {
	s.BaseScreen.SetSize(width, height)
}
