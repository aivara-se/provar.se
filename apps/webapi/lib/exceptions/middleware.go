package exceptions

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/logger"
)

// ErrorHandler is a custom error handler that logs errors on the console
// in detail and returns a simple error message to the client.
func ErrorHandler(c *fiber.Ctx, err error) error {
	e := FromError(err)

	// Log the error message and the wrapped error in detail
	logger.Get().Error("Request failed: "+e.ErrorMessage, "error", e.WrappedError)

	// Send the error response to the client
	return c.Status(e.StatusCode).JSON(e.Response())
}
