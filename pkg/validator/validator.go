package validator

import "github.com/go-playground/validator/v10"

type ValidatorInterface interface {
}

type Validator struct {
	validator *validator.Validate
}

func NewValidator() (ValidatorInterface, error) {
	v := validator.New()
	return &Validator{validator: v}, nil
}
