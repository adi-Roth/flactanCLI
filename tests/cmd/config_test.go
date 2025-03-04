package cmd_test

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/adi-Roth/flactanCLI/cmd"
	"github.com/adi-Roth/flactanCLI/internal/config"
	"github.com/adi-Roth/flactanCLI/internal/system"
	"github.com/adi-Roth/flactanCLI/tests/mocks"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func TestInitCommand(t *testing.T) {
	mockFS := mocks.NewMockFileSystem() // ✅ Use the correct constructor

	// Create a temporary test directory
	tempDir, err := os.MkdirTemp("", "flactancli-test")
	if err != nil {
		t.Fatalf("Failed to create temporary test directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after test

	// Reset any previous errors in the mock
	mockFS.ResetErrors()

	// Debugging: Print expected config.yaml path
	expectedConfigPath := tempDir + "/config.yaml"
	fmt.Println("Expected config.yaml path:", expectedConfigPath)

	// Run init function with mock filesystem and test directory
	err = cmd.RunConfigInit(mockFS, tempDir)
	if err != nil {
		t.Fatalf("Failed to run config init: %v", err)
	}

	// Check if config.yaml exists
	configData, err := mockFS.ReadFile(expectedConfigPath)
	if err != nil {
		t.Fatalf("Expected config.yaml to be created, but it does not exist: %v", err)
	}

	// Read and parse config.yaml
	var config config.Config
	if err := yaml.Unmarshal(configData, &config); err != nil {
		t.Fatalf("Failed to parse config.yaml: %v", err)
	}

	// Verify OS information
	expectedOS, expectedArch, expectedVersion := system.GetOSInfo()
	assert.Equal(t, expectedOS, config.OSName, "OSName mismatch")
	assert.Equal(t, expectedArch, config.OSArch, "OSArch mismatch")
	assert.Equal(t, expectedVersion, config.OSVersion, "OSVersion mismatch")

	// Check if tools.yaml exists
	expectedToolsPath := tempDir + "/tools.yaml"
	_, err = mockFS.ReadFile(expectedToolsPath)
	assert.NoError(t, err, "Expected tools.yaml to be created, but it does not exist")

	// ✅ Simulate WriteFile error (optional)
	mockFS.SetWriteFileError(expectedConfigPath, errors.New("mock write error"))
	err = mockFS.WriteFile(expectedConfigPath, []byte("test"), 0644)
	assert.Error(t, err, "Mock error should have been triggered for config.yaml")

	// ✅ Simulate ReadFile error (optional)
	mockFS.SetReadFileError(expectedConfigPath, errors.New("mock read error"))
	_, err = mockFS.ReadFile(expectedConfigPath)
	assert.Error(t, err, "Mock error should have been triggered for config.yaml read")
}

func TestInitCommand_ErrorCases(t *testing.T) {
	mockFS := mocks.NewMockFileSystem() // ✅ Use NewMockFileSystem()
	tempDir, err := os.MkdirTemp("", "flactancli-test")
	if err != nil {
		t.Fatalf("Failed to create temporary test directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after test

	// 1️⃣ **Simulate directory creation failure**
	fmt.Println("🔍 Testing directory creation failure...")
	mockFS.SetMkdirAllError(fmt.Errorf("failed to create directory"))

	err = cmd.RunConfigInit(mockFS, tempDir) // ✅ Directly capture returned error
	if err == nil {
		t.Errorf("❌ Expected error when creating directory, but got nil")
	} else {
		fmt.Println("✅ Caught expected error:", err) // 🔍 Debugging output
		if !strings.Contains(err.Error(), "failed to create directory") {
			t.Errorf("❌ Expected error message 'failed to create directory', got: %v", err)
		}
	}

	// 2️⃣ **Simulate config file writing failure**
	fmt.Println("🔍 Testing config.yaml writing failure...")
	mockFS.ResetErrors()
	mockFS.SetWriteFileError(tempDir+"/config.yaml", fmt.Errorf("failed to write config file"))

	err = cmd.RunConfigInit(mockFS, tempDir)
	if err == nil {
		t.Errorf("❌ Expected config file write error, but got nil")
	} else {
		fmt.Println("✅ Caught expected error:", err)
		if !strings.Contains(err.Error(), "failed to write config file") {
			t.Errorf("❌ Expected error message 'failed to write config file', got: %v", err)
		}
	}

	// 3️⃣ **Simulate tools.yaml writing failure**
	fmt.Println("🔍 Testing tools.yaml writing failure...")
	mockFS.ResetErrors()
	mockFS.SetWriteFileError(tempDir+"/tools.yaml", fmt.Errorf("failed to write tools file"))

	err = cmd.RunConfigInit(mockFS, tempDir)
	if err == nil {
		t.Errorf("❌ Expected tools file write error, but got nil")
	} else {
		fmt.Println("✅ Caught expected error:", err)
		if !strings.Contains(err.Error(), "failed to write tools file") {
			t.Errorf("❌ Expected error message 'failed to write tools file', got: %v", err)
		}
	}

	// 4️⃣ **Simulate tools.yaml read failure**
	fmt.Println("🔍 Testing tools.yaml read failure...")
	mockFS.ResetErrors()
	mockFS.SetReadFileError(tempDir+"/tools.yaml", fmt.Errorf("failed to read tools file"))

	err = cmd.RunConfigInit(mockFS, tempDir)
	if err != nil {
		t.Errorf("❌ Unexpected error when reading tools file: %v", err)
	} else {
		fmt.Println("✅ No error on tools.yaml read failure (expected behavior)")
	}
}
