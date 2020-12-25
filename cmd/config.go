package cmd

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/touchbasemanager"
)

var (
    configs       = &touchbasemanager.Config{}
    configsUpdate = &touchbasemanager.ConfigUpdate{}
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
