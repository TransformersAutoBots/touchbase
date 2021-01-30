package configs

import (
    "github.com/spf13/viper"

    "github.com/autobots/touchbase/constants"
    log "github.com/autobots/touchbase/logger"
    "github.com/autobots/touchbase/types"
)

var (
    Config = types.Config{}
)

func ParseConfig(configDir, configFileName string) {
    // Set the file name of the configurations file
    viper.SetConfigName(configFileName)
    // Set the path to look for the configurations file
    viper.AddConfigPath(configDir)
    viper.SetConfigType(constants.JsonFormat)

    // Reads all the configuration variables
    if err := viper.ReadInConfig(); err != nil {
        getLogger().Fatal("Error in reading config file", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }

    // Map configs to model
    if err := viper.Unmarshal(&Config); err != nil {
        getLogger().Fatal("Enable to decode config into struct", log.TouchBaseError(&types.Log{
            Reason: err.Error(),
        }))
    }
}
