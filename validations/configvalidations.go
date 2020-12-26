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
    // touchbasemanager.ConfigInit struct json validation names
    validateSpreadsheetTag = "validateSpreadsheet"
    validateDirTag         = "validateDir"
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
    lg.Println(err)
    return err == nil
}

func dirTranslationFunc(ut ut.Translator, fe validator.FieldError) string {
    t, err := ut.T(fe.Tag(), fe.Field(), reflect.ValueOf(fe.Value()).String())
    if err != nil {
        return err.Error()
    }
    return t
}

func isValidDir(fl validator.FieldLevel) bool {
    return validateDirPath(fl.Field().String())
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

    err = validate.RegisterValidation(validateDirTag, isValidDir)
    if err != nil {
        return nil, customValidationError(validateDirTag, err.Error())
    }

    translations := []struct {
        tag             string
        translation     string
        translationFunc validator.TranslationFunc
    }{
        {
            tag:             validateDirTag,
            translation:     "{0}: {1} is not valid dir",
            translationFunc: dirTranslationFunc,
        },
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

func ValidateConfig(config *touchbasemanager.ConfigInit) error {
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
