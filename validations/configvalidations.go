package validations

import (
    lg "log"
    "reflect"

    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    "github.com/pkg/errors"

    "github.com/autobots/touchbase/gcpclients"
    "github.com/autobots/touchbase/touchbasemanager"
)

const (
    // touchbasemanager.Config struct json validation names
    validateSpreadsheetTag = "validateSpreadsheet"
)

func spreadsheetTranslationFunc(ut ut.Translator, fe validator.FieldError) string {
    t, err := ut.T(fe.Tag(), fe.StructField(), reflect.ValueOf(fe.Value()).String())
    if err != nil {
        return err.Error()
    }
    return t
}

func isValidSpreadsheet(fl validator.FieldLevel) bool {
    _, err := gcpclients.RetrieveSpreadsheet(fl.Field().String())
    return err == nil
}

func newConfigValidator() (ut.Translator, error) {
    trans, err := getUniversalTranslator()
    if err != nil {
        return nil, err
    }

    err = validate.RegisterValidation(validateSpreadsheetTag, isValidSpreadsheet)
    if err != nil {
        return nil, customValidationError(validateSpreadsheetTag, err.Error())
    }

    translations := []struct {
        tag             string
        translation     string
        translationFunc validator.TranslationFunc
    }{
        {
            tag:             validateSpreadsheetTag,
            translation:     "{0}: {1} is not valid",
            translationFunc: spreadsheetTranslationFunc,
        },
    }

    for _, t := range translations {
        err = validate.RegisterTranslation(t.tag, trans, registrationFunc(t.tag, t.translation), t.translationFunc)
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
            lg.Println(e.Translate(trans))
            return errors.New(e.Translate(trans))
        }
    }
    return nil
}
