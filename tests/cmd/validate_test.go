package cmd_test

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/adi-Roth/flactanCLI/cmd"
	"github.com/adi-Roth/flactanCLI/internal/utils"
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
	cmd.InitializeConfig(fs, tempDir)

	fmt.Println("[TEST] Running validation with tempDir:", tempDir)

	// Redirect stdout to capture output
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Run validation on the real file system
	cmd.RunValidation(tempDir)

	// Restore stdout
	w.Close()
	os.Stdout = oldStdout
	capturedOutput, _ := io.ReadAll(r)

	// Convert output to a string
	output := string(capturedOutput)
	fmt.Println("Captured Output:\n", output)

	// Assertions
	assert.Contains(t, output, "✔ config.yaml is valid ✅", "Expected validation output missing")
	assert.Contains(t, output, "✔ tools.yaml is valid ✅", "Expected validation output missing")
}
