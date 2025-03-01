package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adi-Roth/flactanCLI/internal/system"
	"github.com/adi-Roth/flactanCLI/internal/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// Config structure for storing system information
type Config struct {
	OSName    string `yaml:"os-name"`
	OSVersion string `yaml:"os-version"`
	Internet  string `yaml:"internet"`
	ToolsPath string `yaml:"tools-path"`
}

// InitializeConfig now takes a FileSystem interface (for mocking in tests)
func InitializeConfig(fs utils.FileSystem, customDir string) {
	// Use customDir in tests, otherwise use $HOME/.flactancli
	var configDir string
	if customDir != "" {
		configDir = customDir
	} else {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Println("Error getting user home directory:", err)
			os.Exit(1)
		}
		configDir = filepath.Join(homeDir, ".flactancli")
	}

	configPath := filepath.Join(configDir, "config.yaml")
	toolsPath := filepath.Join(configDir, "tools.yaml")

	// Ensure the directory exists
	if err := fs.MkdirAll(configDir, 0755); err != nil {
		fmt.Println("Error creating config directory:", err)
		os.Exit(1)
	}

	// Get system information
	osName, osVersion := system.GetOSInfo()
	isConnected := system.CheckInternet()
	internetStatus := "offline"
	if isConnected {
		internetStatus = "connected"
	}

	// Create config file
	config := Config{
		OSName:    osName,
		OSVersion: osVersion,
		Internet:  internetStatus,
		ToolsPath: toolsPath,
	}
	configData, err := yaml.Marshal(&config)
	if err != nil {
		fmt.Println("Error marshaling config data:", err)
		os.Exit(1)
	}

	if err := fs.WriteFile(configPath, configData, 0644); err != nil {
		fmt.Println("Error writing config.yaml:", err)
		os.Exit(1)
	}

	fmt.Println("Configuration saved:", configPath)

	// Create an empty tools.yaml file if it doesn't exist
	_, err = fs.ReadFile(toolsPath)
	if err != nil { // If file doesn't exist, write it
		if err := fs.WriteFile(toolsPath, []byte{}, 0644); err != nil {
			fmt.Println("Error creating tools.yaml:", err)
			os.Exit(1)
		}
		fmt.Println("Tools configuration initialized:", toolsPath)
	}
}

// initCmd represents the "init" command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize FlactanCLI and create configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		InitializeConfig(utils.OSFileSystem{}, "") // Use real filesystem in production
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
