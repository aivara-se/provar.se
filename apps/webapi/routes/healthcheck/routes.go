package healthcheck

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	SetupBasicHealthcheck(app)
}
