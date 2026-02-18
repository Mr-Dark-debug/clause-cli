package governance

import (
	"fmt"
	"sort"
	"sync"
)

// Component represents a project component in the registry.
type Component struct {
	// Name is the unique component identifier
	Name string `yaml:"name" json:"name"`

	// Type is the component type (frontend, backend, service, library, etc.)
	Type string `yaml:"type" json:"type"`

	// Path is the component's file path relative to project root
	Path string `yaml:"path" json:"path"`

	// Description is a brief component description
	Description string `yaml:"description" json:"description"`

	// Dependencies lists other components this depends on
	Dependencies []string `yaml:"dependencies,omitempty" json:"dependencies,omitempty"`

	// Tags are searchable tags for the component
	Tags []string `yaml:"tags,omitempty" json:"tags,omitempty"`

	// TechStack lists technologies used by this component
	TechStack []string `yaml:"tech_stack,omitempty" json:"tech_stack,omitempty"`

	// APIContracts defines API contracts this component implements
	APIContracts []APIContract `yaml:"api_contracts,omitempty" json:"api_contracts,omitempty"`

	// LastModified is when the component was last changed
	LastModified string `yaml:"last_modified,omitempty" json:"last_modified,omitempty"`
}

// APIContract represents an API contract.
type APIContract struct {
	// Name is the contract name
	Name string `yaml:"name" json:"name"`

	// Type is the contract type (rest, graphql, grpc, etc.)
	Type string `yaml:"type" json:"type"`

	// File is the contract definition file
	File string `yaml:"file" json:"file"`
}

// ComponentRegistry manages registered project components.
type ComponentRegistry struct {
	// components stores registered components by name
	components map[string]Component

	// byType indexes components by type
	byType map[string][]string

	// byTag indexes components by tag
	byTag map[string][]string

	// mu protects concurrent access
	mu sync.RWMutex
}

// NewComponentRegistry creates a new component registry.
func NewComponentRegistry() *ComponentRegistry {
	return &ComponentRegistry{
		components: make(map[string]Component),
		byType:     make(map[string][]string),
		byTag:      make(map[string][]string),
	}
}

// Register adds a component to the registry.
func (r *ComponentRegistry) Register(comp Component) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Validate component
	if comp.Name == "" {
		return fmt.Errorf("component name is required")
	}

	// Check if already exists
	if _, exists := r.components[comp.Name]; exists {
		return fmt.Errorf("component %s already registered", comp.Name)
	}

	// Store component
	r.components[comp.Name] = comp

	// Index by type
	if comp.Type != "" {
		r.byType[comp.Type] = append(r.byType[comp.Type], comp.Name)
	}

	// Index by tags
	for _, tag := range comp.Tags {
		r.byTag[tag] = append(r.byTag[tag], comp.Name)
	}

	return nil
}

// Unregister removes a component from the registry.
func (r *ComponentRegistry) Unregister(name string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	comp, exists := r.components[name]
	if !exists {
		return fmt.Errorf("component %s not found", name)
	}

	// Remove from type index
	if comp.Type != "" {
		r.removeFromSlice(r.byType[comp.Type], name)
	}

	// Remove from tag index
	for _, tag := range comp.Tags {
		r.removeFromSlice(r.byTag[tag], name)
	}

	// Remove component
	delete(r.components, name)

	return nil
}

// Get retrieves a component by name.
func (r *ComponentRegistry) Get(name string) (Component, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	comp, exists := r.components[name]
	if !exists {
		return Component{}, fmt.Errorf("component %s not found", name)
	}

	return comp, nil
}

// List returns all registered components.
func (r *ComponentRegistry) List() []Component {
	r.mu.RLock()
	defer r.mu.RUnlock()

	components := make([]Component, 0, len(r.components))
	for _, comp := range r.components {
		components = append(components, comp)
	}

	// Sort by name
	sort.Slice(components, func(i, j int) bool {
		return components[i].Name < components[j].Name
	})

	return components
}

