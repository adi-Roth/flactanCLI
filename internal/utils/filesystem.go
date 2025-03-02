package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

// FileSystem defines an interface for file operations (to enable mocking)
type FileSystem interface {
	MkdirAll(path string, perm os.FileMode) error
	WriteFile(filename string, data []byte, perm os.FileMode) error
	ReadFile(filename string) ([]byte, error)
	FileExists(filename string) bool
	ValidateYAML(filename string) bool
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

// FileExists checks if a file exists
func (OSFileSystem) FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// ValidateYAML checks if a given file is a valid YAML file
func (OSFileSystem) ValidateYAML(filename string) bool {
	data, err := os.ReadFile(filename)
	if err != nil {
		return false // File does not exist or cannot be read
	}

	var parsedData map[string]interface{}
	err = yaml.Unmarshal(data, &parsedData)
	return err == nil // Returns true if the file is valid YAML
}
