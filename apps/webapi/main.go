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

	// Setup shared components (eg: db, jwt, etc.)
	if err := lib.Setup(); err != nil {
		log.Fatal("Error connecting to database")
	}

	// Start the server
	cfg := config.Get()
	app := api.Create()
	if err := app.Listen(cfg.HostPort); err != nil {
		log.Fatal(err)
	}
}
