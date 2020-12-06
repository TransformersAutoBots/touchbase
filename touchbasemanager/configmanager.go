package touchbasemanager

import (
    "fmt"
    "os"

    "github.com/pkg/errors"

    "github.com/autobots/touchbase/configs"
    "github.com/autobots/touchbase/constants"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/utils"
)

const (
    configFileName = "config"
)

func (c *Config) getConfigDirPath() string {
    return fmt.Sprintf("%s/.%s", c.Dir, constants.AppName)
}

func (c *Config) generateConfigDir() error {
    err := utils.Mkdir(c.getConfigDirPath(), 0766)
    if err != nil {
        return err
    }
    return nil
}

func (c *Config) getConfigFilePath() string {
    return fmt.Sprintf("%s/%s", c.getConfigDirPath(), configFileName)
}

func (c *Config) generateConfigFile() error {
    // E.g: ./.{app_name}/config
    configFile := c.getConfigFilePath()
    fileInfo, err := os.Stat(configFile)
    if !os.IsNotExist(err) && !fileInfo.IsDir() {
        return errors.Errorf("Config already exists! Please use the config update command to modify the property")
    }

    file, err := os.Create(configFile)
    if err != nil {
        return err
    }

    //goland:noinspection ALL
    defer file.Close()
    return nil
}

func CreateConfig(c *Config) error {
    if err := c.generateConfigDir(); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", c.getConfigDirPath()),
        ).Error("Failed to create the required config dir", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    if err := c.generateConfigFile(); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", c.getConfigDirPath()),
            log.Attribute("configFileName", configFileName),
        ).Error("Failed to create the config file")
        return err
    }

    config := configs.New(c.getConfigFilePath(), c)
    if err := config.Create(); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", c.getConfigDirPath()),
            log.Attribute("configFileName", configFileName),
            log.Attribute("configData", c),
        ).Error("Failed to write the data to the config file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }
    return nil
}

func UpdateConfig(c *Config) error {
    return nil
}
