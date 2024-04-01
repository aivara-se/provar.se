package ping

import (
	"github.com/gofiber/fiber/v2"
)

func SetupBasicHealthcheck(app *fiber.App) {
	path := "/ping"

	app.Get(path, func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
