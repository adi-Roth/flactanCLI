package config

// Config structure for storing system information
type Config struct {
	OSName         string         `yaml:"os-name"`
	OSArch         string         `yaml:"os-arch"`
	OSVersion      string         `yaml:"os-version"`
	Internet       string         `yaml:"internet"`
	ToolsPath      string         `yaml:"tools-path"`
	GlobalSettings GlobalSettings `yaml:"global-settings"`
}

// GlobalSettings structure for storing global settings
type GlobalSettings struct {
	Sources map[string]string `yaml:"sources"`
}
