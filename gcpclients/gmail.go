package gcpclients

import (
    "google.golang.org/api/gmail/v1"
)

func SendEmail(userID string, message *gmail.Message) (*gmail.Message, error) {
    gmailService := Gmail()
    sentMessage, err := gmailService.Users.Messages.Send(userID, message).Do()
    if err != nil {
        return nil, err
    }
    return sentMessage, nil
}
