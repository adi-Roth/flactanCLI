package config_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/adi-Roth/flactanCLI/internal/config"
)

func TestGetFilePath(t *testing.T) {
	t.Parallel()

	// Test case: custom directory
	customDir := "/tmp"
	fileName := "test.yaml"
	expected := "/tmp/test.yaml"
	if result := config.GetFilePath(customDir, fileName); result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// Test case: default directory
	customDir = ""
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting user home directory:", err)
		os.Exit(1)
	}
	expected = homeDir + "/.flactancli/test.yaml"
	if result := config.GetFilePath(customDir, fileName); result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}

	// ✅ Test case: Simulate `os.UserHomeDir` failure (capture panic)
	config.GetHomeDir = func() (string, error) {
		return "", errors.New("mocked home directory error")
	}

	// Capture the panic properly
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected function to panic on home directory error, but it did not")
		} else {
			errMsg := r.(string)
			expectedErrMsg := "Error getting user home directory: mocked home directory error"
			if errMsg != expectedErrMsg {
				t.Errorf("Expected panic message '%s', but got '%s'", expectedErrMsg, errMsg)
			}
		}
	}()

	_ = config.GetFilePath("", fileName) // This should trigger panic
}
