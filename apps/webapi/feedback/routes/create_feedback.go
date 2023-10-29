package routes

import "github.com/gofiber/fiber/v2"

func SetupCreateFeedback(app *fiber.App) {
	app.Post("/feedback", handler)
}

func handler(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}
