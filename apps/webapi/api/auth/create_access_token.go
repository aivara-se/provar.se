package auth

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/database/repository"
	"provar.se/webapi/lib/token"
)

// CreateAccessTokenRequestBody is the request body for creating access tokens
type CreateAccessTokenRequestBody struct {
	Key string `json:"key"`
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
// @Param       key             body      string  taYxWHrU0YoKu5vtbE4jn     "Key"
// @Success     200             {object}  auth.CreateAccessTokenResponseBody
func SetupCreateAccessToken(app *fiber.App) {
	app.Post("/auth/token", func(c *fiber.Ctx) error {
		organizationID := string(c.Request().Header.Peek("x-organization"))
		reqBody := &CreateAccessTokenRequestBody{}
		if err := c.BodyParser(&reqBody); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if organizationID == "" || reqBody.Key == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		repo := repository.GetCredentialRepository()
		if err := repo.IsValidCredenial(organizationID, reqBody.Key); err != nil {
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
