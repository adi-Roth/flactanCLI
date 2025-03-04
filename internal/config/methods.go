package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Define a variable for getting the home directory (mockable in tests)
var GetHomeDir = os.UserHomeDir

func GetFilePath(customDir string, fileName string) string {
	var filePath string
	if customDir != "" {
		filePath = filepath.Join(customDir, fileName) // Ensure full path
	} else {
		homeDir, err := GetHomeDir()
		if err != nil {
			panic(fmt.Sprintf("Error getting user home directory: %v", err))
		}
		filePath = filepath.Join(homeDir, ".flactancli", fileName)
	}
	return filePath
}
