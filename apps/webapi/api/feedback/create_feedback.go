package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/database/repository"
)

// @Router      /feedback  [post]
// @Summary     Create a new feedback event for an organization.
// @Description Create a new feedback event for an organization.
// @Tags        feedback
// @Accept      json
// @Success     204  "ok"
func SetupCreateFeedback(app *fiber.App) {
	app.Use("/feedback", credential.GetMiddleware())
	app.Post("/feedback", func(c *fiber.Ctx) error {
		organizationID := credential.GetOrganizationID(c)
		repo := repository.GetFeedbackRepository()
		err := repo.CreateFeedback(organizationID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
