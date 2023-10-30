package main

import (
	"log"
	"net/http"
	"os"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/joho/godotenv"
	"provar.se/webapi/routes"
	"provar.se/webapi/shared/database"
)

func main() {
	// Load environment variables from .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
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
	app.Use(filesystem.New(filesystem.Config{
		Root: http.Dir("./public"),
	}))
	// Load all app routes
	routes.SetupRoutes(app)
	// Listen on port $PORT
	PORT := os.Getenv("PORT")
	app.Listen(":" + PORT)
}
