package system_test

import (
	"os"
	"os/exec"
	"testing"

	"github.com/adi-Roth/flactanCLI/internal/system"
)

func TestCheckAdminPrivileges(t *testing.T) {
	// Backup the original getOS function
	originalGetOS := system.GetOS
	defer func() { system.GetOS = originalGetOS }() // Restore after test

	// ✅ **Test Linux/macOS root user**
	t.Run("Unix_Root", func(t *testing.T) {
		system.GetOS = func() string { return "linux" } // Mock Linux

		if os.Geteuid() == 0 {
			if !system.CheckAdminPrivileges() {
				t.Errorf("Expected true when running as root, got false")
			}
		} else {
			if system.CheckAdminPrivileges() {
				t.Errorf("Expected false when not running as root, got true")
			}
		}
	})

	// ✅ **Test Windows Admin**
	t.Run("Windows_Admin", func(t *testing.T) {
		system.GetOS = func() string { return "windows" } // Mock Windows

		cmd := exec.Command("net", "session")
		err := cmd.Run()

		expected := err == nil
		actual := system.CheckAdminPrivileges()

		if expected != actual {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	// ✅ **Test Unsupported OS**
	t.Run("UnsupportedOS", func(t *testing.T) {
		system.GetOS = func() string { return "unsupportedOS" } // Mock an unknown OS

		if system.CheckAdminPrivileges() {
			t.Errorf("Expected false for unsupported OS, got true")
		}
	})
}
