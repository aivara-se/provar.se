package validator

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	validationErrorCode    = "validation_error"
	validationErrorMessage = "The request body is invalid. Please check the errors field for more info."
)

// ValidationError is the response body for validation errors
type ValidationError struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Errors  []string `json:"errors"`
}

// newValidationError returns a new ErrorResponse for the given error
func newValidationError(errs validator.ValidationErrors) *ValidationError {
	errorMessages := make([]string, len(errs))
	for i, err := range errs {
		errorMessages[i] = err.Error()
	}
	return &ValidationError{
		Code:    validationErrorCode,
		Message: validationErrorMessage,
		Errors:  errorMessages,
	}
}

// Error returns the string representation of an ErrorResponse
func (e *ValidationError) Error() string {
	return "validation error: " + strings.Join(e.Errors, ", ")
}
