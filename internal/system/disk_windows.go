//go:build windows
// +build windows

package system

import (
	"golang.org/x/sys/windows"
)

// CheckDiskSpace verifies if the system has at least `requiredGB` of free disk space.
func CheckDiskSpace(requiredGB uint64) bool {
	return checkDiskSpaceWindows(requiredGB)
}

// Windows-specific implementation
func checkDiskSpaceWindows(requiredGB uint64) bool {
	var freeBytesAvailable, totalNumberOfBytes, totalNumberOfFreeBytes uint64

	drive, err := windows.UTF16PtrFromString("C:\\")
	if err != nil {
		return false
	}

	err = windows.GetDiskFreeSpaceEx(drive, &freeBytesAvailable, &totalNumberOfBytes, &totalNumberOfFreeBytes)
	if err != nil {
		return false
	}

	freeSpaceGB := totalNumberOfFreeBytes / (1024 * 1024 * 1024) // Convert to GB
	return freeSpaceGB >= requiredGB
}
