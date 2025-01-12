package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

// ErrorHandler is a custom error handler that logs errors on the console
// in detail and returns a simple error message to the client.
func ErrorHandler(c *fiber.Ctx, err error) error {
	e := FromError(err)
	l := log.WithContext(c.Context())

	// Log the original error in detail
	l.Errorw("Request failed", "error", e.WrappedError)

	// Send the error response to the client
	return c.Status(e.StatusCode).JSON(e.Response())
}
