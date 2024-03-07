package testapp

import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"provar.se/webapi/api"
	"provar.se/webapi/lib"
)

var (
	// cachedTestApp is a cached instance of the test app
	app *fiber.App
)

// Create creates a test app with a test database
func Create() *fiber.App {
	// Return cached app instance if one exists
	if app != nil {
		return app
	}

	// Load environment variables from .env.test
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)
	envfpath := filepath.Join(basepath, "../../.env.test")
	godotenv.Load(envfpath)

	// Load configuration from environment variables
	config := lib.ReadConfig()

	// Setup shared components (eg: db, jwt, etc.)
	if err := lib.Setup(config); err != nil {
		log.Fatal("Error setting up dependencies: ", err)
	}

	app = api.CreateApp()
	go startFileServer()

	return app
}
