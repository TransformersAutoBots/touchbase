package touchbasemanager

type Config struct {
    SpreadsheetID string `json:"spreadsheetid" validate:"required,validateSpreadsheet"`
}

type ConfigUpdate struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

type Apply struct {
    CompanyName string `json:"company_name"`
    StartRow    int64  `json:"start_row"`
    EndRow      int64  `json:"end_row"`
}
