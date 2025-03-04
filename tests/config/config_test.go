package config_test

import (
	"testing"

	"github.com/adi-Roth/flactanCLI/internal/config"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

// TestConfigInitialization verifies that the Config struct initializes correctly
func TestConfigInitialization(t *testing.T) {
	cfg := config.Config{
		OSName:    "darwin",
		OSArch:    "arm64",
		OSVersion: "13.0",
		Internet:  "connected",
		ToolsPath: "/Users/test/.flactancli/tools.yaml",
		GlobalSettings: config.GlobalSettings{
			Sources: map[string]string{
				"artifactory": "http://artifactory/",
			},
		},
	}

	assert.Equal(t, "darwin", cfg.OSName)
	assert.Equal(t, "arm64", cfg.OSArch)
	assert.Equal(t, "13.0", cfg.OSVersion)
	assert.Equal(t, "connected", cfg.Internet)
	assert.Equal(t, "/Users/test/.flactancli/tools.yaml", cfg.ToolsPath)
	assert.Equal(t, "http://artifactory/", cfg.GlobalSettings.Sources["artifactory"])
}

// TestYAMLMarshaling ensures Config correctly marshals into YAML
func TestYAMLMarshaling(t *testing.T) {
	cfg := config.Config{
		OSName:    "linux",
		OSArch:    "amd64",
		OSVersion: "20.04",
		Internet:  "disconnected",
		ToolsPath: "/home/test/.flactancli/tools.yaml",
		GlobalSettings: config.GlobalSettings{
			Sources: map[string]string{
				"custom-repo": "http://custom-repo/",
			},
		},
	}

	yamlData, err := yaml.Marshal(&cfg)
	assert.NoError(t, err)
	assert.Contains(t, string(yamlData), "os-name: linux")
	assert.Contains(t, string(yamlData), "internet: disconnected")
	assert.Contains(t, string(yamlData), "custom-repo: http://custom-repo/")
}

// TestYAMLUnmarshaling ensures Config correctly unmarshals from YAML
func TestYAMLUnmarshaling(t *testing.T) {
	yamlData := `
os-name: windows
os-arch: amd64
os-version: "10"
internet: connected
tools-path: "C:/Users/test/.flactancli/tools.yaml"
global-settings:
  sources:
    corporate-proxy: "http://proxy.company.com"
`
	var cfg config.Config
	err := yaml.Unmarshal([]byte(yamlData), &cfg)

	assert.NoError(t, err)
	assert.Equal(t, "windows", cfg.OSName)
	assert.Equal(t, "amd64", cfg.OSArch)
	assert.Equal(t, "10", cfg.OSVersion)
	assert.Equal(t, "connected", cfg.Internet)
	assert.Equal(t, "C:/Users/test/.flactancli/tools.yaml", cfg.ToolsPath)
	assert.Equal(t, "http://proxy.company.com", cfg.GlobalSettings.Sources["corporate-proxy"])
}

// TestUpdateGlobalSettings verifies that sources can be updated
func TestUpdateGlobalSettings(t *testing.T) {
	cfg := config.Config{
		GlobalSettings: config.GlobalSettings{
			Sources: map[string]string{
				"docker-registry": "http://docker.local/",
			},
		},
	}

	// Add new source
	cfg.GlobalSettings.Sources["new-repo"] = "http://new-repo.com"

	// Verify changes
	assert.Equal(t, "http://docker.local/", cfg.GlobalSettings.Sources["docker-registry"])
	assert.Equal(t, "http://new-repo.com", cfg.GlobalSettings.Sources["new-repo"])
}

// TestDeleteGlobalSetting verifies removing a source from global settings
func TestDeleteGlobalSetting(t *testing.T) {
	cfg := config.Config{
		GlobalSettings: config.GlobalSettings{
			Sources: map[string]string{
				"old-repo": "http://old-repo.com",
				"keep-me":  "http://important-repo.com",
			},
		},
	}

	// Delete a source
	delete(cfg.GlobalSettings.Sources, "old-repo")

	// Verify deletion
	assert.NotContains(t, cfg.GlobalSettings.Sources, "old-repo")
	assert.Contains(t, cfg.GlobalSettings.Sources, "keep-me")
}
