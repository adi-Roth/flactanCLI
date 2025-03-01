package system_test

import (
	"testing"

	"github.com/adi-Roth/flactanCLI/internal/system"
)

// Test GetOSInfo function
func TestGetOSInfo(t *testing.T) {
	osName, osVersion := system.GetOSInfo()

	if osName == "" {
		t.Errorf("Expected a valid OS name, got an empty string")
	}
	if osVersion == "" {
		t.Errorf("Expected a valid OS version, got an empty string")
	}
}
