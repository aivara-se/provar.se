package healthcheck

import "github.com/gofiber/fiber/v2"

// @Summary			Basic healthcheck endpoint
// @Description	Basic healthcheck endpoint that can be used to check whether the service can be reached.
// @Tags				health
// @Success			204			"ok"
// @Router			/ping 	[get]
func SetupBasicHealthcheck(app *fiber.App) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})
}
