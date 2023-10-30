package healthcheck

import "github.com/gofiber/fiber/v2"

// SetupBasicHealthcheck sets up the route for a basic healthcheck
// GET /ping
func SetupBasicHealthcheck(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})
}
