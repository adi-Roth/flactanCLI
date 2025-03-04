package system_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/adi-Roth/flactanCLI/internal/system"
)

// MockCommandOutput simulates command execution with predefined output
func MockCommandOutput(output string, fail bool) func(name string, args ...string) ([]byte, error) {
	return func(name string, args ...string) ([]byte, error) {
		if fail {
			return nil, fmt.Errorf("mocked command failure")
		}
		return []byte(output), nil
	}
}

func TestGetOSInfo(t *testing.T) {
	// Store the original functions
	originalRunCommand := system.RunCommand
	originalGetGOOS := system.GetRuntimeGOOS
	originalGetGOARCH := system.GetRuntimeGOARCH

	// Restore original functions after test
	defer func() {
		system.RunCommand = originalRunCommand
		system.GetRuntimeGOOS = originalGetGOOS
		system.GetRuntimeGOARCH = originalGetGOARCH
	}()

	// ✅ **Test Linux (lsb_release)**
	t.Run("Linux_lsb_release", func(t *testing.T) {
		// Mock the system command execution
		system.RunCommand = MockCommandOutput("Description:    Ubuntu 20.04.6 LTS", false)
		system.GetRuntimeGOOS = func() string { return "linux" }
		system.GetRuntimeGOARCH = func() string { return "amd64" }

		// Call the function
		osName, arch, version := system.GetOSInfo()

		// Ensure Split result has at least two elements before accessing [1]
		parts := strings.Split(version, ":")
		if len(parts) > 1 {
			version = strings.TrimSpace(parts[1])
		}

		if osName != "linux" || arch != "amd64" || version != "Ubuntu 20.04.6 LTS" {
			t.Errorf("Expected Linux Ubuntu, got: %s %s %s", osName, arch, version)
		}
	})
	t.Run("Linux_lsb_release_failure", func(t *testing.T) {
		// Mock the system command execution
		system.RunCommand = MockCommandOutput("", true)
		system.GetRuntimeGOOS = func() string { return "linux" }
		system.GetRuntimeGOARCH = func() string { return "amd64" }

		// Call the function
		osName, arch, version := system.GetOSInfo()

		// Ensure Split result has at least two elements before accessing [1]
		parts := strings.Split(version, ":")
		if len(parts) > 1 {
			version = strings.TrimSpace(parts[1])
		}

		if osName != "linux" || arch != "amd64" || version != "Unknown" {
			t.Errorf("Expected Linux Ubuntu, got: %s %s %s", osName, arch, version)
		}
	})
	t.Run("Linux_Corrupt_os_release", func(t *testing.T) {
		system.RunCommand = MockCommandOutput("BAD_DATA=12345", true) // Corrupt file
		system.GetRuntimeGOOS = func() string { return "linux" }
		system.GetRuntimeGOARCH = func() string { return "amd64" }

		osName, arch, version := system.GetOSInfo()

		if osName != "linux" || arch != "amd64" || version != "Unknown" {
			t.Errorf("Expected Unknown for corrupt os-release, got: %s %s %s", osName, arch, version)
		}
	})

	// ✅ **Test Windows**
	t.Run("Windows", func(t *testing.T) {
		system.RunCommand = MockCommandOutput("Caption\nWindows 11 Pro", false)
		system.GetRuntimeGOOS = func() string { return "windows" }
		system.GetRuntimeGOARCH = func() string { return "amd64" }

		osName, arch, version := system.GetOSInfo()
		fmt.Printf("DEBUG: OS Name: %s, Arch: %s, Version: %s\n", osName, arch, version)

		version = strings.TrimSpace(version) // Ensure no extra spaces
		if osName != "windows" || arch != "amd64" || !strings.Contains(version, "Windows 11") {
			t.Errorf("Expected Windows 11, got: %s %s %s", osName, arch, version)
		}
	})
	t.Run("Windows_failure", func(t *testing.T) {
		system.RunCommand = MockCommandOutput("", true)
		system.GetRuntimeGOOS = func() string { return "windows" }
		system.GetRuntimeGOARCH = func() string { return "amd64" }

		osName, arch, version := system.GetOSInfo()
		fmt.Printf("DEBUG: OS Name: %s, Arch: %s, Version: %s\n", osName, arch, version)

		version = strings.TrimSpace(version) // Ensure no extra spaces
		if osName != "windows" || arch != "amd64" || version != "Unknown" {
			t.Errorf("Expected Windows 11, got: %s %s %s", osName, arch, version)
		}
	})

	// ✅ **Test macOS**
	t.Run("macOS", func(t *testing.T) {
		system.RunCommand = MockCommandOutput("13.0\n", false)
		system.GetRuntimeGOOS = func() string { return "darwin" }
		system.GetRuntimeGOARCH = func() string { return "arm64" }

		osName, arch, version := system.GetOSInfo()
		version = strings.TrimSpace(version) // Ensure no extra spaces

		if osName != "darwin" || arch != "arm64" || version != "13.0" {
			t.Errorf("Expected macOS version 13.0, got: %s %s %s", osName, arch, version)
		}
	})
	t.Run("macOS_failure", func(t *testing.T) {
		system.RunCommand = MockCommandOutput("", true)
		system.GetRuntimeGOOS = func() string { return "darwin" }
		system.GetRuntimeGOARCH = func() string { return "arm64" }

		osName, arch, version := system.GetOSInfo()
		version = strings.TrimSpace(version) // Ensure no extra spaces

		if osName != "darwin" || arch != "arm64" || version != "Unknown" {
			t.Errorf("Expected macOS version 13.0, got: %s %s %s", osName, arch, version)
		}
	})

	// ✅ **Test Unknown OS**
	t.Run("UnknownOS", func(t *testing.T) {
		system.RunCommand = MockCommandOutput("", true) // Simulate command failure
		system.GetRuntimeGOOS = func() string { return "solaris" }
		system.GetRuntimeGOARCH = func() string { return "unknown" }

		osName, arch, version := system.GetOSInfo()

		if osName != "solaris" || arch != "unknown" || version != "Unknown" {
			t.Errorf("Expected Unknown version, got: %s %s %s", osName, arch, version)
		}
	})
}
