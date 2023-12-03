package testapp

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

const FeedbackCSV = `type,time,cnps,csat,text
text,2023-01-21T16:34:59.512Z,,,The experience was average. There's room for enhancement.
text,2023-05-16T15:39:02.789Z,,,"The platform is reliable, but it lacks some features."
text,2023-03-06T02:31:53.609Z,,,"The service is decent, but there's room for enhancement."`

var (
	// IsRunning is true if the test app is running
	isRunning = false
)

// startFileServer starts a server with csv files
func startFileServer() {
	if isRunning {
		return
	}
	isRunning = true

	app := fiber.New(fiber.Config{})

	app.Use("/importable/valid", func(c *fiber.Ctx) error {
		c.SendStatus(fiber.StatusOK)
		c.SendString(FeedbackCSV)
		return nil
	})

	if err := app.Listen(":9001"); err != nil {
		log.Fatal(err)
	}
}
