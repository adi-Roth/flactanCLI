package system

import (
	"os/exec"
	"runtime"
	"strings"
)

// Allow overriding for testing
var GetRuntimeGOOS = func() string { return runtime.GOOS }
var GetRuntimeGOARCH = func() string { return runtime.GOARCH }

// Define function variables for mocking exec.Command in tests
var RunCommand = func(name string, args ...string) ([]byte, error) {
	return exec.Command(name, args...).Output()
}

// GetOSInfo returns the OS name, architecture, and version
func GetOSInfo() (string, string, string) {
	osName := GetRuntimeGOOS()
	arch := GetRuntimeGOARCH()
	var version string

	switch osName {
	case "linux":
		// Try lsb_release first
		out, err := RunCommand("lsb_release", "-d")
		if err == nil {
			version = strings.TrimSpace(strings.Split(string(out), ":")[1])
		} else {
			// Fallback: Read /etc/os-release
			out, err = RunCommand("cat", "/etc/os-release")
			if err == nil {
				lines := strings.SplitSeq(string(out), "\n")
				for line := range lines {
					if strings.HasPrefix(line, "PRETTY_NAME=") {
						version = strings.Trim(strings.Split(line, "=")[1], "\"")
						break
					}
				}
			}
		}

	case "darwin":
		// macOS: Use sw_vers
		out, err := RunCommand("sw_vers", "-productVersion")
		if err == nil {
			version = strings.TrimSpace(string(out))
		}

	case "windows":
		// Windows: Use wmic os get Caption (more reliable than "ver")
		out, err := RunCommand("wmic", "os", "get", "Caption")
		if err == nil {
			lines := strings.Split(string(out), "\n")
			if len(lines) > 1 {
				version = strings.TrimSpace(lines[1])
			}
		}
	}

	// Default to "Unknown" if version retrieval fails
	if version == "" {
		version = "Unknown"
	}

	return osName, arch, version
}
