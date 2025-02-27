package cmd

import (
	"fmt"
	"net"
	"os"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFilePath string

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize CLI with system properties",
	Long: `The init command detects OS properties, checks internet access,
and saves this data into a configuration file at $HOME/.flactan/config.yaml.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the config file path
		configFilePath = getConfigFilePath()

		// Collect system information
		systemInfo := map[string]string{
			"os":      runtime.GOOS,
			"arch":    runtime.GOARCH,
			"version": runtime.Version(),
			"online":  fmt.Sprintf("%v", checkInternetAccess()),
		}

		// Save config
		saveConfig(systemInfo)

		fmt.Println("✅ Initialization complete! System info saved to", configFilePath)
	},
}

func init() {
	RootCmd.AddCommand(initCmd)
}


// checkInternetAccess checks if the system has internet access
func checkInternetAccess() bool {
	_, err := net.LookupHost("google.com")
	return err == nil
}

// saveConfig writes the system info to a YAML file
func saveConfig(data map[string]string) {
	viper.SetConfigFile(configFilePath)
	viper.SetConfigType("yaml")

	for key, value := range data {
		viper.Set(key, value)
	}

	err := viper.WriteConfigAs(configFilePath)
	if err != nil {
		fmt.Println("⚠️ Failed to write config:", err)
		os.Exit(1)
	}
}
