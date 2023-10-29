package healthcheck

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/healthcheck/routes"
)

func SetupRoutes(app *fiber.App) {
	routes.SetupPing(app)
}
