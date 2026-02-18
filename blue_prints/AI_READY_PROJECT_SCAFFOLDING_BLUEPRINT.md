# AI-Ready Project Scaffolding System: Complete Technical & Strategic Blueprint

---

# Executive Vision

## The Problem Space

The modern development landscape faces a critical fragmentation problem. When developers start new projects, they are immediately confronted with an overwhelming cascade of decisions that profoundly impact their entire development lifecycle. Which frontend framework? Which backend technology? How should the folder structure be organized? What about state management, testing infrastructure, CI/CD pipelines, environment configurations, and the countless other architectural decisions that compound into technical debt before a single line of application code is written?

Existing solutions address fragments of this problem. Create React App scaffolds React applications but nothing else. Vite offers lightning-fast builds but leaves backend decisions entirely unaddressed. Yeoman provides generators but requires hunting through an ecosystem of varying quality. Nx and Turborepo excel at monorepo management but come with significant learning curves and presuppose architectural decisions. Cookiecutter provides Python project templates but remains Python-centric.

None of these tools addresses the emerging reality that fundamentally transforms how code is written: the AI-assisted development revolution. Tools like Cursor, Claude Code, GitHub Copilot, and Windsurf have become integral to modern development workflows, yet they operate without structured context about project conventions, architectural decisions, or organizational standards. The result is AI-generated code that, while syntactically correct, often violates project patterns, introduces inconsistent dependencies, and creates what we might call "architectural drift"—the gradual degradation of codebase coherence as AI tools make locally optimal but globally suboptimal decisions.

This blueprint describes a fundamentally different approach: an AI-Operable Project Environment scaffolding system that doesn't just generate project structures but creates comprehensive context systems that guide AI assistants toward consistent, maintainable, and standards-compliant code generation.

## The Core Innovation

The proposed system—provisionally named **Scaffold AI** with the CLI command `scaffold` (alternatives explored in the branding section)—represents a paradigm shift in project scaffolding philosophy. Rather than treating project initialization as a one-time event that produces static artifacts, this system treats project scaffolding as the creation of a persistent context layer that maintains architectural integrity throughout the project's entire lifecycle.

The core innovation lies in the `.ai-prompt-guidelines/` directory (or equivalent naming convention), a structured system of markdown files that serves as the project's architectural constitution. This directory contains not just passive documentation, but active behavioral specifications that AI coding assistants can read, interpret, and follow. It includes:

- **System Prompt Architecture**: A layered prompt structure that combines base behavioral instructions with project-specific constraints, ensuring that AI tools operate within defined boundaries from their first interaction with the codebase.
- **Architecture Decision Records (ADRs)**: Documented decisions about technology choices, patterns, and conventions that inform all subsequent AI-generated code.
- **Component Registry**: A living catalog of reusable components, their APIs, dependencies, and usage patterns that prevents duplicate implementations and encourages consistency.
- **Brainstorm.md**: A designated scratchpad for AI self-reflection, question resolution, and iterative reasoning that eliminates the need for AI tools to interrupt user workflows with clarifying questions.
- **Ops Philosophy**: Guidelines for how operational concerns—logging, monitoring, error handling, security—should be integrated into generated code.
- **Documentation Standards**: Templates and requirements for how complex logic should be explained, ensuring that generated code remains comprehensible to human developers.
- **Brand and Theme Guidelines**: Specifications for visual consistency, preventing the AI from hardcoding values or creating inconsistent user interfaces.

## The Strategic Positioning

Scaffold AI occupies a unique position in the development tools ecosystem. It is not a build tool like Vite or Webpack. It is not a framework like Next.js or Django. It is not an AI assistant like Cursor or Claude Code. Instead, it is infrastructure for AI-assisted development—a meta-tool that enhances the effectiveness of other tools by providing structured context.

This positioning creates multiple strategic advantages:

1. **Complementarity Rather Than Competition**: Scaffold AI enhances rather than replaces existing tools. A developer can use Scaffold AI to initialize a project, then use Vite for builds, Next.js for frontend routing, and Cursor for AI assistance. Each tool does what it does best.

2. **Lock-In Resistance**: Because the output is primarily structured markdown files and conventional project structures, developers are never locked into the Scaffold AI ecosystem. If they stop using the tool, they retain all the context files and project structures.

3. **AI-Agnostic Architecture**: The prompt guidelines system is designed to work with any AI coding assistant—Claude, GPT-4, Gemini, local models, or future systems not yet invented. This future-proofs the investment in creating comprehensive context files.

4. **Progressive Enhancement**: Teams can adopt Scaffold AI incrementally, starting with basic project scaffolding and progressively adding more sophisticated AI context systems as they become comfortable with the approach.

## The Developer Experience Vision

The intended user experience is deliberately simple, masking the complexity of what happens beneath the surface:

```bash
# Global installation (one-time)
curl -fsSL https://get.scaffoldai.dev | bash

# Or via package managers
brew install scaffold-ai
winget install ScaffoldAI.Scaffold
npm install -g @scaffold-ai/cli

# Project initialization
cd ~/projects
scaffold init my-awesome-project
```

Upon running `scaffold init`, the user enters an interactive terminal UI—a polished, responsive, aesthetically pleasing experience powered by Bubble Tea and Lip Gloss (for the Go implementation). The wizard guides users through:

1. **Project Type Selection**: Web application, API, CLI tool, mobile backend, data pipeline, etc.
2. **Frontend Stack Configuration**: Framework choice (React, Vue, Svelte, Angular, Next.js, Nuxt, etc.), state management approach, styling system, component architecture preferences.
3. **Backend Stack Configuration**: Framework choice (FastAPI, Django, Express, NestJS, Go Fiber, etc.), database decisions, API style (REST, GraphQL, tRPC), authentication approach.
4. **Infrastructure Choices**: Container preferences, cloud platform, CI/CD system, monitoring approach.
5. **Development Preferences**: Package manager, linter, formatter, testing framework, environment management.
6. **AI Integration Settings**: Primary AI assistant, context management preferences, documentation verbosity.

The result is not just a folder structure but a comprehensive, immediately productive development environment with all configuration files, boilerplate code, and—most importantly—the AI context system pre-configured and ready to guide subsequent development.

## The Long-Term Vision

Beyond the initial scaffolding capability, Scaffold AI is positioned to become the standard for AI-operable project environments. Future capabilities include:

- **Template Marketplace**: A community-driven ecosystem of project templates optimized for specific use cases—e-commerce platforms, SaaS applications, data dashboards, etc.
- **Team Sync**: Cloud-based synchronization of organizational standards, ensuring that all team members scaffold projects with consistent conventions.
- **AI Analytics**: Insights into how AI tools interact with the codebase, identifying patterns where AI assistants struggle or produce inconsistent code.
- **Continuous Context Evolution**: Tools for updating context files as the project evolves, ensuring that AI guidelines remain synchronized with actual project practices.

This vision positions Scaffold AI as essential infrastructure for the AI-assisted development era—a tool that every developer using AI coding assistants will want to have in their arsenal.

---

# Market Landscape Analysis

## The Scaffolding Tool Ecosystem

### JavaScript/TypeScript Ecosystem Scaffolding Tools

#### Create React App (CRA)

Create React App dominated React project scaffolding for years, establishing the de facto standard for how React applications should be structured. Its contribution was eliminating configuration complexity—developers could run a single command and receive a working React application with hot reloading, build optimization, and testing infrastructure pre-configured.

However, CRA's approach reveals fundamental limitations that inform our understanding of what scaffolding tools should and should not do. First, CRA is deeply opinionated about its technology choices. While this reduces decision fatigue for beginners, it creates friction for teams with different preferences. Want to use a different state management library? Different styling approach? You can, but you're fighting against the grain. Second, CRA provides no guidance for backend development. Modern applications rarely consist of frontend alone, yet CRA treats backend as entirely out of scope. Third, and most critically for our purposes, CRA provides no AI context system. The generated project has no mechanism for informing AI tools about project conventions or architecture.

Vite has largely superseded CRA for new projects, offering dramatically faster build times through native ES modules. Vite's approach is less opinionated, supporting multiple frameworks through plugins. However, Vite remains fundamentally a build tool that happens to include scaffolding capabilities. Its scaffolding is secondary to its primary mission of fast development servers and optimized production builds. The project structure it generates is minimal—intentionally so—leaving teams to establish their own conventions.

#### Vite

Vite represents the modern approach to JavaScript build tooling, achieving popularity through its innovative use of native ES modules during development. When scaffolding projects, Vite offers framework-agnostic templates and community-contributed templates for various frameworks, but its scaffolding capabilities remain basic compared to what a dedicated scaffolding tool could provide.

Vite's strength lies in its plugin architecture, which allows extensive customization of the build process. However, this same flexibility means that Vite-scaffolded projects begin with minimal structure and require additional decisions about folder organization, state management, testing infrastructure, and other architectural concerns. For experienced teams with established conventions, this minimal starting point is an advantage—they can layer their preferences onto the basic structure. For less experienced teams or those seeking guidance, the minimal scaffolding offers insufficient direction.

From an AI context perspective, Vite-generated projects are essentially blank slates. An AI assistant encountering a Vite-scaffolded project has no built-in guidance about the team's preferences, architectural decisions, or coding standards. Every AI interaction begins from zero context, leading to potentially inconsistent code generation.

#### Yeoman

Yeoman pioneered the generator-based approach to project scaffolding, creating an ecosystem where community members could create and share scaffolding generators for various project types. This model has significant advantages: it enables specialized generators for specific use cases, allows community contributions to expand coverage, and provides a consistent interface across different generator types.

The Yeoman ecosystem includes generators for virtually every framework and technology combination imaginable. Need a generator for a React frontend with a Django backend? One probably exists. Want a generator for a specific corporate template? You can create it. This extensibility is Yeoman's greatest strength.

However, Yeoman's approach has notable weaknesses. The quality of generators varies enormously—some are meticulously maintained and updated, while others languish with outdated dependencies and abandoned features. Finding the right generator requires navigating this uneven ecosystem. Yeoman generators are also inherently one-time operations; they generate initial project structure but provide no mechanism for ongoing guidance or context maintenance. Once the scaffolding is complete, Yeoman's involvement ends.

For AI context, Yeoman offers nothing. Its generators produce conventional project structures without any mechanism for informing AI assistants about project conventions. A Yeoman-generated project is indistinguishable from a manually created project from an AI perspective.

#### Nx

Nx represents a different category of tool—not just a scaffolding system but a comprehensive monorepo management platform. Nx excels at managing complex project graphs, enforcing code organization rules, and providing tooling for large-scale development. Its scaffolding capabilities are sophisticated, including generators for creating new libraries, applications, and components within the monorepo structure.

Nx's approach is highly opinionated, prescribing specific organizational patterns that work well for large enterprises but may feel heavyweight for smaller teams. The learning curve is significant; teams must understand Nx's mental models to use it effectively. The payoff for this investment is substantial—Nx provides powerful capabilities for managing complexity that simpler tools cannot match.

From an AI perspective, Nx's structured approach provides implicit context that AI tools can potentially leverage. The enforced patterns create predictability that could theoretically improve AI-generated code consistency. However, Nx has no explicit mechanism for AI context management. The structure is there, but there's no guidance telling AI tools how to interpret it or what conventions to follow when generating new code within the Nx structure.

#### Turborepo

Turborepo focuses on monorepo build optimization rather than scaffolding per se. Its value proposition is making monorepo builds fast through intelligent caching and task orchestration. While Turborepo can be used alongside scaffolding tools, it doesn't provide significant scaffolding capabilities itself.

For teams already committed to monorepo development, Turborepo offers compelling performance benefits. However, its narrow focus means it doesn't address the broader scaffolding question of how to structure and configure new projects or components. Turborepo is complementary to, rather than competitive with, dedicated scaffolding tools.

### Python Ecosystem Scaffolding Tools

#### Cookiecutter

Cookiecutter pioneered Python project templating and remains the dominant tool in its space. Its approach is elegant: define a template directory structure with placeholder variables, then use Cookiecutter to create new projects from the template with user-provided values. The template system is simple but powerful, supporting conditional file generation, variable substitution, and hooks for custom logic.

