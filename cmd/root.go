package cmd

import (
	"fmt"

	"github.com/adi-Roth/flactanCLI/internal/utils"
	"github.com/spf13/cobra"
)

// Global variables
var logLevel string // CLI flag for log level

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "flactancli",
	Short: "FlactanCLI - Automate workstation setup with ease",
	Long: `
FlactanCLI is a cross-platform command-line tool designed to automate
the setup and configuration of development workstations in both online 
and offline environments. It simplifies software installation, network 
setup, and system configurations.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		utils.InitLogger(logLevel) // Initialize logger before anything runs
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
Welcome to FlactanCLI!
----------------------
Run 'flactancli --help' for more information.
		`)
	},
}

// Execute runs the root command
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		utils.Logger.Debug("Cobra error: ", err) // Keep it for debugging but not for users
		utils.ExitHandler(1, nil)                // Exit without re-logging the error
	}
}

// Override the default completion command
var completionCmd = &cobra.Command{
	Use:    "completion",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		// Disable the default completion command
	},
}

// Init function for RootCmd
func init() {
	RootCmd.AddCommand(completionCmd)
	RootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Set the log level (debug, info, warn, error)")
}
