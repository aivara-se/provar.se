package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

var (
	// feedbackKey is the key used to store the loaded feedback
	feedbackKey = "app:feedback"
)

// Loader loads an feedback from the database and attaches it to the context
func Loader(getID router.ParamFetcher) fiber.Handler {
	return func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		invite, err := FindByID(getID(c), org.ID)
		if err != nil {
			return NewNotFoundError(err)
		}
		c.Locals(feedbackKey, invite)
		return c.Next()
	}
}

// GetFeedback returns the loaded feedback from the context
func GetFeedback(c *fiber.Ctx) *Feedback {
	fb := c.Locals(feedbackKey).(*Feedback)
	if fb == nil {
		panic("Feedback is not available in the request context")
	}
	return fb
}
