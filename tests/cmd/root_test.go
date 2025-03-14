package cmd_test

import (
	"testing"

	"github.com/adi-Roth/flactanCLI/cmd"
	"github.com/adi-Roth/flactanCLI/internal/utils"
)

func TestExecute(t *testing.T) {
	// Mock exit handler for testing
	exitCalled := false
	utils.ExitHandler = func(code int, err error) {
		exitCalled = true
	}

	cmd.Execute()

	if exitCalled {
		t.Errorf("Expected Execute() to complete without calling ExitHandler")
	}
}
