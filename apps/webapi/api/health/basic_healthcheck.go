package health

import (
	"github.com/gofiber/fiber/v2"
)

// @Router      /ping  [get]
// @Summary     Ensure that the service can be reached.
// @Description Ensure that the service can be reached.
// @Tags        health
// @Success     204  "ok"
func SetupBasicHealthcheck(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
