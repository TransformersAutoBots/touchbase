package touchbasemanager

import (
    "fmt"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/utils"
)

func getConfigDirEnvVar() string {
    return utils.GetEnv(constants.TouchBaseConfigDir)
}

func getConfigDirPath() string {
    return fmt.Sprintf("%s/.%s", getConfigDirEnvVar(), constants.AppName)
}

func getConfigFilePath() string {
    return fmt.Sprintf("%s/%s", getConfigDirPath(), configFileName)
}

func getConfigFilePathWithExt() string {
    return fmt.Sprintf("%s.%s", getConfigFilePath(), constants.JsonFormat)
}
