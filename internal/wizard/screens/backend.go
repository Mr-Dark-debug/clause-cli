package screens

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// BackendScreen configures backend options.
type BackendScreen struct {
	BaseScreen
	section      int
	cursor       int
	enabled      bool
	frameworkIdx int
	databaseIdx  int
	apiStyleIdx  int
	features     map[string]bool
}

// Backend framework options
var backendFrameworks = []struct {
	name        string
	language    string
	description string
}{
	{"FastAPI", "python", "Modern Python web framework"},
	{"Express", "node", "Minimal Node.js framework"},
	{"NestJS", "node", "Structured Node.js framework"},
	{"Django", "python", "Full-featured Python framework"},
	{"Go Gin", "go", "Fast Go web framework"},
	{"Go Fiber", "go", "Express-like Go framework"},
	{"Rust Axum", "rust", "Ergonomic Rust web framework"},
	{"Rails", "ruby", "Ruby on Rails"},
	{"Spring", "java", "Enterprise Java framework"},
	{"Phoenix", "elixir", "Elixir web framework"},
}

// Database options
var databases = []struct {
	name        string
	description string
}{
	{"PostgreSQL", "Robust relational database"},
	{"MySQL", "Popular relational database"},
	{"SQLite", "Lightweight file database"},
	{"MongoDB", "Document database"},
	{"MariaDB", "MySQL-compatible database"},
}

// API style options
var apiStyles = []struct {
	name        string
	description string
}{
	{"REST", "Traditional REST API"},
	{"GraphQL", "Query language for APIs"},
	{"tRPC", "End-to-end typesafe APIs"},
	{"gRPC", "High-performance RPC"},
}

// Backend feature options
var backendFeatureOptions = []struct {
	key         string
	name        string
	description string
}{
	{"auth", "Authentication", "User auth system"},
	{"websocket", "WebSocket", "Real-time communication"},
	{"jobs", "Background Jobs", "Async task processing"},
	{"file_upload", "File Upload", "File handling"},
	{"email", "Email", "Email sending"},
	{"rate_limiting", "Rate Limiting", "API rate limiting"},
	{"logging", "Logging", "Structured logging"},
	{"metrics", "Metrics", "Performance metrics"},
}

// NewBackendScreen creates a new backend screen.
func NewBackendScreen() *BackendScreen {
	return &BackendScreen{
		BaseScreen:   *NewBaseScreen("Backend", "backend"),
		enabled:      true,
		frameworkIdx: 0,
		databaseIdx:  0,
		apiStyleIdx:  0,
		features: map[string]bool{
			"auth":          true,
			"websocket":     false,
			"jobs":          false,
			"file_upload":   false,
			"email":         false,
			"rate_limiting": true,
			"logging":       true,
			"metrics":       false,
		},
		section: 0,
		cursor:  0,
	}
}

// Init initializes the screen.
func (s *BackendScreen) Init() tea.Cmd {
	return nil
}

// Update handles updates.
func (s *BackendScreen) Update(msg tea.Msg) tea.Cmd {
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
			if s.section < 4 {
				s.section++
				s.cursor = 0
			}
		case "enter", " ":
			s.toggle()
		case "tab":
			if s.section < 4 {
				s.section++
				s.cursor = 0
			}
		}
	}

	s.complete = true
	return nil
}

func (s *BackendScreen) getMaxItems() int {
	switch s.section {
	case 0:
		return 2
	case 1:
		return len(backendFrameworks)
	case 2:
		return len(databases)
	case 3:
		return len(apiStyles)
	case 4:
		return len(backendFeatureOptions)
	}
	return 0
}

func (s *BackendScreen) toggle() {
	switch s.section {
	case 0:
		s.enabled = s.cursor == 0
	case 1:
		s.frameworkIdx = s.cursor
	case 2:
		s.databaseIdx = s.cursor
	case 3:
		s.apiStyleIdx = s.cursor
	case 4:
		if s.cursor < len(backendFeatureOptions) {
			key := backendFeatureOptions[s.cursor].key
			s.features[key] = !s.features[key]
		}
	}
}

