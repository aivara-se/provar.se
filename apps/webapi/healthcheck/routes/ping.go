package routes

import "github.com/gofiber/fiber/v2"

func SetupPing(app *fiber.App) {
	app.Get("/ping", handler)
}

func handler(c *fiber.Ctx) error {
	return c.SendString("pong")
}
