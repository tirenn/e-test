package utils

import (
	"errors"
	"fmt"
	"tirenn/test-efishery/auth-app/models"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func InitValidator() {
	validate = validator.New()
	validate.RegisterValidation("roles", validateRoles)
}

func ValidateVarWithValue(f interface{}, o interface{}, tag string) error {
	err := validate.VarWithValue(fmt.Sprintf("%v", f), fmt.Sprintf("%v", o), tag)
	if err != nil {
		message := fmt.Sprintf("Error:Field validation for %v and %v failed on the '"+tag+"' tag", o, f)
		return errors.New(message)
	}

	return nil
}

func ValidateStruct(s interface{}) error {
	return validate.Struct(s)
}

func validateRoles(fl validator.FieldLevel) bool {
	s := fl.Field().String()

	if len(s) == 0 {
		return false
	}

	result := false
	for _, v := range models.Roles {
		if s == v {
			result = true
		}
	}
	return result
}
