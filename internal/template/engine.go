package template

import (
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/utils"
)

// Engine provides template rendering capabilities.
type Engine struct {
	// TemplateFuncs are custom template functions
	TemplateFuncs template.FuncMap

	// Delims are the left and right delimiters
	LeftDelim  string
	RightDelim string

	// MissingKeyHandling determines how missing keys are handled
	MissingKey string // "error", "default", "zero", "invalid"
}

// EngineOption is a functional option for configuring the engine.
type EngineOption func(*Engine)

// NewEngine creates a new template engine with default settings.
func NewEngine(opts ...EngineOption) *Engine {
	e := &Engine{
		LeftDelim:    "{{",
		RightDelim:   "}}",
		MissingKey:   "error",
		TemplateFuncs: make(template.FuncMap),
	}

	// Add default functions
	e.addDefaultFuncs()

	// Apply options
	for _, opt := range opts {
		opt(e)
	}

	return e
}

// WithDelims sets custom delimiters.
func WithDelims(left, right string) EngineOption {
	return func(e *Engine) {
		e.LeftDelim = left
		e.RightDelim = right
	}
}

// WithMissingKey sets how missing keys are handled.
func WithMissingKey(handling string) EngineOption {
	return func(e *Engine) {
		e.MissingKey = handling
	}
}

// WithFuncs adds custom template functions.
func WithFuncs(funcs template.FuncMap) EngineOption {
	return func(e *Engine) {
		for k, v := range funcs {
			e.TemplateFuncs[k] = v
		}
	}
}

// addDefaultFuncs adds default template functions.
func (e *Engine) addDefaultFuncs() {
	// String functions
	e.TemplateFuncs["lower"] = strings.ToLower
	e.TemplateFuncs["upper"] = strings.ToUpper
	e.TemplateFuncs["title"] = strings.Title
	e.TemplateFuncs["trim"] = strings.TrimSpace
	e.TemplateFuncs["trimPrefix"] = strings.TrimPrefix
	e.TemplateFuncs["trimSuffix"] = strings.TrimSuffix
	e.TemplateFuncs["replace"] = strings.ReplaceAll
	e.TemplateFuncs["split"] = strings.Split
	e.TemplateFuncs["join"] = strings.Join
	e.TemplateFuncs["contains"] = strings.Contains
	e.TemplateFuncs["hasPrefix"] = strings.HasPrefix
	e.TemplateFuncs["hasSuffix"] = strings.HasSuffix

	// Case conversion
	e.TemplateFuncs["camelCase"] = utils.CamelCase
	e.TemplateFuncs["pascalCase"] = utils.PascalCase
	e.TemplateFuncs["snakeCase"] = utils.SnakeCase
	e.TemplateFuncs["kebabCase"] = utils.KebabCase
	e.TemplateFuncs["screamingSnakeCase"] = utils.ScreamingSnakeCase
	e.TemplateFuncs["trainCase"] = utils.TrainCase
	e.TemplateFuncs["dotCase"] = utils.DotCase
	e.TemplateFuncs["pathCase"] = utils.PathCase

	// String manipulation
	e.TemplateFuncs["truncate"] = utils.Truncate
	e.TemplateFuncs["repeat"] = strings.Repeat
	e.TemplateFuncs["default"] = func(defaultVal, val string) string {
		if val == "" {
			return defaultVal
		}
		return val
	}

	// Boolean helpers
	e.TemplateFuncs["ternary"] = func(cond bool, ifTrue, ifFalse interface{}) interface{} {
		if cond {
			return ifTrue
		}
		return ifFalse
	}

	// Type conversion
	e.TemplateFuncs["toString"] = fmt.Sprint
	e.TemplateFuncs["quote"] = utils.Quote
	e.TemplateFuncs["squote"] = utils.SingleQuote

	// List functions
	e.TemplateFuncs["first"] = func(slice []string) string {
		if len(slice) > 0 {
			return slice[0]
		}
		return ""
	}
	e.TemplateFuncs["last"] = func(slice []string) string {
		if len(slice) > 0 {
			return slice[len(slice)-1]
		}
		return ""
	}

	// Indent helper
	e.TemplateFuncs["indent"] = func(spaces int, text string) string {
		pad := strings.Repeat(" ", spaces)
		return pad + strings.ReplaceAll(text, "\n", "\n"+pad)
	}

	// Nindent (newline + indent)
	e.TemplateFuncs["nindent"] = func(spaces int, text string) string {
		pad := strings.Repeat(" ", spaces)
		return "\n" + pad + strings.ReplaceAll(text, "\n", "\n"+pad)
	}
}

