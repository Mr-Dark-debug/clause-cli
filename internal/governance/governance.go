package governance

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/clause-cli/clause/internal/config"
	"github.com/clause-cli/clause/pkg/output"
)

// Governance manages AI governance and compliance for a project.
type Governance struct {
	// ProjectPath is the root path of the project
	ProjectPath string

	// Config is the project configuration
	Config *config.ProjectConfig

	// Registry is the component registry
	Registry *ComponentRegistry

	// ContextManager handles AI context
	ContextManager *ContextManager

	// Logger for output
	Logger *output.Logger

	// mu protects concurrent access
	mu sync.RWMutex
}

// GovernanceOption is a functional option for configuring governance.
type GovernanceOption func(*Governance)

// New creates a new governance manager.
func New(projectPath string, opts ...GovernanceOption) *Governance {
	g := &Governance{
		ProjectPath: projectPath,
		Logger:      output.DefaultLogger,
	}

	for _, opt := range opts {
		opt(g)
	}

	// Initialize registry
	g.Registry = NewComponentRegistry()

	// Initialize context manager
	g.ContextManager = NewContextManager(projectPath)

	return g
}

// WithConfig sets the project configuration.
func WithConfig(cfg *config.ProjectConfig) GovernanceOption {
	return func(g *Governance) {
		g.Config = cfg
	}
}

// WithLogger sets the logger.
func WithLogger(logger *output.Logger) GovernanceOption {
	return func(g *Governance) {
		g.Logger = logger
	}
}

// Initialize sets up governance for a project.
func (g *Governance) Initialize() error {
	g.mu.Lock()
	defer g.mu.Unlock()

	// Create .clause directory
	clauseDir := filepath.Join(g.ProjectPath, ".clause")
	if err := os.MkdirAll(clauseDir, 0755); err != nil {
		return fmt.Errorf("failed to create .clause directory: %w", err)
	}

	// Initialize context
	if err := g.ContextManager.Initialize(); err != nil {
		return fmt.Errorf("failed to initialize context: %w", err)
	}

	// Create governance files if enabled
	if g.Config != nil && g.Config.Governance.Enabled {
		gen := NewGenerator(g.ProjectPath, g.Config)
		if err := gen.Generate(); err != nil {
			return fmt.Errorf("failed to generate governance files: %w", err)
		}
	}

	g.Logger.Info("Governance initialized successfully")
	return nil
}

// RegisterComponent registers a new component in the registry.
func (g *Governance) RegisterComponent(comp Component) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if err := g.Registry.Register(comp); err != nil {
		return err
	}

	// Update context with new component
	if err := g.ContextManager.UpdateComponentContext(comp); err != nil {
		g.Logger.Warn("Failed to update context for component: %v", err)
	}

	return nil
}

// UnregisterComponent removes a component from the registry.
func (g *Governance) UnregisterComponent(name string) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	if err := g.Registry.Unregister(name); err != nil {
		return err
	}

	// Remove from context
	if err := g.ContextManager.RemoveComponentContext(name); err != nil {
		g.Logger.Warn("Failed to remove component context: %v", err)
	}

	return nil
}

// GetComponent retrieves a component by name.
func (g *Governance) GetComponent(name string) (Component, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.Registry.Get(name)
}

// ListComponents returns all registered components.
func (g *Governance) ListComponents() []Component {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.Registry.List()
}

// GetContext returns the current AI context.
func (g *Governance) GetContext() (*AIContext, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()

	return g.ContextManager.GetContext()
}

// UpdateContext updates the AI context.
func (g *Governance) UpdateContext(updates ContextUpdates) error {
	g.mu.Lock()
	defer g.mu.Unlock()

	return g.ContextManager.Update(updates)
}

// Validate validates the governance configuration.
func (g *Governance) Validate() error {
	g.mu.RLock()
	defer g.mu.RUnlock()

	// Check if .clause directory exists
	clauseDir := filepath.Join(g.ProjectPath, ".clause")
	if _, err := os.Stat(clauseDir); os.IsNotExist(err) {
		return fmt.Errorf("governance not initialized: .clause directory not found")
	}

	// Validate registry
	if errs := g.Registry.Validate(); len(errs) > 0 {
		return fmt.Errorf("registry validation failed: %v", errs)
	}

	return nil
}

// Status returns the current governance status.
func (g *Governance) Status() *Status {
	g.mu.RLock()
	defer g.mu.RUnlock()

	status := &Status{
		Initialized:     g.isInitialized(),
		ContextLevel:    "standard",
		ComponentsCount: len(g.Registry.List()),
	}

	if g.Config != nil {
		status.ContextLevel = g.Config.Governance.ContextLevel
		status.ComponentRegistry = g.Config.Governance.ComponentRegistry
		status.BrainstormMd = g.Config.Governance.BrainstormMd
		status.PromptGuidelines = g.Config.Governance.PromptGuidelines
	}

	// Check if context exists
	if ctx, err := g.ContextManager.GetContext(); err == nil {
		status.LastContextUpdate = ctx.UpdatedAt
	}

	return status
}

// isInitialized checks if governance is initialized.
func (g *Governance) isInitialized() bool {
	clauseDir := filepath.Join(g.ProjectPath, ".clause")
	if _, err := os.Stat(clauseDir); err != nil {
		return false
	}

	contextFile := filepath.Join(clauseDir, "context.yaml")
	if _, err := os.Stat(contextFile); err != nil {
		return false
	}

	return true
}

// Status represents governance status.
type Status struct {
	// Initialized indicates if governance is set up
	Initialized bool `json:"initialized"`

	// ContextLevel is the AI context detail level
	ContextLevel string `json:"context_level"`

	// ComponentRegistry indicates if component registry is enabled
	ComponentRegistry bool `json:"component_registry"`

	// BrainstormMd indicates if Brainstorm.md is enabled
	BrainstormMd bool `json:"brainstorm_md"`

	// PromptGuidelines indicates if prompt guidelines exist
	PromptGuidelines bool `json:"prompt_guidelines"`

	// ComponentsCount is the number of registered components
	ComponentsCount int `json:"components_count"`

	// LastContextUpdate is when context was last updated
	LastContextUpdate string `json:"last_context_update,omitempty"`
}
