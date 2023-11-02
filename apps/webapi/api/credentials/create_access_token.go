package credentials

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/creds"
	"provar.se/webapi/lib/token"
)

// CreateAccessTokenRequestBody is the request body for creating access tokens
type CreateAccessTokenRequestBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

// CreateAccessTokenResponseBody is the response body for creating access tokens
type CreateAccessTokenResponseBody struct {
	AccessToken string `json:"access_token"`
}

// @Summary			Creates access tokens
// @Description	Creates access tokens which can be used to authenticate when sending feedback.
// @Tags				credentials
// @Accept			json
// @Param       organization_id   path		string	6541eba0b8857ce9f394cf7e	"Organization ID"
// @Param       client_id   			header	string	taYxWHrU0YoKu5vtbE4jn			"Client ID"
// @Param       client_secret   	header	string	taYxWHrU0YoKu5vtbE4jn			"Client Secret"
// @Success 200 {object} 					credentials.CreateAccessTokenResponseBody
// @Router			/organization/{organization_id}/credential 	[post]
func SetupCreateAccessToken(app *fiber.App) {
	app.Post("/organization/:organization_id/credential", func(c *fiber.Ctx) error {
		organizationID := c.Params("organization_id")
		clientID := string(c.Request().Header.Peek("x-client-id"))
		clientSecret := string(c.Request().Header.Peek("x-client-secret"))
		if !creds.ValidateClientCredentials(organizationID, clientID, clientSecret) {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		token, err := token.CreateAccessToken(organizationID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		responseBody := CreateAccessTokenResponseBody{
			AccessToken: token,
		}
		return c.JSON(responseBody)
	})
}
