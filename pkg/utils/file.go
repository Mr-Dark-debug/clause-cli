// Package utils provides common utility functions for the Clause CLI.
package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// FileExists checks if a file exists at the given path.
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// FileExistsErr checks if a file exists and returns any error.
func FileExistsErr(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// IsDirectory checks if the path points to a directory.
func IsDirectory(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

// IsRegularFile checks if the path points to a regular file.
func IsRegularFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().IsRegular()
}

// IsSymlink checks if the path is a symbolic link.
func IsSymlink(path string) bool {
	info, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return info.Mode()&os.ModeSymlink != 0
}

// EnsureDirectory creates a directory and all parent directories if they don't exist.
func EnsureDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}

// EnsureDirectoryWithPerm creates a directory with specific permissions.
func EnsureDirectoryWithPerm(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// CreateFile creates a new file with the given content.
func CreateFile(path string, content []byte) error {
	// Ensure parent directory exists
	dir := filepath.Dir(path)
	if err := EnsureDirectory(dir); err != nil {
		return fmt.Errorf("failed to create parent directory: %w", err)
	}

	// Write the file
	if err := os.WriteFile(path, content, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// CreateFileWithPerm creates a new file with specific permissions.
func CreateFileWithPerm(path string, content []byte, perm os.FileMode) error {
	dir := filepath.Dir(path)
	if err := EnsureDirectory(dir); err != nil {
		return fmt.Errorf("failed to create parent directory: %w", err)
	}

	return os.WriteFile(path, content, perm)
}

// WriteFile writes content to a file, creating it if necessary.
func WriteFile(path string, content []byte) error {
	return CreateFile(path, content)
}

// ReadFile reads the entire content of a file.
func ReadFile(path string) ([]byte, error) {
	return os.ReadFile(path)
}

// ReadFileString reads the entire content of a file as a string.
func ReadFileString(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// CopyFile copies a file from src to dst.
func CopyFile(src, dst string) error {
	// Open source file
	sourceFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer sourceFile.Close()

	// Get source file info for permissions
	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return fmt.Errorf("failed to get source file info: %w", err)
	}

	// Ensure destination directory exists
	dstDir := filepath.Dir(dst)
	if err := EnsureDirectory(dstDir); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Create destination file
	dstFile, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, sourceInfo.Mode())
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	// Copy content
	if _, err := io.Copy(dstFile, sourceFile); err != nil {
		return fmt.Errorf("failed to copy file content: %w", err)
	}

	return nil
}

// CopyDirectory recursively copies a directory from src to dst.
func CopyDirectory(src, dst string) error {
	// Get source directory info
	srcInfo, err := os.Stat(src)
	if err != nil {
		return fmt.Errorf("failed to get source directory info: %w", err)
	}

	// Create destination directory with same permissions
	if err := os.MkdirAll(dst, srcInfo.Mode()); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Read directory entries
	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read source directory: %w", err)
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())

		if entry.IsDir() {
			if err := CopyDirectory(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, dstPath); err != nil {
				return err
			}
		}
	}

	return nil
}

// MoveFile moves a file from src to dst.
func MoveFile(src, dst string) error {
	// Try rename first (fastest if on same filesystem)
	if err := os.Rename(src, dst); err == nil {
		return nil
	}

	// Fall back to copy + delete
	if err := CopyFile(src, dst); err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	if err := os.Remove(src); err != nil {
		return fmt.Errorf("failed to remove source file: %w", err)
	}

	return nil
}

// DeleteFile removes a file.
func DeleteFile(path string) error {
	return os.Remove(path)
}

// DeleteDirectory removes a directory and all its contents.
func DeleteDirectory(path string) error {
	return os.RemoveAll(path)
}

// EmptyDirectory removes all contents of a directory but keeps the directory itself.
func EmptyDirectory(path string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if err := os.RemoveAll(fullPath); err != nil {
			return err
		}
	}

	return nil
}

// ListFiles lists all files in a directory (non-recursive).
func ListFiles(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}

// ListDirectories lists all directories in a directory (non-recursive).
func ListDirectories(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var dirs []string
	for _, entry := range entries {
		if entry.IsDir() {
			dirs = append(dirs, entry.Name())
		}
	}

	return dirs, nil
}

