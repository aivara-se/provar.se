package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/database/repository"
	"provar.se/webapi/lib/token"
)

// @Summary			Creates a new feedback
// @Description	Creates a new feedback for the organization.
// @Tags				feedbacks
// @Accept			json
// @Param       organization_id   path	string	6541eba0b8857ce9f394cf7e	"Organization ID"
// @Success			204			"ok"
// @Router			/organization/{organization_id}/feedback 	[post]
func SetupCreateFeedback(app *fiber.App) {
	app.Use("/organization/:organization_id/feedback", token.GetMiddleware())
	app.Post("/organization/:organization_id/feedback", func(c *fiber.Ctx) error {
		organizationID := c.Params("organization_id")
		repo := repository.GetFeedbackRepository()
		err := repo.CreateFeedback(organizationID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
