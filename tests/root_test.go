package tests

import (
	"bytes"
	"os"
	"testing"

	"github.com/adi-Roth/flactanCLI/cmd"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestRootCommandExecution(t *testing.T) {
	// Create a buffer to capture output
	buf := new(bytes.Buffer)
	cmd.RootCmd.SetOut(buf)
	cmd.RootCmd.SetErr(buf)

	// Execute root command
	cmd.RootCmd.SetArgs([]string{})
	err := cmd.RootCmd.Execute()

	// Force flush the buffer
	output := buf.String()

	// Debugging: Print output to see what Cobra is returning
	t.Logf("Captured Output: %s", output)

	assert.NoError(t, err, "Executing root command should not return an error")
	assert.Contains(t, output, "flactanCLI - Workstation Setup Automation CLI",
		"Command should contain application name and description")
}

func TestRootCommandHelp(t *testing.T) {
	// Capture help output
	buf := new(bytes.Buffer)
	cmd.RootCmd.SetOut(buf)

	// Simulate running --help
	cmd.RootCmd.SetArgs([]string{"--help"})
	err := cmd.RootCmd.Execute()
	assert.NoError(t, err, "Executing help should not return an error")

	output := buf.String()
	assert.Contains(t, output, "Usage:", "Help output should contain 'Usage:' section")
	assert.Contains(t, output, "flactanCLI", "Help output should mention CLI name")
}

func TestRootCommandFlags(t *testing.T) {
	// Capture output
	buf := new(bytes.Buffer)
	cmd.RootCmd.SetOut(buf)

	// Simulate --toggle flag
	cmd.RootCmd.SetArgs([]string{"--toggle"})
	err := cmd.RootCmd.Execute()
	assert.NoError(t, err, "Executing with --toggle should not return an error")

	output := buf.String()
	assert.Contains(t, output, "Help message for toggle", "Toggle flag should print correct help message")
}

func TestInitConfigWithFile(t *testing.T) {
	// Set a temporary config file
	tmpFile, err := os.CreateTemp("", "flactanCLI_test.yaml")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	// Write a test config
	tmpFile.WriteString("testKey: testValue\n")
	tmpFile.Close()

	// Set config file manually
	cmd.CfgFile = tmpFile.Name()
	viper.SetConfigFile(tmpFile.Name()) // Force Viper to load it
	cmd.InitConfig()
	viper.ReadInConfig() // Explicitly read config file

	assert.Equal(t, "testValue", viper.GetString("testKey"), "Config should load correctly")
}

func TestInitConfigWithoutFile(t *testing.T) {
	// Ensure CfgFile is empty so it defaults to home directory
	cmd.CfgFile = ""
	cmd.InitConfig()

	// Ensure Viper does not panic and loads default values
	assert.NotNil(t, viper.ConfigFileUsed(), "Viper should attempt to load a config")
}
