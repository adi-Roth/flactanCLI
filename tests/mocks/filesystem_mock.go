package mocks

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// MockFileSystem simulates file system operations for testing
type MockFileSystem struct {
	Files map[string][]byte // Simulated in-memory file storage
}

// MkdirAll simulates creating directories (always succeeds)
func (fs *MockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	fmt.Println("[MOCK] Creating directory:", path)
	return nil
}

// WriteFile simulates writing files
func (fs *MockFileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	if fs.Files == nil {
		fs.Files = make(map[string][]byte)
	}
	fmt.Println("[MOCK] Writing file:", filename)
	fs.Files[filename] = data
	return nil
}

// ReadFile simulates reading files
func (fs *MockFileSystem) ReadFile(filename string) ([]byte, error) {
	if data, exists := fs.Files[filename]; exists {
		fmt.Println("[MOCK] Reading file:", filename)
		return data, nil
	}
	fmt.Println("[MOCK] File not found:", filename)
	return nil, &os.PathError{Op: "open", Path: filename, Err: errors.New("no such file or directory")}
}

// FileExists simulates checking if a file exists
func (fs *MockFileSystem) FileExists(filename string) bool {
	_, exists := fs.Files[filename]
	fmt.Printf("[MOCK] Checking if file exists: %s -> %v\n", filename, exists)
	return exists
}

// ValidateYAML simulates checking if a file contains valid YAML content
func (fs *MockFileSystem) ValidateYAML(filename string) bool {
	data, exists := fs.Files[filename]
	if !exists {
		fmt.Printf("[MOCK] YAML validation failed (file not found): %s ❌\n", filename)
		return false
	}

	// Try parsing YAML (simulating behavior of `utils.ValidateYAML`)
	var parsedData map[string]interface{}
	err := yaml.Unmarshal(data, &parsedData)
	if err != nil {
		fmt.Printf("[MOCK] YAML validation failed (invalid format): %s ❌\n", filename)
		return false
	}

	fmt.Printf("[MOCK] YAML validation passed: %s ✅\n", filename)
	return true
}
