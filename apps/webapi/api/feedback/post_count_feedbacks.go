package feedback

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/feedback"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
	"provar.se/webapi/lib/validator"
)

// CountFeedbackResponseBody is the response body for searching feedbacks
type CountFeedbackResponseBody struct {
	Total int `json:"total"`
}

func SetupCountFeedback(app *fiber.App) {
	path := "/organization/:organizationId/feedbacks/count"

	app.Post(path, access.AuthenticatedGuard())
	app.Post(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Post(path, access.MembershipGuard())
	app.Post(path, validator.ValidateMiddleware(CreateSearchFeedbackRequestBody))

	app.Post(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		organizationID := c.Params("organizationId")
		body := validator.GetRequestBody(c).(*SearchFeedbackRequestBody)
		total, err := feedback.Count(organizationID, &feedback.SearchFeedbackData{
			PageLimit:    body.PageLimit,
			PageOffset:   body.PageOffset,
			BegTimestamp: body.BegTimestamp,
			EndTimestamp: body.EndTimestamp,
			FeedbackType: body.FeedbackType,
			FeedbackTags: body.FeedbackTags,
			FeedbackMeta: body.FeedbackMeta,
			FeedbackUser: body.FeedbackUser,
		})
		if err != nil {
			logger.Error("Failed to get feedback count", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to get feedback count")
		}
		return c.JSON(&CountFeedbackResponseBody{
			Total: total,
		})
	})
}
