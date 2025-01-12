package auth

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
)

func SetupOAuth2Login(app *fiber.App) {
	path := "/auth/oauth2/:provider/login"

	app.Get(path, func(c *fiber.Ctx) error {
		provider := c.Params("provider")
		state := c.Query("state")
		if state == "" {
			return fiber.NewError(fiber.StatusBadRequest, "The state parameter is required")
		}
		redirectURL, err := access.GetOAuth2RedirectURL(provider, state, c)
		if err != nil {
			return err
		}
		return c.Redirect(redirectURL, fiber.StatusTemporaryRedirect)
	})
}
