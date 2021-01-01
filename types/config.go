package types

type Config struct {
    SpreadsheetID string `json:"spreadsheetid" validate:"required,validateSpreadsheet"`
}
