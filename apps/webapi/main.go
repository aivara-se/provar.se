package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"provar.se/webapi/api"
	"provar.se/webapi/lib/database"
	"provar.se/webapi/lib/token"

	_ "provar.se/webapi/docs"
)

func main() {
	// Load environment variables from .env
	godotenv.Load()

	// Connect to database using MONGO_URI
	MONGO_URI := os.Getenv("MONGO_URI")
	if err := database.Connect(MONGO_URI); err != nil {
		log.Fatal("Error connecting to database")
	}

	// Configure access token signing secret
	JWT_SECRET := os.Getenv("JWT_SECRET")
	token.SetSigningSecret(JWT_SECRET)

	// Load all app routes
	app := api.CreateApp()

	// Listen on port $PORT
	PORT := os.Getenv("PORT")
	if err := app.Listen(":" + PORT); err != nil {
		log.Fatal(err)
	}
}
