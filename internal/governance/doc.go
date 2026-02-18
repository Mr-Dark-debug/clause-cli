// Package governance provides AI governance and compliance management.
//
// The governance system helps manage AI-assisted development by:
//   - Maintaining consistent context for AI assistants
//   - Tracking components and their relationships
//   - Generating prompt guidelines
//   - Managing brainstorming documents
//
// Usage:
//
//	gov := governance.New(projectPath)
//	if err := gov.Initialize(); err != nil {
//	    log.Fatal(err)
//	}
package governance
