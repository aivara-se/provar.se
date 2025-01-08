package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/validator"
)

// LoginWithEmailConfirmRequestBody is the request body for logging in with email magic link.
// It contains the magic link token that the user received in their email.
type LoginWithEmailConfirmRequestBody struct {
	Token string `json:"token" validate:"required"`
}

// CreateLoginWithEmailConfirmRequestBody returns a new LoginWithEmailConfirmRequestBody
func CreateLoginWithEmailConfirmRequestBody() interface{} {
	return new(LoginWithEmailConfirmRequestBody)
}

// LoginWithEmailConfirmResponseBody is the response body sent when logging in with email
// magic link. It contains the access token that the client should use for further requests.
type LoginWithEmailConfirmResponseBody struct {
	AccessToken string `json:"accessToken" validate:"required"`
}

func SetupLoginWithEmailConfirm(app *fiber.App) {
	path := "/auth/email/confirm"

	app.Post(path, validator.ValidateMiddleware(CreateLoginWithEmailConfirmRequestBody))

	app.Post(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		body := validator.GetRequestBody(c).(*LoginWithEmailConfirmRequestBody)
		accessToken, err := access.ConfirmLoginWithEmail(body.Token)
		if err != nil {
			logger.Info("Failed to confirm login with email", err)
			return fiber.NewError(fiber.StatusForbidden, "Invalid email magic link token")
		}
		responseBody := LoginWithEmailConfirmResponseBody{AccessToken: accessToken}
		return c.JSON(responseBody)
	})
}
