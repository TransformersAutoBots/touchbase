package cmd

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/touchbasemanager"
)

const (
    key   = "key"
    value = "value"
)

var (
    configsUpdate = &touchbasemanager.ConfigUpdate{}
)

// configUpdateCmd represents the touchbase config update command
var configUpdateCmd = &cobra.Command{
    Use:   "update",
    Short: "Update configs",
    Long:  fmt.Sprintf(`%sThe update command will update the config property of touchbase application.`, generateBanner(constants.AppName)),

    PreRunE: func(cmd *cobra.Command, args []string) error {
        // Initialize Logging
        initLogging(constants.ConsoleFormat, debugMode)

        if err := validateEnvVars(); err != nil {
            return err
        }

        return nil
    },

    RunE: func(cmd *cobra.Command, args []string) error {
        getLogger().Debug("Updating config... ")
        err := touchbasemanager.UpdateConfig(configsUpdate)
        if err != nil {
            return err
        }
        return nil
    },
}

func init() {
    configUpdateCmd.Flags().StringVarP(&configsUpdate.Key, key, "k", "", "The key to be updated")
    _ = configInitCmd.MarkFlagRequired(key)

    configUpdateCmd.Flags().StringVarP(&configsUpdate.Value, value, "v", "", "The updated config value")
    _ = configInitCmd.MarkFlagRequired(value)
}
