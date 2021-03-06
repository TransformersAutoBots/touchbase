package configs

import (
    "io/ioutil"

    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/utils"
)

func New(path string, data interface{}) ConfigImpl {
    return &config{
        FilePath: path,
        Data:     data,
    }
}

func (c *config) writeJsonToFile() error {
    data, err := utils.PrettyJson(c.Data)
    if err != nil {
        getLogger().With(
            log.Attribute("data", c.Data),
        ).Error("Failed to marshal and indent json data", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    err = ioutil.WriteFile(c.FilePath, data, 0644)
    if err != nil {
        return err
    }
    return nil
}

func (c *config) Create() error {
    err := c.writeJsonToFile()
    if err != nil {
        return err
    }
    return nil
}

func (c *config) Update() error {
    err := c.writeJsonToFile()
    if err != nil {
        return err
    }
    return nil
}
