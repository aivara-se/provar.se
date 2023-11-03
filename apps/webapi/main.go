package main

import (
	"log"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"provar.se/webapi/api"
	"provar.se/webapi/lib/database"

	_ "provar.se/webapi/docs"
)

func init() {
	// Load environment variables from .env
	godotenv.Load()
}

// @title Provar API
// @version 1.0
// @description Provar.se API
// @host https://api.provar.se
// @BasePath /
// @accept json
// @produce json
func main() {
	// Connect to database using MONGO_URI
	MONGO_URI := os.Getenv("MONGO_URI")
	if err := database.Connect(MONGO_URI); err != nil {
		log.Fatal("Error connecting to database")
	}

	// Configure Fiber app with faster JSON
	app := fiber.New(fiber.Config{
		AppName:         "Provar.se API v1.0",
		ReadTimeout:     time.Second,
		WriteTimeout:    time.Second,
		WriteBufferSize: 2048,
		BodyLimit:       10 * 1024,
		JSONEncoder:     json.Marshal,
		JSONDecoder:     json.Unmarshal,
	})

	// Use rate limiter middleware
	app.Use(limiter.New(limiter.Config{
		Expiration: time.Second,
		Max:        100,
	}))

	// Use logger middleware
	app.Use(logger.New())

	// Serve swagger API UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Load all app routes
	api.SetupRoutes(app)

	// Listen on port $PORT
	PORT := os.Getenv("PORT")
	if err := app.Listen(":" + PORT); err != nil {
		log.Fatal(err)
	}
}
