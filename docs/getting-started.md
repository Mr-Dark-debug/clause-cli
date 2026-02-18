# Getting Started with Clause CLI

## Prerequisites

- Go 1.21+ (if building from source)
- Git
- A terminal (Bash, Zsh, PowerShell, Fish)

## Installation

See [Installation Guide](installation.md) for detailed instructions for your operating system.

## Your First Project

1.  **Initialize a new project**:

    ```bash
    clause init my-app
    ```

2.  **Follow the wizard**:
    - Select your frontend framework (e.g., Next.js)
    - Select your backend framework (e.g., FastAPI)
    - Choose your database (e.g., PostgreSQL)

3.  **Explore the generated structure**:

    ```bash
    cd my-app
    ls -R
    ```

    You'll see a `frontend` folder, a `backend` folder, and importantly, an `ai_prompt_guidelines` folder.

4.  **Start coding with AI**:
    Open the project in your favorite AI-powered editor (e.g., VS Code with Copilot, Cursor). The AI will now be aware of your project's specific constraints and architecture.

## Next Steps

- [Learn about AI Governance](governance.md)
- [Explore Templates](templates.md)
- [Configure the CLI](configuration.md)
