package mocks

import (
	"errors"
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"
)

// MockFileSystem simulates file system operations for testing
type MockFileSystem struct {
	Files        map[string][]byte // Simulated in-memory file storage
	MkdirAllErr  error             // Simulated error for MkdirAll
	WriteFileErr map[string]error  // Simulated errors for WriteFile (per file)
	ReadFileErr  map[string]error  // Simulated errors for ReadFile (per file)
	mu           sync.Mutex        // Mutex for safe concurrent access
}

// NewMockFileSystem creates a new instance of MockFileSystem
func NewMockFileSystem() *MockFileSystem {
	return &MockFileSystem{
		Files:        make(map[string][]byte),
		WriteFileErr: make(map[string]error),
		ReadFileErr:  make(map[string]error),
	}
}

// MkdirAll simulates creating directories (with optional failure)
func (fs *MockFileSystem) MkdirAll(path string, perm os.FileMode) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	if fs.MkdirAllErr != nil {
		fmt.Println("[MOCK] Simulated MkdirAll failure:", fs.MkdirAllErr) // 🔍 Debugging
		return fs.MkdirAllErr
	}
	fmt.Println("[MOCK] Creating directory:", path)
	return nil
}

// WriteFile simulates writing files (with optional failure)
func (fs *MockFileSystem) WriteFile(filename string, data []byte, perm os.FileMode) error {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	if err, exists := fs.WriteFileErr[filename]; exists {
		return err
	}

	fmt.Println("[MOCK] Writing file:", filename)
	fs.Files[filename] = data
	return nil
}

// ReadFile simulates reading files (with optional failure)
func (fs *MockFileSystem) ReadFile(filename string) ([]byte, error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()

	if err, exists := fs.ReadFileErr[filename]; exists {
		return nil, err
	}

	if data, exists := fs.Files[filename]; exists {
		fmt.Println("[MOCK] Reading file:", filename)
		return data, nil
	}
	return nil, &os.PathError{Op: "open", Path: filename, Err: errors.New("no such file or directory")}
}

// FileExists simulates checking if a file exists
func (fs *MockFileSystem) FileExists(filename string) bool {
	fs.mu.Lock()
	defer fs.mu.Unlock()

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

	var parsedData map[string]interface{}
	err := yaml.Unmarshal(data, &parsedData)
	if err != nil {
		fmt.Printf("[MOCK] YAML validation failed (invalid format): %s ❌\n", filename)
		return false
	}

	fmt.Printf("[MOCK] YAML validation passed: %s ✅\n", filename)
	return true
}

// SetMkdirAllError sets an error for MkdirAll (for testing)
func (fs *MockFileSystem) SetMkdirAllError(err error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.MkdirAllErr = err
}

// SetWriteFileError sets an error for WriteFile (per file)
func (fs *MockFileSystem) SetWriteFileError(filename string, err error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.WriteFileErr[filename] = err
}

// SetReadFileError sets an error for ReadFile (per file)
func (fs *MockFileSystem) SetReadFileError(filename string, err error) {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.ReadFileErr[filename] = err
}

// ResetErrors clears all simulated errors
func (fs *MockFileSystem) ResetErrors() {
	fs.mu.Lock()
	defer fs.mu.Unlock()
	fs.MkdirAllErr = nil
	fs.WriteFileErr = make(map[string]error)
	fs.ReadFileErr = make(map[string]error)
}
