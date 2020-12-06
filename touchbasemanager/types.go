package touchbasemanager

type Sender struct {
    User           string `json:"user" validate:"required,email"`
    Password       string `json:"password,omitempty"`
    DataFile       string `json:"data_file" validate:"required,fileExists"`
    ConfigFilePath string `json:"config_file_path,omitempty" validate:"validPath"`
}
