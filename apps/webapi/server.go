package main

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"provar.se/webapi/routes"
)

func init() {
    err := godotenv.Load()
    if err != nil {
      log.Fatal("Error loading .env file")
    }
}

func main() {
    app := fiber.New(fiber.Config{
        JSONEncoder: json.Marshal,
        JSONDecoder: json.Unmarshal,
    })

    routes.SetupRoutes(app)

		port := os.Getenv("PORT")
    app.Listen(":" + port)
}
