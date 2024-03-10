package user

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
)

func SetupSecureHealthcheck(app *fiber.App) {
	app.Get("/user/:userID", access.EnsureUserLoggedInMiddleware())
	app.Get("/user/:userID", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