// View renders the screen.
func (s *BackendScreen) View() string {
	var b strings.Builder

	b.WriteString(s.Renderer().Title("Backend Configuration"))
	b.WriteString("\n\n")

	// Section tabs
	tabs := []string{"Enable", "Framework", "Database", "API", "Features"}
	b.WriteString(s.renderTabs(tabs, s.section))
	b.WriteString("\n\n")

	switch s.section {
	case 0:
		b.WriteString(s.renderEnableSection())
	case 1:
		b.WriteString(s.renderFrameworkSection())
	case 2:
		b.WriteString(s.renderDatabaseSection())
	case 3:
		b.WriteString(s.renderAPISection())
	case 4:
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

func (s *BackendScreen) renderTabs(tabs []string, active int) string {
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

func (s *BackendScreen) renderEnableSection() string {
	var b strings.Builder
	b.WriteString(s.Renderer().Header("Enable Backend?"))
	b.WriteString("\n\n")

	options := []string{"Yes, include backend", "No, skip backend"}
	for i, opt := range options {
		selected := (i == 0 && s.enabled) || (i == 1 && !s.enabled)
		b.WriteString(s.Renderer().RadioButton(opt, selected))
		b.WriteString("\n")
	}

	return b.String()
}

func (s *BackendScreen) renderFrameworkSection() string {
	var b strings.Builder

	if !s.enabled {
		return s.Renderer().Muted("Enable backend to select a framework")
	}

	b.WriteString(s.Renderer().Header("Select Framework"))
	b.WriteString("\n\n")

	for i, fw := range backendFrameworks {
		prefix := "  "
		if i == s.cursor {
			prefix = "▸ "
		}
		b.WriteString(s.Renderer().ListItem(prefix+fw.name+" ("+fw.language+")", i == s.cursor))
		b.WriteString(s.Renderer().Muted(" - "+fw.description))
		b.WriteString("\n")
	}

	return b.String()
}

func (s *BackendScreen) renderDatabaseSection() string {
	var b strings.Builder

	if !s.enabled {
		return s.Renderer().Muted("Enable backend to select a database")
	}

	b.WriteString(s.Renderer().Header("Select Database"))
	b.WriteString("\n\n")

	for i, db := range databases {
		if i == s.cursor {
			b.WriteString(s.Renderer().ListItem("▸ "+db.name, true))
		} else {
			b.WriteString(s.Renderer().ListItem("  "+db.name, false))
		}
		b.WriteString(s.Renderer().Muted(" - "+db.description))
		b.WriteString("\n")
	}

	return b.String()
}

func (s *BackendScreen) renderAPISection() string {
	var b strings.Builder

	if !s.enabled {
		return s.Renderer().Muted("Enable backend to select API style")
	}

	b.WriteString(s.Renderer().Header("Select API Style"))
	b.WriteString("\n\n")

	for i, style := range apiStyles {
		if i == s.cursor {
			b.WriteString(s.Renderer().ListItem("▸ "+style.name, true))
		} else {
			b.WriteString(s.Renderer().ListItem("  "+style.name, false))
		}
		b.WriteString(s.Renderer().Muted(" - "+style.description))
		b.WriteString("\n")
	}

	return b.String()
}

func (s *BackendScreen) renderFeaturesSection() string {
	var b strings.Builder

	if !s.enabled {
		return s.Renderer().Muted("Enable backend to select features")
	}

	b.WriteString(s.Renderer().Header("Select Features"))
	b.WriteString("\n\n")

	for i, feat := range backendFeatureOptions {
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
func (s *BackendScreen) ApplyToConfig() {
	if s.config == nil {
		return
	}

	s.config.Backend.Enabled = s.enabled

	if s.enabled && s.frameworkIdx < len(backendFrameworks) {
		fw := backendFrameworks[s.frameworkIdx]
		s.config.Backend.Framework = strings.ToLower(strings.ReplaceAll(fw.name, " ", "-"))
		s.config.Backend.Language = fw.language
	}

	if s.databaseIdx < len(databases) {
		s.config.Backend.Database.Primary = strings.ToLower(databases[s.databaseIdx].name)
	}

	if s.apiStyleIdx < len(apiStyles) {
		s.config.Backend.API.Style = strings.ToLower(apiStyles[s.apiStyleIdx].name)
	}

	s.config.Backend.Features.WebSocket = s.features["websocket"]
	s.config.Backend.Features.BackgroundJobs = s.features["jobs"]
	s.config.Backend.Features.FileUpload = s.features["file_upload"]
	s.config.Backend.Features.Email = s.features["email"]
	s.config.Backend.Features.RateLimiting = s.features["rate_limiting"]
	s.config.Backend.Features.Logging = s.features["logging"]
	s.config.Backend.Features.Metrics = s.features["metrics"]
}

// SetTheme sets the theme.
func (s *BackendScreen) SetTheme(theme *styles.Theme) {
	s.BaseScreen.SetTheme(theme)
}

// SetConfig sets the config.
func (s *BackendScreen) SetConfig(config *config.ProjectConfig) {
	s.BaseScreen.SetConfig(config)
}

// SetSize sets the size.
func (s *BackendScreen) SetSize(width, height int) {
	s.BaseScreen.SetSize(width, height)
}
