package system

import (
	"os/exec"
	"runtime"
	"strings"
)

// GetOSInfo returns the OS name, architecture, and version
func GetOSInfo() (string, string, string) {
	osName := runtime.GOOS
	arch := runtime.GOARCH
	var version string

	switch osName {
	case "linux":
		// Try lsb_release first
		out, err := exec.Command("lsb_release", "-d").Output()
		if err == nil {
			version = strings.TrimSpace(strings.Split(string(out), ":")[1])
		} else {
			// Fallback: Read /etc/os-release
			out, err = exec.Command("cat", "/etc/os-release").Output()
			if err == nil {
				lines := strings.Split(string(out), "\n")
				for _, line := range lines {
					if strings.HasPrefix(line, "PRETTY_NAME=") {
						version = strings.Trim(strings.Split(line, "=")[1], "\"")
						break
					}
				}
			}
		}

	case "darwin":
		// macOS: Use sw_vers
		out, err := exec.Command("sw_vers", "-productVersion").Output()
		if err == nil {
			version = strings.TrimSpace(string(out))
		}

	case "windows":
		// Windows: Use wmic os get Caption (more reliable than "ver")
		out, err := exec.Command("wmic", "os", "get", "Caption").Output()
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
