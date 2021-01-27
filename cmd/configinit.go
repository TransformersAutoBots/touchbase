package cmd

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/touchbasemanager"
    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/utils"
    "github.com/autobots/touchbase/validations"
)

const (
    spreadsheetID = "spreadsheet-id"
    fullName      = "full-name"
    email         = "email"
    resume        = "resume"
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
        if err := validateEnvVars(); err != nil {
            return err
        }

        resumeAbsPath, err := utils.GetAbsPath(config.User.Resume)
        if err != nil {
            return err
        }
        config.User.Resume = resumeAbsPath

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
    configInitCmd.Flags().StringVar(&config.SpreadsheetID, spreadsheetID, "", "The Google spreadsheet id")
    _ = configInitCmd.MarkFlagRequired(spreadsheetID)

    configInitCmd.Flags().StringVar(&config.User.FullName, fullName, "", "The user full name")
    _ = configInitCmd.MarkFlagRequired(fullName)

    configInitCmd.Flags().StringVar(&config.User.EmailID, email, "", "The user email id (Must be a gmail account)")
    _ = configInitCmd.MarkFlagRequired(email)

    configInitCmd.Flags().StringVar(&config.User.Resume, resume, "", "The user resume file location along with file name")
    _ = configInitCmd.MarkFlagRequired(resume)
}
