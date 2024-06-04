package utils

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Field  string `json:"field"`
	Reason string `json:"reason"`
}

type ValidationErrors []ValidationError

func (e ValidationErrors) Error() string {
	var errMsg string
	for _, err := range e {
		errMsg += err.Field + ": " + err.Reason + ", "
	}
	return strings.TrimSuffix(errMsg, ", ")
}

func ValidateRequest(req interface{}) error {
	validate := validator.New()

	err := validate.Struct(req)
	if err != nil {
		var valErrors validator.ValidationErrors
		if errors.As(err, &valErrors) {
			var validationErrors ValidationErrors
			for _, e := range valErrors {
				validationError := ValidationError{
					Field:  e.Field(),
					Reason: getValidationReason(e),
				}
				validationErrors = append(validationErrors, validationError)
			}
			return validationErrors
		}
		return fmt.Errorf("invalid request")
	}

	return nil
}

func getValidationReason(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "is required"
	case "min":
		return fmt.Sprintf("must be at least %s characters", e.Param())
	case "max":
		return fmt.Sprintf("must be at most %s characters", e.Param())
	case "email":
		return "must be a valid email address"
	case "url":
		return "must be a valid URL"
	default:
		return fmt.Sprintf("failed on the '%s' tag", e.Tag())
	}
}
