package auth

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credentials"
	"provar.se/webapi/lib/token"
)

// CreateAccessTokenRequestBody is the request body for creating access tokens
type CreateAccessTokenRequestBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// CreateAccessTokenResponseBody is the response body for creating access tokens
type CreateAccessTokenResponseBody struct {
	Token string `json:"token"`
}

// @Router      /auth/token  [post]
// @Summary     Create an access token using client credentials.
// @Description Create an access token using client credentials.
// @Tags        auth
// @Accept      json
// @Param       x-organization  path      string  6541eba0b8857ce9f394cf7e  "Organization ID"
// @Param       client_id       body      string  taYxWHrU0YoKu5vtbE4jn    "Client ID"
// @Param       client_secret   body      string  taYxWHrU0YoKu5vtbE4jn    "Client Secret"
// @Success     200             {object}  auth.CreateAccessTokenResponseBody
func SetupCreateAccessToken(app *fiber.App) {
	app.Post("/auth/token", func(c *fiber.Ctx) error {
		organizationID := string(c.Request().Header.Peek("x-organization"))
		reqBody := &CreateAccessTokenRequestBody{}
		if err := c.BodyParser(&reqBody); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if organizationID == "" || reqBody.ClientID == "" || reqBody.ClientSecret == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if !credentials.ValidateClientCredentials(organizationID, reqBody.ClientID, reqBody.ClientSecret) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		token, err := token.CreateAccessToken(organizationID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		resBody := CreateAccessTokenResponseBody{Token: token}
		return c.JSON(resBody)
	})
}
