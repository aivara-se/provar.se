package feedback

import (
	"github.com/gofiber/fiber/v2"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/feedback"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupDeleteInvitation(app *fiber.App) {
	path := "/organization/:organizationId/feedback/:feedbackId"

	app.Delete(path, access.AuthenticatedGuard())
	app.Delete(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Delete(path, access.MembershipGuard())

	app.Delete(path, func(c *fiber.Ctx) error {
		organizationID := c.Params("organizationId")
		feedbackID := c.Params("feedbackId")
		err := feedback.DeleteByID(feedbackID, organizationID)
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