The Cookiecutter ecosystem includes hundreds of templates for various project types—from basic Python packages to Django applications to data science projects. This breadth of coverage is a significant advantage, enabling developers to find templates that match their specific needs.

Cookiecutter's weaknesses mirror those of Yeoman: template quality varies widely, templates can become outdated, and there's no mechanism for ongoing project guidance. Once the project is generated, Cookiecutter's involvement ends. The generated project has no memory of the template or the decisions made during generation.

For AI context, Cookiecutter templates are conventional project structures without any built-in mechanism for AI guidance. Some sophisticated templates might include documentation files that could theoretically inform AI tools, but this is not a standard practice and there's no consistent framework for AI context across templates.

#### Django startproject

Django's built-in scaffolding—`django-admin startproject` and `python manage.py startapp`—represents the framework-integrated approach to scaffolding. These commands create Django-specific project structures optimized for the framework's conventions. The generated structure is opinionated but well-documented, following Django's "batteries included" philosophy.

The strength of framework-integrated scaffolding is its tight alignment with framework best practices. The generated structure is exactly what the framework expects, reducing friction between project organization and framework requirements. However, this approach is limited to the specific framework. Django's scaffolding cannot help with frontend decisions, infrastructure configuration, or the many concerns that fall outside Django's scope.

Django's project structure provides implicit context that experienced developers (and potentially AI tools) can interpret. The standard layout of models, views, templates, and static files creates predictability. However, there's no explicit AI context system, no mechanism for documenting team-specific conventions, and no way to extend the context as the project evolves.

### Framework-Specific Scaffolding

#### Angular CLI

Angular CLI represents perhaps the most comprehensive framework-integrated scaffolding system. It generates not just initial project structure but ongoing artifacts throughout the project lifecycle—components, services, modules, pipes, guards, and more. The CLI enforces Angular conventions and integrates with the framework's build system, testing infrastructure, and deployment tooling.

Angular CLI's comprehensiveness creates a self-reinforcing pattern: because the CLI handles so much, developers are strongly incentivized to follow its conventions. This creates consistency across Angular projects, making it easier for developers to move between codebases. The trade-off is significant opinionation—teams that disagree with Angular CLI's choices face friction when trying to deviate.

For AI context, Angular CLI's generated structure provides substantial implicit guidance. An AI tool familiar with Angular can reasonably infer many conventions from the standard structure. However, there's no explicit mechanism for documenting team-specific patterns, custom architectural decisions, or organizational standards that extend beyond Angular defaults.

#### Rails Generators

Ruby on Rails pioneered the generator pattern that many subsequent tools adopted. Rails generators create not just initial project structure but ongoing scaffolding for models, controllers, migrations, and other artifacts. The generated code follows Rails conventions meticulously, embodying the framework's "convention over configuration" philosophy.

Rails' approach creates exceptionally consistent codebases when teams follow the conventions. The generated structure is comprehensible to any Rails developer, and the framework's opinionated nature means there are often fewer decisions to make. However, teams that need to deviate from Rails conventions sometimes find the framework's opinionated nature constraining.

Rails' comprehensive scaffolding provides rich implicit context, but like other tools, it lacks explicit AI context management. An AI tool working with a Rails project can infer conventions from the structure, but there's no mechanism for documenting deviations from standard conventions or team-specific patterns.

### Go Ecosystem Scaffolding

#### Cobra CLI

Cobra is the de facto standard for building command-line applications in Go, powering major projects including Kubernetes, Docker, GitHub CLI, and Hugo. While Cobra is primarily a library for building CLI applications rather than a scaffolding tool per se, it includes a generator (`cobra-cli init`) that creates initial project structure for CLI applications.

Cobra's generated structure follows established Go conventions, organizing commands into a predictable hierarchy with support for persistent flags, nested commands, and configuration integration via Viper. The structure is minimal but functional—appropriate for CLI tools but not designed for full application scaffolding.

For AI context, Cobra-generated projects are conventional Go applications without any built-in AI guidance mechanism. The predictability of Cobra's structure provides some implicit context for AI tools familiar with Go CLI patterns, but there's no explicit system for documenting conventions or guiding AI behavior.

### Analysis: Common Patterns and Gaps

Across all these scaffolding tools, several patterns emerge:

**Pattern 1: One-Time Generation** - Almost all scaffolding tools operate as one-time generators. They create initial structure and then their involvement ends. There's no mechanism for ongoing guidance or context maintenance as the project evolves.

**Pattern 2: Technology Narrowness** - Most tools focus on a single technology or framework. Frontend tools don't address backend. Backend tools don't address frontend. Monorepo tools assume you've already made the monorepo decision. This fragmentation creates gaps for full-stack projects.

**Pattern 3: Implicit Rather Than Explicit Context** - The context that scaffolding tools provide is implicit in the generated structure. Experienced developers (and potentially sophisticated AI tools) can infer conventions from file organization. But there's no explicit mechanism for documenting conventions, explaining decisions, or guiding future development.

**Pattern 4: No AI Awareness** - Critically, no existing scaffolding tool is designed with AI-assisted development in mind. None generates structured context that AI coding assistants can use to produce more consistent, convention-compliant code.

These patterns reveal the gap that Scaffold AI addresses: a need for scaffolding that transcends single technologies, provides explicit context for ongoing development, and is designed from the ground up for the AI-assisted development era.

## The AI Developer Tool Ecosystem

### AI-Native IDEs

#### Cursor

Cursor has emerged as the leading AI-native code editor, built on VS Code's foundation with deep AI integration. Its standout features include multi-file editing, codebase-aware chat, and the innovative `.cursorrules` system for project-specific AI instructions. The `.cursorrules` file allows teams to define behavioral guidelines that Cursor's AI follows when working with the codebase.

Cursor's `.cursorrules` system represents the closest existing approach to what Scaffold AI proposes. Teams can define coding standards, architectural preferences, and behavioral guidelines that inform AI responses. However, `.cursorrules` is limited in several ways: it's a single file (though Cursor now supports multiple rule files), its format is informal, it's Cursor-specific (not portable to other AI tools), and there's no structured approach to creating and maintaining these rules.

Cursor's approach validates the core thesis of Scaffold AI: teams benefit significantly from providing structured context to AI tools. However, Cursor's implementation is just the beginning. A more comprehensive, AI-agnostic, systematically generated context system could provide even greater benefits.

#### Windsurf

Windsurf, developed by Codeium, represents another AI-native IDE approach. Like Cursor, it provides deep AI integration within the editing experience. Windsurf emphasizes its "Flow" feature for multi-step operations and its memory system for maintaining context across conversations.

Windsurf's approach to context is less structured than Cursor's `.cursorrules`. The tool builds context through its memory system and codebase analysis rather than explicit instruction files. This approach is more automatic but less controllable—teams cannot precisely specify how AI should behave, instead relying on the tool's inference capabilities.

#### Replit Ghostwriter

Replit's Ghostwriter provides AI assistance within the Replit browser-based development environment. Its strength lies in tight integration with Replit's collaborative features and instant deployment capabilities. However, its approach to context is minimal—it primarily relies on the immediate code context rather than project-wide conventions or explicit guidelines.

### AI Coding Assistants

#### GitHub Copilot

GitHub Copilot pioneered AI pair programming and remains widely adopted. Its strength lies in its training on vast amounts of public code and its tight integration with VS Code and JetBrains IDEs. Copilot excels at line-by-line code completion and can generate entire functions based on context.

Copilot's context awareness is limited to the immediately visible code and recently edited files. There's no mechanism for teams to provide explicit guidance about coding standards or architectural conventions. The tool operates more as a sophisticated auto-complete than a context-aware assistant.

Recent additions like Copilot Chat and workspace awareness have expanded Copilot's capabilities, but it still lacks the structured context system that would enable truly convention-compliant code generation.

#### Claude Code

Claude Code is Anthropic's terminal-based AI coding assistant, designed to operate directly in the development environment through the command line. Unlike IDE-integrated tools, Claude Code works with the entire filesystem and can execute commands, edit files, and interact with version control.

Claude Code's strength lies in its ability to understand and manipulate entire projects rather than just visible code. It can read multiple files, understand project structure, and make coordinated changes across the codebase. However, like other tools, it lacks explicit mechanisms for teams to provide behavioral guidance. Claude Code infers conventions from existing code, which works well for established codebases but provides less guidance for new projects.

#### Gemini CLI

Google's Gemini CLI provides command-line access to Gemini models for coding tasks. It offers generous free tier limits (60 requests per minute, 1000 per day at time of research) and integrates with Google's ecosystem.

Gemini CLI's context handling is similar to other AI assistants—it can read and analyze code but has no structured mechanism for receiving project-specific behavioral guidelines. Its strength lies in its accessibility and cost-effectiveness rather than sophisticated context management.

#### Aider

Aider is an open-source AI coding assistant that operates from the command line. Its distinguishing feature is tight integration with git—it stages and commits AI-generated changes, maintains a clear record of what was AI-generated, and can work across multiple files coherently.

Aider's approach is particularly interesting for its transparency and control. Users explicitly specify which files to include in context, preventing AI from making unexpected modifications. However, Aider lacks a structured system for defining behavioral guidelines or architectural constraints.

### AI Web Development Platforms

#### Bolt.new

Bolt.new, developed by StackBlitz, represents a different approach: AI-powered full application generation in the browser. Users describe what they want to build, and Bolt generates a complete application with frontend, backend logic, and deployment configuration.

Bolt.new is impressive for rapid prototyping but limited for production development. Generated code follows Bolt's conventions rather than team-specific standards. There's no mechanism for teams to influence how Bolt generates code or to maintain consistency across multiple Bolt-generated projects.

#### v0 by Vercel

v0 focuses on UI component generation, allowing users to describe interfaces in natural language and receive production-ready React components. It's tightly integrated with Vercel's ecosystem and excels at creating shadcn/ui-based components.

v0 is excellent for accelerating UI development but narrow in scope. It generates isolated components rather than complete applications and has no mechanism for team-specific design systems or coding standards.

### Analysis: The AI Context Gap

The AI developer tool ecosystem reveals a clear pattern: powerful AI capabilities coupled with minimal structured context management. Tools like Cursor are beginning to address this gap with features like `.cursorrules`, but the approach remains:

1. **Tool-Specific**: Each AI tool has its own context mechanism (if any), preventing portability.
2. **Informal**: Context specification formats are ad-hoc without structured schemas.
3. **Manual**: Teams must create and maintain context files manually, without tooling support.
4. **Narrow**: Context systems focus on coding standards without addressing broader concerns like architecture documentation, component registries, or operational philosophy.

This is precisely the gap that Scaffold AI addresses: providing a comprehensive, structured, AI-agnostic context system that teams can generate consistently and maintain systematically.

---

# Competitor Deep Dive

## Detailed Competitor Analysis

### Yeoman: The Generator Pioneer

**What Problem Yeoman Solves**: Yeoman addresses the need for customizable, community-driven project scaffolding. It provides a framework for creating generators that can produce any type of project structure, enabling specialization and community contribution.

**Yeoman's Strengths**:
- **Extensibility**: The generator system allows virtually any project structure to be encoded and reproduced.
- **Community Ecosystem**: Hundreds of generators exist for various technology combinations.
- **Customization**: Generators can include prompts, conditional logic, and post-install hooks.
- **Established Patterns**: Well-documented generator creation process.

**Yeoman's Weaknesses**:
- **Ecosystem Fragmentation**: Generator quality varies enormously; many are unmaintained.
- **One-Time Operation**: Once scaffolding is complete, Yeoman has no ongoing role.
- **JavaScript-Centric**: While generators can produce any output, the tooling is JavaScript-focused.
- **No AI Context**: Generated projects have no mechanism for AI guidance.
- **Manual Maintenance**: Teams must manually update project conventions as they evolve.

