package main

import (
	"log"

	"github.com/joho/godotenv"
	"provar.se/webapi/api"
	"provar.se/webapi/lib"
)

func main() {
	// Load environment variables from .env
	godotenv.Load()

	// Load configuration from environment variables
	config := lib.ReadConfig()

	// Setup shared components (eg: db, jwt, etc.)
	if err := lib.Setup(config); err != nil {
		log.Fatal("Error connecting to database")
	}

	// Start the server
	app := api.CreateApp()
	if err := app.Listen(config.Address); err != nil {
		log.Fatal(err)
	}
}
