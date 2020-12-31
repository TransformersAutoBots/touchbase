package cmd

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/utils"
    "github.com/autobots/touchbase/validations"
)

// configCmd represents the touchbase config command
var configCmd = &cobra.Command{
    Use:   "config",
    Short: "Config required for touchbase application",
    Long: fmt.Sprintf(`%sThe config command along with its subcommand will initialize the necessary 
config files required for the touchbase application to run.`, generateBanner(constants.AppName)),

    PreRun: func(cmd *cobra.Command, args []string) {
        // Initialize Logging
        initLogging(constants.ConsoleFormat, debugMode)
    },

    Run: func(cmd *cobra.Command, args []string) {
        getLogger().Info(`Config command should be used with one of its subcommand. For list of subcommands run "touchbase config --help"`)
    },
}

func init() {
    // Add sub commands
    configCmd.AddCommand(configInitCmd)
    configCmd.AddCommand(configUpdateCmd)
}

func validateEnvVars() error {
    if err := validations.ValidateGoogleApplicationCredentials(utils.GetEnv(constants.GoogleApplicationCredentials)); err != nil {
        return err
    }

    if err := validations.ValidateConfigDir(utils.GetEnv(constants.TouchBaseConfigDir)); err != nil {
        return err
    }
    return nil
}
