// Package utils provides common utility functions for the Clause CLI.
//
// This package contains general-purpose utilities organized into categories:
//
// # File Utilities (file.go)
//
// Functions for file system operations:
//   - FileExists, IsDirectory, IsRegularFile, IsSymlink
//   - EnsureDirectory, CreateFile, WriteFile, ReadFile
//   - CopyFile, CopyDirectory, MoveFile, DeleteFile
//   - ListFiles, ListDirectories, WalkFiles
//   - FileSize, FileHash, AtomicWrite, BackupFile
//
// Example:
//
//	if utils.FileExists("config.yaml") {
//	    config, err := utils.ReadFile("config.yaml")
//	    // handle error and use config
//	}
//
// # Path Utilities (path.go)
//
// Functions for path manipulation and resolution:
//   - ExpandHome, ContractHome, GetHomeDirectory
//   - ToAbsPath, IsAbsPath, JoinPath, CleanPath
//   - FindFileUp, FindGitRoot, IsInDirectory
//   - IsValidFilename, SanitizeFilename
//
// Example:
//
//	// Expand ~ to home directory
//	configPath := utils.ExpandHome("~/.clause/config.yaml")
//
//	// Find .git directory from current location
//	gitRoot := utils.FindGitRoot(".")
//
// # String Utilities (string.go)
//
// Functions for string manipulation and formatting:
//   - Truncate, PadLeft, PadRight, Center
//   - CamelCase, PascalCase, SnakeCase, KebabCase
//   - TitleCase, Capitalize, Upper, Lower
//   - Wrap, Indent, Dedent
//   - IsEmpty, IsNumeric, IsAlpha, IsAlphaNumeric
//
// Example:
//
//	// Convert to kebab-case
//	slug := utils.KebabCase("My Project Name") // "my-project-name"
//
//	// Truncate with ellipsis
//	short := utils.Truncate("Very long text...", 10) // "Very lo..."
//
// # Slice Utilities (slice.go)
//
// Generic functions for slice operations:
//   - Contains, IndexOf, Find, Filter, Map
//   - Unique, Reverse, Chunk, Partition
//   - First, Last, Take, Drop
//   - Concat, Flatten, Difference, Intersection
//
// Example:
//
//	// Filter even numbers
//	numbers := []int{1, 2, 3, 4, 5, 6}
//	evens := utils.Filter(numbers, func(n int) bool { return n%2 == 0 })
//
//	// Map to strings
//	strings := utils.Map(evens, func(n int) string { return strconv.Itoa(n) })
//
// # Map Utilities (map.go)
//
// Generic functions for map operations:
//   - Keys, Values, Entries, FromEntries
//   - GetOrDefault, HasKey, HasValue
//   - Merge, Clone, Filter, MapValues
//   - GroupBy, CountBy, Invert
//
// Example:
//
//	// Group items by category
//	items := []Item{...}
//	grouped := utils.GroupBy(items, func(i Item) string { return i.Category })
//
// # Terminal Utilities (terminal.go)
//
// Functions for terminal detection and control:
//   - IsTerminal, GetTerminalSize, DetectColorDepth
//   - SupportsTrueColor, Supports256Colors
//   - ClearScreen, MoveCursor, HideCursor, ShowCursor
//   - EnableAlternateScreen, DisableAlternateScreen
//   - NotifyResize, StartRawMode
//
// Example:
//
//	// Check terminal capabilities
//	info := utils.DetectTerminalInfo()
//	if info.ColorDepth >= utils.ColorDepthTrue {
//	    // Use true color
//	}
//
// # Version Utilities (version.go)
//
// Functions for semantic version parsing and comparison:
//   - ParseVersion, CompareVersions
//   - IsNewer, IsOlder, VersionsEqual
//   - Version struct with comparison methods
//   - VersionRange for constraint checking
//
// Example:
//
//	// Compare versions
//	if newer, _ := utils.IsNewer("1.0.0", "1.0.1"); newer {
//	    fmt.Println("Update available!")
//	}
//
// # Design Philosophy
//
// This package follows these design principles:
//
//  1. **Type Safety**: Generic functions use Go 1.18+ type parameters
//     for compile-time type safety without reflection.
//
//  2. **Zero Dependencies**: Core utilities have minimal dependencies
//     beyond the standard library and x/term.
//
//  3. **Error Handling**: Functions that can fail return errors.
//     "Must" variants panic on error for convenience in initialization.
//
//  4. **Immutability**: Functions return new values rather than
//     modifying inputs, making them safe for concurrent use.
//
//  5. **Consistency**: Naming follows Go conventions:
//     - "Get" prefix for simple accessors
//     - "Is" prefix for boolean checks
//     - "To" prefix for conversions
//     - "Parse" prefix for string parsing
package utils
