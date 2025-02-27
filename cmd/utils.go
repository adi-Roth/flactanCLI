package cmd

import (
	"fmt"
	"os"
	"path/filepath"
)

// getConfigFilePath returns the full path to the config file inside $HOME/.flactan/
func getConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("❌ Failed to get home directory:", err)
		os.Exit(1)
	}

	configDir := filepath.Join(homeDir, ".flactan")
	configFilePath := filepath.Join(configDir, "config.yaml")

	// Ensure the directory exists
	if err := os.MkdirAll(configDir, os.ModePerm); err != nil {
		fmt.Println("❌ Failed to create config directory:", err)
		os.Exit(1)
	}

	return configFilePath
}