**Why Yeoman Doesn't Solve This Problem**: Yeoman's fundamental limitation is its transactional nature. It generates artifacts but doesn't create persistent context. An AI assistant working with a Yeoman-generated project has no way to know what generator was used, what options were selected, or what conventions the team has established since generation.

**Architectural Philosophy**: Yeoman follows a generator-consumer model where generators define templates and prompts, and consumers receive generated artifacts. The relationship ends at generation time.

**AI Integration Limitations**: Yeoman has no AI integration story. It predates AI-assisted development and hasn't evolved to address AI context needs.

**Gaps We Can Exploit**:
1. Create structured context that persists beyond scaffolding.
2. Build AI-awareness into the generated project structure.
3. Provide tooling for maintaining and evolving context.
4. Design for AI-agnostic context that works across AI tools.

### Nx: The Monorepo Powerhouse

**What Problem Nx Solves**: Nx addresses the complexity of managing large monorepos with multiple applications and libraries. It provides tooling for dependency graph visualization, affected test selection, code generation, and consistent architectural enforcement.

**Nx's Strengths**:
- **Comprehensive Tooling**: Integrated solution for build, test, generate, and analyze operations.
- **Architectural Enforcement**: Can define and enforce module boundaries and dependency rules.
- **Performance**: Intelligent caching and parallelization for fast CI/CD.
- **Enterprise Focus**: Built for the needs of large development organizations.
- **Extensible Generators**: Custom generators for organization-specific patterns.

**Nx's Weaknesses**:
- **Learning Curve**: Significant complexity requires substantial investment to master.
- **Opinionation**: Prescribes specific organizational patterns that may not fit all teams.
- **Overhead**: Heavy tooling that may be excessive for smaller projects.
- **JavaScript/TypeScript Focus**: Less support for polyglot environments.
- **No AI Context**: While structured, Nx doesn't generate AI-readable behavioral guidelines.

**Why Nx Doesn't Solve This Problem**: Nx provides structural organization but not semantic guidance. It tells you where files should go and what can depend on what, but it doesn't tell an AI assistant how to write code, what conventions to follow, or how to document complex logic.

**Architectural Philosophy**: Nx emphasizes constraints and automation—enforcing architectural decisions through tooling rather than documentation. This works well for structural concerns but less well for coding style, documentation standards, and behavioral patterns.

**AI Integration Limitations**: Nx's structure provides implicit context but no explicit AI guidance. An AI tool can observe that certain imports are forbidden by Nx rules, but it has no guidance about preferred patterns, naming conventions, or documentation expectations.

**Gaps We Can Exploit**:
1. Combine structural organization with semantic guidance.
2. Reduce complexity while maintaining power through smart defaults.
3. Generate explicit AI context alongside structural enforcement.
4. Support polyglot projects with unified context.

### Cookiecutter: Python's Template Standard

**What Problem Cookiecutter Solves**: Cookiecutter provides simple, powerful project templating for Python and other languages. It enables reusable templates with variable substitution and conditional generation.

**Cookiecutter's Strengths**:
- **Simplicity**: Minimal concepts to learn; straightforward template structure.
- **Flexibility**: Can template any type of project, not just Python.
- **Template Inheritance**: Supports template chaining and extension.
- **Wide Adoption**: De facto standard in Python ecosystem.
- **Hook System**: Pre- and post-generation hooks for custom logic.

**Cookiecutter's Weaknesses**:
- **Template Fragmentation**: Quality varies; many templates become outdated.
- **One-Shot Generation**: No mechanism for ongoing project evolution.
- **Python-Centric Ecosystem**: While tool is flexible, ecosystem is Python-focused.
- **No Context Persistence**: Generated projects have no memory of template choices.
- **No AI Awareness**: Templates are conventional files without AI context.

**Why Cookiecutter Doesn't Solve This Problem**: Like Yeoman, Cookiecutter is transactional. The template's role ends at generation time. Subsequent development has no access to template context, and there's no mechanism for AI tools to understand template decisions.

**Architectural Philosophy**: Cookiecutter follows a template-instantiation model where templates define structure and prompts, and instances receive generated files. The relationship is fundamentally ephemeral.

**AI Integration Limitations**: Cookiecutter templates can include any files, so theoretically could include AI context files. However, there's no standard for this, no tooling support, and templates in practice don't include AI guidance.

**Gaps We Can Exploit**:
1. Build AI context into the core output, not as an optional add-on.
2. Create tooling for maintaining context post-generation.
3. Establish standards for AI context that transcend individual tools.
4. Provide context that evolves with the project.

### Cursor: The AI-Native Editor

**What Problem Cursor Solves**: Cursor provides deep AI integration within a familiar VS Code-based editor. Its `.cursorrules` system allows teams to define behavioral guidelines that influence AI responses.

**Cursor's Strengths**:
- **AI Integration**: Purpose-built for AI-assisted development.
- **`.cursorrules` System**: Mechanism for project-specific AI instructions.
- **Multi-File Awareness**: Can understand and edit across multiple files.
- **Codebase Chat**: Natural language interaction with entire codebase.
- **VS Code Compatibility**: Familiar interface with AI enhancements.

**Cursor's Weaknesses**:
- **Editor Lock-In**: Works only within Cursor editor.
- **Informal Rules**: `.cursorrules` format is ad-hoc without structured validation.
- **Single Tool**: Rules don't port to other AI assistants.
- **Manual Maintenance**: Teams must create and maintain rules manually.
- **No Scaffolding Integration**: Rules must be created separately from project setup.

**Why Cursor Doesn't Solve This Problem**: Cursor provides a mechanism for AI context but not comprehensive scaffolding. Teams using Cursor must still use other tools for project initialization and manually create `.cursorrules` files. There's no integration between scaffolding and context creation.

**Architectural Philosophy**: Cursor enhances the editing experience with AI capabilities, treating context as an enhancement to existing development workflows rather than a fundamental project artifact.

**AI Integration Limitations**: While Cursor has the most sophisticated AI context system among existing tools, it's limited to Cursor itself. Teams using multiple AI tools (Claude Code for terminal work, Copilot for quick completions, Cursor for complex edits) can't share context across tools.

**Gaps We Can Exploit**:
1. Create AI-agnostic context that works across all tools.
2. Integrate context creation with project scaffolding.
3. Provide structured, validated context format.
4. Build tooling for context maintenance and evolution.

### Vite: The Build-First Scaffolder

**What Problem Vite Solves**: Vite provides fast, modern build tooling with basic scaffolding capabilities. Its primary value is dramatically improved development server performance through native ES modules.

**Vite's Strengths**:
- **Performance**: Orders of magnitude faster than Webpack-based alternatives.
- **Framework Agnostic**: Supports multiple frameworks through plugins.
- **Minimal Configuration**: Works well out-of-the-box with sensible defaults.
- **Active Development**: Rapid improvement and responsive maintainers.
- **Ecosystem**: Rich plugin ecosystem for various needs.

**Vite's Weaknesses**:
- **Build Focus**: Primarily a build tool; scaffolding is secondary.
- **Minimal Structure**: Generated projects are minimal, requiring additional setup.
- **No Backend Story**: Frontend-focused with no backend scaffolding.
- **No AI Context**: Generated projects have no AI guidance mechanism.
- **No Opinionation**: Minimal structure means no guidance on conventions.

**Why Vite Doesn't Solve This Problem**: Vite's minimal approach is intentional—it avoids opinionation to maximize flexibility. This is appropriate for a build tool but means Vite-scaffolded projects need significant additional setup and provide no guidance for subsequent development.

**Architectural Philosophy**: Vite embraces minimalism and plugin-based extensibility. The core provides essential build capabilities, and plugins extend functionality for specific use cases.

**AI Integration Limitations**: Vite has no AI integration story. It generates minimal project structure without any mechanism for AI context.

**Gaps We Can Exploit**:
1. Provide opinionated defaults while maintaining flexibility.
2. Generate comprehensive structure, not minimal scaffolding.
3. Build AI context into the core offering.
4. Support full-stack development, not just frontend.

## Competitive Positioning Matrix

| Dimension | Yeoman | Nx | Cookiecutter | Cursor | Vite | Scaffold AI |
|-----------|--------|-----|--------------|--------|------|-------------|
| Full-Stack Support | Partial | Yes | Partial | N/A | No | Yes |
| AI Context System | No | No | No | Yes (Cursor-only) | No | Yes (AI-agnostic) |
| Ongoing Guidance | No | Structural only | No | Rules system | No | Yes |
| Template Ecosystem | Large | Built-in | Large | N/A | Plugin-based | Planned |
| Cross-Platform | Yes | Yes | Yes | Yes | Yes | Yes |
| Opinionated Defaults | Varies | Yes | Varies | No | No | Yes |
| Customization Depth | High | High | Medium | N/A | High | High |
| Learning Curve | Medium | High | Low | Low | Low | Medium |
| Post-Generation Support | No | Yes | No | N/A | No | Yes |
| Multi-Framework | Yes | Yes | Yes | N/A | Yes | Yes |
| Enterprise Ready | Partial | Yes | Partial | Yes | Yes | Planned |

### Strategic Differentiation

Scaffold AI differentiates from all competitors through its AI-Operable Project Environment concept. While other tools generate static artifacts, Scaffold AI generates living context systems that persist throughout the project lifecycle and guide AI-assisted development.

Key differentiators:

1. **AI-First Architecture**: Every aspect of the generated project is designed with AI context in mind, from the initial structure to ongoing maintenance tooling.

2. **Brainstorm.md Paradigm**: The innovative scratchpad system allows AI tools to maintain internal reasoning, resolve ambiguities, and continue execution without interrupting users—uniquely enabling autonomous AI development.

3. **AI-Agnostic Context**: Unlike Cursor's tool-specific `.cursorrules`, Scaffold AI generates context that works across all AI coding assistants.

4. **Integrated Scaffolding + Context**: The scaffolding process creates both project structure and AI context in a single, integrated operation.

5. **Living Context**: Context files are designed to evolve with the project, with tooling support for maintenance and updates.

6. **Comprehensive Coverage**: Context covers not just coding standards but architecture decisions, component registries, operational philosophy, and documentation requirements.

---

# Architecture Blueprint

## System Architecture Overview

The Scaffold AI system comprises several interconnected subsystems that work together to deliver the complete scaffolding experience. The architecture is designed for extensibility, maintainability, and performance while keeping the core distribution lightweight.

