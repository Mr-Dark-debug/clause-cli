# 📚 CLAUSE CLI - Templates, Governance & Documentation Generation Prompt

---

## 📋 OVERVIEW

You are tasked with creating **ALL content files** for the Clause CLI project. This includes AI governance templates, project templates, documentation, examples, and all supporting files needed for a production-ready open source project.

**Project:** Clause CLI (AI-Native Project Scaffolding Tool)
**Command:** `clause`
**Repository:** https://github.com/Mr-Dark-debug/clause-cli

---

## 🎯 WHAT YOU NEED TO CREATE

### Part 1: AI Governance Templates
These files define how AI assistants should behave when working on Clause-generated projects.

### Part 2: Project Templates
Actual scaffolding templates for different project types (frontend, backend, full-stack).

### Part 3: Documentation
Complete documentation including README, guides, examples, and API references.

### Part 4: Configuration Files
Example configurations, presets, and rule definitions.

---

# PART 1: AI GOVERNANCE TEMPLATES

## 1.1 templates/governance/ai_prompt_guidelines/system_prompt.md

Create the main system prompt template that will be included in every Clause-generated project:

```markdown
# System Prompt for AI Coding Assistants

## Role Definition

You are an expert software engineer working on the {{.ProjectName}} project. You have complete access to the codebase and are responsible for writing, modifying, and maintaining code according to the guidelines defined in this directory.

## Core Behavioral Rules

### 1. Always Read First
Before making ANY code changes, you MUST:
1. Read this entire `ai_prompt_guidelines/` directory
2. Understand the project architecture from `architecture.md`
3. Check the technology whitelist in `technologies.md`
4. Review any existing patterns in `patterns/` directory
5. Check the component registry in `registry.json`

### 2. Follow the Stack
- This project uses: {{.FrontendFramework}} for frontend, {{.BackendFramework}} for backend
- Do NOT introduce new technologies without explicit user permission
- Do NOT suggest alternative frameworks or libraries that are not in the approved list
- If you need a library not listed, ask the user first

### 3. Maintain Architecture
- Respect the folder structure defined in `architecture.md`
- Keep business logic separate from presentation
- Follow the dependency direction: outer layers depend on inner layers
- Never create circular dependencies

### 4. Document Your Work
- Every function must have a doc comment explaining purpose, parameters, and return values
- Complex logic must include inline comments explaining the "why"
- New files must include a file-level comment explaining their purpose
- Update the component registry when creating new reusable components

### 5. Write Clean Code
- Use meaningful variable and function names
- Keep functions small and focused (max 50 lines preferred)
- Follow the project's naming conventions
- Remove dead code and unused imports

## Working with Brainstorm.md

The `brainstorm.md` file is your working memory. Use it as follows:

### When to Write to Brainstorm.md
- When you encounter ambiguity or uncertainty
- When you need to explore multiple approaches
- When you're blocked waiting for information
- When you need to track decisions and their rationale

### How to Use Brainstorm.md
1. Write your current problem under "Open Questions"
2. Continue working on other tasks (don't wait for user input)
3. Periodically check back to see if questions are resolved
4. Move resolved items to "Decision Log" with the resolution
5. Update the "Current Focus" section when switching tasks

### Brainstorm.md Structure
```markdown
# Brainstorm

