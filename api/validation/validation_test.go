package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateSuccess(t *testing.T) {
	assert := assert.New(t)

	type A struct {
		Name string `validate:"required,alphanum,gte=3"`
	}

	toValidate := &A{
		Name: "bob",
	}

	err := Validate(toValidate)
	assert.Len(err, 0)
}

func TestValidateNameDoesNotExist(t *testing.T) {
	assert := assert.New(t)

	type A struct {
		Name string `validate:"required,alphanum,gte=3"`
	}

	toValidate := &A{
		Name: "",
	}

	err := Validate(toValidate)
	assert.Greater(len(err), 0)
}

