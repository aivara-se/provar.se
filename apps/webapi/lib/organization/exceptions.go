package organization

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/exceptions"
)

// NewNotFoundError creates a new error for organization not found
func NewNotFoundError(err error) error {
	return exceptions.NewError(err, fiber.StatusNotFound, "Organization not found")
}
