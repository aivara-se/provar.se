package ping

import (
	"github.com/gofiber/fiber/v2"
)

func SetupBasicHealthcheck(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
