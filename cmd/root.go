package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "flactancli",
	Short: "FlactanCLI - Automate workstation setup with ease",
	Long: `FlactanCLI is a cross-platform command-line tool designed to automate
the setup and configuration of development workstations in both online 
and offline environments. It simplifies software installation, network 
setup, and system configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to FlactanCLI! Use --help to see available commands.")
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This function is called by main.go.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// Override completionCmd to prevent it from appearing
var completionCmd = &cobra.Command{
	Use:    "completion",
	Hidden: true, // Hides the command from --help
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing to effectively disable it
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