### High-Level Architecture Components

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           Scaffold AI CLI                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   Command   │  │     TUI     │  │   Config    │  │   Update    │        │
│  │   Router    │  │   Engine    │  │   Manager   │  │   System    │        │
│  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘        │
├─────────────────────────────────────────────────────────────────────────────┤
│                           Core Engine Layer                                  │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   Template  │  │   Context   │  │   File      │  │   Plugin    │        │
│  │   Engine    │  │   Builder   │  │   Generator │  │   System    │        │
│  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘        │
├─────────────────────────────────────────────────────────────────────────────┤
│                           Template Library                                   │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐  ┌─────────────┐        │
│  │   Frontend  │  │   Backend   │  │    Full     │  │   Context   │        │
│  │  Templates  │  │  Templates  │  │   Stack     │  │  Templates  │        │
│  └─────────────┘  └─────────────┘  └─────────────┘  └─────────────┘        │
└─────────────────────────────────────────────────────────────────────────────┘
```

## CLI Core Subsystem

### Command Router

The command router is built on Cobra, the industry-standard CLI framework for Go applications. Cobra provides hierarchical command structure, flag parsing, help generation, and shell completion out of the box. The router maps user commands to their respective handlers while maintaining a clean separation between command definition and business logic.

**Core Commands**:

- `scaffold init [project-name]` - Initialize a new project with interactive wizard
- `scaffold init [project-name] --template <template>` - Initialize from specific template
- `scaffold init [project-name] --non-interactive` - Initialize with defaults (CI/CD use case)
- `scaffold add <component-type> <name>` - Add new component to existing project
- `scaffold context update` - Update AI context files
- `scaffold context check` - Validate AI context consistency
- `scaffold template list` - List available templates
- `scaffold update` - Self-update the CLI
- `scaffold version` - Display version information

**Command Structure Design**:

The command hierarchy follows a verb-noun pattern that feels natural to developers familiar with tools like git, docker, and kubectl. Each command has a handler function that orchestrates the appropriate subsystems. Commands are designed to be composable—the `add` command, for instance, reuses the same template engine as `init` but in a different context.

### Terminal UI Engine

The terminal UI engine is built on Bubble Tea, Charmbracelet's Elm-inspired TUI framework. Bubble Tea provides a clean abstraction for building interactive terminal applications through its Model-Update-View architecture.

**Model-Update-View Architecture**:

The Model represents the complete state of the wizard—current step, user selections, validation state, available options. The Update function handles events (keyboard input, window resize, timer ticks) and produces updated models along with commands (side effects like API calls or timer scheduling). The View function renders the model as a string, which Bubble Tea displays in the terminal.

**Responsive Design Strategy**:

Terminal sizes vary enormously—from massive monitors to small SSH windows to mobile terminals. The UI engine implements responsive design patterns adapted for terminal contexts:

1. **Minimum Window Detection**: Before rendering, detect terminal dimensions. If below minimum (e.g., 80x24), display a message requesting resize rather than broken UI.

2. **Dynamic Layout Calculation**: Layout algorithms recalculate on window resize events. Multi-column layouts collapse to single column on narrow terminals. Progress indicators adapt to available width.

3. **Content Truncation**: Long option descriptions are truncated with ellipses on narrow displays. Full text is available via tooltip-style expand (pressing a key to see more detail).

4. **Scrollable Views**: When content exceeds available space, implement scrolling rather than truncation. This is particularly important for template selection where many options may be available.

**Aesthetic Implementation with Lip Gloss**:

Lip Gloss provides declarative styling for terminal output with a CSS-like mental model. The styling system implements:

1. **Color Palette**: A carefully chosen palette that works in both light and dark terminal themes. Primary colors for brand identity, accent colors for interactive elements, semantic colors for success/error/warning states.

2. **Typography**: While terminals have limited typography options, strategic use of bold, underline, and color creates hierarchy. Title bars use bold bright colors. Body text uses standard intensity. Subdued text uses dim colors.

3. **Layout System**: Lip Gloss's layout primitives (JoinHorizontal, JoinVertical, Place) enable complex layouts without manual spacing calculations. This is critical for responsive design.

4. **Animation System**: Subtle animations enhance perceived polish—spinning indicators during operations, fade-in effects for new content, transition animations between wizard steps. These are implemented through timer-based update cycles.

**Accessibility Considerations**:

While terminal UI accessibility is limited compared to web interfaces, several practices improve usability:

1. **Keyboard-Only Navigation**: All interactions are keyboard-accessible with clear focus indicators.
2. **Screen Reader Compatibility**: Output plain text alternatives for complex visual elements when requested.
3. **Color Independence**: Never rely solely on color to convey information; use symbols and text labels as well.
4. **High Contrast Mode**: Option to disable colors and use high-contrast text formatting.

### Configuration Management

Configuration management uses Viper, the standard configuration library for Go applications. Viper provides a unified interface across multiple configuration sources with a defined precedence hierarchy.

**Configuration Hierarchy** (highest to lowest precedence):

1. **Command-line flags**: Explicit user overrides for specific settings
2. **Environment variables**: Infrastructure-level configuration for CI/CD
3. **Project configuration file**: `.scaffold/config.yaml` in the project directory
4. **User configuration file**: `~/.config/scaffold/config.yaml` for user preferences
5. **Default values**: Built-in defaults for all settings

**Configuration Schema**:

```yaml
# User/Project configuration structure
scaffold:
  defaults:
    frontend:
      framework: react
      styling: tailwind
      state: zustand
    backend:
      framework: fastapi
      database: postgresql
    ai:
      context_verbosity: detailed
      brainstorm_enabled: true
    preferences:
      package_manager: auto-detect
      ide: cursor
      
  templates:
    registry: https://registry.scaffoldai.dev
    cache_dir: ~/.cache/scaffold/templates
    
  updates:
    check_enabled: true
    channel: stable
```

**Configuration Migration Strategy**:

As Scaffold AI evolves, configuration schemas will change. The configuration system includes migration logic to automatically update older configuration formats to current versions, preserving user preferences while adapting to new capabilities.

### Self-Update System

The self-update system enables Scaffold AI to update itself without requiring users to manually download new versions. This is implemented through the `go-selfupdate` library pattern, which:

1. **Version Checking**: Queries GitHub releases API to check for newer versions
2. **Binary Download**: Downloads the appropriate binary for the current OS/architecture
3. **Verification**: Validates checksums and signatures before applying update
4. **Atomic Replacement**: Replaces the running binary atomically to avoid corruption
5. **Restart Prompt**: Optionally prompts user to restart to use the new version

**Update Channels**:

- **stable**: Production-ready releases (default)
- **beta**: Pre-release versions for early adopters
- **nightly**: Bleeding-edge builds for contributors

**Security Considerations**:

- All releases are signed with a private key; the CLI verifies signatures before applying updates
- HTTPS-only communication with release servers
- Checksum verification for downloaded binaries
- Rate limiting on update checks to avoid API abuse

## Core Engine Layer

### Template Engine

The template engine is responsible for loading, processing, and rendering project templates. It must balance flexibility (supporting diverse project types) with consistency (ensuring all templates follow standards).

**Template Definition Structure**:

```
template-name/
├── template.yaml          # Template metadata and configuration
├── files/                 # Template files with variable substitution
│   ├── src/
│   ├── tests/
│   └── ...
├── hooks/
│   ├── pre-generate.sh    # Run before file generation
│   └── post-generate.sh   # Run after file generation
├── prompts.yaml           # Interactive prompts for this template
└── context-templates/     # AI context templates
    ├── system-prompt.md.tpl
    ├── architecture.md.tpl
    └── ...
```

**Template Processing Pipeline**:

1. **Discovery**: Identify template based on user selections or explicit specification
2. **Loading**: Parse template.yaml and validate structure
3. **Prompt Processing**: Collect user input through prompts.yaml definitions
4. **Variable Resolution**: Build complete variable context from prompts, defaults, and CLI flags
5. **File Processing**: Apply template processing to all files:
   - Variable substitution using Go template syntax
   - Conditional file inclusion based on variables
   - File permission preservation
   - Binary file handling (no template processing)
6. **Hook Execution**: Run pre/post generation hooks if present
7. **Validation**: Verify generated output meets quality standards

**Variable Substitution Engine**:

The engine uses Go's text/template with sprig extensions for rich template functionality:

- **Simple Substitution**: `{{ .ProjectName }}`
- **Conditionals**: `{{ if .Frontend.Enabled }}frontend content{{ end }}`
- **Loops**: `{{ range .Components }}{{ .Name }}{{ end }}`
- **Helpers**: `{{ .ProjectName | kebabcase }}`, `{{ now | date "2006" }}`

**Conditional File Inclusion**:

Files can be conditionally included based on user selections using a simple naming convention or explicit configuration:

```
files/
├── shared/              # Always included
├── frontend+react/      # Included if frontend is react
├── backend+fastapi/     # Included if backend is fastapi
└── database+postgres/   # Included if database is postgres
```

### Context Builder

The context builder is the subsystem that creates the AI context files—the heart of Scaffold AI's differentiation. It generates structured, AI-readable guidance that persists throughout the project lifecycle.

**Context File Architecture**:

```
.ai-context/
├── system-prompt.md           # Core behavioral instructions
├── architecture/
│   ├── decisions.md           # Architecture Decision Records
│   ├── patterns.md            # Approved design patterns
│   └── constraints.md         # What NOT to do
├── components/
│   ├── registry.md            # Catalog of reusable components
│   └── guidelines.md          # Component creation guidelines
├── operations/
│   ├── philosophy.md          # DevOps/SRE approach
│   ├── monitoring.md          # Logging and metrics standards
│   └── error-handling.md      # Error handling patterns
├── documentation/
│   ├── standards.md           # Documentation requirements
│   └── templates/             # Documentation templates
├── brainstorm.md              # AI scratchpad for reasoning
└── brand/
    ├── theme.md               # Visual theme guidelines
    └── voice.md               # Content style guidelines
```

**System Prompt Generation Strategy**:

The system prompt is generated through a layered composition approach:

**Layer 1 - Base Instructions**: Fundamental behavioral rules that apply to all projects:
- Code quality standards
- Documentation requirements
- Testing expectations
- General best practices

**Layer 2 - Technology-Specific Guidance**: Instructions derived from selected technologies:
- Framework-specific patterns
- Library usage conventions
- Technology-specific gotchas

**Layer 3 - Project-Specific Configuration**: Instructions derived from user choices:
- Naming conventions
- Folder structure expectations
- Integration patterns

**Layer 4 - User Extensions**: Custom instructions provided by the user:
- Organizational standards
- Team preferences
- Project-specific requirements

The final system prompt merges these layers with clear section headers and priority indicators, enabling AI tools to understand which instructions are absolute constraints versus preferences.

**Brainstorm.md Design**:

The Brainstorm.md file is perhaps the most innovative aspect of the context system. It serves as an externalized reasoning space for AI tools, enabling them to:

1. **Document Confusion**: When an AI encounters ambiguity, it writes the question to Brainstorm.md rather than interrupting the user
2. **Propose Solutions**: The AI can propose multiple solutions and reason through tradeoffs in the file
3. **Make Decisions**: The AI documents its decision and reasoning, creating an audit trail
4. **Continue Execution**: With the decision documented, the AI proceeds without waiting
5. **Enable Review**: Users can later review Brainstorm.md to understand AI reasoning and correct if needed

The Brainstorm.md template includes sections for:
- Active questions requiring resolution
- Resolved questions with reasoning
- Design decisions made
- Future considerations
- Links to relevant context

**Context Validation**:

The context builder includes validation logic to ensure generated context is:
- Syntactically correct (valid markdown, properly formatted)
- Semantically consistent (no contradictory instructions)
- Complete (all referenced files exist)
- Up-to-date (synchronized with project structure)

### File Generator

The file generator subsystem handles the actual creation of files and directories on disk. It must handle:

1. **Directory Creation**: Creating nested directory structures with appropriate permissions
2. **File Writing**: Writing processed template content to files
3. **Permission Preservation**: Maintaining executable permissions for scripts
4. **Conflict Handling**: What to do when files already exist
5. **Atomic Operations**: Ensuring generation either completes fully or rolls back

**Conflict Resolution Strategies**:

- **fail**: Stop generation if any conflict exists (safest, default for new projects)
- **skip**: Skip conflicting files, leave existing files unchanged
- **overwrite**: Replace existing files with generated content
- **merge**: Attempt to merge generated content with existing (complex, template-specific)

**Rollback Support**:

For complex projects with many files, the generator tracks all created/modified files. If generation fails partway, it can roll back to the initial state, avoiding partial project generation that leaves the directory in an inconsistent state.

### Plugin System

While the initial release will have a monolithic architecture, the design anticipates a future plugin system for extensibility. The plugin system would enable:

1. **Custom Template Sources**: Load templates from private registries, git repositories, or local directories
2. **Language Extensions**: Add support for languages not covered by core templates
3. **Framework Integration**: Deep integration with specific frameworks (e.g., Next.js, Django)
4. **AI Tool Integration**: Custom context formats for specific AI tools
5. **Organizational Templates**: Enterprise-specific template libraries

**Plugin Architecture Design**:

```go
// Plugin interface definition
type Plugin interface {
    // Metadata returns plugin information
    Metadata() PluginMetadata
    
    // Initialize is called when plugin is loaded
    Initialize(ctx context.Context, config PluginConfig) error
    
    // Hooks for extending functionality
    Hooks() []PluginHook
}

