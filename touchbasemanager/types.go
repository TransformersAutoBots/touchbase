package touchbasemanager

type Config struct {
    SpreadsheetID string `json:"spreadsheet_id" validate:"required,validateSpreadsheet"`
    Dir           string `json:"dir" validate:"required,validateDir"`
}

type configToSave struct {
    SpreadsheetID string `deepcopier:"field:SpreadsheetID"`
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
