# Architecture Guidelines

## Project Architecture Overview

This document defines the architectural principles for {{.ProjectName}}. All code changes must respect these principles.

## Layer Architecture

```
┌─────────────────────────────────────────────────────────────┐
│                    PRESENTATION LAYER                        │
│     UI Components, Pages, API Endpoints                      │
│     (Depends on Application Layer)                           │
├─────────────────────────────────────────────────────────────┤
│                   APPLICATION LAYER                          │
│     Use Cases, Application Services, Orchestration           │
│     (Depends on Domain Layer)                                │
├─────────────────────────────────────────────────────────────┤
│                      DOMAIN LAYER                            │
│     Business Logic, Domain Entities, Domain Services         │
│     (No external dependencies)                               │
├─────────────────────────────────────────────────────────────┤
│                   INFRASTRUCTURE LAYER                       │
│     Database, External APIs, File System, Third-party        │
│     (Implements Domain Interfaces)                           │
└─────────────────────────────────────────────────────────────┘
```

## Dependency Rules

1. **Dependencies flow inward**: Outer layers can depend on inner layers, never the reverse
2. **Domain independence**: The domain layer has NO dependencies on frameworks, databases, or external services
3. **Interface segregation**: Use interfaces to decouple layers
4. **Dependency injection**: Inject dependencies, don't instantiate them directly

## Folder Structure

```
project-root/
├── frontend/                    # Frontend application
│   ├── src/
│   │   ├── app/                 # Pages/routes (Next.js App Router)
│   │   ├── components/          # Reusable UI components
│   │   │   ├── ui/              # Base UI components (buttons, inputs)
│   │   │   ├── forms/           # Form-related components
│   │   │   ├── layouts/         # Layout components
│   │   │   └── features/        # Feature-specific components
│   │   ├── hooks/               # Custom React hooks
│   │   ├── lib/                 # Utilities and helpers
│   │   ├── types/               # TypeScript type definitions
│   │   ├── styles/              # Global styles
│   │   └── utils/               # Utility functions
│   ├── public/                  # Static assets
│   └── tests/                   # Frontend tests
│
├── backend/                     # Backend application
│   ├── api/                     # API layer
│   │   ├── routes/              # Route handlers
│   │   ├── middleware/          # Request/response middleware
│   │   └── schemas/             # Request/response schemas
│   ├── services/                # Business logic
│   ├── models/                  # Data models
│   ├── repositories/            # Data access
│   ├── utils/                   # Utilities
│   └── tests/                   # Backend tests
│
├── infrastructure/              # Infrastructure configuration
│   ├── docker/                  # Docker configurations
│   ├── kubernetes/              # K8s manifests (if applicable)
│   └── ci/                      # CI/CD configurations
│
├── ai_prompt_guidelines/        # AI governance (this directory)
│
├── docs/                        # Project documentation
│
└── scripts/                     # Utility scripts
```

## Module Boundaries

### Frontend Modules
- **Components** should be self-contained with their own styles
- **Hooks** should handle a single concern
- **Utils** should be pure functions with no side effects

### Backend Modules
- **Routes** should only handle HTTP concerns, delegate to services
- **Services** should contain business logic only
- **Repositories** should handle data access only
- **Models** should represent data structure, not behavior

## Naming Conventions

### Files
- React components: `PascalCase.tsx` (e.g., `UserProfile.tsx`)
- Hooks: `camelCase.ts` with `use` prefix (e.g., `useAuth.ts`)
- Utilities: `camelCase.ts` (e.g., `formatDate.ts`)
- Types: `camelCase.ts` or `types.ts` for collections
- API routes: `kebab-case.ts` (e.g., `user-routes.ts`)

### Code
- Components: `PascalCase`
- Functions: `camelCase`
- Constants: `SCREAMING_SNAKE_CASE`
- Types/Interfaces: `PascalCase`
- Private members: prefix with `_`

## Code Organization Principles

1. **Colocation**: Keep related files together
2. **Single Responsibility**: Each file does one thing
3. **Small Files**: Aim for files under 200 lines
4. **Clear Exports**: Use index files for clean imports
5. **No Circular Dependencies**: Structure to prevent cycles

## Refactoring Guidelines

Before refactoring:
1. Ensure tests exist for the code being refactored
2. Document the reason for refactoring
3. Make small, incremental changes
4. Run tests after each change
5. Update documentation if interfaces change
