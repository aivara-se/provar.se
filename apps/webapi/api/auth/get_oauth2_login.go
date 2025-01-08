package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
)

func SetupOAuth2Login(app *fiber.App) {
	path := "/auth/oauth2/:provider/login"

	app.Get(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		provider := c.Params("provider")
		state := c.Query("state")
		if state == "" {
			return fiber.NewError(fiber.StatusBadRequest, "The state parameter is required")
		}
		redirectURL, err := access.GetOAuth2RedirectURL(provider, state, c)
		if err != nil {
			logger.Error("Failed to get OAuth2 redirect URL", err)
			return fiber.NewError(fiber.StatusInternalServerError, "OAuth2 login failed")
		}
		return c.Redirect(redirectURL, fiber.StatusTemporaryRedirect)
	})
}
