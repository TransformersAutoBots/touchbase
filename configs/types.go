package configs

type ConfigImpl interface {
	Create() error
	Update() error
}

type Config struct {
	FilePath string      `json:"file_path"`
	Data     interface{} `json:"data"`
}
