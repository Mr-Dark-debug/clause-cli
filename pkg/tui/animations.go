package tui

import (
	"strings"
	"sync"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/clause-cli/clause/pkg/styles"
)

// AnimationFrame represents a single frame of an animation.
type AnimationFrame struct {
	Content string
	Delay   time.Duration
}

// Animation represents a sequence of frames.
type Animation struct {
	frames    []AnimationFrame
	current   int
	loop      bool
	startTime time.Time
	mu        sync.RWMutex
}

// NewAnimation creates a new animation.
func NewAnimation(frames []AnimationFrame, loop bool) *Animation {
	return &Animation{
		frames: frames,
		loop:   loop,
	}
}

// Update advances the animation.
func (a *Animation) Update(t time.Time) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.frames) == 0 {
		return
	}

	if a.startTime.IsZero() {
		a.startTime = t
	}

	elapsed := t.Sub(a.startTime)
	var totalDuration time.Duration
	for _, f := range a.frames {
		totalDuration += f.Delay
	}

	if totalDuration == 0 {
		return
	}

	// Calculate current frame
	elapsedMs := elapsed.Milliseconds()
	totalMs := totalDuration.Milliseconds()

	if a.loop {
		elapsedMs = elapsedMs % totalMs
	} else if elapsedMs > totalMs {
		a.current = len(a.frames) - 1
		return
	}

	var accumulated int64
	for i, f := range a.frames {
		accumulated += f.Delay.Milliseconds()
		if elapsedMs < accumulated {
			a.current = i
			return
		}
	}
}

// Current returns the current frame content.
func (a *Animation) Current() string {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if len(a.frames) == 0 {
		return ""
	}
	if a.current >= len(a.frames) {
		return a.frames[len(a.frames)-1].Content
	}
	return a.frames[a.current].Content
}

// Done returns true if the animation is complete (non-looping only).
func (a *Animation) Done() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.loop {
		return false
	}
	return a.current >= len(a.frames)-1
}

// Reset resets the animation to the beginning.
func (a *Animation) Reset() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.current = 0
	a.startTime = time.Time{}
}

// SpinnerStyles contains common spinner animations.
var SpinnerStyles = map[string][]AnimationFrame{
	"dots": {
		{Content: "⠋", Delay: 80 * time.Millisecond},
		{Content: "⠙", Delay: 80 * time.Millisecond},
		{Content: "⠹", Delay: 80 * time.Millisecond},
		{Content: "⠸", Delay: 80 * time.Millisecond},
		{Content: "⠼", Delay: 80 * time.Millisecond},
		{Content: "⠴", Delay: 80 * time.Millisecond},
		{Content: "⠦", Delay: 80 * time.Millisecond},
		{Content: "⠧", Delay: 80 * time.Millisecond},
		{Content: "⠇", Delay: 80 * time.Millisecond},
		{Content: "⠏", Delay: 80 * time.Millisecond},
	},
	"line": {
		{Content: "|", Delay: 100 * time.Millisecond},
		{Content: "/", Delay: 100 * time.Millisecond},
		{Content: "-", Delay: 100 * time.Millisecond},
		{Content: "\\", Delay: 100 * time.Millisecond},
	},
	"bounce": {
		{Content: "[    ]", Delay: 100 * time.Millisecond},
		{Content: "[=   ]", Delay: 100 * time.Millisecond},
		{Content: "[==  ]", Delay: 100 * time.Millisecond},
		{Content: "[ ===]", Delay: 100 * time.Millisecond},
		{Content: "[  ==]", Delay: 100 * time.Millisecond},
		{Content: "[   =]", Delay: 100 * time.Millisecond},
		{Content: "[    ]", Delay: 100 * time.Millisecond},
		{Content: "[   =]", Delay: 100 * time.Millisecond},
		{Content: "[  ==]", Delay: 100 * time.Millisecond},
		{Content: "[ ===]", Delay: 100 * time.Millisecond},
		{Content: "[==  ]", Delay: 100 * time.Millisecond},
		{Content: "[=   ]", Delay: 100 * time.Millisecond},
	},
	"pulse": {
		{Content: "●", Delay: 200 * time.Millisecond},
		{Content: "○", Delay: 200 * time.Millisecond},
	},
	"points": {
		{Content: ".", Delay: 300 * time.Millisecond},
		{Content: "..", Delay: 300 * time.Millisecond},
		{Content: "...", Delay: 300 * time.Millisecond},
	},
}

