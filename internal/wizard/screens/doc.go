// Package screens provides the individual screens for the wizard.
//
// Each screen implements the wizard.Screen interface and handles a specific
// aspect of project configuration. Screens are composable and reusable.
//
// Available screens:
//   - WelcomeScreen: Initial welcome and preset selection
//   - ProjectScreen: Basic project information
//   - FrontendScreen: Frontend framework and features
//   - BackendScreen: Backend framework and database
//   - InfrastructureScreen: Hosting and CI/CD
//   - GovernanceScreen: AI governance settings
//   - SummaryScreen: Configuration review and confirmation
package screens
