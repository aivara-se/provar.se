package validator

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/exceptions"
)

// NewParseFailedError creates a new error for failed request body parsing
func NewParseFailedError(err error) error {
	return exceptions.NewError(err, fiber.StatusBadRequest, "Failed to parse request body")
}

// NewValidateFailedError creates a new error for failed request body validation
func NewValidateFailedError(err error, details ...string) error {
	return exceptions.NewError(err, fiber.StatusBadRequest, "Failed to validate request body", details...)
}
