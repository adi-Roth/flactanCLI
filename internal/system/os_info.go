package system

import (
	"runtime"
)

// GetOSInfo returns the OS name and version
func GetOSInfo() (string, string) {
	return runtime.GOOS, runtime.GOARCH
}
