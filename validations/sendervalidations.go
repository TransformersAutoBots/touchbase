package validations

import (
    "os"

    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    "github.com/pkg/errors"

    "github.com/autobots/touchbase/types"
    "github.com/autobots/touchbase/utils"
)

const (
    // types.Sender struct json validation names
    fileExistsTag = "fileExists"

    // types.Sender struct json keys
    dataFilePath = "data_file"
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
    // return utils.CheckFileExists(fl.Field().String())
}

func newSenderValidator() (ut.Translator, error) {
    trans, err := getUniversalTranslator()
    if err != nil {
        return nil, err
    }

    err = validate.RegisterValidation(fileExistsTag, validateFileExists)
    if err != nil {
        return nil, errors.Errorf("failed to register %s custom validations. Reason: %s", fileExistsTag, err.Error())
    }

    translations := []struct {
        tag         string
        translation string
    }{
        {
            tag:         fileExistsTag,
            translation: "data file does not exists at location",
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

func ValidateSender(sender *types.Sender) error {
    trans, err := newSenderValidator()
    if err != nil {
        return err
    }

    if err := validate.Struct(sender); err != nil {
        errs := err.(validator.ValidationErrors)
        for _, e := range errs {
            if e.Field() == dataFilePath {
                return errors.Errorf("%s %s", e.Translate(trans), e.Value())
            }
            return errors.Errorf("%s. Current Value: %s", e.Translate(trans), e.Value())
        }
    }
    return nil
}
