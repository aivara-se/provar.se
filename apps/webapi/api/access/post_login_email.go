package access

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/validator"
)

// LoginWithEmailRequestBody is the request body for logging in with email magic link
type LoginWithEmailRequestBody struct {
	Email string `json:"email" validate:"required,email"`
}

// NewLoginWithEmailRequestBody returns a new LoginWithEmailRequestBody
func NewLoginWithEmailRequestBody() interface{} {
	return new(LoginWithEmailRequestBody)
}

func SetupLoginWithEmail(app *fiber.App) {
	app.Post("/access/login/email", validator.Middleware(NewLoginWithEmailRequestBody))
	app.Post("/access/login/email", func(c *fiber.Ctx) error {
		body := validator.GetRequestBody(c).(*LoginWithEmailRequestBody)
		access.PrepareLoginWithEmail(body.Email)
		return c.SendStatus(fiber.StatusNoContent)
	})
}
