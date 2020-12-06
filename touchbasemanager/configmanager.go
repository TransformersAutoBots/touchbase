package touchbasemanager

import (
    "fmt"
    "os"

    "github.com/autobots/touchbase/configs"
    "github.com/autobots/touchbase/constants"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/utils"
    "github.com/pkg/errors"
)

const (
    configFileName = "config"
)

func (s *Sender) getConfigDirAbsolutePath() string {
    return fmt.Sprintf("%s/.%s", s.ConfigFilePath, constants.AppName)
}

func (s *Sender) generateConfigDir() error {
    err := utils.Mkdir(s.getConfigDirAbsolutePath(), 0766)
    if err != nil {
        return err
    }
    return nil
}

func (s *Sender) getConfigFileAbsolutePath() string {
    return fmt.Sprintf("%s/%s", s.getConfigDirAbsolutePath(), configFileName)
}

func (s *Sender) generateConfigFile() error {
    // E.g: ./.{app_name}/config
    configFile := s.getConfigFileAbsolutePath()
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

func CreateConfig(s *Sender) error {
    if err := s.generateConfigDir(); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", s.getConfigDirAbsolutePath()),
        ).Error("Failed to create the required config dir", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    if err := s.generateConfigFile(); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", s.getConfigDirAbsolutePath()),
            log.Attribute("configFileName", configFileName),
        ).Error("Failed to create the config file")
        return err
    }

    configData, err := utils.PrettyJson(s)
    if err != nil {
        getLogger().With(
            log.Attribute("sender", s),
        ).Error("Failed to marshal and indent json data", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    if err := configs.CreateConfig(s.getConfigFileAbsolutePath(), configData); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", s.getConfigDirAbsolutePath()),
            log.Attribute("configFileName", configFileName),
            log.Attribute("configData", s),
        ).Error("Failed to write the data to the config file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }
    return nil
}
