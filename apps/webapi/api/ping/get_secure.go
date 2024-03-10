package ping

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
)

func SetupSecureHealthcheck(app *fiber.App) {
	app.Get("/ping/secure", access.AuthenticatedGuard())
	app.Get("/ping/secure", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
