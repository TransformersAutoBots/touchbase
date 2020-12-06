package configs

import (
    "io/ioutil"
)

func CreateConfig(configFile string, configData []byte) error {
    err := ioutil.WriteFile(configFile, configData, 0644)
    if err != nil {
        return err
    }
    return nil
}
