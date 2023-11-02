package feedback

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the routes for the feedback API
func SetupRoutes(app *fiber.App) {
	SetupCreateFeedback(app)
}
