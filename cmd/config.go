package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/adi-Roth/flactanCLI/internal/config"
	"github.com/adi-Roth/flactanCLI/internal/system"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage FlactanCLI configuration",
}

func init() {
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
		configPath := getConfigPath()
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
		configPath := getConfigPath()
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano"
		}
		exec.Command(editor, configPath).Run()
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration to system-detected values",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := getConfigPath()
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
		os.WriteFile(configPath, data, 0644)
		fmt.Println("Configuration reset to system-detected values.")
	},
}

var addCmd = &cobra.Command{
	Use:   "add source <key=value>",
	Short: "Add a new source to global settings",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := getConfigPath()
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
		os.WriteFile(configPath, updatedData, 0644)
		fmt.Println("Source added successfully!")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete source <key>",
	Short: "Delete a source from global settings",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := getConfigPath()
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
		os.WriteFile(configPath, updatedData, 0644)
		fmt.Println("Source deleted successfully!")
	},
}

func getConfigPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".flactancli", "config.yaml")
}