// Render renders a template string with the given data.
func (e *Engine) Render(tmpl string, data interface{}) (string, error) {
	t, err := template.New("template").
		Delims(e.LeftDelim, e.RightDelim).
		Funcs(e.TemplateFuncs).
		Option("missingkey=" + e.MissingKey).
		Parse(tmpl)
	if err != nil {
		return "", fmt.Errorf("failed to parse template: %w", err)
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute template: %w", err)
	}

	return buf.String(), nil
}

// RenderFile renders a template file with the given data.
func (e *Engine) RenderFile(path string, data interface{}) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read template file: %w", err)
	}

	return e.Render(string(content), data)
}

// RenderToFile renders a template and writes to a file.
func (e *Engine) RenderToFile(tmpl, outputPath string, data interface{}) error {
	result, err := e.Render(tmpl, data)
	if err != nil {
		return err
	}

	// Ensure directory exists
	dir := filepath.Dir(outputPath)
	if err := utils.EnsureDirectory(dir); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Write file
	if err := os.WriteFile(outputPath, []byte(result), 0644); err != nil {
		return fmt.Errorf("failed to write output file: %w", err)
	}

	return nil
}

// RenderFS renders all templates from a filesystem.
func (e *Engine) RenderFS(fsys fs.FS, root string, data interface{}, outputDir string) error {
	return fs.WalkDir(fsys, root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Read template content
		content, err := fs.ReadFile(fsys, path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		// Determine output path (remove .tmpl extension if present)
		relPath := strings.TrimPrefix(path, root+"/")
		outputPath := filepath.Join(outputDir, strings.TrimSuffix(relPath, ".tmpl"))

		// Check if this is a template file
		if strings.HasSuffix(path, ".tmpl") {
			// Render template
			result, err := e.Render(string(content), data)
			if err != nil {
				return fmt.Errorf("failed to render %s: %w", path, err)
			}

			// Ensure directory exists
			if err := utils.EnsureDirectory(filepath.Dir(outputPath)); err != nil {
				return err
			}

			// Write rendered content
			return os.WriteFile(outputPath, []byte(result), 0644)
		}

		// Copy non-template files
		if err := utils.EnsureDirectory(filepath.Dir(outputPath)); err != nil {
			return err
		}
		return os.WriteFile(outputPath, content, 0644)
	})
}

// RenderDir renders all templates in a directory.
func (e *Engine) RenderDir(templateDir, outputDir string, data interface{}) error {
	return filepath.WalkDir(templateDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if d.IsDir() {
			return nil
		}

		// Read template content
		content, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", path, err)
		}

		// Determine output path
		relPath := strings.TrimPrefix(path, templateDir)
		relPath = strings.TrimPrefix(relPath, string(os.PathSeparator))
		outputPath := filepath.Join(outputDir, strings.TrimSuffix(relPath, ".tmpl"))

		// Check if this is a template file
		if strings.HasSuffix(path, ".tmpl") {
			// Render template
			result, err := e.Render(string(content), data)
			if err != nil {
				return fmt.Errorf("failed to render %s: %w", path, err)
			}

			// Ensure directory exists
			if err := utils.EnsureDirectory(filepath.Dir(outputPath)); err != nil {
				return err
			}

			// Write rendered content
			return os.WriteFile(outputPath, []byte(result), 0644)
		}

		// Copy non-template files
		if err := utils.EnsureDirectory(filepath.Dir(outputPath)); err != nil {
			return err
		}
		return utils.CopyFile(path, outputPath)
	})
}

// RenderWithConfig renders a template with project configuration.
func (e *Engine) RenderWithConfig(tmpl string, cfg *config.ProjectConfig) (string, error) {
	data := NewTemplateData(cfg)
	return e.Render(tmpl, data)
}

// AddFunc adds a custom template function.
func (e *Engine) AddFunc(name string, fn interface{}) {
	e.TemplateFuncs[name] = fn
}

// AddFuncs adds multiple custom template functions.
func (e *Engine) AddFuncs(funcs template.FuncMap) {
	for k, v := range funcs {
		e.TemplateFuncs[k] = v
	}
}
