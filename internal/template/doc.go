// Package template provides a powerful template engine for project scaffolding.
//
// The template engine supports:
//   - Go templates with custom functions
//   - File and directory templating
//   - Conditional file generation
//   - Template inheritance and composition
//   - Built-in helpers for common operations
//
// Usage:
//
//	engine := template.NewEngine()
//	result, err := engine.RenderFile("template.yaml.tmpl", data)
package template
