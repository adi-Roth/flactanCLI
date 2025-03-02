package system

import (
	"os"
	"os/exec"
	"runtime"
)

// CheckAdminPrivileges returns true if the user has admin/sudo rights
func CheckAdminPrivileges() bool {
	switch runtime.GOOS {
	case "windows":
		// Windows: Check if running as Administrator
		out, err := exec.Command("net", "session").Output()
		return err == nil && len(out) > 0
	case "linux", "darwin":
		// Unix-based: Check if running as root
		return os.Geteuid() == 0
	default:
		return false
	}
}
