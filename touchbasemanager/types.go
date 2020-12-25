package touchbasemanager

type Config struct {
    User     string `json:"user" validate:"required,email"`
    Password string `json:"password" validate:"required"`
    DataFile string `json:"data_file" validate:"required,fileExists"`
    Dir      string `json:"dir" validate:"dirExists"`
}

type ConfigUpdate struct {
    Key   string `json:"key"`
    Value string `json:"value"`
}
