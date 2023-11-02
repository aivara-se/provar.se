package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the healthcheck routes
func SetupRoutes(app *fiber.App) {
	SetupBasicHealthcheck(app)
}
