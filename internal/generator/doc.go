// Package generator provides project scaffolding capabilities.
//
// The generator creates project files and directories based on templates
// and project configuration. It handles:
//   - Directory structure creation
//   - File generation from templates
//   - Configuration file creation
//   - Git initialization
//   - Dependency installation
//
// Usage:
//
//	gen := generator.New(cfg)
//	if err := gen.Generate("/path/to/project"); err != nil {
//	    log.Fatal(err)
//	}
package generator