// NewSpinner creates a new spinner animation.
func NewSpinner(style string) *Animation {
	frames, ok := SpinnerStyles[style]
	if !ok {
		frames = SpinnerStyles["dots"]
	}
	return NewAnimation(frames, true)
}

// SpinnerModel is a spinner component.
type SpinnerModel struct {
	Animation *Animation
	Style     string
	Text      string
	width     int
}

// NewSpinnerModel creates a new spinner model.
func NewSpinnerModel(style string) SpinnerModel {
	return SpinnerModel{
		Animation: NewSpinner(style),
		Style:     style,
	}
}

// Init initializes the spinner.
func (s SpinnerModel) Init() tea.Cmd {
	return Tick(80 * time.Millisecond)
}

// Update handles messages.
func (s SpinnerModel) Update(msg tea.Msg) (SpinnerModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		s.width = msg.Width
		return s, nil
	case TickMsg:
		s.Animation.Update(msg.Time)
		return s, Tick(80 * time.Millisecond)
	}
	return s, nil
}

// View renders the spinner.
func (s SpinnerModel) View() string {
	theme := styles.GetTheme()
	spinnerStyle := theme.Component.Spinner
	spinner := spinnerStyle.Render(s.Animation.Current())

	if s.Text != "" {
		return spinner + " " + s.Text
	}
	return spinner
}

// SetText sets the spinner text.
func (s *SpinnerModel) SetText(text string) {
	s.Text = text
}

// ProgressAnimation animates a progress bar.
type ProgressAnimation struct {
	percent   float64
	target    float64
	increment float64
	width     int
	filled    string
	empty     string
	animating bool
}

// NewProgressAnimation creates a new progress animation.
func NewProgressAnimation(width int) *ProgressAnimation {
	return &ProgressAnimation{
		width:     width,
		increment: 0.02,
		filled:    "█",
		empty:     "░",
	}
}

// SetPercent sets the current percentage.
func (p *ProgressAnimation) SetPercent(percent float64) {
	p.percent = percent
	if p.percent > 1 {
		p.percent = 1
	}
	if p.percent < 0 {
		p.percent = 0
	}
}

// SetTarget sets the target percentage for animation.
func (p *ProgressAnimation) SetTarget(target float64) {
	p.target = target
	if p.target > 1 {
		p.target = 1
	}
	if p.target < 0 {
		p.target = 0
	}
}

// AnimateTo starts animation toward target.
func (p *ProgressAnimation) AnimateTo(target float64) {
	p.SetTarget(target)
	p.animating = true
}

// Update advances the animation.
func (p *ProgressAnimation) Update() bool {
	if !p.animating {
		return false
	}

	if p.percent < p.target {
		p.percent += p.increment
		if p.percent >= p.target {
			p.percent = p.target
			p.animating = false
		}
		return true
	}

	p.animating = false
	return false
}

// View renders the progress bar.
func (p *ProgressAnimation) View() string {
	filledWidth := int(float64(p.width) * p.percent)
	emptyWidth := p.width - filledWidth

	theme := styles.GetTheme()
	filledStyle := theme.Component.ProgressFilled
	emptyStyle := theme.Component.Progress

	bar := filledStyle.Render(strings.Repeat(p.filled, filledWidth)) +
		emptyStyle.Render(strings.Repeat(p.empty, emptyWidth))

	return bar
}

// IsAnimating returns whether the animation is in progress.
func (p *ProgressAnimation) IsAnimating() bool {
	return p.animating
}

// FadeIn creates a fade-in animation.
func FadeIn(content string, steps int) []string {
	if steps <= 0 {
		steps = 10
	}

	frames := make([]string, steps)
	for i := 0; i < steps; i++ {
		frames[i] = content
	}
	return frames
}

