package touchbasemanager

import (
    "github.com/autobots/touchbase/email"
)

func (a *application) send() {
    getLogger().Debug("Initiating connection to reach out to recruiters/manager... ")
    for i := a.StartRow - 1; i < a.EndRow; i++ {
        newEmail := email.New(a.Company.Recruiters.Values[i][0].(string), a.Company.Recruiters.Values[i][1].(string), a.Company.Recruiters.Values[i][2].(string), a.Company.Name, getConfigDirEnvVar())

        err := newEmail.SendWithAttachments()
        if err != nil {
            // Log and ignore. Reach out to the next recruiter/manager
            continue
        }
    }
}