// ListByType returns components of a specific type.
func (r *ComponentRegistry) ListByType(compType string) []Component {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names, exists := r.byType[compType]
	if !exists {
		return []Component{}
	}

	components := make([]Component, 0, len(names))
	for _, name := range names {
		if comp, ok := r.components[name]; ok {
			components = append(components, comp)
		}
	}

	return components
}

// ListByTag returns components with a specific tag.
func (r *ComponentRegistry) ListByTag(tag string) []Component {
	r.mu.RLock()
	defer r.mu.RUnlock()

	names, exists := r.byTag[tag]
	if !exists {
		return []Component{}
	}

	components := make([]Component, 0, len(names))
	for _, name := range names {
		if comp, ok := r.components[name]; ok {
			components = append(components, comp)
		}
	}

	return components
}

// FindDependencies returns all dependencies of a component (recursive).
func (r *ComponentRegistry) FindDependencies(name string) ([]Component, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.findDependenciesRecursive(name, make(map[string]bool))
}

// findDependenciesRecursive recursively finds dependencies.
func (r *ComponentRegistry) findDependenciesRecursive(name string, visited map[string]bool) ([]Component, error) {
	if visited[name] {
		return nil, nil // Avoid circular dependency
	}
	visited[name] = true

	comp, exists := r.components[name]
	if !exists {
		return nil, fmt.Errorf("component %s not found", name)
	}

	var deps []Component
	for _, depName := range comp.Dependencies {
		subDeps, err := r.findDependenciesRecursive(depName, visited)
		if err != nil {
			return nil, err
		}
		deps = append(deps, subDeps...)

		if dep, ok := r.components[depName]; ok {
			deps = append(deps, dep)
		}
	}

	return deps, nil
}

// Validate validates the registry state.
func (r *ComponentRegistry) Validate() []error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var errs []error

	// Check for missing dependencies
	for name, comp := range r.components {
		for _, dep := range comp.Dependencies {
			if _, exists := r.components[dep]; !exists {
				errs = append(errs, fmt.Errorf("component %s has missing dependency: %s", name, dep))
			}
		}
	}

	// Check for circular dependencies
	for name := range r.components {
		if hasCycle := r.checkCircularDependency(name, make(map[string]bool)); hasCycle {
			errs = append(errs, fmt.Errorf("circular dependency detected involving: %s", name))
		}
	}

	return errs
}

// checkCircularDependency checks for circular dependencies.
func (r *ComponentRegistry) checkCircularDependency(name string, visited map[string]bool) bool {
	if visited[name] {
		return true
	}
	visited[name] = true

	comp, exists := r.components[name]
	if !exists {
		return false
	}

	for _, dep := range comp.Dependencies {
		if r.checkCircularDependency(dep, visited) {
			return true
		}
	}

	delete(visited, name)
	return false
}

// removeFromSlice removes a string from a slice.
func (r *ComponentRegistry) removeFromSlice(slice []string, item string) []string {
	for i, s := range slice {
		if s == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// ToMap returns the registry as a map for serialization.
func (r *ComponentRegistry) ToMap() map[string]interface{} {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make(map[string]interface{})
	components := make([]Component, 0, len(r.components))
	for _, comp := range r.components {
		components = append(components, comp)
	}
	result["components"] = components
	result["types"] = r.byType
	result["tags"] = r.byTag
	return result
}

// LoadFromMap loads the registry from a map.
func (r *ComponentRegistry) LoadFromMap(data map[string]interface{}) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	components, ok := data["components"].([]Component)
	if !ok {
		return fmt.Errorf("invalid components data")
	}

	for _, comp := range components {
		r.components[comp.Name] = comp

		if comp.Type != "" {
			r.byType[comp.Type] = append(r.byType[comp.Type], comp.Name)
		}

		for _, tag := range comp.Tags {
			r.byTag[tag] = append(r.byTag[tag], comp.Name)
		}
	}

	return nil
}
