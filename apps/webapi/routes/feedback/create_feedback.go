package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/shared/database"
)

// SetupCreateFeedback sets up the route for creating a feedback.
// POST /organization/:organizationID/feedback
// Request body: tbd
func SetupCreateFeedback(app *fiber.App) {
	app.Post("/organization/:organizationID/feedback", func(c *fiber.Ctx) error {
		repo := database.GetFeedbackRepository()
		repo.CreateFeedback(&database.CreateFeedbackParams{
			OrganizationID: c.Params("organizationID"),
		})
		return c.SendStatus(fiber.StatusNoContent)
	})
}
