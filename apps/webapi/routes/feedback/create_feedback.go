package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/shared/database"
)

// @Summary			Creates a new feedback
// @Description	Creates a new feedback for the organization.
// @Tags				feedback
// @Accept			json
// @Param       organization_id   path	string	6541eba0b8857ce9f394cf7e	"Organization ID"
// @Success			204			"ok"
// @Router			/organization/{organization_id}/feedback 	[post]
func SetupCreateFeedback(app *fiber.App) {
	app.Post("/organization/:organization_id/feedback", func(c *fiber.Ctx) error {
		repo := database.GetFeedbackRepository()
		repo.CreateFeedback(&database.CreateFeedbackParams{
			OrganizationID: c.Params("organization_id"),
		})
		return c.SendStatus(fiber.StatusNoContent)
	})
}
