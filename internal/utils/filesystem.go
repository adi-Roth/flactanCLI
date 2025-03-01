package utils

import "os"

// FileSystem defines an interface for file operations (to enable mocking)
type FileSystem interface {
	MkdirAll(path string, perm os.FileMode) error
	WriteFile(filename string, data []byte, perm os.FileMode) error
	ReadFile(filename string) ([]byte, error)
}

// OSFileSystem is the real implementation for actual file operations
type OSFileSystem struct{}

// MkdirAll creates directories
func (OSFileSystem) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

// WriteFile writes data to a file
func (OSFileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

// ReadFile reads a file's content
func (OSFileSystem) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}
