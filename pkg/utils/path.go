package utils

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"
)

// ExpandHome expands the ~ character in a path to the user's home directory.
func ExpandHome(path string) string {
	if path == "" {
		return ""
	}

	if path[0] == '~' {
		if len(path) == 1 {
			return GetHomeDirectory()
		}

		if path[1] == '/' || path[1] == '\\' {
			return filepath.Join(GetHomeDirectory(), path[2:])
		}

		// Handle ~username format
		parts := strings.SplitN(path[1:], string(filepath.Separator), 2)
		username := parts[0]

		u, err := user.Lookup(username)
		if err != nil {
			return path // Can't expand, return as-is
		}

		if len(parts) == 1 {
			return u.HomeDir
		}
		return filepath.Join(u.HomeDir, parts[1])
	}

	return path
}

// ContractHome contracts the home directory to ~ if applicable.
func ContractHome(path string) string {
	home := GetHomeDirectory()
	if home == "" {
		return path
	}

	// Normalize paths for comparison
	absPath, err := filepath.Abs(path)
	if err != nil {
		return path
	}

	absHome, err := filepath.Abs(home)
	if err != nil {
		return path
	}

	// Check if path is in home directory
	if strings.HasPrefix(absPath, absHome) {
		relative := strings.TrimPrefix(absPath, absHome)
		if relative == "" {
			return "~"
		}
		return "~" + relative
	}

	return path
}

// GetHomeDirectory returns the current user's home directory.
func GetHomeDirectory() string {
	// Try USERPROFILE first (Windows)
	if home := os.Getenv("USERPROFILE"); home != "" {
		return home
	}

	// Try HOME (Unix-like)
	if home := os.Getenv("HOME"); home != "" {
		return home
	}

	// Fallback to user.Current()
	u, err := user.Current()
	if err == nil && u.HomeDir != "" {
		return u.HomeDir
	}

	return ""
}

// GetWorkingDirectory returns the current working directory.
func GetWorkingDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	return dir
}

// IsAbsPath checks if a path is absolute.
func IsAbsPath(path string) bool {
	return filepath.IsAbs(path)
}

