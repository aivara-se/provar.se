package main

import (
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/feedback"
	"provar.se/webapi/healthcheck"
)

func main() {
    app := fiber.New(fiber.Config{
        JSONEncoder: json.Marshal,
        JSONDecoder: json.Unmarshal,
    })

    feedback.SetupRoutes(app)
    healthcheck.SetupRoutes(app)

		port := os.Getenv("PORT")
    app.Listen(":" + port)
}
