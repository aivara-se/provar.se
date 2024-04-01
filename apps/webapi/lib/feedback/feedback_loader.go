package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/router"
)

var (
	// feedbackKey is the key used to store the loaded feedback
	feedbackKey = "app:feedback"
)

// Loader loads an feedback from the database and attaches it to the context
func Loader(getID router.ParamFetcher) fiber.Handler {
	return func(c *fiber.Ctx) error {
		invite, err := FindByID(getID(c))
		if err != nil {
			return c.SendStatus(fiber.StatusNotFound)
		}
		c.Locals(feedbackKey, invite)
		return c.Next()
	}
}

// GetFeedback returns the loaded feedback from the context
func GetFeedback(c *fiber.Ctx) *Feedback {
	return c.Locals(feedbackKey).(*Feedback)
}
