package cmd

import (
    "fmt"

    "github.com/spf13/cobra"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/touchbasemanager"
    "github.com/autobots/touchbase/utils"
    "github.com/autobots/touchbase/validations"
)

// TODO:
// subject from user input
// Add test email capability using a flag

const (
    emailSubject = "subject"
)

var (
    subject string
)

var reachOutCmd = &cobra.Command{
    Use:   "reach-out",
    Short: "Reach out to the company recruiters/managers",
    Long: fmt.Sprintf(`%sThe apply-command will send emails to recuiters/managers of the company provided
in the data sheet.`, generateBanner(constants.AppName)),

    PreRunE: func(cmd *cobra.Command, args []string) error {
        if err := validateEnvVars(); err != nil {
            return err
        }

        if err := validateIntroduceHtmlFileExists(); err != nil {
            return err
        }
        return nil
    },

    RunE: func(cmd *cobra.Command, args []string) error {
        err := touchbasemanager.ReachOutRecruiters(&subject)
        if err != nil {
            return err
        }
        return nil
    },
}

func init() {
    reachOutCmd.Flags().StringVarP(&subject, emailSubject, "", "Application for Software Engineer/Software Developer", "The subject of the email to be sent")
}

func validateIntroduceHtmlFileExists() error {
    introduceHtmlFilePath := fmt.Sprintf("%s/%s%s", utils.GetEnv(constants.TouchBaseConfigDir), constants.IntroduceTemplateName, constants.DotHtml)
    if err := validations.ValidateIntroduceHtmlFileExists(introduceHtmlFilePath); err != nil {
        return err
    }
    return nil
}
