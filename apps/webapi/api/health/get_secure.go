package health

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
)

func SetupSecureHealthcheck(app *fiber.App) {
	app.Get("/health/secure", access.AuthenticatedGuard())
	app.Get("/health/secure", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
