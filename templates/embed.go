// Package templates provides embedded template files for Clause.
package templates

import (
	"embed"
	"io/fs"
)

//go:embed governance/*.tmpl
var templatesFS embed.FS

// FS returns the embedded filesystem for templates.
func FS() (fs.FS, error) {
	return fs.Sub(templatesFS, "governance")
}

// GetTemplate retrieves a template by name.
func GetTemplate(name string) ([]byte, error) {
	return templatesFS.ReadFile("governance/" + name)
}

// ListTemplates returns a list of available template names.
func ListTemplates() ([]string, error) {
	entries, err := fs.ReadDir(templatesFS, "governance")
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			names = append(names, entry.Name())
		}
	}

	return names, nil
}
