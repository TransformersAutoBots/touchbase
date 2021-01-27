package touchbasemanager

import (
    "io/ioutil"
    "os"

    "github.com/pkg/errors"
    "github.com/tidwall/gjson"
    "github.com/tidwall/sjson"

    "github.com/autobots/touchbase/configs"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/utils"
    "github.com/autobots/touchbase/validations"
)

const (
    configFileName = "config"
)

func generateConfigDir() error {
    err := utils.Mkdir(getConfigDirPath(), 0766)
    if err != nil {
        return err
    }
    return nil
}

func generateConfigFile() error {
    // E.g: ./.{app_name}/config
    configFile := getConfigFilePathWithExt()

    // Config already exists
    fileInfo, err := os.Stat(configFile)
    if !os.IsNotExist(err) && !fileInfo.IsDir() {
        return errors.Errorf("Config already exists! Please use the config update command to modify the property")
    }

    file, err := os.Create(configFile)
    if err != nil {
        return err
    }

    //goland:noinspection GoNilness
    defer file.Close()
    return nil
}

func CreateConfig(c *types.Config) error {
    if err := generateConfigDir(); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", getConfigDirPath()),
        ).Error("Failed to create the required config dir", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    if err := generateConfigFile(); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", getConfigDirPath()),
            log.Attribute("configFileName", configFileName),
        ).Error("Failed to create the config file")
        return err
    }

    configImpl := configs.New(getConfigFilePathWithExt(), c)
    if err := configImpl.Create(); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", getConfigDirPath()),
            log.Attribute("configFileName", configFileName),
            log.Attribute("configData", c),
        ).Error("Failed to write the data to the config file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }
    return nil
}

func readConfigFile(configFilePath string) ([]byte, error) {
    configData, err := ioutil.ReadFile(configFilePath)
    if err != nil {
        return nil, err
    }
    return configData, nil
}

func checkKeyExists(bytesData []byte, key string) bool {
    return gjson.GetBytes(bytesData, key).Exists()
}

func UpdateConfig(c *ConfigUpdate) error {
    configFilePath := getConfigFilePathWithExt()
    oldConfig, err := readConfigFile(configFilePath)
    if err != nil {
        getLogger().With(
            log.Attribute("configDirPath", getConfigDirPath()),
            log.Attribute("configFileName", configFileName),
        ).Error("Error in reading existing config file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    if !checkKeyExists(oldConfig, c.Key) {
        return errors.Errorf("Key: %s not found in config file", c.Key)
    }

    newConfig, err := sjson.SetBytes(oldConfig, c.Key, c.Value)
    if err != nil {
        getLogger().With(
            log.Attribute("configDirPath", getConfigDirPath()),
            log.Attribute("configFileName", configFileName),
            log.Attribute("input", c),
        ).Error("Error in updating key value", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    updatedConfig := &types.Config{}
    err = utils.UnmarshalJson(newConfig, updatedConfig)
    if err != nil {
        getLogger().Error("Failed to unmarshal json", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }

    // Validate the email address and data file path
    if err := validations.ValidateConfig(updatedConfig); err != nil {
        return err
    }

    configImpl := configs.New(configFilePath, updatedConfig)
    if err := configImpl.Update(); err != nil {
        getLogger().With(
            log.Attribute("configDirPath", getConfigDirPath()),
            log.Attribute("configFileName", configFileName),
            log.Attribute("configData", updatedConfig),
        ).Error("Failed to update the config data to the config file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
        return err
    }
    return nil
}
