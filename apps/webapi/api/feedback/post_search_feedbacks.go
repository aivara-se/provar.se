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

// SearchFeedbackRequestBody is the request body for searching feedbacks
type SearchFeedbackRequestBody struct {
	PageLimit    int                     `json:"pageLimit" validate:"gte=0,lte=1000"`
	PageOffset   int                     `json:"pageOffset" validate:"gte=0"`
	BegTimestamp time.Time               `json:"begTimestamp"`
	EndTimestamp time.Time               `json:"endTimestamp"`
	FeedbackType []feedback.FeedbackType `json:"feedbackType"`
	FeedbackTags map[string]string       `json:"feedbackTags"`
	FeedbackMeta map[string]string       `json:"feedbackMeta"`
	FeedbackUser map[string]string       `json:"feedbackUser"`
}

// SearchFeedbackResponseBody is the response body for searching feedbacks
type SearchFeedbackResponseBody struct {
	Total     int                  `json:"total"`
	Feedbacks []*feedback.Feedback `json:"feedbacks"`
}

// CreateSearchFeedbackRequestBody returns a new SearchFeedbackRequestBody
func CreateSearchFeedbackRequestBody() interface{} {
	return new(SearchFeedbackRequestBody)
}

func SetupSearchFeedback(app *fiber.App) {
	path := "/organization/:organizationId/feedbacks"

	app.Post(path, access.AuthenticatedGuard())
	app.Post(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Post(path, access.MembershipGuard())
	app.Post(path, validator.ValidateMiddleware(CreateSearchFeedbackRequestBody))

	app.Post(path, func(c *fiber.Ctx) error {
		orgID := c.Params("organizationId")
		body := validator.GetRequestBody(c).(*SearchFeedbackRequestBody)
		fbs, total, err := feedback.Search(orgID, &feedback.SearchFeedbackData{
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
			return c.SendStatus(fiber.StatusInternalServerError)
		}
		return c.JSON(&SearchFeedbackResponseBody{
			Total:     total,
			Feedbacks: fbs,
		})
	})
}
