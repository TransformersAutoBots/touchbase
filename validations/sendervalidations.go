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
    // touchbasemanager.Sender struct json validation names
    fileExistsTag = "fileExists"
    validPathTag  = "validPath"

    // touchbasemanager.Sender struct json keys
    dataFilePath   = "data_file"
    configFilePath = "config_file_path"
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

func newSenderValidator() (ut.Translator, error) {
    trans, err := getUniversalTranslator()
    if err != nil {
        return nil, err
    }

    err = validate.RegisterValidation(fileExistsTag, validateFileExists)
    if err != nil {
        return nil, customValidationError(fileExistsTag, err.Error())
    }

    err = validate.RegisterValidation(validPathTag, validatePath)
    if err != nil {
        return nil, customValidationError(validPathTag, err.Error())
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
            tag:         validPathTag,
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

func ValidateSender(sender *touchbasemanager.Sender) error {
    trans, err := newSenderValidator()
    if err != nil {
        return err
    }

    if err := validate.Struct(sender); err != nil {
        errs := err.(validator.ValidationErrors)
        for _, e := range errs {
            if e.Field() == dataFilePath || e.Field() == configFilePath {
                return errors.Errorf("%s %s", e.Translate(trans), e.Value())
            }
            return errors.Errorf("%s. Current Value: %s", e.Translate(trans), e.Value())
        }
    }
    return nil
}
