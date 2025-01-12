package exceptions

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ErrorResponse is the error response sent to the client.
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Error is an interface for errors occured in the application.
type Error struct {
	WrappedError error
	ErrorCode    string
	ErrorMessage string
	StatusCode   int
}

// NewError creates a new Error instance with given parameters.
func NewError(err error, status int, code string, message string) *Error {
	return &Error{
		WrappedError: err,
		ErrorCode:    code,
		ErrorMessage: message,
		StatusCode:   status,
	}
}

// FromError creates a new Error instance from an error type.
func FromError(err error) *Error {
	// Search for a fiber.Error in the chain
	var fe *fiber.Error
	if errors.As(err, &fe) {
		return FromFiberError(fe)
	}
	return FromUnknownError(err)
}

// FromFiberError creates a new Error instance from a fiber.Error instance.
func FromFiberError(err *fiber.Error) *Error {
	return NewError(err, err.Code, http.StatusText(err.Code), err.Message)
}

// FromUnknownError creates a new Error instance from an unknown error.
func FromUnknownError(err error) *Error {
	code := 500
	text := http.StatusText(code)
	return NewError(err, code, text, text)
}

// Error returns the error message that can be sent to the client.
func (e *Error) Error() string {
	return e.ErrorMessage
}

// Details returns the error details that can be sent to the client.
func (e *Error) Details() string {
	return e.Error()
}

// Unwrap returns the original (wrapped) error with internal details.
// This may contain sensitive information and should not be sent to the client.
func (e *Error) Unwrap() error {
	return e.WrappedError
}

// Response returns the error response that can be sent to the client.
func (e *Error) Response() *ErrorResponse {
	return &ErrorResponse{
		Code:    e.ErrorCode,
		Message: e.ErrorMessage,
	}
}
