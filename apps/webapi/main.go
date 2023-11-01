package main

import (
	"log"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	"provar.se/webapi/routes"
	"provar.se/webapi/shared/database"

	_ "provar.se/webapi/docs"
)

func init() {
	// Load environment variables from .env
	godotenv.Load()
}

// @title Provar API
// @version 1.0
// @description Provar.se public API
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
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
	// Serve swagger API UI
	app.Get("/swagger/*", swagger.HandlerDefault)
	// Load all app routes
	routes.SetupRoutes(app)
	// Listen on port $PORT
	PORT := os.Getenv("PORT")
	app.Listen(":" + PORT)
}