type PluginHook struct {
    Name        string
    Trigger     string  // "pre-init", "post-init", "pre-generate", etc.
    Handler     HookHandler
    Priority    int
}
```

The plugin system would use hashicorp/go-plugin for secure, out-of-process plugin execution, enabling plugins written in any language while maintaining isolation.

---

# AI Prompt Engineering System

## Context Engineering Philosophy

The emergence of AI coding assistants has created a new discipline: context engineering. Just as software engineering evolved from programming to encompass architecture, testing, and maintenance, working with AI assistants requires thoughtful curation of the context they receive.

Context engineering recognizes that AI models are powerful but context-dependent. The same model can produce brilliant or mediocre output depending entirely on the context provided. A model that receives clear instructions, relevant examples, and well-defined constraints will consistently outperform one that must infer everything from scratch.

This philosophy shapes every aspect of Scaffold AI's context system:

1. **Explicit Over Implicit**: Rather than hoping AI tools infer conventions from code, explicitly document expectations
2. **Structured Over Unstructured**: Use consistent formats that AI tools can reliably parse
3. **Living Over Static**: Context should evolve with the project, not fossilize at creation
4. **Layered Over Flat**: Organize context hierarchically to enable efficient retrieval
5. **Model-Agnostic Over Specific**: Design context that works across AI tools and models

## System Prompt Architecture

The system prompt is the most critical context artifact. It sets the behavioral foundation for all AI interactions with the project. Scaffold AI generates system prompts through a sophisticated composition system.

### Prompt Layer Composition

**Layer 0: Identity and Purpose**

Every system prompt begins with identity establishment—telling the AI what role it should assume:

```markdown
# AI Development Assistant for [PROJECT_NAME]

You are an expert software developer working on the [PROJECT_NAME] project.
Your role is to write clean, maintainable, well-documented code that follows
the architectural patterns and conventions established for this project.

## Critical Instruction

Before generating ANY code or making ANY changes, you MUST:
1. Read this entire document
2. Read all files in the .ai-context/ directory
3. Review the Brainstorm.md file for any outstanding questions or decisions
4. Check the Component Registry before creating new components

This project uses AI-assisted development workflows. The context files in
.ai-context/ are not optional documentation—they are behavioral contracts
that you must follow.
```

This layer establishes the AI's identity and, critically, forces attention to the context system. Without explicit instruction to read context files, AI tools often ignore them.

**Layer 1: Project Architecture Overview**

```markdown
## Project Architecture

[PROJECT_NAME] is a [PROJECT_TYPE] built with:

**Frontend**: [FRONTEND_FRAMEWORK] with [STYLING_APPROACH] styling
**Backend**: [BACKEND_FRAMEWORK] with [DATABASE] database
**Infrastructure**: [INFRASTRUCTURE_TYPE]

### Key Architectural Decisions

[Brief summary of major architectural decisions with links to ADRs]

### Critical Constraints

- Do NOT add dependencies without checking approved libraries list
- Do NOT create new architectural patterns without documenting in ADRs
- Do NOT bypass established error handling patterns
- Do NOT hardcode configuration values
```

This layer provides high-level orientation, helping AI tools understand the project's shape before diving into specifics.

**Layer 2: Technology-Specific Guidelines**

Each selected technology contributes its guidelines:

```markdown
## [TECHNOLOGY_NAME] Guidelines

### Patterns to Follow
[Technology-specific patterns and best practices]

### Patterns to Avoid
[Anti-patterns specific to this technology]

### Dependency Management
[How to handle dependencies for this technology]

### Common Gotchas
[Technology-specific pitfalls and how to avoid them]
```

These guidelines are curated from official documentation, community best practices, and hard-won experience. They're not generic technology introductions but practical guidance for using each technology effectively within the project context.

**Layer 3: Code Style and Documentation**

```markdown
## Code Style Requirements

### Documentation Standards

Every function MUST include documentation that explains:
1. Purpose: What the function does and why it exists
2. Parameters: What each parameter represents and constraints
3. Returns: What the function returns and in what format
4. Errors: What errors might be thrown and when
5. Side Effects: Any side effects the function has

For complex logic, include inline comments explaining the reasoning:
- Why this approach was chosen
- What alternatives were considered
- Any performance or correctness considerations

### Naming Conventions

- Components: PascalCase with descriptive names
- Functions: camelCase starting with verb
- Constants: SCREAMING_SNAKE_CASE
- Files: kebab-case matching primary export
```

**Layer 4: Component and Pattern Registry**

```markdown
## Component Registry

Before creating any new component, check the existing registry:
[Link to component registry]

### Reusability Requirements

- All components must be designed for reuse
- Props must be typed and documented
- Components must not contain business logic
- Styling must use theme variables, not hardcoded values
```

**Layer 5: Operational Concerns**

```markdown
## Operational Philosophy

### Logging Standards

All code must include appropriate logging:
- Entry/exit logging for significant functions
- Error logging with context
- Performance logging for slow operations

### Error Handling

Follow the established error handling patterns:
[Link to error handling documentation]

Never:
- Swallow errors silently
- Return raw errors without context
- Use panics for expected error conditions
```

### Prompt Merge Strategy

When multiple layers contribute to the final prompt, the merge strategy ensures:

1. **No Conflicts**: Later layers cannot contradict earlier layers without explicit override
2. **Proper Ordering**: Foundational instructions come before specific ones
3. **Clear Sections**: Each concern has its own section for easy scanning
4. **Reasonable Length**: Total prompt length stays within context window limits

The merge process also injects project-specific values (project name, selected technologies, user preferences) into template slots, creating a customized prompt from standardized templates.

## Brainstorm.md Lifecycle

The Brainstorm.md file is designed to support autonomous AI operation while maintaining human oversight. Its lifecycle reflects the rhythm of AI-assisted development.

### Initial State

At project creation, Brainstorm.md contains a template with empty sections:

```markdown
# Brainstorm.md

This file serves as an external reasoning space for AI assistants.
When encountering ambiguity, write questions here, propose solutions,
and document decisions. Users can review this file to understand
AI reasoning and provide corrections.

---

## Active Questions

_Questions the AI is currently trying to resolve_

(No active questions)

---

## Resolved Decisions

_Decisions the AI has made, with reasoning_

(No decisions yet)

---

## Design Rationale

_Document the reasoning behind significant design choices_

(No entries yet)

---

## Future Considerations

_Things to revisit or improve later_

(No entries yet)
```

### Question Resolution Cycle

When an AI encounters ambiguity, it follows this cycle:

1. **Identify Ambiguity**: The AI recognizes it needs information not available in context
2. **Document Question**: Write the question to "Active Questions" with context
3. **Propose Options**: List possible approaches with tradeoffs
4. **Select Approach**: Choose the most reasonable option given available information
5. **Document Decision**: Move to "Resolved Decisions" with reasoning
6. **Continue Execution**: Proceed with implementation

Example Brainstorm.md entry:

```markdown
## Active Questions

### How should user authentication be implemented?

Context: The project requires user authentication but no specific method
was specified during scaffolding.

Options:
1. JWT-based authentication (stateless, good for APIs)
2. Session-based authentication (simpler, requires session storage)
3. OAuth integration (delegated to providers, more complex)

Analysis:
- The project uses FastAPI backend with PostgreSQL
- Session storage adds database complexity
- JWT aligns with RESTful API patterns
- No OAuth providers specified in requirements

Decision: Implementing JWT-based authentication with refresh tokens.
This can be revisited if requirements change.

---

## Resolved Decisions

### 2024-01-15: Authentication Approach
**Question**: How should user authentication be implemented?
**Decision**: JWT-based authentication with refresh tokens
**Reasoning**: Aligns with RESTful API patterns, works well with
PostgreSQL for token invalidation, no OAuth requirements specified.
```

### Human Review Integration

Users can review Brainstorm.md at any time to:
- See what decisions the AI has made
- Correct decisions that were wrong
- Provide guidance for open questions
- Update preferences based on new understanding

When a user makes changes to Brainstorm.md, the AI should:
- Acknowledge the changes
- Update its behavior to match new guidance
- Document the correction in Resolved Decisions

## AI Compatibility Abstraction Layer

Different AI tools have different capabilities and expectations. Scaffold AI's context system includes an abstraction layer that adapts context to different tools while maintaining semantic consistency.

### Tool-Specific Outputs

**For Cursor (.cursorrules format)**:

```
---
description: Project AI Guidelines
globs: ["**/*"]
---

[CORE SYSTEM PROMPT CONTENT]
```

**For Continue.dev (.continuerc format)**:

```json
{
  "systemPrompt": "[CORE SYSTEM PROMPT CONTENT]",
  "contextProviders": [
    {
      "name": "ai-context",
      "params": {
        "contextDir": ".ai-context"
      }
    }
  ]
}
```

**For Claude Code (CLAUDE.md format)**:

```markdown
# Project Context for Claude Code

[CORE SYSTEM PROMPT CONTENT]

## Additional Context Files

Please read these files for complete project context:
- .ai-context/architecture/decisions.md
- .ai-context/components/registry.md
- .ai-context/operations/philosophy.md
```

### Semantic Preservation

While output formats differ, the abstraction layer ensures semantic equivalence:

1. **All instructions preserved**: Nothing is lost in translation
2. **Priority maintained**: Constraints remain constraints, preferences remain preferences
3. **Links work**: References between files resolve correctly
4. **Updates synchronized**: Changes to core context propagate to all formats

## Preventing AI Drift

One of the key challenges in AI-assisted development is drift—the gradual degradation of codebase coherence as AI tools make decisions that are locally optimal but globally inconsistent. Scaffold AI's context system includes specific mechanisms to prevent drift.

### Constraint Enforcement

Hard constraints are distinguished from preferences:

```markdown
## HARD CONSTRAINTS (MUST follow)

These rules are non-negotiable. Violating them requires explicit
override from a human developer with documented justification.

1. All new dependencies must be added to approved-dependencies.md
2. Database queries must use the query builder, never raw SQL
3. API endpoints must follow the established error response format
4. Frontend components must use the theme system, no hardcoded colors

## PREFERENCES (SHOULD follow)

These guidelines represent best practices but can be deviated from
when there's good reason.

1. Prefer functional components over class components
2. Prefer composition over inheritance
3. Prefer explicit returns over implicit
```

AI tools are instructed to treat hard constraints as inviolable and to document any deviations from preferences.

### Dependency Control

The context system includes an approved dependencies list:

```markdown
## Approved Dependencies

### Frontend
| Package | Version | Purpose | Alternatives Considered |
|---------|---------|---------|------------------------|
| react | ^18.2 | UI framework | Vue, Svelte |
| zustand | ^4.4 | State management | Redux, Jotai |
| tanstack-query | ^5.0 | Data fetching | SWR, Apollo |

### Backend
| Package | Version | Purpose | Alternatives Considered |
|---------|---------|---------|------------------------|
| fastapi | ^0.104 | API framework | Flask, Express |
| sqlalchemy | ^2.0 | ORM | Django ORM, Prisma |

### Adding New Dependencies

New dependencies require:
1. Justification in ADR format
2. Security review
3. Bundle size impact analysis (frontend)
4. License compatibility check
```

This prevents the common AI behavior of suggesting new libraries for every task, which can lead to dependency bloat and inconsistency.

### Consistency Checking

The `scaffold context check` command validates AI-generated code against context rules:

- Import analysis: Are all imports from approved sources?
- Pattern matching: Does code follow established patterns?
- Documentation check: Are new functions documented?
- Component uniqueness: Does a similar component already exist?

This doesn't replace code review but provides an automated first pass at catching drift.

---

# UX Strategy

## Design Philosophy

The user experience of a CLI tool might seem like an afterthought—after all, command-line interfaces are by definition text-based. However, in an era of beautiful graphical interfaces and polished developer experiences, CLI tools have evolved significantly. Tools like `gh` (GitHub CLI), `vercel`, and `railway` have demonstrated that terminal applications can be both powerful and delightful.

Scaffold AI's UX philosophy centers on three principles:

1. **Progressive Disclosure**: Complexity should reveal itself gradually. New users can run `scaffold init` and be guided through every decision. Power users can skip the wizard entirely with flags and config files.

2. **Confidence Through Clarity**: Every step should make the user feel confident about what's happening. Clear progress indicators, explanatory text, and confirmation prompts prevent the anxiety of "what is this thing doing to my filesystem?"

3. **Delight in Details**: Small moments of polish—a smooth spinner animation, a well-formatted success message, ASCII art that doesn't feel childish—create an emotional connection that makes the tool feel professional and trustworthy.

## Terminal UI Implementation

### Visual Language

**Color System**:

The color palette is designed to work across terminal themes (light, dark, and various color schemes):

```
Primary Brand:    #6C63FF (Vibrant purple - modern, distinctive)
Primary Dark:     #5549E8 (Darker variant for contrast)
Primary Light:    #8B85FF (Lighter variant for backgrounds)

