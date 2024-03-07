package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	// requestBodyKey is the key used to store the request body in locals
	requestBodyKey = "app:requestBody"

	// validate is the validator instance used to validate requests
	validate = validator.New(validator.WithRequiredStructEnabled())
)

// BodyFactory is a function that returns a new instance of a struct
// that will be used to parse and validate the request body.
type BodyFactory func() interface{}

// GetMiddleware returns a fiber middleware to validate credentials
func GetMiddleware(factory BodyFactory) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := factory()
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := validate.Struct(body); err != nil {
			validationErrors, ok := err.(validator.ValidationErrors)
			if !ok {
				return c.SendStatus(fiber.StatusBadRequest)
			}
			responseBody := newValidationError(validationErrors)
			return c.Status(fiber.StatusBadRequest).JSON(responseBody)
		}
		c.Locals(requestBodyKey, body)
		return c.Next()
	}
}

// GetRequestBody returns the organization id from the fiber context
func GetRequestBody(c *fiber.Ctx) interface{} {
	return c.Locals(requestBodyKey)
}
