package cmd_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/adi-Roth/flactanCLI/internal/config"
	"github.com/adi-Roth/flactanCLI/internal/utils"
	"github.com/adi-Roth/flactanCLI/internal/validation"
	"github.com/stretchr/testify/assert"
)

func TestValidateCommand(t *testing.T) {
	fs := utils.OSFileSystem{} // Use real file system

	// Create a temporary test directory
	tempDir, err := os.MkdirTemp("", "flactancli-test")
	if err != nil {
		t.Fatalf("Failed to create temporary test directory: %v", err)
	}
	defer os.RemoveAll(tempDir) // Clean up after test

	// Use InitializeConfig to create actual files in tempDir
	err = config.InitializeConfig(fs, tempDir)
	if err != nil {
		t.Fatalf("Failed to run config init: %v", err)
	}

	fmt.Println("[TEST] Running validation with tempDir:", tempDir)

	// 🔍 Capture logs instead of stdout
	var logBuffer bytes.Buffer
	utils.Logger.SetOutput(&logBuffer) // Redirect logrus output to buffer
	defer func() {
		utils.Logger.SetOutput(os.Stderr) // Restore normal logging after test
	}()

	// Run validation
	validation.RunValidation(tempDir)

	// Read captured logs
	logOutput := logBuffer.String()
	fmt.Println("Captured Logs:\n", logOutput)

	// Assertions
	assert.Contains(t, logOutput, "✔ config.yaml is valid ✅", "Expected validation output missing")
	assert.Contains(t, logOutput, "✔ tools.yaml is valid ✅", "Expected validation output missing")
}
