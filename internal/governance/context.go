package governance

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

// AIContext represents the AI context for a project.
type AIContext struct {
	// ProjectName is the project name
	ProjectName string `yaml:"project_name" json:"project_name"`

	// Description is the project description
	Description string `yaml:"description" json:"description"`

	// TechStack lists all technologies used
	TechStack []string `yaml:"tech_stack" json:"tech_stack"`

	// Architecture describes the project architecture
	Architecture ArchitectureInfo `yaml:"architecture" json:"architecture"`

	// Conventions lists coding conventions
	Conventions []Convention `yaml:"conventions" json:"conventions"`

	// Patterns lists design patterns used
	Patterns []string `yaml:"patterns" json:"patterns"`

	// KeyFiles lists important files to know about
	KeyFiles []KeyFile `yaml:"key_files" json:"key_files"`

	// Components lists registered components
	Components []ComponentSummary `yaml:"components" json:"components"`

	// BestPractices lists project-specific best practices
	BestPractices []string `yaml:"best_practices" json:"best_practices"`

	// KnownIssues lists known issues or limitations
	KnownIssues []string `yaml:"known_issues,omitempty" json:"known_issues,omitempty"`

	// CreatedAt is when context was created
	CreatedAt string `yaml:"created_at" json:"created_at"`

	// UpdatedAt is when context was last updated
	UpdatedAt string `yaml:"updated_at" json:"updated_at"`
}

// ArchitectureInfo contains architecture details.
type ArchitectureInfo struct {
	// Style is the architectural style (monolith, microservices, serverless, etc.)
	Style string `yaml:"style" json:"style"`

	// Frontend describes the frontend architecture
	Frontend string `yaml:"frontend,omitempty" json:"frontend,omitempty"`

	// Backend describes the backend architecture
	Backend string `yaml:"backend,omitempty" json:"backend,omitempty"`

	// Database describes the database setup
	Database string `yaml:"database,omitempty" json:"database,omitempty"`

	// Layers describes the application layers
	Layers []string `yaml:"layers,omitempty" json:"layers,omitempty"`
}

// Convention represents a coding convention.
type Convention struct {
	// Category is the convention category (naming, structure, etc.)
	Category string `yaml:"category" json:"category"`

	// Rule describes the convention rule
	Rule string `yaml:"rule" json:"rule"`

	// Examples provides examples
	Examples []string `yaml:"examples,omitempty" json:"examples,omitempty"`
}

// KeyFile represents an important file.
type KeyFile struct {
	// Path is the file path
	Path string `yaml:"path" json:"path"`

	// Purpose describes what the file is for
	Purpose string `yaml:"purpose" json:"purpose"`
}

// ComponentSummary is a summary of a registered component.
type ComponentSummary struct {
	// Name is the component name
	Name string `yaml:"name" json:"name"`

	// Type is the component type
	Type string `yaml:"type" json:"type"`

	// Description is the component description
	Description string `yaml:"description" json:"description"`
}

// ContextUpdates contains updates to apply to context.
type ContextUpdates struct {
	// TechStack adds technologies to the stack
	TechStack []string `yaml:"tech_stack,omitempty" json:"tech_stack,omitempty"`

	// Patterns adds design patterns
	Patterns []string `yaml:"patterns,omitempty" json:"patterns,omitempty"`

	// BestPractices adds best practices
	BestPractices []string `yaml:"best_practices,omitempty" json:"best_practices,omitempty"`

	// KnownIssues adds known issues
	KnownIssues []string `yaml:"known_issues,omitempty" json:"known_issues,omitempty"`

	// Conventions adds conventions
	Conventions []Convention `yaml:"conventions,omitempty" json:"conventions,omitempty"`
}

// ContextManager manages the AI context for a project.
type ContextManager struct {
	// ProjectPath is the root path of the project
	ProjectPath string

	// contextFile is the path to the context file
	contextFile string

	// cache holds the loaded context
	cache *AIContext
}

// NewContextManager creates a new context manager.
func NewContextManager(projectPath string) *ContextManager {
	return &ContextManager{
		ProjectPath: projectPath,
		contextFile: filepath.Join(projectPath, ".clause", "context.yaml"),
	}
}

