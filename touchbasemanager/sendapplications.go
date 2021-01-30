package touchbasemanager

import (
    "fmt"

    "github.com/autobots/touchbase/email"
    "github.com/autobots/touchbase/utils"
)

func getRecruiterFullName(recruiterFirstName, recruiterLastName string) string {
    if utils.IsEmptyString(recruiterFirstName) || utils.IsEmptyString(recruiterLastName) {
        return fmt.Sprintf("%s%s", recruiterFirstName, recruiterLastName)
    }
    return fmt.Sprintf("%s %s", recruiterFirstName, recruiterLastName)
}

func (a *application) send() {
    getLogger().Debug("Initiating connection to reach out to recruiters/manager... ")
    for i := a.StartRow - 1; i < a.EndRow; i++ {
        recruiterName := getRecruiterFullName(a.Company.Recruiters.Values[i][0].(string), a.Company.Recruiters.Values[i][1].(string))

        newEmail := email.New(recruiterName, a.Company.Recruiters.Values[i][2].(string), a.Company.Name, a.Subject, getConfigDirEnvVar())
        err := newEmail.SendWithAttachments()
        if err != nil {
            // Log and ignore. Reach out to the next recruiter/manager
            continue
        }
    }
}
