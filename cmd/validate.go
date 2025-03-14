package cmd

import (
	"fmt"

	"github.com/adi-Roth/flactanCLI/internal/validation"
	"github.com/spf13/cobra"
)

// Flag for specifying a custom config directory
var configDir string

// ValidateCmd represents the validate command
var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate the system and configuration files",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("unknown subcommand: %s", args[0])
		}
		return cmd.Help()
	},
	Run: func(cmd *cobra.Command, args []string) {
		validation.RunValidation(configDir)
	},
}

func init() {
	RootCmd.AddCommand(ValidateCmd)
	ValidateCmd.Flags().StringVarP(&configDir, "config-dir", "c", "", "Specify a custom config directory")
}