// Initialize creates the initial context.
func (m *ContextManager) Initialize() error {
	// Check if context already exists
	if _, err := os.Stat(m.contextFile); err == nil {
		// Load existing context
		return m.load()
	}

	// Create initial context
	ctx := &AIContext{
		TechStack:      []string{},
		Conventions:    []Convention{},
		Patterns:       []string{},
		KeyFiles:       []KeyFile{},
		Components:     []ComponentSummary{},
		BestPractices:  []string{},
		KnownIssues:    []string{},
		CreatedAt:      time.Now().Format(time.RFC3339),
		UpdatedAt:      time.Now().Format(time.RFC3339),
	}

	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(m.contextFile), 0755); err != nil {
		return fmt.Errorf("failed to create context directory: %w", err)
	}

	m.cache = ctx
	return m.save()
}

// GetContext returns the current AI context.
func (m *ContextManager) GetContext() (*AIContext, error) {
	if m.cache == nil {
		if err := m.load(); err != nil {
			return nil, err
		}
	}
	return m.cache, nil
}

// Update applies updates to the context.
func (m *ContextManager) Update(updates ContextUpdates) error {
	if err := m.load(); err != nil {
		return err
	}

	// Apply tech stack updates
	for _, tech := range updates.TechStack {
		if !contains(m.cache.TechStack, tech) {
			m.cache.TechStack = append(m.cache.TechStack, tech)
		}
	}

	// Apply pattern updates
	for _, pattern := range updates.Patterns {
		if !contains(m.cache.Patterns, pattern) {
			m.cache.Patterns = append(m.cache.Patterns, pattern)
		}
	}

	// Apply best practices
	for _, practice := range updates.BestPractices {
		if !contains(m.cache.BestPractices, practice) {
			m.cache.BestPractices = append(m.cache.BestPractices, practice)
		}
	}

	// Apply known issues
	for _, issue := range updates.KnownIssues {
		if !contains(m.cache.KnownIssues, issue) {
			m.cache.KnownIssues = append(m.cache.KnownIssues, issue)
		}
	}

	// Apply conventions
	for _, conv := range updates.Conventions {
		m.cache.Conventions = append(m.cache.Conventions, conv)
	}

	m.cache.UpdatedAt = time.Now().Format(time.RFC3339)

	return m.save()
}

// UpdateComponentContext updates context when a component changes.
func (m *ContextManager) UpdateComponentContext(comp Component) error {
	if err := m.load(); err != nil {
		return err
	}

	// Update or add component summary
	found := false
	summary := ComponentSummary{
		Name:        comp.Name,
		Type:        comp.Type,
		Description: comp.Description,
	}

	for i, c := range m.cache.Components {
		if c.Name == comp.Name {
			m.cache.Components[i] = summary
			found = true
			break
		}
	}

	if !found {
		m.cache.Components = append(m.cache.Components, summary)
	}

	// Add tech stack from component
	for _, tech := range comp.TechStack {
		if !contains(m.cache.TechStack, tech) {
			m.cache.TechStack = append(m.cache.TechStack, tech)
		}
	}

	m.cache.UpdatedAt = time.Now().Format(time.RFC3339)

	return m.save()
}

// RemoveComponentContext removes a component from context.
func (m *ContextManager) RemoveComponentContext(name string) error {
	if err := m.load(); err != nil {
		return err
	}

	// Remove component summary
	var updated []ComponentSummary
	for _, c := range m.cache.Components {
		if c.Name != name {
			updated = append(updated, c)
		}
	}
	m.cache.Components = updated

	m.cache.UpdatedAt = time.Now().Format(time.RFC3339)

	return m.save()
}

// SetProjectInfo sets basic project information in context.
func (m *ContextManager) SetProjectInfo(name, description string) error {
	if err := m.load(); err != nil {
		return err
	}

	m.cache.ProjectName = name
	m.cache.Description = description
	m.cache.UpdatedAt = time.Now().Format(time.RFC3339)

	return m.save()
}

