package cmd_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/adi-Roth/flactanCLI/cmd"
	"github.com/stretchr/testify/assert"
)

func TestExecute(t *testing.T) {
	// Set environment variable to avoid os.Exit(1)
	os.Setenv("FLACTANCLI_TEST", "true")
	defer os.Unsetenv("FLACTANCLI_TEST")

	// Capture output
	var buf bytes.Buffer
	cmd.RootCmd.SetOut(&buf)
	cmd.RootCmd.SetArgs([]string{"--help"}) // Simulate running `flactancli --help`

	// Run the command
	cmd.Execute()

	// Validate output
	output := buf.String()
	assert.Contains(t, output, "FlactanCLI", "Expected CLI help output not found")
}

// Test Execute() with an invalid command without terminating the test
func TestExecuteWithInvalidCommand(t *testing.T) {
	// Override os.Exit to prevent test termination
	exitCalled := false
	cmd.ExitFunc = func(code int) {
		exitCalled = true // Track if exit was called
	}

	// Capture output
	var buf bytes.Buffer
	cmd.RootCmd.SetOut(&buf)
	cmd.RootCmd.SetErr(&buf)
	cmd.RootCmd.SetArgs([]string{"invalid-command"})

	// Run the command (should not crash)
	cmd.Execute()

	// Validate output
	output := buf.String()
	assert.Contains(t, output, "unknown command", "Expected an error message for an invalid command")
	assert.True(t, exitCalled, "Expected os.Exit(1) to be called")
}

func TestExecuteWithNoArgs(t *testing.T) {
	// Set test mode
	os.Setenv("FLACTANCLI_TEST", "true")
	defer os.Unsetenv("FLACTANCLI_TEST")

	// Capture output
	var buf bytes.Buffer
	cmd.RootCmd.SetOut(&buf)
	cmd.RootCmd.SetArgs([]string{}) // Simulating no arguments

	// Run the command
	cmd.Execute()

	// Validate output
	output := buf.String()
	assert.Contains(t, output, "Usage", "Expected usage instructions when no arguments are provided")
}
