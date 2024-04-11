package feedback

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/feedback"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
	"provar.se/webapi/lib/validator"
)

const (
	// MaxFeedbacksPerBatch is the maximum number of feedbacks that can be
	// created in a single request. The request will be rejected if more than
	// this number of feedbacks are included in the request.
	MaxFeedbacksPerBatch = 10
)

// FeedbackToCreate is the data needed to create a feedback
type FeedbackToCreate struct {
	Type feedback.FeedbackType `json:"type" validate:"required"`
	Time time.Time             `json:"time" validate:"required"`
	Data map[string]string     `json:"data" validate:"required"`
	Tags map[string]string     `json:"tags"`
	Meta map[string]string     `json:"meta"`
	User map[string]string     `json:"user"`
}

// CreateFeedbackRequestBody is the request body for creating an feedback
type CreateFeedbackRequestBody struct {
	Feedbacks []FeedbackToCreate `json:"feedbacks" validate:"dive"`
}

// CreateCreateFeedbackRequestBody returns a new CreateFeedbackRequestBody
func CreateCreateFeedbackRequestBody() interface{} {
	return new(CreateFeedbackRequestBody)
}

func SetupCreateFeedback(app *fiber.App) {
	path := "/organization/:organizationId/feedback"

	app.Post(path, access.AuthenticatedGuard())
	app.Post(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Post(path, access.MembershipGuard())
	app.Post(path, validator.ValidateMiddleware(CreateCreateFeedbackRequestBody))

	app.Post(path, func(c *fiber.Ctx) error {
		orgID := c.Params("organizationId")
		body := validator.GetRequestBody(c).(*CreateFeedbackRequestBody)
		if body == nil || body.Feedbacks == nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		if len(body.Feedbacks) > MaxFeedbacksPerBatch {
			return c.SendStatus(fiber.StatusRequestEntityTooLarge)
		}
		errd := false
		for _, fb := range body.Feedbacks {
			if fb.Meta == nil {
				fb.Meta = make(map[string]string)
			}
			if fb.Tags == nil {
				fb.Tags = make(map[string]string)
			}
			if fb.User == nil {
				fb.User = make(map[string]string)
			}
			_, err := feedback.Create(orgID, &feedback.CreateFeedbackData{
				FeedbackTime:   fb.Time,
				FeedbackType:   fb.Type,
				FeedbackData:   fb.Data,
				FeedbackTags:   fb.Tags,
				FeedbackMeta:   fb.Meta,
				FeedbackUser:   fb.User,
				RequestIP:      c.IP(),
				RequestHeaders: c.GetReqHeaders(),
			})
			if err != nil {
				errd = true
			}
		}
		if errd {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
