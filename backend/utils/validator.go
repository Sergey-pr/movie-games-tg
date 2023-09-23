package utils

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func NewValidator() *validator.Validate {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	return v
}

type ValidateErrors []ValidateError

func (ValidateErrors) Error() string { return "validation errors" }

type ValidateError struct {
	Field string `json:"field"`
	Err   string `json:"error"`
}

func (o ValidateError) Error() string {
	return fmt.Sprintf("%s: %s", o.Field, o.Err)
}

func ValidateStruct(v *validator.Validate, obj interface{}) ValidateErrors {
	var validateErrors ValidateErrors

	if err := v.Struct(obj); err != nil {
		switch err.(type) {
		case validator.ValidationErrors:
			for _, err := range err.(validator.ValidationErrors) {
				validateErrors = append(validateErrors, ValidateError{err.Field(), err.Tag()})
			}
		case *validator.InvalidValidationError:
			validateErrors = append(validateErrors, ValidateError{err.Error(), "Struct"})
		}

	}

	return validateErrors
}