// ListAll lists all entries in a directory (non-recursive).
func ListAll(path string) ([]string, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	var names []string
	for _, entry := range entries {
		names = append(names, entry.Name())
	}

	return names, nil
}

// WalkFiles walks a directory tree and returns all files matching the pattern.
func WalkFiles(root, pattern string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			matched, err := filepath.Match(pattern, filepath.Base(path))
			if err != nil {
				return err
			}
			if matched {
				files = append(files, path)
			}
		}

		return nil
	})

	return files, err
}

// WalkAll walks a directory tree and calls the callback for each entry.
func WalkAll(root string, callback func(path string, d fs.DirEntry) error) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		return callback(path, d)
	})
}

// FileSize returns the size of a file in bytes.
func FileSize(path string) (int64, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Size(), nil
}

// FileHash returns the SHA256 hash of a file.
func FileHash(path string) (string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// FilePermissions returns the file permissions.
func FilePermissions(path string) (os.FileMode, error) {
	info, err := os.Stat(path)
	if err != nil {
		return 0, err
	}
	return info.Mode().Perm(), nil
}

// SetFilePermissions sets the file permissions.
func SetFilePermissions(path string, perm os.FileMode) error {
	return os.Chmod(path, perm)
}

// IsExecutable checks if a file is executable.
func IsExecutable(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.Mode().Perm()&0111 != 0
}

// MakeExecutable makes a file executable.
func MakeExecutable(path string) error {
	return SetFilePermissions(path, 0755)
}

// Touch creates an empty file or updates the modification time.
func Touch(path string) error {
	// Ensure parent directory exists
	dir := filepath.Dir(path)
	if err := EnsureDirectory(dir); err != nil {
		return err
	}

	// Try to update timestamp if file exists
	if FileExists(path) {
		now := time.Now()
		return os.Chtimes(path, now, now)
	}

	// Create empty file
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	return file.Close()
}

// TempDir creates a temporary directory with the given prefix.
func TempDir(prefix string) (string, error) {
	return os.MkdirTemp("", prefix)
}

// TempFile creates a temporary file with the given prefix.
func TempFile(prefix string) (*os.File, error) {
	return os.CreateTemp("", prefix)
}

// AtomicWrite writes content to a file atomically using a temp file.
func AtomicWrite(path string, content []byte) error {
	// Create temp file in same directory for atomic rename
	dir := filepath.Dir(path)
	if err := EnsureDirectory(dir); err != nil {
		return err
	}

	tempFile, err := os.CreateTemp(dir, ".tmp-")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tempPath := tempFile.Name()

	// Ensure cleanup on failure
	defer func() {
		tempFile.Close()
		if FileExists(tempPath) {
			os.Remove(tempPath)
		}
	}()

	// Write content
	if _, err := tempFile.Write(content); err != nil {
		return fmt.Errorf("failed to write to temp file: %w", err)
	}

	// Ensure data is written to disk
	if err := tempFile.Sync(); err != nil {
		return fmt.Errorf("failed to sync temp file: %w", err)
	}

	// Close before rename
	if err := tempFile.Close(); err != nil {
		return fmt.Errorf("failed to close temp file: %w", err)
	}

	// Atomic rename
	if err := os.Rename(tempPath, path); err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}

// BackupFile creates a backup of a file with a .bak extension.
func BackupFile(path string) (string, error) {
	backupPath := path + ".bak"

	// If backup exists, add a number
	counter := 1
	for FileExists(backupPath) {
		backupPath = fmt.Sprintf("%s.bak.%d", path, counter)
		counter++
	}

	if err := CopyFile(path, backupPath); err != nil {
		return "", err
	}

	return backupPath, nil
}

// FileExtension returns the file extension (without the dot).
func FileExtension(path string) string {
	ext := filepath.Ext(path)
	return strings.TrimPrefix(ext, ".")
}

// FileNameWithoutExtension returns the filename without extension.
func FileNameWithoutExtension(path string) string {
	base := filepath.Base(path)
	ext := filepath.Ext(base)
	return strings.TrimSuffix(base, ext)
}

// HasExtension checks if a file has one of the given extensions.
func HasExtension(path string, extensions ...string) bool {
	ext := strings.ToLower(FileExtension(path))
	for _, e := range extensions {
		if strings.ToLower(strings.TrimPrefix(e, ".")) == ext {
			return true
		}
	}
	return false
}
