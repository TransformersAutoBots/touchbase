package email

import (
    "encoding/base64"
    "fmt"
    "net/http"
    "path/filepath"

    "google.golang.org/api/gmail/v1"

    "github.com/autobots/touchbase/configs"
    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/utils"
)

func getRecruiterFullName(recruiterFirstName, recruiterLastName string) string {
    if utils.IsEmptyString(recruiterFirstName) || utils.IsEmptyString(recruiterLastName) {
        return fmt.Sprintf("%s%s", recruiterFirstName, recruiterLastName)
    }
    return fmt.Sprintf("%s %s", recruiterFirstName, recruiterLastName)
}

func New(recruiterFirstName, recruiterLastName, recruiterEmailID, recruiterCompanyName, templateDirPath string) *email {
    return &email{
        Recruiter: recruiter{
            FullName:    getRecruiterFullName(recruiterFirstName, recruiterLastName),
            EmailID:     recruiterEmailID,
            CompanyName: recruiterCompanyName,
        },
        User: user{
            FullName:      configs.Config.User.FullName,
            EmailID:       configs.Config.User.EmailID,
            Resume:        configs.Config.User.Resume,
            IntroTemplate: templateDirPath,
        },
    }
}

func (e *email) To() string {
    return e.Recruiter.EmailID
}

func (e *email) From() string {
    return fmt.Sprintf("%s<%s>", e.User.FullName, e.User.EmailID)
}

func (e *email) Subject() string {
    return "Application for Software Engineer/Software Developer"
}

func (e *email) MIMEBoundary() string {
    if utils.IsEmptyString(e.Boundary) {
        e.Boundary = randomString(32, constants.AlphaNumericType)
    }
    return e.Boundary
}

func (e *email) EmailBody() (string, error) {
    return e.parseTemplates()
}

func (e *email) ResumeFileName() string {
    _, resumeFileName := filepath.Split(e.User.Resume)
    return resumeFileName
}

func (e *email) MIMEType(data []byte) string {
    return http.DetectContentType(data)
}

func (e *email) Attachment(fileBytes []byte) string {
    return base64.StdEncoding.EncodeToString(fileBytes)
}

func (e *email) ComposeMessage(messageBody []byte) *gmail.Message {
    message := &gmail.Message{}
    message.Raw = base64.URLEncoding.EncodeToString(messageBody)
    return message
}