// ToAbsPath converts a relative path to absolute.
func ToAbsPath(path string) (string, error) {
	if filepath.IsAbs(path) {
		return filepath.Clean(path), nil
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return filepath.Clean(abs), nil
}

// JoinPath joins path elements.
func JoinPath(elements ...string) string {
	return filepath.Join(elements...)
}

// SplitPath splits a path into directory and file components.
func SplitPath(path string) (dir, file string) {
	return filepath.Split(path)
}

// DirPath returns the directory part of a path.
func DirPath(path string) string {
	return filepath.Dir(path)
}

// BasePath returns the last element of a path.
func BasePath(path string) string {
	return filepath.Base(path)
}

// CleanPath cleans a path by removing . and .. elements.
func CleanPath(path string) string {
	return filepath.Clean(path)
}

// RelPath returns a relative path from base to target.
func RelPath(base, target string) (string, error) {
	return filepath.Rel(base, target)
}

// FindFileUp searches for a file by walking up the directory tree.
// Returns the path to the file if found, or empty string if not found.
func FindFileUp(name string, start string) string {
	dir := start
	if dir == "" {
		dir = GetWorkingDirectory()
	}

	for {
		candidate := filepath.Join(dir, name)
		if FileExists(candidate) {
			return candidate
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached root
			return ""
		}
		dir = parent
	}
}

// FindDirectoryUp searches for a directory by walking up the directory tree.
func FindDirectoryUp(name string, start string) string {
	dir := start
	if dir == "" {
		dir = GetWorkingDirectory()
	}

	for {
		candidate := filepath.Join(dir, name)
		if IsDirectory(candidate) {
			return candidate
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return ""
		}
		dir = parent
	}
}

// FindGitRoot finds the root of the git repository containing the given path.
func FindGitRoot(start string) string {
	return FindDirectoryUp(".git", start)
}

// IsInDirectory checks if a path is inside a given directory.
func IsInDirectory(path, dir string) (bool, error) {
	absPath, err := ToAbsPath(path)
	if err != nil {
		return false, err
	}

	absDir, err := ToAbsPath(dir)
	if err != nil {
		return false, err
	}

	rel, err := filepath.Rel(absDir, absPath)
	if err != nil {
		return false, nil
	}

	// If relative path starts with .., it's outside the directory
	return !strings.HasPrefix(rel, "..") && !strings.HasPrefix(rel, "."+string(filepath.Separator)), nil
}

// IsSubPath checks if subPath is a subdirectory of parent.
func IsSubPath(parent, subPath string) bool {
	rel, err := filepath.Rel(parent, subPath)
	if err != nil {
		return false
	}
	return !strings.HasPrefix(rel, "..")
}

// NormalizePath normalizes a path by expanding home and cleaning.
func NormalizePath(path string) string {
	path = ExpandHome(path)
	return filepath.Clean(path)
}

// NormalizePathToAbs normalizes a path and converts to absolute.
func NormalizePathToAbs(path string) (string, error) {
	path = NormalizePath(path)
	return ToAbsPath(path)
}

// CommonPrefix finds the common prefix of multiple paths.
func CommonPrefix(paths ...string) string {
	if len(paths) == 0 {
		return ""
	}

	if len(paths) == 1 {
		return paths[0]
	}

	prefix := paths[0]

	for _, path := range paths[1:] {
		for !strings.HasPrefix(path+string(filepath.Separator), prefix+string(filepath.Separator)) {
			prefix = filepath.Dir(prefix)
			if prefix == "." || prefix == string(filepath.Separator) {
				return ""
			}
		}
	}

	return prefix
}

// IsPathSeparator checks if a character is a path separator.
func IsPathSeparator(c byte) bool {
	return c == '/' || (runtime.GOOS == "windows" && c == '\\')
}

// SplitPathList splits a PATH-like environment variable into individual paths.
func SplitPathList(pathList string) []string {
	if pathList == "" {
		return nil
	}
	return filepath.SplitList(pathList)
}

// JoinPathList joins paths into a PATH-like environment variable.
func JoinPathList(paths []string) string {
	return strings.Join(paths, string(filepath.ListSeparator))
}

// IsValidFilename checks if a string is a valid filename.
func IsValidFilename(name string) bool {
	if name == "" || name == "." || name == ".." {
		return false
	}

	// Check for invalid characters based on OS
	if runtime.GOOS == "windows" {
		invalid := []string{"<", ">", ":", "\"", "|", "?", "*"}
		for _, char := range invalid {
			if strings.Contains(name, char) {
				return false
			}
		}

		// Check for reserved names on Windows
		reserved := []string{
			"CON", "PRN", "AUX", "NUL",
			"COM1", "COM2", "COM3", "COM4", "COM5", "COM6", "COM7", "COM8", "COM9",
			"LPT1", "LPT2", "LPT3", "LPT4", "LPT5", "LPT6", "LPT7", "LPT8", "LPT9",
		}

		base := strings.ToUpper(name)
		for _, r := range reserved {
			if base == r || strings.HasPrefix(base, r+".") {
				return false
			}
		}
	}

	// Check for path separators
	if strings.ContainsAny(name, "/\\") {
		return false
	}

	return true
}

// SanitizeFilename replaces invalid characters in a filename.
func SanitizeFilename(name string) string {
	if name == "" {
		return "unnamed"
	}

	// Replace common invalid characters
	replacer := strings.NewReplacer(
		"/", "_",
		"\\", "_",
		":", "_",
		"*", "_",
		"?", "_",
		"\"", "_",
		"<", "_",
		">", "_",
		"|", "_",
	)

	result := replacer.Replace(name)

	// Trim spaces and dots from ends
	result = strings.Trim(result, " .")

	if result == "" {
		return "unnamed"
	}

	return result
}

// PathDepth returns the depth of a path (number of components).
func PathDepth(path string) int {
	if path == "" {
		return 0
	}

	path = filepath.Clean(path)
	if path == "." || path == string(filepath.Separator) {
		return 0
	}

	count := 0
	for {
		path = filepath.Dir(path)
		count++
		if path == "." || path == string(filepath.Separator) {
			break
		}
	}

	return count
}

// EnsureSuffix ensures a path ends with the given suffix.
func EnsureSuffix(path, suffix string) string {
	if strings.HasSuffix(path, suffix) {
		return path
	}
	return path + suffix
}

// EnsurePrefix ensures a path starts with the given prefix.
func EnsurePrefix(path, prefix string) string {
	if strings.HasPrefix(path, prefix) {
		return path
	}
	return prefix + path
}

// RemoveSuffix removes a suffix from a path if present.
func RemoveSuffix(path, suffix string) string {
	if strings.HasSuffix(path, suffix) {
		return path[:len(path)-len(suffix)]
	}
	return path
}

// EnsureTrailingSeparator ensures a path ends with a separator.
func EnsureTrailingSeparator(path string) string {
	if path == "" {
		return string(filepath.Separator)
	}
	if !os.IsPathSeparator(path[len(path)-1]) {
		return path + string(filepath.Separator)
	}
	return path
}

// RemoveTrailingSeparator removes trailing path separators.
func RemoveTrailingSeparator(path string) string {
	path = strings.TrimRight(path, string(filepath.Separator))
	if path == "" {
		return string(filepath.Separator)
	}
	return path
}
