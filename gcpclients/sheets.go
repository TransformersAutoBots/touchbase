package gcpclients

import (
    "google.golang.org/api/sheets/v4"
)

func RetrieveSpreadsheet(spreadsheetID string) (*sheets.Spreadsheet, error) {
    sheetsService := Sheets()
    spreadsheet, err := sheetsService.Spreadsheets.Get(spreadsheetID).Do()
    if err != nil {
        return nil, err
    }
    return spreadsheet, nil
}

func RetrieveSheetsInSpreadsheet(spreadsheetID string) ([]*sheets.Sheet, error) {
    spreadsheet, err := RetrieveSpreadsheet(spreadsheetID)
    if err != nil {
        return nil, err
    }
    return spreadsheet.Sheets, nil
}

func RetrieveSheetData(spreadsheetID, readRange string) (*sheets.ValueRange, error) {
    sheetsService := Sheets()
    sheetData, err := sheetsService.Spreadsheets.Values.Get(spreadsheetID, readRange).Do()
    if err != nil {
        return nil, err
    }
    return sheetData, nil
}
