package invitation

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/exceptions"
)

// NewNotFoundError creates a new error for when an invitation is not found
func NewNotFoundError(err error) error {
	return exceptions.NewError(err, fiber.StatusNotFound, "Invitation not found")
}
