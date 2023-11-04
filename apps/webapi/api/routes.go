package api

import (
	"github.com/gofiber/fiber/v2"

	"provar.se/webapi/api/auth"
	"provar.se/webapi/api/feedback"
	"provar.se/webapi/api/health"
)

func SetupRoutes(app *fiber.App) {
	auth.SetupRoutes(app)
	feedback.SetupRoutes(app)
	health.SetupRoutes(app)
}
