package main

import (
	"log"

	"github.com/joho/godotenv"
	"provar.se/webapi/api"
	"provar.se/webapi/lib"
	"provar.se/webapi/lib/config"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()

	// Load configuration from environment variables
	cfg := config.FromEnv()

	// Setup shared components (eg: db, jwt, etc.)
	if err := lib.Setup(cfg); err != nil {
		log.Fatal("Error connecting to database")
	}

	// Start the server
	app := api.Create()
	if err := app.Listen(cfg.HostPort); err != nil {
		log.Fatal(err)
	}
}
