package configs

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"

    "github.com/pkg/errors"

    "github.com/autobots/touchbase/constants"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/utils"
)

const (
    configPath     = "./.%s/"
    configFileName = "config"
)

func generateConfigDir() error {
    err := utils.Mkdir(fmt.Sprintf(configPath, constants.AppName), 0766)
    if err != nil {
        return err
    }
    return nil
}

func generateConfigFile() error {
    // E.g: ./.{app_name}/config
    configFile := fmt.Sprintf(configPath+"%s", constants.AppName, configFileName)
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

func CreateConfig(sender *types.Sender) error {
    if err := generateConfigDir(); err != nil {
        getLogger().With(
            log.Attribute("configPath", fmt.Sprintf(configPath, constants.AppName)),
        ).Error("Failed to create the required config dir")
        return err
    }

    if err := generateConfigFile(); err != nil {
        getLogger().With(
            log.Attribute("configPath", fmt.Sprintf(configPath, constants.AppName)),
            log.Attribute("configFileName", configFileName),
        ).Error("Failed to create the config file")
        return err
    }
    writeFile(sender)
    return nil
}

func writeFile(sender *types.Sender) {
    configFile := fmt.Sprintf(configPath+"%s", constants.AppName, configFileName)
    r, err := json.MarshalIndent(sender, constants.JsonPrefix, constants.JsonIntend)
    err = ioutil.WriteFile(configFile, r, 0644)
    if err != nil {
        return
    }
}
