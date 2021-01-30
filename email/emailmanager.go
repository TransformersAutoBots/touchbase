package email

import (
    "fmt"
    "io/ioutil"

    "github.com/autobots/touchbase/gcpclients"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
)

func (e *email) getMessageWithAttachment(fileBytes []byte) (string, error) {
    emailBody, err := e.EmailBody()
    if err != nil {
        getLogger().Error("Error in getting the email body", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return "", err
    }

    return fmt.Sprintf(`Content-Type: multipart/mixed; boundary=%s
MIME-Version: 1.0
to: %s
from: %s
subject: %s

--%s
Content-Type: text/html; charset="UTF-8"
MIME-Version: 1.0
Content-Transfer-Encoding: 7bit

%s

--%s
Content-Type: %s; name="%s"
MIME-Version: 1.0
Content-Transfer-Encoding: base64
Content-Disposition: attachment; filename="%s"

%s

--%s--`, e.MIMEBoundary(), e.To(), e.From(), e.EmailSubject(), e.MIMEBoundary(), emailBody, e.MIMEBoundary(), e.MIMEType(fileBytes), e.ResumeFileName(), e.ResumeFileName(), chunkSplit(e.Attachment(fileBytes), 76, "\n"), e.MIMEBoundary()), nil
}

func (e *email) SendWithAttachments() error {
    getLogger().With(
        log.Attribute("name", e.Recruiter.FullName),
        log.Attribute("email", e.Recruiter.EmailID),
    ).Debug("Sending email to recruiter/manager... ")
    fileBytes, err := ioutil.ReadFile(e.User.Resume)
    if err != nil {
        return err
    }

    message, err := e.getMessageWithAttachment(fileBytes)
    if err != nil {
        getLogger().With(
            log.Attribute("name", e.Recruiter.FullName),
            log.Attribute("email", e.Recruiter.EmailID),
        ).Error("Failed to get email message to be sent to recruiter/manager", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }
    gmailMessage := e.ComposeMessage([]byte(message))

    sentMessage, err := gcpclients.SendEmail(e.User.EmailID, gmailMessage)
    if err != nil {
        getLogger().With(
            log.Attribute("name", e.Recruiter.FullName),
            log.Attribute("email", e.Recruiter.EmailID),
        ).Error("Failed to send email message to recruiter/manager", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    getLogger().With(
        log.Attribute("name", e.Recruiter.FullName),
        log.Attribute("email", e.Recruiter.EmailID),
        log.Attribute("sentMessageID", sentMessage.Id),
    ).Debug("Successfully sent email to recruiter/manager")
    return nil
}
