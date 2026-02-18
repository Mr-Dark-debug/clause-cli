// Package tui provides terminal UI utilities for building interactive
// command-line applications using Bubble Tea and Lip Gloss.
//
// This package builds on the excellent Charmbracelet libraries to provide
// higher-level abstractions and common components used throughout Clause.
//
// # Core Types
//
// The package provides several core types:
//
//   - BaseModel: Common functionality for all Bubble Tea models
//   - Responsive: Responsive layout handling for different terminal sizes
//   - Animation: Frame-based animation system
//   - Renderer: Common rendering utilities
//
// # Getting Started
//
// Create a base model with responsive support:
//
//	type MyModel struct {
//	    tui.BaseModel
//	    responsive *tui.Responsive
//	}
//
//	func (m MyModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
//	    switch msg := msg.(type) {
//	    case tea.WindowSizeMsg:
//	        m.HandleResize(msg)
//	        return m, nil
//	    }
//	    return m, nil
//	}
//
// # Responsive Design
//
// The Responsive type handles adaptive layouts:
//
//	responsive := tui.NewResponsive(tui.DefaultResponsiveConfig())
//	responsive.Update(width, height)
//
//	if responsive.IsCompact() {
//	    // Use compact layout
//	}
//
//	padding := responsive.Padding()
//	contentWidth := responsive.ContentWidth()
//
// # Animations
//
// Create animations from frames:
//
//	anim := tui.NewAnimation([]tui.AnimationFrame{
//	    {Content: "⠋", Delay: 80 * time.Millisecond},
//	    {Content: "⠙", Delay: 80 * time.Millisecond},
//	}, true) // loop = true
//
// Or use built-in spinner styles:
//
//	spinner := tui.NewSpinner("dots")
//
// # Rendering
//
// Use the Renderer for consistent styling:
//
//	renderer := tui.NewRenderer(theme, width, height)
//	title := renderer.Title("Welcome")
//	content := renderer.Body("Configure your project")
//	button := renderer.Button("Continue", true, false)
//
// # Key Bindings
//
// Use KeyBinding for consistent keyboard handling:
//
//	bindings := tui.NavigationKeyBindings()
//	help := bindings.Help() // Formatted help text
//
// # Focus Management
//
// FocusManager handles focus between elements:
//
//	fm := tui.NewFocusManager("input1", "input2", "button")
//	current := fm.Current()
//	fm.Next()
//	if fm.IsFocused("input1") {
//	    // Render focused state
//	}
//
// # Message Types
//
// Common message types are provided for use in Update functions:
//
//   - TickMsg: Timer tick for animations
//   - FrameMsg: Frame tick for 60fps animations
//   - FocusMsg/BlurMsg: Focus state changes
//   - CompleteMsg: Operation completion
//   - ErrorMsg: Error display
//
// # Design Philosophy
//
// This package follows these principles:
//
//  1. **Composition over Inheritance**: BaseModel is designed to be
//     embedded, not inherited.
//
//  2. **Responsive First**: All components adapt to terminal size.
//
//  3. **Theme Aware**: All rendering respects the current theme.
//
//  4. **Performance**: Minimal allocations, efficient updates.
//
//  5. **Accessibility**: Keyboard-only navigation, color-independent indicators.
package tui
