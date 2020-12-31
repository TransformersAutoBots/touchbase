package gcpclients

import (
    "context"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"

    "golang.org/x/oauth2"
    "golang.org/x/oauth2/google"
    "google.golang.org/api/gmail/v1"
    "google.golang.org/api/sheets/v4"

    "github.com/autobots/touchbase/constants"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/utils"
)

const (
    sheetsTokenFileName = "sheetsToken" + constants.DotSeparator + constants.JsonFormat
    gmailTokenFileName  = "gmailToken" + constants.DotSeparator + constants.JsonFormat
)

func getAppCredentialsFilePath() string {
    return utils.GetEnv(constants.GoogleApplicationCredentials)
}

func getTokenBasePath() string {
    return fmt.Sprintf(".%s", constants.AppName)
}

func getSheetsTokenPath() string {
    return fmt.Sprintf("%s/%s", getTokenBasePath(), sheetsTokenFileName)
}

func getGmailTokenPath() string {
    return fmt.Sprintf("%s/%s", getTokenBasePath(), gmailTokenFileName)
}

// Retrieves a token from a local file.
func tokenFromFile(tokenFile string) (*oauth2.Token, error) {
    f, err := os.Open(tokenFile)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    oauthToken := &oauth2.Token{}
    err = utils.JsonDecoder(f, oauthToken)
    if err != nil {
        return nil, err
    }
    return oauthToken, nil
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
    authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
    fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

    var authCode string
    if _, err := fmt.Scan(&authCode); err != nil {
        getLogger().Fatal("Error in reading authorization code", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }

    oauthToken, err := config.Exchange(context.TODO(), authCode)
    if err != nil {
        getLogger().Fatal("Error in retrieving token from web", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }
    return oauthToken
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) error {
    getLogger().With(
        log.Attribute("location", path),
    ).Debug("Saving credential file to location")

    f, err := os.Create(path)
    if err != nil {
        getLogger().Fatal("Error in caching oauth token", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }

    //goland:noinspection GoNilness
    defer f.Close()

    err = utils.JsonEncoder(f, token)
    if err != nil {
        return err
    }
    return nil
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(tokenFilePath string, config *oauth2.Config) (*http.Client, error) {
    // The file token.json stores the user's access and refresh tokens, and is
    // created automatically when the authorization flow completes for the first
    // time.
    token, err := tokenFromFile(tokenFilePath)
    if err != nil {
        token = getTokenFromWeb(config)
        saveErr := saveToken(tokenFilePath, token)
        if saveErr != nil {
            return nil, saveErr
        }
    }
    return config.Client(context.Background(), token), nil
}

func Sheets() *sheets.Service {
    appCredentialsFilePath := getAppCredentialsFilePath()
    credentialsFile, err := ioutil.ReadFile(appCredentialsFilePath)
    if err != nil {
        getLogger().Fatal("Error in reading the client secret file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }

    // If modifying these scopes, delete your previously saved token.json.
    config, err := google.ConfigFromJSON(credentialsFile, sheets.SpreadsheetsScope)
    if err != nil {
        getLogger().Fatal("Error in parsing the client secret file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }
    client, err := getClient(getSheetsTokenPath(), config)
    if err != nil {
        getLogger().Fatal("Error in getting the client using provided token", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }

    //goland:noinspection GoDeprecation
    sheetsService, err := sheets.New(client)
    if err != nil {
        getLogger().Fatal("Error in retrieving Google Sheets service object", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }
    return sheetsService
}

func Gmail() *gmail.Service {
    appCredentialsFilePath := getAppCredentialsFilePath()
    credentialsFile, err := ioutil.ReadFile(appCredentialsFilePath)
    if err != nil {
        getLogger().Fatal("Error in reading the client secret file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }

    // If modifying these scopes, delete your previously saved token.json.
    config, err := google.ConfigFromJSON(credentialsFile, gmail.GmailSendScope)
    if err != nil {
        getLogger().Fatal("Error in parsing the client secret file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }

    client, err := getClient(getGmailTokenPath(), config)
    if err != nil {
        getLogger().Fatal("Error in getting the client using provided token", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }

    //goland:noinspection GoDeprecation
    gmailService, err := gmail.New(client)
    if err != nil {
        getLogger().Fatal("Error in retrieving Google Gmail service object", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }
    return gmailService
}