Success:          #10B981 (Green - universally positive)
Warning:          #F59E0B (Amber - attention without alarm)
Error:            #EF4444 (Red - clear but not harsh)
Info:             #3B82F6 (Blue - neutral, informational)

Text Primary:     Terminal default (adapts to theme)
Text Secondary:   Dim variant
Text Accent:      Bold primary color
```

For terminals that don't support true color, the system gracefully degrades to 256-color or 16-color approximations.

**Typography Strategy**:

Terminal typography is limited, but strategic use of available options creates hierarchy:

- **Titles**: Bold + Bright + Primary Color
- **Section Headers**: Bold + Primary Color
- **Option Labels**: Bold
- **Descriptions**: Normal intensity
- **Subdued Text**: Dim
- **Code/Commands**: Italic or surrounded by backticks in display

**Spacing and Layout**:

Generous spacing prevents the cramped feeling of many CLI tools:
- Double line breaks between major sections
- Single line break between options
- Indentation for hierarchical relationships
- Consistent padding around borders

### Wizard Flow

The interactive wizard is the heart of the `init` experience. Its flow is carefully designed to minimize cognitive load while collecting all necessary information.

**Step 1: Welcome and Project Naming**

```
╔═══════════════════════════════════════════════════════════════════════╗
║                                                                       ║
║   ███████╗ ██████╗ █████╗ ██╗     ███████╗███████╗██████╗ ██╗        ║
║   ██╔════╝██╔════╝██╔══██╗██║     ██╔════╝██╔════╝██╔══██╗██║        ║
║   ███████╗██║     ███████║██║     ███████╗█████╗  ██████╔╝██║        ║
║   ╚════██║██║     ██╔══██║██║     ╚════██║██╔══╝  ██╔══██╗██║        ║
║   ███████║╚██████╗██║  ██║███████╗███████║███████╗██║  ██║███████╗   ║
║   ╚══════╝ ╚═════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝╚═╝  ╚═╝╚══════╝   ║
║                                                                       ║
║                    AI-Ready Project Scaffolding                       ║
║                                                                       ║
╚═══════════════════════════════════════════════════════════════════════╝

Welcome to Scaffold AI! Let's create your AI-ready project.

? What would you like to name your project? (my-awesome-project)
```

The ASCII art logo establishes brand identity. The welcome message is brief. The first question is simple but sets expectations.

**Step 2: Project Type Selection**

```
┌─────────────────────────────────────────────────────────────────────┐
│  Project Type                                                        │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  What type of project are you building?                             │
│                                                                      │
│    ◉ Full-Stack Web Application                                     │
│      Frontend Only (SPA/SSR)                                        │
│      Backend API Service                                            │
│      CLI Application                                                │
│      Library/Package                                                │
│      Data Pipeline/ETL                                              │
│                                                                      │
│  ─────────────────────────────────────────────────────────────────  │
│  Full-Stack Web Application includes both frontend and backend     │
│  with database integration, authentication, and deployment config.  │
└─────────────────────────────────────────────────────────────────────┘
```

Options are presented with radio buttons (single select). A description panel at the bottom explains the currently selected option. Keyboard navigation (arrows, j/k, tab) moves between options.

**Step 3: Frontend Configuration** (if applicable)

```
┌─────────────────────────────────────────────────────────────────────┐
│  Frontend Stack                                                      │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  ? Framework?                                                        │
│                                                                      │
│    ◉ React (with Vite)      ○ Angular        ○ Svelte               │
│      Vue (with Vite)          Next.js          Nuxt                  │
│      Solid                    Astro            Remix                 │
│                                                                      │
│  ─────────────────────────────────────────────────────────────────  │
│  React with Vite offers excellent DX, huge ecosystem, and fast     │
│  builds. Ideal for SPAs and compatible with many hosting options.   │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  Frontend Details                                                    │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  ? Styling approach?        ? State management?                     │
│                                                                      │
│    ◉ Tailwind CSS            ◉ Zustand                              │
│      CSS Modules               Redux Toolkit                        │
│      Styled Components         Jotai                                │
│      Vanilla Extract           React Query + Context                │
│                                                                      │
│  ? Component library?       ? Form handling?                        │
│                                                                      │
│    ◉ shadcn/ui              ◉ React Hook Form                       │
│      MUI                       Formik                               │
│      Chakra UI                 Native                               │
│      Headless UI                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

Multiple related questions can be displayed simultaneously using a split layout. This speeds up the process for experienced users while still providing guidance.

**Step 4: Backend Configuration** (if applicable)

```
┌─────────────────────────────────────────────────────────────────────┐
│  Backend Stack                                                       │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  ? Framework?                                                        │
│                                                                      │
│    ◉ FastAPI (Python)       ○ Express (Node)    ○ Go Fiber          │
│      Django                   NestJS              Gin                │
│      Flask                    Fastify             Echo               │
│                                                                      │
│  ─────────────────────────────────────────────────────────────────  │
│  FastAPI provides automatic OpenAPI docs, type safety with Pydantic,│
│  and excellent async support. Great for modern API development.     │
└─────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────┐
│  Backend Details                                                     │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  ? Database?                ? ORM/Query Tool?                       │
│                                                                      │
│    ◉ PostgreSQL             ◉ SQLAlchemy 2.0                        │
│      MySQL                    Prisma                                │
│      SQLite                   Django ORM                            │
│      MongoDB                  Tortoise ORM                          │
│                                                                      │
│  ? Cache?                   ? Task Queue?                           │
│                                                                      │
│    ◉ Redis                  ○ Celery                                │
│      None                     None                                  │
│                                                                      │
│  ? Environment management?                                           │
│                                                                      │
│    ◉ uv (recommended)       ○ poetry       ○ pipenv                 │
│      pip + venv                                                │
└─────────────────────────────────────────────────────────────────────┘
```

**Step 5: AI Context Configuration**

```
┌─────────────────────────────────────────────────────────────────────┐
│  AI Development Context                                              │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  Scaffold AI creates context files to guide AI assistants like      │
│  Cursor, Claude Code, and GitHub Copilot.                           │
│                                                                      │
│  ? Primary AI assistant?                                             │
│                                                                      │
│    ◉ Any/AI-agnostic        ○ Cursor        ○ Continue.dev          │
│      Claude Code              GitHub Copilot                        │
│                                                                      │
│  ? Context verbosity?                                                │
│                                                                      │
│    ◉ Comprehensive          ○ Standard      ○ Minimal               │
│                                                                      │
│  ? Enable Brainstorm.md (AI self-reasoning)?                        │
│                                                                      │
│    ◉ Yes (recommended)      ○ No                                   │
│                                                                      │
│  ─────────────────────────────────────────────────────────────────  │
│  AI-agnostic context works with any AI tool. Comprehensive verbosity│
│  includes detailed patterns, examples, and explanations.            │
└─────────────────────────────────────────────────────────────────────┘
```

This step explains the AI context feature and configures it appropriately.

**Step 6: Review and Generate**

```
┌─────────────────────────────────────────────────────────────────────┐
│  Review Configuration                                                │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  Project: my-awesome-project                                        │
│  Type: Full-Stack Web Application                                   │
│                                                                      │
│  Frontend:                                                          │
│    • Framework: React + Vite                                        │
│    • Styling: Tailwind CSS + shadcn/ui                             │
│    • State: Zustand + React Query                                   │
│    • Forms: React Hook Form                                         │
│                                                                      │
│  Backend:                                                           │
│    • Framework: FastAPI                                             │
│    • Database: PostgreSQL + SQLAlchemy                              │
│    • Cache: Redis                                                   │
│    • Environment: uv                                                │
│                                                                      │
│  AI Context:                                                        │
│    • Format: AI-agnostic                                            │
│    • Verbosity: Comprehensive                                       │
│    • Brainstorm.md: Enabled                                         │
│                                                                      │
│  ─────────────────────────────────────────────────────────────────  │
│                                                                      │
│  ? Generate project with these settings?                            │
│                                                                      │
│    ◉ Yes, generate project                                          │
│      Modify configuration                                           │
│      Save as template                                               │
│      Cancel                                                         │
└─────────────────────────────────────────────────────────────────────┘
```

A final review gives users confidence before the tool makes any changes.

**Step 7: Generation Progress**

```
┌─────────────────────────────────────────────────────────────────────┐
│  Generating Project                                                  │
├─────────────────────────────────────────────────────────────────────┤
│                                                                      │
│  ✓ Creating directory structure                                     │
│  ✓ Generating frontend files                                        │
│  ✓ Generating backend files                                         │
│  ◐ Building AI context files                                        │
│  ○ Installing dependencies                                          │
│  ○ Finalizing configuration                                         │
│                                                                      │
│  ─────────────────────────────────────────────────────────────────  │
│  Writing system prompt with 47 technology-specific rules...         │
└─────────────────────────────────────────────────────────────────────┘
```

Progress is shown in real-time with checkmarks for completed items and spinners for in-progress items.

**Step 8: Success and Next Steps**

```
╔═══════════════════════════════════════════════════════════════════════╗
║                                                                       ║
║                        ✓ Project Created!                            ║
║                                                                       ║
╚═══════════════════════════════════════════════════════════════════════╝

  my-awesome-project/ has been created with:
  
  • Frontend: React + Vite + Tailwind + shadcn/ui
  • Backend: FastAPI + PostgreSQL + SQLAlchemy
  • AI Context: Comprehensive, AI-agnostic format
  • Brainstorm.md: Enabled for autonomous AI reasoning

  Next steps:
  
    cd my-awesome-project
    
    # Frontend
    cd frontend && npm install && npm run dev
    
    # Backend  
    cd backend && uv sync && uv run uvicorn app.main:dev
    
  AI Context Files:
    .ai-context/system-prompt.md - Core behavioral instructions
    .ai-context/brainstorm.md    - AI reasoning scratchpad
    .ai-context/                 - Full context directory
    
  Documentation:
    README.md                    - Project overview and setup
    .ai-context/architecture/    - Architecture decision records
```

### Responsive Behavior

The wizard adapts to terminal width:

**Wide Terminal (>100 columns)**:
- Side-by-side option groups
- Description panel always visible
- Full ASCII art logo

**Medium Terminal (80-100 columns)**:
- Stacked option groups
- Description panel toggleable
- Simplified logo

**Narrow Terminal (60-80 columns)**:
- Single column options
- Descriptions on demand (press 'd')
- Text-only header

**Minimum Terminal (<60 columns)**:
- Warning message requesting resize
- Option to continue with minimal UI

### Accessibility Features

1. **Keyboard Navigation**: All interactions work via keyboard
   - Arrow keys or j/k for list navigation
   - Tab between sections
   - Enter to select
   - Escape to go back
   - '?' for help

2. **Screen Reader Support**: 
   - Semantic structure for screen readers
   - Status announcements for changes
   - Plain text mode option

3. **Color Independence**:
   - Never rely solely on color
   - Use symbols (◉ ○ ✓ ✗) alongside colors
   - High contrast mode option

4. **Motion Sensitivity**:
   - Animations can be disabled
   - No flashing or rapid movement

## Command Line Interface

For users who prefer flags over wizards or need to integrate with CI/CD:

