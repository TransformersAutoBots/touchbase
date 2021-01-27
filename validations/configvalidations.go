package validations

import (
    "fmt"
    "path/filepath"
    "reflect"

    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    "github.com/pkg/errors"

    "github.com/autobots/touchbase/gcpclients"
    "github.com/autobots/touchbase/types"
)

const (
    pdfFormat = ".pdf"

    // touchbasemanager.Config struct json validation names
    validateSpreadsheetTag = "validateSpreadsheet"
    validateResumeFileTag  = "validateResumeFile"
)

func spreadsheetTranslationFunc(ut ut.Translator, fe validator.FieldError) string {
    t, err := ut.T(fe.Tag(), fe.StructField(), reflect.ValueOf(fe.Value()).String())
    if err != nil {
        return err.Error()
    }
    return t
}

func resumeFileTranslationFunc(ut ut.Translator, fe validator.FieldError) string {
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

func isValidResumeFile(fl validator.FieldLevel) bool {
    return validateFileExists(fl.Field().String()) && filepath.Ext(fl.Field().String()) == pdfFormat
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

    err = validate.RegisterValidation(validateResumeFileTag, isValidResumeFile)
    if err != nil {
        return nil, customValidationError(validateResumeFileTag, err.Error())
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
        {
            tag:             validateResumeFileTag,
            translation:     fmt.Sprintf("{0}: {1} is not a valid resume file. Either the file does not exists or it is not a %s file", pdfFormat),
            translationFunc: resumeFileTranslationFunc,
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

func ValidateConfig(config *types.Config) error {
    trans, err := newConfigValidator()
    if err != nil {
        return err
    }

    if err := validate.Struct(config); err != nil {
        errs := err.(validator.ValidationErrors)
        for _, e := range errs {
            return errors.New(e.Translate(trans))
        }
    }
    return nil
}
