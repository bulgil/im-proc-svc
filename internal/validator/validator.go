package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	val *validator.Validate
}

func New() *Validator {
	return &Validator{
		val: validator.New(validator.WithRequiredStructEnabled()),
	}
}

func (v *Validator) Validate(data any) error {
	return v.val.Struct(data)
}
