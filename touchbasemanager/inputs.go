package touchbasemanager

import (
    "fmt"
    "strconv"
    "strings"

    "github.com/manifoldco/promptui"
    "github.com/pkg/errors"
    "google.golang.org/api/sheets/v4"

    "github.com/autobots/touchbase/configs"
    "github.com/autobots/touchbase/gcpclients"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
)

const (
    initSelectLabel = "Select a company from the list"
    pointerUnicode  = "\U0001F449"
    activeColor     = "cyan"
    inactiveColor   = "cyan"
    selectColor     = "green"
    backgroundColor = "cyan"
)

func populateCompanyList(spreadsheetID string) ([]string, error) {
    getLogger().With(
        log.Attribute("spreadsheetID", spreadsheetID),
    ).Debug("Retrieving list of companies in the spreadsheet... ")
    var companyNames []string
    sheetsInSpreadsheet, err := gcpclients.RetrieveSheetsInSpreadsheet(spreadsheetID)
    if err != nil {
        return nil, err
    }

    for _, sheet := range sheetsInSpreadsheet {
        companyNames = append(companyNames, sheet.Properties.Title)
    }
    getLogger().With(
        log.Attribute("spreadsheetID", spreadsheetID),
    ).Debug("Successfully retrieved list of companies from the spreadsheet")
    return companyNames, nil
}

func selectCompany() (string, error) {
    selectTemplate := &promptui.SelectTemplates{
        Label:    "{{ . }}?",
        Active:   fmt.Sprintf("%s {{ . | %s }}", pointerUnicode, activeColor),
        Inactive: fmt.Sprintf("  {{ . | %s }}", inactiveColor),
        Selected: fmt.Sprintf("%s {{ . | %s | %s }}", pointerUnicode, selectColor, backgroundColor),
    }

    companiesList, err := populateCompanyList(configs.Config.SpreadsheetID)
    if err != nil {
        getLogger().With(
            log.Attribute("spreadsheetID", configs.Config.SpreadsheetID),
        ).Error("Error in retrieving list of companies from spreadsheet", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return "", err
    }

    searcher := func(input string, index int) bool {
        name := strings.Replace(strings.ToLower(companiesList[index]), " ", "", -1)
        input = strings.Replace(strings.ToLower(input), " ", "", -1)
        return strings.Contains(name, input)
    }

    selectPrompt := promptui.Select{
        Label:     initSelectLabel,
        Items:     companiesList,
        Templates: selectTemplate,
        Size:      10,
        Searcher:  searcher,
    }

    _, companyName, err := selectPrompt.Run()
    if err != nil {
        getLogger().Error("Error in selecting the company from the list", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return "", err
    }
    return companyName, nil
}

func getLimits(companyDetails *sheets.ValueRange) (int64, int64, error) {
    getLogger().Debug("Retrieving the start row to send emails from the company sheet")
    prompt := promptui.Prompt{
        Label: "Start Row (Start from 2 to ignore the header line)",
        Validate: func(input string) error {
            startNum, err := strconv.ParseInt(input, 10, 64)
            if err != nil {
                return errors.New("Not a valid number")
            }
            if startNum < 2 {
                return errors.New("Start row cannot be less than 2")
            }
            return nil
        },
    }

    min, err := prompt.Run()
    if err != nil {
        return 0, 0, err
    }

    start, err := strconv.ParseInt(min, 10, 64)
    if err != nil {
        return 0, 0, err
    }

    getLogger().Debug("Retrieving the end row to send emails from the company sheet")
    prompt = promptui.Prompt{
        Label: fmt.Sprintf("End Row (Max allowed: %d)", len(companyDetails.Values)),
        Validate: func(input string) error {
            endNum, err := strconv.ParseInt(input, 10, 64)
            if err != nil {
                return errors.New("Not a valid number")
            }
            if endNum > int64(len(companyDetails.Values)) {
                return errors.Errorf("End row cannot be greater than %d", len(companyDetails.Values))
            } else if endNum < start {
                return errors.New("End row cannot be smaller than start row")
            }
            return nil
        },
    }

    max, err := prompt.Run()
    if err != nil {
        return 0, 0, err
    }

    end, err := strconv.ParseInt(max, 10, 64)
    if err != nil {
        return 0, 0, err
    }

    getLogger().Debug("Successfully retrieved the start and end row to send emails from the company sheet")
    return start, end, nil
}
