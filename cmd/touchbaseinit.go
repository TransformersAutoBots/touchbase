package cmd

import (
    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/configs"
    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/validations"
)

const (
    user         = "user"
    dataFilePath = "file-path"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
    Use:   "init",
    Short: "Initialize the touchbase application",
    Long: `The init command will initialize the touchbase application and generate the
necessary config files required for the application to run.`,

    PreRunE: func(cmd *cobra.Command, args []string) error {
        // Initialize Logging
        initLogging(constants.ConsoleFormat, debugMode)

        // Validate the email address and data file path
        err := validations.ValidateSender(sender)
        if err != nil {
            return err
        }
        return nil
    },

    RunE: func(cmd *cobra.Command, args []string) error {
        getLogger().Debug("Initializing application configs... ")
        err := configs.CreateConfig(sender)
        if err != nil {
            return err
        }
        return nil
    },
}

func init() {
    initCmd.Flags().StringVarP(&sender.User, user, "u", "", "The sender email address")
    _ = initCmd.MarkFlagRequired(user)

    initCmd.Flags().StringVarP(&sender.DataFile, dataFilePath, "p", "", "The data file path")
    _ = initCmd.MarkFlagRequired(dataFilePath)
}
