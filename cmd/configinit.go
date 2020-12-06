package cmd

import (
    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/touchbasemanager"
    "github.com/autobots/touchbase/utils"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/validations"
)

const (
    user          = "user"
    password      = "password"
    dataFilePath  = "data-file"
    configDirPath = "config-dir"
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

        if err := validations.ValidateEnvKey(constants.TouchBaseToken); err != nil {
            return err
        }

        // Clean the dir/file path
        if err := ensureAbsPath(configs); err != nil {
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
    configInitCmd.Flags().StringVarP(&configs.User, user, "u", "", "The user email address")
    _ = configInitCmd.MarkFlagRequired(user)

    configInitCmd.Flags().StringVarP(&configs.Password, password, "p", "", "The user password")
    _ = configInitCmd.MarkFlagRequired(password)

    configInitCmd.Flags().StringVarP(&configs.DataFile, dataFilePath, "d", "", "The data file path")
    _ = configInitCmd.MarkFlagRequired(dataFilePath)

    configInitCmd.Flags().StringVarP(&configs.Dir, configDirPath, "", ".", "The config dir path")
}

func ensureAbsPath(config *touchbasemanager.Config) error {
    absPath, err := utils.GetAbsPath(config.DataFile)
    if err != nil {
        return err
    }
    config.DataFile = absPath

    absPath, err = utils.GetAbsPath(config.Dir)
    if err != nil {
        return err
    }
    config.Dir = absPath
    return nil
}
