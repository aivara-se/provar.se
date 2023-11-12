package health

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
)

// @Router      /ping/secure  [get]
// @Summary     Ensure that the api key is valid
// @Description Ensure that the api key is valid
// @Tags        health
// @Success     204  "ok"
func SetupSecureHealthcheck(app *fiber.App) {
	app.Get("/ping/secure", credential.GetMiddleware())
	app.Get("/ping/secure", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
