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

// CreateFeedbackRequestBody is the request body for creating an feedback
type CreateFeedbackRequestBody struct {
	FeedbackType feedback.FeedbackType `json:"feedbackType" validate:"required"`
	FeedbackTime time.Time             `json:"feedbackTime" validate:"required"`
	FeedbackData map[string]string     `json:"feedbackData" validate:"required"`
	FeedbackTags map[string]string     `json:"feedbackTags"`
	FeedbackMeta map[string]string     `json:"feedbackMeta"`
	FeedbackUser map[string]string     `json:"feedbackUser"`
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
		if body.FeedbackMeta == nil {
			body.FeedbackMeta = make(map[string]string)
		}
		if body.FeedbackTags == nil {
			body.FeedbackTags = make(map[string]string)
		}
		if body.FeedbackUser == nil {
			body.FeedbackUser = make(map[string]string)
		}
		feedback.SetRequestIP(&body.FeedbackMeta, c.IP())
		feedback.SetRequestHeaders(&body.FeedbackMeta, c.GetReqHeaders())
		feedback.SetMetadataFromIP(&body.FeedbackMeta)
		feedback.SetMetadataFromUA(&body.FeedbackMeta)
		_, err := feedback.Create(orgID, &feedback.CreateFeedbackData{
			FeedbackTime: body.FeedbackTime,
			FeedbackType: body.FeedbackType,
			FeedbackData: body.FeedbackData,
			FeedbackTags: body.FeedbackTags,
			FeedbackMeta: body.FeedbackMeta,
			FeedbackUser: body.FeedbackUser,
		})
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
