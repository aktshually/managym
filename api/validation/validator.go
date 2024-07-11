package validation

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// The representation of `validator.Validate` that will be used to validate
// the JSON request bodies provided by the user.
var Validator *validator.Validate = validator.New()

// Validates a struct's values and returns the errors, if they exist.
func Validate[T any](toValidate T) []string {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(Validator, trans)

	Validator.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} must have a value", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.StructField())
		return t
	})
	Validator.RegisterTranslation("gte", trans, func(ut ut.Translator) error {
		return ut.Add("gte", "\"{0}\" must be greater than or equal {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("gte", fe.StructField(), fe.Value().(string))
		return t
	})

	err := Validator.Struct(toValidate)
	if err != nil {
		var formattedErrors []string
		errs := err.(validator.ValidationErrors)

		for _, e := range errs {
			formattedErrors = append(formattedErrors, e.Translate(trans))
		}

		return formattedErrors
	}

	return []string{}
}
