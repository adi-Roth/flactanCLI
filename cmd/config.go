package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/adi-Roth/flactanCLI/internal/config"
	"github.com/adi-Roth/flactanCLI/internal/utils"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage FlactanCLI configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			return fmt.Errorf("unknown subcommand: %s", args[0])
		}
		return cmd.Help()
	},
}

func init() {
	configCmd.AddCommand(initCmd, showCmd, editCmd, resetCmd, addCmd)
	RootCmd.AddCommand(configCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize FlactanCLI and create configuration files",
	Run: func(cmd *cobra.Command, args []string) {
		err := config.InitializeConfig(utils.OSFileSystem{}, "")
		if err != nil {
			fmt.Println("Error initializing configuration:", err)
		}
	},
}

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show FlactanCLI configuration",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetFilePath("", "config.yaml")
		fs := utils.OSFileSystem{}

		cfg, err := config.ReadConfig(fs, configPath)
		if err != nil {
			fmt.Println("Error reading config:", err)
			return
		}

		data, _ := yaml.Marshal(cfg)
		fmt.Println(string(data))
	},
}

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Open the configuration file in the system editor",
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetFilePath("", "config.yaml")
		editor := os.Getenv("EDITOR")
		if editor == "" {
			editor = "nano"
		}
		if err := exec.Command(editor, configPath).Run(); err != nil {
			fmt.Println("Error opening editor:", err)
		}
	},
}

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the configuration to system-detected values",
	Run: func(cmd *cobra.Command, args []string) {
		fs := utils.OSFileSystem{}

		err := config.InitializeConfig(fs, "")
		if err != nil {
			fmt.Println("Error resetting configuration:", err)
		} else {
			fmt.Println("Configuration reset to system-detected values.")
		}
	},
}

var addCmd = &cobra.Command{
	Use:   "add source <key=value>",
	Short: "Add a new source to global settings",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		configPath := config.GetFilePath("", "config.yaml")
		fs := utils.OSFileSystem{}

		cfg, err := config.ReadConfig(fs, configPath)
		if err != nil {
			fmt.Println("Error reading config:", err)
			return
		}

		parts := strings.SplitN(args[0], "=", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid format. Use key=value.")
			return
		}
		cfg.GlobalSettings.Sources[parts[0]] = parts[1]

		err = config.WriteConfig(fs, configPath, *cfg)
		if err != nil {
			fmt.Println("Error saving configuration:", err)
		} else {
			fmt.Println("Source added successfully!")
		}
	},
}
