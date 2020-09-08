package types

// Log model to represent error in touchbase with error code, error message
// and reason.
type Log struct {
    ErrorCode    string `json:"error_code,omitempty"`
    ErrorMessage string `json:"error_message,omitempty"`
    Reason       string `json:"reason,omitempty"`
}
