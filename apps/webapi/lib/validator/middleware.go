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

// ValidateMiddleware returns a fiber middleware to validate credentials
func ValidateMiddleware(factory BodyFactory) fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := factory()
		if err := c.BodyParser(body); err != nil {
			return NewParseFailedError(err)
		}
		if err := validate.Struct(body); err != nil {
			validationErrors, ok := err.(validator.ValidationErrors)
			if !ok {
				return NewValidateFailedError(err)
			}
			validationError := newValidationError(validationErrors)
			return NewValidateFailedError(err, validationError.Error())
		}
		c.Locals(requestBodyKey, body)
		return c.Next()
	}
}

// GetRequestBody returns the organization id from the fiber context
func GetRequestBody(c *fiber.Ctx) interface{} {
	return c.Locals(requestBodyKey)
}
