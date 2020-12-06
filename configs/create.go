package configs

import (
	"io/ioutil"

	log "github.com/autobots/touchbase/logger"
	"github.com/autobots/touchbase/types"
	"github.com/autobots/touchbase/utils"
)

func New(path string, data interface{}) ConfigImpl {
	return &Config{
		FilePath: path,
		Data:     data,
	}
}

func (c *Config) Create() error {
	configData, err := utils.PrettyJson(c.Data)
	if err != nil {
		getLogger().With(
			log.Attribute("configs", c.Data),
		).Error("Failed to marshal and indent json data", log.TouchBaseError(&types.Log{
			Reason: err.Error(),
		}))
		return err
	}

	err = ioutil.WriteFile(c.FilePath, configData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) Update() error {
	panic("implement me")
}
