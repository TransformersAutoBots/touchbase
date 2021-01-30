package email

import (
    "google.golang.org/api/gmail/v1"
)

type recruiter struct {
    FullName string `json:"full_name"`
    EmailID  string `json:"email_id"`
    Company  string `json:"company"`
}

type user struct {
    FullName      string `json:"full_name"`
    EmailID       string `json:"email_id"`
    Resume        string `json:"resume"`
    IntroTemplate string `json:"intro_template"`
}

type email struct {
    Recruiter recruiter `json:"recruiter"`
    User      user      `json:"user"`

    Subject  string `json:"subject"`
    Boundary string `json:"boundary"`
}

type Email interface {
    To() string
    From() string
    EmailSubject() string
    MIMEBoundary() string

    EmailBody() string

    ResumeFileName() string
    MIMEType(data []byte) string
    Attachment(fileBytes []byte) string

    ComposeMessage(messageBody []byte) *gmail.Message
}
