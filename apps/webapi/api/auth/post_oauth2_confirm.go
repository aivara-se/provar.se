package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/validator"
)

// OAuth2ConfirmRequestBody is the request body for logging in with oauth2.
type OAuth2ConfirmRequestBody struct {
	State string `json:"state" validate:"required"`
	Code  string `json:"code"`
}

// CreateOAuth2ConfirmRequestBody returns a new OAuth2ConfirmRequestBody
func CreateOAuth2ConfirmRequestBody() interface{} {
	return new(OAuth2ConfirmRequestBody)
}

// OAuth2ConfirmResponseBody is the response body sent when logging in with
// OAuth2. It contains the access token that the client should use for further
// requests.
type OAuth2ConfirmResponseBody struct {
	AccessToken string `json:"accessToken" validate:"required"`
}

func SetupOAuth2Callback(app *fiber.App) {
	path := "/auth/oauth2/:provider/confirm"

	app.Post(path, validator.ValidateMiddleware(CreateOAuth2ConfirmRequestBody))

	app.Post(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		provider := c.Params("provider")
		body := validator.GetRequestBody(c).(*OAuth2ConfirmRequestBody)
		token, err := access.CompleteOAuth2Flow(provider, body.State, body.Code, c)
		if err != nil {
			logger.Error("Failed to complete OAuth2 flow", err)
			return fiber.NewError(fiber.StatusInternalServerError, "OAuth2 login failed")
		}
		return c.JSON(&OAuth2ConfirmResponseBody{
			AccessToken: token,
		})
	})
}
