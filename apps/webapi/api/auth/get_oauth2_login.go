package auth

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
)

func SetupOAuth2Login(app *fiber.App) {
	path := "/auth/oauth2/:provider/login"

	app.Get(path, func(c *fiber.Ctx) error {
		providerName := c.Params("provider")
		state := c.Query("state")
		if state == "" {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		redirectURL, err := access.GetOAuth2RedirectURL(providerName, state, c)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.Redirect(redirectURL, fiber.StatusTemporaryRedirect)
	})
}
