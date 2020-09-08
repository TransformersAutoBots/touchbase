package validations

import (
    "reflect"
    "strings"

    "github.com/go-playground/locales/en"
    ut "github.com/go-playground/universal-translator"
    "github.com/go-playground/validator/v10"
    entranslations "github.com/go-playground/validator/v10/translations/en"
    "github.com/pkg/errors"

    "github.com/autobots/touchbase/constants"
)

var validate *validator.Validate

func InitValidator() {
    validate = validator.New()
    validate.RegisterTagNameFunc(func(field reflect.StructField) string {
        jsonTag := strings.SplitN(field.Tag.Get(constants.JsonFormat), constants.CommaSeparator, 2)[0]
        if jsonTag == "-" {
            return ""
        }
        return jsonTag
    })
}

func getUniversalTranslator() (ut.Translator, error) {
    enLocaleTranslator := en.New()
    uni := ut.New(enLocaleTranslator, enLocaleTranslator)
    trans, _ := uni.GetTranslator("en")
    err := entranslations.RegisterDefaultTranslations(validate, trans)
    if err != nil {
        return nil, errors.Errorf("failed to register default translations. Reason: %s", err.Error())
    }
    return trans, nil
}

func registrationFunc(tag string, translation string) validator.RegisterTranslationsFunc {
    return func(ut ut.Translator) (err error) {
        if err = ut.Add(tag, translation, true); err != nil {
            return
        }
        return
    }
}

func translationFunc(ut ut.Translator, fe validator.FieldError) string {
    t, err := ut.T(fe.Tag(), reflect.ValueOf(fe.Value()).String(), fe.Param())
    if err != nil {
        return fe.(error).Error()
    }
    return t
}
