package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adi-Roth/flactanCLI/internal/system"
	"github.com/adi-Roth/flactanCLI/internal/utils"
	"github.com/spf13/cobra"
)

// ValidateCmd represents the validate command
var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the system and configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		RunValidation("")
	},
}

func init() {
	RootCmd.AddCommand(ValidateCmd)
}

// RunValidation performs system and config validation
func RunValidation(configDir string) {
	fmt.Printf("🔍 Running validation checks (config path: %s)...\n", configDir)
	if configDir == "" {
		homeDir, _ := os.UserHomeDir()
		configDir = filepath.Join(homeDir, ".flactancli")
	}

	configFile := filepath.Join(configDir, "config.yaml")
	toolsFile := filepath.Join(configDir, "tools.yaml")

	fmt.Println("🔍 Running validation checks...")

	// Check OS Compatibility
	osName, arch, osVersion := system.GetOSInfo()
	fmt.Printf("✔ OS Detected: %s (%s) - %s ✅\n", osName, arch, osVersion)

	// Check Internet Connectivity
	if system.CheckInternet() {
		fmt.Println("✔ Internet Access: Connected ✅")
	} else {
		fmt.Println("✖ Internet Access: Not Connected ❌")
	}

	// Check Admin Privileges
	if system.CheckAdminPrivileges() {
		fmt.Println("✔ Sudo Privileges: Available ✅")
	} else {
		fmt.Println("✖ Sudo Privileges: Not Available ❌")
	}

	// Check Disk Space
	requiredSpace := uint64(50) // Minimum 50GB free space
	if system.CheckDiskSpace(requiredSpace) {
		fmt.Printf("✔ Disk Space: %dGB Free ✅\n", requiredSpace)
	} else {
		fmt.Printf("✖ Disk Space: Less than %dGB Free ❌\n", requiredSpace)
	}

	fs := utils.OSFileSystem{}

	if fs.FileExists(configFile) {
		if fs.ValidateYAML(configFile) {
			fmt.Println("✔ config.yaml is valid ✅")
		} else {
			fmt.Println("✖ config.yaml is invalid ❌")
		}
	} else {
		fmt.Println("✖ config.yaml is missing ❌")
	}

	if fs.FileExists(toolsFile) {
		if fs.ValidateYAML(toolsFile) {
			fmt.Println("✔ tools.yaml is valid ✅")
		} else {
			fmt.Println("✖ tools.yaml is invalid ❌")
		}
	} else {
		fmt.Println("✖ tools.yaml is missing ❌")
	}

	fmt.Println("\n✅ Validation Complete!")
}
