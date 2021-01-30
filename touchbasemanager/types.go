package touchbasemanager

import (
    "google.golang.org/api/sheets/v4"
)

type ConfigUpdate struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}

type company struct {
    Name       string             `json:"name"`
    Recruiters *sheets.ValueRange `json:"recruiters"`
}

type application struct {
    Company company `json:"company"`

    StartRow int64  `json:"start_row"`
    EndRow   int64  `json:"end_row"`
    Subject  string `json:"subject"`
}
