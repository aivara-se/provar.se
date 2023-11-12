package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/database/repository"
)

// RequestBody is the request body for creating a feedback event
type RequestBody struct {
	Type repository.FeedbackType `json:"type"`
	Data repository.FeedbackData `json:"data"`
}

// @Router      /feedback  [post]
// @Summary     Create a new feedback event for an organization.
// @Description Create a new feedback event for an organization.
// @Tags        feedback
// @Accept      json
// @Param       body  body  RequestBody  true  "The request body"
// @Success     204  "ok"
func SetupCreateFeedback(app *fiber.App) {
	app.Use("/feedback", credential.GetMiddleware())
	app.Post("/feedback", func(c *fiber.Ctx) error {
		organizationID := credential.GetOrganizationID(c)
		body := new(RequestBody)
		if err := c.BodyParser(body); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		repo := repository.GetFeedbackRepository()
		err := repo.CreateFeedback(organizationID, body.Type, body.Data)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
