package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/adi-Roth/flactanCLI/internal/system"
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

// initCmd represents the "init" command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize FlactanCLI and create configuration files",
	Long: `The "init" command collects system information, checks internet connectivity, 
and sets up necessary configuration files in $HOME/.flactancli/.`,
	Run: func(cmd *cobra.Command, args []string) {
		initializeConfig()
	},
}

// initializeConfig performs the setup tasks
func initializeConfig() {
	// Get OS details
	osName, osVersion := system.GetOSInfo()
	fmt.Printf("Detected OS: %s %s\n", osName, osVersion)

	// Check internet connectivity
	isConnected := system.CheckInternet()
	internetStatus := "offline"
	if isConnected {
		internetStatus = "connected"
	}
	fmt.Printf("Internet Status: %s\n", internetStatus)

	// Define config file path
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		os.Exit(1)
	}
	configDir := filepath.Join(homeDir, ".flactancli")
	configPath := filepath.Join(configDir, "config.yaml")
	toolsPath := filepath.Join(configDir, "tools.yaml")

	// Ensure the directory exists
	if err := os.MkdirAll(configDir, 0755); err != nil {
		fmt.Println("Error creating config directory:", err)
		os.Exit(1)
	}

	// Create config.yaml file
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

	if err := os.WriteFile(configPath, configData, 0644); err != nil {
		fmt.Println("Error writing config.yaml:", err)
		os.Exit(1)
	}

	fmt.Println("Configuration saved:", configPath)

	// Create an empty tools.yaml file if it doesn't exist
	if _, err := os.Stat(toolsPath); os.IsNotExist(err) {
		if err := os.WriteFile(toolsPath, []byte{}, 0644); err != nil {
			fmt.Println("Error creating tools.yaml:", err)
			os.Exit(1)
		}
		fmt.Println("Tools configuration initialized:", toolsPath)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
