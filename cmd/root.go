package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// CfgFile stores the path to the configuration file, set via the `--config` flag.
var CfgFile string
var ConfigFilePath string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "flactanCLI",
	Short: "flactanCLI - Workstation Setup Automation CLI",
	Long: `flactanCLI is a CLI tool for automating the setup of a new workstation.
It can install packages, configure settings, and more.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("flactanCLI - Workstation Setup Automation CLI")
	},
}

// Execute runs the root command, initializing all CLI subcommands.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(InitConfig)

	RootCmd.PersistentFlags().StringVar(&CfgFile, "config", "", "config file (default is $HOME/.flactan/config.yaml)")
}

// InitConfig reads in config file and ENV variables if set.
func InitConfig() {
	// Get the config file path
	ConfigFilePath = getConfigFilePath()

	if CfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(CfgFile)
	} else {
		viper.SetConfigFile(ConfigFilePath)
	}

	viper.SetConfigType("yaml")
	viper.AutomaticEnv() // Read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
