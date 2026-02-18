# AI Governance

Clause CLI includes a powerful AI governance system designed to guide AI coding assistants (like Copilot, Cursor, etc.) to produce code that aligns with your project's standards.

## Overview

When you initialize a project with `clause init`, a `.clause/governance` directory (aliased as `ai_prompt_guidelines`) is created. This directory contains:

- **System Prompt**: Defines the core role and behavior of the AI.
- **Architecture Guidelines**: Specifies the project's architectural layers and dependency rules.
- **Technology Stack**: Whitelists approved technologies and frameworks.
- **Brainstorming Template**: Provides a structured way for the AI to "think" before coding.
- **Registry**: Tracks created components to promote reuse.

## How it Works

Modern AI assistants are context-aware. By placing these guidelines in your project, the AI reads them and (ideally) adheres to the rules defined therein.

### Key Components

1.  **System Prompt (`system_prompt.md`)**: The "constitution" for the AI. It sets the tone, rules, and prohibited actions.
2.  **Architecture (`architecture.md`)**: Visual and textual description of the system design.
3.  **Brainstorming (`brainstorm.md`)**: A scratchpad for the AI to document its reasoning process, specialized for complex tasks.

## Customization

You can customize the governance rules by editing the files in `ai_prompt_guidelines/`.

For example, to enforce stricter testing requirements, edit `system_prompt.md` to include:

```markdown
### Testing Rules
- 100% coverage required for all new utility functions.
- Integration tests must be written for all API endpoints.
```
