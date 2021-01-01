package cmd

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/touchbasemanager"
    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/validations"
)

const (
    spreadSheet   = "spreadsheet"
)

var (
    config = &types.Config{}
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

        if err := validateEnvVars(); err != nil {
            return err
        }

        // Validate the email address and data file path
        if err := validations.ValidateConfig(config); err != nil {
            return err
        }
        return nil
    },

    RunE: func(cmd *cobra.Command, args []string) error {
        getLogger().Debug("Initializing application configs... ")
        err := touchbasemanager.CreateConfig(config)
        if err != nil {
            return err
        }
        return nil
    },
}

func init() {
    configInitCmd.Flags().StringVar(&config.SpreadsheetID, spreadSheet, "", "The Google spreadsheet id")
    _ = configInitCmd.MarkFlagRequired(spreadSheet)
}
