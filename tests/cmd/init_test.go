package cmd_test

import (
	"os"
	"testing"

	"github.com/adi-Roth/flactanCLI/cmd"
	"github.com/adi-Roth/flactanCLI/internal/system"
	"github.com/adi-Roth/flactanCLI/tests/mocks"
	"gopkg.in/yaml.v2"
)

// Config structure matching the one in init.go
type Config struct {
	OSName    string `yaml:"os-name"`
	OSVersion string `yaml:"os-version"`
	Internet  string `yaml:"internet"`
	ToolsPath string `yaml:"tools-path"`
}

// TestInitCommand verifies that the init command correctly generates config files
func TestInitCommand(t *testing.T) {
	mockFS := &mocks.MockFileSystem{}

	// Create a temporary test directory
	tempDir, err := os.MkdirTemp("", "flactancli-test")
	if err != nil {
		t.Fatalf("Failed to create temporary test directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after test

	// Run init function with mock filesystem and test directory
	cmd.InitializeConfig(mockFS, tempDir)

	// Check if config.yaml exists
	configData, err := mockFS.ReadFile(tempDir + "/config.yaml")
	if err != nil {
		t.Fatalf("Expected config.yaml to be created, but it does not exist")
	}

	// Read and parse config.yaml
	var config Config
	if err := yaml.Unmarshal(configData, &config); err != nil {
		t.Fatalf("Failed to parse config.yaml: %v", err)
	}

	// Verify OS information
	expectedOS, expectedVersion := system.GetOSInfo()
	if config.OSName != expectedOS {
		t.Errorf("Expected OSName to be %s, got %s", expectedOS, config.OSName)
	}
	if config.OSVersion != expectedVersion {
		t.Errorf("Expected OSVersion to be %s, got %s", expectedVersion, config.OSVersion)
	}

	// Check if tools.yaml exists
	if _, err := mockFS.ReadFile(tempDir + "/tools.yaml"); err != nil {
		t.Fatalf("Expected tools.yaml to be created, but it does not exist")
	}
}
