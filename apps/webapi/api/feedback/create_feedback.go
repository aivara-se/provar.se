package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/credential"
	"provar.se/webapi/lib/feedback"
	"provar.se/webapi/lib/validator"
)

// CreateFeedbackRequestBody is the request body for creating a feedback event
type CreateFeedbackRequestBody struct {
	ProjectID string                `json:"projectId"`
	Type      feedback.FeedbackType `json:"type" validate:"required,oneof=text cnps csat"`
	Data      feedback.FeedbackData `json:"data" validate:"required"`
	Meta      feedback.FeedbackMeta `json:"meta"`
	User      feedback.FeedbackUser `json:"user"`
	Tags      feedback.FeedbackTags `json:"tags"`
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
		repo := feedback.GetFeedbackRepository()
		data := feedback.CreateFeedbackData{
			OrganizationID: organizationID,
			ProjectID:      body.ProjectID,
			Type:           body.Type,
			Data:           body.Data,
			Meta:           body.Meta,
			User:           body.User,
			Tags:           body.Tags,
		}
		if err := data.Validate(); err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if data.Meta == nil {
			data.Meta = feedback.FeedbackMeta{}
		}
		data.Meta.SetRequestIP(c.IP())
		data.Meta.SetRequestHeaders(c.GetReqHeaders())
		data.Meta.SetLocationFromIP()
		if err := repo.CreateFeedback(data); err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
