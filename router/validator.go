package router

import "gopkg.in/go-playground/validator.v9"

// NewValidator return an instance of Validator
func NewValidator() *Validator {
	return &Validator{
		validator: validator.New(),
	}
}

// Validator type
type Validator struct {
	validator *validator.Validate
}

// Validate and argument
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}
