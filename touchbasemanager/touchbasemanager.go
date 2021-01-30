package touchbasemanager

import (
    "fmt"

    "google.golang.org/api/sheets/v4"

    "github.com/autobots/touchbase/configs"
    "github.com/autobots/touchbase/gcpclients"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
)

// initConfig reads the config file specified using constants.TouchBaseConfigDir
// env variable.
func initConfig() {
    configs.ParseConfig(getConfigDirPath(), configFileName)
}

func defaultReadRange(sheetName string) string {
    return fmt.Sprintf("%s!A1:C", sheetName)
}

func retrieveCompanyDetails(companyName string) (*sheets.ValueRange, error) {
    getLogger().Debug("Retrieving the recruiters details from the company sheet")

    companyDetails, err := gcpclients.RetrieveSheetData(configs.Config.SpreadsheetID, defaultReadRange(companyName))
    if err != nil {
        return nil, err
    }

    getLogger().Debug("Successfully retrieved the recruiters details from the company sheet")
    return companyDetails, nil
}

func ReachOutRecruiters(subject *string) error {
    getLogger().Debug("Initializing config... ")
    initConfig()

    // Prompt all the company names in the Google Spread Sheet
    companyName, err := selectCompany()
    if err != nil {
        getLogger().Error("Failed to select a company for applications", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }
    log.AddCompanyName(companyName, logger)

    companyDetails, err := retrieveCompanyDetails(companyName)
    if err != nil {
        getLogger().Error("Failed to retrieve the recruiters details from the company sheet", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    start, end, err := getLimits(companyDetails)
    if err != nil {
        getLogger().Error("Failed to get start and end row to retrieve recruiters emails from the sheet", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    application := application{
        Company: company{
            Name:       companyName,
            Recruiters: companyDetails,
        },
        StartRow: start,
        EndRow:   end,
        Subject:  *subject,
    }

    application.send()
    return nil
}
