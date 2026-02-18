package screens

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/styles"
	"github.com/clause-cli/clause/pkg/tui"
)

// FrontendScreen configures frontend options.
type FrontendScreen struct {
	BaseScreen
	cursor       int
	section      int // 0: enabled, 1: framework, 2: styling, 3: features
	enabled      bool
	frameworkIdx int
	stylingIdx   int
	features     map[string]bool
}

// Framework options with descriptions
var frameworks = []struct {
	name        string
	description string
}{
	{"React", "Component-based UI library"},
	{"Vue", "Progressive JavaScript framework"},
	{"Svelte", "Compiled frontend framework"},
	{"Next.js", "React with SSR and routing"},
	{"Nuxt", "Vue with SSR and routing"},
	{"SvelteKit", "Svelte with SSR and routing"},
	{"Angular", "Full-featured framework"},
	{"Remix", "React with web fundamentals"},
	{"Astro", "Static site generator"},
	{"Solid", "Reactive UI library"},
}

// Styling options
var stylingOptions = []struct {
	name        string
	description string
}{
	{"Tailwind CSS", "Utility-first CSS framework"},
	{"CSS Modules", "Scoped CSS for components"},
	{"Styled Components", "CSS-in-JS styling"},
	{"SCSS/Sass", "CSS preprocessor"},
	{"Emotion", "CSS-in-JS library"},
}

// Frontend feature options
var frontendFeatureOptions = []struct {
	key         string
	name        string
	description string
}{
	{"typescript", "TypeScript", "Type-safe JavaScript"},
	{"ssr", "SSR", "Server-side rendering"},
	{"ssg", "SSG", "Static site generation"},
	{"pwa", "PWA", "Progressive web app"},
	{"i18n", "i18n", "Internationalization"},
	{"dark_mode", "Dark Mode", "Theme switching"},
	{"storybook", "Storybook", "Component documentation"},
}

// NewFrontendScreen creates a new frontend screen.
func NewFrontendScreen() *FrontendScreen {
	s := &FrontendScreen{
		BaseScreen:  *NewBaseScreen("Frontend", "frontend"),
		enabled:     true,
		frameworkIdx: 0,
		stylingIdx:   0,
		features: map[string]bool{
			"typescript": true,
			"ssr":        false,
			"ssg":        false,
			"pwa":        false,
			"i18n":       false,
			"dark_mode":  true,
			"storybook":  false,
		},
		section: 0,
		cursor:  0,
	}

	return s
}

// Init initializes the screen.
func (s *FrontendScreen) Init() tea.Cmd {
	return nil
}

// Update handles updates for the screen.
func (s *FrontendScreen) Update(msg tea.Msg) tea.Cmd {
	switch m := msg.(type) {
	case tea.KeyMsg:
		switch m.String() {
		case "up", "k":
			s.moveUp()
		case "down", "j":
			s.moveDown()
		case "left", "h":
			if s.section > 0 {
				s.section--
				s.cursor = 0
			}
		case "right", "l":
			if s.section < 3 {
				s.section++
				s.cursor = 0
			}
		case "enter", " ":
			s.toggle()
		case "tab":
			if s.section < 3 {
				s.section++
				s.cursor = 0
			}
		}
	}

	s.updateCompletion()
	return nil
}

func (s *FrontendScreen) moveUp() {
	switch s.section {
	case 0:
		// No up for enabled toggle
	case 1:
		if s.cursor > 0 {
			s.cursor--
		}
	case 2:
		if s.cursor > 0 {
			s.cursor--
		}
	case 3:
		if s.cursor > 0 {
			s.cursor--
		}
	}
}

func (s *FrontendScreen) moveDown() {
	switch s.section {
	case 0:
		// No down for enabled toggle
	case 1:
		if s.cursor < len(frameworks)-1 {
			s.cursor++
		}
	case 2:
		if s.cursor < len(stylingOptions)-1 {
			s.cursor++
		}
	case 3:
		if s.cursor < len(frontendFeatureOptions)-1 {
			s.cursor++
		}
	}
}

func (s *FrontendScreen) toggle() {
	switch s.section {
	case 0:
		s.enabled = s.cursor == 0
	case 1:
		s.frameworkIdx = s.cursor
	case 2:
		s.stylingIdx = s.cursor
	case 3:
		if s.cursor < len(frontendFeatureOptions) {
			key := frontendFeatureOptions[s.cursor].key
			s.features[key] = !s.features[key]
		}
	}
}

func (s *FrontendScreen) updateCompletion() {
	s.complete = true
}

