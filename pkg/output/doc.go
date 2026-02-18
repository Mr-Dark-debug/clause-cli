// Package output provides utilities for formatted console output.
//
// This package includes three main components:
//
// # Printer
//
// Printer provides styled output functions for common use cases:
//
//	printer := output.NewPrinter(nil, os.Stdout)
//	printer.PrintSuccess("Operation completed")
//	printer.PrintError("Something went wrong")
//	printer.PrintHeader("Welcome to Clause")
//	printer.PrintBullet("Item one")
//	printer.PrintCheckmark("Done")
//
// # Logger
//
// Logger provides structured logging with levels:
//
//	logger := output.NewLogger(
//	    output.WithLevel(output.LevelDebug),
//	    output.WithShowTime(true),
//	)
//	logger.Debug("Debug information")
//	logger.Info("General information")
//	logger.Warn("Warning message")
//	logger.Error("Error message")
//
// Use WithFields for structured logging:
//
//	logger.WithFields(output.Fields{"user": "john", "action": "login"}).Info("User logged in")
//
// # Table
//
// Table provides styled table rendering:
//
//	columns := []output.TableColumn{
//	    {Title: "Name", Width: 20},
//	    {Title: "Status", Width: 10},
//	    {Title: "Date", Width: 15},
//	}
//
//	table := output.NewTable(columns)
//	table.AddRow("Project A", "Active", "2024-01-15")
//	table.AddRow("Project B", "Pending", "2024-01-16")
//	fmt.Println(table.Render())
//
// # Convenience Functions
//
// Package-level functions are available for quick access:
//
//	output.PrintSuccess("Done")
//	output.PrintError("Failed")
//	output.PrintHeader("Title")
//
// # Design Philosophy
//
// This package follows these principles:
//
//  1. **Theme Aware**: All output respects the current theme colors.
//
//  2. **Terminal Detection**: Automatically adapts to terminal capabilities.
//
//  3. **Thread Safe**: Logger is safe for concurrent use.
//
//  4. **Flexible**: Options pattern for configuration.
//
//  5. **Consistent**: Consistent styling across all output methods.
package output
