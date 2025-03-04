package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/adi-Roth/flactanCLI/internal/config"
	"github.com/adi-Roth/flactanCLI/internal/system"
	"github.com/adi-Roth/flactanCLI/internal/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage FlactanCLI configuration",
}

func init() {
	configCmd.AddCommand(initCmd)
	configCmd.AddCommand(showCmd)
	configCmd.AddCommand(editCmd)
	configCmd.AddCommand(resetCmd)
	configCmd.AddCommand(addCmd)
	configCmd.AddCommand(deleteCmd)
	RootCmd.AddCommand(configCmd)
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show FlactanCLI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetFilePath("", "config.yaml")
		data, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}
		fmt.Println(string(data))
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the configuration file in the system editor",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetFilePath("", "config.yaml")
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano"
		}
		if err := exec.Command(editor, configPath).Run(); err != nil {
			fmt.Println("Error opening editor:", err)
		}
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration to system-detected values",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetFilePath("", "config.yaml")
		osName, osArch, osVersion := system.GetOSInfo()
		internet := "disconnected"
		if system.CheckInternet() {
			internet = "connected"
		}

		defaultConfig := map[string]interface{}{
			"os-name":    osName,
			"os-arch":    osArch,
			"os-version": osVersion,
			"internet":   internet,
			"tools-path": filepath.Join(filepath.Dir(configPath), "tools.yaml"),
			"global-settings": map[string]interface{}{
				"sources": map[string]string{},
			},
		}
		data, _ := yaml.Marshal(defaultConfig)
		if err := os.WriteFile(configPath, data, 0644); err != nil {
			fmt.Println("Error writing config file:", err)
		}
		fmt.Println("Configuration reset to system-detected values.")
	},
}

var addCmd = &cobra.Command{
	Use:   "add source <key=value>",
	Short: "Add a new source to global settings",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetFilePath("", "config.yaml")
		data, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		var config config.Config
		if err := yaml.Unmarshal(data, &config); err != nil {
			fmt.Println("Error parsing YAML:", err)
			return
		}

		parts := strings.SplitN(args[0], "=", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format. Use key=value.")
			return
		}
		config.GlobalSettings.Sources[parts[0]] = parts[1]

		updatedData, _ := yaml.Marshal(config)
		if err := os.WriteFile(configPath, updatedData, 0644); err != nil {
			fmt.Println("Error writing config file:", err)
		}
		fmt.Println("Source added successfully!")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete source <key>",
	Short: "Delete a source from global settings",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetFilePath("", "config.yaml")
		data, err := os.ReadFile(configPath)
		if err != nil {
			fmt.Println("Error reading config file:", err)
			return
		}

		var config config.Config
		if err := yaml.Unmarshal(data, &config); err != nil {
			fmt.Println("Error parsing YAML:", err)
			return
		}

		delete(config.GlobalSettings.Sources, args[0])

		updatedData, _ := yaml.Marshal(config)
		if err := os.WriteFile(configPath, updatedData, 0644); err != nil {
			fmt.Println("Error writing config file:", err)
		}
		fmt.Println("Source deleted successfully!")
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize FlactanCLI and create configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		RunConfigInit(utils.OSFileSystem{}, "") // Use real filesystem in production
	},
}

// InitializeConfig now takes a FileSystem interface (for mocking in tests)
func RunConfigInit(fs utils.FileSystem, customDir string) error {
	configDir := config.GetFilePath(customDir, "")
	configPath := config.GetFilePath(customDir, "config.yaml")
	toolsPath := config.GetFilePath(customDir, "tools.yaml")

	// Ensure the directory exists
	if err := fs.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	// Get system information
	osName, osArch, osVersion := system.GetOSInfo()
	isConnected := system.CheckInternet()
	internetStatus := "offline"
	if isConnected {
		internetStatus = "connected"
	}

	// Create config file
	config := config.Config{
		OSName:    osName,
		OSArch:    osArch,
		OSVersion: osVersion,
		Internet:  internetStatus,
		ToolsPath: toolsPath,
		GlobalSettings: config.GlobalSettings{
			Sources: make(map[string]string), // ✅ Initialize empty map to prevent nil panic
		},
	}

	configData, err := yaml.Marshal(&config)
	if err != nil {
		return fmt.Errorf("error marshaling config data: %w", err)
	}

	if err := fs.WriteFile(configPath, configData, 0644); err != nil {
		return fmt.Errorf("error writing config.yaml: %w", err)
	}

	fmt.Println("Configuration saved:", configPath)

	// Create an empty tools.yaml file if it doesn't exist
	_, err = fs.ReadFile(toolsPath)
	if err != nil { // If file doesn't exist, write it
		if err := fs.WriteFile(toolsPath, []byte{}, 0644); err != nil {
			return fmt.Errorf("error creating tools.yaml: %w", err)
		}
		fmt.Println("Tools configuration initialized:", toolsPath)
	}

	return nil
}