// View renders the screen.
func (s *FrontendScreen) View() string {
	var b strings.Builder

	b.WriteString(s.Renderer().Title("Frontend Configuration"))
	b.WriteString("\n\n")

	if !s.enabled {
		b.WriteString(s.Renderer().Muted("Frontend is disabled"))
		b.WriteString("\n\n")
	}

	// Section tabs
	tabs := []string{"Enable", "Framework", "Styling", "Features"}
	b.WriteString(s.renderTabs(tabs, s.section))
	b.WriteString("\n\n")

	// Render current section
	switch s.section {
	case 0:
		b.WriteString(s.renderEnableSection())
	case 1:
		b.WriteString(s.renderFrameworkSection())
	case 2:
		b.WriteString(s.renderStylingSection())
	case 3:
		b.WriteString(s.renderFeaturesSection())
	}

	b.WriteString("\n")

	// Help
	kb := tui.NewKeyBindings()
	kb.Add("←/→", "Switch sections")
	kb.Add("↑/↓", "Navigate options")
	kb.Add("Enter/Space", "Select")
	if s.IsComplete() {
		kb.Add("Enter", "Continue")
	}
	b.WriteString(s.Renderer().HelpText(kb))

	return b.String()
}

func (s *FrontendScreen) renderTabs(tabs []string, active int) string {
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

func (s *FrontendScreen) renderEnableSection() string {
	var b strings.Builder
	b.WriteString(s.Renderer().Header("Enable Frontend?"))
	b.WriteString("\n\n")

	options := []string{"Yes, include frontend", "No, skip frontend"}
	for i, opt := range options {
		selected := (i == 0 && s.enabled) || (i == 1 && !s.enabled)
		b.WriteString(s.Renderer().RadioButton(opt, selected))
		b.WriteString("\n")
	}

	return b.String()
}

func (s *FrontendScreen) renderFrameworkSection() string {
	var b strings.Builder

	if !s.enabled {
		return s.Renderer().Muted("Enable frontend to select a framework")
	}

	b.WriteString(s.Renderer().Header("Select Framework"))
	b.WriteString("\n\n")

	for i, fw := range frameworks {
		if i == s.cursor {
			b.WriteString(s.Renderer().ListItem("▸ "+fw.name, true))
		} else {
			b.WriteString(s.Renderer().ListItem("  "+fw.name, false))
		}
		b.WriteString(s.Renderer().Muted(" - "+fw.description))
		b.WriteString("\n")
	}

	return b.String()
}

func (s *FrontendScreen) renderStylingSection() string {
	var b strings.Builder

	if !s.enabled {
		return s.Renderer().Muted("Enable frontend to select styling")
	}

	b.WriteString(s.Renderer().Header("Select Styling Approach"))
	b.WriteString("\n\n")

	for i, style := range stylingOptions {
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

func (s *FrontendScreen) renderFeaturesSection() string {
	var b strings.Builder

	if !s.enabled {
		return s.Renderer().Muted("Enable frontend to select features")
	}

	b.WriteString(s.Renderer().Header("Select Features"))
	b.WriteString("\n\n")

	for i, feat := range frontendFeatureOptions {
		checked := s.features[feat.key]
		line := s.Renderer().Checkbox(feat.name+" "+s.Renderer().Muted("- "+feat.description), checked)
		if i == s.cursor {
			b.WriteString("▸ " + line + "\n")
		} else {
			b.WriteString("  " + line + "\n")
		}
	}

	return b.String()
}

// ApplyToConfig applies settings to config.
func (s *FrontendScreen) ApplyToConfig() {
	if s.config == nil {
		return
	}

	s.config.Frontend.Enabled = s.enabled

	if s.enabled {
		s.config.Frontend.Framework = strings.ToLower(strings.ReplaceAll(frameworks[s.frameworkIdx].name, ".", ""))
		s.config.Frontend.Styling = strings.ToLower(strings.Split(stylingOptions[s.stylingIdx].name, " ")[0])
		s.config.Frontend.TypeScript = s.features["typescript"]
		s.config.Frontend.Features.SSR = s.features["ssr"]
		s.config.Frontend.Features.SSG = s.features["ssg"]
		s.config.Frontend.Features.PWA = s.features["pwa"]
		s.config.Frontend.Features.I18n = s.features["i18n"]
		s.config.Frontend.Features.DarkMode = s.features["dark_mode"]
		s.config.Frontend.Features.Storybook = s.features["storybook"]
	}
}

// SetTheme sets the theme.
func (s *FrontendScreen) SetTheme(theme *styles.Theme) {
	s.BaseScreen.SetTheme(theme)
}

// SetConfig sets the config.
func (s *FrontendScreen) SetConfig(config *config.ProjectConfig) {
	s.BaseScreen.SetConfig(config)
}

// SetSize sets the size.
func (s *FrontendScreen) SetSize(width, height int) {
	s.BaseScreen.SetSize(width, height)
}
