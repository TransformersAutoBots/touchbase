package types

type Sender struct {
    User         string `json:"user" validate:"email"`
    Password     string `json:"password,omitempty"`
    DataFile     string `json:"data_file" validate:"required,fileExists"`
}
