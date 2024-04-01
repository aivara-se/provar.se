package ping

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
)

func SetupSecureHealthcheck(app *fiber.App) {
	path := "/ping/secure"

	app.Get(path, access.AuthenticatedGuard())

	app.Get(path, func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
