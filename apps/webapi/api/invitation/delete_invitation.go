package invitation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"provar.se/webapi/lib/access"
	"provar.se/webapi/lib/invitation"
	"provar.se/webapi/lib/organization"
	"provar.se/webapi/lib/router"
)

func SetupDeleteInvitation(app *fiber.App) {
	path := "/organization/:organizationId/invitation/:invitationId"

	app.Delete(path, access.AuthenticatedGuard())
	app.Delete(path, organization.Loader(router.FromPathParam("organizationId")))
	app.Delete(path, access.MembershipGuard())

	app.Delete(path, func(c *fiber.Ctx) error {
		logger := log.WithContext(c.Context())
		inviteID := c.Params("invitationId")
		err := invitation.DeleteByID(inviteID)
		if err != nil {
			logger.Error("Failed to delete invitation", err)
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete invitation")
		}
		return c.SendStatus(fiber.StatusNoContent)
	})
}
