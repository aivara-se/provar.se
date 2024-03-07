package health

import (
	"github.com/gofiber/fiber/v2"
)

func SetupBasicHealthcheck(app *fiber.App) {
	app.Get("/health/basic", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
