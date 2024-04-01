package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/feedback"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupFeedbackDetails(app *fiber.App) {
	path := "/organization/:organizationId/feedback/:feedbackId"

	app.Get(path, access.AuthenticatedGuard())
	app.Get(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Get(path, access.MembershipGuard())
	app.Get(path, feedback.Loader(router.FromPathParam("feedbackId")))

	app.Get(path, func(c *fiber.Ctx) error {
		org := organization.GetOrganization(c)
		fb := feedback.GetFeedback(c)
		if fb.OrganizationID != org.ID {
			return c.SendStatus(fiber.StatusNotFound)
		}
		return c.JSON(fb)
	})
}
