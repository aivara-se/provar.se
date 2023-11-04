package health

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/token"
)

// @Router      /ping/secure  [get]
// @Summary     Ensure that an access token is valid.
// @Description Ensure that an access token is valid.
// @Tags        health
// @Success     204  "ok"
func SetupSecureHealthcheck(app *fiber.App) {
	app.Use("/ping/secure", token.GetMiddleware())
	app.Get("/ping/secure", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
