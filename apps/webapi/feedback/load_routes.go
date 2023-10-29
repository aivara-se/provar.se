package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/feedback/routes"
)

func SetupRoutes(app *fiber.App) {
	routes.SetupCreateFeedback(app)
}
