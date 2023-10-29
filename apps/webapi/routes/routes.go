package routes

import (
	"github.com/gofiber/fiber/v2"

	"provar.se/webapi/routes/feedback"
	"provar.se/webapi/routes/healthcheck"
)

func SetupRoutes(app *fiber.App) {
	  feedback.SetupRoutes(app)
    healthcheck.SetupRoutes(app)
}