// SetArchitecture sets architecture information.
func (m *ContextManager) SetArchitecture(arch ArchitectureInfo) error {
	if err := m.load(); err != nil {
		return err
	}

	m.cache.Architecture = arch
	m.cache.UpdatedAt = time.Now().Format(time.RFC3339)

	return m.save()
}

// AddKeyFile adds a key file to the context.
func (m *ContextManager) AddKeyFile(path, purpose string) error {
	if err := m.load(); err != nil {
		return err
	}

	kf := KeyFile{Path: path, Purpose: purpose}

	// Check if already exists
	for i, f := range m.cache.KeyFiles {
		if f.Path == path {
			m.cache.KeyFiles[i] = kf
			return m.save()
		}
	}

	m.cache.KeyFiles = append(m.cache.KeyFiles, kf)
	m.cache.UpdatedAt = time.Now().Format(time.RFC3339)

	return m.save()
}

// GeneratePrompt generates an AI context prompt.
func (m *ContextManager) GeneratePrompt() (string, error) {
	if err := m.load(); err != nil {
		return "", err
	}

	var sb strings.Builder

	sb.WriteString("# Project Context\n\n")

	if m.cache.ProjectName != "" {
		sb.WriteString(fmt.Sprintf("**Project:** %s\n\n", m.cache.ProjectName))
	}

	if m.cache.Description != "" {
		sb.WriteString(fmt.Sprintf("**Description:** %s\n\n", m.cache.Description))
	}

	// Tech Stack
	if len(m.cache.TechStack) > 0 {
		sb.WriteString("## Technology Stack\n\n")
		for _, tech := range m.cache.TechStack {
			sb.WriteString(fmt.Sprintf("- %s\n", tech))
		}
		sb.WriteString("\n")
	}

	// Architecture
	if m.cache.Architecture.Style != "" {
		sb.WriteString("## Architecture\n\n")
		sb.WriteString(fmt.Sprintf("**Style:** %s\n\n", m.cache.Architecture.Style))

		if m.cache.Architecture.Frontend != "" {
			sb.WriteString(fmt.Sprintf("**Frontend:** %s\n\n", m.cache.Architecture.Frontend))
		}
		if m.cache.Architecture.Backend != "" {
			sb.WriteString(fmt.Sprintf("**Backend:** %s\n\n", m.cache.Architecture.Backend))
		}
		if m.cache.Architecture.Database != "" {
			sb.WriteString(fmt.Sprintf("**Database:** %s\n\n", m.cache.Architecture.Database))
		}
	}

	// Components
	if len(m.cache.Components) > 0 {
		sb.WriteString("## Components\n\n")
		for _, comp := range m.cache.Components {
			sb.WriteString(fmt.Sprintf("- **%s** (%s): %s\n", comp.Name, comp.Type, comp.Description))
		}
		sb.WriteString("\n")
	}

	// Best Practices
	if len(m.cache.BestPractices) > 0 {
		sb.WriteString("## Best Practices\n\n")
		for _, practice := range m.cache.BestPractices {
			sb.WriteString(fmt.Sprintf("- %s\n", practice))
		}
		sb.WriteString("\n")
	}

	// Known Issues
	if len(m.cache.KnownIssues) > 0 {
		sb.WriteString("## Known Issues\n\n")
		for _, issue := range m.cache.KnownIssues {
			sb.WriteString(fmt.Sprintf("- %s\n", issue))
		}
		sb.WriteString("\n")
	}

	return sb.String(), nil
}

// load loads the context from file.
func (m *ContextManager) load() error {
	data, err := os.ReadFile(m.contextFile)
	if err != nil {
		return fmt.Errorf("failed to read context file: %w", err)
	}

	var ctx AIContext
	if err := yaml.Unmarshal(data, &ctx); err != nil {
		return fmt.Errorf("failed to parse context file: %w", err)
	}

	m.cache = &ctx
	return nil
}

// save saves the context to file.
func (m *ContextManager) save() error {
	data, err := yaml.Marshal(m.cache)
	if err != nil {
		return fmt.Errorf("failed to marshal context: %w", err)
	}

	return os.WriteFile(m.contextFile, data, 0644)
}

// contains checks if a string is in a slice.
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
