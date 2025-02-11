package config

import (
	"github.com/markusressel/fan2go/cmd/global"
	"github.com/markusressel/fan2go/internal/configuration"
	"github.com/markusressel/fan2go/internal/ui"
	"github.com/spf13/cobra"
	"os"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates the current configuration",
	Long:  ``,
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		var configPath string
		if len(global.CfgFile) <= 0 {
			configPath = configuration.DetectConfigFile()
		} else {
			configPath = global.CfgFile
		}

		ui.Info("Using configuration file at: %s", configPath)
		configuration.LoadConfig()

		if err := configuration.Validate(configPath); err != nil {
			ui.Error("Validation failed: %v", err)
			os.Exit(1)
		}

		ui.Success("Config looks good! :)")
		return nil
	},
}

func init() {
	Command.AddCommand(validateCmd)
}
