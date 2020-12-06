package cmd

import (
    "github.com/autobots/touchbase/touchbasemanager"
    "github.com/autobots/touchbase/utils"
    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/validations"
)

const (
    user           = "user"
    dataFilePath   = "file-path"
    configFilePath = "config-path"
)

// configInitCmd represents the touchbase config init command
var configInitCmd = &cobra.Command{
    Use:   "init",
    Short: "Initialize configs",
    Long: `The init command will initialize the touchbase application and generate the
necessary config files required for the application to run.`,

    PreRunE: func(cmd *cobra.Command, args []string) error {
        // Initialize Logging
        initLogging(constants.ConsoleFormat, debugMode)

        // Clean the dir/file path
        if err := ensureAbsPath(sender); err != nil {
            return err
        }

        // Validate the email address and data file path
        if err := validations.ValidateSender(sender); err != nil {
            return err
        }
        return nil
    },

    RunE: func(cmd *cobra.Command, args []string) error {
        getLogger().Debug("Initializing application configs... ")
        err := touchbasemanager.CreateConfig(sender)
        if err != nil {
            return err
        }
        return nil
    },
}

func init() {
    configInitCmd.Flags().StringVarP(&sender.User, user, "u", "", "The sender email address")
    _ = configInitCmd.MarkFlagRequired(user)

    configInitCmd.Flags().StringVarP(&sender.DataFile, dataFilePath, "p", "", "The data file path")
    _ = configInitCmd.MarkFlagRequired(dataFilePath)

    configInitCmd.Flags().StringVarP(&sender.ConfigFilePath, configFilePath, "", ".", "The config file path.")
}

func ensureAbsPath(sender *touchbasemanager.Sender) error {
    absPath, err := utils.GetAbsPath(sender.DataFile)
    if err != nil {
        return err
    }
    sender.DataFile = absPath

    absPath, err = utils.GetAbsPath(sender.ConfigFilePath)
    if err != nil {
        return err
    }
    sender.ConfigFilePath = absPath
    return nil
}
