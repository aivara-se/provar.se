package api

import (
	"github.com/gofiber/fiber/v2"

	"provar.se/webapi/api/credentials"
	"provar.se/webapi/api/feedback"
	"provar.se/webapi/api/healthcheck"
)

func SetupRoutes(app *fiber.App) {
	credentials.SetupRoutes(app)
	feedback.SetupRoutes(app)
	healthcheck.SetupRoutes(app)
}
