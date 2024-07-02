package validation

import (
	"github.com/go-playground/validator/v10"
)

// The representation of `validator.Validate` that will be used to validate
// the JSON request bodies provided by the user.
var Validator *validator.Validate

// Validates a struct's values and returns the errors, if they exist.
func Validate[T any](toValidate *T) validator.ValidationErrors {
	Validator = validator.New()

	err := Validator.Struct(toValidate)
	if err != nil {
		return err.(validator.ValidationErrors)
	}

	return nil
}
