package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/touchbasemanager"
    "github.com/autobots/touchbase/validations"
)

const (
    key   = "key"
    value = "value"
)

// configUpdateCmd represents the touchbase config update command
var configUpdateCmd = &cobra.Command{
    Use:   "update",
    Short: "Update configs",
    Long:  fmt.Sprintf(`%sThe update command will update the config property of touchbase application.`, generateBanner(constants.AppName)),

    PreRunE: func(cmd *cobra.Command, args []string) error {
        // Initialize Logging
        initLogging(constants.ConsoleFormat, debugMode)

        if err := validations.ValidateAppToken(constants.TouchBaseToken); err != nil {
            return err
        }

        if err := validations.ValidateEnvConfigDir(constants.TouchBaseConfigDir); err != nil {
            return err
        }
        return nil
    },

    RunE: func(cmd *cobra.Command, args []string) error {
        log.AddConfigUpdateDetails(configsUpdate, logger)
        getLogger().Debug("Updating config... ")
        err := touchbasemanager.UpdateConfig(os.Getenv(constants.TouchBaseConfigDir), configsUpdate)
        if err != nil {
            return err
        }
        return nil
    },
}

func init() {
    configUpdateCmd.Flags().StringVarP(&configsUpdate.Key, key, "k", "", "The config key to be updated")
    _ = configInitCmd.MarkFlagRequired(key)

    configUpdateCmd.Flags().StringVarP(&configsUpdate.Value, value, "v", "", "The updated config value")
    _ = configInitCmd.MarkFlagRequired(value)
}
