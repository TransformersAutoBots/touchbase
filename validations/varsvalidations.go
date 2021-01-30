package validations

import (
    "os"

    "github.com/pkg/errors"

    "github.com/autobots/touchbase/constants"
    "github.com/autobots/touchbase/utils"
)

func validateFileExists(filePath string) bool {
    fileInfo, err := os.Stat(filePath)
    if err != nil && os.IsNotExist(err) {
        return false
    }
    return !fileInfo.IsDir()
}

func ValidateGoogleApplicationCredentials(path string) error {
    if utils.IsEmptyString(path) {
        return errors.Errorf(`Env var %s is empty or missing`, constants.GoogleApplicationCredentials)
    }

    absPath, err := utils.GetAbsPath(path)
    if err != nil {
        return err
    }

    if !validateFileExists(absPath) {
        return errors.Errorf("%s is not a valid file path", absPath)
    }
    return nil
}

func validateDirPath(path string) bool {
    if utils.IsEmptyString(path) {
        return false
    }
    fileInfo, err := os.Stat(path)
    if err != nil && os.IsNotExist(err) {
        return false
    }
    return fileInfo.IsDir()
}

func ValidateConfigDir(configDirPath string) error {
    if utils.IsEmptyString(configDirPath) {
        return errors.Errorf(`Env var %s is empty or missing`, constants.TouchBaseConfigDir)
    }

    absPath, err := utils.GetAbsPath(configDirPath)
    if err != nil {
        return err
    }

    if !validateDirPath(absPath) {
        return errors.Errorf("%s is not a valid config dir path", absPath)
    }
    return nil
}

func ValidateIntroduceHtmlFileExists(introduceHtmlFilePath string) error {
    absPath, err := utils.GetAbsPath(introduceHtmlFilePath)
    if err != nil {
        return err
    }

    if !validateFileExists(absPath) {
        return errors.Errorf(`introduce.html file not found at location %s`, absPath)
    }
    return nil
}
