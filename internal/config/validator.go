package config

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidationError represents a single validation error.
type ValidationError struct {
	// Field is the field path that failed validation
	Field string `json:"field"`

	// Message describes the validation error
	Message string `json:"message"`

	// Value is the invalid value (optional)
	Value interface{} `json:"value,omitempty"`

	// Severity indicates the error severity (error, warning)
	Severity string `json:"severity"`
}

// Error implements the error interface.
func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// ValidationErrors is a collection of validation errors.
type ValidationErrors []ValidationError

// Error implements the error interface.
func (e ValidationErrors) Error() string {
	if len(e) == 0 {
		return ""
	}

	var msgs []string
	for _, err := range e {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// HasErrors returns true if there are any error-level validation errors.
func (e ValidationErrors) HasErrors() bool {
	for _, err := range e {
		if err.Severity == "error" {
			return true
		}
	}
	return false
}

// HasWarnings returns true if there are any warning-level validation errors.
func (e ValidationErrors) HasWarnings() bool {
	for _, err := range e {
		if err.Severity == "warning" {
			return true
		}
	}
	return false
}

// Validator validates configuration values.
type Validator struct {
	// Strict enables strict validation (warnings become errors)
	Strict bool
}

// NewValidator creates a new configuration validator.
func NewValidator() *Validator {
	return &Validator{
		Strict: false,
	}
}

// Validate validates the entire project configuration.
func (v *Validator) Validate(config *ProjectConfig) ValidationErrors {
	var errors ValidationErrors

	// Validate metadata
	errors = append(errors, v.validateMetadata(&config.Metadata)...)

	// Validate frontend
	if config.Frontend.Enabled {
		errors = append(errors, v.validateFrontend(&config.Frontend)...)
	}

	// Validate backend
	if config.Backend.Enabled {
		errors = append(errors, v.validateBackend(&config.Backend)...)
	}

	// Validate infrastructure
	errors = append(errors, v.validateInfrastructure(&config.Infrastructure)...)

	// Validate governance
	errors = append(errors, v.validateGovernance(&config.Governance)...)

	// Validate cross-field dependencies
	errors = append(errors, v.validateDependencies(config)...)

	return errors
}

// validateMetadata validates project metadata.
func (v *Validator) validateMetadata(m *ProjectMetadata) ValidationErrors {
	var errors ValidationErrors

	// Project name is required
	if m.Name == "" {
		errors = append(errors, ValidationError{
			Field:    "metadata.name",
			Message:  "project name is required",
			Severity: "error",
		})
	} else if !isValidProjectName(m.Name) {
		errors = append(errors, ValidationError{
			Field:    "metadata.name",
			Message:  "project name must contain only lowercase letters, numbers, and hyphens",
			Value:    m.Name,
			Severity: "error",
		})
	}

	// Version should be valid semver if set
	if m.Version != "" && !isValidSemver(m.Version) {
		errors = append(errors, ValidationError{
			Field:    "metadata.version",
			Message:  "version should follow semantic versioning (e.g., 1.0.0)",
			Value:    m.Version,
			Severity: "warning",
		})
	}

	return errors
}

// validateFrontend validates frontend configuration.
func (v *Validator) validateFrontend(f *FrontendConfig) ValidationErrors {
	var errors ValidationErrors

	// Framework is required
	if f.Framework == "" {
		errors = append(errors, ValidationError{
			Field:    "frontend.framework",
			Message:  "frontend framework is required when frontend is enabled",
			Severity: "error",
		})
	} else if !isValidFrontendFramework(f.Framework) {
		errors = append(errors, ValidationError{
			Field:    "frontend.framework",
			Message:  fmt.Sprintf("unsupported frontend framework: %s (supported: react, vue, svelte, angular, nextjs, nuxt, sveltekit)", f.Framework),
			Value:    f.Framework,
			Severity: "error",
		})
	}

	// Styling approach validation
	if f.Styling != "" && !isValidStyling(f.Styling) {
		errors = append(errors, ValidationError{
			Field:    "frontend.styling",
			Message:  fmt.Sprintf("unsupported styling approach: %s (supported: tailwind, css-modules, styled-components, scss)", f.Styling),
			Value:    f.Styling,
			Severity: "error",
		})
	}

	// Package manager validation
	if f.PackageManager != "" && !isValidPackageManager(f.PackageManager) {
		errors = append(errors, ValidationError{
			Field:    "frontend.package_manager",
			Message:  fmt.Sprintf("unsupported package manager: %s (supported: npm, yarn, pnpm, bun)", f.PackageManager),
			Value:    f.PackageManager,
			Severity: "error",
		})
	}

	// Build tool validation
	if f.BuildTool != "" && !isValidBuildTool(f.BuildTool) {
		errors = append(errors, ValidationError{
			Field:    "frontend.build_tool",
			Message:  fmt.Sprintf("unsupported build tool: %s (supported: vite, webpack, esbuild, rollup, turbopack)", f.BuildTool),
			Value:    f.BuildTool,
			Severity: "error",
		})
	}

	// Directory validation
	if f.Directory == "" {
		errors = append(errors, ValidationError{
			Field:    "frontend.directory",
			Message:  "frontend directory is required",
			Severity: "error",
		})
	}

	// Feature compatibility checks
	if f.Features.SSR && !supportsSSR(f.Framework) {
		errors = append(errors, ValidationError{
			Field:    "frontend.features.ssr",
			Message:  fmt.Sprintf("SSR is not supported by framework: %s", f.Framework),
			Value:    f.Features.SSR,
			Severity: "warning",
		})
	}

	return errors
}

// validateBackend validates backend configuration.
func (v *Validator) validateBackend(b *BackendConfig) ValidationErrors {
	var errors ValidationErrors

	// Framework is required
	if b.Framework == "" {
		errors = append(errors, ValidationError{
			Field:    "backend.framework",
			Message:  "backend framework is required when backend is enabled",
			Severity: "error",
		})
	} else if !isValidBackendFramework(b.Framework) {
		errors = append(errors, ValidationError{
			Field:    "backend.framework",
			Message:  fmt.Sprintf("unsupported backend framework: %s (supported: fastapi, express, nestjs, go-gin, go-fiber, rust-axum, django, rails)", b.Framework),
			Value:    b.Framework,
			Severity: "error",
		})
	}

	// Language validation
	if b.Language == "" {
		errors = append(errors, ValidationError{
			Field:    "backend.language",
			Message:  "backend language is required when backend is enabled",
			Severity: "error",
		})
	}

	// Database validation
	errors = append(errors, v.validateDatabase(&b.Database)...)

	// Auth validation
	errors = append(errors, v.validateAuth(&b.Auth)...)

	// API validation
	errors = append(errors, v.validateAPI(&b.API)...)

	// Directory validation
	if b.Directory == "" {
		errors = append(errors, ValidationError{
			Field:    "backend.directory",
			Message:  "backend directory is required",
			Severity: "error",
		})
	}

	return errors
}

// validateDatabase validates database configuration.
func (v *Validator) validateDatabase(d *DatabaseConfig) ValidationErrors {
	var errors ValidationErrors

	// Primary database validation
	if d.Primary != "" && !isValidDatabase(d.Primary) {
		errors = append(errors, ValidationError{
			Field:    "backend.database.primary",
			Message:  fmt.Sprintf("unsupported database: %s (supported: postgresql, mysql, sqlite, mongodb, mariadb, cockroachdb)", d.Primary),
			Value:    d.Primary,
			Severity: "error",
		})
	}

	// ORM validation based on database
	if d.Primary != "" && d.ORM != "" && !isValidORMForDatabase(d.ORM, d.Primary) {
		errors = append(errors, ValidationError{
			Field:    "backend.database.orm",
			Message:  fmt.Sprintf("ORM %s may not be compatible with database %s", d.ORM, d.Primary),
			Value:    d.ORM,
			Severity: "warning",
		})
	}

	return errors
}

// validateAuth validates authentication configuration.
func (v *Validator) validateAuth(a *AuthConfig) ValidationErrors {
	var errors ValidationErrors

	// Provider validation
	if a.Provider != "" && !isValidAuthProvider(a.Provider) {
		errors = append(errors, ValidationError{
			Field:    "backend.auth.provider",
			Message:  fmt.Sprintf("unsupported auth provider: %s (supported: jwt, oauth, clerk, auth0, firebase, nextauth, passport)", a.Provider),
			Value:    a.Provider,
			Severity: "error",
		})
	}

	// Session duration validation
	if a.SessionDuration < 0 {
		errors = append(errors, ValidationError{
			Field:    "backend.auth.session_duration",
			Message:  "session duration must be non-negative",
			Value:    a.SessionDuration,
			Severity: "error",
		})
	}

	return errors
}

// validateAPI validates API configuration.
func (v *Validator) validateAPI(a *APIConfig) ValidationErrors {
	var errors ValidationErrors

	// Style validation
	if a.Style != "" && !isValidAPIStyle(a.Style) {
		errors = append(errors, ValidationError{
			Field:    "backend.api.style",
			Message:  fmt.Sprintf("unsupported API style: %s (supported: rest, graphql, grpc, trpc)", a.Style),
			Value:    a.Style,
			Severity: "error",
		})
	}

	// Versioning validation
	if a.Versioning != "" && !isValidAPIVersioning(a.Versioning) {
		errors = append(errors, ValidationError{
			Field:    "backend.api.versioning",
			Message:  fmt.Sprintf("unsupported API versioning: %s (supported: url, header, query, none)", a.Versioning),
			Value:    a.Versioning,
			Severity: "error",
		})
	}

	return errors
}

// validateInfrastructure validates infrastructure configuration.
func (v *Validator) validateInfrastructure(i *InfrastructureConfig) ValidationErrors {
	var errors ValidationErrors

	// CI platform validation
	if i.CI != "" && !isValidCI(i.CI) {
		errors = append(errors, ValidationError{
			Field:    "infrastructure.ci",
			Message:  fmt.Sprintf("unsupported CI platform: %s (supported: github-actions, gitlab-ci, circleci, jenkins, azure-pipelines)", i.CI),
			Value:    i.CI,
			Severity: "error",
		})
	}

	// Hosting platform validation
	if i.Hosting != "" && !isValidHosting(i.Hosting) {
		errors = append(errors, ValidationError{
			Field:    "infrastructure.hosting",
			Message:  fmt.Sprintf("unsupported hosting platform: %s (supported: vercel, netlify, aws, gcp, azure, digitalocean, railway, render, fly)", i.Hosting),
			Value:    i.Hosting,
			Severity: "error",
		})
	}

	// Kubernetes requires Docker
	if i.Kubernetes && !i.Docker {
		errors = append(errors, ValidationError{
			Field:    "infrastructure.kubernetes",
			Message:  "Kubernetes requires Docker to be enabled",
			Severity: "warning",
		})
	}

	return errors
}

// validateGovernance validates governance configuration.
func (v *Validator) validateGovernance(g *GovernanceConfig) ValidationErrors {
	var errors ValidationErrors

	// Context level validation
	if g.ContextLevel != "" && !isValidContextLevel(g.ContextLevel) {
		errors = append(errors, ValidationError{
			Field:    "governance.context_level",
			Message:  fmt.Sprintf("invalid context level: %s (supported: minimal, standard, comprehensive)", g.ContextLevel),
			Value:    g.ContextLevel,
			Severity: "error",
		})
	}

	return errors
}

// validateDependencies validates cross-field dependencies.
func (v *Validator) validateDependencies(config *ProjectConfig) ValidationErrors {
	var errors ValidationErrors

	// Check if at least one of frontend or backend is enabled
	if !config.Frontend.Enabled && !config.Backend.Enabled {
		errors = append(errors, ValidationError{
			Field:    "config",
			Message:  "at least one of frontend or backend must be enabled",
			Severity: "error",
		})
	}

	// SSR requires a backend
	if config.Frontend.Enabled && config.Frontend.Features.SSR && !config.Backend.Enabled {
		errors = append(errors, ValidationError{
			Field:    "frontend.features.ssr",
			Message:  "SSR requires a backend to be enabled",
			Severity: "warning",
		})
	}

	// Docker Compose is useful with backend
	if config.Backend.Enabled && !config.Infrastructure.DockerCompose {
		errors = append(errors, ValidationError{
			Field:    "infrastructure.docker_compose",
			Message:  "Docker Compose is recommended for backend development",
			Severity: "warning",
		})
	}

	// Monitoring is recommended for production
	if !config.Infrastructure.Monitoring.Enabled {
		errors = append(errors, ValidationError{
			Field:    "infrastructure.monitoring.enabled",
			Message:  "monitoring is recommended for production applications",
			Severity: "warning",
		})
	}

	return errors
}

// Helper validation functions

var projectNameRegex = regexp.MustCompile(`^[a-z][a-z0-9-]*$`)

func isValidProjectName(name string) bool {
	return projectNameRegex.MatchString(name) && len(name) <= 100
}

var semverRegex = regexp.MustCompile(`^v?(\d+)\.(\d+)\.(\d+)(-([a-zA-Z0-9.-]+))?(\+([a-zA-Z0-9.-]+))?$`)

func isValidSemver(version string) bool {
	return semverRegex.MatchString(version)
}

func isValidFrontendFramework(framework string) bool {
	validFrameworks := []string{
		"react", "vue", "svelte", "angular",
		"nextjs", "nuxt", "sveltekit", "remix",
		"astro", "solid",
	}
	return contains(validFrameworks, framework)
}

func isValidBackendFramework(framework string) bool {
	validFrameworks := []string{
		"fastapi", "express", "nestjs", "django",
		"go-gin", "go-fiber", "go-echo",
		"rust-axum", "rust-actix", "rust-rocket",
		"rails", "phoenix", "spring",
	}
	return contains(validFrameworks, framework)
}

func isValidStyling(styling string) bool {
	validStyling := []string{
		"tailwind", "css-modules", "styled-components",
		"scss", "sass", "less", "emotion", "stitches",
	}
	return contains(validStyling, styling)
}

func isValidPackageManager(pm string) bool {
	validPM := []string{"npm", "yarn", "pnpm", "bun"}
	return contains(validPM, pm)
}

func isValidBuildTool(tool string) bool {
	validTools := []string{
		"vite", "webpack", "esbuild", "rollup",
		"turbo", "turboPack", "parcel", "swc",
	}
	return contains(validTools, tool)
}

func supportsSSR(framework string) bool {
	ssrFrameworks := []string{
		"nextjs", "nuxt", "sveltekit", "remix",
		"astro", "angular",
	}
	return contains(ssrFrameworks, framework)
}

func isValidDatabase(db string) bool {
	validDB := []string{
		"postgresql", "mysql", "sqlite", "mongodb",
		"mariadb", "cockroachdb", "planetscale",
	}
	return contains(validDB, db)
}

func isValidORMForDatabase(orm, db string) bool {
	// Define ORM compatibility
	ormDBMap := map[string][]string{
		"prisma":     {"postgresql", "mysql", "sqlite", "mongodb", "cockroachdb"},
		"sqlalchemy": {"postgresql", "mysql", "sqlite", "mariadb"},
		"typeorm":    {"postgresql", "mysql", "sqlite", "mongodb", "mariadb"},
		"drizzle":    {"postgresql", "mysql", "sqlite"},
		"mongoose":   {"mongodb"},
		"gorm":       {"postgresql", "mysql", "sqlite"},
		"sqlboiler":  {"postgresql", "mysql", "sqlite"},
		"ent":        {"postgresql", "mysql", "sqlite"},
	}

	supportedDBs, ok := ormDBMap[orm]
	if !ok {
		return true // Unknown ORM, assume compatible
	}
	return contains(supportedDBs, db)
}

func isValidAuthProvider(provider string) bool {
	validProviders := []string{
		"jwt", "oauth", "oidc",
		"clerk", "auth0", "firebase",
		"nextauth", "passport", "lucia",
		"supabase", "cognito",
	}
	return contains(validProviders, provider)
}

func isValidAPIStyle(style string) bool {
	validStyles := []string{"rest", "graphql", "grpc", "trpc", "tsoa"}
	return contains(validStyles, style)
}

func isValidAPIVersioning(versioning string) bool {
	validVersioning := []string{"url", "header", "query", "none"}
	return contains(validVersioning, versioning)
}

func isValidCI(ci string) bool {
	validCI := []string{
		"github-actions", "gitlab-ci", "circleci",
		"jenkins", "azure-pipelines", "travis",
		"bitbucket-pipelines", "buildkite",
	}
	return contains(validCI, ci)
}

func isValidHosting(hosting string) bool {
	validHosting := []string{
		"vercel", "netlify", "aws", "gcp", "azure",
		"digitalocean", "railway", "render", "fly",
		"heroku", "cloudflare", "self-hosted",
	}
	return contains(validHosting, hosting)
}

func isValidContextLevel(level string) bool {
	validLevels := []string{"minimal", "standard", "comprehensive"}
	return contains(validLevels, level)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Validate validates a configuration and returns any errors.
func Validate(config *ProjectConfig) ValidationErrors {
	return NewValidator().Validate(config)
}

// ValidateStrict validates a configuration in strict mode.
func ValidateStrict(config *ProjectConfig) ValidationErrors {
	v := NewValidator()
	v.Strict = true
	return v.Validate(config)
}

// IsValid returns true if the configuration has no errors.
func IsValid(config *ProjectConfig) bool {
	errors := Validate(config)
	return !errors.HasErrors()
}