## Current Focus
[What you're working on right now]

## Open Questions
- [Question 1?]
- [Question 2?]

## Reasoning Space
[Your working notes and analysis]

## Decision Log
- [Date] Decision: [What was decided] - Reason: [Why]

## Blockers
- [What's preventing progress]
```

## Code Generation Rules

### Frontend Code ({{.FrontendFramework}})
{{if eq .FrontendFramework "nextjs"}}
- Use App Router (app/ directory)
- Components go in `components/` with subdirectories by category
- Use TypeScript for all new files
- Follow the component structure in `patterns/react-component.md`
- Style with {{.StylingApproach}}
{{end}}

{{if eq .FrontendFramework "react"}}
- Use functional components with hooks
- Keep components in `src/components/`
- Use TypeScript for all new files
- Follow React best practices (avoid prop drilling, use context wisely)
{{end}}

### Backend Code ({{.BackendFramework}})
{{if eq .BackendFramework "fastapi"}}
- Use Pydantic models for all schemas
- Keep routers in `api/routes/`
- Business logic in `services/`
- Database models in `models/`
- Always include type hints
{{end}}

{{if eq .BackendFramework "express"}}
- Use Express Router for route organization
- Middleware in `middleware/`
- Controllers in `controllers/`
- Use TypeScript with proper typing
{{end}}

## Error Handling

When you encounter errors:
1. Read the error message carefully
2. Check if this is a known pattern in the codebase
3. Search for similar implementations
4. If stuck, write to brainstorm.md and continue with other tasks
5. Only ask the user for help if truly blocked

## Prohibited Actions

NEVER:
- Delete files without explicit permission
- Rewrite large portions of code without discussion
- Introduce breaking changes without warning
- Ignore existing patterns and conventions
- Skip documentation
- Leave TODO comments without creating tracking issues

## Communication Style

- Be concise but thorough
- Explain your reasoning for significant decisions
- Warn about potential side effects of changes
- Suggest improvements but don't implement without approval
- Use the brainstorm.md for extended reasoning, not chat messages
```

---

## 1.2 templates/governance/ai_prompt_guidelines/architecture.md

Create the architecture principles template:

```markdown
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
```

---

## 1.3 templates/governance/ai_prompt_guidelines/technologies.md

Create the technology constraints template:

```markdown
# Technology Stack & Constraints

## Approved Technologies

This document defines what technologies are approved for use in {{.ProjectName}}. AI assistants must NOT introduce technologies not listed here without explicit user approval.

## Frontend Stack

### Core Framework
{{if eq .FrontendFramework "nextjs"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| Next.js | 14.x | React framework with SSR/SSG |
| React | 18.x | UI library |
| TypeScript | 5.x | Type-safe JavaScript |
{{end}}

{{if eq .FrontendFramework "react"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| React | 18.x | UI library |
| Vite | 5.x | Build tool |
| TypeScript | 5.x | Type-safe JavaScript |
{{end}}

{{if eq .FrontendFramework "vue"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| Vue.js | 3.x | UI framework |
| Vite | 5.x | Build tool |
| TypeScript | 5.x | Type-safe JavaScript |
{{end}}

### Styling
{{if eq .StylingApproach "tailwind"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| Tailwind CSS | 3.x | Utility-first CSS |
| PostCSS | 8.x | CSS processing |
{{end}}

{{if eq .StylingApproach "css-modules"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| CSS Modules | - | Scoped CSS |
| PostCSS | 8.x | CSS processing |
{{end}}

### State Management
{{if eq .StateManager "redux"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| Redux Toolkit | 2.x | State management |
| React-Redux | 9.x | React bindings |
{{end}}

{{if eq .StateManager "zustand"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| Zustand | 4.x | State management |
{{end}}

### UI Components
{{if eq .UILibrary "shadcn"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| shadcn/ui | latest | Component library |
| Radix UI | latest | Headless components |
| Lucide React | latest | Icons |
{{end}}

## Backend Stack

### Core Framework
{{if eq .BackendFramework "fastapi"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| FastAPI | 0.109+ | Python web framework |
| Pydantic | 2.x | Data validation |
| Uvicorn | 0.27+ | ASGI server |
| Python | 3.11+ | Runtime |
{{end}}

{{if eq .BackendFramework "express"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| Express | 4.x | Node.js web framework |
| TypeScript | 5.x | Type-safe JavaScript |
| Node.js | 20.x LTS | Runtime |
{{end}}

{{if eq .BackendFramework "django"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| Django | 5.x | Python web framework |
| Django REST Framework | 3.x | API toolkit |
| Python | 3.11+ | Runtime |
{{end}}

### Database
{{if eq .Database "postgresql"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| PostgreSQL | 15+ | Primary database |
{{if eq .BackendFramework "fastapi"}}
| SQLAlchemy | 2.x | ORM |
| Alembic | 1.x | Migrations |
{{end}}
{{if eq .BackendFramework "django"}}
| psycopg2 | 2.x | PostgreSQL adapter |
{{end}}
{{end}}

{{if eq .Database "mongodb"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| MongoDB | 6+ | Primary database |
| Mongoose | 8.x | ODM (Node.js) |
{{end}}

### Authentication
{{if eq .Auth "jwt"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| JWT | - | Token-based auth |
{{if eq .BackendFramework "fastapi"}}
| python-jose | 3.x | JWT handling |
| passlib | 1.x | Password hashing |
{{end}}
{{end}}

## Infrastructure

### Containerization
| Technology | Version | Purpose |
|------------|---------|---------|
| Docker | 24+ | Containerization |
| Docker Compose | 2.x | Multi-container orchestration |

### CI/CD
{{if eq .CICD "github"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| GitHub Actions | - | CI/CD platform |
{{end}}

## Testing

### Frontend Testing
| Technology | Version | Purpose |
|------------|---------|---------|
| Vitest | 1.x | Test runner |
| React Testing Library | 14.x | Component testing |
| Playwright | 1.x | E2E testing |

### Backend Testing
{{if eq .BackendFramework "fastapi"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| pytest | 8.x | Test framework |
| pytest-asyncio | 0.23+ | Async testing |
| httpx | 0.26+ | API testing |
{{end}}

{{if eq .BackendFramework "express"}}
| Technology | Version | Purpose |
|------------|---------|---------|
| Jest | 29.x | Test framework |
| Supertest | 6.x | API testing |
{{end}}

## Code Quality

| Technology | Version | Purpose |
|------------|---------|---------|
| ESLint | 8.x | Linting |
| Prettier | 3.x | Code formatting |
{{if eq .BackendFramework "fastapi"}}
| Ruff | 0.1+ | Python linting |
| Black | 24.x | Python formatting |
{{end}}

## Forbidden Technologies

The following are NOT allowed in this project:

- **jQuery** - Use native DOM or React instead
- **Moment.js** - Use date-fns or native Date
- **Lodash** - Use native JavaScript methods
- **Enzyme** - Use React Testing Library instead
- **Class components** - Use functional components with hooks
- **var** - Use const or let
- **Any** type in TypeScript - Define proper types

## Version Pinning Rules

1. Pin major versions in package.json / requirements.txt
2. Use lock files (package-lock.json, poetry.lock)
3. Document why specific versions are pinned if not latest
4. Regularly update dependencies for security patches
```

---

## 1.4 templates/governance/ai_prompt_guidelines/brainstorm.md

Create the brainstorm template:

```markdown
# 🧠 Brainstorm

This file serves as external working memory for AI assistants working on this project. Use it to reason through complex problems, track decisions, and work autonomously.

---

## 📌 Current Focus

*What are you working on right now?*

**Task:**
**Status:**
**Started:**

---

## ❓ Open Questions

*Questions that need answers. Write them here and continue working - check back later.*

1. 
2. 
3. 

---

## 💭 Reasoning Space

*Use this area for thinking through complex problems, exploring approaches, and working through logic.*

### Active Investigation



### Approaches Considered



### Trade-offs Being Evaluated



---

## 📋 Decision Log

*Record decisions made and their rationale. This creates a history of architectural choices.*

| Date | Decision | Rationale | Alternatives Considered |
|------|----------|-----------|------------------------|
| | | | |

---

## 🚧 Blockers

*What's preventing progress? List here and check periodically if resolved.*

1. **Blocker:** 
   - **Impact:** 
   - **Possible Solutions:** 
   - **Status:** 

---

## 📝 Notes & Observations

*General notes about the codebase, patterns noticed, things to remember.*

- 
- 
- 

---

## 🔄 Tasks to Revisit

*Things to come back to later.*

- [ ] 
- [ ] 
- [ ] 

---

## 📚 Reference Links

*Useful links, documentation, or resources.*

- 
- 
- 

---

## Usage Instructions for AI Assistants

1. **When stuck**: Write your question in "Open Questions" and continue with other work
2. **When deciding**: Document the decision in "Decision Log" with rationale
3. **When blocked**: Add to "Blockers" and check back periodically
4. **When reasoning**: Use "Reasoning Space" to think through complex problems
5. **Keep it clean**: Remove resolved items, keep only active context

**Remember**: This file is your second brain. Use it to maintain context across long sessions and to avoid asking the user for every small decision.
```

---

## 1.5 templates/governance/ai_prompt_guidelines/documentation.md

Create documentation standards:

```markdown
# Documentation Standards

## Overview

All code in {{.ProjectName}} must be documented according to these standards. Documentation is not optional - it's a core part of the development process.

## File-Level Documentation

Every source file must start with a file-level comment:

```typescript
/**
 * @fileoverview Brief description of what this file contains
 * 
 * This file provides [main purpose]. It is responsible for:
 * - [responsibility 1]
 * - [responsibility 2]
 * 
 * @module path/to/module
 * @author Project Team
 * @created YYYY-MM-DD
 */
```

```python
"""
Module: module_name

This module provides [main purpose]. It is responsible for:
- [responsibility 1]
- [responsibility 2]

Example:
    from module import function
    result = function(param)
"""
```

## Function Documentation

### TypeScript/JavaScript

```typescript
/**
 * Brief description of what the function does.
 * 
 * @description More detailed description if needed.
 * Can span multiple lines for complex functions.
 * 
 * @param paramName - Description of the parameter
 * @param anotherParam - Description of another parameter
 * @returns Description of what is returned
 * 
 * @throws {ErrorType} When this error occurs
 * 
 * @example
 * ```ts
 * const result = functionName('value', 123);
 * console.log(result); // Expected output
 * ```
 * 
 * @since 1.0.0
 * @see relatedFunction
 */
function functionName(paramName: string, anotherParam: number): ReturnType {
  // implementation
}
```

### Python

```python
def function_name(param_name: str, another_param: int) -> ReturnType:
    """
    Brief description of what the function does.
    
    More detailed description if needed. Can span multiple
    lines for complex functions.
    
    Args:
        param_name: Description of the parameter.
        another_param: Description of another parameter.
    
    Returns:
        Description of what is returned.
    
    Raises:
        ErrorType: When this error occurs.
    
    Example:
        >>> result = function_name("value", 123)
        >>> print(result)
        expected_output
    
    Note:
        Any additional notes about the function.
    """
    # implementation
```

## Class Documentation

```typescript
/**
 * Brief description of the class.
 * 
 * @description More detailed description of the class purpose
 * and how it should be used.
 * 
 * @example
 * ```ts
 * const instance = new ClassName(config);
 * instance.method();
 * ```
 */
class ClassName {
  /**
   * Description of the property.
   */
  public property: Type;

  /**
   * Creates an instance of ClassName.
   * 
   * @param config - Configuration object
   */
  constructor(config: Config) {
    // implementation
  }
}
```

## Component Documentation (React)

```typescript
/**
 * ComponentName - Brief description of what it renders.
 * 
 * @description More detailed description of the component's
 * purpose, behavior, and any important notes.
 * 
 * @example
 * ```tsx
 * <ComponentName 
 *   prop1="value"
 *   prop2={123}
 * />
 * ```
 */
interface ComponentNameProps {
  /** Description of prop1 */
  prop1: string;
  /** Description of prop2 */
  prop2?: number;
}

export function ComponentName({ prop1, prop2 = 0 }: ComponentNameProps) {
  // implementation
}
```

## API Documentation

### Endpoint Documentation

```typescript
/**
 * @route GET /api/users/:id
 * @description Retrieves a user by their ID
 * 
 * @param {string} id - The user's unique identifier
 * 
 * @returns {User} The user object
 * @throws {404} User not found
 * @throws {500} Server error
 * 
 * @example
 * // Request
 * GET /api/users/abc123
 * 
 * // Response
 * {
 *   "id": "abc123",
 *   "name": "John Doe",
 *   "email": "john@example.com"
 * }
 */
```

## README Sections

Every major module should have a README.md with:

1. **Title and Description** - What is this module?
2. **Installation** - How to install/enable it
3. **Usage** - How to use it with examples
4. **API Reference** - Public API documentation
5. **Configuration** - Available configuration options
6. **Examples** - Common use cases
7. **Contributing** - How to contribute (if applicable)

## Inline Comments

Use inline comments for:

1. **Complex logic explanation**: Why this approach was chosen
2. **Non-obvious behavior**: Things that might surprise readers
3. **TODOs**: With ticket/issue reference
4. **Bug workarounds**: With link to issue

```typescript
// Using binary search because the list is sorted
// This gives O(log n) instead of O(n)
const index = binarySearch(sortedList, target);

// TODO(#123): Remove this workaround when API is fixed
// Currently the API returns null for empty strings
if (response.data === null) {
  response.data = '';
}

// HACK: Safari doesn't support this CSS property
// Using vendor prefix as a fallback
element.style.webkitTransform = 'translateX(0)';
```

## Documentation Review Checklist

Before submitting code, verify:

- [ ] All public functions have JSDoc/docstrings
- [ ] Complex logic has inline comments
- [ ] README files are up to date
- [ ] Examples are working and tested
- [ ] Type definitions include descriptions
- [ ] API endpoints are documented
- [ ] Breaking changes are noted
```

---

## 1.6 templates/governance/ai_prompt_guidelines/registry.json

Create the component registry template:

```json
{
  "$schema": "https://clause.dev/schemas/registry.json",
  "version": "1.0.0",
  "project": "{{.ProjectName}}",
  "lastUpdated": "{{.GeneratedAt}}",
  "components": {
    "frontend": {
      "ui": [],
      "forms": [],
      "layouts": [],
      "features": []
    },
    "backend": {
      "routes": [],
      "services": [],
      "models": [],
      "middleware": []
    },
    "shared": {
      "types": [],
      "utils": [],
      "constants": []
    }
  },
  "dependencies": {},
  "metadata": {
    "generatedBy": "Clause CLI v{{.ClauseVersion}}",
    "template": "{{.TemplateName}}"
  }
}
```

---

## 1.7 templates/governance/ai_prompt_guidelines/context.yaml

Create context configuration:

```yaml
# AI Context Configuration
# This file provides structured context for AI assistants

project:
  name: {{.ProjectName}}
  description: {{.ProjectDescription}}
  version: 0.1.0
  created: {{.GeneratedAt}}

stack:
  frontend:
    framework: {{.FrontendFramework}}
    styling: {{.StylingApproach}}
    state: {{.StateManager}}
  backend:
    framework: {{.BackendFramework}}
    database: {{.Database}}
    auth: {{.AuthMethod}}
  infrastructure:
    containerization: {{.Containerization}}
    cicd: {{.CICD}}

conventions:
  language: {{.PrimaryLanguage}}
  testing: {{.TestingFramework}}
  linting: {{.Linter}}
  formatting: {{.Formatter}}

governance:
  strictness: {{.GovernanceStrictness}}
  documentation: {{.DocumentationLevel}}
  
paths:
  frontend: ./frontend
  backend: ./backend
  shared: ./shared
  tests: ./tests
  docs: ./docs

rules:
  # AI behavior rules
  max_file_lines: 200
  require_docs: true
  require_tests: true
  require_types: true
  
  # Technology restrictions
  allow_new_dependencies: ask  # always | never | ask
  allow_framework_changes: ask
  allow_db_changes: ask
  
  # Documentation requirements
  doc_functions: always        # always | public | complex | never
  doc_classes: always
  doc_components: always
  doc_endpoints: always

quick_context:
  # Summary for AI context window optimization
  summary: |
    {{.ProjectName}} is a {{.ProjectType}} built with {{.FrontendFramework}} 
    frontend and {{.BackendFramework}} backend. It uses {{.Database}} for 
    persistence and follows clean architecture principles.
    
  key_patterns:
    - Layer-based architecture with dependency inversion
    - Repository pattern for data access
    - Service layer for business logic
    - Component-based frontend with hooks
    
  important_files:
    - ai_prompt_guidelines/system_prompt.md    # AI behavior rules
    - ai_prompt_guidelines/architecture.md     # Architecture constraints
    - ai_prompt_guidelines/technologies.md     # Allowed technologies
    - ai_prompt_guidelines/brainstorm.md       # Working memory
    - ai_prompt_guidelines/registry.json       # Component registry
```

---

# PART 2: PROJECT TEMPLATES

## 2.1 templates/frontend/nextjs/ Template Files

Create a complete Next.js project template:

### templates/frontend/nextjs/package.json.template

```json
{
  "name": "{{.ProjectName}}-frontend",
  "version": "0.1.0",
  "private": true,
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "start": "next start",
    "lint": "next lint",
    "test": "vitest",
    "test:coverage": "vitest --coverage",
    "type-check": "tsc --noEmit"
  },
  "dependencies": {
    "next": "^14.1.0",
    "react": "^18.2.0",
    "react-dom": "^18.2.0"{{if eq .UILibrary "shadcn"}},
    "@radix-ui/react-slot": "^1.0.2",
    "class-variance-authority": "^0.7.0",
    "clsx": "^2.1.0",
    "tailwind-merge": "^2.2.0",
    "lucide-react": "^0.312.0"{{end}}
  },
  "devDependencies": {
    "@types/node": "^20.11.0",
    "@types/react": "^18.2.0",
    "@types/react-dom": "^18.2.0",
    "typescript": "^5.3.0",
    "eslint": "^8.56.0",
    "eslint-config-next": "^14.1.0"{{if eq .StylingApproach "tailwind"}},
    "tailwindcss": "^3.4.0",
    "postcss": "^8.4.0",
    "autoprefixer": "^10.4.0"{{end}},
    "vitest": "^1.2.0",
    "@testing-library/react": "^14.1.0",
    "@vitejs/plugin-react": "^4.2.0"
  }
}
```

### templates/frontend/nextjs/README.md.template

```markdown
# {{.ProjectName}} Frontend

A modern web application built with Next.js 14, React 18, and TypeScript.

## 🚀 Quick Start

### Prerequisites

- Node.js 20+ (LTS recommended)
- npm, yarn, or pnpm

### Installation

```bash
# Install dependencies
npm install

# Start development server
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) to see the application.

## 📁 Project Structure

```
frontend/
├── src/
│   ├── app/                    # Next.js App Router pages
│   │   ├── layout.tsx          # Root layout
│   │   ├── page.tsx            # Home page
│   │   └── (routes)/           # Route groups
│   ├── components/
│   │   ├── ui/                 # Base UI components
│   │   ├── forms/              # Form components
│   │   ├── layouts/            # Layout components
│   │   └── features/           # Feature components
│   ├── hooks/                  # Custom React hooks
│   ├── lib/                    # Utilities and helpers
│   ├── types/                  # TypeScript types
│   ├── styles/                 # Global styles
│   └── utils/                  # Utility functions
├── public/                     # Static assets
└── tests/                      # Test files

## 🧪 Testing

```bash
# Run tests
npm test

# Run tests with coverage
npm run test:coverage
```

## 📝 Scripts

| Command | Description |
|---------|-------------|
| `npm run dev` | Start development server |
| `npm run build` | Build for production |
| `npm run start` | Start production server |
| `npm run lint` | Run ESLint |
| `npm run test` | Run tests |
| `npm run type-check` | Run TypeScript type checking |

## 🔧 Configuration

Environment variables are defined in `.env.local`:

```env
NEXT_PUBLIC_API_URL=http://localhost:8000
```

## 📚 Learn More

- [Next.js Documentation](https://nextjs.org/docs)
- [React Documentation](https://react.dev)
- [TypeScript Documentation](https://www.typescriptlang.org/docs)

## 📄 License

MIT
```

### templates/frontend/nextjs/src/app/layout.tsx.template

```tsx
import type { Metadata } from 'next'
import { Inter } from 'next/font/google'
import './globals.css'

const inter = Inter({ subsets: ['latin'] })

export const metadata: Metadata = {
  title: '{{.ProjectName}}',
  description: '{{.ProjectDescription}}',
}

/**
 * Root Layout Component
 * 
 * This is the top-level layout that wraps all pages.
 * It includes global styles, fonts, and any providers.
 * 
 * @param children - Child components to render
 */
export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <body className={inter.className}>{children}</body>
    </html>
  )
}
```

### templates/frontend/nextjs/src/app/page.tsx.template

```tsx
/**
 * Home Page
 * 
 * The main landing page for the application.
 * Modify this to reflect your project's content.
 */

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <div className="text-center">
        <h1 className="text-4xl font-bold mb-4">
          Welcome to {{.ProjectName}}
        </h1>
        <p className="text-lg text-gray-600">
          {{.ProjectDescription}}
        </p>
        <div className="mt-8">
          <a
            href="/docs"
            className="rounded-md bg-orange-500 px-4 py-2 text-white hover:bg-orange-600"
          >
            Get Started
          </a>
        </div>
      </div>
    </main>
  )
}
```

---

## 2.2 templates/backend/fastapi/ Template Files

### templates/backend/fastapi/requirements.txt.template

```
# Core Framework
fastapi==0.109.0
uvicorn[standard]==0.27.0
pydantic==2.5.0
pydantic-settings==2.1.0

# Database
sqlalchemy==2.0.25
alembic==1.13.0
{{if eq .Database "postgresql"}}
psycopg2-binary==2.9.9
asyncpg==0.29.0
{{end}}
{{if eq .Database "mongodb"}}
motor==3.3.2
pymongo==4.6.0
{{end}}

# Authentication
{{if eq .Auth "jwt"}}
python-jose[cryptography]==3.3.0
passlib[bcrypt]==1.7.4
python-multipart==0.0.6
{{end}}

# Utilities
python-dotenv==1.0.0
httpx==0.26.0

# Testing
pytest==8.0.0
pytest-asyncio==0.23.0
pytest-cov==4.1.0

# Code Quality
ruff==0.1.0
black==24.1.0
mypy==1.8.0
```

### templates/backend/fastapi/README.md.template

```markdown
# {{.ProjectName}} Backend

A modern REST API built with FastAPI, Python 3.11+, and {{.Database}}.

## 🚀 Quick Start

### Prerequisites

- Python 3.11+
- {{.Database}} installed and running
- Virtual environment (recommended)

### Installation

```bash
# Create virtual environment
python -m venv venv

# Activate virtual environment
# On macOS/Linux:
source venv/bin/activate
# On Windows:
.\venv\Scripts\activate

# Install dependencies
pip install -r requirements.txt

# Run database migrations
alembic upgrade head

# Start development server
uvicorn app.main:app --reload
```

The API will be available at [http://localhost:8000](http://localhost:8000).

## 📁 Project Structure

```
backend/
├── app/
│   ├── api/
│   │   ├── routes/              # API route handlers
│   │   │   ├── __init__.py
│   │   │   └── health.py
│   │   └── dependencies.py      # Dependency injection
│   ├── core/
│   │   ├── config.py            # Configuration settings
│   │   └── security.py          # Security utilities
│   ├── models/                  # SQLAlchemy models
│   ├── schemas/                 # Pydantic schemas
│   ├── services/                # Business logic
│   ├── repositories/            # Data access
│   └── main.py                  # Application entry point
├── alembic/                     # Database migrations
├── tests/                       # Test files
├── alembic.ini                  # Alembic configuration
├── requirements.txt             # Python dependencies
└── .env.example                 # Environment variables template

## 🧪 Testing

```bash
# Run tests
pytest

# Run tests with coverage
pytest --cov=app tests/

# Run specific test file
pytest tests/test_health.py -v
```

## 📝 API Documentation

Once the server is running, access:

- **Swagger UI**: http://localhost:8000/docs
- **ReDoc**: http://localhost:8000/redoc
- **OpenAPI JSON**: http://localhost:8000/openapi.json

## 🔧 Configuration

Create a `.env` file from the example:

```bash
cp .env.example .env
```

Environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | Database connection string | Required |
| `SECRET_KEY` | Secret key for JWT | Required |
| `DEBUG` | Enable debug mode | `false` |
| `CORS_ORIGINS` | Allowed CORS origins | `["http://localhost:3000"]` |

## 📚 Learn More

- [FastAPI Documentation](https://fastapi.tiangolo.com)
- [SQLAlchemy Documentation](https://docs.sqlalchemy.org)
- [Alembic Documentation](https://alembic.sqlalchemy.org)

## 📄 License

MIT
```

### templates/backend/fastapi/app/main.py.template

```python
"""
Main FastAPI Application

This module initializes the FastAPI application and configures
middleware, routers, and exception handlers.
"""

from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from app.api.routes import health
from app.core.config import settings

# Create FastAPI application instance
app = FastAPI(
    title="{{.ProjectName}} API",
    description="{{.ProjectDescription}}",
    version="0.1.0",
    docs_url="/docs",
    redoc_url="/redoc",
)

# Configure CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=settings.CORS_ORIGINS,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Include routers
app.include_router(health.router, prefix="/api", tags=["Health"])


@app.on_event("startup")
async def startup_event():
    """
    Application startup event handler.
    
    Use this for initializing connections, caches, etc.
    """
    # Add startup logic here
    pass


@app.on_event("shutdown")
async def shutdown_event():
    """
    Application shutdown event handler.
    
    Use this for cleanup, closing connections, etc.
    """
    # Add cleanup logic here
    pass


@app.get("/", tags=["Root"])
async def root():
    """
    Root endpoint.
    
    Returns basic API information.
    """
    return {
        "name": "{{.ProjectName}} API",
        "version": "0.1.0",
        "docs": "/docs",
    }
```

### templates/backend/fastapi/app/api/routes/health.py.template

```python
"""
Health Check Routes

Provides endpoints for monitoring application health.
"""

from fastapi import APIRouter, Response
from pydantic import BaseModel
from typing import Dict

router = APIRouter()


class HealthResponse(BaseModel):
    """Health check response model."""
    status: str
    version: str
    services: Dict[str, str]


@router.get("/health", response_model=HealthResponse)
async def health_check() -> HealthResponse:
    """
    Health check endpoint.
    
    Returns the current health status of the API.
    
    Returns:
        HealthResponse: Current health status
    """
    return HealthResponse(
        status="healthy",
        version="0.1.0",
        services={
            "api": "up",
            "database": "up",  # Check actual database connection
        }
    )


@router.get("/ready")
async def readiness_check(response: Response):
    """
    Readiness check endpoint.
    
    Verifies that all required services are available.
    
    Returns:
        dict: Readiness status
    """
    # Check database connection
    # Check other dependencies
    
    response.status_code = 200
    return {"ready": True}
```

---

# PART 3: DOCUMENTATION

## 3.1 Root README.md

Create the main project README:

```markdown
<p align="center">
  <img src="website/assets/images/logo/forge-icon.svg" alt="Clause Logo" width="120" />
</p>

<h1 align="center">Clause CLI</h1>

<p align="center">
  <strong>Structure Your AI's Intelligence</strong>
</p>

<p align="center">
  The AI-native project scaffolding tool that guides your AI coding assistant
  to produce consistent, maintainable, architecturally sound code.
</p>

<p align="center">
  <a href="#-features">Features</a> •
  <a href="#-installation">Installation</a> •
  <a href="#-quick-start">Quick Start</a> •
  <a href="#-documentation">Documentation</a> •
  <a href="#-contributing">Contributing</a>
</p>

<p align="center">
  <a href="https://github.com/Mr-Dark-debug/clause-cli/actions">
    <img src="https://github.com/Mr-Dark-debug/clause-cli/workflows/CI/badge.svg" alt="CI Status" />
  </a>
  <a href="https://github.com/Mr-Dark-debug/clause-cli/releases">
    <img src="https://img.shields.io/github/v/release/Mr-Dark-debug/clause-cli" alt="Release" />
  </a>
  <a href="https://github.com/Mr-Dark-debug/clause-cli/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/Mr-Dark-debug/clause-cli" alt="License" />
  </a>
  <a href="https://goreportcard.com/report/github.com/Mr-Dark-debug/clause-cli">
    <img src="https://goreportcard.com/badge/github.com/Mr-Dark-debug/clause-cli" alt="Go Report Card" />
  </a>
</p>

---

## 🎯 Why Clause?

Modern AI coding assistants are powerful, but their power is undirected. They can write virtually any code, but without proper guidance, that code may violate project standards, introduce inconsistencies, or create technical debt.

**Clause provides the direction that transforms raw AI capability into disciplined, project-aligned engineering output.**

### Key Differentiators

- 🤖 **AI-Native Design** - Built from the ground up for AI-assisted development
- 📋 **Behavioral Governance** - Rules that guide AI behavior within your project
- 🧠 **Brainstorm.md** - Novel self-reflection mechanism for autonomous AI problem-solving
- 📦 **Component Registry** - Living inventory of created components for reusability
- 🎨 **Beautiful TUI** - Modern terminal interface with responsive design

---

## ✨ Features

### Interactive Project Wizard

Clause guides you through project configuration with a beautiful terminal UI:

```bash
$ clause init my-project
```

- Choose frontend framework (Next.js, React, Vue, Svelte)
- Choose backend framework (FastAPI, Express, Django)
- Configure database, authentication, and infrastructure
- Set AI governance strictness level

### AI Governance System

Every Clause-generated project includes an `ai_prompt_guidelines/` directory that:

- Defines how AI should behave in your project
- Specifies approved technologies and forbidden patterns
- Provides architectural constraints
- Includes documentation standards
- Offers a Brainstorm.md file for AI self-reflection

### Cross-Platform Support

Works on macOS, Linux, and Windows with native package managers:

- **macOS**: Homebrew
- **Linux**: APT, Snap, AUR
- **Windows**: Winget, Scoop

---

## 📦 Installation

### macOS

```bash
# Homebrew (recommended)
brew install clause-cli/tap/clause

# Or using curl
curl -fsSL https://clause.dev/install.sh | bash
```

### Linux

```bash
# curl | bash
curl -fsSL https://clause.dev/install.sh | bash

# APT (Debian/Ubuntu)
curl -fsSL https://clause.dev/apt/gpg.key | sudo gpg --dearmor -o /usr/share/keyrings/clause.gpg
echo "deb [signed-by=/usr/share/keyrings/clause.gpg] https://clause.dev/apt stable main" | sudo tee /etc/apt/sources.list.d/clause.list
sudo apt update && sudo apt install clause
```

### Windows

```powershell
# Winget
winget install Clause.ClauseCLI

# Scoop
scoop bucket add clause-cli https://github.com/Mr-Dark-debug/scoop-bucket
scoop install clause

# PowerShell
irm https://clause.dev/install.ps1 | iex
```

### Verify Installation

```bash
clause --version
# Output: clause version 1.0.0
```

---

## 🚀 Quick Start

### Create a New Project

```bash
# Interactive mode (recommended)
clause init my-awesome-project

# With a preset
clause init my-saas --preset saas

# Non-interactive mode
clause init my-project --non-interactive \
  --frontend nextjs \
  --backend fastapi \
  --database postgresql
```

### What Clause Creates

```
my-awesome-project/
├── frontend/                    # Your frontend application
│   ├── src/
│   │   ├── app/
│   │   ├── components/
│   │   └── lib/
│   └── package.json
│
├── backend/                     # Your backend application
│   ├── app/
│   │   ├── api/
│   │   ├── models/
│   │   └── services/
│   └── requirements.txt
│
├── ai_prompt_guidelines/        # AI reads this!
│   ├── system_prompt.md         # Core AI behavior rules
│   ├── architecture.md          # Architectural constraints
│   ├── technologies.md          # Approved tech stack
│   ├── documentation.md         # Documentation standards
│   ├── brainstorm.md            # AI working memory
│   ├── registry.json            # Component registry
│   └── context.yaml             # Structured context
│
├── infrastructure/              # Docker, CI/CD configs
│   ├── docker/
│   └── .github/
│
└── README.md                    # Project documentation
```

### Work with AI Assistants

When you open your Clause-generated project in an AI coding assistant (Claude Code, Cursor, Windsurf, etc.), the AI will automatically find the `ai_prompt_guidelines/` directory and follow the rules defined there.

---

## 📚 Documentation

Full documentation is available at [clause.dev/docs](https://clause.dev/docs)

- [Getting Started](https://clause.dev/docs/getting-started)
- [Installation Guide](https://clause.dev/docs/installation)
- [CLI Reference](https://clause.dev/docs/cli-reference)
- [AI Governance](https://clause.dev/docs/governance)
- [Templates](https://clause.dev/docs/templates)
- [Contributing](https://clause.dev/pages/contributing)

---

## 🛠️ CLI Commands

```bash
# Initialize a new project
clause init [project-name]

# Add components to existing project
clause add [component-type]

# Validate project compliance
clause validate

# Manage configuration
clause config [command]

# Update Clause CLI
clause update

# Show version
clause version
```

Run `clause --help` for more information.

---

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](https://clause.dev/pages/contributing) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/Mr-Dark-debug/clause-cli.git
cd clause-cli

# Build
make build

# Run tests
make test

# Run locally
./bin/clause init test-project
```

### Project Structure

```
clause-cli/
├── cmd/forge/              # Entry point
├── internal/               # Private packages
│   ├── cmd/                # CLI commands
│   ├── wizard/             # Interactive wizard
│   ├── config/             # Configuration management
│   ├── template/           # Template engine
│   ├── generator/          # Project generation
│   └── governance/         # AI governance system
├── pkg/                    # Public packages
│   ├── styles/             # Theme and styling
│   ├── tui/                # Terminal UI components
│   └── utils/              # Utilities
└── templates/              # Project templates
```

---

## 📝 License

Clause CLI is released under the [MIT License](LICENSE).

---

## 🙏 Acknowledgments

Built with:
- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - Terminal UI
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling

---

## 📬 Connect

- **GitHub**: [github.com/Mr-Dark-debug/clause-cli](https://github.com/Mr-Dark-debug/clause-cli)
- **Discord**: [Join our community](https://discord.gg/clause)
- **Twitter**: [@ClauseCLI](https://twitter.com/ClauseCLI)

---

<p align="center">
  Made with ❤️ for the developer community
</p>
```

---

## 3.2 CONTRIBUTING.md

```markdown
# Contributing to Clause CLI

Thank you for your interest in contributing to Clause CLI! This document provides guidelines and instructions for contributing.

## 📜 Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment. Please be considerate of others and follow standard open-source community guidelines.

## 🐛 Reporting Bugs

Before creating a bug report, please check existing issues to avoid duplicates.

### Bug Report Template

```markdown
**Description**
A clear description of the bug.

**Steps to Reproduce**
1. Run `clause init test`
2. Select options X, Y, Z
3. Observe error

**Expected Behavior**
What should happen instead.

**Actual Behavior**
What actually happens.

**Environment**
- OS: [e.g., macOS 14]
- Clause version: [e.g., 1.0.0]
- Terminal: [e.g., iTerm2]

**Screenshots**
If applicable, add screenshots.

**Additional Context**
Any other relevant information.
```

## 💡 Requesting Features

Feature requests are welcome! Please use the feature request template:

```markdown
**Is your feature request related to a problem?**
A clear description of the problem.

**Describe the solution you'd like**
A clear description of what you want to happen.

**Describe alternatives you've considered**
Any alternative solutions or features you've considered.

**Additional context**
Any other context, screenshots, or examples.
```

## 🔧 Development Setup

### Prerequisites

- Go 1.21 or higher
- Make (optional, for using Makefile)
- Git

### Getting Started

```bash
# Fork and clone
git clone https://github.com/YOUR_USERNAME/clause-cli.git
cd clause-cli

# Create a branch
git checkout -b feature/your-feature-name

# Install dependencies
go mod download

# Build
make build
# or
go build -o bin/clause ./cmd/forge

# Run tests
make test
# or
go test ./...

# Run locally
./bin/clause init test-project
```

### Project Structure

```
clause-cli/
├── cmd/forge/main.go        # Entry point
├── internal/                # Private packages
│   ├── cmd/                 # CLI commands
│   ├── wizard/              # TUI wizard
│   ├── config/              # Configuration
│   ├── template/            # Template engine
│   ├── generator/           # Project generation
│   └── governance/          # AI governance
├── pkg/                     # Public packages
│   ├── styles/              # Styling
│   ├── tui/                 # Terminal UI
│   └── utils/               # Utilities
└── templates/               # Embedded templates
```

## 📝 Coding Standards

### Go Style

- Follow [Effective Go](https://golang.org/doc/effective_go)
- Use `gofmt` for formatting
- Use `golint` for linting
- Run `go vet` before committing

### Code Organization

- Keep files under 200 lines
- One responsibility per file
- Use clear, descriptive names
- Add documentation comments

### Documentation

```go
// Package example provides example functionality.
//
// This package demonstrates documentation standards.
// Use it as a reference for documenting your code.
package example

// FunctionName does something specific.
//
// More detailed description if needed.
//
// Example:
//
//   result := FunctionName(arg1, arg2)
//   fmt.Println(result)
func FunctionName(arg1, arg2 Type) ReturnType {
    // implementation
}
```

### Testing

- Write tests for all new functionality
- Aim for 80%+ coverage on new code
- Use table-driven tests for multiple cases

```go
func TestFunctionName(t *testing.T) {
    tests := []struct {
        name     string
        input    Type
        expected ReturnType
        wantErr  bool
    }{
        {
            name:     "valid input",
            input:    validInput,
            expected: expectedOutput,
            wantErr:  false,
        },
        {
            name:     "invalid input",
            input:    invalidInput,
            expected: zeroValue,
            wantErr:  true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result, err := FunctionName(tt.input)
            if tt.wantErr {
                assert.Error(t, err)
                return
            }
            assert.NoError(t, err)
            assert.Equal(t, tt.expected, result)
        })
    }
}
```

## 📋 Commit Guidelines

We follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <subject>

<body>

<footer>
```

### Types

- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation only
- `style`: Code style changes (formatting, etc.)
- `refactor`: Code refactoring
- `test`: Adding/updating tests
- `chore`: Maintenance tasks

### Examples

```
feat(wizard): add backend framework selection screen

Add a new wizard screen for selecting backend frameworks
including FastAPI, Express, and Django.

Closes #123
```

```
fix(generator): handle empty project name correctly

Fix a panic that occurred when project name was empty.
Now shows a validation error instead.
```

## 🔀 Pull Request Process

1. **Fork and Branch**: Create a feature branch from `main`
2. **Make Changes**: Implement your changes with tests
3. **Run Tests**: Ensure all tests pass
4. **Update Docs**: Update documentation if needed
5. **Commit**: Use conventional commit messages
6. **Push**: Push to your fork
7. **Open PR**: Create a pull request

### PR Checklist

- [ ] Code compiles correctly
- [ ] All tests pass
- [ ] New code has tests
- [ ] Documentation updated
- [ ] Commit messages follow conventions
- [ ] PR description is clear

### PR Template

```markdown
## Description
Brief description of changes.

## Type of Change
- [ ] Bug fix
- [ ] New feature
- [ ] Breaking change
- [ ] Documentation update

## Testing
Describe how you tested these changes.

## Checklist
- [ ] Tests pass
- [ ] Code follows style guidelines
- [ ] Documentation updated
```

## 🏗️ Release Process

1. Version bumps follow [SemVer](https://semver.org/)
2. Releases are automated via GoReleaser
3. Changelog is updated automatically

## 📞 Getting Help

- Open a GitHub Discussion for questions
- Join our Discord for real-time help
- Check existing issues/PRs before creating new ones

---

Thank you for contributing! 🎉
```

---

## 3.3 CHANGELOG.md

```markdown
# Changelog

All notable changes to Clause CLI are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-01-15

### Added
- Initial stable release
- Interactive TUI wizard with 7 configuration screens
- AI governance system with prompt guidelines
- Brainstorm.md self-reflection mechanism
- Component registry for tracking created components
- Cross-platform support (macOS, Linux, Windows)
- Multiple installation methods (Homebrew, Winget, APT, Scoop)
- Project templates for Next.js, React, Vue, FastAPI, Express, Django
- Non-interactive mode for CI/CD integration
- Preset configurations for common project types
- Self-update mechanism
- Configuration management commands
- Validation commands

### Documentation
- Complete README with quick start guide
- Contributing guidelines
- CLI reference documentation
- AI governance documentation

### Infrastructure
- GoReleaser configuration for automated releases
- GitHub Actions CI/CD pipeline
- Makefile for common tasks

## [0.9.0] - 2025-01-01

### Added
- Beta release for testing
- Core CLI structure with Cobra
- Basic project generation
- Template engine with Go templates
- Configuration loading and saving
- Terminal UI with Bubble Tea and Lip Gloss

### Changed
- Improved error handling
- Enhanced TUI responsiveness

## [0.5.0] - 2024-12-15

### Added
- Alpha release
- Basic `init` command
- Simple project scaffolding
- Initial governance templates

---

## Version History Summary

| Version | Date | Description |
|---------|------|-------------|
| 1.0.0 | 2025-01-15 | Initial stable release |
| 0.9.0 | 2025-01-01 | Beta release |
| 0.5.0 | 2024-12-15 | Alpha release |

---

[1.0.0]: https://github.com/Mr-Dark-debug/clause-cli/releases/tag/v1.0.0
[0.9.0]: https://github.com/Mr-Dark-debug/clause-cli/releases/tag/v0.9.0
[0.5.0]: https://github.com/Mr-Dark-debug/clause-cli/releases/tag/v0.5.0
```

---

## 3.4 LICENSE

```
MIT License

Copyright (c) 2025 Clause CLI Contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

---

## 3.5 docs/cli-reference.md

Create complete CLI reference documentation:

```markdown
# CLI Reference

Complete command reference for Clause CLI.

## Global Flags

```bash
clause [global flags] <command> [command flags] [arguments]
```

| Flag | Short | Description |
|------|-------|-------------|
| `--config` | `-c` | Path to configuration file |
| `--verbose` | `-v` | Enable verbose output |
| `--quiet` | `-q` | Suppress non-essential output |
| `--no-color` | | Disable colored output |
| `--help` | `-h` | Show help |
| `--version` | | Show version |

---

## clause init

Initialize a new project with AI governance.

### Usage

```bash
clause init [project-name] [flags]
```

### Arguments

| Argument | Description | Required |
|----------|-------------|----------|
| `project-name` | Name of the project to create | Yes |

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--preset` | | Use a preset configuration |
| `--non-interactive` | `false` | Skip interactive wizard |
| `--frontend` | | Frontend framework (nextjs, react, vue, svelte, none) |
| `--backend` | | Backend framework (fastapi, express, django, none) |
| `--database` | | Database type (postgresql, mongodb, mysql, none) |
| `--output` | `.` | Output directory |

### Examples

```bash
# Interactive mode
clause init my-project

# With preset
clause init my-saas --preset saas

# Non-interactive with all options
clause init my-project --non-interactive \
  --frontend nextjs \
  --backend fastapi \
  --database postgresql

# In current directory
clause init . --frontend react --backend express
```

### Presets

| Preset | Frontend | Backend | Database |
|--------|----------|---------|----------|
| `saas` | Next.js | FastAPI | PostgreSQL |
| `startup` | React | Express | MongoDB |
| `enterprise` | Next.js | Django | PostgreSQL |
| `minimal` | React | FastAPI | None |

---

## clause add

Add components to an existing project.

### Usage

```bash
clause add <component-type> [flags]
```

### Component Types

| Type | Description |
|------|-------------|
| `frontend` | Add frontend components |
| `backend` | Add backend modules |
| `governance` | Add governance rules |
| `infrastructure` | Add infrastructure configs |

### Examples

```bash
# Add a frontend component
clause add frontend --component auth

# Add backend module
clause add backend --module users

# Add governance rule
clause add governance --rule no-any-type

# Add Docker configuration
clause add infrastructure --docker
```

---

## clause validate

Validate project for governance compliance.

### Usage

```bash
clause validate [path] [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--strict` | `false` | Enable strict validation |
| `--fix` | `false` | Auto-fix issues where possible |
| `--output` | `text` | Output format (text, json) |

### Examples

```bash
# Validate current directory
clause validate

# Validate specific project
clause validate ./my-project

# Strict validation with JSON output
clause validate --strict --output json

# Auto-fix issues
clause validate --fix
```

### Exit Codes

| Code | Meaning |
|------|---------|
| 0 | All checks passed |
| 1 | Validation errors found |
| 2 | Configuration error |

---

## clause config

Manage Clause configuration.

### Usage

```bash
clause config <command> [flags]
```

### Subcommands

| Command | Description |
|---------|-------------|
| `list` | List all configuration values |
| `get` | Get a specific configuration value |
| `set` | Set a configuration value |
| `init` | Initialize configuration file |

### Examples

```bash
# List all config
clause config list

# Get specific value
clause config get default.frontend

# Set a value
clause config set default.frontend nextjs

# Initialize config file
clause config init
```

---

## clause update

Update Clause CLI to the latest version.

### Usage

```bash
clause update [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--channel` | `stable` | Update channel (stable, beta, nightly) |
| `--check` | `false` | Only check for updates, don't install |
| `--force` | `false` | Force update even if already latest |

### Examples

```bash
# Update to latest stable
clause update

# Check for updates
clause update --check

# Update to beta channel
clause update --channel beta
```

---

## clause version

Display version information.

### Usage

```bash
clause version [flags]
```

### Flags

| Flag | Default | Description |
|------|---------|-------------|
| `--json` | `false` | Output in JSON format |

### Examples

```bash
clause version
# Output: clause version 1.0.0 (built 2025-01-15)

clause version --json
# Output: {"version": "1.0.0", "buildDate": "2025-01-15", ...}
```

---

## clause completion

Generate shell completion scripts.

### Usage

```bash
clause completion <shell>
```

### Shells

| Shell | Description |
|-------|-------------|
| `bash` | Bash completion |
| `zsh` | Zsh completion |
| `fish` | Fish completion |
| `powershell` | PowerShell completion |

### Examples

```bash
# Generate bash completion
clause completion bash > /etc/bash_completion.d/clause

# Generate zsh completion
clause completion zsh > "${fpath[1]}/_clause"

# Add to PowerShell profile
clause completion powershell >> $PROFILE
```

---

## Environment Variables

| Variable | Description |
|----------|-------------|
| `CLAUSE_CONFIG` | Path to configuration file |
| `CLAUSE_NO_COLOR` | Disable colored output |
| `CLAUSE_CACHE_DIR` | Cache directory location |

---

## Configuration File

Configuration is stored in YAML format:

```yaml
# ~/.clause/config.yaml

defaults:
  frontend: nextjs
  backend: fastapi
  database: postgresql

presets:
  my-custom:
    frontend: react
    backend: express
    styling: tailwind

updates:
  channel: stable
  checkOnStartup: true
```

---

## Exit Codes

All commands use standard exit codes:

| Code | Meaning |
|------|---------|
| 0 | Success |
| 1 | Error |
| 2 | Configuration error |
| 130 | Interrupted (Ctrl+C) |
```

---

## 3.6 Example Generated Files

### Example .env.example

```env
# Application
APP_NAME={{.ProjectName}}
APP_ENV=development
DEBUG=true

# Frontend
NEXT_PUBLIC_API_URL=http://localhost:8000

# Backend
DATABASE_URL=postgresql://user:password@localhost:5432/{{.ProjectName}}
SECRET_KEY=your-secret-key-change-in-production
CORS_ORIGINS=http://localhost:3000

# Authentication
JWT_SECRET=your-jwt-secret-change-in-production
JWT_EXPIRATION=3600

# External Services (add as needed)
# REDIS_URL=redis://localhost:6379
# AWS_ACCESS_KEY=
# AWS_SECRET_KEY=
```

### Example docker-compose.yml.template

```yaml
version: '3.8'

services:
  frontend:
    build:
      context: ./frontend
      dockerfile: Dockerfile
    ports:
      - "3000:3000"
    environment:
      - NEXT_PUBLIC_API_URL=http://backend:8000
    depends_on:
      - backend
    volumes:
      - ./frontend:/app
      - /app/node_modules

  backend:
    build:
      context: ./backend
      dockerfile: Dockerfile
    ports:
      - "8000:8000"
    environment:
      - DATABASE_URL=postgresql://postgres:postgres@db:5432/{{.ProjectName}}
      - SECRET_KEY=${SECRET_KEY}
    depends_on:
      - db
    volumes:
      - ./backend:/app

  db:
    image: postgres:15-alpine
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=postgres
      - POSTGRES_PASSWORD=postgres
      - POSTGRES_DB={{.ProjectName}}
    volumes:
      - postgres_data:/var/lib/postgresql/data

volumes:
  postgres_data:
```

---

## 📋 DELIVERABLE CHECKLIST

Create all the following files:

### AI Governance Templates
- [ ] `templates/governance/ai_prompt_guidelines/system_prompt.md`
- [ ] `templates/governance/ai_prompt_guidelines/architecture.md`
- [ ] `templates/governance/ai_prompt_guidelines/technologies.md`
- [ ] `templates/governance/ai_prompt_guidelines/brainstorm.md`
- [ ] `templates/governance/ai_prompt_guidelines/documentation.md`
- [ ] `templates/governance/ai_prompt_guidelines/registry.json`
- [ ] `templates/governance/ai_prompt_guidelines/context.yaml`

### Frontend Templates
- [ ] `templates/frontend/nextjs/package.json.template`
- [ ] `templates/frontend/nextjs/README.md.template`
- [ ] `templates/frontend/nextjs/src/app/layout.tsx.template`
- [ ] `templates/frontend/nextjs/src/app/page.tsx.template`
- [ ] `templates/frontend/nextjs/src/app/globals.css.template`
- [ ] `templates/frontend/nextjs/tsconfig.json.template`
- [ ] `templates/frontend/nextjs/next.config.js.template`
- [ ] `templates/frontend/nextjs/tailwind.config.js.template`
- [ ] Similar templates for React, Vue, Svelte

### Backend Templates
- [ ] `templates/backend/fastapi/requirements.txt.template`
- [ ] `templates/backend/fastapi/README.md.template`
- [ ] `templates/backend/fastapi/app/main.py.template`
- [ ] `templates/backend/fastapi/app/core/config.py.template`
- [ ] `templates/backend/fastapi/app/api/routes/health.py.template`
- [ ] Similar templates for Express, Django

### Root Documentation
- [ ] `README.md`
- [ ] `CONTRIBUTING.md`
- [ ] `CHANGELOG.md`
- [ ] `LICENSE`
- [ ] `CODE_OF_CONDUCT.md`
- [ ] `SECURITY.md`

### Documentation Pages
- [ ] `docs/cli-reference.md`
- [ ] `docs/governance.md`
- [ ] `docs/getting-started.md`
- [ ] `docs/installation.md`
- [ ] `docs/templates.md`
- [ ] `docs/configuration.md`

### Example Files
- [ ] `.env.example`
- [ ] `docker-compose.yml.template`
- [ ] `.gitignore.template`
- [ ] `.editorconfig`

---

**Now, create all these files following the templates and standards defined above. Make sure every file is properly formatted, documented, and ready for production use.** 🚀
