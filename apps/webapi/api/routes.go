package api

import (
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"provar.se/webapi/api/feedback"
	"provar.se/webapi/api/health"
)

// @title       Provar API
// @version		  1.0
// @description	Provar.se API
// @host			  https://api.provar.se
// @BasePath		/
// @accept			json
// @produce     json
func CreateApp() *fiber.App {
	// Create a new fiber application
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

	// Use cors middleware
	app.Use(cors.New())

	// Use logger middleware
	app.Use(logger.New())

	// Serve swagger API UI
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Load all app routes
	feedback.SetupCreateFeedback(app)
	feedback.SetupImportFeedback(app)
	health.SetupBasicHealthcheck(app)
	health.SetupSecureHealthcheck(app)

	return app
}
