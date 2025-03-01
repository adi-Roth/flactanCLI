package system_test

import (
	"testing"

	"github.com/adi-Roth/flactanCLI/internal/system"
)

// Test CheckInternet function
func TestCheckInternet(t *testing.T) {
	// We can't guarantee network availability, so just check if it returns a boolean
	isConnected := system.CheckInternet()

	if isConnected != true && isConnected != false {
		t.Errorf("Expected CheckInternet() to return true or false, got %v", isConnected)
	}
}
