package feedback

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
		logger := log.WithContext(c.Context())
		feedbackID := c.Params("feedbackId")
		err := feedback.DeleteByID(feedbackID)
		if err != nil {
			logger.Error("Failed to delete feedback", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete feedback")
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
