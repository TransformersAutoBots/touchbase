package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/touchbasemanager"
    "github.com/autobots/touchbase/utils"
    "github.com/autobots/touchbase/validations"
)

const (
    spreadSheet   = "spreadsheet"
    configDirPath = "config-dir"
)

// configInitCmd represents the touchbase config init command
var configInitCmd = &cobra.Command{
    Use:   "init",
    Short: "Initialize configs",
    Long: fmt.Sprintf(`%sThe init command will initialize the touchbase application and generate 
the necessary config files required for the application to run.`, generateBanner(constants.AppName)),

    PreRunE: func(cmd *cobra.Command, args []string) error {
        // Initialize Logging
        initLogging(constants.ConsoleFormat, debugMode)

        if err := ensureAbsPath(configs); err != nil {
            return err
        }

        if err := validations.ValidateGoogleApplicationCredentials(os.Getenv(constants.GoogleApplicationCredentials)); err != nil {
            return err
        }

        // Validate the email address and data file path
        if err := validations.ValidateConfig(configs); err != nil {
            return err
        }
        return nil
    },

    RunE: func(cmd *cobra.Command, args []string) error {
        getLogger().Debug("Initializing application configs... ")
        err := touchbasemanager.CreateConfig(configs)
        if err != nil {
            return err
        }
        return nil
    },
}

func init() {
    configInitCmd.Flags().StringVar(&configs.SpreadsheetID, spreadSheet, "", "The Google spreadsheet id")
    _ = configInitCmd.MarkFlagRequired(spreadSheet)

    configInitCmd.Flags().StringVar(&configs.Dir, configDirPath, ".", "The config dir path")
}

func ensureAbsPath(config *touchbasemanager.Config) error {
    absPath, err := utils.GetAbsPath(config.Dir)
    if err != nil {
        return err
    }
    config.Dir = absPath
    return nil
}