```bash
# Minimal with all defaults
scaffold init my-project --defaults

# Full specification
scaffold init my-project \
  --type fullstack \
  --frontend react \
  --frontend-styling tailwind \
  --frontend-state zustand \
  --backend fastapi \
  --database postgres \
  --orm sqlalchemy \
  --ai-context comprehensive \
  --non-interactive

# From saved configuration
scaffold init my-project --config .scaffold-template.yaml

# Using a template
scaffold init my-project --template saas-starter

# Extending user's custom prompt
scaffold init my-project --user-prompt ./my-company-guidelines.md
```

### User Prompt Append Strategy

A key feature is the ability to append custom organizational guidelines to the generated system prompt:

```bash
scaffold init my-project --user-prompt ./company-standards.md
```

The user's prompt is appended to the base Scaffold AI prompt:

```
[SCAFFOLD AI BASE PROMPT]

---

# Organization-Specific Guidelines

[USER'S CUSTOM PROMPT CONTENT]

---

Continue following the guidelines above while respecting all constraints
defined in the base prompt.
```

This enables organizations to layer their standards on top of Scaffold AI's comprehensive base.

---

# Branding Strategy

## Name Exploration

The name must work across multiple contexts:
- As a CLI command (short, memorable, easy to type)
- As a brand (distinctive, ownable, not generic)
- As a noun (when discussing the tool)
- As a verb (when describing the action)

### Primary Recommendation: Scaffold AI

**Command**: `scaffold`

**Rationale**:
- "Scaffold" directly describes what the tool does—creating scaffolding for projects
- Familiar term in development context (scaffolding is a well-understood concept)
- Works naturally as a verb ("let me scaffold a new project")
- ".ai" extension positions it as an AI-era tool
- Clear, professional, not trying too hard

**Concerns**:
- "Scaffold" is a common word; SEO may be challenging initially
- Some might expect traditional scaffolding without AI context features

**Mitigation**: The differentiation comes through in the positioning—this isn't just scaffolding, it's AI-ready scaffolding with persistent context.

### Alternative Options

**Clause AI** (`Clause`)
- Evokes creation, craftsmanship, transformation
- Strong verb potential ("Clause a new project")
- Historic/industrial vibe suggests solidity
- Slightly more abstract than "scaffold"
- Risk: "Clause" is used by other dev tools (Forgit, Clause for Minecraft, etc.)

**Blueprint AI** (`blueprint` or `bp`)
- Emphasizes planning and architectural guidance
- Works well with the context/decision documentation aspect
- "Blueprint" is longer to type; "bp" is terse but less memorable
- Positive associations with construction and intentional design

**Architect AI** (`arch`)
- Emphasizes the architectural guidance aspect
- Short command ("arch init")
- May imply the tool does more design than it actually does
- "Arch" has other associations (arch Linux, architecture in general)

**Seed AI** (`seed`)
- Evokes growth, beginnings, potential
- Very short command
- Organic metaphor may not fit all users
- "Seed" has database seeding connotations that could confuse

**Foundry AI** (`foundry`)
- Similar to Clause—evokes creation and craft
- Less common than Clause in dev tools space
- Slightly longer command
- Industrial associations suggest reliability

## Visual Identity Direction

### Logo Concept

The logo should combine:
1. **Structural element**: Representing scaffolding/architecture
2. **AI element**: Representing intelligent guidance
3. **Terminal context**: Acknowledging the CLI nature

**Concept A**: ASCII-style scaffolding structure with a glowing "brain" or circuit pattern integrated
**Concept B**: Minimal geometric brackets `[ ]` forming a structure with an AI spark
**Concept C**: Terminal prompt `$` transforming into a building structure

### Color System

