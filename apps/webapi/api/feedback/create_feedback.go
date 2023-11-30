package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/database/repository"
	"provar.se/webapi/lib/validator"
)

// RequestBody is the request body for creating a feedback event
type RequestBody struct {
	ProjectID string                  `json:"projectId"`
	Type      repository.FeedbackType `json:"type" validate:"required,oneof=text cnps csat"`
	Data      repository.FeedbackData `json:"data" validate:"required"`
	Tags      repository.FeedbackTags `json:"tags"`
}

// NewRequestBody returns a new RequestBody for parsing and validating
func NewRequestBody() interface{} {
	return new(RequestBody)
}

// @Router      /feedback  [post]
// @Summary     Create a new feedback event for an organization.
// @Description Create a new feedback event for an organization.
// @Tags        feedback
// @Accept      json
// @Param       body  body  RequestBody  true  "The request body"
// @Success     204  "ok"
func SetupCreateFeedback(app *fiber.App) {
	app.Post("/feedback", credential.GetMiddleware())
	app.Post("/feedback", validator.GetMiddleware(NewRequestBody))

	app.Post("/feedback", func(c *fiber.Ctx) error {
		organizationID := credential.GetOrganizationID(c)
		body := validator.GetRequestBody(c).(*RequestBody)
		repo := repository.GetFeedbackRepository()
		err := repo.CreateFeedback(organizationID, body.ProjectID, body.Type, body.Data, body.Tags)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
