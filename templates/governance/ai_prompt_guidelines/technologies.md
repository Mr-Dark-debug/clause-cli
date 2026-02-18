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
