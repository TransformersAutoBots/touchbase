package validations

import (
    "os"

    "github.com/autobots/touchbase/touchbasemanager"
    "github.com/autobots/touchbase/utils"

    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    "github.com/pkg/errors"
)

const (
    // touchbasemanager.Config struct json validation names
    fileExistsTag = "fileExists"
    dirExistsTag  = "dirExists"

    // touchbasemanager.Config struct json keys
    dataFilePath  = "data_file"
    configDirPath = "dir"
)

func validateFileExists(fl validator.FieldLevel) bool {
    if utils.IsEmptyString(fl.Field().String()) {
        return false
    }
    fileInfo, err := os.Stat(fl.Field().String())
    if err != nil && os.IsNotExist(err) {
        return false
    }
    return !fileInfo.IsDir()
}

func validatePath(fl validator.FieldLevel) bool {
    if utils.IsEmptyString(fl.Field().String()) {
        return false
    }
    fileInfo, err := os.Stat(fl.Field().String())
    if err != nil && os.IsNotExist(err) {
        return false
    }
    return fileInfo.IsDir()
}

// customValidationError returns the custom validation error message.
func customValidationError(tag, errorMessage string) error {
    return errors.Errorf("failed to register %s custom validation. Reason: %s", tag, errorMessage)
}

func newConfigValidator() (ut.Translator, error) {
    trans, err := getUniversalTranslator()
    if err != nil {
        return nil, err
    }

    err = validate.RegisterValidation(fileExistsTag, validateFileExists)
    if err != nil {
        return nil, customValidationError(fileExistsTag, err.Error())
    }

    err = validate.RegisterValidation(dirExistsTag, validatePath)
    if err != nil {
        return nil, customValidationError(dirExistsTag, err.Error())
    }

    translations := []struct {
        tag         string
        translation string
    }{
        {
            tag:         fileExistsTag,
            translation: "data file does not exists at location",
        },
        {
            tag:         dirExistsTag,
            translation: "invalid config file path",
        },
    }

    for _, t := range translations {
        err = validate.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation), translationFunc)
        if err != nil {
            panic(err)
        }
    }
    return trans, nil
}

func ValidateConfig(config *touchbasemanager.Config) error {
    trans, err := newConfigValidator()
    if err != nil {
        return err
    }

    if err := validate.Struct(config); err != nil {
        errs := err.(validator.ValidationErrors)
        for _, e := range errs {
            if e.Field() == dataFilePath || e.Field() == configDirPath {
                return errors.Errorf("%s %s", e.Translate(trans), e.Value())
            }
            return errors.Errorf("%s. Current Value: %s", e.Translate(trans), e.Value())
        }
    }
    return nil
}