**Primary**: Deep purple (#6C63FF)
- Modern, distinctive, tech-forward
- Works in both light and dark contexts
- Not oversaturated; maintains professionalism

**Accent**: Electric blue (#3B82F6)
- Used for interactive elements and highlights
- Provides contrast with primary

**Neutral**: Cool grays
- Background, text, borders
- Avoid warm grays that might clash with purple

### Typography

**Logo**: Custom or modified monospace font
- Reinforces the developer/terminal identity
- Should be distinctive even in plain text

**Marketing**: Geometric sans-serif
- Clean, modern, readable
- Examples: Inter, Satoshi, Space Grotesk

**Terminal**: Terminal default
- Never force specific fonts in terminal output
- Use weight and color for hierarchy

## Tagline Exploration

**Primary Tagline**: "Build with AI from day one."

**Rationale**:
- Directly states the value proposition
- "From day one" emphasizes that AI context is built in, not retrofitted
- Active, forward-looking

**Alternative Taglines**:

- "Projects that AI understands." (Emphasizes the context system)
- "Scaffold once. AI forever." (Emphasizes persistence)
- "Where projects meet AI intelligence." (More corporate)
- "The last scaffolding tool you'll need." (Bold claim, potentially overpromising)

## Emotional Positioning

Scaffold AI should feel:

**Professional**: Not toy-like, not overly playful. Developers use it for serious work.

**Intelligent**: The AI context system is sophisticated; the brand should reflect that intelligence.

**Helpful**: The tool exists to reduce friction and improve developer experience.

**Trustworthy**: The tool makes significant changes to filesystems; users must trust it.

**Modern**: This is a tool for 2024 and beyond; it should feel contemporary.

Not:

**Gimmicky**: No excessive animations, no cute mascot, no forced "personality"

**Cold**: While professional, not robotic or impersonal

**Generic**: Avoiding the "startup blue" or "tech green" clichés

---

# Monetization Strategy

## Open Source Foundation

Scaffold AI is built on an open source foundation. The core scaffolding engine, basic templates, and AI context system are freely available under a permissive license (MIT or Apache 2.0). This approach:

1. **Enables adoption**: Developers can try the tool without commitment
2. **Builds community**: Open source encourages contribution and advocacy
3. **Establishes standard**: Ubiquity makes the AI context format a de facto standard
4. **Reduces friction**: No procurement process for individual developers

## Premium Tier: Scaffold AI Pro

**Price**: $19/month or $190/year (individual), $49/user/month (team)

**Features**:

1. **Template Marketplace Access**: Premium templates for specific use cases:
   - SaaS Starter (authentication, billing, dashboard)
   - E-commerce Platform (cart, checkout, inventory)
   - Data Dashboard (charts, filters, real-time updates)
   - Mobile Backend (push notifications, offline sync)
   - Enterprise API (rate limiting, versioning, documentation)

2. **Team Synchronization**: Cloud-based sharing of:
   - Organization-specific templates
   - Shared AI context configurations
   - Coding standards and conventions
   - Component libraries and patterns

3. **Advanced AI Context**:
   - Automatic context updates based on codebase changes
   - AI analytics dashboard (how AI interacts with context)
   - Context optimization suggestions
   - Multiple context profiles per project

4. **Priority Support**:
   - Faster response times on GitHub issues
   - Access to private Discord/Slack community
   - Feature request prioritization

## Enterprise Tier: Scaffold AI Enterprise

**Price**: Custom pricing (typically $20K-$100K/year based on team size)

**Features**:

1. **Self-Hosted Option**:
   - Deploy template registry on-premises
   - Air-gapped environment support
   - Custom branding and white-labeling

2. **Governance and Compliance**:
   - AI context audit logging
   - Approved technology enforcement
   - Security scanning integration
   - Compliance report generation

3. **Integration**:
   - SSO (SAML, OIDC)
   - IDE plugin for VS Code, JetBrains
   - CI/CD pipeline integration
   - Jira/Linear issue tracking integration

4. **Custom Template Development**:
   - Professional services for custom templates
   - Organization-specific technology stack support
   - Migration from existing project templates

5. **Support**:
   - Dedicated customer success manager
   - Custom training sessions
   - SLA guarantees
   - Phone support

## Additional Revenue Streams

### Template Marketplace

A marketplace for community-created templates with revenue sharing:

- Free templates: No cost, community contribution
- Paid templates: Creator sets price (typically $9-$49), 70/30 split
- Verified templates: Vetted by Scaffold AI team, premium pricing

This creates an ecosystem where template creators can monetize their expertise while expanding Scaffold AI's capabilities.

### AI Context Services

**Context Auditing**: Professional review of AI context files to improve effectiveness (one-time fee, $299/project)

**Context Migration**: Service to migrate existing projects to include AI context ($499-$1999 depending on complexity)

**Custom Context Development**: Bespoke AI context systems for unusual technology stacks (hourly consulting)

### Educational Content

**Course**: "Effective AI-Assisted Development with Scaffold AI" ($99)
- How to structure projects for AI understanding
- Best practices for context maintenance
- Advanced Brainstorm.md workflows
- Team collaboration patterns

**Book**: "The AI-Operable Codebase" ($39)
- Philosophy and methodology
- Case studies
- Architecture patterns
- Future of AI-assisted development

## Pricing Philosophy

### Individual Developer Pricing

The $19/month individual price point is intentional:
- Under the "lunch money" threshold where purchasing requires little thought
- Comparable to other developer tools (Cursor, GitHub Pro)
- Significant value for serious AI-assisted development

### Team Pricing

The $49/user/month team pricing:
- Encourages team adoption over individual
- Justified by collaboration features
- Competitive with team plans for similar tools

### Enterprise Pricing

Enterprise pricing is consultative because:
- Requirements vary enormously
- Procurement processes differ
- Value is tied to specific use cases

### Conversion Strategy

**Free to Pro**:
- Limit of 5 projects with free tier
- Premium templates visible but inaccessible
- Team sync features grayed out
- In-app messaging about upgrade benefits

**Pro to Enterprise**:
- Direct sales outreach for teams >25 users
- Custom demos and pilots
- Case studies and ROI documentation

---

# Distribution Strategy

## Installation Methods

### curl | bash (Unix/macOS)

The canonical installation method for developer tools on Unix-like systems:

```bash
curl -fsSL https://get.scaffoldai.dev | bash
```

**Implementation**:
- Shell script detects OS and architecture
- Downloads appropriate binary from GitHub Releases
- Verifies checksum
- Installs to `/usr/local/bin` or `~/.local/bin`
- Adds to PATH if necessary
- Handles updates via `scaffold update`

**Security Considerations**:
- Script is served over HTTPS
- Binary is signed; signature verified before installation
- Checksum verification prevents tampering
- Clear documentation of what the script does

### Homebrew (macOS/Linux)

```bash
brew install scaffold-ai/tap/scaffold
```

**Advantages**:
- Familiar installation method for macOS users
- Automatic updates via `brew upgrade`
- Dependency management handled by Homebrew
- Trusted by security-conscious users

**Implementation**:
- Homebrew formula in scaffold-ai tap
- Formula references GitHub Releases binaries
- Includes man page and shell completions

### Winget (Windows)

```powershell
winget install ScaffoldAI.Scaffold
```

**Advantages**:
- Native Windows package manager
- Pre-installed on Windows 11
- Automatic updates
- Enterprise deployment support

**Implementation**:
- Manifest in winget-pkgs repository
- MSI installer for professional Windows experience
- Code signing certificate required

### Chocolatey (Windows)

```powershell
choco install scaffold-ai
```

**Advantages**:
- Popular among Windows power users
- Enterprise-friendly
- Supports internal repositories

### Scoop (Windows)

```powershell
scoop bucket add scaffold-ai https://github.com/scaffold-ai/scoop-bucket
scoop install scaffold
```

**Advantages**:
- Popular with developers on Windows
- No admin privileges required
- Clean installations without polluting system

### APT (Debian/Ubuntu)

```bash
curl -fsSL https://apt.scaffoldai.dev/gpg | sudo gpg --dearmor -o /etc/apt/trusted.gpg.d/scaffold.gpg
echo "deb https://apt.scaffoldai.dev stable main" | sudo tee /etc/apt/sources.list.d/scaffold.list
sudo apt update
sudo apt install scaffold-ai
```

**Advantages**:
- Native package manager for Debian-based systems
- Automatic updates via apt
- Familiar to Linux administrators

### NPM (Cross-platform)

```bash
npm install -g @scaffold-ai/cli
```

**Implementation**:
- Node.js wrapper around native binaries
- Downloads appropriate binary during install
- Works anywhere npm works

**Advantages**:
- Familiar to JavaScript developers
- No compilation required
- Works in environments where npm is already available

## Release Pipeline

### Versioning Strategy

**Semantic Versioning**: MAJOR.MINOR.PATCH

- **MAJOR**: Breaking changes to CLI interface or generated output
- **MINOR**: New features, new templates, new context capabilities
- **PATCH**: Bug fixes, minor improvements

**Pre-release Channels**:
- `alpha`: Internal testing, unstable
- `beta`: Public testing, feature complete but potentially buggy
- `rc`: Release candidate, final testing before stable

### CI/CD Pipeline

**On Pull Request**:
1. Run linting and formatting checks
2. Run unit tests
3. Run integration tests
4. Build binaries for all platforms
5. Run binary size checks

**On Merge to Main**:
1. All PR checks
2. Build development binaries
3. Deploy documentation updates

**On Version Tag**:
1. All previous checks
2. Build release binaries for all platforms
3. Sign binaries with GPG key
4. Create GitHub Release with:
   - Release notes
   - Binary downloads
   - Checksums
   - Signature files
5. Update package managers:
   - Submit to Homebrew
   - Update winget manifest
   - Update Chocolatey package
   - Update apt repository
   - Publish npm package

### Cross-Compilation Strategy

Go's cross-compilation capabilities enable building for all platforms from a single machine:

**Target Platforms**:
- darwin/amd64 (macOS Intel)
- darwin/arm64 (macOS Apple Silicon)
- linux/amd64 (Linux x64)
- linux/arm64 (Linux ARM64)
- windows/amd64 (Windows x64)
- windows/arm64 (Windows ARM64)

**Build Command**:
```bash
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o scaffold-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o scaffold-darwin-arm64
# etc.
```

### Binary Size Optimization

Go binaries can be large, but several techniques reduce size:

1. **Strip Debug Symbols**: `-ldflags="-s -w"` strips symbol table and DWARF
2. **UPX Compression**: Optional compression reduces size by 50-70%
3. **Build Tags**: Exclude unnecessary features for specific builds
4. **Dependency Audit**: Avoid heavy dependencies when lighter alternatives exist

**Target Size**: <15MB uncompressed, <6MB with UPX

### Security Signing

All release binaries are signed:

**GPG Signing** (Linux/macOS):
- Release manager's GPG key signs binaries
- Public key published on website and in repository
- Users can verify with `gpg --verify`

**Code Signing** (Windows):
- Authenticode certificate from trusted CA
- Prevents SmartScreen warnings
- Required for winget acceptance

**Notarization** (macOS):
- Apple notarization for Gatekeeper approval
- Required for distribution outside App Store
- Prevents "unverified developer" warnings

---

# Risk Analysis

## Market Risks

### Risk: AI Tool Evolution

**Description**: AI coding assistants are evolving rapidly. Features we're building (structured context, reasoning scratchpads) might become native to AI tools, making our approach redundant.

**Probability**: Medium-High

**Impact**: Could render core differentiation obsolete

**Mitigation**:
- Build strong relationships with AI tool vendors; position as complement
- Focus on technology-agnostic context that works across tools
- Build deeper value: templates, team sync, governance features
- Monitor AI tool development closely; adapt quickly
- Consider partnership/acquisition as exit strategy

### Risk: Market Saturation

**Description**: The developer tools space is crowded. Many scaffolding tools exist, and new AI tools launch frequently.

**Probability**: Medium

**Impact**: Difficulty gaining traction and awareness

**Mitigation**:
- Clear differentiation (AI context system is unique)
- Strong content marketing and developer advocacy
- Community building through open source
- Integration with popular AI tools (make Scaffold AI the recommended way to structure projects)
- Focus on developer experience excellence

### Risk: CLI Fatigue

**Description**: Developers may resist installing yet another CLI tool, preferring to stick with familiar tools even if less capable.

**Probability**: Medium

**Impact**: Slower adoption than projected

**Mitigation**:
- Make installation as frictionless as possible
- Clear, immediate value demonstration
- Integration with existing workflows (don't require tool switching)
- IDE plugins for those who prefer GUI
- Support for `npx` style usage without installation

## Technical Risks

### Risk: Context Window Limitations

**Description**: AI models have context window limits. The AI context files we generate might exceed limits for smaller models or older model versions.

**Probability**: Medium

**Impact**: AI tools might not read complete context, reducing effectiveness

**Mitigation**:
- Tiered context verbosity (comprehensive/standard/minimal)
- Smart context selection (include only relevant sections)
- Context summarization for smaller context windows
- Documentation of context sizes and recommendations
- Monitoring of model context window evolution

### Risk: AI Model API Changes

**Description**: AI models change behavior with updates. Context optimized for one model version might not work well with others.

**Probability**: High (continuous)

**Impact**: Degraded AI context effectiveness over time

**Mitigation**:
- Model-agnostic context design
- Regular testing across major AI tools and models
- Community feedback mechanisms for context effectiveness
- Iterative improvement based on real-world usage
- Version tracking for context templates

### Risk: Prompt Brittleness

**Description**: AI models can be sensitive to prompt wording. Small changes in context files might significantly change AI behavior, making maintenance tricky.

**Probability**: Medium

**Impact**: Unpredictable AI behavior, difficult debugging

**Mitigation**:
- Extensive testing of context templates
- Versioned templates with change documentation
- Monitoring of AI output quality
- Conservative changes; extensive validation before release

## Business Risks

### Risk: Dependency on AI Ecosystem

**Description**: Our value proposition depends on AI coding assistants continuing to be important. If AI coding loses momentum or faces regulatory/business challenges, our tool becomes less valuable.

**Probability**: Low

**Impact**: Existential threat to business model

**Mitigation**:
- The scaffolding value exists independent of AI context
- Build strong standalone scaffolding capabilities
- Diversify into related areas (project governance, team standards)
- Stay adaptable to market shifts

### Risk: Pricing Model Sustainability

**Description**: Our chosen pricing model might not generate sufficient revenue or might face resistance from target customers.

**Probability**: Medium

**Impact**: Revenue below projections, difficulty sustaining development

**Mitigation**:
- Start with conservative projections
- Test pricing with early adopters
- Offer annual discounts to improve cash flow
- Enterprise tier provides significant revenue potential
- Monitor conversion metrics and adjust

### Risk: Support Burden

**Description**: As an open source project with paid tiers, we may face disproportionate support demands from free users.

**Probability**: Medium-High

**Impact**: Developer time consumed by support, reduced feature development

**Mitigation**:
- Comprehensive documentation
- Community support forums
- Clear support tier boundaries
- Automated diagnostics and troubleshooting
- Prioritized support for paying customers

## Operational Risks

### Risk: Maintenance Overhead

**Description**: Maintaining templates for rapidly evolving frameworks and technologies requires continuous effort.

**Probability**: High (continuous)

**Impact**: Templates become outdated, users have poor experience

**Mitigation**:
- Community contribution model for templates
- Automated testing of templates
- Deprecation policy for outdated technologies
- Focus on most popular combinations
- Clear versioning of templates

### Risk: Security Vulnerabilities

**Description**: Generated projects might include vulnerable dependencies or insecure configurations.

**Probability**: Medium

**Impact**: Security incidents for users, reputation damage

**Mitigation**:
- Security scanning of templates
- Dependency version pinning
- Regular security audits
- Automated dependency updates (Dependabot)
- Security advisory process

---

# Roadmap

## Phase 1: Foundation (Months 1-3)

**Goal**: Deliver core scaffolding functionality with initial AI context system

**Deliverables**:
- CLI with `init` command and interactive wizard
- Basic template support (React, Vue, FastAPI, Express, Next.js, Django)
- AI context system with system prompt, architecture docs, brainstorm.md
- Installation via curl, Homebrew, npm
- Basic documentation

**Success Metrics**:
- 1,000 GitHub stars
- 500+ successful project generations
- <5% error rate on generation

## Phase 2: Polish (Months 4-6)

**Goal**: Improve user experience and expand template coverage

**Deliverables**:
- Enhanced TUI with animations and better responsiveness
- Additional templates (Go backend, mobile backends, Svelte)
- Self-update system
- Windows installers (winget, Chocolatey, Scoop)
- Template validation and testing infrastructure
- Improved documentation with examples

**Success Metrics**:
- 5,000 GitHub stars
- 5,000+ project generations
- Positive user feedback on UX
- Templates for 10+ technology combinations

## Phase 3: Growth (Months 7-12)

**Goal**: Launch premium features and expand market reach

**Deliverables**:
- Pro tier with template marketplace
- Team synchronization features
- Additional CLI commands (`add`, `context update`)
- IDE plugin for VS Code
- Enterprise pilot program
- Content marketing campaign
- Conference talks and community building

**Success Metrics**:
- 10,000 GitHub stars
- 500 paid subscribers
- 3+ enterprise pilots
- Coverage in developer media

## Phase 4: Scale (Year 2)

**Goal**: Establish as standard for AI-ready development

**Deliverables**:
- Full enterprise tier with self-hosting
- Plugin system for extensibility
- Advanced AI context features
- API for programmatic access
- Training and certification programs
- Partnership integrations (Cursor, Continue.dev, etc.)

**Success Metrics**:
- 50,000 GitHub stars
- 5,000+ paid subscribers
- 20+ enterprise customers
- Recognized standard for AI context

## Feature Prioritization Framework

Features are prioritized using a modified RICE framework:

**Reach**: How many users will this feature benefit?
**Impact**: How much will this feature improve user outcomes?
**Confidence**: How confident are we in our estimates?
**Effort**: How much work is required?

Score = (Reach × Impact × Confidence) / Effort

Features are prioritized within each phase based on this score, adjusted for strategic alignment and dependencies.

---

# Strategic Conclusion

## The Opportunity

The emergence of AI coding assistants has created a fundamental shift in how software is developed. This shift creates an opportunity that didn't exist five years ago: the need for structured context that enables AI tools to produce consistent, high-quality, convention-compliant code.

Existing tools address fragments of this need. Cursor has `.cursorrules`. Yeoman has generators. Nx has architectural enforcement. But no tool comprehensively addresses the full lifecycle of creating and maintaining AI-operable project environments.

Scaffold AI is positioned to fill this gap with a unique combination of capabilities:

1. **AI-First Architecture**: Every feature designed with AI context in mind
2. **Brainstorm.md Paradigm**: Enabling autonomous AI reasoning without user interruption
3. **AI-Agnostic Approach**: Working across all major AI tools and models
4. **Integrated Experience**: Scaffolding and context creation in one seamless flow

## The Strategy

Our strategy is to:

1. **Establish the Standard**: Make the Scaffold AI context format the de facto standard for AI-ready projects through open source adoption
2. **Build Moats**: Create switching costs through template ecosystems, team synchronization, and organizational investment in context files
3. **Capture Value**: Convert a portion of users to paid tiers while maintaining strong free tier value
4. **Expand Scope**: Grow from scaffolding to comprehensive AI development governance

## The Ask

Building this vision requires:

**Investment**:
- $2-3M seed funding for 18-month runway
- Team: 4-6 engineers, 1 designer, 1 developer advocate

**Focus**:
- Year 1: Product-market fit for core scaffolding + AI context
- Year 2: Growth through template ecosystem and enterprise features

**Success Criteria**:
- 50,000+ active users
- $1M+ ARR
- Recognition as the standard for AI-ready development

## The Vision

In five years, Scaffold AI should be as essential to development as Git is today. Every new project should start with `scaffold init`. Every AI assistant should expect to find `.ai-context/` in the projects it works with. Every team should have their organizational standards encoded in Scaffold AI templates.

This is an ambitious vision, but the market opportunity is real, the timing is right, and the execution plan is sound. The question is not whether someone will build this—someone will. The question is whether we will be the ones to build it correctly.

---

*This blueprint represents our current understanding and strategy. It will evolve as we learn from users, observe market changes, and iterate on the product. The core insight—structured context for AI-assisted development—remains constant. Everything else is execution.*

---

**Document Version**: 1.0
**Last Updated**: January 2025
**Author**: Strategic Planning & Architecture Team
