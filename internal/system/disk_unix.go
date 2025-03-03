//go:build !windows

package system

import (
	"syscall"
)

// CheckDiskSpace verifies if the system has at least `requiredGB` of free disk space.
func CheckDiskSpace(requiredGB uint64) bool {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		return false
	}

	freeSpace := (stat.Bavail * uint64(stat.Bsize)) / (1024 * 1024 * 1024) // Convert to GB
	return freeSpace >= requiredGB
}
