package auth

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/validator"
)

// LoginWithEmailRequestBody is the request body for logging in with email magic link
type LoginWithEmailRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}

// CreateLoginWithEmailRequestBody returns a new LoginWithEmailRequestBody
func CreateLoginWithEmailRequestBody() interface{} {
	return new(LoginWithEmailRequestBody)
}

func SetupLoginWithEmail(app *fiber.App) {
	path := "/auth/email/prepare"

	app.Post(path, validator.ValidateMiddleware(CreateLoginWithEmailRequestBody))

	app.Post(path, func(c *fiber.Ctx) error {
		body := validator.GetRequestBody(c).(*LoginWithEmailRequestBody)
		access.PrepareLoginWithEmail(body.Email)
		return c.SendStatus(fiber.StatusNoContent)
	})
}
