package credentials

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes sets up the credential routes
func SetupRoutes(app *fiber.App) {
	SetupCreateAccessToken(app)
}
