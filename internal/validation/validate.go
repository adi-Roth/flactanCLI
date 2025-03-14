package validation

import (
	"github.com/adi-Roth/flactanCLI/internal/config"
	"github.com/adi-Roth/flactanCLI/internal/system"
	"github.com/adi-Roth/flactanCLI/internal/utils"
)

// RunValidation performs system and config validation
func RunValidation(configDir string) {
	fs := utils.OSFileSystem{}
	// get config files path
	configFile := config.GetFilePath(configDir, "config.yaml")
	toolsFile := config.GetFilePath(configDir, "tools.yaml")

	utils.Logger.Info("🔍 Running validation checks...")

	// Check OS Compatibility
	osName, arch, osVersion := system.GetOSInfo()
	utils.Logger.Infof("✔ OS Detected: %s (%s) - %s ✅", osName, arch, osVersion)

	// Check Internet Connectivity
	if system.CheckInternet() {
		utils.Logger.Info("✔ Internet Access: Connected ✅")
	} else {
		utils.Logger.Warn("✖ Internet Access: Not Connected ❌")
	}

	// Check Admin Privileges
	if system.CheckAdminPrivileges() {
		utils.Logger.Info("✔ Sudo Privileges: Available ✅")
	} else {
		utils.Logger.Warn("✖ Sudo Privileges: Not Available ❌")
	}

	// Check Disk Space
	requiredSpace := uint64(10) // Minimum 50GB free space
	if system.CheckDiskSpace(requiredSpace) {
		utils.Logger.Infof("✔ Disk Space: %dGB Free ✅", requiredSpace)
	} else {
		utils.Logger.Warnf("✖ Disk Space: Less than %dGB Free ❌", requiredSpace)
	}

	// Validate config files
	validateFile(fs, configFile, "config.yaml")
	validateFile(fs, toolsFile, "tools.yaml")
}

// validateFile checks if the file exists and is valid
var validateFile = func(fs utils.OSFileSystem, filename string, displayName string) {
	switch {
	case !fs.FileExists(filename):
		utils.Logger.Warnf("✖ %s is missing ❌\n", displayName)
	case !fs.ValidateYAML(filename):
		utils.Logger.Warnf("✖ %s is invalid ❌\n", displayName)
	default:
		utils.Logger.Infof("✔ %s is valid ✅\n", displayName)
	}
}
