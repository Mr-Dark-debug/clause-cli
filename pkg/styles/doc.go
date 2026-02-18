// Package styles provides styling utilities for the Clause CLI.
//
// This package defines the visual design system for Clause, including:
//   - Color palette with primary, semantic, and accent colors
//   - Theme system supporting dark and light modes
//   - Typography helpers for consistent text styling
//   - Layout utilities for responsive terminal UI
//
// The package uses lipgloss for declarative styling and automatically
// adapts to terminal color capabilities (true color, 256-color, or 16-color).
//
// Example usage:
//
//	package main
//
//	import (
//	    "github.com/clause-cli/clause/pkg/styles"
//	    "github.com/charmbracelet/lipgloss"
//	)
//
//	func main() {
//	    // Get the default theme
//	    theme := styles.GetTheme()
//
//	    // Create styled text
//	    title := theme.Typography.Title.Render("Welcome to Clause")
//	    body := theme.Typography.Body.Render("AI-native project scaffolding")
//
//	    // Use color helpers
//	    success := styles.Success("Operation completed!")
//	    error := styles.Error("Something went wrong")
//
//	    // Create a card
//	    card := theme.Layout.Card.Render("Card content here")
//
//	    // Use the layout system
//	    layout := styles.NewLayout(theme, 80, 24)
//	    content := layout.BorderBox("Title", "Content", 60)
//	}
//
// # Color System
//
// The color system is designed to work across different terminal capabilities:
//   - True color terminals: Full hex color support
//   - 256-color terminals: Closest color approximation
//   - 16-color terminals: Basic ANSI color mapping
//
// Primary colors:
//   - PrimaryOrange (#FF6B35): Main brand color
//   - BackgroundNavy (#0D1117): Dark mode background
//   - TextPrimary (#F0F6FC): Primary text color
//
// Semantic colors:
//   - SuccessGreen (#3FB950): Success states
//   - WarningAmber (#D29922): Warnings
//   - ErrorRed (#F85149): Errors
//   - InfoBlue (#58A6FF): Information
//
// # Theme System
//
// Themes provide a consistent visual identity. Use GetTheme() to access
// the current theme and SetTheme() to change it.
//
//	// Use default dark theme
//	theme := styles.GetTheme()
//
//	// Switch to light theme
//	styles.SetThemeMode(styles.ModeLight)
//
// # Responsive Layout
//
// The layout system adapts to terminal size with three breakpoints:
//   - Compact (<80 columns): Simplified layouts
//   - Standard (80-120 columns): Full experience
//   - Wide (>120 columns): Enhanced layouts with more information
//
//	layout := styles.NewLayout(theme, terminalWidth, terminalHeight)
//	if layout.IsCompact() {
//	    // Use simplified layout
//	}
package styles