// Typewriter creates a typewriter effect.
func Typewriter(text string) []AnimationFrame {
	frames := make([]AnimationFrame, len(text)+1)
	for i := 0; i <= len(text); i++ {
		frames[i] = AnimationFrame{
			Content: text[:i],
			Delay:   50 * time.Millisecond,
		}
	}
	return frames
}

// NewTypewriter creates a typewriter animation.
func NewTypewriter(text string) *Animation {
	return NewAnimation(Typewriter(text), false)
}

// TypewriterModel is a typewriter effect component.
type TypewriterModel struct {
	Animation *Animation
	Text      string
	Complete  bool
}

// NewTypewriterModel creates a new typewriter model.
func NewTypewriterModel(text string) TypewriterModel {
	return TypewriterModel{
		Animation: NewTypewriter(text),
		Text:      text,
	}
}

// Init initializes the typewriter.
func (t TypewriterModel) Init() tea.Cmd {
	return Tick(50 * time.Millisecond)
}

// Update handles messages.
func (t TypewriterModel) Update(msg tea.Msg) (TypewriterModel, tea.Cmd) {
	switch msg := msg.(type) {
	case TickMsg:
		t.Animation.Update(msg.Time)
		if t.Animation.Done() {
			t.Complete = true
			return t, nil
		}
		return t, Tick(50 * time.Millisecond)
	}
	return t, nil
}

// View renders the typewriter.
func (t TypewriterModel) View() string {
	return t.Animation.Current()
}

// Celebration creates a celebration effect.
func Celebration(width, height int) []AnimationFrame {
	// Simple celebration frames
	frames := []AnimationFrame{
		{Content: "🎉", Delay: 200 * time.Millisecond},
		{Content: "✨", Delay: 200 * time.Millisecond},
		{Content: "🎊", Delay: 200 * time.Millisecond},
		{Content: "⭐", Delay: 200 * time.Millisecond},
	}
	return frames
}

// NewCelebration creates a celebration animation.
func NewCelebration() *Animation {
	frames := []AnimationFrame{
		{Content: "🎉", Delay: 150 * time.Millisecond},
		{Content: "✨", Delay: 150 * time.Millisecond},
		{Content: "🎊", Delay: 150 * time.Millisecond},
		{Content: "⭐", Delay: 150 * time.Millisecond},
		{Content: "🌟", Delay: 150 * time.Millisecond},
		{Content: "💫", Delay: 150 * time.Millisecond},
	}
	return NewAnimation(frames, true)
}

// PulseEffect creates a pulsing effect on text.
type PulseEffect struct {
	text      string
	phase     float64
	increment float64
	color1    string
	color2    string
}

// NewPulseEffect creates a new pulse effect.
func NewPulseEffect(text string) *PulseEffect {
	return &PulseEffect{
		text:      text,
		phase:     0,
		increment: 0.1,
		color1:    styles.PrimaryOrange,
		color2:    styles.AccentPurple,
	}
}

// Update advances the pulse phase.
func (p *PulseEffect) Update() {
	p.phase += p.increment
	if p.phase > 1 {
		p.phase = 0
	}
}

// View renders the pulsed text.
func (p *PulseEffect) View() string {
	// Simple implementation: alternate between colors
	if p.phase < 0.5 {
		return styles.ColorFunc(p.color1)(p.text)
	}
	return styles.ColorFunc(p.color2)(p.text)
}

// ApplyFade applies a fade effect to content based on alpha value (0.0-1.0).
// This creates a dimming effect where 0.0 is invisible and 1.0 is fully visible.
func ApplyFade(content string, alpha float64) string {
	if alpha >= 1.0 {
		return content
	}
	if alpha <= 0.0 {
		return ""
	}

	// Simple implementation: apply dim effect
	// In a real implementation, this would manipulate ANSI codes
	// For now, we just return the content as-is if alpha is above threshold
	if alpha < 0.3 {
		return strings.Repeat(" ", len(strings.Split(content, "\n")[0]))
	}
	return content
}
