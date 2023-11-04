package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/database/repository"
	"provar.se/webapi/lib/token"
)

// @Router      /feedback  [post]
// @Summary     Create a new feedback event for an organization.
// @Description Create a new feedback event for an organization.
// @Tags        feedback
// @Accept      json
// @Param       x-organization  path  string  6541eba0b8857ce9f394cf7e  "Organization ID"
// @Success     204  "ok"
func SetupCreateFeedback(app *fiber.App) {
	app.Use("/feedback", token.GetMiddleware())
	app.Post("/feedback", func(c *fiber.Ctx) error {
		organizationID := token.ExtractOrganizationId(c)
		repo := repository.GetFeedbackRepository()
		err := repo.CreateFeedback(organizationID)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
