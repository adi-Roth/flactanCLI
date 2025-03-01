package mocks

import (
	"errors"
	"fmt"
	"os"
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
