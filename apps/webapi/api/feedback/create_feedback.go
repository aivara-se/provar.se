package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/database/repository"
	"provar.se/webapi/lib/validator"
)

// CreateFeedbackRequestBody is the request body for creating a feedback event
type CreateFeedbackRequestBody struct {
	ProjectID string                  `json:"projectId"`
	Type      repository.FeedbackType `json:"type" validate:"required,oneof=text cnps csat"`
	Data      repository.FeedbackData `json:"data" validate:"required"`
	Tags      repository.FeedbackTags `json:"tags"`
}

// NewCreateFeedbackRequestBody returns a new CreateFeedbackRequestBody
func NewCreateFeedbackRequestBody() interface{} {
	return new(CreateFeedbackRequestBody)
}

// @Router      /feedback  [post]
// @Summary     Create a new feedback event for an organization.
// @Description Create a new feedback event for an organization.
// @Tags        feedback
// @Accept      json
// @Param       body  body  CreateFeedbackRequestBody  true  "Request body"
// @Success     204  "ok"
func SetupCreateFeedback(app *fiber.App) {
	app.Post("/feedback", credential.GetMiddleware())
	app.Post("/feedback", validator.GetMiddleware(NewCreateFeedbackRequestBody))

	app.Post("/feedback", func(c *fiber.Ctx) error {
		organizationID := credential.GetOrganizationID(c)
		body := validator.GetRequestBody(c).(*CreateFeedbackRequestBody)
		repo := repository.GetFeedbackRepository()
		data := repository.CreateFeedbackData{
			OrganizationID: organizationID,
			ProjectID:      body.ProjectID,
			Type:           body.Type,
			Data:           body.Data,
			Tags:           body.Tags,
		}
		if err := data.Validate(); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if err := repo.CreateFeedback(data); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
