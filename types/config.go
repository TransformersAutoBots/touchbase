package types

type user struct {
    FullName string `json:"fullname" validate:"required"`
    EmailID  string `json:"emailid" validate:"required,email"`
    Resume   string `json:"resume" validate:"required,validateResumeFile"`
}

type Config struct {
    SpreadsheetID string `json:"spreadsheetid" validate:"required,validateSpreadsheet"`
    User          user   `json:"user"`
}
