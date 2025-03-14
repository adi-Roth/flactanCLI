package config

import (
	"fmt"

	"github.com/adi-Roth/flactanCLI/internal/system"
	"github.com/adi-Roth/flactanCLI/internal/utils"
	"gopkg.in/yaml.v2"
)

// InitializeConfig creates config.yaml and tools.yaml if not exists
func InitializeConfig(fs utils.FileSystem, customDir string) error {
	configDir := GetFilePath(customDir, "")
	configPath := GetFilePath(customDir, "config.yaml")
	toolsPath := GetFilePath(customDir, "tools.yaml")

	// Ensure config directory exists
	if err := fs.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	// Fetch system details
	osName, osArch, osVersion := system.GetOSInfo()
	internetStatus := "offline"
	if system.CheckInternet() {
		internetStatus = "connected"
	}

	// Create default config
	defaultConfig := Config{
		OSName:    osName,
		OSArch:    osArch,
		OSVersion: osVersion,
		Internet:  internetStatus,
		ToolsPath: toolsPath,
		GlobalSettings: GlobalSettings{
			Sources: make(map[string]string),
		},
	}

	// Write config.yaml
	if err := writeYAML(fs, configPath, defaultConfig); err != nil {
		return fmt.Errorf("error writing config.yaml: %w", err)
	}
	fmt.Println("Configuration saved:", configPath)

	// Ensure tools.yaml exists
	if !fs.FileExists(toolsPath) {
		if err := fs.WriteFile(toolsPath, []byte{}, 0644); err != nil {
			return fmt.Errorf("error creating tools.yaml: %w", err)
		}
		fmt.Println("Tools configuration initialized:", toolsPath)
	}

	return nil
}

// ReadConfig reads the config.yaml file into a Config struct
func ReadConfig(fs utils.FileSystem, path string) (*Config, error) {
	data, err := fs.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("error parsing YAML: %w", err)
	}
	return &config, nil
}

// WriteConfig saves the Config struct to the config.yaml file
func WriteConfig(fs utils.FileSystem, path string, config Config) error {
	return writeYAML(fs, path, config)
}

// Helper to write YAML files
func writeYAML(fs utils.FileSystem, path string, data interface{}) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshaling YAML: %w", err)
	}
	return fs.WriteFile(path, yamlData, 0644)
}
